package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/ai"
	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/bpk"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/jdihn"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/mkri"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/peraturan"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/setneg"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/drive"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/imagegen"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/publishing"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/telegram"
	"github.com/rezahanif/hukum-aneh/backend/internal/validator"
	"github.com/rezahanif/hukum-aneh/backend/internal/workflow"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

func main() {
	var limit int
	var verbose bool
	var statusFilter string
	var onlyEmpty bool
	flag.IntVar(&limit, "limit", 5, "max number of documents to process (0 = all)")
	flag.BoolVar(&verbose, "verbose", false, "enable debug logging")
	flag.StringVar(&statusFilter, "status", "", "filter by status (e.g. discovered, parse_failed)")
	flag.BoolVar(&onlyEmpty, "only-empty", true, "only process documents with empty text content")
	flag.Parse()

	logLevel := slog.LevelInfo
	if verbose {
		logLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load: %v", err)
	}

	ctx := context.Background()

	// Init repositories (reads STORAGE_MODE from .env)
	repos, err := repository.NewRepoSet(ctx, cfg)
	if err != nil {
		log.Fatalf("repos init: %v", err)
	}
	defer repos.Closer.Close()

	pg, ok := repos.LawRepo.(*repository.PostgresRepo)
	if !ok {
		// If dual_write is active, use the underlying PG repository
		if dw, ok := repos.LawRepo.(*repository.DualWriteRepo); ok {
			pg = dw.Postgres()
		}
	}
	if pg == nil {
		log.Fatalf("PostgreSQL repository is not active. Check STORAGE_MODE in .env.")
	}

	// Init Scraper & Connectors
	scr := scraper.New(cfg.Scraper.PythonPath, cfg.Scraper.ScriptPath, logger)
	registry := connectors.NewRegistry()
	registry.Register("Peraturan.go.id", peraturan.New(scr, logger))
	registry.Register("JDIHN", jdihn.New(scr, logger))
	registry.Register("JDIH BPK", bpk.New(scr, logger))
	registry.Register("Mahkamah Konstitusi", mkri.New(scr, logger))
	registry.Register("JDIH Setneg", setneg.New(scr, logger))

	// Init Parser
	p := parser.New(logger)

	// Init Qdrant Client
	var qdrantClient *retrieval.QdrantClient
	if cfg.IsPostgres() {
		qdrantClient, err = retrieval.NewQdrantClient(ctx, cfg.Qdrant.Host, cfg.Qdrant.Port, cfg.Qdrant.Collection, cfg.Qdrant.APIKey, cfg.Qdrant.VectorSize, logger)
		if err != nil {
			log.Fatalf("qdrant client: %v", err)
		}
		defer qdrantClient.Close()
	}

	// Init Services
	ret, err := retrieval.New(ctx, cfg, repos.EmbedRepo, qdrantClient)
	if err != nil {
		log.Fatalf("retrieval: %v", err)
	}
	aiSvc := ai.New(cfg)
	imgGen := imagegen.New(cfg)
	tgSvc := telegram.New(cfg)
	pubSvc := publishing.New(cfg)
	val := validator.New()

	// Init Google Drive
	var driveSvc *drive.Service
	if _, err := os.Stat(cfg.Google.CredentialsPath); err == nil {
		driveSvc, err = drive.New(ctx, cfg.Google.CredentialsPath, cfg.Google.TokenPath, cfg.Google.FolderID)
		if err != nil {
			log.Fatalf("google drive init: %v", err)
		}
		logger.Info("Google Drive service initialized")
	}

	// Engine
	engine := workflow.NewEngine(cfg, repos, registry, p, ret, qdrantClient, aiSvc, imgGen, tgSvc, pubSvc, driveSvc, val, logger)

	// Fetch candidates to reparse from PG
	query := `
		SELECT ld.id, ld.law_number, ld.title, ld.source_url, ld.source, ld.level, ld.document_type, ld.raw_file_path, ld.published_date, ld.status
		FROM law_documents ld
		LEFT JOIN law_versions lv ON ld.id = lv.law_document_id
	`
	var args []interface{}
	whereAdded := false

	if statusFilter != "" {
		query += " WHERE ld.status = $1"
		args = append(args, statusFilter)
		whereAdded = true
	}

	if onlyEmpty {
		if whereAdded {
			query += " AND (lv.text_content IS NULL OR lv.text_content = '')"
		} else {
			query += " WHERE (lv.text_content IS NULL OR lv.text_content = '')"
		}
	}

	query += " ORDER BY ld.created_at ASC"

	if limit > 0 {
		if onlyEmpty || statusFilter != "" {
			query += fmt.Sprintf(" LIMIT $%d", len(args)+1)
		} else {
			query += " LIMIT $1"
		}
		args = append(args, limit)
	}

	rows, err := pg.Pool().Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("query candidates failed: %v", err)
	}
	defer rows.Close()

	var docs []models.LawDocument
	for rows.Next() {
		var doc models.LawDocument
		err := rows.Scan(
			&doc.ID, &doc.LawNumber, &doc.Title, &doc.SourceURL, &doc.Source,
			&doc.Level, &doc.DocumentType, &doc.RawFilePath, &doc.PublishedDate, &doc.Status,
		)
		if err != nil {
			log.Fatalf("scan candidate failed: %v", err)
		}
		docs = append(docs, doc)
	}

	logger.Info("found candidates to re-download and re-parse", "count", len(docs))

	successCount := 0
	for i, doc := range docs {
		logger.Info("reparsing document", "index", i+1, "total", len(docs), "id", doc.ID, "number", doc.LawNumber)

		// Create processing context with timeout
		procCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
		err := engine.ProcessDocument(procCtx, &doc)
		cancel()

		if err != nil {
			logger.Error("reparse failed", "id", doc.ID, "number", doc.LawNumber, "error", err)
			continue
		}

		successCount++
		logger.Info("reparse succeeded", "id", doc.ID, "number", doc.LawNumber)
	}

	logger.Info("reparse run completed", "total", len(docs), "succeeded", successCount)
}
