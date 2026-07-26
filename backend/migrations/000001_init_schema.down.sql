-- Rollback: drop all tables in reverse dependency order
DROP TRIGGER IF EXISTS law_documents_updated_at ON law_documents;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS embedding_metadata;
DROP TABLE IF EXISTS publishing_jobs;
DROP TABLE IF EXISTS approvals;
DROP TABLE IF EXISTS image_assets;
DROP TABLE IF EXISTS captions;
DROP TABLE IF EXISTS content_drafts;
DROP TABLE IF EXISTS law_analyses;
DROP TABLE IF EXISTS law_relationships;
DROP TABLE IF EXISTS law_versions;
DROP TABLE IF EXISTS law_documents;
