package peraturan

import (
	"context"
	"log/slog"
	"testing"
)

// TestCheckUpdates_RealSite fetches the real peraturan.go.id and verifies
// that the scraper can parse law listings.
// Skip if offline: go test -short
func TestCheckUpdates_RealSite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping real site test in short mode")
	}

	conn := New(nil, slog.Default())
	docs, err := conn.CheckUpdates(context.Background())
	if err != nil {
		t.Fatalf("CheckUpdates failed: %v", err)
	}

	if len(docs) == 0 {
		t.Fatal("expected at least 1 law, got 0")
	}

	t.Logf("found %d laws", len(docs))

	// Verify first result has required fields
	d := docs[0]
	if d.LawNumber == "" {
		t.Error("LawNumber is empty")
	}
	if d.Title == "" {
		t.Error("Title is empty")
	}
	if d.SourceURL == "" {
		t.Error("SourceURL is empty")
	}
	if !contains(d.SourceURL, ".pdf") {
		t.Errorf("SourceURL should be a PDF link, got: %s", d.SourceURL)
	}

	t.Logf("first law: %s — %s", d.LawNumber, d.Title)
	t.Logf("pdf url: %s", d.SourceURL)
}

func TestSlugToLawNumber(t *testing.T) {
	cases := []struct {
		slug string
		want string
	}{
		{"uu-no-3-tahun-2026", "UU No. 3 Tahun 2026"},
		{"uu-no-1-tahun-2026", "UU No. 1 Tahun 2026"},
		{"uud-no-2-tahun-2026", "UUD No. 2 Tahun 2026"},
		{"pp-no-5-tahun-2025", "PP No. 5 Tahun 2025"},
		{"perppu-no-1-tahun-2025", "Perppu No. 1 Tahun 2025"},
	}

	for _, c := range cases {
		got := slugToLawNumber(c.slug)
		if got != c.want {
			t.Errorf("slugToLawNumber(%q) = %q, want %q", c.slug, got, c.want)
		}
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
