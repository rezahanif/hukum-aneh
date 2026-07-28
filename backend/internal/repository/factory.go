package repository

import (
        "context"
        "errors"
        "fmt"
        "log/slog"
        "time"

        "cloud.google.com/go/firestore"

        "github.com/rezahanif/hukum-aneh/backend/internal/config"
        "github.com/rezahanif/hukum-aneh/backend/internal/models"
)

// ============================================================================
// factory.go — Phase 5.1
// ============================================================================
//
// NewRepoSet(ctx, cfg) is the single entry point for constructing a RepoSet.
// It reads cfg.StorageMode and returns the appropriate concrete backend:
//
//   "firestore"  → *FirestoreRepo wrapped in RepoSet
//   "postgres"   → *PostgresRepo wrapped in RepoSet
//   "dual_write" → *DualWriteRepo wrapped in RepoSet
//
// DualWriteRepo is the migration cutover mode: writes go to BOTH Firestore
// (source of truth) and Postgres (best-effort); reads come from Firestore.
// Once PG is validated, flip STORAGE_MODE to "postgres" and remove Firestore.
//
// All cmd/ entry points should call NewRepoSet instead of NewFirestoreRepo
// directly. This is wired up in Phase 5.2.
// ============================================================================

// NewRepoSet is the storage factory. Reads cfg.StorageMode and constructs
// the appropriate concrete RepoSet.
//
// Caller must defer repos.Closer.Close() to release resources.
func NewRepoSet(ctx context.Context, cfg *config.Config) (*RepoSet, error) {
        if cfg == nil {
                return nil, errors.New("nil config")
        }

        mode := cfg.StorageMode
        if mode == "" {
                mode = config.StorageModeFirestore
        }

        switch mode {
        case config.StorageModeFirestore:
                return newFirestoreRepoSet(ctx, cfg)
        case config.StorageModePostgres:
                return newPostgresRepoSet(ctx, cfg)
        case config.StorageModeDualWrite:
                return newDualWriteRepoSet(ctx, cfg)
        default:
                return nil, fmt.Errorf("unknown STORAGE_MODE %q (want firestore|postgres|dual_write)", mode)
        }
}

func newFirestoreRepoSet(ctx context.Context, cfg *config.Config) (*RepoSet, error) {
        repo, err := NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath)
        if err != nil {
                return nil, fmt.Errorf("firestore init: %w", err)
        }
        return NewRepoSetFromFirestore(repo), nil
}

func newPostgresRepoSet(ctx context.Context, cfg *config.Config) (*RepoSet, error) {
        repo, err := NewPostgresRepo(ctx, cfg)
        if err != nil {
                return nil, fmt.Errorf("postgres init: %w", err)
        }
        return NewRepoSetFromPostgres(repo), nil
}

func newDualWriteRepoSet(ctx context.Context, cfg *config.Config) (*RepoSet, error) {
        fs, err := NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath)
        if err != nil {
                return nil, fmt.Errorf("firestore init (dual-write primary): %w", err)
        }
        pg, err := NewPostgresRepo(ctx, cfg)
        if err != nil {
                // Don't leak the Firestore client if PG fails.
                _ = fs.Close()
                return nil, fmt.Errorf("postgres init (dual-write secondary): %w", err)
        }
        dw := &DualWriteRepo{primary: fs, secondary: pg}
        return &RepoSet{
                LawRepo:      dw,
                VersionRepo:  dw,
                AnalysisRepo: dw,
                DraftRepo:    dw,
                ImageRepo:    dw,
                ApprovalRepo: dw,
                PublishRepo:  dw,
                EmbedRepo:    dw,
                Closer:       dw,
        }, nil
}

// ============================================================================
// DualWriteRepo
// ============================================================================
//
// Migration cutover backend. Wraps a Firestore primary + Postgres secondary.
//
// Write semantics:
//   1. Write to Firestore (primary). If it fails, return the error — Firestore
//      is still source of truth, so a failed primary write fails the op.
//   2. Write to Postgres (secondary). If it fails, log a warning but DO NOT
//      fail the op. PG will catch up on the next migrate_to_pg run.
//
// Read semantics:
//   - All reads go to Firestore (primary). PG is write-only during cutover.
//   - Reads switch to PG only when STORAGE_MODE flips to "postgres".
//
// Why not read from PG to validate?
//   Reading from PG during dual-write would surface PG-specific bugs before
//   cutover, but it doubles query load and risks inconsistency if PG is
//   behind. The migrate_to_pg tool (Phase 6.1) does explicit count/diff
//   validation instead — safer and more thorough than runtime spot-checks.
//
// The struct implements all 8 repository interfaces. Compile-time assertions
// below verify the signatures match.
// ============================================================================

// Compile-time assertions for DualWriteRepo.
var (
        _ LawDocumentRepo   = (*DualWriteRepo)(nil)
        _ LawVersionRepo    = (*DualWriteRepo)(nil)
        _ LawAnalysisRepo   = (*DualWriteRepo)(nil)
        _ ContentDraftRepo  = (*DualWriteRepo)(nil)
        _ ImageAssetRepo    = (*DualWriteRepo)(nil)
        _ ApprovalRepo      = (*DualWriteRepo)(nil)
        _ PublishingJobRepo = (*DualWriteRepo)(nil)
        _ EmbeddingRepo     = (*DualWriteRepo)(nil)
        _ Closer            = (*DualWriteRepo)(nil)
)

// DualWriteRepo dual-writes to Firestore (primary) + Postgres (secondary).
type DualWriteRepo struct {
        primary   *FirestoreRepo // source of truth — all reads go here
        secondary *PostgresRepo  // best-effort writes only
}

// Postgres returns the underlying Postgres repository wrapper.
func (d *DualWriteRepo) Postgres() *PostgresRepo {
        return d.secondary
}

// Close releases both backends. Errors from either are logged; final error
// returned is the primary's (matches non-dual-write semantics).
func (d *DualWriteRepo) Close() error {
        priErr := d.primary.Close()
        secErr := d.secondary.Close()
        if secErr != nil {
                slog.Warn("dual-write: secondary close failed", "error", secErr)
        }
        return priErr
}

// logSecFail logs a secondary write failure without failing the op.
// Called after primary write succeeded.
func logSecFail(what string, id string, err error) {
        slog.Warn("dual-write: secondary write failed (primary succeeded)",
                "op", what, "id", id, "error", err)
}

// ============================================================================
// LawDocumentRepo — dual-write
// ============================================================================

func (d *DualWriteRepo) SaveLawDocument(ctx context.Context, doc *models.LawDocument) (string, error) {
        id, err := d.primary.SaveLawDocument(ctx, doc)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SaveLawDocument(ctx, doc); err != nil {
                logSecFail("SaveLawDocument", id, err)
        }
        return id, nil
}

func (d *DualWriteRepo) GetLawDocument(ctx context.Context, id string) (*models.LawDocument, error) {
        return d.primary.GetLawDocument(ctx, id)
}

func (d *DualWriteRepo) FindByLawNumber(ctx context.Context, lawNumber string) (*models.LawDocument, error) {
        return d.primary.FindByLawNumber(ctx, lawNumber)
}

func (d *DualWriteRepo) ListLawsByStatus(ctx context.Context, status string) ([]models.LawDocument, error) {
        return d.primary.ListLawsByStatus(ctx, status)
}

func (d *DualWriteRepo) ListAllLaws(ctx context.Context) ([]models.LawDocument, error) {
        return d.primary.ListAllLaws(ctx)
}

func (d *DualWriteRepo) FindStuckDocuments(ctx context.Context, status string, before time.Time) ([]models.LawDocument, error) {
        return d.primary.FindStuckDocuments(ctx, status, before)
}

// ============================================================================
// LawVersionRepo — dual-write
// ============================================================================

func (d *DualWriteRepo) GetLatestLawVersion(ctx context.Context, lawID string) (*models.LawVersion, error) {
        return d.primary.GetLatestLawVersion(ctx, lawID)
}

func (d *DualWriteRepo) SaveLawVersion(ctx context.Context, lawID string, v *models.LawVersion) (string, error) {
        id, err := d.primary.SaveLawVersion(ctx, lawID, v)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SaveLawVersion(ctx, lawID, v); err != nil {
                logSecFail("SaveLawVersion", id, err)
        }
        return id, nil
}

// ============================================================================
// LawAnalysisRepo — dual-write
// ============================================================================

func (d *DualWriteRepo) SaveLawAnalysis(ctx context.Context, lawID string, a *models.LawAnalysis) (string, error) {
        id, err := d.primary.SaveLawAnalysis(ctx, lawID, a)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SaveLawAnalysis(ctx, lawID, a); err != nil {
                logSecFail("SaveLawAnalysis", id, err)
        }
        return id, nil
}

func (d *DualWriteRepo) GetLawAnalysisByDraft(ctx context.Context, draftID string) (*models.LawAnalysis, error) {
        return d.primary.GetLawAnalysisByDraft(ctx, draftID)
}

// ============================================================================
// ContentDraftRepo — dual-write
// ============================================================================

func (d *DualWriteRepo) GetContentDraft(ctx context.Context, id string) (*models.ContentDraft, error) {
        return d.primary.GetContentDraft(ctx, id)
}

func (d *DualWriteRepo) SaveContentDraft(ctx context.Context, draft *models.ContentDraft) (string, error) {
        id, err := d.primary.SaveContentDraft(ctx, draft)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SaveContentDraft(ctx, draft); err != nil {
                logSecFail("SaveContentDraft", id, err)
        }
        return id, nil
}

// ============================================================================
// ImageAssetRepo — dual-write
// ============================================================================

func (d *DualWriteRepo) GetImageAssetsByDraft(ctx context.Context, draftID string) ([]models.ImageAsset, error) {
        return d.primary.GetImageAssetsByDraft(ctx, draftID)
}

func (d *DualWriteRepo) SaveImageAsset(ctx context.Context, asset *models.ImageAsset) (string, error) {
        id, err := d.primary.SaveImageAsset(ctx, asset)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SaveImageAsset(ctx, asset); err != nil {
                logSecFail("SaveImageAsset", id, err)
        }
        return id, nil
}

// ============================================================================
// ApprovalRepo — dual-write
// ============================================================================

func (d *DualWriteRepo) SaveApproval(ctx context.Context, a *models.Approval) (string, error) {
        id, err := d.primary.SaveApproval(ctx, a)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SaveApproval(ctx, a); err != nil {
                logSecFail("SaveApproval", id, err)
        }
        return id, nil
}

// ============================================================================
// PublishingJobRepo — dual-write
// ============================================================================

func (d *DualWriteRepo) SavePublishingJob(ctx context.Context, j *models.PublishingJob) (string, error) {
        id, err := d.primary.SavePublishingJob(ctx, j)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SavePublishingJob(ctx, j); err != nil {
                logSecFail("SavePublishingJob", id, err)
        }
        return id, nil
}

// ============================================================================
// EmbeddingRepo — dual-write
// ============================================================================
//
// Embeddings need special handling in dual-write:
//   - Primary (Firestore): stores the actual vector. Used for brute-force
//     search until Qdrant is wired (Stream B B-4.1).
//   - Secondary (Postgres): stores metadata only (vector is in Qdrant per
//     schema). The Vector field is preserved in the model so the primary
//     write still works; PG silently ignores it.
//
// Save* writes both. List* reads from primary (which has vectors).

func (d *DualWriteRepo) SaveEmbedding(ctx context.Context, emb *models.EmbeddingEntry) (string, error) {
        id, err := d.primary.SaveEmbedding(ctx, emb)
        if err != nil {
                return "", err
        }
        if _, err := d.secondary.SaveEmbedding(ctx, emb); err != nil {
                logSecFail("SaveEmbedding", id, err)
        }
        return id, nil
}

func (d *DualWriteRepo) ListAllEmbeddings(ctx context.Context) ([]models.EmbeddingEntry, error) {
        return d.primary.ListAllEmbeddings(ctx)
}

func (d *DualWriteRepo) ListEmbeddingsBatch(
        ctx context.Context,
        cursor *firestore.DocumentSnapshot,
        limit int,
) ([]models.EmbeddingEntry, *firestore.DocumentSnapshot, error) {
        return d.primary.ListEmbeddingsBatch(ctx, cursor, limit)
}

func (d *DualWriteRepo) ListEmbeddingsByDocType(ctx context.Context, docType string) ([]models.EmbeddingEntry, error) {
        return d.primary.ListEmbeddingsByDocType(ctx, docType)
}
