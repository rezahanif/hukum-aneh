package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"cloud.google.com/go/firestore" // only for the EmbeddingRepo cursor signature

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
)

// ============================================================================
// Compile-time interface assertions (Phase 3.1)
// ============================================================================
//
// Mirror the assertions in firestore.go. If any method is missing or has a
// mismatched signature, the build fails here.
var (
	_ LawDocumentRepo   = (*PostgresRepo)(nil)
	_ LawVersionRepo    = (*PostgresRepo)(nil)
	_ LawAnalysisRepo   = (*PostgresRepo)(nil)
	_ ContentDraftRepo  = (*PostgresRepo)(nil)
	_ ImageAssetRepo    = (*PostgresRepo)(nil)
	_ ApprovalRepo      = (*PostgresRepo)(nil)
	_ PublishingJobRepo = (*PostgresRepo)(nil)
	_ EmbeddingRepo     = (*PostgresRepo)(nil)
	_ Closer            = (*PostgresRepo)(nil)
)

// ============================================================================
// Schema assumptions (must match Stream B migration 000001_init_schema.up.sql)
// ============================================================================
//
// law_documents(id TEXT PK, law_number, title, source_url, source, level,
//   document_type, raw_file_path, published_date, status, created_at, updated_at)
// law_versions(id TEXT PK, law_document_id, version_number, text_content,
//   embedding_id NULLABLE, parsed_at)
// law_relationships(id TEXT PK, law_document_id, related_law_number,
//   relationship_type, article_ref)  -- NOTE: repo interface doesn't expose this yet
// law_analyses(id TEXT PK, law_document_id, summary, affected_laws JSONB,
//   overall_score, controversy_score, economic_score, legal_consistency,
//   confidence, raw_json, created_at)
// content_drafts(id TEXT PK, law_analysis_id NULLABLE, caption, hashtags JSONB,
//   hook, image_prompt, status, created_at)
// image_assets(id TEXT PK, content_draft_id, prompt_used, file_path, validated,
//   design_guide_version, created_at)
// approvals(id TEXT PK, content_draft_id, reviewer_id, decision, reason, timestamp)
// publishing_jobs(id TEXT PK, content_draft_id, platform, status, posted_at NULLABLE,
//   external_post_id)
// embedding_metadata(id TEXT PK, law_document_id, is_mock, qdrant_point_id, created_at)
//
// NOTE: embedding_metadata has NO vector column. Vectors live in Qdrant only.
// The EmbeddingRepo methods here return EmbeddingEntry with Vector=nil.
// Retrieval service must use Qdrant for actual vector search (Stream B B-4.1).
//
// NOTE: ListEmbeddingsBatch returns an error — the signature uses
// *firestore.DocumentSnapshot which is meaningless in PG mode. After Qdrant
// integration (Stream B Phase 4), brute-force batched search is deprecated.
// ============================================================================

// PostgresRepo implements all 8 repository interfaces against PostgreSQL.
type PostgresRepo struct {
	pool *pgxpool.Pool
}

// NewPostgresRepo opens a pgxpool against cfg.Postgres and returns a *PostgresRepo.
// The pool is sized by cfg.Postgres.MaxConns (default 10).
// Caller must defer Close().
func NewPostgresRepo(ctx context.Context, cfg *config.Config) (*PostgresRepo, error) {
	if cfg == nil {
		return nil, errors.New("config is nil")
	}
	if !cfg.IsPostgres() {
		return nil, fmt.Errorf("postgres repo requires STORAGE_MODE=postgres or dual_write, got %q", cfg.StorageMode)
	}
	if cfg.Postgres.Host == "" {
		return nil, errors.New("postgres host not configured")
	}

	poolCfg, err := pgxpool.ParseConfig(cfg.PostgresDSN())
	if err != nil {
		return nil, fmt.Errorf("parse pgxpool config: %w", err)
	}
	if cfg.Postgres.MaxConns > 0 {
		poolCfg.MaxConns = int32(cfg.Postgres.MaxConns)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("open pgxpool: %w", err)
	}

	// Verify connectivity before returning. A misconfigured DSN should fail
	// here at startup, not on the first query.
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &PostgresRepo{pool: pool}, nil
}

// Close releases the pgxpool. Safe to call multiple times.
func (r *PostgresRepo) Close() error {
	if r.pool == nil {
		return nil
	}
	r.pool.Close()
	return nil
}

// Pool exposes the underlying *pgxpool.Pool for tooling that needs direct
// SQL access (e.g. migrate_to_pg, smoke tests, admin cleanup scripts).
//
// Production code should prefer the interface methods — direct pool access
// bypasses the abstraction and will not work with FirestoreRepo or
// DualWriteRepo. Use only in cmd/ utilities that explicitly know they're
// running against PG.
func (r *PostgresRepo) Pool() *pgxpool.Pool {
	return r.pool
}

// ============================================================================
// Helpers
// ============================================================================

// newID returns a fresh UUID string. Used when callers pass an empty ID to
// a Save* method (matches Firestore's auto-ID-on-Add behavior).
func newID() string {
	return uuid.NewString()
}

// errNotFound is returned by Get* methods when the row is missing.
// We don't use pgx.ErrNoRows directly because callers check `err != nil`
// rather than errors.Is — wrapping keeps the error chain clean.
func wrapNoRows(err error, what string) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%s not found", what)
	}
	return err
}

// ============================================================================
// LawDocumentRepo
// ============================================================================

func (r *PostgresRepo) SaveLawDocument(ctx context.Context, doc *models.LawDocument) (string, error) {
	if doc == nil {
		return "", errors.New("nil law document")
	}
	if doc.ID == "" {
		doc.ID = newID()
	}
	if doc.CreatedAt.IsZero() {
		doc.CreatedAt = time.Now()
	}
	doc.UpdatedAt = time.Now()

	const q = `
                INSERT INTO law_documents
                    (id, law_number, title, source_url, source, level, document_type,
                     raw_file_path, published_date, status, created_at, updated_at)
                VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
                ON CONFLICT (id) DO UPDATE SET
                    law_number     = EXCLUDED.law_number,
                    title          = EXCLUDED.title,
                    source_url     = EXCLUDED.source_url,
                    source         = EXCLUDED.source,
                    level          = EXCLUDED.level,
                    document_type  = EXCLUDED.document_type,
                    raw_file_path  = EXcluded.raw_file_path,
                    published_date = EXCLUDED.published_date,
                    status         = EXCLUDED.status,
                    updated_at     = EXCLUDED.updated_at
        `
	_, err := r.pool.Exec(ctx, q,
		doc.ID, doc.LawNumber, doc.Title, doc.SourceURL, doc.Source,
		doc.Level, doc.DocumentType, doc.RawFilePath, doc.PublishedDate,
		doc.Status, doc.CreatedAt, doc.UpdatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("upsert law document: %w", err)
	}
	return doc.ID, nil
}

func (r *PostgresRepo) GetLawDocument(ctx context.Context, id string) (*models.LawDocument, error) {
	const q = `
                SELECT id, law_number, title, source_url, source, level, document_type,
                       raw_file_path, published_date, status, created_at, updated_at
                FROM law_documents WHERE id = $1
        `
	var doc models.LawDocument
	err := r.pool.QueryRow(ctx, q, id).Scan(
		&doc.ID, &doc.LawNumber, &doc.Title, &doc.SourceURL, &doc.Source,
		&doc.Level, &doc.DocumentType, &doc.RawFilePath, &doc.PublishedDate,
		&doc.Status, &doc.CreatedAt, &doc.UpdatedAt,
	)
	if err != nil {
		return nil, wrapNoRows(err, "law document")
	}
	return &doc, nil
}

func (r *PostgresRepo) FindByLawNumber(ctx context.Context, lawNumber string) (*models.LawDocument, error) {
	const q = `
                SELECT id, law_number, title, source_url, source, level, document_type,
                       raw_file_path, published_date, status, created_at, updated_at
                FROM law_documents WHERE law_number = $1 LIMIT 1
        `
	var doc models.LawDocument
	err := r.pool.QueryRow(ctx, q, lawNumber).Scan(
		&doc.ID, &doc.LawNumber, &doc.Title, &doc.SourceURL, &doc.Source,
		&doc.Level, &doc.DocumentType, &doc.RawFilePath, &doc.PublishedDate,
		&doc.Status, &doc.CreatedAt, &doc.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil // not found is not an error here — matches Firestore
	}
	if err != nil {
		return nil, fmt.Errorf("find by law number: %w", err)
	}
	return &doc, nil
}

func (r *PostgresRepo) ListLawsByStatus(ctx context.Context, status string) ([]models.LawDocument, error) {
	const q = `
                SELECT id, law_number, title, source_url, source, level, document_type,
                       raw_file_path, published_date, status, created_at, updated_at
                FROM law_documents WHERE status = $1
                ORDER BY created_at DESC
        `
	rows, err := r.pool.Query(ctx, q, status)
	if err != nil {
		return nil, fmt.Errorf("list laws by status: %w", err)
	}
	defer rows.Close()

	var result []models.LawDocument
	for rows.Next() {
		var doc models.LawDocument
		if err := rows.Scan(
			&doc.ID, &doc.LawNumber, &doc.Title, &doc.SourceURL, &doc.Source,
			&doc.Level, &doc.DocumentType, &doc.RawFilePath, &doc.PublishedDate,
			&doc.Status, &doc.CreatedAt, &doc.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan law: %w", err)
		}
		result = append(result, doc)
	}
	return result, rows.Err()
}

func (r *PostgresRepo) ListAllLaws(ctx context.Context) ([]models.LawDocument, error) {
	const q = `
                SELECT id, law_number, title, source_url, source, level, document_type,
                       raw_file_path, published_date, status, created_at, updated_at
                FROM law_documents ORDER BY created_at DESC
        `
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("list all laws: %w", err)
	}
	defer rows.Close()

	var result []models.LawDocument
	for rows.Next() {
		var doc models.LawDocument
		if err := rows.Scan(
			&doc.ID, &doc.LawNumber, &doc.Title, &doc.SourceURL, &doc.Source,
			&doc.Level, &doc.DocumentType, &doc.RawFilePath, &doc.PublishedDate,
			&doc.Status, &doc.CreatedAt, &doc.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan law: %w", err)
		}
		result = append(result, doc)
	}
	return result, rows.Err()
}

func (r *PostgresRepo) FindStuckDocuments(ctx context.Context, status string, before time.Time) ([]models.LawDocument, error) {
	const q = `
                SELECT id, law_number, title, source_url, source, level, document_type,
                       raw_file_path, published_date, status, created_at, updated_at
                FROM law_documents
                WHERE status = $1 AND updated_at < $2
                ORDER BY updated_at ASC
        `
	rows, err := r.pool.Query(ctx, q, status, before)
	if err != nil {
		return nil, fmt.Errorf("find stuck docs: %w", err)
	}
	defer rows.Close()

	var result []models.LawDocument
	for rows.Next() {
		var doc models.LawDocument
		if err := rows.Scan(
			&doc.ID, &doc.LawNumber, &doc.Title, &doc.SourceURL, &doc.Source,
			&doc.Level, &doc.DocumentType, &doc.RawFilePath, &doc.PublishedDate,
			&doc.Status, &doc.CreatedAt, &doc.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan stuck: %w", err)
		}
		result = append(result, doc)
	}
	return result, rows.Err()
}

// ============================================================================
// LawVersionRepo
// ============================================================================

func (r *PostgresRepo) GetLatestLawVersion(ctx context.Context, lawID string) (*models.LawVersion, error) {
	const q = `
                SELECT id, law_document_id, version_number, text_content, embedding_id, parsed_at
                FROM law_versions
                WHERE law_document_id = $1
                ORDER BY version_number DESC
                LIMIT 1
        `
	var v models.LawVersion
	// embedding_id is nullable in schema; use a pointer to handle NULL.
	var embeddingID *string
	err := r.pool.QueryRow(ctx, q, lawID).Scan(
		&v.ID, &v.LawDocumentID, &v.VersionNumber, &v.TextContent,
		&embeddingID, &v.ParsedAt,
	)
	if err != nil {
		return nil, wrapNoRows(err, "law version")
	}
	if embeddingID != nil {
		v.EmbeddingID = *embeddingID
	}
	return &v, nil
}

func (r *PostgresRepo) SaveLawVersion(ctx context.Context, lawID string, v *models.LawVersion) (string, error) {
	if v == nil {
		return "", errors.New("nil law version")
	}
	if v.VersionNumber == 0 {
		v.VersionNumber = int(time.Now().Unix())
	}
	if v.ID == "" {
		v.ID = newID()
	}
	if v.LawDocumentID == "" {
		v.LawDocumentID = lawID
	}
	if v.ParsedAt.IsZero() {
		v.ParsedAt = time.Now()
	}

	const q = `
                INSERT INTO law_versions
                    (id, law_document_id, version_number, text_content, embedding_id, parsed_at)
                VALUES ($1,$2,$3,$4,$5,$6)
                ON CONFLICT (id) DO UPDATE SET
                    law_document_id = EXCLUDED.law_document_id,
                    version_number  = EXCLUDED.version_number,
                    text_content    = EXCLUDED.text_content,
                    embedding_id    = EXCLUDED.embedding_id,
                    parsed_at       = EXCLUDED.parsed_at
        `
	var embeddingID interface{}
	if v.EmbeddingID != "" {
		embeddingID = v.EmbeddingID
	}
	_, err := r.pool.Exec(ctx, q,
		v.ID, v.LawDocumentID, v.VersionNumber, v.TextContent, embeddingID, v.ParsedAt,
	)
	if err != nil {
		return "", fmt.Errorf("upsert law version: %w", err)
	}
	return v.ID, nil
}

// ============================================================================
// LawAnalysisRepo
// ============================================================================

func (r *PostgresRepo) SaveLawAnalysis(ctx context.Context, lawID string, a *models.LawAnalysis) (string, error) {
	if a == nil {
		return "", errors.New("nil law analysis")
	}
	if a.ID == "" {
		a.ID = newID()
	}
	if a.LawDocumentID == "" {
		a.LawDocumentID = lawID
	}
	a.CreatedAt = time.Now()

	// affected_laws is JSONB — marshal as []byte for pgx.
	var affectedJSON []byte
	var err error
	if len(a.AffectedLaws) > 0 {
		affectedJSON, err = json.Marshal(a.AffectedLaws)
		if err != nil {
			return "", fmt.Errorf("marshal affected_laws: %w", err)
		}
	}

	const q = `
                INSERT INTO law_analyses
                    (id, law_document_id, summary, affected_laws, overall_score,
                     controversy_score, economic_score, legal_consistency, confidence,
                     raw_json, created_at)
                VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
                ON CONFLICT (id) DO UPDATE SET
                    law_document_id    = EXCLUDED.law_document_id,
                    summary            = EXCLUDED.summary,
                    affected_laws      = EXCLUDED.affected_laws,
                    overall_score      = EXCLUDED.overall_score,
                    controversy_score  = EXCLUDED.controversy_score,
                    economic_score     = EXCLUDED.economic_score,
                    legal_consistency  = EXCLUDED.legal_consistency,
                    confidence         = EXCLUDED.confidence,
                    raw_json           = EXCLUDED.raw_json
        `
	_, err = r.pool.Exec(ctx, q,
		a.ID, a.LawDocumentID, a.Summary, affectedJSON,
		a.OverallScore, a.ControversyScore, a.EconomicScore, a.LegalConsistency,
		a.Confidence, a.RawJSON, a.CreatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("upsert law analysis: %w", err)
	}
	return a.ID, nil
}

func (r *PostgresRepo) GetLawAnalysisByDraft(ctx context.Context, draftID string) (*models.LawAnalysis, error) {
	// Two-step: fetch draft to get law_analysis_id, then fetch the analysis.
	// In Firestore this was a CollectionGroup scan; in PG it's two indexed lookups.
	draft, err := r.GetContentDraft(ctx, draftID)
	if err != nil {
		return nil, fmt.Errorf("load draft for analysis lookup: %w", err)
	}
	if draft.LawAnalysisID == "" {
		return nil, fmt.Errorf("draft %s has no law_analysis_id", draftID)
	}

	const q = `
                SELECT id, law_document_id, summary, affected_laws, overall_score,
                       controversy_score, economic_score, legal_consistency, confidence,
                       raw_json, created_at
                FROM law_analyses WHERE id = $1
        `
	var a models.LawAnalysis
	var affectedJSON []byte
	err = r.pool.QueryRow(ctx, q, draft.LawAnalysisID).Scan(
		&a.ID, &a.LawDocumentID, &a.Summary, &affectedJSON,
		&a.OverallScore, &a.ControversyScore, &a.EconomicScore, &a.LegalConsistency,
		&a.Confidence, &a.RawJSON, &a.CreatedAt,
	)
	if err != nil {
		return nil, wrapNoRows(err, "law analysis")
	}
	if len(affectedJSON) > 0 {
		if err := json.Unmarshal(affectedJSON, &a.AffectedLaws); err != nil {
			return nil, fmt.Errorf("unmarshal affected_laws: %w", err)
		}
	}
	return &a, nil
}

// ============================================================================
// ContentDraftRepo
// ============================================================================

func (r *PostgresRepo) GetContentDraft(ctx context.Context, id string) (*models.ContentDraft, error) {
	const q = `
                SELECT id, law_analysis_id, caption, hashtags, hook, image_prompt,
                       status, created_at
                FROM content_drafts WHERE id = $1
        `
	var d models.ContentDraft
	var lawAnalysisID *string
	var hashtagsJSON []byte
	err := r.pool.QueryRow(ctx, q, id).Scan(
		&d.ID, &lawAnalysisID, &d.Caption, &hashtagsJSON,
		&d.Hook, &d.ImagePrompt, &d.Status, &d.CreatedAt,
	)
	if err != nil {
		return nil, wrapNoRows(err, "content draft")
	}
	if lawAnalysisID != nil {
		d.LawAnalysisID = *lawAnalysisID
	}
	if len(hashtagsJSON) > 0 {
		if err := json.Unmarshal(hashtagsJSON, &d.Hashtags); err != nil {
			return nil, fmt.Errorf("unmarshal hashtags: %w", err)
		}
	}
	return &d, nil
}

func (r *PostgresRepo) SaveContentDraft(ctx context.Context, draft *models.ContentDraft) (string, error) {
	if draft == nil {
		return "", errors.New("nil content draft")
	}
	if draft.ID == "" {
		draft.ID = newID()
	}
	draft.CreatedAt = time.Now()

	var hashtagsJSON []byte
	var err error
	if len(draft.Hashtags) > 0 {
		hashtagsJSON, err = json.Marshal(draft.Hashtags)
		if err != nil {
			return "", fmt.Errorf("marshal hashtags: %w", err)
		}
	}

	var lawAnalysisID interface{}
	if draft.LawAnalysisID != "" {
		lawAnalysisID = draft.LawAnalysisID
	}

	const q = `
                INSERT INTO content_drafts
                    (id, law_analysis_id, caption, hashtags, hook, image_prompt, status, created_at)
                VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
                ON CONFLICT (id) DO UPDATE SET
                    law_analysis_id = EXCLUDED.law_analysis_id,
                    caption         = EXCLUDED.caption,
                    hashtags        = EXCLUDED.hashtags,
                    hook            = EXCLUDED.hook,
                    image_prompt    = EXCLUDED.image_prompt,
                    status          = EXCLUDED.status
        `
	_, err = r.pool.Exec(ctx, q,
		draft.ID, lawAnalysisID, draft.Caption, hashtagsJSON,
		draft.Hook, draft.ImagePrompt, draft.Status, draft.CreatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("upsert content draft: %w", err)
	}
	return draft.ID, nil
}

// ============================================================================
// ImageAssetRepo
// ============================================================================

func (r *PostgresRepo) GetImageAssetsByDraft(ctx context.Context, draftID string) ([]models.ImageAsset, error) {
	const q = `
                SELECT id, content_draft_id, prompt_used, file_path, validated,
                       design_guide_version, created_at
                FROM image_assets WHERE content_draft_id = $1 ORDER BY created_at ASC
        `
	rows, err := r.pool.Query(ctx, q, draftID)
	if err != nil {
		return nil, fmt.Errorf("list image assets: %w", err)
	}
	defer rows.Close()

	var result []models.ImageAsset
	for rows.Next() {
		var a models.ImageAsset
		if err := rows.Scan(
			&a.ID, &a.ContentDraftID, &a.PromptUsed, &a.FilePath,
			&a.Validated, &a.DesignGuideVersion, &a.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan image asset: %w", err)
		}
		result = append(result, a)
	}
	return result, rows.Err()
}

func (r *PostgresRepo) SaveImageAsset(ctx context.Context, asset *models.ImageAsset) (string, error) {
	if asset == nil {
		return "", errors.New("nil image asset")
	}
	if asset.ID == "" {
		asset.ID = newID()
	}
	asset.CreatedAt = time.Now()

	const q = `
                INSERT INTO image_assets
                    (id, content_draft_id, prompt_used, file_path, validated,
                     design_guide_version, created_at)
                VALUES ($1,$2,$3,$4,$5,$6,$7)
                ON CONFLICT (id) DO UPDATE SET
                    content_draft_id     = EXCLUDED.content_draft_id,
                    prompt_used          = EXCLUDED.prompt_used,
                    file_path            = EXCLUDED.file_path,
                    validated            = EXCLUDED.validated,
                    design_guide_version = EXCLUDED.design_guide_version
        `
	_, err := r.pool.Exec(ctx, q,
		asset.ID, asset.ContentDraftID, asset.PromptUsed, asset.FilePath,
		asset.Validated, asset.DesignGuideVersion, asset.CreatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("upsert image asset: %w", err)
	}
	return asset.ID, nil
}

// ============================================================================
// ApprovalRepo
// ============================================================================

func (r *PostgresRepo) SaveApproval(ctx context.Context, a *models.Approval) (string, error) {
	if a == nil {
		return "", errors.New("nil approval")
	}
	if a.ID == "" {
		a.ID = newID()
	}
	a.Timestamp = time.Now()

	const q = `
                INSERT INTO approvals
                    (id, content_draft_id, reviewer_id, decision, reason, timestamp)
                VALUES ($1,$2,$3,$4,$5,$6)
                ON CONFLICT (id) DO UPDATE SET
                    content_draft_id = EXCLUDED.content_draft_id,
                    reviewer_id      = EXCLUDED.reviewer_id,
                    decision         = EXCLUDED.decision,
                    reason           = EXCLUDED.reason,
                    timestamp        = EXCLUDED.timestamp
        `
	_, err := r.pool.Exec(ctx, q,
		a.ID, a.ContentDraftID, a.ReviewerID, a.Decision, a.Reason, a.Timestamp,
	)
	if err != nil {
		return "", fmt.Errorf("upsert approval: %w", err)
	}
	return a.ID, nil
}

// ============================================================================
// PublishingJobRepo
// ============================================================================

func (r *PostgresRepo) SavePublishingJob(ctx context.Context, j *models.PublishingJob) (string, error) {
	if j == nil {
		return "", errors.New("nil publishing job")
	}
	if j.ID == "" {
		j.ID = newID()
	}

	// posted_at is nullable — pass *time.Time directly; pgx handles nil.
	const q = `
                INSERT INTO publishing_jobs
                    (id, content_draft_id, platform, status, posted_at, external_post_id)
                VALUES ($1,$2,$3,$4,$5,$6)
                ON CONFLICT (id) DO UPDATE SET
                    content_draft_id = EXCLUDED.content_draft_id,
                    platform         = EXCLUDED.platform,
                    status           = EXCLUDED.status,
                    posted_at        = EXCLUDED.posted_at,
                    external_post_id = EXCLUDED.external_post_id
        `
	_, err := r.pool.Exec(ctx, q,
		j.ID, j.ContentDraftID, j.Platform, j.Status, j.PostedAt, j.ExternalPostID,
	)
	if err != nil {
		// Surface CHECK constraint violations with context (e.g. tiktok platform).
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23514" { // check_violation
			return "", fmt.Errorf("publishing job check failed (platform=%q status=%q): %w",
				j.Platform, j.Status, err)
		}
		return "", fmt.Errorf("upsert publishing job: %w", err)
	}
	return j.ID, nil
}

// ============================================================================
// EmbeddingRepo
// ============================================================================
//
// PG mode: vectors live in Qdrant, NOT in PG. The embedding_metadata table
// only stores id, law_document_id, is_mock, qdrant_point_id, created_at.
//
// All methods below return EmbeddingEntry with Vector=nil. The retrieval
// service (Stream B B-4.1) must fetch actual vectors from Qdrant using
// qdrant_point_id.

func (r *PostgresRepo) SaveEmbedding(ctx context.Context, emb *models.EmbeddingEntry) (string, error) {
	if emb == nil {
		return "", errors.New("nil embedding")
	}
	if emb.ID == "" {
		emb.ID = newID()
	}
	if emb.CreatedAt.IsZero() {
		emb.CreatedAt = time.Now()
	}

	// qdrant_point_id is NOT NULL in schema. If caller didn't set it, use emb.ID
	// as the Qdrant point ID (1:1 mapping). Stream B's Qdrant upsert (B-4.2)
	// must use the same ID when pushing the vector to Qdrant.
	qdrantID := emb.ID
	if qdrantID == "" {
		qdrantID = newID()
	}

	const q = `
                INSERT INTO embedding_metadata
                    (id, law_document_id, is_mock, qdrant_point_id, created_at)
                VALUES ($1,$2,$3,$4,$5)
                ON CONFLICT (id) DO UPDATE SET
                    law_document_id = EXCLUDED.law_document_id,
                    is_mock         = EXCLUDED.is_mock,
                    qdrant_point_id = EXCLUDED.qdrant_point_id
        `
	_, err := r.pool.Exec(ctx, q,
		emb.ID, emb.LawDocumentID, emb.IsMock, qdrantID, emb.CreatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("upsert embedding metadata: %w", err)
	}
	return emb.ID, nil
}

func (r *PostgresRepo) ListAllEmbeddings(ctx context.Context) ([]models.EmbeddingEntry, error) {
	const q = `
                SELECT id, law_document_id, is_mock, qdrant_point_id, created_at
                FROM embedding_metadata ORDER BY created_at ASC
        `
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("list all embeddings: %w", err)
	}
	defer rows.Close()

	var result []models.EmbeddingEntry
	for rows.Next() {
		var e models.EmbeddingEntry
		var qdrantID string
		if err := rows.Scan(&e.ID, &e.LawDocumentID, &e.IsMock, &qdrantID, &e.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan embedding: %w", err)
		}
		// Vector stays nil in PG mode — retrieval uses Qdrant for vectors.
		result = append(result, e)
	}
	return result, rows.Err()
}

// ListEmbeddingsBatch is NOT supported in PG mode.
//
// The signature uses *firestore.DocumentSnapshot for cursor pagination, which
// is Firestore-specific and meaningless against PostgreSQL. After Qdrant
// integration (Stream B Phase 4), brute-force batched search is deprecated —
// the retrieval service uses Qdrant's HNSW index directly.
//
// PostgresRepo returns an error. Callers in PG mode must use Qdrant for
// vector search. This will be cleaned up in Phase 7.2 (Firestore removal)
// where the interface itself is refactored to drop this method.
func (r *PostgresRepo) ListEmbeddingsBatch(
	ctx context.Context,
	cursor *firestore.DocumentSnapshot,
	limit int,
) ([]models.EmbeddingEntry, *firestore.DocumentSnapshot, error) {
	return nil, nil, errors.New("ListEmbeddingsBatch not supported in PG mode — use Qdrant for vector search")
}

// ListEmbeddingsByDocType returns embedding metadata for all laws of the given
// document_type. Vector field is nil — caller (retrieval service) must fetch
// actual vectors from Qdrant using the returned IDs (Stream B B-4.1).
func (r *PostgresRepo) ListEmbeddingsByDocType(ctx context.Context, docType string) ([]models.EmbeddingEntry, error) {
	const q = `
                SELECT em.id, em.law_document_id, em.is_mock, em.qdrant_point_id, em.created_at
                FROM embedding_metadata em
                JOIN law_documents ld ON ld.id = em.law_document_id
                WHERE ld.document_type = $1
                ORDER BY em.created_at ASC
        `
	rows, err := r.pool.Query(ctx, q, docType)
	if err != nil {
		return nil, fmt.Errorf("list embeddings by doc type: %w", err)
	}
	defer rows.Close()

	var result []models.EmbeddingEntry
	for rows.Next() {
		var e models.EmbeddingEntry
		var qdrantID string
		if err := rows.Scan(&e.ID, &e.LawDocumentID, &e.IsMock, &qdrantID, &e.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan embedding by doc type: %w", err)
		}
		result = append(result, e)
	}
	return result, rows.Err()
}
