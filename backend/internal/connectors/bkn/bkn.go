package bkn

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

// BKNConnector scrapes jdih.bkn.go.id.
// Covers: Peraturan Badan Kepegawaian Negara.
type BKNConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	perPage  int
	docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *BKNConnector {
	return &BKNConnector{
		scraper:  s,
		logger:   logger,
		client:   &http.Client{Timeout: 60 * time.Second},
		baseURL:  "https://jdih.bkn.go.id",
		perPage:  5,
		docTypes: []string{"Peraturan Badan Kepegawaian Negara"},
	}
}

func (b *BKNConnector) Name() string { return "JDIH BKN" }

// resultLinkRe matches BKN law links, e.g. /Detail_peraturan/breaking/2643
var resultLinkRe = regexp.MustCompile(`href="(/Detail_peraturan/breaking/|/index.php/Detail_peraturan/breaking/)(\d+)"`)

// pdfLinkRe matches direct PDF link on BKN detail page, e.g. http://jdih.bkn.go.id/common/dokumen/...pdf
var pdfLinkRe = regexp.MustCompile(`href="(https?://jdih\.bkn\.go\.id/common/dokumen/[^"]+\.pdf)"`)

func (b *BKNConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	for _, docType := range b.docTypes {
		cursor, hasCursor := cursors.Get(docType)
		b.logger.Info("scraping JDIH BKN",
			"type", docType,
			"has_cursor", hasCursor, "cursor_law", cursor.LastKnownID,
		)

		docs, newest, caughtUp, err := b.scrapeRecent(ctx, docType, cursor, hasCursor)
		if err != nil {
			b.logger.Warn("scrape failed", "type", docType, "error", err)
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

	if len(cursorUpdates) > 0 {
		if err := connectors.SaveAll(cursorUpdates); err != nil {
			b.logger.Warn("batch save cursors failed", "error", err)
		}
	}

	return allDocs, nil
}

func (b *BKNConnector) scrapeRecent(
	ctx context.Context,
	docType string,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	// pagination offsets are multiples of 5: 0, 5, 10, 15...
	const maxPages = 5
	for page := 0; page < maxPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		offset := page * b.perPage
		url := fmt.Sprintf("%s/index.php/Home/peraturan/%d", b.baseURL, offset)

		html, err := b.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch offset=%d: %w", offset, err)
		}

		pageDocs := b.parseListing(html, docType)
		if len(pageDocs) == 0 {
			break
		}

		if page == 0 && len(pageDocs) > 0 {
			newestLaw = pageDocs[0].LawNumber
		}

		for _, d := range pageDocs {
			if hasCursor && d.LawNumber == cursor.LastKnownID {
				b.logger.Info("hit last known law, caught up",
					"type", docType, "law", d.LawNumber)
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

func (b *BKNConnector) parseListing(html string, docType string) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := resultLinkRe.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		if len(m) < 3 {
			continue
		}
		id := m[2]

		detailURL := fmt.Sprintf("%s/index.php/Detail_peraturan/breaking/%s", b.baseURL, id)
		lawNum := fmt.Sprintf("Peraturan BKN No. %s", id)

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNum,
			Title:         fmt.Sprintf("Peraturan Badan Kepegawaian Negara Nomor %s", id),
			SourceURL:     detailURL,
			Source:        b.Name(),
			Level:         "sectoral",
			DocumentType:  docType,
			PublishedDate: "",
		})
	}
	return docs
}

func (b *BKNConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Fetch the detail page to extract PDF URL
	html, err := b.fetchWithRetry(ctx, meta.SourceURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail page: %w", err)
	}

	m := pdfLinkRe.FindStringSubmatch(html)
	if m == nil {
		return connectors.RawDocument{}, fmt.Errorf("no PDF link found on detail page for %s", meta.LawNumber)
	}

	pdfURL := m[1]

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pdfURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := b.client.Do(req)
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
		Filename: fmt.Sprintf("%s.pdf", meta.LawNumber),
	}, nil
}

func (b *BKNConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (b *BKNConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return b.Download(ctx, meta)
}

func (b *BKNConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

		resp, err := b.client.Do(req)
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
