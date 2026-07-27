package kpu

import (
        "context"
        "fmt"
        "io"
        "log/slog"
        "net/http"
        "regexp"
        "strconv"
        "strings"
        "sync"
        "time"

        "github.com/rezahanif/hukum-aneh/backend/internal/connectors"
        "github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// KPUConnector scrapes jdih.kpu.go.id.
// Covers: Peraturan KPU and Keputusan KPU.
type KPUConnector struct {
        scraper  *scraper.Scraper
        logger   *slog.Logger
        client   *http.Client
        baseURL  string
        docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *KPUConnector {
        return &KPUConnector{
                scraper:  s,
                logger:   logger,
                client:   &http.Client{Timeout: 60 * time.Second},
                baseURL:  "https://jdih.kpu.go.id",
                docTypes: []string{"Peraturan KPU", "Keputusan KPU"},
        }
}

func (k *KPUConnector) Name() string { return "JDIH KPU" }

// resultLinkRe matches detail link on KPU JDIH, e.g. /peraturan-kpu/detail/3tgf5A2viTbUzsFSjXTCqWtzL0FSbWwwS3dkNE1lQ3Jub3R5N0E9PQ
// or /keputusan-kpu/detail/...
var resultLinkRe = regexp.MustCompile(`href="https://jdih.kpu.go.id/(peraturan-kpu|keputusan-kpu)/detail/([a-zA-Z0-9_-]+)"`)

// downloadLinkRe matches download link on detail page, e.g. /peraturan-kpu/download/462
var downloadLinkRe = regexp.MustCompile(`href="https://jdih.kpu.go.id/(peraturan-kpu|keputusan-kpu)/download/(\d+)"`)

func (k *KPUConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
        var allDocs []connectors.DocumentMeta
        // Dedup by numeric ID, not by LawNumber. The KPU listing's ?year=N param
        // is ignored by the server — same numeric IDs are returned regardless of
        // year — so iterating multiple years produces the same set of PDFs with
        // only the year suffix in LawNumber changing. Dedup by the stable numeric
        // ID (extracted from SourceURL) to avoid downloading the same PDF twice.
        seenNumericID := make(map[string]bool)
        cursors := connectors.LoadCursors()
        cursorUpdates := make(map[string]connectors.Cursor)

        // Single year — the year filter is ignored by the server, so iterating
        // multiple years just produces duplicates. Use current year for the
        // LawNumber suffix (purely cosmetic — the actual PDF is the same).
        year := time.Now().Year()

        for _, docType := range k.docTypes {
                cursor, hasCursor := cursors.Get(docType)
                k.logger.Info("scraping JDIH KPU",
                        "type", docType,
                        "has_cursor", hasCursor, "cursor_law", cursor.LastKnownID,
                )

                docs, newest, caughtUp, err := k.scrapeYear(ctx, docType, year, cursor, hasCursor)
                if err != nil {
                        k.logger.Warn("scrape year failed", "type", docType, "year", year, "error", err)
                        continue
                }

                for _, d := range docs {
                        // Extract numeric ID from SourceURL for dedup.
                        // URL format: https://jdih.kpu.go.id/<path>/download/<numericID>
                        parts := strings.Split(d.SourceURL, "/")
                        numericID := ""
                        if len(parts) > 0 {
                                numericID = parts[len(parts)-1]
                        }
                        if numericID == "" || seenNumericID[numericID] {
                                continue
                        }
                        seenNumericID[numericID] = true
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
                        k.logger.Warn("batch save cursors failed", "error", err)
                }
        }

        return allDocs, nil
}

func (k *KPUConnector) scrapeYear(
        ctx context.Context,
        docType string,
        year int,
        cursor connectors.Cursor,
        hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
        var docs []connectors.DocumentMeta
        var newestLaw string
        caughtUp := false

        pathSegment := "peraturan-kpu"
        pageParam := "page_peraturan"
        if docType == "Keputusan KPU" {
                pathSegment = "keputusan-kpu"
                pageParam = "page_keputusan"
        }

        // Try up to 25 pages per docType. Each page has 10 entries, so up to
        // 250 Peraturan KPU + 250 Keputusan KPU candidates (in practice KPU has
        // far fewer Keputusan, so most keputusan-kpu pages will return empty).
        const maxPages = 25
        for page := 1; page <= maxPages; page++ {
                select {
                case <-ctx.Done():
                        return nil, "", false, ctx.Err()
                default:
                }

                url := fmt.Sprintf("%s/%s?%s=%d&year=%d", k.baseURL, pathSegment, pageParam, page, year)

                html, err := k.fetchWithRetry(ctx, url, 3)
                if err != nil {
                        return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
                }

                pageDocs := k.parseListing(html, docType, year)
                if len(pageDocs) == 0 {
                        break
                }

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

func (k *KPUConnector) parseListing(html string, docType string, year int) []connectors.DocumentMeta {
        type result struct {
                idx       int
                numericID string
                dlURL     string
                hashShort string
                err       error
        }

        // First pass: extract all (pathSegment, hashID) pairs from the listing.
        type listingItem struct {
                pathSegment string
                hashID      string
                detailURL   string
        }
        var items []listingItem
        matches := resultLinkRe.FindAllStringSubmatch(html, -1)
        for _, m := range matches {
                if len(m) < 3 {
                        continue
                }
                items = append(items, listingItem{
                        pathSegment: m[1],
                        hashID:      m[2],
                        detailURL:   fmt.Sprintf("%s/%s/detail/%s", k.baseURL, m[1], m[2]),
                })
        }
        if len(items) == 0 {
                return nil
        }

        // Second pass: fetch detail pages CONCURRENTLY to extract stable numeric IDs.
        // The hashID in the listing URL is a session token that changes every scrape,
        // but the numeric download ID is stable. We use it as the LawNumber so
        // skip-existing works across runs.
        results := make([]result, len(items))
        const concurrency = 8
        sem := make(chan struct{}, concurrency)
        var wg sync.WaitGroup
        ctx := context.Background()

        for i, item := range items {
                wg.Add(1)
                go func(i int, item listingItem) {
                        defer wg.Done()
                        sem <- struct{}{}
                        defer func() { <-sem }()

                        hashShort := item.hashID
                        if len(hashShort) > 8 {
                                hashShort = hashShort[:8]
                        }

                        detailHTML, err := k.fetchWithRetry(ctx, item.detailURL, 2)
                        if err != nil {
                                results[i] = result{idx: i, hashShort: hashShort, err: err}
                                return
                        }
                        dlMatch := downloadLinkRe.FindStringSubmatch(detailHTML)
                        if dlMatch == nil {
                                results[i] = result{idx: i, hashShort: hashShort, err: fmt.Errorf("no download link on detail page")}
                                return
                        }
                        results[i] = result{
                                idx:       i,
                                numericID: dlMatch[2],
                                dlURL:     fmt.Sprintf("%s/%s/download/%s", k.baseURL, item.pathSegment, dlMatch[2]),
                                hashShort: hashShort,
                        }
                }(i, item)
        }
        wg.Wait()

        // Third pass: build DocumentMeta for items that succeeded.
        prefix := "PKPU"
        if docType == "Keputusan KPU" {
                prefix = "Keputusan_KPU"
        }
        var docs []connectors.DocumentMeta
        for _, r := range results {
                if r.err != nil {
                        k.logger.Warn("fetch detail page failed in CheckUpdates, skipping",
                                "type", docType, "year", year, "hashID", r.hashShort, "error", r.err)
                        continue
                }
                lawNum := fmt.Sprintf("%s_%s_%s", prefix, r.numericID, strconv.Itoa(year))
                docs = append(docs, connectors.DocumentMeta{
                        LawNumber:     lawNum,
                        Title:         fmt.Sprintf("%s #%s (%s)", docType, r.numericID, strconv.Itoa(year)),
                        SourceURL:     r.dlURL,
                        Source:        k.Name(),
                        Level:         "sectoral",
                        DocumentType:  docType,
                        PublishedDate: strconv.Itoa(year),
                })
        }
        return docs
}

func min(a, b int) int {
        if a < b {
                return a
        }
        return b
}

func (k *KPUConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
        // SourceURL is already the direct download URL (set by CheckUpdates
        // after fetching the detail page). No need to fetch the detail page again.
        downloadURL := meta.SourceURL

        // Extract the numeric ID from the URL for the stable filename.
        // URL format: https://jdih.kpu.go.id/<path>/download/<numericID>
        parts := strings.Split(downloadURL, "/")
        numericID := ""
        if len(parts) > 0 {
                numericID = parts[len(parts)-1]
        }

        prefix := "PKPU"
        if meta.DocumentType == "Keputusan KPU" {
                prefix = "Keputusan_KPU"
        }
        year := meta.PublishedDate
        if year == "" {
                year = strconv.Itoa(time.Now().Year())
        }
        stableFilename := fmt.Sprintf("%s_%s_%s.pdf", prefix, numericID, year)

        req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
        if err != nil {
                return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
        }
        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

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
                return connectors.RawDocument{}, fmt.Errorf("no PDF available for %s (got HTML)", meta.LawNumber)
        }

        return connectors.RawDocument{
                Meta:     meta,
                Content:  resp.Body,
                MimeType: mime,
                Filename: stableFilename,
        }, nil
}

func (k *KPUConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
        return raw.Meta, nil
}

func (k *KPUConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
        return k.Download(ctx, meta)
}

func (k *KPUConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
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
