package lkpp

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

// LKPPConnector scrapes jdih.lkpp.go.id via Python scraper bridge.
// Covers: Peraturan LKPP / Keputusan Kepala LKPP.
type LKPPConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	client  *http.Client
}

func New(s *scraper.Scraper, logger *slog.Logger) *LKPPConnector {
	return &LKPPConnector{
		scraper: s,
		logger:  logger,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (l *LKPPConnector) Name() string { return "JDIH LKPP" }

func (l *LKPPConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	resp, err := l.scraper.Call(ctx, scraper.ScrapeRequest{
		URL:    "https://jdih.lkpp.go.id/regulation/index",
		Action: "check_updates",
		Source: l.Name(),
	})
	if err != nil {
		l.logger.Warn("lkpp scraper failed", "error", err)
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

func (l *LKPPConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	var resp *http.Response
	var lastErr error

	// We use the Googlebot whitelisted User-Agent to bypass firewall blocks on LKPP download
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
		if err != nil {
			return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

		resp, err = l.client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}

		if err != nil {
			lastErr = err
		} else {
			resp.Body.Close()
			lastErr = fmt.Errorf("status %d for %s", resp.StatusCode, meta.SourceURL)
		}

		l.logger.Warn("lkpp download retry", "attempt", attempt+1, "url", meta.SourceURL, "error", lastErr)
		time.Sleep(time.Duration(attempt+1) * 2 * time.Second)
	}

	if lastErr != nil && (resp == nil || resp.StatusCode != http.StatusOK) {
		return connectors.RawDocument{}, fmt.Errorf("download failed after 3 attempts: %w", lastErr)
	}

	mime := resp.Header.Get("Content-Type")
	if mime == "" {
		mime = "application/pdf"
	}
	if strings.Contains(mime, "text/html") {
		resp.Body.Close()
		return connectors.RawDocument{}, fmt.Errorf("no PDF available for %s (got HTML)", meta.LawNumber)
	}

	return connectors.RawDocument{
		Meta:     meta,
		Content:  resp.Body,
		MimeType: mime,
		Filename: fmt.Sprintf("%s.pdf", meta.LawNumber),
	}, nil
}

func (l *LKPPConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (l *LKPPConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return l.Download(ctx, meta)
}
