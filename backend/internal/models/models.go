package models

import "time"

// LawDocument represents a discovered law at the top level.
// Firestore collection: laws
type LawDocument struct {
	ID           string    `json:"id" firestore:"-"`
	LawNumber    string    `json:"law_number"`
	Title        string    `json:"title"`
	SourceURL    string    `json:"source_url"`
	Source       string    `json:"source"`
	Level        string    `json:"level"`        // national, sectoral, local
	DocumentType string    `json:"document_type"` // UUD, UU, PP, Perpres, etc.
	RawFilePath  string    `json:"raw_file_path"`
	PublishedDate string   `json:"published_date"`
	Status       string    `json:"status"` // discovered, downloaded, parsed, analyzed, published
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// LawVersion represents a parsed version of a law's text.
// Firestore subcollection: laws/{lawId}/versions
type LawVersion struct {
	ID            string    `json:"id" firestore:"-"`
	LawDocumentID string    `json:"law_document_id"`
	VersionNumber int       `json:"version_number"`
	TextContent   string    `json:"text_content"`
	Embedding     []float32 `json:"embedding,omitempty" firestore:"-"`
	EmbeddingID   string    `json:"embedding_id,omitempty"` // ref to vector store
	ParsedAt      time.Time `json:"parsed_at"`
}

// LawRelationship represents a link between laws.
// Firestore subcollection: laws/{lawId}/relationships
type LawRelationship struct {
	ID              string `json:"id" firestore:"-"`
	LawDocumentID   string `json:"law_document_id"`
	RelatedLawNumber string `json:"related_law_number"`
	RelationshipType string `json:"relationship_type"` // amends, repeals, supersedes, references
	ArticleRef      string `json:"article_ref"`
}

// LawAnalysis is the output of the Law Analysis Agent.
// Firestore subcollection: laws/{lawId}/analyses
type LawAnalysis struct {
	ID               string         `json:"id" firestore:"-"`
	LawDocumentID    string         `json:"law_document_id"`
	Summary          string         `json:"summary"`
	AffectedLaws     []AffectedLaw  `json:"affected_laws"`
	OverallScore     int            `json:"overall_score"`
	ControversyScore int            `json:"controversy_score"`
	EconomicScore    int            `json:"economic_score"`
	LegalConsistency int            `json:"legal_consistency"`
	Confidence       float64        `json:"confidence"`
	RawJSON          string         `json:"raw_json"`
	CreatedAt        time.Time      `json:"created_at"`
}

type AffectedLaw struct {
	Law     string  `json:"law"`
	Article string  `json:"article"`
	Reason  string  `json:"reason"`
	Severity float64 `json:"severity"`
}

// ContentDraft is the output of the Content Strategy Agent.
// Firestore collection: content_drafts
type ContentDraft struct {
	ID           string    `json:"id" firestore:"-"`
	LawAnalysisID string   `json:"law_analysis_id"`
	Caption      string    `json:"caption"`
	Hashtags     []string  `json:"hashtags"`
	Hook         string    `json:"hook"`
	Status       string    `json:"status"` // draft, pending_approval, approved, rejected, published
	CreatedAt    time.Time `json:"created_at"`
}

// Caption variant within a draft.
// Firestore subcollection: content_drafts/{draftId}/captions
type Caption struct {
	ID            string `json:"id" firestore:"-"`
	ContentDraftID string `json:"content_draft_id"`
	Text          string `json:"text"`
	VariantNumber int    `json:"variant_number"`
}

// ImageAsset represents a generated image.
// Firestore collection: image_assets
type ImageAsset struct {
	ID                string    `json:"id" firestore:"-"`
	ContentDraftID    string    `json:"content_draft_id"`
	PromptUsed        string    `json:"prompt_used"`
	FilePath          string    `json:"file_path"`
	Validated         bool      `json:"validated"`
	DesignGuideVersion string   `json:"design_guide_version"`
	CreatedAt         time.Time `json:"created_at"`
}

// Approval records the human review decision.
// Firestore collection: approvals
type Approval struct {
	ID            string    `json:"id" firestore:"-"`
	ContentDraftID string   `json:"content_draft_id"`
	ReviewerID    string    `json:"reviewer_id"`
	Decision      string    `json:"decision"` // approve, reject, regenerate_caption, regenerate_image
	Reason        string    `json:"reason"`
	Timestamp     time.Time `json:"timestamp"`
}

// PublishingJob tracks social media publishing status.
// Firestore collection: publishing_jobs
type PublishingJob struct {
	ID             string     `json:"id" firestore:"-"`
	ContentDraftID string     `json:"content_draft_id"`
	Platform       string     `json:"platform"` // instagram, tiktok
	Status         string     `json:"status"`   // pending, published, failed
	PostedAt       *time.Time `json:"posted_at,omitempty"`
	ExternalPostID string     `json:"external_post_id"`
}

// EmbeddingEntry stores the raw float vector in a separate collection.
// Firestore collection: embeddings
type EmbeddingEntry struct {
	ID            string    `json:"id" firestore:"-"`
	LawDocumentID string    `json:"law_document_id"`
	Vector        []float32 `json:"vector"`
}
