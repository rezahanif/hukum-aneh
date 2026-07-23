package bpk

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

type BpkConnector struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
	client  *http.Client
	baseURL string
}

func New(s *scraper.Scraper, logger *slog.Logger) *BpkConnector {
	return &BpkConnector{
		scraper: s,
		logger:  logger,
		client:  &http.Client{},
		baseURL: "https://peraturan.bpk.go.id",
	}
}

func (b *BpkConnector) Name() string { return "JDIH BPK" }

// Regex to capture Details links on BPK home page: href="/Details/347856/peraturan-bpk-no-1-tahun-2026"
var detailsLinkRe = regexp.MustCompile(`href="/Details/(\d+)/([^"]+)"`)

// Regex to capture Download links on BPK details page: href="/Download/411430/Salinan%20PBPK%201%20Tahun%202026.pdf"
var downloadLinkRe = regexp.MustCompile(`href="/Download/(\d+)/([^"]+)"`)

func (b *BpkConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	// Fetch home page listing
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch home page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	html := string(bodyBytes)
	matches := detailsLinkRe.FindAllStringSubmatch(html, -1)

	var docs []connectors.DocumentMeta
	seen := make(map[string]bool)

	for _, m := range matches {
		id := m[1]
		slug := m[2]
		detailPath := fmt.Sprintf("/Details/%s/%s", id, slug)

		// Convert slug to a human readable title
		title := strings.ReplaceAll(slug, "-", " ")
		title = strings.ToUpper(title)

		lawNumber := b.slugToLawNumber(slug)
		if lawNumber == "" {
			continue
		}

		if seen[lawNumber] {
			continue
		}
		seen[lawNumber] = true

		// Fetch details page to get the real PDF download link
		pdfURL, err := b.getDownloadURL(ctx, detailPath)
		if err != nil {
			b.logger.Warn("failed to fetch download url for detail", "path", detailPath, "error", err)
			continue
		}

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNumber,
			Title:         title,
			SourceURL:     pdfURL,
			Source:        b.Name(),
			Level:         "national",
			DocumentType:  "Database lintas jenis",
			PublishedDate: "", // will be parsed later or left blank
		})
	}

	b.logger.Info("bpk check complete", "found", len(docs))
	return docs, nil
}

func (b *BpkConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := b.client.Do(req)
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

func (b *BpkConnector) getDownloadURL(ctx context.Context, detailPath string) (string, error) {
	detailURL := b.baseURL + detailPath
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, detailURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := b.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	m := downloadLinkRe.FindStringSubmatch(string(bodyBytes))
	if m == nil {
		return "", fmt.Errorf("download link not found in detail page")
	}

	// Download link: /Download/411430/Salinan%20PBPK%201%20Tahun%202026.pdf
	return fmt.Sprintf("%s/Download/%s/%s", b.baseURL, m[1], m[2]), nil
}

func (b *BpkConnector) slugToLawNumber(slug string) string {
	parts := strings.Split(slug, "-")
	if len(parts) < 3 {
		return ""
	}
	// e.g. "peraturan-bpk-no-1-tahun-2026" -> "PERATURAN BPK No. 1 Tahun 2026"
	// Find index of "no"
	noIdx := -1
	for i, p := range parts {
		if p == "no" {
			noIdx = i
			break
		}
	}
	if noIdx == -1 || noIdx+1 >= len(parts) {
		return ""
	}

	prefixParts := parts[:noIdx]
	prefix := strings.ToUpper(strings.Join(prefixParts, " "))

	num := parts[noIdx+1]

	// Find year
	year := ""
	for i := noIdx + 2; i < len(parts); i++ {
		if parts[i] == "tahun" && i+1 < len(parts) {
			year = parts[i+1]
			break
		}
	}

	if year == "" {
		return fmt.Sprintf("%s No. %s", prefix, num)
	}

	return fmt.Sprintf("%s No. %s Tahun %s", prefix, num, year)
}

func (b *BpkConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (b *BpkConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return b.Download(ctx, meta)
}

func extractFilename(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "document"
	}
	return parts[len(parts)-1]
}
