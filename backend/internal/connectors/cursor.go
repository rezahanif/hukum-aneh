package connectors

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// CursorFile is the default location for all connector cursors.
// Connectors share this file via namespaced keys.
const CursorFile = "backend/configs/scrape_cursor.json"

// Cursor records the last known position for a given source key.
// Use SourceType (e.g. "Undang-Undang (UU)") as the key.
type Cursor struct {
	LastKnownID      string    `json:"last_known_id"`       // last known law number or document ID
	LastKnownTitle   string    `json:"last_known_title,omitempty"`
	Timestamp        time.Time `json:"timestamp"`
	Extra            map[string]string `json:"extra,omitempty"` // connector-specific extras
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

// SaveCursors writes the cursor map atomically. Creates parent dir if missing.
func SaveCursors(s CursorStore) error {
	if err := os.MkdirAll(filepath.Dir(CursorFile), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(CursorFile, data, 0644)
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
func (s CursorStore) Save(key string, c Cursor) error {
	updated := s.Set(key, c)
	return SaveCursors(updated)
}
