package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/rezahanif/hukum-aneh/backend/internal/models"
)

// ============================================================================
// Repository Interfaces
// ============================================================================
//
// These 8 interfaces decouple consumers (Engine, retrieval.Service, cmd/)
// from the concrete storage backend (FirestoreRepo today, PostgresRepo
// tomorrow, DualWriteRepo during migration).
//
// Each interface covers one domain aggregate. Method signatures match
// FirestoreRepo exactly — verified by compile-time assertions in
// firestore.go (Phase 0.2) and postgres.go (Phase 3.1).
//
// Adding a new method: add to the interface, then implement on both
// FirestoreRepo and PostgresRepo. The compile-time assertions will catch
// any implementation that falls behind.
// ============================================================================

// LawDocumentRepo covers the laws collection CRUD + status-based queries.
type LawDocumentRepo interface {
	SaveLawDocument(ctx context.Context, doc *models.LawDocument) (string, error)
	GetLawDocument(ctx context.Context, id string) (*models.LawDocument, error)
	FindByLawNumber(ctx context.Context, lawNumber string) (*models.LawDocument, error)
	ListLawsByStatus(ctx context.Context, status string) ([]models.LawDocument, error)
	ListAllLaws(ctx context.Context) ([]models.LawDocument, error)
	FindStuckDocuments(ctx context.Context, status string, before time.Time) ([]models.LawDocument, error)
}

// LawVersionRepo covers the law_versions subcollection.
type LawVersionRepo interface {
	GetLatestLawVersion(ctx context.Context, lawID string) (*models.LawVersion, error)
	SaveLawVersion(ctx context.Context, lawID string, v *models.LawVersion) (string, error)
}

// LawAnalysisRepo covers the law_analyses subcollection.
// GetLawAnalysisByDraft uses a CollectionGroup scan in Firestore; in PostgreSQL
// it becomes a JOIN content_drafts -> law_analyses.
type LawAnalysisRepo interface {
	SaveLawAnalysis(ctx context.Context, lawID string, a *models.LawAnalysis) (string, error)
	GetLawAnalysisByDraft(ctx context.Context, draftID string) (*models.LawAnalysis, error)
}

// ContentDraftRepo covers the content_drafts collection.
type ContentDraftRepo interface {
	GetContentDraft(ctx context.Context, id string) (*models.ContentDraft, error)
	SaveContentDraft(ctx context.Context, draft *models.ContentDraft) (string, error)
}

// ImageAssetRepo covers the image_assets collection.
type ImageAssetRepo interface {
	GetImageAssetsByDraft(ctx context.Context, draftID string) ([]models.ImageAsset, error)
	SaveImageAsset(ctx context.Context, asset *models.ImageAsset) (string, error)
}

// ApprovalRepo covers the approvals collection.
type ApprovalRepo interface {
	SaveApproval(ctx context.Context, a *models.Approval) (string, error)
}

// PublishingJobRepo covers the publishing_jobs collection.
type PublishingJobRepo interface {
	SavePublishingJob(ctx context.Context, j *models.PublishingJob) (string, error)
}

// EmbeddingRepo covers the embeddings collection.
//
// Note: Vector data lives in Firestore today, in Qdrant tomorrow.
// ListEmbeddingsBatch + ListEmbeddingsByDocType return EmbeddingEntry
// including the Vector field (needed for brute-force search fallback).
// After Phase 7.2 (Firestore removal), these methods will return
// metadata-only entries (Vector=nil) — callers must use Qdrant for vectors.
type EmbeddingRepo interface {
	SaveEmbedding(ctx context.Context, emb *models.EmbeddingEntry) (string, error)
	ListAllEmbeddings(ctx context.Context) ([]models.EmbeddingEntry, error)
	ListEmbeddingsBatch(
		ctx context.Context,
		cursor *firestore.DocumentSnapshot,
		limit int,
	) ([]models.EmbeddingEntry, *firestore.DocumentSnapshot, error)
	ListEmbeddingsByDocType(ctx context.Context, docType string) ([]models.EmbeddingEntry, error)
}

// ============================================================================
// IMPORTANT: ListEmbeddingsBatch signature uses *firestore.DocumentSnapshot
// for cursor pagination. This is Firestore-specific.
//
// Phase 3 PostgresRepo will need to implement this same signature to satisfy
// the interface, but it cannot use firestore.DocumentSnapshot meaningfully.
// Options:
//   (a) PostgresRepo returns an error (not used in PG mode)
//   (b) Refactor interface to use a generic cursor type (opaque interface{})
//   (c) Split EmbeddingRepo into FirestoreEmbeddingRepo + PGEmbeddingRepo
//
// Decision: option (a) for now. After Qdrant is integrated (Stream B Phase 4),
// brute-force search is deprecated and ListEmbeddingsBatch becomes dead code.
// PostgresRepo returns an error: "ListEmbeddingsBatch not supported in PG
// mode — use Qdrant for vector search". Clean up in Phase 7.2.
// ============================================================================

// Closer is implemented by all repos that hold resources (Firestore client,
// pgxpool, etc.). Callers should defer Close().
type Closer interface {
	Close() error
}
