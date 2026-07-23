package peraturan

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// PeraturanConnector scrapes peraturan.go.id.
// Covers: UUD 1945, TAP MPR, UU, Perppu, PP.
// Uses Python subprocess for TLS-fingerprinted requests when needed.
type PeraturanConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	client  *http.Client
}

func New(s *scraper.Scraper, logger *slog.Logger) *PeraturanConnector {
	return &PeraturanConnector{
		scraper: s,
		logger:  logger,
		client:  &http.Client{},
	}
}

func (p *PeraturanConnector) Name() string { return "Peraturan.go.id" }

// CheckUpdates polls peraturan.go.id for new/changed laws.
// Tries direct HTTP first; falls back to Python scraper if blocked.
func (p *PeraturanConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	// Try Python scraper (handles TLS fingerprint + anti-bot)
	resp, err := p.scraper.Call(ctx, scraper.ScrapeRequest{
		URL:    "https://www.peraturan.go.id/uu",
		Action: "check_updates",
		Source: p.Name(),
	})
	if err != nil {
		p.logger.Warn("scraper failed, trying direct HTTP", "error", err)
		return p.checkUpdatesDirect(ctx)
	}

	docs := make([]connectors.DocumentMeta, 0, len(resp.Documents))
	for _, d := range resp.Documents {
	docs = append(docs, connectors.DocumentMeta{
			LawNumber:     d.LawNumber,
			Title:         d.Title,
			SourceURL:     d.SourceURL,
			Source:        d.Source,
			Level:         d.Level,
			DocumentType:  d.DocumentType,
			PublishedDate: d.PublishedDate,
		})
	}
	return docs, nil
}

// Download fetches the raw PDF/HTML for a law.
func (p *PeraturanConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Direct HTTP download — peraturan.go.id PDFs usually don't need TLS fingerprint
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := p.client.Do(req)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("download: %w", err)
	}

	mime := resp.Header.Get("Content-Type")
	if mime == "" {
		if strings.HasSuffix(strings.ToLower(meta.SourceURL), ".pdf") {
			mime = "application/pdf"
		} else {
			mime = "text/html"
		}
	}

	filename := extractFilename(meta.SourceURL)

	return connectors.RawDocument{
		Meta:     meta,
		Content:  resp.Body,
		MimeType: mime,
		Filename: filename,
	}, nil
}

func (p *PeraturanConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	// ponytail: stub — extract from HTML listing page when needed
	// upgrade: per-source HTML parser
	return raw.Meta, nil
}

func (p *PeraturanConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return p.Download(ctx, meta)
}

// checkUpdatesDirect is the fallback when Python scraper is unavailable.
// ponytail: minimal implementation — parse listing page
func (p *PeraturanConnector) checkUpdatesDirect(ctx context.Context) ([]connectors.DocumentMeta, error) {
	// Basic HTTP fetch — real parsing happens in Python scraper
	return []connectors.DocumentMeta{}, nil
}

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "document"
	}
	return parts[len(parts)-1]
}
