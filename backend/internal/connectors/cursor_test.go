package connectors

import (
	"os"
	"sync"
	"testing"
	"time"
)

// TestCursorSaveAll_Concurrent proves that SaveAll() serializes multi-key
// writes safely. SaveAll is the recommended API for batch updates
// (e.g. connector scrape runs that touch several keys at once) because
// it does a single Load→modify→Save under the lock.
func TestCursorSaveAll_Concurrent(t *testing.T) {
	os.Remove(CursorFile)
	defer os.Remove(CursorFile)

	const n = 50
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			_ = SaveAll(map[string]Cursor{
				keyFor(idx): {
					LastKnownID: keyFor(idx),
					Timestamp:   time.Now(),
				},
			})
		}(i)
	}
	wg.Wait()

	final := LoadCursors()
	if len(final.Cursors) != n {
		t.Fatalf("expected %d cursors after %d concurrent SaveAll calls, got %d (lost updates)",
			n, n, len(final.Cursors))
	}
}

// TestCursorSave_SingleGoroutine proves the simple Save() flow works for
// the single-goroutine case (which is the current batch -workers=1 mode).
// Concurrent Save() calls without SaveAll are NOT safe — each goroutine
// re-loads, mutates, and saves, and last-writer-wins.
func TestCursorSave_SingleGoroutine(t *testing.T) {
	os.Remove(CursorFile)
	defer os.Remove(CursorFile)

	store := LoadCursors()
	if err := store.Save("test-key", Cursor{
		LastKnownID: "TEST-1",
		Timestamp:   time.Now(),
	}); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	loaded := LoadCursors()
	c, ok := loaded.Get("test-key")
	if !ok || c.LastKnownID != "TEST-1" {
		t.Fatalf("roundtrip failed: %+v", loaded)
	}
}

func keyFor(i int) string {
	if i == 0 {
		return "concurrent-test-0"
	}
	digits := []byte{}
	for i > 0 {
		digits = append([]byte{byte('0' + i%10)}, digits...)
		i /= 10
	}
	return "concurrent-test-" + string(digits)
}
