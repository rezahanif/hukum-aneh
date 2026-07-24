package peraturan

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// Cursor types now live in internal/connectors/cursor.go (shared with other connectors).

// PeraturanConnector scrapes peraturan.go.id.
// Covers: UU, PP, Perppu (active/berlaku status only).
// Site has no anti-bot — direct HTTP works.
type PeraturanConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	client  *http.Client
	baseURL string
}

func New(s *scraper.Scraper, logger *slog.Logger) *PeraturanConnector {
	return &PeraturanConnector{
		scraper: s,
		logger:  logger,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
		baseURL: "https://www.peraturan.go.id",
	}
}

func (p *PeraturanConnector) Name() string { return "Peraturan.go.id" }

// SourceType defines a law type to scrape.
type SourceType struct {
	Path     string // /uu, /pp, /perppu
	DocType  string // Undang-Undang (UU), Peraturan Pemerintah (PP), Perppu
	LastPage int    // known last page for active laws
}

var sourceTypes = []SourceType{
	{Path: "/uud", DocType: "UUD 1945", LastPage: 5},
	{Path: "/uu", DocType: "Undang-Undang (UU)", LastPage: 85},
	{Path: "/pp", DocType: "Peraturan Pemerintah (PP)", LastPage: 214},
	{Path: "/perppu", DocType: "Perppu", LastPage: 10},
	{Path: "/perpres", DocType: "Peraturan Presiden (Perpres)", LastPage: 200},
	{Path: "/keppres", DocType: "Keputusan Presiden (Keppres)", LastPage: 200},
	{Path: "/inpres", DocType: "Instruksi Presiden (Inpres)", LastPage: 50},
}

// lawLink matches: /id/uu-no-3-tahun-2026 + title text
var lawLinkRe = regexp.MustCompile(`href="/id/((?:uu|uud|tap-mpr|perppu|pp|perpres|keppres|inpres)-no-\d+-tahun-\d+)"[^>]*title="lihat detail"[^>]*>([^<]*)</a>`)

// statusRe extracts status from law detail page
var statusRe = regexp.MustCompile(`<th[^>]*>Status</th><td>([^<]+)</td>`)

// CheckUpdates polls peraturan.go.id for all active (Berlaku) laws.
// Scrapes pages for each source type with status filter until caught up.
func (p *PeraturanConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)

	cursors := connectors.LoadCursors()

	for _, st := range sourceTypes {
		cursor, hasCursor := cursors.Get(st.DocType)
		p.logger.Info("scraping source type", 
			"type", st.DocType, 
			"last_page_safety_cap", st.LastPage, 
			"has_cursor", hasCursor, 
			"cursor_law", cursor.LastKnownID,
		)

		var newestLawNumber string
		caughtUp := false

		for page := 1; ; page++ {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			// Safety cap check
			if page > st.LastPage {
				p.logger.Warn("exceeded safety page cap without matching cursor, stopping source type", 
					"type", st.DocType, 
					"page", page, 
					"cap", st.LastPage,
				)
				break
			}

			url := fmt.Sprintf("%s%s?PeraturanSearch%%5Bstatus%%5D=Berlaku&page=%d", p.baseURL, st.Path, page)
			html, err := p.fetchURLWithRetry(ctx, url, 3)
			if err != nil {
				p.logger.Warn("fetch page failed after retries", "type", st.DocType, "page", page, "error", err)
				continue
			}

			docs := p.parseListing(html, st.DocType)
			if len(docs) == 0 {
				p.logger.Info("no more laws on page, stopping", "type", st.DocType, "page", page)
				break
			}

			// Track the newest law number (first law seen on page 1)
			if page == 1 && len(docs) > 0 {
				newestLawNumber = docs[0].LawNumber
			}

			for _, d := range docs {
				if hasCursor && d.LawNumber == cursor.LastKnownID {
					p.logger.Info("hit last known law number, caught up", "type", st.DocType, "law_number", d.LawNumber)
					caughtUp = true
					break
				}

				if seen[d.LawNumber] {
					continue
				}
				seen[d.LawNumber] = true
				allDocs = append(allDocs, d)
			}

			if caughtUp {
				break
			}

			// Rate limit — be respectful
			time.Sleep(500 * time.Millisecond)
		}

		// Update cursor with the newest law seen in this run for this source type
		if newestLawNumber != "" {
			if err := cursors.Save(st.DocType, connectors.Cursor{
				LastKnownID: newestLawNumber,
				Timestamp:   time.Now(),
			}); err != nil {
				p.logger.Warn("save cursor failed", "type", st.DocType, "error", err)
			}
		}

		p.logger.Info("source type complete", "type", st.DocType, "total_unique", len(allDocs))
	}

	p.logger.Info("peraturan.go.id scrape complete", "total_laws", len(allDocs))
	return allDocs, nil
}

// parseListing extracts law metadata from listing HTML.
func (p *PeraturanConnector) parseListing(html string, docType string) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta

	matches := lawLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		slug := m[1]
		title := strings.TrimSpace(m[2])

		lawNumber := slugToLawNumber(slug)
		if lawNumber == "" {
			continue
		}

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:    lawNumber,
			Title:        title,
			SourceURL:    fmt.Sprintf("%s/files/%s.pdf", p.baseURL, slug),
			Source:       p.Name(),
			Level:        "national",
			DocumentType: docType,
		})
	}

	return docs
}

// slugToLawNumber converts "uu-no-3-tahun-2026" → "UU No. 3 Tahun 2026"
func slugToLawNumber(slug string) string {
	numRe := regexp.MustCompile(`no-(\d+)-tahun-(\d+)`)
	m := numRe.FindStringSubmatch(slug)
	if m == nil {
		return ""
	}

	prefix := strings.ToUpper(strings.Split(slug, "-")[0])
	switch prefix {
	case "UU":
		prefix = "UU"
	case "UUD":
		prefix = "UUD"
	case "TAP":
		prefix = "TAP MPR"
	case "PERPPU":
		prefix = "Perppu"
	case "PP":
		prefix = "PP"
	case "PERPRES":
		prefix = "Perpres"
	case "KEPPRES":
		prefix = "Keppres"
	case "INPRES":
		prefix = "Inpres"
	}

	return fmt.Sprintf("%s No. %s Tahun %s", prefix, m[1], m[2])
}

// Download fetches the raw PDF for a law.
// Returns error if response is not actually a PDF (some laws have no PDF file).
func (p *PeraturanConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	resp, err := p.fetchURLRaw(ctx, meta.SourceURL)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("download: %w", err)
	}

	mime := resp.Header.Get("Content-Type")
	if mime == "" {
		mime = "application/pdf"
	}

	// Check if response is actually a PDF, not an HTML redirect/error page
	if strings.Contains(mime, "text/html") {
		resp.Body.Close()
		return connectors.RawDocument{}, fmt.Errorf("no PDF available for %s (server returned HTML)", meta.LawNumber)
	}

	return connectors.RawDocument{
		Meta:     meta,
		Content:  resp.Body,
		MimeType: mime,
		Filename: extractFilename(meta.SourceURL),
	}, nil
}

// CheckStatus fetches the law detail page and extracts status.
// Returns "Berlaku" or "Tidak Berlaku".
func (p *PeraturanConnector) CheckStatus(ctx context.Context, slug string) (string, error) {
	url := fmt.Sprintf("%s/id/%s", p.baseURL, slug)
	html, err := p.fetchURL(ctx, url)
	if err != nil {
		return "", fmt.Errorf("fetch detail: %w", err)
	}

	m := statusRe.FindStringSubmatch(html)
	if m == nil {
		return "", fmt.Errorf("status not found for %s", slug)
	}

	return strings.TrimSpace(m[1]), nil
}

func (p *PeraturanConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (p *PeraturanConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return p.Download(ctx, meta)
}

// fetchURL fetches a URL and returns the HTML as string.
func (p *PeraturanConnector) fetchURL(ctx context.Context, url string) (string, error) {
	resp, err := p.fetchURLRaw(ctx, url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body: %w", err)
	}

	return string(body), nil
}

// fetchURLWithRetry retries fetch on failure with backoff.
func (p *PeraturanConnector) fetchURLWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		html, err := p.fetchURL(ctx, url)
		if err == nil {
			return html, nil
		}
		lastErr = err
		p.logger.Warn("fetch retry", "attempt", attempt+1, "url", url, "error", err)
		time.Sleep(time.Duration(attempt+1) * 2 * time.Second)
	}
	return "", lastErr
}

// fetchURLRaw fetches a URL and returns the raw response.
func (p *PeraturanConnector) fetchURLRaw(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("status %d for %s", resp.StatusCode, url)
	}

	return resp, nil
}

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "document"
	}
	return parts[len(parts)-1]
}
