package ma

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
)

// MAConnector scrapes putusan3.mahkamahagung.go.id.
// Covers: Putusan Pengadilan (Mahkamah Agung court decisions).
//
// Site behavior: paginated search by category, 10 results per page.
// URL pattern: /direktori/putusan/page/1?kategori=pidana
type MAConnector struct {
	logger  *slog.Logger
	client  *http.Client
	baseURL string
	perPage int
}

func New(logger *slog.Logger) *MAConnector {
	return &MAConnector{
		logger:  logger,
		client:  &http.Client{Timeout: 60 * time.Second},
		baseURL: "https://putusan3.mahkamahagung.go.id",
		perPage: 10,
	}
}

func (m *MAConnector) Name() string { return "Mahkamah Agung" }

// resultLinkRe matches putusan detail links.
// Pattern: /direktori/putusan/putusan_xxx... or similar
var resultLinkRe = regexp.MustCompile(`href="(/direktori/putusan/putusan[^"]+)"`)

// pdfLinkRe matches PDF download on detail page.
var pdfLinkRe = regexp.MustCompile(`href="(https?://[^"]+\.pdf[^"]*)"`)

// pageInfoRe extracts total entries.
var pageInfoRe = regexp.MustCompile(`of\s+(\d+)\s+(?:results|putusan|records)`)

func (m *MAConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	cursor, hasCursor := cursors.Get(m.Name())
	m.logger.Info("scraping Mahkamah Agung putusan",
		"has_cursor", hasCursor, "cursor_law", cursor.LastKnownID,
	)

	docs, newest, caughtUp, err := m.scrapeRecent(ctx, cursor, hasCursor)
	if err != nil {
		m.logger.Warn("scrape failed", "error", err)
		return nil, err
	}

	for _, d := range docs {
		if seen[d.LawNumber] {
			continue
		}
		seen[d.LawNumber] = true
		allDocs = append(allDocs, d)
	}

	if !caughtUp && newest != "" {
		cursorUpdates[m.Name()] = connectors.Cursor{
			LastKnownID: newest,
			Timestamp:   time.Now(),
		}
	}

	if len(cursorUpdates) > 0 {
		if err := connectors.SaveAll(cursorUpdates); err != nil {
			m.logger.Warn("batch save cursors failed", "error", err)
		}
	}

	return allDocs, nil
}

func (m *MAConnector) scrapeRecent(
	ctx context.Context,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestID string
	caughtUp := false

	// Cap at 50 pages to avoid runaway.
	const maxPages = 50
	totalPages := 1

	for page := 1; page <= totalPages && page <= maxPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		url := fmt.Sprintf("%s/direktori/putusan/paginate/?page=%d", m.baseURL, page)
		html, err := m.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch page=%d: %w", page, err)
		}

		if page == 1 {
			if m2 := pageInfoRe.FindStringSubmatch(html); m2 != nil {
				if total, err := strconv.Atoi(m2[1]); err == nil {
					totalPages = (total + m.perPage - 1) / m.perPage
				}
			}
		}

		pageDocs := parseMAListing(html, m.baseURL)
		if len(pageDocs) == 0 {
			break
		}

		if page == 1 && len(pageDocs) > 0 {
			newestID = pageDocs[0].LawNumber
		}

		for _, d := range pageDocs {
			if hasCursor && d.LawNumber == cursor.LastKnownID {
				m.logger.Info("hit last known, caught up", "law", d.LawNumber)
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

	return docs, newestID, caughtUp, nil
}

func parseMAListing(html, baseURL string) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := resultLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		if len(m) < 2 {
			continue
		}
		href := m[1]
		// Extract nominal ID from URL slug for law_number
		// e.g. /direktori/putusan/putusan_zon4f0jc9cn3vqle3d3an5me4 -> use slug hash
		slug := strings.TrimPrefix(href, "/direktori/putusan/putusan_")
		lawNum := "Putusan MA " + slug[:min(12, len(slug))]

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:    lawNum,
			Title:        "Putusan Mahkamah Agung " + slug[:min(8, len(slug))],
			SourceURL:    baseURL + href,
			Source:       "Mahkamah Agung",
			Level:        "national",
			DocumentType: "Putusan Pengadilan",
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

func (m *MAConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Fetch detail page to get PDF URL
	html, err := m.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail: %w", err)
	}

	pdfURL := extractMAPDFURL(html)
	if pdfURL == "" {
		return connectors.RawDocument{}, fmt.Errorf("no PDF link found for %s", meta.LawNumber)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pdfURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

	resp, err := m.client.Do(req)
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
		Filename: extractFilename(pdfURL),
	}, nil
}

func extractMAPDFURL(html string) string {
	if m := pdfLinkRe.FindStringSubmatch(html); m != nil {
		return m[1]
	}
	return ""
}

func (m *MAConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (m *MAConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return m.Download(ctx, meta)
}

func (m *MAConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

		resp, err := m.client.Do(req)
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

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]
	if idx := strings.Index(filename, "?"); idx != -1 {
		filename = filename[:idx]
	}
	return filename
}
