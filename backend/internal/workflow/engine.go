package workflow

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"bytes"
	"io"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

// Engine orchestrates the full pipeline. Owns all control flow.
// AI agents are workers the engine calls — they never orchestrate. Spec §2.
type Engine struct {
	cfg      *config.Config
	repo     *repository.FirestoreRepo
	registry *connectors.Registry
	parser   *parser.Parser
	logger   *slog.Logger
}

func NewEngine(cfg *config.Config, repo *repository.FirestoreRepo, registry *connectors.Registry, p *parser.Parser, logger *slog.Logger) *Engine {
	return &Engine{
		cfg:      cfg,
		repo:     repo,
		registry: registry,
		parser:   p,
		logger:   logger,
	}
}

// RunDiscovery is the entry point triggered by the Scheduler.
// Iterates all registered connectors, checks for updates, and writes new
// LawDocuments to Firestore. Does NOT parse or analyze — that's event-driven
// off subsequent steps. Spec §3 pipeline.
func (e *Engine) RunDiscovery(ctx context.Context) error {
	e.logger.Info("discovery run started")

	for name, conn := range e.registry.All() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e.logger.Info("checking source", "connector", name)
		docs, err := conn.CheckUpdates(ctx)
		if err != nil {
			e.logger.Error("connector check failed", "connector", name, "error", err)
			continue
		}

		for _, meta := range docs {
			existing, err := e.repo.FindByLawNumber(ctx, meta.LawNumber)
			if err != nil {
				e.logger.Error("dup check failed", "law_number", meta.LawNumber, "error", err)
				continue
			}
			if existing != nil {
				continue
			}

			doc := &models.LawDocument{
				LawNumber:     meta.LawNumber,
				Title:         meta.Title,
				SourceURL:     meta.SourceURL,
				Source:        meta.Source,
				Level:         meta.Level,
				DocumentType:  meta.DocumentType,
				PublishedDate: meta.PublishedDate,
				Status:        "discovered",
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
			id, err := e.repo.SaveLawDocument(ctx, doc)
			if err != nil {
				e.logger.Error("save law doc failed", "law_number", meta.LawNumber, "error", err)
				continue
			}
			e.logger.Info("discovered new law", "id", id, "law_number", meta.LawNumber, "title", meta.Title)
		}
	}

	e.logger.Info("discovery run complete")
	return nil
}

// ProcessDocument handles the download → parse → save pipeline for a single law.
// Called after discovery finds a new law. Spec §3: Document Download → Document Parsing → Save to Database.
func (e *Engine) ProcessDocument(ctx context.Context, doc *models.LawDocument) error {
	conn, ok := e.registry.Get(doc.Source)
	if !ok {
		return fmt.Errorf("no connector for source: %s", doc.Source)
	}

	meta := connectors.DocumentMeta{
		LawNumber:     doc.LawNumber,
		Title:         doc.Title,
		SourceURL:     doc.SourceURL,
		Source:        doc.Source,
		Level:         doc.Level,
		DocumentType:  doc.DocumentType,
		PublishedDate: doc.PublishedDate,
	}

	// Download
	raw, err := conn.Download(ctx, meta)
	if err != nil {
		doc.Status = "download_failed"
		doc.UpdatedAt = time.Now()
		e.repo.SaveLawDocument(ctx, doc)
		return fmt.Errorf("download: %w", err)
	}
	defer raw.Content.Close()

	// Save raw file to storage
	storageDir := "backend/internal/storage"
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return fmt.Errorf("mkdir storage: %w", err)
	}
	rawPath := filepath.Join(storageDir, doc.ID+"_"+raw.Filename)
	rawFile, err := os.Create(rawPath)
	if err != nil {
		return fmt.Errorf("create raw file: %w", err)
	}
	defer rawFile.Close()

	// Tee reader — write to file while parsing
	rawBytes, err := io.ReadAll(raw.Content)
	if err != nil {
		return fmt.Errorf("read raw content: %w", err)
	}
	raw.Content.Close()

	if _, err := rawFile.Write(rawBytes); err != nil {
		return fmt.Errorf("write raw file: %w", err)
	}

	doc.RawFilePath = rawPath
	doc.Status = "downloaded"
	doc.UpdatedAt = time.Now()
	if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
		return fmt.Errorf("save doc status: %w", err)
	}

	// Parse
	result, err := e.parser.Parse(ctx, bytes.NewReader(rawBytes), raw.MimeType, raw.Filename)
	if err != nil {
		doc.Status = "parse_failed"
		doc.UpdatedAt = time.Now()
		e.repo.SaveLawDocument(ctx, doc)
		return fmt.Errorf("parse: %w", err)
	}

	// Save version to Firestore
	version := &models.LawVersion{
		LawDocumentID: doc.ID,
		VersionNumber: int(time.Now().Unix()),
		TextContent:   result.TextContent,
		ParsedAt:      time.Now(),
	}
	if _, err := e.repo.SaveLawVersion(ctx, doc.ID, version); err != nil {
		return fmt.Errorf("save version: %w", err)
	}

	doc.Status = "parsed"
	doc.UpdatedAt = time.Now()
	if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
		return fmt.Errorf("update doc status: %w", err)
	}

	e.logger.Info("document processed", "id", doc.ID, "law_number", doc.LawNumber, "sections", len(result.Sections))
	return nil
}

// CheckStuckJobs finds documents stuck in a stage too long and re-drives them.
// Spec §5.1 responsibility.
func (e *Engine) CheckStuckJobs(ctx context.Context) error {
	threshold, err := time.ParseDuration(e.cfg.Scheduler.StuckJobThreshold)
	if err != nil {
		threshold = 6 * time.Hour
	}
	cutoff := time.Now().Add(-threshold)

	stuck, err := e.repo.FindStuckDocuments(ctx, "discovered", cutoff)
	if err != nil {
		return fmt.Errorf("query stuck: %w", err)
	}

	for _, doc := range stuck {
		e.logger.Warn("stuck document detected", "id", doc.ID, "law_number", doc.LawNumber, "status", doc.Status)
		// Re-trigger: mark for download retry
		doc.Status = "discovered"
		doc.UpdatedAt = time.Now()
		if _, err := e.repo.SaveLawDocument(ctx, &doc); err != nil {
			e.logger.Error("re-queue stuck doc failed", "id", doc.ID, "error", err)
		}
	}

	return nil
}
