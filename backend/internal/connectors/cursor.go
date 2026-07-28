package connectors

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// CursorFile is the default location for all connector cursors.
// Connectors share this file via namespaced keys.
const CursorFile = "backend/configs/scrape_cursor.json"

// cursorMu serializes Save operations within this process.
// Cheap and safe; protects against concurrent goroutines in a single batch run.
var cursorMu sync.Mutex

// Cursor records the last known position for a given source key.
// Use SourceType (e.g. "Undang-Undang (UU)") as the key.
type Cursor struct {
	LastKnownID    string            `json:"last_known_id"`           // last known law number or document ID
	LastKnownTitle string            `json:"last_known_title,omitempty"`
	Timestamp      time.Time         `json:"timestamp"`
	Extra          map[string]string `json:"extra,omitempty"` // connector-specific extras
}

// CursorStore is the on-disk cursor map keyed by source identifier.
type CursorStore struct {
	Cursors map[string]Cursor `json:"cursors"`
}

// LoadCursors reads the shared cursor file. Returns empty store if file missing.
func LoadCursors() CursorStore {
	s := CursorStore{Cursors: make(map[string]Cursor)}
	data, err := os.ReadFile(CursorFile)
	if err != nil {
		return s
	}
	_ = json.Unmarshal(data, &s)
	if s.Cursors == nil {
		s.Cursors = make(map[string]Cursor)
	}
	return s
}

// SaveCursors writes the cursor map atomically under a process-wide lock.
// Cross-process safety: writes go to a .tmp file then rename, so a partial
// write cannot corrupt the cursor file. In-process safety: cursorMu prevents
// two goroutines from clobbering each other's updates when running with
// -workers > 1. For true cross-process safety, wrap the call site in
// external file locking (e.g. flock) when running multiple binaries.
func SaveCursors(s CursorStore) error {
	cursorMu.Lock()
	defer cursorMu.Unlock()

	if err := os.MkdirAll(filepath.Dir(CursorFile), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	tmp := CursorFile + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, CursorFile)
}

// Get returns the cursor for a key and whether it exists.
func (s CursorStore) Get(key string) (Cursor, bool) {
	c, ok := s.Cursors[key]
	return c, ok
}

// Set updates the cursor for a key and returns the updated store.
// Call SaveCursors(store) afterwards to persist.
func (s CursorStore) Set(key string, c Cursor) CursorStore {
	s.Cursors[key] = c
	return s
}

// Save is a convenience: set + persist in one call.
// Locking is handled inside SaveCursors.
//
// NOTE: This is safe for the current sequential batch mode (-workers=1).
// For concurrent writes (multiple connectors running in parallel), use
// SaveAll() instead — otherwise last-writer-wins will lose updates.
func (s CursorStore) Save(key string, c Cursor) error {
	updated := s.Set(key, c)
	return SaveCursors(updated)
}

// SaveAll performs a read-modify-write of multiple keys atomically under
// the same lock. Use this when a connector updates several cursor keys in
// one scrape run to prevent lost updates from interleaved writes.
func SaveAll(updates map[string]Cursor) error {
	cursorMu.Lock()
	defer cursorMu.Unlock()

	store := loadCursorsLocked()
	for k, v := range updates {
		store.Cursors[k] = v
	}
	return saveCursorsLocked(store)
}

// loadCursorsLocked and saveCursorsLocked are the lock-free primitives
// that the public Load/Save wrap. They MUST be called with cursorMu held.
func loadCursorsLocked() CursorStore {
	s := CursorStore{Cursors: make(map[string]Cursor)}
	data, err := os.ReadFile(CursorFile)
	if err != nil {
		return s
	}
	_ = json.Unmarshal(data, &s)
	if s.Cursors == nil {
		s.Cursors = make(map[string]Cursor)
	}
	return s
}

func saveCursorsLocked(s CursorStore) error {
	if err := os.MkdirAll(filepath.Dir(CursorFile), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	tmp := CursorFile + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, CursorFile)
}
