package jdihn

import (
	"log/slog"
	"testing"
	"time"
)

func TestNewJdihnConnector(t *testing.T) {
	conn := New(nil, slog.Default())
	if conn.client == nil {
		t.Fatal("expected http client, got nil")
	}
	if conn.client.Timeout != 60*time.Second {
		t.Errorf("expected 60s timeout, got: %s", conn.client.Timeout)
	}
}
