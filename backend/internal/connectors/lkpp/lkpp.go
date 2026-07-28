package lkpp

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

// LKPPConnector scrapes jdih.lkpp.go.id.
// Covers: Peraturan LKPP, Keputusan Kepala LKPP, Peraturan Presiden (terjemahan),
// Instruksi Presiden, Surat Edaran, etc.
//
// Site structure: /regulation/{category} listing pages with links to detail pages.
// Detail pages contain download-regulation links that serve PDFs.
// Pagination: ?page=N&per-page=M
type LKPPConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	perPage  int
}

// regulationCategory defines a category of regulations on LKPP.
type regulationCategory struct {
	Path     string // URL path, e.g. "/regulation/peraturan-lkpp"
	Label    string // human-readable label
	DocType  string // DocumentType for DocumentMeta
	Prefix   string // law number prefix
	Level    string // "sectoral" or "national"
}

func New(s *scraper.Scraper, logger *slog.Logger) *LKPPConnector {
	return &LKPPConnector{
		scraper: s,
		logger:  logger,
		client:  &http.Client{Timeout: 60 * time.Second},
		baseURL: "https://jdih.lkpp.go.id",
		perPage: 10,
	}
}

func (l *LKPPConnector) Name() string { return "JDIH LKPP" }

var detailLinkRe = regexp.MustCompile(`href="(/regulation/[^"]+/[^"]+)"`)
var slugNumRe = regexp.MustCompile(`nomor-(\d+)-tahun-(\d+)`)

var categories = []regulationCategory{
	{Path: "/regulation/peraturan-lkpp", Label: "Peraturan LKPP", DocType: "Peraturan LKPP", Prefix: "Peraturan LKPP", Level: "sectoral"},
	{Path: "/regulation/keputusan-kepala-lkpp", Label: "Keputusan Kepala LKPP", DocType: "Keputusan Kepala LKPP", Prefix: "Keputusan Kepala LKPP", Level: "sectoral"},
	{Path: "/regulation/peraturan-presiden", Label: "Perpres (LKPP)", DocType: "Peraturan Presiden (Perpres)", Prefix: "Perpres", Level: "national"},
	{Path: "/regulation/instruksi-presiden", Label: "Inpres (LKPP)", DocType: "Instruksi Presiden (Inpres)", Prefix: "Inpres", Level: "national"},
}

const maxPages = 20

func (l *LKPPConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	for _, cat := range categories {
		cursor, hasCursor := cursors.Get(cat.DocType)
		l.logger.Info("scraping JDIH LKPP",
			"type", cat.Label,
			"has_cursor", hasCursor,
			"cursor_law", cursor.LastKnownID,
		)

		docs, newest, caughtUp, err := l.scrapeCategory(ctx, cat, cursor, hasCursor)
		if err != nil {
			l.logger.Warn("scrape category failed", "type", cat.Label, "error", err)
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
			cursorUpdates[cat.DocType] = connectors.Cursor{
				LastKnownID: newest,
				Timestamp:   time.Now(),
			}
		}
	}

	if len(cursorUpdates) > 0 {
		if err := connectors.SaveAll(cursorUpdates); err != nil {
			l.logger.Warn("batch save cursors failed", "error", err)
		}
	}

	return allDocs, nil
}

func (l *LKPPConnector) scrapeCategory(
	ctx context.Context,
	cat regulationCategory,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	for page := 1; page <= maxPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		url := fmt.Sprintf("%s%s?page=%d&per-page=%d",
			l.baseURL, cat.Path, page, l.perPage)
		l.logger.Debug("fetching LKPP listing", "url", url, "page", page)

		html, err := l.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch page=%d: %w", page, err)
		}

		pageDocs := l.parseListing(html, cat)
		if len(pageDocs) == 0 {
			l.logger.Debug("no results on page, stopping", "page", page)
			break
		}

		l.logger.Debug("parsed LKPP listing page", "page", page, "count", len(pageDocs))

		if page == 1 && len(pageDocs) > 0 {
			newestLaw = pageDocs[0].LawNumber
		}

		for _, d := range pageDocs {
			if hasCursor && d.LawNumber == cursor.LastKnownID {
				l.logger.Info("hit last known law, caught up",
					"type", cat.Label, "law", d.LawNumber)
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

func (l *LKPPConnector) parseListing(html string, cat regulationCategory) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := detailLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		if len(m) < 2 {
			continue
		}
		href := m[1]

		// Extract regulation number and year from the slug.
		slug := href[strings.LastIndex(href, "/")+1:]
		slugNum := slugNumRe.FindStringSubmatch(slug)
		if slugNum == nil || len(slugNum) < 3 {
			continue
		}
		num := slugNum[1]
		slugYear := slugNum[2]

		lawNum := fmt.Sprintf("%s No. %s Tahun %s", cat.Prefix, num, slugYear)

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNum,
			Title:          strings.ReplaceAll(slug, "-", " "),
			SourceURL:      l.baseURL + href,
			Source:         "JDIH LKPP",
			Level:          cat.Level,
			DocumentType:   cat.DocType,
			PublishedDate:  slugYear,
		})
	}
	return docs
}

// detailDownloadRe extracts the download-regulation link from detail page.
var detailDownloadRe = regexp.MustCompile(`href="(/regulation/download-regulation\?id=[^"]+)"`)

func (l *LKPPConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// SourceURL is the detail page. Fetch it to extract the download link.
	detailHTML, err := l.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail page: %w", err)
	}

	var downloadURL string
	if dlMatch := detailDownloadRe.FindStringSubmatch(detailHTML); dlMatch != nil {
		downloadURL = l.baseURL + dlMatch[1]
	} else {
		l.logger.Warn("no download link on detail page",
			"law", meta.LawNumber, "source_url", meta.SourceURL)
		return connectors.RawDocument{}, fmt.Errorf("no download link on detail page for %s", meta.LawNumber)
	}

	l.logger.Debug("found download link", "law", meta.LawNumber, "url", downloadURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/pdf,*/*")

	resp, err := l.client.Do(req)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("download PDF: %w", err)
	}

	// Follow redirect if present (LKPP returns 302 → presigned S3 URL).
	if resp.StatusCode == http.StatusFound || resp.StatusCode == http.StatusMovedPermanently {
		location := resp.Header.Get("Location")
		resp.Body.Close()
		if location == "" {
			return connectors.RawDocument{}, fmt.Errorf("redirect with no Location for %s", meta.LawNumber)
		}
		if !strings.HasPrefix(location, "http") {
			location = l.baseURL + location
		}
		l.logger.Debug("following redirect", "law", meta.LawNumber, "location", location)

		req, err = http.NewRequestWithContext(ctx, http.MethodGet, location, nil)
		if err != nil {
			return connectors.RawDocument{}, fmt.Errorf("build redirect request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
		resp, err = l.client.Do(req)
		if err != nil {
			return connectors.RawDocument{}, fmt.Errorf("follow redirect: %w", err)
		}
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return connectors.RawDocument{}, fmt.Errorf("download returned status %d for %s", resp.StatusCode, meta.LawNumber)
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
		Filename: fmt.Sprintf("%s.pdf", meta.LawNumber),
	}, nil
}

func (l *LKPPConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (l *LKPPConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return l.Download(ctx, meta)
}

func (l *LKPPConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

		resp, err := l.client.Do(req)
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
