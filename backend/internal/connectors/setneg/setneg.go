package setneg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// SetnegConnector scrapes jdih.setneg.go.id (JDIH Sekretariat Negara).
// Primary source for Peraturan Presiden (Perpres), Keppres, Inpres, PP, Perpu.
//
// Site migrated to Next.js with JSON API. Data is fetched via POST to
// /api/hukumproduk/produkhukum with JSON body:
//
//	{"tentang":"","p_lihan":"semua","jns":["PERPRES"],"thn":["2026"],
//	 "status":"","terx":"All","sortOrder":"desc","length":10,"start":0}
//
// The API returns: {data: [{idperaturan, no_peraturan, tahun, tentang, jns, nama_jenis, ...}]}
type SetnegConnector struct {
	scraper  *scraper.Scraper
	logger   *slog.Logger
	client   *http.Client
	baseURL  string
	perPage  int
	docTypes []docTypeConfig
}

type docTypeConfig struct {
	Label    string // display label for logging
	JNS      string // API jns filter value
	DocType  string // DocumentType for DocumentMeta
	Prefix   string // law number prefix
}

func New(s *scraper.Scraper, logger *slog.Logger) *SetnegConnector {
	return &SetnegConnector{
		scraper: s,
		logger:  logger,
		client:  &http.Client{Timeout: 30 * time.Second},
		baseURL: "https://jdih.setneg.go.id",
		perPage: 10,
		docTypes: []docTypeConfig{
			{Label: "Perpres", JNS: "PERPRES", DocType: "Peraturan Presiden (Perpres)", Prefix: "Perpres"},
			{Label: "Keppres", JNS: "KEPPRES", DocType: "Keputusan Presiden (Keppres)", Prefix: "Keppres"},
			{Label: "Inpres", JNS: "INPRES", DocType: "Instruksi Presiden (Inpres)", Prefix: "Inpres"},
			{Label: "PP", JNS: "PP", DocType: "Peraturan Pemerintah (PP)", Prefix: "PP"},
			{Label: "Perpu", JNS: "PERPU", DocType: "Perppu", Prefix: "Perppu"},
		},
	}
}

func (s *SetnegConnector) Name() string { return "JDIH Setneg" }

// CheckUpdates polls Setneg API for new laws. Crawls current year + previous year.
func (s *SetnegConnector) CheckUpdates(ctx context.Context) ([]connectors.DocumentMeta, error) {
	var allDocs []connectors.DocumentMeta
	seen := make(map[string]bool)
	cursors := connectors.LoadCursors()
	cursorUpdates := make(map[string]connectors.Cursor)

	now := time.Now()
	years := []int{now.Year(), now.Year() - 1}

	for _, dt := range s.docTypes {
		cursor, hasCursor := cursors.Get(dt.DocType)
		s.logger.Info("scraping JDIH Setneg",
			"type", dt.Label,
			"has_cursor", hasCursor,
			"cursor_law", cursor.LastKnownID,
		)

		for _, year := range years {
			docs, newest, caughtUp, err := s.scrapeYear(ctx, dt, year, cursor, hasCursor)
			if err != nil {
				s.logger.Warn("scrape year failed", "type", dt.Label, "year", year, "error", err)
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
				cursorUpdates[dt.DocType] = connectors.Cursor{
					LastKnownID: newest,
					Timestamp:   time.Now(),
				}
			}
		}
	}

	if len(cursorUpdates) > 0 {
		if err := connectors.SaveAll(cursorUpdates); err != nil {
			s.logger.Warn("batch save cursors failed", "error", err)
		}
	}

	return allDocs, nil
}

// apiRequest is the POST body for /api/hukumproduk/produkhukum.
type apiRequest struct {
	Tentang  string   `json:"tentang"`
	PLihan   string   `json:"p_lihan"`
	JNS      []string `json:"jns"`
	Thn      []string `json:"thn"`
	Status   string   `json:"status"`
	Terx     string   `json:"terx"`
	SortOrder string  `json:"sortOrder"`
	Length   int      `json:"length"`
	Start    int      `json:"start"`
}

// apiResponse is the JSON response from the API.
type apiResponse struct {
	Data []apiItem `json:"data"`
	Jml  int       `json:"jml"` // total count
}

// apiItem represents a single regulation from the API.
type apiItem struct {
	IDPeraturan  string `json:"idperaturan"`
	NoPeraturan  string `json:"no_peraturan"`
	Tahun        string `json:"tahun"`
	Tentang      string `json:"tentang"`
	JNS          string `json:"jns"`
	NamaJenis    string `json:"nama_jenis"`
	Files        string `json:"files"`
	StatusHukum  string `json:"status_hukum"`
}

func (s *SetnegConnector) scrapeYear(
	ctx context.Context,
	dt docTypeConfig,
	year int,
	cursor connectors.Cursor,
	hasCursor bool,
) ([]connectors.DocumentMeta, string, bool, error) {
	var docs []connectors.DocumentMeta
	var newestLaw string
	caughtUp := false

	// Fetch first page to get total count
	start := 0
	for {
		select {
		case <-ctx.Done():
			return nil, "", false, ctx.Err()
		default:
		}

		reqBody := apiRequest{
			Tentang:   "",
			PLihan:    "semua",
			JNS:       []string{dt.JNS},
			Thn:       []string{strconv.Itoa(year)},
			Status:    "",
			Terx:      "All",
			SortOrder: "desc",
			Length:    s.perPage,
			Start:     start,
		}

		bodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			return nil, "", false, fmt.Errorf("marshal request: %w", err)
		}

		url := fmt.Sprintf("%s/api/hukumproduk/produkhukum", s.baseURL)
		respData, err := s.postWithRetry(ctx, url, bodyBytes, 3)
		if err != nil {
			return nil, "", false, fmt.Errorf("fetch year=%d start=%d: %w", year, start, err)
		}

		var apiResp apiResponse
		if err := json.Unmarshal(respData, &apiResp); err != nil {
			s.logger.Warn("failed to parse API response", "error", err, "body_len", len(respData))
			break
		}

		if len(apiResp.Data) == 0 {
			break
		}

		for _, item := range apiResp.Data {
			num, err := strconv.Atoi(item.NoPeraturan)
			if err != nil {
				continue
			}
			lawNum := fmt.Sprintf("%s No. %d Tahun %s", dt.Prefix, num, item.Tahun)

			d := connectors.DocumentMeta{
				LawNumber:     lawNum,
				Title:          item.Tentang,
				SourceURL:      fmt.Sprintf("%s/api/hukumproduk/detailperaturan?jns=%s&no=%s&thn=%s", s.baseURL, dt.JNS, item.NoPeraturan, item.Tahun),
				Source:         "JDIH Setneg",
				Level:          "national",
				DocumentType:   dt.DocType,
				PublishedDate:  item.Tahun,
			}
			docs = append(docs, d)

			if hasCursor && d.LawNumber == cursor.LastKnownID {
				s.logger.Info("hit last known law, caught up",
					"type", dt.Label, "year", year, "law", d.LawNumber)
				caughtUp = true
				return docs, "", true, nil
			}
		}

		// Track newest (first item on first page of current year)
		if start == 0 && year == time.Now().Year() && len(apiResp.Data) > 0 {
			item := apiResp.Data[0]
			num, _ := strconv.Atoi(item.NoPeraturan)
			newestLaw = fmt.Sprintf("%s No. %d Tahun %s", dt.Prefix, num, item.Tahun)
		}

		// Check if we've fetched all
		start += s.perPage
		if start >= apiResp.Jml {
			break
		}

		if caughtUp {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	return docs, newestLaw, caughtUp, nil
}

// Download fetches the PDF for a Setneg regulation.
// The detail API may be WAF-protected, so we construct a direct download URL
// from the detail API endpoint, then fetch the detail page to find PDF links.
func (s *SetnegConnector) Download(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	// Try fetching the detail page to find PDF download links
	// The detail page URL pattern: /peraturan/{jns}-{no}-{tahun}
	detailURL := fmt.Sprintf("%s/peraturan?jns=%s&no=%s&thn=%s",
		s.baseURL, extractJNSFromMeta(meta), extractNoFromMeta(meta), extractYearFromMeta(meta))

	html, err := s.fetchWithRetry(ctx, detailURL, 3)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("fetch detail: %w", err)
	}

	// Look for PDF links in the rendered HTML
	pdfURL := extractPDFURL(html, s.baseURL)
	if pdfURL == "" {
		// Also try the detail API directly
		detailAPIURL := fmt.Sprintf("%s/api/hukumproduk/detailperaturan?jns=%s&no=%s&thn=%s",
			s.baseURL, extractJNSFromMeta(meta), extractNoFromMeta(meta), extractYearFromMeta(meta))
		apiHTML, err := s.fetchWithRetry(ctx, detailAPIURL, 2)
		if err == nil {
			pdfURL = extractPDFURL(apiHTML, s.baseURL)
		}
	}

	if pdfURL == "" {
		return connectors.RawDocument{}, fmt.Errorf("no PDF link found for %s (detailURL: %s)", meta.LawNumber, detailURL)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pdfURL, nil)
	if err != nil {
		return connectors.RawDocument{}, fmt.Errorf("build PDF request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

	resp, err := s.client.Do(req)
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

// extractJNSFromMeta extracts the JNS code from DocumentMeta fields.
func extractJNSFromMeta(meta connectors.DocumentMeta) string {
	dt := strings.ToLower(meta.DocumentType)
	switch {
	case strings.Contains(dt, "perpres"):
		return "PERPRES"
	case strings.Contains(dt, "keppres"):
		return "KEPPRES"
	case strings.Contains(dt, "inpres"):
		return "INPRES"
	case strings.Contains(dt, "peraturan pemerintah"):
		return "PP"
	case strings.Contains(dt, "perppu"):
		return "PERPU"
	default:
		return "PERPRES"
	}
}

// extractNoFromMeta extracts the regulation number from LawNumber.
func extractNoFromMeta(meta connectors.DocumentMeta) string {
	re := regexp.MustCompile(`No\.\s*(\d+)\s*Tahun`)
	if m := re.FindStringSubmatch(meta.LawNumber); len(m) >= 2 {
		return m[1]
	}
	return ""
}

// extractYearFromMeta extracts the year from LawNumber.
func extractYearFromMeta(meta connectors.DocumentMeta) string {
	re := regexp.MustCompile(`Tahun\s+(\d+)`)
	if m := re.FindStringSubmatch(meta.LawNumber); len(m) >= 2 {
		return m[1]
	}
	return ""
}

var pdfLinkRe = regexp.MustCompile(`href="([^"]+\.pdf[^"]*)"`)

func extractPDFURL(html, baseURL string) string {
	if m := pdfLinkRe.FindStringSubmatch(html); m != nil {
		href := m[1]
		if strings.HasPrefix(href, "http") {
			return href
		}
		return baseURL + href
	}
	return ""
}

func (s *SetnegConnector) ExtractMetadata(ctx context.Context, raw connectors.RawDocument) (connectors.DocumentMeta, error) {
	return raw.Meta, nil
}

func (s *SetnegConnector) ExtractDocument(ctx context.Context, meta connectors.DocumentMeta) (connectors.RawDocument, error) {
	return s.Download(ctx, meta)
}

func (s *SetnegConnector) fetchWithRetry(ctx context.Context, url string, maxRetries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

		resp, err := s.client.Do(req)
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

func (s *SetnegConnector) postWithRetry(ctx context.Context, url string, body []byte, maxRetries int) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
		req.Header.Set("Accept", "application/json")

		resp, err := s.client.Do(req)
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

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return respBody, nil
	}
	return nil, lastErr
}
