package dpr

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// DPRConnector scrapes jdih.dpr.go.id via Python scraper bridge.
// Covers: Keputusan Presiden (Keppres), Instruksi Presiden (Inpres).
type DPRConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	client  *http.Client
}

func New(s *scraper.Scraper, logger *slog.Logger) *DPRConnector {
	return &DPRConnector{
		scraper: s,
		logger:  logger,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (d *DPRConnector) Name() string { return "JDIH DPR RI" }

func (d *DPRConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	resp, err := d.scraper.Call(ctx, scraper.ScrapeRequest{
		URL:    "https://jdih.dpr.go.id/",
		Action: "check_updates",
		Source: d.Name(),
	})
	if err != nil {
		d.logger.Warn("dpr scraper failed", "error", err)
		return []connectors.DocumentMeta{}, nil
	}

	docs := make([]connectors.DocumentMeta, 0, len(resp.Documents))
	for _, doc := range resp.Documents {
		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     doc.LawNumber,
			Title:         doc.Title,
			SourceURL:     doc.SourceURL,
			Source:        doc.Source,
			Level:         doc.Level,
			DocumentType:  doc.DocumentType,
			PublishedDate: doc.PublishedDate,
		})
	}
	return docs, nil
}

func (d *DPRConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	var resp *http.Response
	var lastErr error

	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
		if err != nil {
			return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

		resp, err = d.client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		if err != nil {
			lastErr = err
		} else {
			resp.Body.Close()
			lastErr = fmt.Errorf("status %d for %s", resp.StatusCode, meta.SourceURL)
		}

		d.logger.Warn("dpr download retry", "attempt", attempt+1, "url", meta.SourceURL, "error", lastErr)
		time.Sleep(time.Duration(attempt+1) * 2 * time.Second)
	}

	if lastErr != nil && (resp == nil || resp.StatusCode != http.StatusOK) {
		return connectors.RawDocument{}, fmt.Errorf("download failed after 3 attempts: %w", lastErr)
	}

	mime := resp.Header.Get("Content-Type")
	if mime == "" {
		mime = "application/pdf"
	}

	return connectors.RawDocument{
		Meta:     meta,
		Content:  resp.Body,
		MimeType: mime,
		Filename: fmt.Sprintf("%s.pdf", meta.LawNumber),
	}, nil
}

func (d *DPRConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (d *DPRConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return d.Download(ctx, meta)
}
