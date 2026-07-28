package models

// Firestore collection map for all models.
// Used as single source of truth for collection names.
const (
	ColLaws           = "laws"
	ColContentDrafts  = "content_drafts"
	ColImageAssets    = "image_assets"
	ColApprovals      = "approvals"
	ColPublishingJobs = "publishing_jobs"
	ColEmbeddings     = "embeddings"

	SubVersions       = "versions"
	SubRelationships  = "relationships"
	SubAnalyses       = "analyses"
	SubCaptions       = "captions"
)

// SourceEntry defines a government source to poll.
type SourceEntry struct {
	Name          string `json:"name"`
	Level         string `json:"level"`          // national, sectoral, local
	DocumentType  string `json:"document_type"`
	OfficialSource string `json:"official_source"`
	OfficialURL   string `json:"official_url"`
	Notes         string `json:"notes"`
}
