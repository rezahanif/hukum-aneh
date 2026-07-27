package kemnaker

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

// KemnakerConnector scrapes jdih.kemnaker.go.id.
// Covers: Peraturan Menteri/Keputusan Menteri Ketenagakerjaan.
type KemnakerConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	perPage  int
	docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *KemnakerConnector {
	return &KemnakerConnector{
		scraper:  s,
		logger:   logger,
		client:   &http.Client{Timeout: 60 * time.Second},
		baseURL:  "https://jdih.kemnaker.go.id",
		perPage:  15,
		docTypes: []string{"Peraturan Menteri", "Keputusan Menteri"},
	}
}

func (k *KemnakerConnector) Name() string { return "JDIH Kemnaker" }

// resultLinkRe matches detail links from Kemnaker listing page.
// Matches: href="https://jdih.kemnaker.go.id/peraturan/detail/3043/slug-text"
// Capture groups: (1)id (2)slug
var resultLinkRe = regexp.MustCompile(`href="https://jdih\.kemnaker\.go\.id/peraturan/detail/(\d+)/([a-z0-9-]+)"`)

// linkTextRe extracts the visible text content between > and </a>.
var linkTextRe = regexp.MustCompile(`>([^<]+)</a>`)

// slugNumRe extracts number and year from the slug, e.g.
// "undang-undang-nomor-2-tahun-2026" → "2", "2026"
var slugNumRe = regexp.MustCompile(`nomor-(\d+)-tahun-(\d+)`)

func (k *KemnakerConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	now := time.Now()
	years := []int{now.Year(), now.Year() - 1}

	for _, docType := range k.docTypes {
		cursor, hasCursor := cursors.Get(docType)
		k.logger.Info("scraping JDIH Kemnaker",
			"type", docType,
			"has_cursor", hasCursor, "cursor_law", cursor.LastKnownID,
		)

		for _, year := range years {
			docs, newest, caughtUp, err := k.scrapeYear(ctx, docType, year, cursor, hasCursor)
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

func (k *KemnakerConnector) scrapeYear(
	ctx context.Context,
	docType string,
	year int,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	// Bumped from 10 to 20 for more complete coverage.
	const maxPages = 20
	for page := 1; page <= maxPages; page++ {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		url := fmt.Sprintf("%s/peraturan?keyword=&nomor=&tahun=%d&status=&terjemahan=&per_page=%d&hal=%d",
			k.baseURL, year, k.perPage, page)
		k.logger.Debug("fetching kemnaker listing", "url", url, "page", page)

		html, err := k.fetchWithRetry(ctx, url, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
		}

		pageDocs := k.parseListing(html, docType, year)
		if len(pageDocs) == 0 {
			k.logger.Debug("no results on page, stopping", "page", page)
			break
		}

		k.logger.Debug("parsed kemnaker listing page", "page", page, "count", len(pageDocs))

		if page == 1 && len(pageDocs) > 0 {
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

func (k *KemnakerConnector) parseListing(html string, docType string, year int) []connectors.DocumentMeta {
	var docs []connectors.DocumentMeta
	matches := resultLinkRe.FindAllStringSubmatchIndex(html, -1)

	for _, loc := range matches {
		// loc[2:4] is group 1 (id), loc[4:6] is group 2 (slug)
		id := html[loc[2]:loc[3]]
		slug := html[loc[4]:loc[5]]

		// Extract the link text: find the > ... </a> after the href match.
		afterMatch := html[loc[1]:]
		text := ""
		if textMatch := linkTextRe.FindStringSubmatchIndex(afterMatch); textMatch != nil {
			text = strings.TrimSpace(afterMatch[textMatch[2]:textMatch[3]])
		}

		// If no text found, build a title from the slug.
		if text == "" {
			text = strings.ReplaceAll(slug, "-", " ")
		}

		// Extract number and year from the slug first (most reliable).
		slugNum := slugNumRe.FindStringSubmatch(slug)
		var num, yr string
		if slugNum != nil && len(slugNum) >= 3 {
			num = slugNum[1]
			yr = slugNum[2]
		}

		// If slug parsing failed, try the text.
		if num == "" {
			re := regexp.MustCompile(`(?i)(?:nomor|no\.?)\s*(\d+)(?:\s*(?:tahun\s*(\d+)))?`)
			if m := re.FindStringSubmatch(text); m != nil {
				num = m[1]
				yr = m[2]
			}
		}

		if num == "" {
			continue
		}

		if yr == "" {
			yr = strconv.Itoa(year)
		}

		// Classify by text/slug content.
		textLower := strings.ToLower(text)
		slugLower := strings.ToLower(slug)
		isKeputusan := strings.Contains(textLower, "keputusan") || strings.Contains(slugLower, "keputusan")

		if docType == "Keputusan Menteri" && !isKeputusan {
			continue
		}
		if docType == "Peraturan Menteri" && isKeputusan {
			continue
		}

		prefix := "Permenaker"
		if docType == "Keputusan Menteri" {
			prefix = "Kepmenaker"
		}
		lawNum := fmt.Sprintf("%s No. %s Tahun %s", prefix, num, yr)

		// Direct download URL — no need to load detail page.
		sourceURL := fmt.Sprintf("%s/download.php?id=%s", k.baseURL, id)

		docs = append(docs, connectors.DocumentMeta{
			LawNumber:     lawNum,
			Title:         text,
			SourceURL:     sourceURL,
			Source:        k.Name(),
			Level:         "sectoral",
			DocumentType:  docType,
			PublishedDate: yr,
		})
	}
	return docs
}

func (k *KemnakerConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	k.logger.Debug("downloading PDF", "url", meta.SourceURL, "law", meta.LawNumber)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := k.client.Do(req)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("download PDF: %w", err)
	}

	// Follow redirect if the server returns one.
	if resp.StatusCode == http.StatusFound || resp.StatusCode == http.StatusMovedPermanently {
		location := resp.Header.Get("Location")
		resp.Body.Close()
		if location == "" {
			return connectors.RawDocument{}, fmt.Errorf("redirect with no Location for %s", meta.LawNumber)
		}
		if !strings.HasPrefix(location, "http") {
			location = k.baseURL + location
		}
		k.logger.Debug("following redirect", "law", meta.LawNumber, "location", location)

		req, err = http.NewRequestWithContext(ctx, http.MethodGet, location, nil)
		if err != nil {
			return connectors.RawDocument{}, fmt.Errorf("build redirect request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
		resp, err = k.client.Do(req)
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

func (k *KemnakerConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (k *KemnakerConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return k.Download(ctx, meta)
}

func (k *KemnakerConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
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
