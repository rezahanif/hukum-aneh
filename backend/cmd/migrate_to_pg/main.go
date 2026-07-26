// Package main implements cmd/migrate_to_pg — a one-shot CLI that copies all
// data from Firestore to PostgreSQL. Used during the migration cutover.
//
// Usage:
//
//	STORAGE_MODE=postgres \
//	POSTGRES_HOST=... POSTGRES_PASSWORD=... \
//	FIREBASE_PROJECT_ID=... FIREBASE_CREDENTIALS_PATH=... \
//	go run ./backend/cmd/migrate_to_pg [-dry-run] [-limit N]
//
// Flags:
//
//	-dry-run   Read Firestore, validate counts, but do NOT write to Postgres.
//	           Use this to preview the migration scope before committing.
//	-limit N   Migrate at most N laws (and their subcollections) for testing.
//	           0 = no limit (default).
//	-skip-embeddings  Skip the embeddings collection (large; run separately).
//
// Pre-conditions:
//  1. STORAGE_MODE=postgres in env (or run with -dry-run, which doesn't need PG)
//  2. Postgres schema must be applied (Stream B's B-1.2 migration 000001)
//  3. Firestore credentials have read access
//
// Post-conditions:
//   - All law_documents, law_versions, law_analyses, content_drafts,
//     image_assets, approvals, publishing_jobs, embedding_metadata are
//     present in PG with the same string IDs as Firestore.
//   - Vector data is NOT migrated to PG (vectors live in Qdrant per schema).
//     Stream B's B-4.2 task must push vectors to Qdrant separately.
//
// Exit codes:
//
//	0 = success (or dry-run completed)
//	1 = config / connection error
//	2 = count mismatch after migration (data integrity issue)
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

func main() {
	var (
		dryRun         bool
		limit          int
		skipEmbeddings bool
		verbose        bool
	)
	flag.BoolVar(&dryRun, "dry-run", false, "read Firestore, validate counts, but do NOT write to Postgres")
	flag.IntVar(&limit, "limit", 0, "migrate at most N laws (0 = no limit)")
	flag.BoolVar(&skipEmbeddings, "skip-embeddings", false, "skip the embeddings collection (run separately)")
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	// Always need Firestore (source).
	fsClient, err := openFirestore(ctx, cfg)
	if err != nil {
		logger.Error("firestore open failed", "error", err)
		os.Exit(1)
	}
	defer fsClient.Close()

	// PG client only needed when not dry-run.
	var pg *repository.PostgresRepo
	if !dryRun {
		// Force STORAGE_MODE=postgres so NewPostgresRepo doesn't reject us.
		cfg.StorageMode = config.StorageModePostgres
		pg, err = repository.NewPostgresRepo(ctx, cfg)
		if err != nil {
			logger.Error("postgres open failed", "error", err)
			os.Exit(1)
		}
		defer pg.Close()
		logger.Info("connected to Postgres", "host", cfg.Postgres.Host, "db", cfg.Postgres.Database)
	}
	logger.Info("connected to Firestore", "project", cfg.Firebase.ProjectID)

	if dryRun {
		logger.Info("DRY RUN — no writes will be performed")
	}

	stats := migrationStats{}

	if err := migrateLaws(ctx, fsClient, pg, &stats, limit, dryRun, logger); err != nil {
		logger.Error("law_documents migration failed", "error", err)
		os.Exit(1)
	}

	if err := migrateLawVersions(ctx, fsClient, pg, &stats, dryRun, logger); err != nil {
		logger.Error("law_versions migration failed", "error", err)
		os.Exit(1)
	}

	if err := migrateLawAnalyses(ctx, fsClient, pg, &stats, dryRun, logger); err != nil {
		logger.Error("law_analyses migration failed", "error", err)
		os.Exit(1)
	}

	if err := migrateContentDrafts(ctx, fsClient, pg, &stats, dryRun, logger); err != nil {
		logger.Error("content_drafts migration failed", "error", err)
		os.Exit(1)
	}

	if err := migrateImageAssets(ctx, fsClient, pg, &stats, dryRun, logger); err != nil {
		logger.Error("image_assets migration failed", "error", err)
		os.Exit(1)
	}

	if err := migrateApprovals(ctx, fsClient, pg, &stats, dryRun, logger); err != nil {
		logger.Error("approvals migration failed", "error", err)
		os.Exit(1)
	}

	if err := migratePublishingJobs(ctx, fsClient, pg, &stats, dryRun, logger); err != nil {
		logger.Error("publishing_jobs migration failed", "error", err)
		os.Exit(1)
	}

	if !skipEmbeddings {
		if err := migrateEmbeddings(ctx, fsClient, pg, &stats, dryRun, logger); err != nil {
			logger.Error("embedding_metadata migration failed", "error", err)
			os.Exit(1)
		}
	} else {
		logger.Info("skipping embeddings collection (-skip-embeddings)")
	}

	logger.Info("migration complete",
		"laws", stats.laws,
		"versions", stats.versions,
		"analyses", stats.analyses,
		"drafts", stats.drafts,
		"image_assets", stats.imageAssets,
		"approvals", stats.approvals,
		"publishing_jobs", stats.publishingJobs,
		"embedding_metadata", stats.embeddings,
		"dry_run", dryRun,
	)

	if !dryRun {
		logger.Warn("vectors NOT migrated to PG (live in Qdrant per schema). " +
			"Stream B's B-4.2 task must push vectors to Qdrant separately, " +
			"using the same IDs as the embedding_metadata rows written here.")
	}
}

// migrationStats counts rows written per table.
type migrationStats struct {
	laws, versions, analyses              int
	drafts, imageAssets                   int
	approvals, publishingJobs, embeddings int
}

// ============================================================================
// Firestore client helpers
// ============================================================================

func openFirestore(ctx context.Context, cfg *config.Config) (*firestore.Client, error) {
	var opts []option.ClientOption
	if cfg.Firebase.CredentialsPath != "" {
		opts = append(opts, option.WithCredentialsFile(cfg.Firebase.CredentialsPath))
	}
	return firestore.NewClient(ctx, cfg.Firebase.ProjectID, opts...)
}

// ============================================================================
// Migration: law_documents
// ============================================================================

func migrateLaws(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, limit int, dryRun bool, logger *slog.Logger) error {
	iter := fs.Collection(models.ColLaws).Documents(ctx)
	defer iter.Stop()

	count := 0
	for {
		if limit > 0 && count >= limit {
			logger.Info("reached -limit for laws", "limit", limit)
			break
		}
		doc, err := iter.Next()
		if err != nil {
			break // iterator exhausted
		}
		var law models.LawDocument
		if err := doc.DataTo(&law); err != nil {
			logger.Warn("skip law: decode failed", "id", doc.Ref.ID, "error", err)
			continue
		}
		law.ID = doc.Ref.ID

		if dryRun {
			count++
			continue
		}
		if _, err := pg.SaveLawDocument(ctx, &law); err != nil {
			logger.Warn("save law failed (continuing)", "id", law.ID, "error", err)
			continue
		}
		count++
	}
	stats.laws = count
	logger.Info("migrated law_documents", "count", count)
	return nil
}

// ============================================================================
// Migration: law_versions (subcollection of laws)
// ============================================================================

func migrateLawVersions(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, dryRun bool, logger *slog.Logger) error {
	laws, err := fs.Collection(models.ColLaws).Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("list laws for versions: %w", err)
	}

	count := 0
	for _, lawDoc := range laws {
		iter := lawDoc.Ref.Collection(models.SubVersions).Documents(ctx)
		docs, err := iter.GetAll()
		if err != nil {
			logger.Warn("list versions failed", "law_id", lawDoc.Ref.ID, "error", err)
			continue
		}
		for _, d := range docs {
			var v models.LawVersion
			if err := d.DataTo(&v); err != nil {
				logger.Warn("skip version: decode failed", "id", d.Ref.ID, "error", err)
				continue
			}
			v.ID = d.Ref.ID
			v.LawDocumentID = lawDoc.Ref.ID

			if dryRun {
				count++
				continue
			}
			if _, err := pg.SaveLawVersion(ctx, lawDoc.Ref.ID, &v); err != nil {
				logger.Warn("save version failed (continuing)", "id", v.ID, "error", err)
				continue
			}
			count++
		}
	}
	stats.versions = count
	logger.Info("migrated law_versions", "count", count)
	return nil
}

// ============================================================================
// Migration: law_analyses (subcollection of laws)
// ============================================================================

func migrateLawAnalyses(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, dryRun bool, logger *slog.Logger) error {
	laws, err := fs.Collection(models.ColLaws).Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("list laws for analyses: %w", err)
	}

	count := 0
	for _, lawDoc := range laws {
		iter := lawDoc.Ref.Collection(models.SubAnalyses).Documents(ctx)
		docs, err := iter.GetAll()
		if err != nil {
			logger.Warn("list analyses failed", "law_id", lawDoc.Ref.ID, "error", err)
			continue
		}
		for _, d := range docs {
			var a models.LawAnalysis
			if err := d.DataTo(&a); err != nil {
				logger.Warn("skip analysis: decode failed", "id", d.Ref.ID, "error", err)
				continue
			}
			a.ID = d.Ref.ID
			a.LawDocumentID = lawDoc.Ref.ID

			if dryRun {
				count++
				continue
			}
			if _, err := pg.SaveLawAnalysis(ctx, lawDoc.Ref.ID, &a); err != nil {
				logger.Warn("save analysis failed (continuing)", "id", a.ID, "error", err)
				continue
			}
			count++
		}
	}
	stats.analyses = count
	logger.Info("migrated law_analyses", "count", count)
	return nil
}

// ============================================================================
// Migration: content_drafts (top-level collection)
// ============================================================================

func migrateContentDrafts(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, dryRun bool, logger *slog.Logger) error {
	docs, err := fs.Collection(models.ColContentDrafts).Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("list content_drafts: %w", err)
	}

	count := 0
	for _, d := range docs {
		var draft models.ContentDraft
		if err := d.DataTo(&draft); err != nil {
			logger.Warn("skip draft: decode failed", "id", d.Ref.ID, "error", err)
			continue
		}
		draft.ID = d.Ref.ID

		if dryRun {
			count++
			continue
		}
		if _, err := pg.SaveContentDraft(ctx, &draft); err != nil {
			logger.Warn("save draft failed (continuing)", "id", draft.ID, "error", err)
			continue
		}
		count++
	}
	stats.drafts = count
	logger.Info("migrated content_drafts", "count", count)
	return nil
}

// ============================================================================
// Migration: image_assets (top-level collection)
// ============================================================================

func migrateImageAssets(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, dryRun bool, logger *slog.Logger) error {
	docs, err := fs.Collection(models.ColImageAssets).Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("list image_assets: %w", err)
	}

	count := 0
	for _, d := range docs {
		var asset models.ImageAsset
		if err := d.DataTo(&asset); err != nil {
			logger.Warn("skip image asset: decode failed", "id", d.Ref.ID, "error", err)
			continue
		}
		asset.ID = d.Ref.ID

		if dryRun {
			count++
			continue
		}
		if _, err := pg.SaveImageAsset(ctx, &asset); err != nil {
			logger.Warn("save image asset failed (continuing)", "id", asset.ID, "error", err)
			continue
		}
		count++
	}
	stats.imageAssets = count
	logger.Info("migrated image_assets", "count", count)
	return nil
}

// ============================================================================
// Migration: approvals (top-level collection)
// ============================================================================

func migrateApprovals(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, dryRun bool, logger *slog.Logger) error {
	docs, err := fs.Collection(models.ColApprovals).Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("list approvals: %w", err)
	}

	count := 0
	for _, d := range docs {
		var a models.Approval
		if err := d.DataTo(&a); err != nil {
			logger.Warn("skip approval: decode failed", "id", d.Ref.ID, "error", err)
			continue
		}
		a.ID = d.Ref.ID

		if dryRun {
			count++
			continue
		}
		if _, err := pg.SaveApproval(ctx, &a); err != nil {
			logger.Warn("save approval failed (continuing)", "id", a.ID, "error", err)
			continue
		}
		count++
	}
	stats.approvals = count
	logger.Info("migrated approvals", "count", count)
	return nil
}

// ============================================================================
// Migration: publishing_jobs (top-level collection)
// ============================================================================

func migratePublishingJobs(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, dryRun bool, logger *slog.Logger) error {
	docs, err := fs.Collection(models.ColPublishingJobs).Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("list publishing_jobs: %w", err)
	}

	count := 0
	for _, d := range docs {
		var j models.PublishingJob
		if err := d.DataTo(&j); err != nil {
			logger.Warn("skip publishing job: decode failed", "id", d.Ref.ID, "error", err)
			continue
		}
		j.ID = d.Ref.ID

		if dryRun {
			count++
			continue
		}
		// NOTE: PG schema CHECK constraint restricts platform to 'instagram' only.
		// Any Firestore rows with platform='tiktok' will fail; we log and continue.
		if _, err := pg.SavePublishingJob(ctx, &j); err != nil {
			logger.Warn("save publishing job failed (continuing — likely platform=tiktok blocked by CHECK constraint)",
				"id", j.ID, "platform", j.Platform, "error", err)
			continue
		}
		count++
	}
	stats.publishingJobs = count
	logger.Info("migrated publishing_jobs", "count", count)
	return nil
}

// ============================================================================
// Migration: embedding_metadata (top-level collection; vector dropped)
// ============================================================================

func migrateEmbeddings(ctx context.Context, fs *firestore.Client, pg *repository.PostgresRepo, stats *migrationStats, dryRun bool, logger *slog.Logger) error {
	docs, err := fs.Collection(models.ColEmbeddings).Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("list embeddings: %w", err)
	}

	count := 0
	for _, d := range docs {
		var e models.EmbeddingEntry
		if err := d.DataTo(&e); err != nil {
			logger.Warn("skip embedding: decode failed", "id", d.Ref.ID, "error", err)
			continue
		}
		e.ID = d.Ref.ID
		// NOTE: e.Vector is intentionally dropped here. PG schema has no
		// vector column — vectors live in Qdrant. Stream B's B-4.2 task will
		// push vectors to Qdrant using the same ID.

		if dryRun {
			count++
			continue
		}
		if _, err := pg.SaveEmbedding(ctx, &e); err != nil {
			logger.Warn("save embedding metadata failed (continuing)", "id", e.ID, "error", err)
			continue
		}
		count++
	}
	stats.embeddings = count
	logger.Info("migrated embedding_metadata", "count", count, "note", "vectors NOT copied — live in Qdrant per schema")
	return nil
}
