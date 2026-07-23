package scraper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"time"
)

// Scraper bridges Go → Python subprocess for TLS-fingerprinted scraping.
// Python handles anti-bot evasion (curl_cffi/patchright); Go handles orchestration.
// Spec §7: scraper is deterministic, not AI.
type Scraper struct {
	pythonPath string
	scriptPath string
	logger     *slog.Logger
	timeout    time.Duration
}

func New(pythonPath, scriptPath string, logger *slog.Logger) *Scraper {
	return &Scraper{
		pythonPath: pythonPath,
		scriptPath: scriptPath,
		logger:     logger,
		timeout:    60 * time.Second,
	}
}

// ScrapeRequest is the JSON payload sent to the Python script via stdin.
type ScrapeRequest struct {
	URL    string `json:"url"`
	Action string `json:"action"` // "check_updates", "download", "extract_metadata"
	Source string `json:"source"`
}

// ScrapeResponse is the JSON returned by the Python script via stdout.
type ScrapeResponse struct {
	Success   bool        `json:"success"`
	Error     string      `json:"error,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Documents []ScrapedDoc `json:"documents,omitempty"`
}

type ScrapedDoc struct {
	LawNumber    string `json:"law_number"`
	Title        string `json:"title"`
	SourceURL    string `json:"source_url"`
	Source       string `json:"source"`
	Level        string `json:"level"`
	DocumentType string `json:"document_type"`
	PublishedDate string `json:"published_date"`
	Content      string `json:"content,omitempty"` // raw HTML/PDF bytes (base64 if binary)
	MimeType     string `json:"mime_type,omitempty"`
}

// Call invokes the Python scraper subprocess with the given request.
func (s *Scraper) Call(ctx context.Context, req ScrapeRequest) (*ScrapeResponse, error) {
	input, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	callCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	cmd := exec.CommandContext(callCtx, s.pythonPath, s.scriptPath)
	cmd.Stdin = bytes.NewReader(input)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	s.logger.Debug("calling python scraper", "action", req.Action, "url", req.URL)

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("python scraper failed: %w; stderr: %s", err, stderr.String())
	}

	var resp ScrapeResponse
	if err := json.Unmarshal(stdout.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf("parse scraper response: %w; raw: %s", err, stdout.String())
	}

	if !resp.Success {
		return &resp, fmt.Errorf("scraper error: %s", resp.Error)
	}

	return &resp, nil
}

// SetTimeout overrides the default subprocess timeout.
func (s *Scraper) SetTimeout(d time.Duration) {
	s.timeout = d
}
