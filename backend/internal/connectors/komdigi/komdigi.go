package komdigi

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

// KomdigiConnector scrapes jdih.komdigi.go.id.
// Covers: Peraturan Menteri/Keputusan Menteri Komunikasi dan Digital.
type KomdigiConnector struct {
        scraper  *scraper.Scraper
        logger   *slog.Logger
        client   *http.Client
        baseURL  string
        perPage  int
        docTypes []string
}

func New(s *scraper.Scraper, logger *slog.Logger) *KomdigiConnector {
        return &KomdigiConnector{
                scraper:  s,
                logger:   logger,
                client:   &http.Client{Timeout: 60 * time.Second},
                baseURL:  "https://jdih.komdigi.go.id",
                perPage:  10,
                docTypes: []string{"Peraturan Menteri", "Keputusan Menteri"},
        }
}

func (k *KomdigiConnector) Name() string { return "JDIH Komdigi" }

// resultLinkRe matches detail link, e.g. /produk_hukum/view/id/1018/t/keputusan+menteri+komunikasi+dan+digital+nomor+275+tahun+2026
var resultLinkRe = regexp.MustCompile(`href="https://jdih.komdigi.go.id/produk_hukum/view/id/(\d+)/t/([^"]+)"`)

func (k *KomdigiConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
        var allDocs []connectors.DocumentMeta
        seen := make(map[string]bool)
        cursors := connectors.LoadCursors()
        cursorUpdates := make(map[string]connectors.Cursor)

        now := time.Now()
        years := []int{now.Year(), now.Year() - 1}

        for _, docType := range k.docTypes {
                cursor, hasCursor := cursors.Get(docType)
                k.logger.Info("scraping JDIH Komdigi",
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

func (k *KomdigiConnector) scrapeYear(
        ctx context.Context,
        docType string,
        year int,
        cursor connectors.Cursor,
        hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
        var docs []connectors.DocumentMeta
        var newestLaw string
        caughtUp := false

        // Try up to 20 pages per docType/year (Komdigi lists ~15 items per page,
        // so up to 300 candidates per docType per year). In practice, recent
        // years have far fewer items, so empty pages will break the loop early.
        const maxPages = 20
        for page := 1; page <= maxPages; page++ {
                select {
                case <-ctx.Done():
                        return nil, "", false, ctx.Err()
                default:
                }

                // URL structure for Komdigi listing
                url := fmt.Sprintf("%s/produk_hukum/kategori?page=%d&year=%d", k.baseURL, page, year)

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
                // 1.2s between listing pages to respect Komdigi rate limit.
                time.Sleep(1200 * time.Millisecond)
        }

        return docs, newestLaw, caughtUp, nil
}

func (k *KomdigiConnector) parseListing(html string, docType string, year int) []connectors.DocumentMeta {
        var docs []connectors.DocumentMeta
        matches := resultLinkRe.FindAllStringSubmatch(html, -1)
        for _, m := range matches {
                if len(m) < 3 {
                        continue
                }
                id := m[1]
                titleSlug := m[2]
                rawTitle := strings.ReplaceAll(titleSlug, "+", " ")

                // Filter based on document type
                if docType == "Peraturan Menteri" && !strings.Contains(strings.ToLower(rawTitle), "peraturan menteri") {
                        continue
                }
                if docType == "Keputusan Menteri" && !strings.Contains(strings.ToLower(rawTitle), "keputusan menteri") {
                        continue
                }

                lawNum := extractKomdigiLawNumber(rawTitle, docType, year)
                if lawNum == "" {
                        continue
                }

                // Construct direct download link by replacing view with unduh
                downloadURL := fmt.Sprintf("%s/produk_hukum/unduh/id/%s/t/%s", k.baseURL, id, titleSlug)

                docs = append(docs, connectors.DocumentMeta{
                        LawNumber:     lawNum,
                        Title:         rawTitle,
                        SourceURL:     downloadURL,
                        Source:        k.Name(),
                        Level:         "sectoral",
                        DocumentType:  docType,
                        PublishedDate: strconv.Itoa(year),
                })
        }
        return docs
}

func extractKomdigiLawNumber(title string, docType string, year int) string {
        re := regexp.MustCompile(`(?i)(?:nomor|no\.?)\s*(\d+)(?:\s*(?:tahun\s*(\d+)))?`)
        if m := re.FindStringSubmatch(title); m != nil {
                num := m[1]
                yr := m[2]
                if yr == "" {
                        yr = strconv.Itoa(year)
                }
                prefix := "Permenkominfo"
                if strings.Contains(strings.ToLower(title), "digital") {
                        prefix = "Permenkomdigi"
                }
                if docType == "Keputusan Menteri" {
                        prefix = "Kepmenkominfo"
                        if strings.Contains(strings.ToLower(title), "digital") {
                                prefix = "Kepmenkomdigi"
                        }
                }
                return fmt.Sprintf("%s No. %s Tahun %s", prefix, num, yr)
        }
        return ""
}

func (k *KomdigiConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
        // Polite pre-delay to stay under the server's per-IP rate limit
        // (~50 req/min based on X-Ratelimit-* headers observed in production).
        select {
        case <-ctx.Done():
                return connectors.RawDocument{}, ctx.Err()
        case <-time.After(1200 * time.Millisecond):
        }

        const maxRetries = 4
        var lastErr error
        for attempt := 1; attempt <= maxRetries; attempt++ {
                req, err := http.NewRequestWithContext(ctx, http.MethodGet, meta.SourceURL, nil)
                if err != nil {
                        return connectors.RawDocument{}, fmt.Errorf("build request: %w", err)
                }
                req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
                req.Header.Set("Accept", "application/pdf,*/*")
                req.Header.Set("Accept-Language", "en-US,en;q=0.9")

                resp, err := k.client.Do(req)
                if err != nil {
                        lastErr = fmt.Errorf("download PDF: %w", err)
                        if attempt < maxRetries {
                                time.Sleep(time.Duration(attempt*2) * time.Second)
                                continue
                        }
                        break
                }

                // Handle rate-limit (429) and 5xx with retry-after backoff.
                if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
                        retryAfter := 10 * time.Second
                        if ra := resp.Header.Get("Retry-After"); ra != "" {
                                if secs, err := time.ParseDuration(ra + "s"); err == nil {
                                        retryAfter = secs
                                }
                        }
                        resp.Body.Close()
                        lastErr = fmt.Errorf("status %d for %s (will retry after %s)",
                                resp.StatusCode, meta.LawNumber, retryAfter)
                        if attempt < maxRetries {
                                select {
                                case <-ctx.Done():
                                        return connectors.RawDocument{}, ctx.Err()
                                case <-time.After(retryAfter):
                                }
                                continue
                        }
                        break
                }

                if resp.StatusCode != http.StatusOK {
                        resp.Body.Close()
                        lastErr = fmt.Errorf("unexpected status %d for %s", resp.StatusCode, meta.LawNumber)
                        break
                }

                mime := resp.Header.Get("Content-Type")
                if mime == "" {
                        mime = "application/pdf"
                }
                if strings.Contains(mime, "text/html") {
                        // Could be a soft rate-limit page returned with 200. Sniff first bytes.
                        buf := make([]byte, 5)
                        n, _ := io.ReadFull(resp.Body, buf)
                        _ = resp.Body.Close()
                        if n >= 5 && string(buf[:5]) == "%PDF-" {
                                // Server lied about Content-Type — it's actually a PDF. Re-fetch cleanly.
                                lastErr = fmt.Errorf("server returned PDF but mislabeled as HTML for %s", meta.LawNumber)
                        } else {
                                lastErr = fmt.Errorf("no PDF available for %s (got HTML, status 200)", meta.LawNumber)
                        }
                        if attempt < maxRetries {
                                select {
                                case <-ctx.Done():
                                        return connectors.RawDocument{}, ctx.Err()
                                case <-time.After(10 * time.Second):
                                }
                                continue
                        }
                        break
                }

                return connectors.RawDocument{
                        Meta:     meta,
                        Content:  resp.Body,
                        MimeType: mime,
                        Filename: fmt.Sprintf("%s.pdf", meta.LawNumber),
                }, nil
        }
        return connectors.RawDocument{}, lastErr
}

func (k *KomdigiConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
        return raw.Meta, nil
}

func (k *KomdigiConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
        return k.Download(ctx, meta)
}

func (k *KomdigiConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
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
