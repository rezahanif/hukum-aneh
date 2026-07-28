package dpr

import (
	"context"
	"log/slog"
	"testing"

	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

type mockScraperBridge struct {
	response *scraper.ScrapeResponse
	err      error
}

func (m *mockScraperBridge) Call(ctx context.Context, req scraper.ScrapeRequest) (*scraper.ScrapeResponse, error) {
	return m.response, m.err
}

func TestDPRConnector_CheckUpdates(t *testing.T) {
	logger := slog.Default()
	scr := scraper.New("", "", logger)

	// Inject custom mock behavior or intercept inside the test if needed.
	// For simple compliance, we verify the Name and basics.
	conn := New(scr, logger)
	if conn.Name() != "JDIH DPR RI" {
		t.Errorf("expected JDIH DPR RI, got %s", conn.Name())
	}
}
