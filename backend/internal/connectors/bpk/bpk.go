package bpk

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
)

// BPKConnector scrapes peraturan.bpk.go.id.
// Used as fallback source for laws missing PDFs on peraturan.go.id.
// BPK has advanced search by jenis+nomor+tahun.
type BPKConnector struct {
	logger  *slog.Logger
	client  *http.Client
	baseURL string
}

// jenis codes from BPK search form
const (
	jenisUU     = "8"
	jenisPP     = "10"
	jenisPerpu  = "9"
)

func New(logger *slog.Logger) *BPKConnector {
	return &BPKConnector{
		logger: logger,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
		baseURL: "https://peraturan.bpk.go.id",
	}
}

func (b *BPKConnector) Name() string { return "JDIH BPK" }

// detailLinkRe matches: /Details/134563/uu-no-1-tahun-2020
var detailLinkRe = regexp.MustCompile(`href="/Details/(\d+)/([a-z0-9-]+)"`)

// pdfLinkRe matches: /Download/125355/UU Nomor 1 Tahun 2020.pdf
var pdfLinkRe = regexp.MustCompile(`href="(/Download/\d+/[^"]+\.pdf)"`)

// CheckUpdates is not used for BPK — it's a fallback source.
// Implemented to satisfy Connector interface.
func (b *BPKConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	return nil, nil
}

// SearchByLawNumber searches BPK for a specific law by jenis + nomor + tahun.
// Returns DocumentMeta with PDF download URL if found.
func (b *BPKConnector) SearchByLawNumber(ctx context.Context, lawNumber string, docType string) (*connectors.DocumentMeta, error) {
	jenis, nomor, tahun, err := parseLawNumber(lawNumber, docType)
	if err != nil {
		return nil, fmt.Errorf("parse law number: %w", err)
	}

	// Build search URL
	searchURL := fmt.Sprintf("%s/Search?jenis=%s&nomor=%s&tahun=%s",
		b.baseURL, jenis, nomor, tahun)

	html, err := b.fetchURL(ctx, searchURL)
	if err != nil {
		return nil, fmt.Errorf("search: %w", err)
	}

	// Find detail page link
	matches := detailLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		id := m[1]
		slug := m[2]

		// Verify slug matches our law
		if !slugMatches(slug, nomor, tahun) {
			continue
		}

		// Fetch detail page to get PDF download link
		detailURL := fmt.Sprintf("%s/Details/%s/%s", b.baseURL, id, slug)
		detailHTML, err := b.fetchURL(ctx, detailURL)
		if err != nil {
			b.logger.Warn("fetch detail failed", "id", id, "error", err)
			continue
		}

		pdfMatch := pdfLinkRe.FindStringSubmatch(detailHTML)
		if pdfMatch == nil {
			continue
		}

		pdfURL := b.baseURL + pdfMatch[1]

		// Extract title from detail page
		title := b.extractTitle(detailHTML)

		return &connectors.DocumentMeta{
			LawNumber:    lawNumber,
			Title:        title,
			SourceURL:    pdfURL,
			Source:       b.Name(),
			Level:        "national",
			DocumentType: docType,
		}, nil
	}

	return nil, nil // not found
}

// Download fetches the raw PDF from BPK.
func (b *BPKConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	resp, err := b.fetchURLRaw(ctx, meta.SourceURL)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("download: %w", err)
	}

	mime := resp.Header.Get("Content-Type")
	if mime == "" {
		mime = "application/pdf"
	}

	// Check if response is actually a PDF
	if strings.Contains(mime, "text/html") {
		resp.Body.Close()
		return connectors.RawDocument{}, fmt.Errorf("no PDF available for %s (server returned HTML)", meta.LawNumber)
	}

	filename := extractFilename(meta.SourceURL)

	return connectors.RawDocument{
		Meta:     meta,
		Content:  resp.Body,
		MimeType: "application/pdf",
		Filename: filename,
	}, nil
}

func (b *BPKConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (b *BPKConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return b.Download(ctx, meta)
}

// parseLawNumber extracts jenis, nomor, tahun from law number string.
// "UU No. 1 Tahun 2020" → jenis="8", nomor="1", tahun="2020"
func parseLawNumber(lawNumber string, docType string) (jenis, nomor, tahun string, err error) {
	// Extract nomor and tahun using regex
	re := regexp.MustCompile(`No\.\s*(\d+)\s*Tahun\s*(\d+)`)
	m := re.FindStringSubmatch(lawNumber)
	if m == nil {
		return "", "", "", fmt.Errorf("could not parse: %s", lawNumber)
	}
	nomor = m[1]
	tahun = m[2]

	// Map doc type to jenis code
	switch {
	case strings.Contains(docType, "Undang-Undang") || strings.HasPrefix(lawNumber, "UU "):
		jenis = jenisUU
	case strings.Contains(docType, "Peraturan Pemerintah") || strings.HasPrefix(lawNumber, "PP "):
		jenis = jenisPP
	case strings.Contains(docType, "Perppu") || strings.HasPrefix(lawNumber, "Perppu "):
		jenis = jenisPerpu
	default:
		return "", "", "", fmt.Errorf("unknown doc type: %s", docType)
	}

	return jenis, nomor, tahun, nil
}

// slugMatches checks if BPK slug matches our law number/year.
func slugMatches(slug, nomor, tahun string) bool {
	// slug like "uu-no-1-tahun-2020"
	return strings.Contains(slug, "no-"+nomor+"-tahun-"+tahun)
}

// extractTitle gets the law title from BPK detail page.
func (b *BPKConnector) extractTitle(html string) string {
	re := regexp.MustCompile(`<title>([^<]+)</title>`)
	m := re.FindStringSubmatch(html)
	if m != nil {
		title := strings.TrimSpace(m[1])
		// Remove " - JDIH BPK" suffix
		if idx := strings.Index(title, " - "); idx > 0 {
			title = title[:idx]
		}
		return title
	}
	return ""
}

// fetchURL fetches a URL and returns HTML as string.
func (b *BPKConnector) fetchURL(ctx context.Context, url string) (string, error) {
	resp, err := b.fetchURLRaw(ctx, url)
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

// fetchURLRaw fetches a URL with retry.
func (b *BPKConnector) fetchURLRaw(ctx context.Context, url string) (*http.Response, error) {
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

		resp, err := b.client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Duration(attempt+1) * 2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			lastErr = fmt.Errorf("status %d for %s", resp.StatusCode, url)
			time.Sleep(time.Duration(attempt+1) * 2 * time.Second)
			continue
		}

		return resp, nil
	}
	return nil, lastErr
}

func extractFilename(rawURL string) string {
	parts := strings.Split(rawURL, "/")
	if len(parts) == 0 {
		return "document.pdf"
	}
	name := parts[len(parts)-1]
	decoded, err := url.QueryUnescape(name)
	if err != nil {
		return name
	}
	return decoded
}

// strconv import to avoid unused
var _ = strconv.Atoi
