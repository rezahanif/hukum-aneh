package pemda

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

// PemdaConnector scrapes peraturan regional via JDIHN gateway.
// Covers: Peraturan Daerah (Perda Prov/Kab/Kota) and Peraturan Kepala Daerah (Perkada).
//
// JDIHN aggregates regional regulations. Site behavior similar to other JDIH sites:
// paginated search with year filter.
// URL pattern: search results for "Peraturan Daerah" with regional scope.
type PemdaConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	perPage  int
	docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *PemdaConnector {
	return &PemdaConnector{
		scraper:  s,
		logger:   logger,
		client:   &http.Client{Timeout: 60 * time.Second},
		baseURL:  "https://jdihn.go.id",
		perPage:  10,
		docTypes: []string{"Peraturan Daerah", "Peraturan Kepala Daerah"},
	}
}

func (p *PemdaConnector) Name() string { return "JDIH Pemda" }

// resultLinkRe matches regulation links on JDIHN regional search.
var resultLinkRe = regexp.MustCompile(`href="(/peraturan[^"]+)"[^>]*>\s*([^<]+)\s*</a>`)

// pdfLinkRe extracts PDF URL from detail page.
var pdfLinkRe = regexp.MustCompile(`href="(https?://[^"]+\.pdf[^"]*)"`)

// pageInfoRe for total count.
var pageInfoRe = regexp.MustCompile(`of\s+(\d+)\s+(?:results|peraturan|entries)`)

func (p *PemdaConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	now := time.Now()
	years := []int{now.Year(), now.Year() - 1}

	for _, docType := range p.docTypes {
		cursor, hasCursor := cursors.Get(docType)
		p.logger.Info("scraping JDIH Pemda",
			"type", docType,
			"has_cursor", hasCursor, "cursor_law", cursor.LastKnownID,
		)

		for _, year := range years {
			docs, newest, caughtUp, err := p.scrapeYear(ctx, docType, year, cursor, hasCursor)
			if err != nil {
				p.logger.Warn("scrape year failed", "type", docType, "year", year, "error", err)
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
			p.logger.Warn("batch save cursors failed", "error", err)
		}
	}

	return allDocs, nil
}

func (p *PemdaConnector) scrapeYear(
	ctx context.Context,
	docType string,
	year int,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	totalPages := 1
	for page := 1; page <= totalPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		jenis := "perda"
		if docType == "Peraturan Kepala Daerah" {
			jenis = "perkada"
		}

		url := fmt.Sprintf("%s/search?jenis=%s&tahun=%d&page=%d",
			p.baseURL, jenis, year, page)
		html, err := p.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
		}

		if page == 1 {
			if m := pageInfoRe.FindStringSubmatch(html); m != nil {
				if total, err := strconv.Atoi(m[1]); err == nil {
					totalPages = (total + p.perPage - 1) / p.perPage
				}
			}
		}

		pageDocs := parsePemdaListing(html, docType, year, p.baseURL)
		if len(pageDocs) == 0 {
			break
		}

		if page == 1 && year == time.Now().Year() && len(pageDocs) > 0 {
			newestLaw = pageDocs[0].LawNumber
		}

		for _, d := range pageDocs {
			if hasCursor && d.LawNumber == cursor.LastKnownID {
				p.logger.Info("hit last known law, caught up",
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

func parsePemdaListing(html, docType string, year int, baseURL string) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := resultLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		if len(m) < 3 {
			continue
		}
		href := m[1]
		text := strings.TrimSpace(m[2])

		lawNum := extractPemdaLawNumber(text, href, docType, year)
		if lawNum == "" {
			continue
		}

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNum,
			Title:         text,
			SourceURL:     baseURL + href,
			Source:        "JDIH Pemda",
			Level:         "local",
			DocumentType:  docType,
			PublishedDate: strconv.Itoa(year),
		})
	}
	return docs
}

func extractPemdaLawNumber(text, href, docType string, year int) string {
	re := regexp.MustCompile(`(?i)(?:nomor|no\.?)\s*(\d+)(?:\s*(?:tahun\s*(\d+)))?`)
	if m := re.FindStringSubmatch(text); m != nil {
		num := m[1]
		yr := m[2]
		if yr == "" {
			yr = strconv.Itoa(year)
		}
		prefix := "Perda"
		if docType == "Peraturan Kepala Daerah" {
			prefix = "Perkada"
		}
		return fmt.Sprintf("%s No. %s Tahun %s", prefix, num, yr)
	}
	// Fallback to URL slug
	slugMatch := regexp.MustCompile(`-nomor-(\d+)(?:-tahun-(\d+))?`).FindStringSubmatch(href)
	if len(slugMatch) >= 2 {
		num := slugMatch[1]
		yr := slugMatch[2]
		if yr == "" {
			yr = strconv.Itoa(year)
		}
		prefix := "Perda"
		if docType == "Peraturan Kepala Daerah" {
			prefix = "Perkada"
		}
		return fmt.Sprintf("%s No. %s Tahun %s", prefix, num, yr)
	}
	return ""
}

func (p *PemdaConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Fetch detail page first to get PDF URL
	html, err := p.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail: %w", err)
	}

	pdfURL := extractPemdaPDFURL(html)
	if pdfURL == "" {
		return connectors.RawDocument{}, fmt.Errorf("no PDF link found on detail page for %s", meta.LawNumber)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pdfURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

	resp, err := p.client.Do(req)
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

func extractPemdaPDFURL(html string) string {
	if m := pdfLinkRe.FindStringSubmatch(html); m != nil {
		return m[1]
	}
	return ""
}

func (p *PemdaConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (p *PemdaConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return p.Download(ctx, meta)
}

func (p *PemdaConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

		resp, err := p.client.Do(req)
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
