# Phase 7.2 — Remove Firestore Dependencies

**Status:** DEFERRED — execute only after production migration is validated.
**Owner:** Stream A (after Stream B completes B-4.1 + B-4.2)

## When to execute

Phase 7.2 is the **final cutover step**. Execute ONLY when ALL of these are true:

1. Stream B has completed B-2.1 (Qdrant client wrapper)
2. Stream B has completed B-4.1 (retrieval.Search uses Qdrant)
3. Stream B has completed B-4.2 (Engine upserts vectors to Qdrant)
4. Phase 6.1 migrate_to_pg has been run against production Firestore
5. Vectors have been pushed to Qdrant (Stream B B-4.2 task)
6. `STORAGE_MODE=postgres` has been running in production for ≥1 week with no issues
7. Dual-write mode has been disabled (no new Firestore writes for ≥1 week)
8. Firestore backup has been verified restorable

If any of the above are not yet true, **DO NOT execute Phase 7.2**. The Firestore
code must remain compiled in until you are 100% committed to PG-only operation.

## What "remove Firestore deps" means

Three things, in order of severity:

1. **Runtime**: don't initialize Firestore client. Already done — the factory
   in `factory.go` only creates a Firestore client when `STORAGE_MODE=firestore`
   or `STORAGE_MODE=dual_write`. With `STORAGE_MODE=postgres`, no Firestore
   calls are made at all.

2. **Compile time**: don't compile `firestore.go` (and the Firestore branches
   of `factory.go`) into the binary. Requires build tags. Optional but produces
   a smaller binary and removes the `cloud.google.com/go/firestore` import.

3. **Source code**: delete `firestore.go`, the Firestore branches of
   `factory.go`, and the `*firestore.DocumentSnapshot` cursor type from
   `interfaces.go`. Destructive — only do this after steps 1 and 2 have been
   running in production for a while.

## Step-by-step procedure

### Step 1: Refactor `EmbeddingRepo.ListEmbeddingsBatch` cursor type

The only Firestore-specific type in the repository interfaces is
`*firestore.DocumentSnapshot` in `EmbeddingRepo.ListEmbeddingsBatch`. To remove
Firestore entirely, replace it with an opaque cursor type.

**File:** `backend/internal/repository/interfaces.go`

Replace:
```go
ListEmbeddingsBatch(
    ctx context.Context,
    cursor *firestore.DocumentSnapshot,
    limit int,
) ([]models.EmbeddingEntry, *firestore.DocumentSnapshot, error)
```

With:
```go
// EmbeddingCursor is an opaque cursor for batched embedding reads.
// Concrete type depends on backend: *firestore.DocumentSnapshot for Firestore,
// nil (unsupported) for Postgres (use Qdrant for vector search instead).
type EmbeddingCursor interface{}

ListEmbeddingsBatch(
    ctx context.Context,
    cursor EmbeddingCursor,
    limit int,
) ([]models.EmbeddingEntry, EmbeddingCursor, error)
```

Update `FirestoreRepo.ListEmbeddingsBatch` signature to match (the body stays
the same — `*firestore.DocumentSnapshot` satisfies `interface{}`).

Update `PostgresRepo.ListEmbeddingsBatch` signature to match (still returns
an error — `ListEmbeddingsBatch not supported in PG mode`).

Update `DualWriteRepo.ListEmbeddingsBatch` signature to match (delegates to
primary).

Update `retrieval.go` `Search()` to use the new `EmbeddingCursor` type
instead of `*firestore.DocumentSnapshot`.

Remove the `cloud.google.com/go/firestore` import from `interfaces.go`.

### Step 2: Add build tag for postgres-only builds

**New file:** `backend/internal/repository/factory_postgres_only.go`

```go
//go:build postgres_only

package repository

import (
    "context"
    "errors"

    "github.com/rezahanif/hukum-aneh/backend/internal/config"
)

// In postgres_only builds, Firestore code is excluded entirely.
// NewRepoSet only accepts STORAGE_MODE=postgres.
func NewRepoSet(ctx context.Context, cfg *config.Config) (*RepoSet, error) {
    if cfg.StorageMode != config.StorageModePostgres {
        return nil, errors.New("postgres_only build requires STORAGE_MODE=postgres")
    }
    return newPostgresRepoSet(ctx, cfg)
}
```

**Edit:** `backend/internal/repository/factory.go`

Add `//go:build !postgres_only` at the top of the file. This excludes the
Firestore + DualWrite branches from postgres_only builds.

**Edit:** `backend/internal/repository/firestore.go`

Add `//go:build !postgres_only` at the top of the file. This excludes the
FirestoreRepo entirely from postgres_only builds.

### Step 3: Verify both builds work

```bash
# Default build (all 3 modes supported)
go build ./...

# Postgres-only build (smaller binary, no Firestore deps)
go build -tags postgres_only ./...
```

Both should produce a working binary. The postgres_only binary should be
~5-10 MB smaller (no `cloud.google.com/go/firestore` and its transitive deps).

### Step 4: Update Dockerfile / CI to use postgres_only tag

Once the postgres_only build is verified, update the production Dockerfile
to use `-tags postgres_only`. This is the actual cutover point.

**File:** `Dockerfile` (owned by Stream B B-7.1)

```dockerfile
# Before: builds with all 3 modes supported
RUN go build -o /app/pipeline ./backend/cmd/pipeline

# After: postgres-only build
RUN go build -tags postgres_only -o /app/pipeline ./backend/cmd/pipeline
```

### Step 5 (optional, after 30+ days stable in production): Delete Firestore source

Only after postgres_only builds have been running in production for 30+ days
with zero rollback events:

1. Delete `backend/internal/repository/firestore.go`
2. Remove the Firestore + DualWrite branches from `backend/internal/repository/factory.go`
3. Remove the `//go:build !postgres_only` tags (no longer needed)
4. Remove `cfg.Firebase` from `backend/internal/config/config.go`
5. Remove `FIREBASE_*` env vars from `.env.example`
6. Remove `cloud.google.com/go/firestore` from `go.mod` via `go mod tidy`
7. Run `go build ./...` — should still pass

This is the truly destructive step. Do NOT do this until you are 100%
committed to PG-only operation and have a verified backup of Firestore data.

## Rollback

If Phase 7.2 needs to be rolled back:

- **Steps 1-4**: Revert the git commit(s). Binary rebuild picks up Firestore code again.
- **Step 5**: Restore `firestore.go` from git history. Re-add the `cloud.google.com/go/firestore` dep via `go get`.

The dual-write mode (STORAGE_MODE=dual_write) is the safest rollback target —
it re-enables Firestore as source of truth while keeping PG as secondary.
