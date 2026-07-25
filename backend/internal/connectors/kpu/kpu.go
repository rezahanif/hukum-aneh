package kpu

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// KPUConnector scrapes jdih.kpu.go.id.
// Covers: Peraturan KPU and Keputusan KPU.
type KPUConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *KPUConnector {
	return &KPUConnector{
		scraper:  s,
		logger:   logger,
		client:   &http.Client{Timeout: 60 * time.Second},
		baseURL:  "https://jdih.kpu.go.id",
		docTypes: []string{"Peraturan KPU", "Keputusan KPU"},
	}
}

func (k *KPUConnector) Name() string { return "JDIH KPU" }

// resultLinkRe matches detail link on KPU JDIH, e.g. /peraturan-kpu/detail/3tgf5A2viTbUzsFSjXTCqWtzL0FSbWwwS3dkNE1lQ3Jub3R5N0E9PQ
// or /keputusan-kpu/detail/...
var resultLinkRe = regexp.MustCompile(`href="https://jdih.kpu.go.id/(peraturan-kpu|keputusan-kpu)/detail/([a-zA-Z0-9_-]+)"`)

// downloadLinkRe matches download link on detail page, e.g. /peraturan-kpu/download/462
var downloadLinkRe = regexp.MustCompile(`href="https://jdih.kpu.go.id/(peraturan-kpu|keputusan-kpu)/download/(\d+)"`)

func (k *KPUConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	now := time.Now()
	years := []int{now.Year(), now.Year() - 1}

	for _, docType := range k.docTypes {
		cursor, hasCursor := cursors.Get(docType)
		k.logger.Info("scraping JDIH KPU",
			"type", docType,
			"has_cursor", hasCursor, "cursor_law", cursor.LastKnownID,
		)

		for _, year := range years {
			docs, newest, caughtUp, err := k.scrapeYear(ctx, docType, year, cursor, hasCursor)
			if err != nil {
				k.logger.Warn("scrape year failed", "type", docType, "year", year, "error", err)
				continue
			}

			for _, d := range docs {
				if seen[d.LawNumber] {
					continue
				}
				seen[d.LawNumber] = true
				allDocs = append(allDocs, d)
			}

			if !caughtUp && newest != "" {
				cursorUpdates[docType] = connectors.Cursor{
					LastKnownID: newest,
					Timestamp:   time.Now(),
				}
			}
		}
	}

	if len(cursorUpdates) > 0 {
		if err := connectors.SaveAll(cursorUpdates); err != nil {
			k.logger.Warn("batch save cursors failed", "error", err)
		}
	}

	return allDocs, nil
}

func (k *KPUConnector) scrapeYear(
	ctx context.Context,
	docType string,
	year int,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	pathSegment := "peraturan-kpu"
	pageParam := "page_peraturan"
	if docType == "Keputusan KPU" {
		pathSegment = "keputusan-kpu"
		pageParam = "page_keputusan"
	}

	// Try up to 5 pages per docType/year
	const maxPages = 5
	for page := 1; page <= maxPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		url := fmt.Sprintf("%s/%s?%s=%d&year=%d", k.baseURL, pathSegment, pageParam, page, year)

		html, err := k.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
		}

		pageDocs := k.parseListing(html, docType, year)
		if len(pageDocs) == 0 {
			break
		}

		if page == 1 && len(pageDocs) > 0 {
			newestLaw = pageDocs[0].LawNumber
		}

		for _, d := range pageDocs {
			if hasCursor && d.LawNumber == cursor.LastKnownID {
				k.logger.Info("hit last known law, caught up",
					"type", docType, "year", year, "law", d.LawNumber)
				caughtUp = true
				return docs, "", true, nil
			}
			docs = append(docs, d)
		}

		if caughtUp {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	return docs, newestLaw, caughtUp, nil
}

func (k *KPUConnector) parseListing(html string, docType string, year int) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := resultLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		if len(m) < 3 {
			continue
		}
		pathSegment := m[1]
		hashID := m[2]

		detailURL := fmt.Sprintf("%s/%s/detail/%s", k.baseURL, pathSegment, hashID)

		// Create a best-guess LawNumber based on title / slug elements or use a temporary name
		// (We will download the PDF from detail page)
		prefix := "PKPU"
		if docType == "Keputusan KPU" {
			prefix = "Keputusan KPU"
		}
		lawNum := fmt.Sprintf("%s %s %s", prefix, hashID[:min(8, len(hashID))], strconv.Itoa(year))

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNum,
			Title:         fmt.Sprintf("%s Detail %s", docType, hashID[:min(8, len(hashID))]),
			SourceURL:     detailURL,
			Source:        k.Name(),
			Level:         "sectoral",
			DocumentType:  docType,
			PublishedDate: strconv.Itoa(year),
		})
	}
	return docs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (k *KPUConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Fetch the detail page to find the numeric download link
	html, err := k.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail page: %w", err)
	}

	m := downloadLinkRe.FindStringSubmatch(html)
	if m == nil {
		return connectors.RawDocument{}, fmt.Errorf("no download link found on detail page for %s", meta.LawNumber)
	}

	downloadURL := m[0]

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := k.client.Do(req)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("download PDF: %w", err)
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

func (k *KPUConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (k *KPUConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return k.Download(ctx, meta)
}

func (k *KPUConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

		resp, err := k.client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Duration(attempt+1) * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			lastErr = fmt.Errorf("status %d", resp.StatusCode)
			time.Sleep(time.Duration(attempt+1) * time.Second)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return "", err
		}
		return string(body), nil
	}
	return "", lastErr
}
