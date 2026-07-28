package kemenkeu

import (
        "context"
        "encoding/json"
        "fmt"
        "io"
        "log/slog"
        "net/http"
        "net/url"
        "strconv"
        "strings"
        "time"

        "github.com/rezahanif/hukum-aneh/backend/internal/connectors"
        "github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// KemenkeuConnector scrapes jdih.kemenkeu.go.id via its JSON API.
// Covers: Peraturan Menteri Keuangan (PMK) and Keputusan Menteri Keuangan (KMK).
//
// The site exposes a JSON API at /api/search with params:
//
//      bentuk=Peraturan+Menteri|Keputusan+Menteri&tahun=2026&page=1&size=50
//
// Response includes full_text_pdf with direct download links like:
//
//      /api/download/{uuid}/{filename}.pdf
type KemenkeuConnector struct {
        scraper *scraper.Scraper
        logger  *slog.Logger
        client  *http.Client
        baseURL string
}

// apiSearchResponse is the top-level JSON response from /api/search.
type apiSearchResponse struct {
        Page struct {
                Size       int `json:"size"`
                Total      int `json:"total"`
                TotalPages int `json:"total_pages"`
                Current    int `json:"current"`
        } `json:"page"`
        Data []apiItem `json:"data"`
}

// apiItem is a single document from the Kemenkeu API.
type apiItem struct {
        Slug        string `json:"slug"`
        Bentuk      string `json:"bentuk"`
        No          int    `json:"no"`
        Tahun       int    `json:"tahun"`
        Nomor       string `json:"nomor"`
        Status      string `json:"status"`
        Judul       string `json:"judul"`
        FullTextPDF string `json:"full_text_pdf"`
}

// bentukFilter maps internal doc type names to API bentuk filter values.
type bentukFilter struct {
        DocType  string // internal DocumentType
        APIName  string // bentuk value for API query
        Prefix   string // law number prefix (PMK/KMK)
}

var bentukFilters = []bentukFilter{
        {DocType: "Peraturan Menteri Keuangan", APIName: "Peraturan Menteri", Prefix: "PMK"},
        {DocType: "Keputusan Menteri Keuangan", APIName: "Keputusan Menteri", Prefix: "KMK"},
}

const maxPages = 20
const apiPageSize = 50

func New(s *scraper.Scraper, logger *slog.Logger) *KemenkeuConnector {
        return &KemenkeuConnector{
                scraper: s,
                logger:  logger,
                client:  &http.Client{Timeout: 60 * time.Second},
                baseURL: "https://jdih.kemenkeu.go.id",
        }
}

func (k *KemenkeuConnector) Name() string { return "JDIH Kemenkeu" }

func (k *KemenkeuConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
        var allDocs []connectors.DocumentMeta
        seen := make(map[string]bool)
        cursors := connectors.LoadCursors()
        cursorUpdates := make(map[string]connectors.Cursor)

        now := time.Now()
        years := []int{now.Year(), now.Year() - 1}

        for _, bf := range bentukFilters {
                cursor, hasCursor := cursors.Get(bf.DocType)
                k.logger.Info("scraping JDIH Kemenkeu",
                        "type", bf.DocType,
                        "has_cursor", hasCursor,
                        "cursor_law", cursor.LastKnownID,
                )

                for _, year := range years {
                        docs, newest, caughtUp, err := k.scrapeYear(ctx, bf, year, cursor, hasCursor)
                        if err != nil {
                                k.logger.Warn("scrape year failed", "type", bf.DocType, "year", year, "error", err)
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
                                cursorUpdates[bf.DocType] = connectors.Cursor{
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

func (k *KemenkeuConnector) scrapeYear(
        ctx context.Context,
        bf bentukFilter,
        year int,
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

                url := fmt.Sprintf("%s/api/search?bentuk=%s&tahun=%d&page=%d&size=%d",
                        k.baseURL, url.QueryEscape(bf.APIName), year, page, apiPageSize)
                k.logger.Debug("fetching kemenkeu API", "url", url, "page", page)

                respData, err := k.fetchJSON(ctx, url)
                if err != nil {
                        return nil, "", false, fmt.Errorf("fetch year=%d page=%d: %w", year, page, err)
                }

                if len(respData.Data) == 0 {
                        k.logger.Debug("no results on page, stopping", "page", page)
                        break
                }

                // Cap pages based on API response.
                if page == 1 && respData.Page.TotalPages > 0 {
                        capped := respData.Page.TotalPages
                        if capped > maxPages {
                                capped = maxPages
                        }
                        k.logger.Debug("API reports total pages", "total", respData.Page.TotalPages, "capped", capped)
                }

                if page == 1 && year == time.Now().Year() && len(respData.Data) > 0 {
                        item := respData.Data[0]
                        newestLaw = fmt.Sprintf("%s No. %d Tahun %d", bf.Prefix, item.No, item.Tahun)
                }

                for _, item := range respData.Data {
                        lawNum := fmt.Sprintf("%s No. %d Tahun %d", bf.Prefix, item.No, item.Tahun)

                        // SourceURL: if full_text_pdf exists, use direct PDF URL so
                        // Download() can fetch it without a second detail-page request.
                        sourceURL := item.FullTextPDF
                        if sourceURL != "" && !strings.HasPrefix(sourceURL, "http") {
                                sourceURL = k.baseURL + sourceURL
                        }
                        // Fallback: use detail page URL if no PDF link.
                        if sourceURL == "" {
                                sourceURL = fmt.Sprintf("%s/dok/%s", k.baseURL, item.Slug)
                        }

                        d := connectors.DocumentMeta{
                                LawNumber:    lawNum,
                                Title:        fmt.Sprintf("%s - %s", item.Nomor, item.Judul),
                                SourceURL:    sourceURL,
                                Source:       "JDIH Kemenkeu",
                                Level:        "national",
                                DocumentType: bf.DocType,
                                PublishedDate: strconv.Itoa(item.Tahun),
                        }
                        docs = append(docs, d)

                        if hasCursor && d.LawNumber == cursor.LastKnownID {
                                k.logger.Info("hit last known law, caught up",
                                        "type", bf.DocType, "year", year, "law", d.LawNumber)
                                caughtUp = true
                                return docs, "", true, nil
                        }
                }

                if caughtUp {
                        break
                }
                time.Sleep(300 * time.Millisecond)
        }

        return docs, newestLaw, caughtUp, nil
}

// Download fetches the PDF. If SourceURL is a direct /api/download/... link
// (set by CheckUpdates), it downloads directly. Otherwise falls back to
// fetching the detail page and extracting the PDF link.
func (k *KemenkeuConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
        // Fast path: SourceURL is already a direct PDF URL from the API.
        if strings.Contains(meta.SourceURL, "/api/download/") {
                return k.downloadDirect(ctx, meta.SourceURL, meta)
        }

        // Slow path: fetch detail page to find PDF link.
        html, err := k.fetchHTML(ctx, meta.SourceURL)
        if err != nil {
                return connectors.RawDocument{}, fmt.Errorf("fetch detail page %s: %w", meta.SourceURL, err)
        }

        pdfPath := extractPDFPathFromHTML(html)
        if pdfPath == "" {
                return connectors.RawDocument{}, fmt.Errorf("no PDF link found on detail page for %s (url: %s)", meta.LawNumber, meta.SourceURL)
        }

        pdfURL := k.baseURL + pdfPath
        return k.downloadDirect(ctx, pdfURL, meta)
}

func (k *KemenkeuConnector) downloadDirect(ctx context.Context, pdfURL string, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
        k.logger.Debug("downloading PDF", "url", pdfURL, "law", meta.LawNumber)

        req, err := http.NewRequestWithContext(ctx, http.MethodGet, pdfURL, nil)
        if err != nil {
                return connectors.RawDocument{}, fmt.Errorf("build PDF request: %w", err)
        }
        req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

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
                return connectors.RawDocument{}, fmt.Errorf("no PDF available for %s (got HTML, content-type: %s)", meta.LawNumber, mime)
        }

        return connectors.RawDocument{
                Meta:     meta,
                Content:  resp.Body,
                MimeType: mime,
                Filename: extractFilename(pdfURL),
        }, nil
}

// extractPDFPathFromHTML extracts /api/download/... from a detail page.
var pdfLinkRe = strings.NewReplacer(
        "href=\"", "",
        "\"", "",
)

func extractPDFPathFromHTML(html string) string {
        // Simple substring search for /api/download/ links.
        idx := strings.Index(html, "/api/download/")
        if idx == -1 {
                return ""
        }
        // Extract the full path up to the next quote or space.
        path := html[idx:]
        end := len(path)
        for i, c := range path {
                if c == '"' || c == '\'' || c == ' ' || c == '>' {
                        end = i
                        break
                }
        }
        return path[:end]
}

func (k *KemenkeuConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
        return raw.Meta, nil
}

func (k *KemenkeuConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
        return k.Download(ctx, meta)
}

func (k *KemenkeuConnector) fetchJSON(ctx context.Context, url string) (*apiSearchResponse, error) {
        var lastErr error
        for attempt := 0; attempt < 3; attempt++ {
                req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
                if err != nil {
                        return nil, err
                }
                req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
                req.Header.Set("Accept", "application/json")

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
                        return nil, err
                }

                var result apiSearchResponse
                if err := json.Unmarshal(body, &result); err != nil {
                        k.logger.Warn("failed to parse API response", "error", err, "body_len", len(body))
                        return nil, fmt.Errorf("parse JSON: %w", err)
                }
                return &result, nil
        }
        return nil, lastErr
}

func (k *KemenkeuConnector) fetchHTML(ctx context.Context, url string) (string, error) {
        var lastErr error
        for attempt := 0; attempt < 3; attempt++ {
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

func extractFilename(url string) string {
        parts := strings.Split(url, "/")
        filename := parts[len(parts)-1]
        if idx := strings.Index(filename, "?"); idx != -1 {
                filename = filename[:idx]
        }
        return filename
}
