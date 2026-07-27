package kemenkeu

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

// KemenkeuConnector scrapes jdih.kemenkeu.go.id.
// Covers: Peraturan Menteri Keuangan (PMK) and Keputusan Menteri Keuangan (KMK).
//
// Site is now a Next.js app. The listing page embeds structured JSON data with
// document metadata (slug, no, tahun, nomor, judul). Detail pages at /dok/{slug}
// contain PDF download links at /api/download/{uuid}/{filename}.pdf.
type KemenkeuConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	perPage  int
	docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *KemenkeuConnector {
	return &KemenkeuConnector{
		scraper:  s,
		logger:   logger,
		client:   &http.Client{Timeout: 60 * time.Second},
		baseURL:  "https://jdih.kemenkeu.go.id",
		perPage:  10,
		docTypes: []string{"Peraturan Menteri Keuangan", "Keputusan Menteri Keuangan"},
	}
}

func (k *KemenkeuConnector) Name() string { return "JDIH Kemenkeu" }

// listingItemRe extracts embedded JSON metadata from the Next.js listing page.
// Matches blocks like:
//
//	"slug":"pmk-31-tahun-2026","bentuk":"Peraturan Menteri","no":31,"tahun":2026,
//	"nomor":"PMK 31 TAHUN 2026","status":"Berlaku","judul":"Perubahan atas..."
//
// Capture groups: (1)slug (2)bentuk (3)no (4)tahun (5)nomor (6)judul
var listingItemRe = regexp.MustCompile(
	`"slug":"([a-z0-9-]+)"[^}]*?"bentuk":"([^"]+)"[^}]*?"no":(\d+)[^}]*?"tahun":(\d+)[^}]*?"nomor":"([^"]+)"[^}]*?"judul":"([^"]*)"`,
)

// docLinkRe is a fallback regex for <a> tags linking to /dok/... pages.
// Capture groups: (1)full href path (/dok/slug) (2)slug
var docLinkRe = regexp.MustCompile(`href="(/dok/([a-z0-9-]+))"`)

// slugPartsRe extracts the type prefix, number, and year from a slug like "pmk-31-tahun-2026".
// Capture groups: (1)prefix (pmk|kmk) (2)number (3)year
var slugPartsRe = regexp.MustCompile(`^(pmk|kmk)-(\d+)-tahun-(\d+)$`)

// pdfLinkRe extracts PDF download link from detail page.
// Matches: href="/api/download/{uuid}/{filename}.pdf"
var pdfLinkRe = regexp.MustCompile(`href="(/api/download/[^"]+\.pdf)"`)

// maxPageRe detects page numbers from pagination links in the HTML.
var maxPageRe = regexp.MustCompile(`page=(\d+)`)

// maxPages caps pagination to a reasonable limit (site shows up to page 12).
const maxPages = 15

func (k *KemenkeuConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	now := time.Now()
	years := []int{now.Year(), now.Year() - 1}

	jenisMap := map[string]string{
		"Peraturan Menteri Keuangan": "PMK",
		"Keputusan Menteri Keuangan": "KMK",
	}

	for _, docType := range k.docTypes {
		jenis, ok := jenisMap[docType]
		if !ok {
			continue
		}
		cursor, hasCursor := cursors.Get(docType)
		k.logger.Info("scraping JDIH Kemenkeu",
			"type", docType, "jenis", jenis,
			"has_cursor", hasCursor, "cursor_law", cursor.LastKnownID,
		)

		for _, year := range years {
			docs, newest, caughtUp, err := k.scrapeYear(ctx, docType, jenis, year, cursor, hasCursor)
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

func (k *KemenkeuConnector) scrapeYear(
	ctx context.Context,
	docType, jenis string,
	year int,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	totalPages := maxPages
	for page := 1; page <= totalPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		url := fmt.Sprintf("%s/search?jenis=%s&tahun=%d&page=%d",
			k.baseURL, jenis, year, page)
		k.logger.Debug("fetching kemenkeu listing", "url", url, "page", page)

		html, err := k.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
		}

		// Detect total pages from pagination links on first page.
		if page == 1 {
			if detected := detectMaxPage(html); detected > 0 {
				totalPages = detected
				if totalPages > maxPages {
					totalPages = maxPages
				}
				k.logger.Debug("detected max pages from pagination", "total", totalPages)
			}
		}

		pageDocs := parseKemenkeuListing(html, docType, year, k.baseURL, k.logger)
		if len(pageDocs) == 0 {
			k.logger.Debug("no results on page, stopping", "page", page)
			break
		}

		if page == 1 && year == time.Now().Year() && len(pageDocs) > 0 {
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

// detectMaxPage finds the highest page number from pagination links in the HTML.
func detectMaxPage(html string) int {
	matches := maxPageRe.FindAllStringSubmatch(html, -1)
	highest := 0
	for _, m := range matches {
		if len(m) >= 2 {
			if p, err := strconv.Atoi(m[1]); err == nil && p > highest {
				highest = p
			}
		}
	}
	return highest
}

func parseKemenkeuListing(html, docType string, year int, baseURL string, logger *slog.Logger) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta

	// Strategy 1: Extract from embedded JSON (Next.js data blob).
	jsonMatches := listingItemRe.FindAllStringSubmatch(html, -1)
	if len(jsonMatches) > 0 {
		for _, m := range jsonMatches {
			if len(m) < 7 {
				continue
			}
			slug := m[1]
			num := m[3]
			yr := m[4]
			nomor := m[5]  // e.g. "PMK 31 TAHUN 2026"
			judul := m[6] // e.g. "Perubahan atas Peraturan..."

			// Filter by document type based on the nomor field or slug prefix.
			nomorLower := strings.ToLower(nomor)
			slugLower := strings.ToLower(slug)
			isPMK := strings.HasPrefix(nomorLower, "pmk") || strings.HasPrefix(slugLower, "pmk")
			isKMK := strings.HasPrefix(nomorLower, "kmk") || strings.HasPrefix(slugLower, "kmk")

			if docType == "Peraturan Menteri Keuangan" && !isPMK {
				continue
			}
			if docType == "Keputusan Menteri Keuangan" && !isKMK {
				continue
			}

			prefix := "PMK"
			if docType == "Keputusan Menteri Keuangan" {
				prefix = "KMK"
			}
			lawNum := fmt.Sprintf("%s No. %s Tahun %s", prefix, num, yr)

			docs = append(docs, connectors.DocumentMeta{
				LawNumber:     lawNum,
				Title:         fmt.Sprintf("%s - %s", nomor, judul),
				SourceURL:     fmt.Sprintf("%s/dok/%s", baseURL, slug),
				Source:        "JDIH Kemenkeu",
				Level:         "national",
				DocumentType:  docType,
				PublishedDate: yr,
			})
		}

		if len(docs) > 0 {
			logger.Debug("parsed listing from embedded JSON", "count", len(docs))
			return docs
		}
	}

	// Strategy 2 (fallback): Extract from <a href="/dok/..."> tags in the rendered HTML.
	linkMatches := docLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range linkMatches {
		if len(m) < 3 {
			continue
		}
		slug := m[2]

		lawNum := extractKemenkeuLawNumberFromSlug(slug, docType, year)
		if lawNum == "" {
			continue
		}

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNum,
			Title:         strings.ReplaceAll(slug, "-", " "),
			SourceURL:     fmt.Sprintf("%s/dok/%s", baseURL, slug),
			Source:        "JDIH Kemenkeu",
			Level:         "national",
			DocumentType:  docType,
			PublishedDate: strconv.Itoa(year),
		})
	}

	logger.Debug("parsed listing from link tags (fallback)", "count", len(docs))
	return docs
}

// extractKemenkeuLawNumberFromSlug builds a law number from a slug like "pmk-31-tahun-2026".
func extractKemenkeuLawNumberFromSlug(slug, docType string, year int) string {
	m := slugPartsRe.FindStringSubmatch(slug)
	if m == nil || len(m) < 4 {
		return ""
	}

	slugPrefix := strings.ToUpper(m[1]) // "PMK" or "KMK"

	// Verify slug prefix matches the requested document type.
	if docType == "Peraturan Menteri Keuangan" && slugPrefix != "PMK" {
		return ""
	}
	if docType == "Keputusan Menteri Keuangan" && slugPrefix != "KMK" {
		return ""
	}

	return fmt.Sprintf("%s No. %s Tahun %s", slugPrefix, m[2], m[3])
}

func (k *KemenkeuConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// SourceURL points to the detail page: /dok/{slug}
	// Fetch it to extract the /api/download/... PDF link.
	html, err := k.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail page %s: %w", meta.SourceURL, err)
	}

	pdfPath := extractKemenkeuPDFPath(html)
	if pdfPath == "" {
		return connectors.RawDocument{}, fmt.Errorf("no PDF link found on detail page for %s (url: %s)", meta.LawNumber, meta.SourceURL)
	}

	pdfURL := k.baseURL + pdfPath
	k.logger.Debug("downloading PDF", "url", pdfURL, "law", meta.LawNumber)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pdfURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build PDF request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

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
		return connectors.RawDocument{}, fmt.Errorf("no PDF available for %s (got HTML, content-type: %s)", meta.LawNumber, mime)
	}

	return connectors.RawDocument{
		Meta:     meta,
		Content:  resp.Body,
		MimeType: mime,
		Filename: extractFilename(pdfURL),
	}, nil
}

// extractKemenkeuPDFPath returns the /api/download/... path from a detail page,
// or empty string if not found. Uses the first matching PDF link (Dokumen).
func extractKemenkeuPDFPath(html string) string {
	if m := pdfLinkRe.FindStringSubmatch(html); m != nil {
		return m[1]
	}
	return ""
}

func (k *KemenkeuConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (k *KemenkeuConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return k.Download(ctx, meta)
}

func (k *KemenkeuConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
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

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]
	if idx := strings.Index(filename, "?"); idx != -1 {
		filename = filename[:idx]
	}
	return filename
}
