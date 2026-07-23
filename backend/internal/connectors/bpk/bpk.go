package bpk

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

type BpkConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	client  *http.Client
}

func New(s *scraper.Scraper, logger *slog.Logger) *BpkConnector {
	return &BpkConnector{
		scraper: s,
		logger:  logger,
		client:  &http.Client{},
	}
}

func (b *BpkConnector) Name() string { return "JDIH BPK" }

func (b *BpkConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	resp, err := b.scraper.Call(ctx, scraper.ScrapeRequest{
		URL:    "https://peraturan.bpk.go.id/",
		Action: "check_updates",
		Source: b.Name(),
	})
	if err != nil {
		b.logger.Warn("bpk scraper failed", "error", err)
		return []connectors.DocumentMeta{}, nil
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

func (b *BpkConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := b.client.Do(req)
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

func (b *BpkConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (b *BpkConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return b.Download(ctx, meta)
}

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "document"
	}
	return parts[len(parts)-1]
}
