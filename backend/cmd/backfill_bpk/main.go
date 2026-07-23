package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
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

// Backfill missing PDFs from BPK (peraturan.bpk.go.id).
// Scrapes peraturan.go.id for all active laws, checks which are missing from
// Firestore, searches BPK for each, downloads + parses + saves.
// Uses local cache file to avoid burning Firestore read quota on restarts.
//
// Usage:
//   go run ./backend/cmd/backfill_bpk -workers=1
//   go run ./backend/cmd/backfill_bpk -dry-run

const parsedCacheFile = "backend/configs/parsed_cache.txt"

func main() {
	var (
		workers int
		dryRun  bool
		verbose bool
	)
	flag.IntVar(&workers, "workers", 1, "parallel download workers")
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

	repo, err := repository.NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath)
	if err != nil {
		logger.Error("firestore init failed", "error", err)
		os.Exit(1)
	}
	defer repo.Close()

	bpkConn := bpk.New(logger)
	p := parser.New(logger)

	// Step 1: Load parsed law numbers from cache, fallback to Firestore
	logger.Info("loading parsed law numbers")
	haveParsed := loadParsedCache()
	if len(haveParsed) == 0 {
		logger.Info("cache empty — querying Firestore")
		for _, status := range []string{"parsed", "analyzed", "pending_approval"} {
			docs, err := repo.ListLawsByStatus(ctx, status)
			if err != nil {
				logger.Warn("list laws by status failed", "status", status, "error", err)
				continue
			}
			for _, d := range docs {
				haveParsed[d.LawNumber] = true
			}
		}
		saveParsedCache(haveParsed)
	}
	logger.Info("already parsed laws", "count", len(haveParsed))

	// Step 2: Scrape peraturan.go.id for all active laws
	scr := scraper.New(cfg.Scraper.PythonPath, cfg.Scraper.ScriptPath, logger)
	peraturanConn := peraturan.New(scr, logger)
	logger.Info("scraping peraturan.go.id for all active laws")
	allListed, err := peraturanConn.CheckUpdates(ctx)
	if err != nil {
		logger.Error("scrape peraturan failed", "error", err)
		os.Exit(1)
	}
	logger.Info("scraped from peraturan.go.id", "total", len(allListed))

	// Step 3: Find missing laws
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

	// Step 4: Search BPK for each missing law
	os.MkdirAll("backend/internal/storage", 0755)

	var (
		found        int64
		notFound     int64
		failed       int64
		wg           sync.WaitGroup
		sem          = make(chan struct{}, workers)
		haveParsedMu sync.Mutex
	)

	for _, m := range missing {
		wg.Add(1)
		go func(meta connectors.DocumentMeta) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

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

			raw, err := bpkConn.Download(ctx, *bpkMeta)
			if err != nil {
				logger.Warn("bpk download failed", "law_number", meta.LawNumber, "error", err)
				atomic.AddInt64(&failed, 1)
				return
			}

			doc := &models.LawDocument{
				LawNumber:    meta.LawNumber,
				Title:        meta.Title,
				SourceURL:    bpkMeta.SourceURL,
				Source:       bpkConn.Name(),
				Level:        meta.Level,
				DocumentType: meta.DocumentType,
				Status:       "downloaded",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			docID, err := repo.SaveLawDocument(ctx, doc)
			if err != nil {
				logger.Error("save doc failed", "law_number", meta.LawNumber, "error", err)
				raw.Content.Close()
				atomic.AddInt64(&failed, 1)
				return
			}
			doc.ID = docID

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
			logger.Info("backfilled from BPK",
				"law_number", meta.LawNumber,
				"found", atomic.LoadInt64(&found),
				"not_found", atomic.LoadInt64(&notFound),
				"chars", len(text),
			)

			haveParsedMu.Lock()
			haveParsed[meta.LawNumber] = true
			saveParsedCache(haveParsed)
			haveParsedMu.Unlock()

			// Force GC to release PDF text memory before next law
			runtime.GC()
			debug.FreeOSMemory()
			}(m)
	}

	wg.Wait()

	fmt.Println("\n=== BACKFILL COMPLETE ===")
	fmt.Printf("Total missing:  %d\n", len(missing))
	fmt.Printf("Found on BPK:   %d\n", atomic.LoadInt64(&found))
	fmt.Printf("Not found:      %d\n", atomic.LoadInt64(&notFound))
	fmt.Printf("Failed:         %d\n", atomic.LoadInt64(&failed))
}

func loadParsedCache() map[string]bool {
	data, err := os.ReadFile(parsedCacheFile)
	if err != nil {
		return make(map[string]bool)
	}
	result := make(map[string]bool)
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			result[line] = true
		}
	}
	return result
}

func saveParsedCache(m map[string]bool) {
	var sb strings.Builder
	for k := range m {
		sb.WriteString(k)
		sb.WriteString("\n")
	}
	os.WriteFile(parsedCacheFile, []byte(sb.String()), 0644)
}
