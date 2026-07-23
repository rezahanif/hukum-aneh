package peraturan

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strings"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// PeraturanConnector scrapes peraturan.go.id.
// Covers: UUD 1945, TAP MPR, UU, Perppu, PP.
// Site has no anti-bot — direct HTTP works.
// Python scraper still wired for future sources that need TLS fingerprinting.
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
		client:  &http.Client{},
		baseURL: "https://www.peraturan.go.id",
	}
}

func (p *PeraturanConnector) Name() string { return "Peraturan.go.id" }

// lawLink matches: /id/uu-no-3-tahun-2026 + title text
var lawLinkRe = regexp.MustCompile(`href="/id/((?:uu|uud|tap-mpr|perppu|pp)-no-\d+-tahun-\d+)"[^>]*title="lihat detail"[^>]*>([^<]*)</a>`)

// pdfLink matches: /files/uu-no-3-tahun-2026.pdf
var pdfLinkRe = regexp.MustCompile(`href="/files/((?:uu|uud|tap-mpr|perppu|pp)-no-\d+-tahun-\d+)\.pdf"`)

// CheckUpdates polls peraturan.go.id for new/changed laws.
func (p *PeraturanConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	// Fetch UU listing page
	html, err := p.fetchPage(ctx, "/uu")
	if err != nil {
		return nil, fmt.Errorf("fetch uu listing: %w", err)
	}

	docs := p.parseListing(html, "Undang-Undang (UU)")

	// Also fetch PP listing
	ppHTML, err := p.fetchPage(ctx, "/pp")
	if err != nil {
		p.logger.Warn("failed to fetch pp listing", "error", err)
	} else {
		docs = append(docs, p.parseListing(ppHTML, "Peraturan Pemerintah (PP)")...)
	}

	// Fetch Perppu
	perppuHTML, err := p.fetchPage(ctx, "/perppu")
	if err != nil {
		p.logger.Warn("failed to fetch perppu listing", "error", err)
	} else {
		docs = append(docs, p.parseListing(perppuHTML, "Perppu")...)
	}

	// Deduplicate by law number
	seen := make(map[string]bool)
	var result []connectors.DocumentMeta
	for _, d := range docs {
		if seen[d.LawNumber] {
			continue
		}
		seen[d.LawNumber] = true
		result = append(result, d)
	}

	p.logger.Info("peraturan.go.id check complete", "found", len(result))
	return result, nil
}

// parseListing extracts law metadata from listing HTML.
func (p *PeraturanConnector) parseListing(html string, docType string) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta

	matches := lawLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		slug := m[1]    // e.g. uu-no-3-tahun-2026
		title := strings.TrimSpace(m[2])

		// Convert slug to law number: "uu-no-3-tahun-2026" → "UU No. 3 Tahun 2026"
		lawNumber := slugToLawNumber(slug)
		if lawNumber == "" {
			continue
		}

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNumber,
			Title:         title,
			SourceURL:     fmt.Sprintf("%s/files/%s.pdf", p.baseURL, slug),
			Source:        p.Name(),
			Level:         "national",
			DocumentType:  docType,
		})
	}

	return docs
}

// slugToLawNumber converts "uu-no-3-tahun-2026" → "UU No. 3 Tahun 2026"
func slugToLawNumber(slug string) string {
	parts := strings.Split(slug, "-")
	if len(parts) < 5 {
		return ""
	}
	// parts: [type, no, num, tahun, year]
	prefix := strings.ToUpper(parts[0])
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
	default:
		prefix = strings.ToUpper(parts[0])
	}

	// Find the number and year
	// Pattern: type-no-N-tahun-YYYY
	numRe := regexp.MustCompile(`no-(\d+)-tahun-(\d+)`)
	m := numRe.FindStringSubmatch(slug)
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s No. %s Tahun %s", prefix, m[1], m[2])
}

// Download fetches the raw PDF for a law.
func (p *PeraturanConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	resp, err := p.client.Do(req)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("download: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return connectors.RawDocument{}, fmt.Errorf("download failed: status %d", resp.StatusCode)
	}

	mime := resp.Header.Get("Content-Type")
	if mime == "" {
		mime = "application/pdf"
	}

	filename := extractFilename(meta.SourceURL)

	return connectors.RawDocument{
		Meta:     meta,
		Content:  resp.Body,
		MimeType: mime,
		Filename: filename,
	}, nil
}

func (p *PeraturanConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (p *PeraturanConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return p.Download(ctx, meta)
}

// fetchPage fetches a page from peraturan.go.id and returns the HTML.
func (p *PeraturanConnector) fetchPage(ctx context.Context, path string) (string, error) {
	url := p.baseURL + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	resp, err := p.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetch %s: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("fetch %s: status %d", path, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}

	return string(body), nil
}

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "document"
	}
	return parts[len(parts)-1]
}
