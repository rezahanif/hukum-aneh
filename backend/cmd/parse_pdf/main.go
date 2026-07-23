// standalone PDF parser — runs as subprocess, exits after each PDF.
// Usage: parse_pdf <input.pdf> <output.txt>
// Returns text to stdout, exits 0 on success.
package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: parse_pdf <input.pdf> <output.txt>")
		os.Exit(1)
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelWarn}))
	p := parser.New(logger)

	f, err := os.Open(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open input: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	result, err := p.Parse(os.Stderr.Context(), f, "application/pdf", inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputPath, []byte(result.TextContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write output: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "parsed %d chars\n", len(result.TextContent))
}
