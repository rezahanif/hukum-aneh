package parser

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
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

// parsePDF extracts text from a PDF file using ledongthuc/pdf.
// Writes to temp file first — the library requires file path or io.ReaderAt.
func (p *Parser) parsePDF(ctx context.Context, r io.Reader, filename string) (*ParseResult, error) {
	tmpPath := filepath.Join(os.TempDir(), "hukum-aneh-"+filename)
	if err := os.MkdirAll(filepath.Dir(tmpPath), 0755); err != nil {
		return nil, fmt.Errorf("mkdir temp: %w", err)
	}

	f, err := os.Create(tmpPath)
	if err != nil {
		return nil, fmt.Errorf("create temp file: %w", err)
	}
	defer os.Remove(tmpPath)
	defer f.Close()

	if _, err := io.Copy(f, r); err != nil {
		return nil, fmt.Errorf("write temp file: %w", err)
	}
	f.Close()

	// Open with ledongthuc/pdf
	pdfFile, pdfReader, err := pdf.Open(tmpPath)
	if err != nil {
		// ponytail: if PDF is scanned image, this will fail.
		// upgrade: add OCR fallback (tesseract) when text extraction returns empty
		return nil, fmt.Errorf("open pdf: %w", err)
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

	rawText := sb.String()
	sections := detectStructure(rawText)

	return &ParseResult{
		TextContent: rawText,
		Sections:    sections,
	}, nil
}

// parseHTML extracts text from HTML documents.
// ponytail: stub — uses basic tag stripping. Upgrade: use goquery for proper DOM parsing.
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
				Type:   "chapter",
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
			// Stop at first space or newline group
			if idx := strings.IndexAny(num, " \t"); idx > 0 {
				num = num[:idx]
			}
			current = &Section{
				Type:   "article",
				Number: num,
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
// ponytail: minimal implementation. Upgrade: use goquery.
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
	// Decode common entities
	result = strings.ReplaceAll(result, "&amp;", "&")
	result = strings.ReplaceAll(result, "&lt;", "<")
	result = strings.ReplaceAll(result, "&gt;", ">")
	result = strings.ReplaceAll(result, "&quot;", "\"")
	result = strings.ReplaceAll(result, "&#39;", "'")
	result = strings.ReplaceAll(result, "&nbsp;", " ")
	// Collapse whitespace
	for strings.Contains(result, "  ") {
		result = strings.ReplaceAll(result, "  ", " ")
	}
	for strings.Contains(result, "\n\n\n") {
		result = strings.ReplaceAll(result, "\n\n\n", "\n\n")
	}
	return strings.TrimSpace(result)
}
