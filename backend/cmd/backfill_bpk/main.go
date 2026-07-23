package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/bpk"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/peraturan"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// Backfill missing PDFs from BPK.
// Queries Firestore for laws that were never parsed, searches BPK for each,
// downloads PDF, parses, saves to Firestore.
//
// Usage:
//   go run ./backend/cmd/backfill_bpk -workers=4
//   go run ./backend/cmd/backfill_bpk -workers=4 -dry-run

func main() {
	var (
		workers int
		dryRun  bool
		verbose bool
	)
	flag.IntVar(&workers, "workers", 4, "parallel download workers")
	flag.BoolVar(&dryRun, "dry-run", false, "search BPK only, no download/parse/save")
	flag.BoolVar(&verbose, "verbose", false, "debug logging")
	flag.Parse()

	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("config load failed", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()

	// Firestore
	repo, err := repository.NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath)
	if err != nil {
		logger.Error("firestore init failed", "error", err)
		os.Exit(1)
	}
	defer repo.Close()

	// BPK connector
	bpkConn := bpk.New(logger)

	// Parser
	p := parser.New(logger)

	// Step 1: Get all laws from Firestore to know what we already have
	logger.Info("fetching all laws from Firestore")
	allDocs, err := repo.ListAllLaws(ctx)
	if err != nil {
		logger.Error("list laws failed", "error", err)
		os.Exit(1)
	}

	// Build set of law numbers we already have parsed
	haveParsed := make(map[string]bool)
	for _, d := range allDocs {
		if d.Status == "parsed" || d.Status == "analyzed" || d.Status == "pending_approval" {
			haveParsed[d.LawNumber] = true
		}
	}

	logger.Info("laws in Firestore", "total", len(allDocs), "already_parsed", len(haveParsed))

	// Step 2: Scrape full listing from peraturan.go.id to get all active laws
	scr := scraper.New(cfg.Scraper.PythonPath, cfg.Scraper.ScriptPath, logger)
	peraturanConn := peraturan.New(scr, logger)
	logger.Info("scraping peraturan.go.id for all active laws")
	allListed, err := peraturanConn.CheckUpdates(ctx)
	if err != nil {
		logger.Error("scrape peraturan failed", "error", err)
		os.Exit(1)
	}
	logger.Info("scraped from peraturan.go.id", "total", len(allListed))

	// Step 3: Find laws that are missing (listed but not parsed)
	var missing []connectors.DocumentMeta
	for _, m := range allListed {
		if !haveParsed[m.LawNumber] {
			missing = append(missing, m)
		}
	}

	logger.Info("missing laws to backfill from BPK", "count", len(missing))

	if len(missing) == 0 {
		logger.Info("nothing to backfill — all laws already parsed")
		return
	}

	if dryRun {
		for i, m := range missing {
			if i >= 20 {
				fmt.Printf("... and %d more\n", len(missing)-20)
				break
			}
			fmt.Printf("%s\t%s\t%s\n", m.LawNumber, m.DocumentType, m.Title)
		}
		return
	}

	// Step 4: Search BPK for each missing law, download + parse + save
	storageDir := "backend/internal/storage"
	os.MkdirAll(storageDir, 0755)

	var (
		found    int64
		notFound int64
		failed   int64
		wg       sync.WaitGroup
		sem      = make(chan struct{}, workers)
	)

	for _, m := range missing {
		wg.Add(1)
		go func(meta connectors.DocumentMeta) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			// Search BPK
			bpkMeta, err := bpkConn.SearchByLawNumber(ctx, meta.LawNumber, meta.DocumentType)
			if err != nil {
				logger.Warn("bpk search error", "law_number", meta.LawNumber, "error", err)
				atomic.AddInt64(&failed, 1)
				return
			}
			if bpkMeta == nil {
				atomic.AddInt64(&notFound, 1)
				return
			}

			// Download from BPK
			raw, err := bpkConn.Download(ctx, *bpkMeta)
			if err != nil {
				logger.Warn("bpk download failed", "law_number", meta.LawNumber, "error", err)
				atomic.AddInt64(&failed, 1)
				return
			}
			defer raw.Content.Close()

			// Save LawDocument
			doc := &models.LawDocument{
				LawNumber:     meta.LawNumber,
				Title:         meta.Title,
				SourceURL:     bpkMeta.SourceURL,
				Source:        bpkConn.Name(),
				Level:         meta.Level,
				DocumentType:  meta.DocumentType,
				Status:        "downloaded",
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
			docID, err := repo.SaveLawDocument(ctx, doc)
			if err != nil {
				logger.Error("save doc failed", "law_number", meta.LawNumber, "error", err)
				atomic.AddInt64(&failed, 1)
				return
			}
			doc.ID = docID

			// Parse
			result, err := p.Parse(ctx, raw.Content, raw.MimeType, raw.Filename)
			raw.Content.Close()
			if err != nil {
				doc.Status = "parse_failed"
				doc.UpdatedAt = time.Now()
				repo.SaveLawDocument(ctx, doc)
				logger.Warn("parse failed", "law_number", meta.LawNumber, "error", err)
				atomic.AddInt64(&failed, 1)
				return
			}

			// Truncate to Firestore limit (1,048,487 bytes)
			text := result.TextContent
			if len(text) > 1048000 {
				text = text[:1048000]
				logger.Warn("text truncated to firestore limit", "law_number", meta.LawNumber, "original_chars", len(result.TextContent))
			}

			// Save version
			version := &models.LawVersion{
				LawDocumentID: doc.ID,
				VersionNumber: int(time.Now().Unix()),
				TextContent:   text,
				ParsedAt:      time.Now(),
			}
			if _, err := repo.SaveLawVersion(ctx, doc.ID, version); err != nil {
				logger.Error("save version failed", "law_number", meta.LawNumber, "error", err)
				atomic.AddInt64(&failed, 1)
				return
			}

			doc.Status = "parsed"
			doc.UpdatedAt = time.Now()
			repo.SaveLawDocument(ctx, doc)

			atomic.AddInt64(&found, 1)
			f := atomic.LoadInt64(&found)
			nf := atomic.LoadInt64(&notFound)
			logger.Info("backfilled from BPK", "law_number", meta.LawNumber, "found", f, "not_found", nf, "chars", len(text))

			// Force GC to release PDF text memory before next law
			runtime.GC()
		}(m)
	}

	wg.Wait()

	fmt.Println("\n=== BACKFILL COMPLETE ===")
	fmt.Printf("Total missing:  %d\n", len(missing))
	fmt.Printf("Found on BPK:   %d\n", atomic.LoadInt64(&found))
	fmt.Printf("Not found:      %d\n", atomic.LoadInt64(&notFound))
	fmt.Printf("Failed:         %d\n", atomic.LoadInt64(&failed))
}
