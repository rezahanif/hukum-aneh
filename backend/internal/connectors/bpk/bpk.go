package bpk

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"regexp"
	"strings"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// BPKConnector scrapes peraturan.bpk.go.id via Python curl_cffi bridge.
// Used as fallback source for laws missing PDFs on peraturan.go.id.
// ⚠️ IMPORTANT: BPK is behind Cloudflare — all HTTP must go through
// the Python scraper bridge (curl_cffi with Chrome TLS fingerprint).
// Direct net/http requests WILL be blocked.
type BPKConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	baseURL string
}

const (
	jenisUU    = "8"
	jenisPP    = "10"
	jenisPerpu = "9"
)

func New(s *scraper.Scraper, logger *slog.Logger) *BPKConnector {
	return &BPKConnector{
		scraper: s,
		logger:  logger,
		baseURL: "https://peraturan.bpk.go.id",
	}
}

func (b *BPKConnector) Name() string { return "JDIH BPK" }

var detailLinkRe = regexp.MustCompile(`href="/Details/(\d+)/([a-z0-9-]+)"`)

// CheckUpdates is not used for BPK — it's a fallback source.
func (b *BPKConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	return nil, nil
}

// SearchByLawNumber searches BPK for a specific law via Python bridge.
func (b *BPKConnector) SearchByLawNumber(ctx context.Context, lawNumber string, docType string) (*connectors.DocumentMeta, error) {
	jenis, nomor, tahun, err := parseLawNumber(lawNumber, docType)
	if err != nil {
		return nil, fmt.Errorf("parse law number: %w", err)
	}

	searchURL := fmt.Sprintf("%s/Search?jenis=%s&nomor=%s&tahun=%s",
		b.baseURL, jenis, nomor, tahun)

	// Use Python bridge to bypass Cloudflare
	resp, err := b.scraper.Call(ctx, scraper.ScrapeRequest{
		URL:       searchURL,
		Action:    "search_bpk",
		Source:    b.Name(),
		LawNumber: lawNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("python scraper search: %w", err)
	}

	// Parse HTML from Python response
	dataMap, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response type from scraper")
	}
	html, ok := dataMap["content"].(string)
	if !ok {
		return nil, fmt.Errorf("no HTML content in scraper response")
	}

	// Find detail page link
	matches := detailLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		id := m[1]
		slug := m[2]

		if !slugMatches(slug, nomor, tahun) {
			continue
		}

		// Fetch detail page via Python bridge
		detailURL := fmt.Sprintf("%s/Details/%s/%s", b.baseURL, id, slug)
		metaResp, err := b.scraper.Call(ctx, scraper.ScrapeRequest{
			URL:    detailURL,
			Action: "extract_metadata",
			Source: b.Name(),
		})
		if err != nil {
			b.logger.Warn("fetch detail failed", "id", id, "error", err)
			continue
		}

		metaMap, ok := metaResp.Data.(map[string]interface{})
		if !ok {
			continue
		}

		title, _ := metaMap["title"].(string)
		pdfURL, _ := metaMap["pdf_url"].(string)

		if pdfURL == "" {
			continue
		}

		return &connectors.DocumentMeta{
			LawNumber:    lawNumber,
			Title:        title,
			SourceURL:    pdfURL,
			Source:       b.Name(),
			Level:        "national",
			DocumentType: docType,
		}, nil
	}

	return nil, nil
}

// Download fetches the raw PDF from BPK via Python curl_cffi bridge.
func (b *BPKConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	resp, err := b.scraper.Call(ctx, scraper.ScrapeRequest{
		URL:    meta.SourceURL,
		Action: "download",
		Source: b.Name(),
	})
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("python scraper download: %w", err)
	}

	dataMap, ok := resp.Data.(map[string]interface{})
	if !ok {
		return connectors.RawDocument{}, fmt.Errorf("unexpected download response type")
	}

	contentB64, ok := dataMap["content"].(string)
	if !ok || contentB64 == "" {
		return connectors.RawDocument{}, fmt.Errorf("no content in download response")
	}

	mime, _ := dataMap["mime_type"].(string)
	if mime == "" {
		mime = "application/pdf"
	}

	// Decode base64 content
	rawBytes, err := base64.StdEncoding.DecodeString(contentB64)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("base64 decode: %w", err)
	}

	// Verify it's actually a PDF (not Cloudflare HTML)
	if len(rawBytes) > 4 && string(rawBytes[:4]) != "%PDF" {
		if strings.Contains(string(rawBytes[:min(500, len(rawBytes))]), "cloudflare") {
			return connectors.RawDocument{}, fmt.Errorf("BPK returned Cloudflare challenge instead of PDF")
		}
		b.logger.Warn("BPK response doesn't look like a PDF", "first_bytes", string(rawBytes[:min(20, len(rawBytes))]))
	}

	filename, _ := dataMap["filename"].(string)
	if filename == "" {
		filename = extractFilename(meta.SourceURL)
	}

	return connectors.RawDocument{
		Meta:     meta,
		Content:  io.NopCloser(bytes.NewReader(rawBytes)),
		MimeType: mime,
		Filename: filename,
	}, nil
}

func (b *BPKConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (b *BPKConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return b.Download(ctx, meta)
}

func parseLawNumber(lawNumber string, docType string) (jenis, nomor, tahun string, err error) {
	re := regexp.MustCompile(`No\.\s*(\d+)\s*Tahun\s*(\d+)`)
	m := re.FindStringSubmatch(lawNumber)
	if m == nil {
		return "", "", "", fmt.Errorf("could not parse: %s", lawNumber)
	}
	nomor = m[1]
	tahun = m[2]

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

func slugMatches(slug, nomor, tahun string) bool {
	return strings.Contains(slug, "no-"+nomor+"-tahun-"+tahun)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
