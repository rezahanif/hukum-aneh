package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

type queuedLaw struct {
	Meta     connectors.DocumentMeta `json:"meta"`
	Document *models.LawDocument     `json:"document"`
	Version  *models.LawVersion      `json:"version,omitempty"`
	Error    string                  `json:"error"`
}

func main() {
	var queueDir string
	flag.StringVar(&queueDir, "queue", "backend/internal/storage/local_queue", "local queued law JSON directory")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
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

	files, err := filepath.Glob(filepath.Join(queueDir, "*.json"))
	if err != nil {
		logger.Error("glob queue", "error", err)
		os.Exit(1)
	}

	var pushed, skipped, failed int
	for _, path := range files {
		b, err := os.ReadFile(path)
		if err != nil {
			failed++
			logger.Error("read queue", "file", path, "error", err)
			continue
		}
		var q queuedLaw
		if err := json.Unmarshal(b, &q); err != nil {
			failed++
			logger.Error("decode queue", "file", path, "error", err)
			continue
		}
		if q.Document == nil {
			failed++
			logger.Error("missing document", "file", path)
			continue
		}

		existing, err := repo.FindByLawNumber(ctx, q.Document.LawNumber)
		if err != nil {
			failed++
			logger.Error("dedup check", "law_number", q.Document.LawNumber, "error", err)
			continue
		}
		if existing != nil {
			_ = os.Remove(path)
			skipped++
			continue
		}

		q.Document.ID = ""
		docID, err := repo.SaveLawDocument(ctx, q.Document)
		if err != nil {
			failed++
			logger.Error("save doc", "law_number", q.Document.LawNumber, "error", err)
			continue
		}
		if q.Version != nil {
			q.Version.ID = ""
			q.Version.LawDocumentID = docID
			if _, err := repo.SaveLawVersion(ctx, docID, q.Version); err != nil {
				failed++
				logger.Error("save version", "law_number", q.Document.LawNumber, "error", err)
				continue
			}
		}
		_ = os.Remove(path)
		pushed++
	}
	fmt.Printf("local flush complete: pushed=%d skipped_existing=%d failed=%d remaining=%d\n", pushed, skipped, failed, len(files)-pushed-skipped)
}
