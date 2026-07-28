package parser

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"testing"
)

// TestParsePDF_RealDownload downloads a real PDF from peraturan.go.id
// and runs the full parser (text extraction + OCR fallback).
func TestParsePDF_RealDownload(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping real download test in short mode")
	}

	url := "https://www.peraturan.go.id/files/uu-no-1-tahun-2026.pdf"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("download failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("download status: %d", resp.StatusCode)
	}

	p := New(slog.Default())
	result, err := p.Parse(context.Background(), resp.Body, "application/pdf", "uu-no-1-tahun-2026.pdf")
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	if len(result.TextContent) == 0 {
		t.Fatal("text content is empty")
	}

	t.Logf("source: %s", result.Source)
	t.Logf("text length: %d chars", len(result.TextContent))
	t.Logf("sections found: %d", len(result.Sections))

	// Show first 500 chars
	preview := result.TextContent
	if len(preview) > 500 {
		preview = preview[:500]
	}
	t.Logf("preview:\n%s", preview)

	// Show first 3 sections
	for i, s := range result.Sections {
		if i >= 3 {
			break
		}
		t.Logf("section[%d]: type=%s number=%s content_len=%d", i, s.Type, s.Number, len(s.Content))
	}
}

// TestParsePDF_LocalFile tests parser on a pre-downloaded file.
func TestParsePDF_LocalFile(t *testing.T) {
	path := "/tmp/test_law.pdf"
	if _, err := os.Stat(path); err != nil {
		t.Skipf("test file not found: %s", path)
	}

	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer f.Close()

	p := New(slog.Default())
	result, err := p.Parse(context.Background(), f, "application/pdf", "test_law.pdf")
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	if len(result.TextContent) == 0 {
		t.Fatal("text content is empty")
	}

	t.Logf("source: %s", result.Source)
	t.Logf("text length: %d chars", len(result.TextContent))
	t.Logf("sections found: %d", len(result.Sections))

	preview := result.TextContent
	if len(preview) > 500 {
		preview = preview[:500]
	}
	fmt.Printf("preview:\n%s\n", preview)
}

// Keep io.ReadAll reference to prevent unused import in some builds
var _ = io.ReadAll
