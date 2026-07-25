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
// Site behavior: paginated search listing filtered by year + document type.
// URL pattern: /search?jenis=PMK&tahun=2024&page=1
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

// resultLinkRe matches links to law detail pages on Kemenkeu JDIH.
// Pattern: /view/12345/peraturan-menteri-keuangan-nomor-xxx
var resultLinkRe = regexp.MustCompile(`href="(/view/\d+/[a-z0-9-]+)"[^>]*>\s*([^<]+)\s*</a>`)

// pdfLinkRe extracts PDF download link from detail page.
var pdfLinkRe = regexp.MustCompile(`href="(/download/\d+/[^"]+\.pdf[^"]*)"`)

// pageInfoRe extracts total entries for pagination.
var pageInfoRe = regexp.MustCompile(`of\s+(\d+)\s+(?:results|entries|results\.)`)

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

	totalPages := 1
	for page := 1; page <= totalPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		url := fmt.Sprintf("%s/search?jenis=%s&tahun=%d&page=%d",
			k.baseURL, jenis, year, page)
		html, err := k.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
		}

		if page == 1 {
			if m := pageInfoRe.FindStringSubmatch(html); m != nil {
				if total, err := strconv.Atoi(m[1]); err == nil {
					totalPages = (total + k.perPage - 1) / k.perPage
				}
			}
		}

		pageDocs := parseKemenkeuListing(html, docType, year, k.baseURL)
		if len(pageDocs) == 0 {
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

func parseKemenkeuListing(html, docType string, year int, baseURL string) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := resultLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		if len(m) < 3 {
			continue
		}
		href := m[1]
		text := strings.TrimSpace(m[2])

		lawNum := extractKemenkeuLawNumber(text, href, docType, year)
		if lawNum == "" {
			continue
		}

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNum,
			Title:         text,
			SourceURL:     baseURL + href,
			Source:        "JDIH Kemenkeu",
			Level:         "national",
			DocumentType:  docType,
			PublishedDate: strconv.Itoa(year),
		})
	}
	return docs
}

func extractKemenkeuLawNumber(text, href, docType string, year int) string {
	re := regexp.MustCompile(`(?i)(?:nomor|no\.?)\s*(\d+)(?:\s*(?:tahun\s*(\d+)))?`)
	if m := re.FindStringSubmatch(text); m != nil {
		num := m[1]
		yr := m[2]
		if yr == "" {
			yr = strconv.Itoa(year)
		}
		prefix := "PMK"
		if docType == "Keputusan Menteri Keuangan" {
			prefix = "KMK"
		}
		return fmt.Sprintf("%s No. %s Tahun %s", prefix, num, yr)
	}
	slugMatch := regexp.MustCompile(`-nomor-(\d+)(?:-tahun-(\d+))?`).FindStringSubmatch(href)
	if len(slugMatch) >= 2 {
		num := slugMatch[1]
		yr := slugMatch[2]
		if yr == "" {
			yr = strconv.Itoa(year)
		}
		prefix := "PMK"
		if docType == "Keputusan Menteri Keuangan" {
			prefix = "KMK"
		}
		return fmt.Sprintf("%s No. %s Tahun %s", prefix, num, yr)
	}
	return ""
}

func (k *KemenkeuConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Fetch detail page first to get PDF URL
	html, err := k.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail: %w", err)
	}

	pdfURL := extractKemenkeuPDFURL(html, k.baseURL)
	if pdfURL == "" {
		return connectors.RawDocument{}, fmt.Errorf("no PDF link found on detail page for %s", meta.LawNumber)
	}

	resp, err := k.client.Get(pdfURL)
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

func extractKemenkeuPDFURL(html, baseURL string) string {
	if m := pdfLinkRe.FindStringSubmatch(html); m != nil {
		href := m[1]
		if strings.HasPrefix(href, "http") {
			return href
		}
		return baseURL + href
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
