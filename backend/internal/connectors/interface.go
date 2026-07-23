package connectors

import (
	"context"
	"io"
)

// Connector is the interface every government source connector must implement.
// Add a new source by implementing these 4 methods — no other module changes.
type Connector interface {
	// Name returns the connector identifier (e.g. "Peraturan.go.id").
	Name() string

	// CheckUpdates polls the source for new/changed laws.
	// Returns list of document metadata for laws not yet in DB.
	CheckUpdates(ctx context.Context) ([]DocumentMeta, error)

	// Download fetches the raw document file (PDF/HTML).
	Download(ctx context.Context, meta DocumentMeta) (RawDocument, error)

	// ExtractMetadata pulls structured metadata from the source page.
	ExtractMetadata(ctx context.Context, raw RawDocument) (DocumentMeta, error)

	// ExtractDocument returns the raw file content for parsing.
	// Deprecated alias for Download — kept for spec compliance.
	ExtractDocument(ctx context.Context, meta DocumentMeta) (RawDocument, error)
}

// DocumentMeta is the metadata extracted from a source listing page.
type DocumentMeta struct {
	LawNumber    string `json:"law_number"`
	Title        string `json:"title"`
	SourceURL    string `json:"source_url"`
	Source       string `json:"source"`
	Level        string `json:"level"`
	DocumentType string `json:"document_type"`
	PublishedDate string `json:"published_date"`
}

// RawDocument is the downloaded raw file.
type RawDocument struct {
	Meta     DocumentMeta
	Content  io.ReadCloser
	MimeType string
	Filename string
}
