package mkri

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

type MkriConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	client  *http.Client
}

func New(s *scraper.Scraper, logger *slog.Logger) *MkriConnector {
	return &MkriConnector{
		scraper: s,
		logger:  logger,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (m *MkriConnector) Name() string { return "Mahkamah Konstitusi" }

func (m *MkriConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	resp, err := m.scraper.Call(ctx, scraper.ScrapeRequest{
		URL:    "https://www.mkri.id/perkara/persidangan/putusan",
		Action: "check_updates",
		Source: m.Name(),
	})
	if err != nil {
		m.logger.Warn("mkri scraper failed", "error", err)
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

func (m *MkriConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	var resp *http.Response
	var lastErr error

	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
		if err != nil {
			return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

		resp, err = m.client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		if err != nil {
			lastErr = err
		} else {
			resp.Body.Close()
			lastErr = fmt.Errorf("status %d for %s", resp.StatusCode, meta.SourceURL)
		}

		m.logger.Warn("mkri download retry", "attempt", attempt+1, "url", meta.SourceURL, "error", lastErr)
		time.Sleep(time.Duration(attempt+1) * 2 * time.Second)
	}

	if lastErr != nil && (resp == nil || resp.StatusCode != http.StatusOK) {
		return connectors.RawDocument{}, fmt.Errorf("download failed after 3 attempts: %w", lastErr)
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

func (m *MkriConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (m *MkriConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return m.Download(ctx, meta)
}

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "document"
	}
	return parts[len(parts)-1]
}
