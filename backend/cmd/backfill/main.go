package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/bpk"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/jdihn"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/mkri"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/peraturan"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
	"github.com/rezahanif/hukum-aneh/backend/internal/ai"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/imagegen"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/publishing"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/telegram"
	"github.com/rezahanif/hukum-aneh/backend/internal/validator"
	"github.com/rezahanif/hukum-aneh/backend/internal/workflow"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "enable debug logging")
	flag.Parse()

	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

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

	registry := connectors.NewRegistry()
	scr := scraper.New(cfg.Scraper.PythonPath, cfg.Scraper.ScriptPath, logger)
	registry.Register("Peraturan.go.id", peraturan.New(scr, logger))
	registry.Register("JDIHN", jdihn.New(scr, logger))
	registry.Register("JDIH BPK", bpk.New(scr, logger))
	registry.Register("Mahkamah Konstitusi", mkri.New(scr, logger))

	p := parser.New(logger)
	ret := retrieval.New(cfg, repo)
	aiSvc := ai.New(cfg)
	imgGen := imagegen.New(cfg)
	tgSvc := telegram.New(cfg)
	pubSvc := publishing.New(cfg)
	val := validator.New()

	engine := workflow.NewEngine(cfg, repo, registry, p, ret, aiSvc, imgGen, tgSvc, pubSvc, val, logger)

	// Fetch parsed documents
	logger.Info("fetching parsed documents from Firestore...")
	parsedDocs, err := repo.ListLawsByStatus(ctx, "parsed")
	if err != nil {
		logger.Error("failed to list parsed docs", "error", err)
		os.Exit(1)
	}

	logger.Info("found parsed documents", "count", len(parsedDocs))

	// Run backfill processing
	for i, doc := range parsedDocs {
		logger.Info("processing parsed doc", "index", i+1, "total", len(parsedDocs), "law_number", doc.LawNumber)
		
		// Fetch the corresponding LawVersion
		version, err := repo.GetLatestLawVersion(ctx, doc.ID)
		if err != nil {
			logger.Error("failed to get law version", "doc_id", doc.ID, "error", err)
			continue
		}

		err = engine.ProcessParsedDocument(ctx, &doc, version)
		if err != nil {
			logger.Error("failed to process parsed doc", "law_number", doc.LawNumber, "error", err)
			// Wait 5 seconds before next one to allow rate limits to reset if needed
			time.Sleep(5 * time.Second)
			continue
		}

		logger.Info("successfully processed doc", "law_number", doc.LawNumber)
		time.Sleep(2 * time.Second) // rate limit spacer
	}

	logger.Info("backfill complete")
}
