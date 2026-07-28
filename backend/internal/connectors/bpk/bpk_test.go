package bpk

import (
	"context"
	"log/slog"
	"testing"

	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

func TestSearchByLawNumber_BPK(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping real site test in short mode")
	}

	logger := slog.Default()
	scr := scraper.New("python3", "../../../python/scraper/scrape.py", logger)
	conn := New(scr, logger)
	meta, err := conn.SearchByLawNumber(context.Background(), "UU No. 1 Tahun 2020", "Undang-Undang (UU)")
	if err != nil {
		t.Fatalf("SearchByLawNumber failed: %v", err)
	}

	if meta == nil {
		t.Fatal("expected to find UU No. 1 Tahun 2020 on BPK, got nil")
	}

	t.Logf("found BPK law: %s — %s", meta.LawNumber, meta.Title)
	t.Logf("pdf url: %s", meta.SourceURL)

	if meta.LawNumber != "UU No. 1 Tahun 2020" {
		t.Errorf("expected LawNumber 'UU No. 1 Tahun 2020', got: %s", meta.LawNumber)
	}
	if meta.SourceURL == "" {
		t.Error("SourceURL is empty")
	}
}
