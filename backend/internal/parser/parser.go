package parser

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/otiai10/gosseract/v2"
)

// Parser converts raw PDF/HTML documents into clean text/markdown.
// Spec §5.3: OCR for scanned PDFs, structure detection (articles, chapters, clauses).
// No AI — deterministic tool calls only.
type Parser struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Parser {
	return &Parser{logger: logger}
}

// ParseResult holds the extracted text and detected structure.
type ParseResult struct {
	TextContent string
	Sections    []Section
	Source      string // "pdf_text" or "pdf_ocr" or "html"
}

type Section struct {
	Type    string // chapter, article, clause, paragraph
	Number  string
	Title   string
	Content string
}

// Parse reads a raw document and extracts text + structure.
// mimeType determines the parsing strategy.
func (p *Parser) Parse(ctx context.Context, r io.Reader, mimeType string, filename string) (*ParseResult, error) {
	switch {
	case strings.Contains(mimeType, "application/pdf"):
		return p.parsePDF(ctx, r, filename)
	case strings.Contains(mimeType, "text/html"):
		return p.parseHTML(ctx, r)
	case strings.Contains(mimeType, "text/plain"):
		return p.parseText(ctx, r)
	default:
		return nil, fmt.Errorf("unsupported mime type: %s", mimeType)
	}
}

// parsePDF extracts text from a PDF.
// Strategy: try text extraction first → if empty, fall back to OCR via tesseract.
func (p *Parser) parsePDF(ctx context.Context, r io.Reader, filename string) (*ParseResult, error) {
	tmpDir, err := os.MkdirTemp("", "hukum-aneh-pdf-")
	if err != nil {
		return nil, fmt.Errorf("mkdir temp: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	pdfPath := filepath.Join(tmpDir, sanitizeFilename(filename))
	f, err := os.Create(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("create temp file: %w", err)
	}
	if _, err := io.Copy(f, r); err != nil {
		f.Close()
		return nil, fmt.Errorf("write temp file: %w", err)
	}
	f.Close()

	// Phase 1: Try text extraction
	text, err := p.extractPDFText(pdfPath)
	if err == nil && len(strings.TrimSpace(text)) > 100 {
		p.logger.Info("pdf text extraction succeeded", "chars", len(text), "filename", filename)
		sections := detectStructure(text)
		return &ParseResult{
			TextContent: text,
			Sections:    sections,
			Source:      "pdf_text",
		}, nil
	}

	p.logger.Info("text extraction empty or failed, falling back to OCR", "filename", filename, "text_len", len(text), "err", err)

	// Phase 2: OCR fallback
	ocrText, err := p.ocrPDF(ctx, pdfPath)
	if err != nil {
		return nil, fmt.Errorf("ocr fallback failed: %w", err)
	}

	p.logger.Info("pdf ocr extraction succeeded", "chars", len(ocrText), "filename", filename)
	sections := detectStructure(ocrText)
	return &ParseResult{
		TextContent: ocrText,
		Sections:    sections,
		Source:      "pdf_ocr",
	}, nil
}

// extractPDFText uses ledongthuc/pdf to extract embedded text.
// Recovers from panics on corrupt PDFs — falls through to OCR.
func (p *Parser) extractPDFText(pdfPath string) (text string, err error) {
	defer func() {
		if r := recover(); r != nil {
			p.logger.Warn("pdf text extraction panicked, falling back to OCR", "panic", r)
			text = ""
			err = fmt.Errorf("pdf panic: %v", r)
		}
	}()

	pdfFile, pdfReader, err := pdf.Open(pdfPath)
	if err != nil {
		return "", fmt.Errorf("open pdf: %w", err)
	}
	defer pdfFile.Close()

	var sb strings.Builder
	numPages := pdfReader.NumPage()
	for i := 1; i <= numPages; i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		text, err := page.GetPlainText(nil)
		if err != nil {
			p.logger.Warn("failed to extract page text", "page", i, "error", err)
			continue
		}
		sb.WriteString(text)
		sb.WriteString("\n\n--- Page Break ---\n\n")
	}

	return sb.String(), nil
}

// ocrPDF converts PDF pages to images via pdftoppm, then OCRs each image with tesseract.
// Uses Indonesian+English language pack.
func (p *Parser) ocrPDF(ctx context.Context, pdfPath string) (string, error) {
	imgDir, err := os.MkdirTemp("", "hukum-aneh-img-")
	if err != nil {
		return "", fmt.Errorf("mkdir img temp: %w", err)
	}
	defer os.RemoveAll(imgDir)

	// Convert PDF to images (PNG, 150 DPI — lower memory, still readable for OCR)
	// ponytail: 300 DPI gives better accuracy but OOMs on 100+ page PDFs. Upgrade to 300 when running on host with >16GB RAM.
	// Limit to first 50 pages to avoid OOM on massive PDFs.
	imgPrefix := filepath.Join(imgDir, "page")
	cmd := exec.CommandContext(ctx, "pdftoppm", "-png", "-r", "150", "-l", "50", pdfPath, imgPrefix)
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("pdftoppm failed: %w; output: %s", err, string(output))
	}

	// Find generated images
	entries, err := os.ReadDir(imgDir)
	if err != nil {
		return "", fmt.Errorf("read img dir: %w", err)
	}

	var images []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".png") {
			images = append(images, filepath.Join(imgDir, e.Name()))
		}
	}

	if len(images) == 0 {
		return "", fmt.Errorf("no images generated from pdftoppm")
	}

	p.logger.Info("ocr processing pages", "count", len(images))

	// OCR each image
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("ind", "eng")

	var sb strings.Builder
	for i, imgPath := range images {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		client.SetImage(imgPath)
		text, err := client.Text()
		os.Remove(imgPath) // free disk immediately after OCR
		if err != nil {
			p.logger.Warn("ocr failed for page", "page", i+1, "error", err)
			continue
		}

		sb.WriteString(text)
		sb.WriteString("\n\n--- Page Break ---\n\n")
	}

	return sb.String(), nil
}

// parseHTML extracts text from HTML documents.
func (p *Parser) parseHTML(ctx context.Context, r io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read html: %w", err)
	}
	text := stripHTMLTags(string(data))
	sections := detectStructure(text)
	return &ParseResult{
		TextContent: text,
		Sections:    sections,
		Source:      "html",
	}, nil
}

func (p *Parser) parseText(ctx context.Context, r io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read text: %w", err)
	}
	text := string(data)
	sections := detectStructure(text)
	return &ParseResult{
		TextContent: text,
		Sections:    sections,
		Source:      "text",
	}, nil
}

// detectStructure parses Indonesian legal document structure.
// Looks for patterns: BAB, PASAL, ayat, huruf.
func detectStructure(text string) []Section {
	var sections []Section
	lines := strings.Split(text, "\n")

	var current *Section
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		upper := strings.ToUpper(trimmed)

		// BAB (Chapter)
		if strings.HasPrefix(upper, "BAB ") {
			if current != nil {
				sections = append(sections, *current)
			}
			current = &Section{
				Type:    "chapter",
				Number:  strings.TrimSpace(trimmed[4:]),
				Content: trimmed + "\n",
			}
			continue
		}

		// PASAL (Article)
		if strings.HasPrefix(upper, "PASAL ") {
			if current != nil {
				sections = append(sections, *current)
			}
			num := strings.TrimSpace(trimmed[6:])
			if idx := strings.IndexAny(num, " \t"); idx > 0 {
				num = num[:idx]
			}
			current = &Section{
				Type:    "article",
				Number:  num,
				Content: trimmed + "\n",
			}
			continue
		}

		// Ayat (paragraph/clause) — starts with (1), (2), etc.
		if len(trimmed) > 2 && trimmed[0] == '(' && trimmed[1] >= '0' && trimmed[1] <= '9' {
			if current != nil {
				current.Content += trimmed + "\n"
			}
			continue
		}

		// Regular text — append to current section
		if current != nil {
			current.Content += trimmed + "\n"
		}
	}

	if current != nil {
		sections = append(sections, *current)
	}

	return sections
}

// stripHTMLTags removes HTML tags and decodes entities.
func stripHTMLTags(s string) string {
	var sb strings.Builder
	inTag := false
	for _, r := range s {
		switch r {
		case '<':
			inTag = true
		case '>':
			inTag = false
			sb.WriteRune(' ')
		default:
			if !inTag {
				sb.WriteRune(r)
			}
		}
	}
	result := sb.String()
	result = strings.ReplaceAll(result, "&amp;", "&")
	result = strings.ReplaceAll(result, "&lt;", "<")
	result = strings.ReplaceAll(result, "&gt;", ">")
	result = strings.ReplaceAll(result, "&quot;", "\"")
	result = strings.ReplaceAll(result, "&#39;", "'")
	result = strings.ReplaceAll(result, "&nbsp;", " ")
	for strings.Contains(result, "  ") {
		result = strings.ReplaceAll(result, "  ", " ")
	}
	for strings.Contains(result, "\n\n\n") {
		result = strings.ReplaceAll(result, "\n\n\n", "\n\n")
	}
	return strings.TrimSpace(result)
}

func sanitizeFilename(s string) string {
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "/", "_")
	return s
}
