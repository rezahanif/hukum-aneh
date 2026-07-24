package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/peraturan"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// Batch tool: scrape all active laws from peraturan.go.id,
// download PDFs, parse text, save to Firestore.
// Dedup by law_number — no copies.
//
// Usage:
//   go run ./backend/cmd/batch -workers=4
//   go run ./backend/cmd/batch -workers=4 -dry-run   # scrape only, no download/parse

func main() {
	var (
		workers int
		dryRun  bool
		verbose bool
	)
	flag.IntVar(&workers, "workers", 4, "parallel download workers")
	flag.BoolVar(&dryRun, "dry-run", false, "scrape listing only, no download/parse/save")
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

	// Connector
	scr := scraper.New(cfg.Scraper.PythonPath, cfg.Scraper.ScriptPath, logger)
	conn := peraturan.New(scr, logger)

	// Parser
	p := parser.New(logger)

	// Step 1: Scrape all active laws
	logger.Info("scraping all active laws from peraturan.go.id")
	startTime := time.Now()

	docs, err := conn.CheckUpdates(ctx)
	if err != nil {
		logger.Error("scrape failed", "error", err)
		os.Exit(1)
	}

	scrapeDuration := time.Since(startTime)
	logger.Info("scrape complete", "total_laws", len(docs), "duration", scrapeDuration)

	if dryRun {
		for _, d := range docs {
			fmt.Printf("%s\t%s\t%s\n", d.LawNumber, d.DocumentType, d.Title)
		}
		return
	}

	// Step 2: Check which laws already exist in Firestore
	var toProcess []connectors.DocumentMeta
	for _, d := range docs {
		existing, err := repo.FindByLawNumber(ctx, d.LawNumber)
		if err != nil {
			logger.Warn("dup check failed", "law_number", d.LawNumber, "error", err)
			toProcess = append(toProcess, d)
			continue
		}
		if existing != nil {
			continue
		}
		toProcess = append(toProcess, d)
	}

	logger.Info("laws to process after dedup", "count", len(toProcess), "already_in_db", len(docs)-len(toProcess))

	if len(toProcess) == 0 {
		logger.Info("nothing to process — all laws already in DB")
		return
	}

	// Step 3: Download + parse + save with worker pool
	storageDir := "backend/internal/storage"
	os.MkdirAll(storageDir, 0755)

	var (
		success  int64
		failed   int64
		wg       sync.WaitGroup
		sem      = make(chan struct{}, workers)
		failedMu sync.Mutex
		failures []string
	)

	for _, meta := range toProcess {
		wg.Add(1)
		go func(m connectors.DocumentMeta) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			if err := processOne(ctx, conn, p, repo, m, storageDir, logger); err != nil {
				errStr := err.Error()
				atomic.AddInt64(&failed, 1)
				failedMu.Lock()
				failures = append(failures, fmt.Sprintf("%s: %v", m.LawNumber, err))
				failedMu.Unlock()
				if strings.Contains(errStr, "no PDF available") {
					logger.Warn("no PDF for law", "law_number", m.LawNumber)
				} else {
					logger.Error("process failed", "law_number", m.LawNumber, "error", err)
				}
			} else {
				atomic.AddInt64(&success, 1)
				logger.Info("processed", "law_number", m.LawNumber, "success", atomic.LoadInt64(&success), "failed", atomic.LoadInt64(&failed))
			}
		}(meta)
	}

	wg.Wait()

	totalDuration := time.Since(startTime)

	// Summary
	fmt.Println("\n=== BATCH COMPLETE ===")
	fmt.Printf("Scraped:    %d laws\n", len(docs))
	fmt.Printf("Processed:  %d (already in DB: %d)\n", len(toProcess), len(docs)-len(toProcess))
	fmt.Printf("Success:    %d\n", atomic.LoadInt64(&success))
	fmt.Printf("Failed:     %d\n", atomic.LoadInt64(&failed))
	fmt.Printf("Duration:   %s\n", totalDuration)

	if len(failures) > 0 {
		fmt.Println("\nFailures:")
		for _, f := range failures {
			fmt.Printf("  - %s\n", f)
		}
		// Write failures to file for retry
		f, _ := os.Create("backend/scripts/failures.txt")
		for _, fail := range failures {
			fmt.Fprintln(f, fail)
		}
		f.Close()
		fmt.Println("\nFailures saved to backend/scripts/failures.txt")
	}
}

func processOne(
	ctx context.Context,
	conn connectors.Connector,
	p *parser.Parser,
	repo *repository.FirestoreRepo,
	meta connectors.DocumentMeta,
	storageDir string,
	logger *slog.Logger,
) error {
	// Download
	raw, err := conn.Download(ctx, meta)
	if err != nil {
		return fmt.Errorf("download: %w", err)
	}
	defer raw.Content.Close()

	// Read raw content
	rawBytes, err := io.ReadAll(raw.Content)
	if err != nil {
		return fmt.Errorf("read raw: %w", err)
	}

	// Save raw file
	rawPath := filepath.Join(storageDir, raw.Filename)
	if err := os.WriteFile(rawPath, rawBytes, 0644); err != nil {
		return fmt.Errorf("save raw: %w", err)
	}

	// Parse before Firestore write so quota failures can still be queued locally.
	result, err := p.Parse(ctx, bytes.NewReader(rawBytes), raw.MimeType, raw.Filename)

	doc := &models.LawDocument{
		LawNumber:    meta.LawNumber,
		Title:        meta.Title,
		SourceURL:    meta.SourceURL,
		Source:       meta.Source,
		Level:        meta.Level,
		DocumentType: meta.DocumentType,
		RawFilePath:  rawPath,
		Status:       "parsed",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err != nil {
		doc.Status = "parse_failed"
		if qerr := saveLocalQueue(meta, doc, nil, fmt.Sprintf("parse: %v", err)); qerr != nil {
			return fmt.Errorf("parse: %w; local queue: %v", err, qerr)
		}
		return fmt.Errorf("parse: %w", err)
	}

	version := &models.LawVersion{
		VersionNumber: int(time.Now().Unix()),
		TextContent:   result.TextContent,
		ParsedAt:      time.Now(),
	}

	docID, err := repo.SaveLawDocument(ctx, doc)
	if err != nil {
		if qerr := saveLocalQueue(meta, doc, version, fmt.Sprintf("save doc: %v", err)); qerr != nil {
			return fmt.Errorf("save doc: %w; local queue: %v", err, qerr)
		}
		logger.Warn("firestore quota/error, saved parsed law to local queue", "law_number", meta.LawNumber, "error", err)
		return nil
	}
	doc.ID = docID
	version.LawDocumentID = doc.ID

	if _, err := repo.SaveLawVersion(ctx, doc.ID, version); err != nil {
		if qerr := saveLocalQueue(meta, doc, version, fmt.Sprintf("save version: %v", err)); qerr != nil {
			return fmt.Errorf("save version: %w; local queue: %v", err, qerr)
		}
		logger.Warn("firestore quota/error, saved parsed law to local queue", "law_number", meta.LawNumber, "error", err)
		return nil
	}

	if _, err := repo.SaveLawDocument(ctx, doc); err != nil {
		if qerr := saveLocalQueue(meta, doc, version, fmt.Sprintf("update status: %v", err)); qerr != nil {
			return fmt.Errorf("update status: %w; local queue: %v", err, qerr)
		}
		logger.Warn("firestore quota/error, saved parsed law to local queue", "law_number", meta.LawNumber, "error", err)
		return nil
	}

	logger.Debug("law processed", "law_number", meta.LawNumber, "text_chars", len(result.TextContent), "source", result.Source)
	return nil
}

type queuedLaw struct {
	Meta     connectors.DocumentMeta `json:"meta"`
	Document *models.LawDocument     `json:"document"`
	Version  *models.LawVersion      `json:"version,omitempty"`
	Error    string                  `json:"error"`
	QueuedAt time.Time               `json:"queued_at"`
}

func saveLocalQueue(meta connectors.DocumentMeta, doc *models.LawDocument, version *models.LawVersion, reason string) error {
	queueDir := "backend/internal/storage/local_queue"
	if err := os.MkdirAll(queueDir, 0755); err != nil {
		return err
	}

	q := queuedLaw{
		Meta:     meta,
		Document: doc,
		Version:  version,
		Error:    reason,
		QueuedAt: time.Now(),
	}
	b, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return err
	}

	name := strings.NewReplacer(" ", "_", "/", "_", "\\", "_", ":", "_", ".", "").Replace(meta.LawNumber)
	if name == "" {
		name = fmt.Sprintf("queued_%d", time.Now().UnixNano())
	}
	return os.WriteFile(filepath.Join(queueDir, name+".json"), b, 0644)
}
