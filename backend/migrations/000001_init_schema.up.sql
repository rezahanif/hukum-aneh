-- Hukum-Aneh PostgreSQL Schema v1
-- Migration: 000001_init_schema.up.sql
-- Replaces: Firestore collections (laws, content_drafts, image_assets, approvals, publishing_jobs, embeddings, + subcollections)

CREATE EXTENSION IF NOT EXISTS "pgcrypto";  -- for gen_random_uuid() if needed

-- ============================================================================
-- 1. law_documents (replaces Firestore 'laws' collection)
-- ============================================================================
CREATE TABLE law_documents (
    id            TEXT PRIMARY KEY,
    law_number    TEXT NOT NULL,
    title         TEXT NOT NULL,
    source_url    TEXT,
    source        TEXT,
    level         TEXT,  -- national, sectoral, local
    document_type TEXT,  -- UUD, UU, PP, Perpres, Permen, Kepmen, PBI, etc.
    raw_file_path TEXT,
    published_date TEXT,  -- kept as TEXT to preserve source format (e.g. "2020-01-15" or "15 Januari 2020")
    status        TEXT NOT NULL DEFAULT 'discovered',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT law_documents_status_chk CHECK (
        status IN (
            'discovered', 'downloaded', 'parsed', 'analyzed',
            'no_conflict', 'pending_prompt_approval',
            'pending_approval', 'approved', 'archived',
            'rejected', 'prompt_rejected',
            'download_failed', 'parse_failed'
        )
    )
);

CREATE UNIQUE INDEX law_documents_law_number_uidx ON law_documents (law_number);
CREATE INDEX law_documents_status_idx ON law_documents (status);
CREATE INDEX law_documents_status_updated_idx ON law_documents (status, updated_at);  -- FindStuckDocuments + ListLawsByStatus
CREATE INDEX law_documents_doc_type_idx ON law_documents (document_type);  -- ListEmbeddingsByDocType join

-- ============================================================================
-- 2. law_versions (replaces Firestore subcollection 'laws/{id}/versions')
-- ============================================================================
CREATE TABLE law_versions (
    id              TEXT PRIMARY KEY,
    law_document_id TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,
    version_number  INTEGER NOT NULL,
    text_content    TEXT NOT NULL,  -- TOAST handles >1MB Indonesian regulation texts automatically
    embedding_id    TEXT,  -- ref to Qdrant point ID; vector itself lives in Qdrant, NOT here
    parsed_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX law_versions_doc_ver_idx ON law_versions (law_document_id, version_number DESC);

-- ============================================================================
-- 3. law_relationships (replaces Firestore subcollection 'laws/{id}/relationships')
-- ============================================================================
CREATE TABLE law_relationships (
    id                 TEXT PRIMARY KEY,
    law_document_id    TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,
    related_law_number TEXT,
    relationship_type  TEXT,  -- amends, repeals, supersedes, references
    article_ref        TEXT
);

CREATE INDEX law_relationships_doc_idx ON law_relationships (law_document_id);

-- ============================================================================
-- 4. law_analyses (replaces Firestore subcollection 'laws/{id}/analyses')
-- ============================================================================
CREATE TABLE law_analyses (
    id               TEXT PRIMARY KEY,
    law_document_id  TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,
    summary          TEXT,
    affected_laws    JSONB,  -- array of {law, article, reason, severity} objects
    overall_score    INTEGER,
    controversy_score INTEGER,
    economic_score   INTEGER,
    legal_consistency INTEGER,
    confidence       DOUBLE PRECISION,
    raw_json         TEXT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT law_analyses_overall_chk CHECK (overall_score BETWEEN 0 AND 100),
    CONSTRAINT law_analyses_controv_chk CHECK (controversy_score BETWEEN 0 AND 100),
    CONSTRAINT law_analyses_econ_chk CHECK (economic_score BETWEEN 0 AND 100),
    CONSTRAINT law_analyses_legal_chk CHECK (legal_consistency BETWEEN 0 AND 100)
);

CREATE INDEX law_analyses_doc_idx ON law_analyses (law_document_id);

-- ============================================================================
-- 5. content_drafts (replaces Firestore 'content_drafts' collection)
-- ============================================================================
CREATE TABLE content_drafts (
    id              TEXT PRIMARY KEY,
    law_analysis_id TEXT REFERENCES law_analyses(id) ON DELETE SET NULL,  -- nullable: draft may exist before analysis is finalized
    caption         TEXT,
    hashtags        JSONB,  -- array of strings
    hook            TEXT,
    image_prompt    TEXT,
    status          TEXT NOT NULL DEFAULT 'draft',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT content_drafts_status_chk CHECK (
        status IN (
            'draft', 'pending_prompt_approval', 'pending_approval',
            'approved', 'rejected', 'published',
            'prompt_rejected'
        )
    )
);

CREATE INDEX content_drafts_status_idx ON content_drafts (status);
CREATE INDEX content_drafts_analysis_idx ON content_drafts (law_analysis_id);

-- ============================================================================
-- 6. captions (replaces Firestore subcollection 'content_drafts/{id}/captions')
-- ============================================================================
CREATE TABLE captions (
    id               TEXT PRIMARY KEY,
    content_draft_id TEXT NOT NULL REFERENCES content_drafts(id) ON DELETE CASCADE,
    text             TEXT,
    variant_number   INTEGER
);

CREATE INDEX captions_draft_idx ON captions (content_draft_id);

-- ============================================================================
-- 7. image_assets (replaces Firestore 'image_assets' collection)
-- ============================================================================
CREATE TABLE image_assets (
    id                  TEXT PRIMARY KEY,
    content_draft_id    TEXT NOT NULL REFERENCES content_drafts(id) ON DELETE CASCADE,
    prompt_used         TEXT,
    file_path           TEXT,
    validated           BOOLEAN NOT NULL DEFAULT FALSE,
    design_guide_version TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX image_assets_draft_idx ON image_assets (content_draft_id);

-- ============================================================================
-- 8. approvals (replaces Firestore 'approvals' collection)
-- ============================================================================
CREATE TABLE approvals (
    id              TEXT PRIMARY KEY,
    content_draft_id TEXT NOT NULL REFERENCES content_drafts(id) ON DELETE CASCADE,
    reviewer_id     TEXT,  -- Telegram user ID
    decision        TEXT NOT NULL,
    reason          TEXT,
    timestamp       TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT approvals_decision_chk CHECK (
        decision IN (
            'approve', 'reject',
            'regenerate_caption', 'regenerate_image',
            'prompt_approve', 'prompt_reject', 'prompt_regen'
        )
    )
);

CREATE INDEX approvals_draft_idx ON approvals (content_draft_id);
CREATE INDEX approvals_reviewer_idx ON approvals (reviewer_id);

-- ============================================================================
-- 9. publishing_jobs (replaces Firestore 'publishing_jobs' collection)
-- ============================================================================
CREATE TABLE publishing_jobs (
    id              TEXT PRIMARY KEY,
    content_draft_id TEXT NOT NULL REFERENCES content_drafts(id) ON DELETE CASCADE,
    platform        TEXT NOT NULL,
    status          TEXT NOT NULL DEFAULT 'pending',
    posted_at       TIMESTAMPTZ,
    external_post_id TEXT,

    CONSTRAINT publishing_jobs_platform_chk CHECK (platform = 'instagram'),  -- explicitly NO tiktok
    CONSTRAINT publishing_jobs_status_chk CHECK (status IN ('pending', 'published', 'failed', 'pending_image_hosting'))
);

CREATE INDEX publishing_jobs_draft_idx ON publishing_jobs (content_draft_id);
CREATE INDEX publishing_jobs_status_idx ON publishing_jobs (status);

-- ============================================================================
-- 10. embedding_metadata (replaces Firestore 'embeddings' collection)
-- Vectors themselves live in Qdrant, NOT in PostgreSQL.
-- ============================================================================
CREATE TABLE embedding_metadata (
    id              TEXT PRIMARY KEY,  -- matches Qdrant point ID (UUID)
    law_document_id TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,
    is_mock         BOOLEAN NOT NULL DEFAULT FALSE,  -- critical: filters out mock embeddings from search
    qdrant_point_id TEXT NOT NULL,  -- UUID reference to Qdrant point
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX embedding_metadata_doc_idx ON embedding_metadata (law_document_id);
CREATE INDEX embedding_metadata_mock_idx ON embedding_metadata (is_mock);
CREATE UNIQUE INDEX embedding_metadata_qdrant_idx ON embedding_metadata (qdrant_point_id);

-- ============================================================================
-- Updated_at trigger for law_documents
-- ============================================================================
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER law_documents_updated_at
    BEFORE UPDATE ON law_documents
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
