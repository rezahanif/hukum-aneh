package setneg

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

// SetnegConnector scrapes jdih.setneg.go.id (JDIH Sekretariat Negara).
// Primary source for Peraturan Presiden (Perpres).
//
// Site behavior: search-results page with year filter, paginated 10 per page.
// Cursor key: law number (e.g. "PERPRES NOMOR 12 TAHUN 2026").
// We crawl current year + previous year — older years are static and can be
// backfilled on demand.
type SetnegConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	perPage  int
	docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *SetnegConnector {
	return &SetnegConnector{
		scraper:  s,
		logger:   logger,
		client:   &http.Client{Timeout: 30 * time.Second},
		baseURL:  "https://jdih.setneg.go.id",
		perPage:  10,
		docTypes: []string{"Peraturan Presiden (Perpres)"},
	}
}

func (s *SetnegConnector) Name() string { return "JDIH Setneg" }

// resultLinkRe matches links to perpres detail page on Setneg.
var resultLinkRe = regexp.MustCompile(`href="(/peraturan[^"]+)"[^>]*>([^<]+)<`)

// pageInfoRe extracts "Showing X to Y of Z entries" for pagination calc.
var pageInfoRe = regexp.MustCompile(`Showing\s+\d+\s+to\s+\d+\s+of\s+(\d+)\s+entries`)

// CheckUpdates polls Setneg for new Perpres by year. We crawl current year
// first; if cursor indicates a year boundary, also crawl previous year.
func (s *SetnegConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()

	now := time.Now()
	years := []int{now.Year(), now.Year() - 1}

	for _, docType := range s.docTypes {
		cursor, hasCursor := cursors.Get(docType)
		s.logger.Info("scraping JDIH Setneg",
			"type", docType,
			"has_cursor", hasCursor,
			"cursor_law", cursor.LastKnownID,
		)

		for _, year := range years {
			docs, newest, caughtUp, err := s.scrapeYear(ctx, docType, year, cursor, hasCursor)
			if err != nil {
				s.logger.Warn("scrape year failed", "type", docType, "year", year, "error", err)
				continue
			}

			for _, d := range docs {
				if seen[d.LawNumber] {
					continue
				}
				seen[d.LawNumber] = true
				allDocs = append(allDocs, d)
			}

			// Only update cursor if we made progress (didn't get caught up early)
			if !caughtUp && newest != "" {
				if err := cursors.Save(docType, connectors.Cursor{
					LastKnownID: newest,
					Timestamp:   time.Now(),
				}); err != nil {
					s.logger.Warn("save cursor failed", "type", docType, "error", err)
				}
			}
		}

		s.logger.Info("JDIH Setneg scrape complete", "type", docType, "total_unique", len(allDocs))
	}

	return allDocs, nil
}

// scrapeYear crawls a single year's listing. Returns (docs, newestLaw, caughtUp, err).
func (s *SetnegConnector) scrapeYear(
	ctx context.Context,
	docType string,
	year int,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	// First, find total pages from page 1
	totalPages := 1
	for page := 1; page <= totalPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		url := fmt.Sprintf("%s/?perPage=%d&page=%d&tahun=%d&jenis=Perpres",
			s.baseURL, s.perPage, page, year)
		html, err := s.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
		}

		// On first page, learn total count
		if page == 1 {
			if m := pageInfoRe.FindStringSubmatch(html); m != nil {
				if total, err := strconv.Atoi(m[1]); err == nil {
					totalPages = (total + s.perPage - 1) / s.perPage
				}
			}
		}

		pageDocs := parseListing(html, year, s.baseURL)
		if len(pageDocs) == 0 {
			break
		}

		// Track newest (first law on page 1 of current year)
		if page == 1 && year == time.Now().Year() && len(pageDocs) > 0 {
			newestLaw = pageDocs[0].LawNumber
		}

		for _, d := range pageDocs {
			if hasCursor && d.LawNumber == cursor.LastKnownID {
				s.logger.Info("hit last known law, caught up",
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

// parseListing extracts law metadata from a Setneg search-results page.
func parseListing(html string, year int, baseURL string) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := resultLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		if len(m) < 3 {
			continue
		}
		href := m[1]
		text := strings.TrimSpace(m[2])

		// Extract law number from link text (e.g. "PERPRES NOMOR 12 TAHUN 2026")
		// Or from URL slug
		lawNum := extractLawNumber(text, href)
		if lawNum == "" {
			continue
		}

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:    lawNum,
			Title:        text,
			SourceURL:    baseURL + href,
			Source:       "JDIH Setneg",
			Level:        "national",
			DocumentType: "Peraturan Presiden (Perpres)",
			PublishedDate: strconv.Itoa(year),
		})
	}
	return docs
}

func extractLawNumber(text, href string) string {
	// Try parsing from text first
	re := regexp.MustCompile(`(?i)(?:nomor|no\.?)\s*(\d+)\s*(?:tahun\s*(\d+))?`)
	if m := re.FindStringSubmatch(text); m != nil {
		num := m[1]
		yr := m[2]
		if yr == "" {
			yr = strconv.Itoa(time.Now().Year())
		}
		return fmt.Sprintf("Perpres No. %s Tahun %s", num, yr)
	}
	// Fall back to URL slug
	slugMatch := regexp.MustCompile(`/perpres[-/](\d+)[-/](\d+)`).FindStringSubmatch(href)
	if len(slugMatch) == 3 {
		return fmt.Sprintf("Perpres No. %s Tahun %s", slugMatch[1], slugMatch[2])
	}
	return ""
}

func (s *SetnegConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Fetch detail page first to get PDF URL
	html, err := s.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail: %w", err)
	}

	pdfURL := extractPDFURL(html, s.baseURL)
	if pdfURL == "" {
		return connectors.RawDocument{}, fmt.Errorf("no PDF link found on detail page for %s", meta.LawNumber)
	}

	resp, err := s.client.Get(pdfURL)
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

var pdfLinkRe = regexp.MustCompile(`href="([^"]+\.pdf[^"]*)"`)

func extractPDFURL(html, baseURL string) string {
	if m := pdfLinkRe.FindStringSubmatch(html); m != nil {
		href := m[1]
		if strings.HasPrefix(href, "http") {
			return href
		}
		return baseURL + href
	}
	return ""
}

func (s *SetnegConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (s *SetnegConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return s.Download(ctx, meta)
}

func (s *SetnegConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

		resp, err := s.client.Do(req)
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
