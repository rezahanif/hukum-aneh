package bpk

import (
	"context"
	"log/slog"
	"testing"
)

func TestCheckUpdates_BPK(t *testing.T) {
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

	t.Logf("found %d laws from BPK", len(docs))

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

	t.Logf("first BPK law: %s — %s", d.LawNumber, d.Title)
	t.Logf("pdf url: %s", d.SourceURL)
}
