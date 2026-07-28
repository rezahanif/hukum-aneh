// Package main implements cmd/collect_pdfs — a standalone utility that runs
// all registered connectors' CheckUpdates() + Download() pipeline and saves
// the fetched PDFs into per-connector subdirectories under an output dir.
//
// It does NOT touch Firestore / Postgres / Qdrant — it is purely a PDF
// harvester for offline archival / migration / Drive upload.
//
// Usage:
//   collect_pdfs -out /home/z/my-project/download/pdfs -per-connector 5
//
// Flags:
//   -out             output root directory (default: ./pdfs)
//   -per-connector   max PDFs to keep per connector (0 = unlimited)
//   -only            comma-separated connector names; if set, only those run
//   -skip             comma-separated connector names to skip
//   -timeout         per-Download timeout (default: 120s)
//   -verbose         debug logging
package main

import (
        "context"
        "encoding/json"
        "flag"
        "fmt"
        "io"
        "log/slog"
        "os"
        "os/exec"
        "os/signal"
        "path/filepath"
        "strings"
        "sync"
        "syscall"
        "time"

        "github.com/rezahanif/hukum-aneh/backend/internal/connectors"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/bkn"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/bpk"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/dpr"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/jdihn"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kemendag"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kemenkeu"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kemnaker"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/komdigi"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kpu"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/lkpp"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/ma"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/mkri"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/peraturan"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/setneg"
        "github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
)

// manifestEntry records one downloaded PDF.
type manifestEntry struct {
        Connector    string    `json:"connector"`
        LawNumber    string    `json:"law_number"`
        Title        string    `json:"title"`
        SourceURL    string    `json:"source_url"`
        DocumentType string    `json:"document_type"`
        Level        string    `json:"level"`
        Filename     string    `json:"filename"`
        Bytes        int64     `json:"bytes"`
        MimeType     string    `json:"mime_type"`
        SavedAt      time.Time `json:"saved_at"`
        Error        string    `json:"error,omitempty"`
        Skipped      bool      `json:"skipped,omitempty"`
}

func main() {
        var (
                outDir        string
                perConnector  int
                onlyNames     string
                skipNames     string
                perDownloadTO time.Duration
                verbose       bool
                dryRun        bool
        )
        flag.StringVar(&outDir, "out", "./pdfs", "output root directory")
        flag.IntVar(&perConnector, "per-connector", 0, "max PDFs per connector (0 = unlimited)")
        flag.StringVar(&onlyNames, "only", "", "comma-separated connector names to run (empty = all)")
        flag.StringVar(&skipNames, "skip", "", "comma-separated connector names to skip")
        flag.DurationVar(&perDownloadTO, "timeout", 120*time.Second, "per-Download timeout")
        flag.BoolVar(&verbose, "verbose", false, "debug logging")
        flag.BoolVar(&dryRun, "dry-run", false, "list discovered laws but don't download")
        flag.Parse()

        level := slog.LevelInfo
        if verbose {
                level = slog.LevelDebug
        }
        logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
        slog.SetDefault(logger)

        // Build set filters
        onlySet := parseSet(onlyNames)
        skipSet := parseSet(skipNames)

        // Python scraper bridge — required by most connectors.
        scr := scraper.New("python3", "backend/python/scraper/scrape.py", logger)

        // Wire all 14 connectors (same as cmd/pipeline/main.go, minus DB).
        registry := connectors.NewRegistry()
        register := func(c connectors.Connector) {
                name := c.Name()
                if len(onlySet) > 0 && !onlySet[name] {
                        return
                }
                if skipSet[name] {
                        return
                }
                registry.Register(name, c)
        }

        register(peraturan.New(scr, logger))
        register(jdihn.New(scr, logger))
        register(bpk.New(scr, logger))
        register(mkri.New(scr, logger))
        register(setneg.New(scr, logger))
        register(kemenkeu.New(scr, logger))
        register(ma.New(logger))
        register(kemnaker.New(scr, logger))
        register(kemendag.New(scr, logger))
        register(komdigi.New(scr, logger))
        register(kpu.New(scr, logger))
        register(bkn.New(scr, logger))
        register(lkpp.New(scr, logger))
        register(dpr.New(scr, logger))

        if len(registry.All()) == 0 {
                logger.Error("no connectors matched filter")
                os.Exit(1)
        }

        logger.Info("collect_pdfs starting",
                "out_dir", outDir,
                "per_connector_cap", perConnector,
                "connectors", len(registry.All()),
                "per_download_timeout", perDownloadTO,
                "dry_run", dryRun,
        )

        if err := os.MkdirAll(outDir, 0o755); err != nil {
                logger.Error("create out dir failed", "error", err)
                os.Exit(1)
        }

        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()

        sigCh := make(chan os.Signal, 1)
        signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
        go func() {
                <-sigCh
                logger.Info("interrupt received, cancelling")
                cancel()
        }()

        // Per-connector output dir + global manifest.
        var (
                manifestMu sync.Mutex
                manifest   []manifestEntry
        )

        // Load the Drive upload manifest (if present) so we can skip
        // re-downloading PDFs that have already been uploaded to Drive and
        // deleted locally. The Drive manifest is written by upload_to_drive.py.
        uploadedOnDrive := loadDriveManifest(outDir)
        if len(uploadedOnDrive) > 0 {
                logger.Info("found Drive manifest, will skip already-uploaded files",
                        "count", len(uploadedOnDrive))
        }

        for _, name := range sortedNames(registry.Names()) {
                c, _ := registry.Get(name)
                connDir := filepath.Join(outDir, sanitizeDirName(name))
                if err := os.MkdirAll(connDir, 0o755); err != nil {
                        logger.Error("mkdir connector dir failed", "connector", name, "error", err)
                        continue
                }

                logger.Info("connector start", "name", name, "dir", connDir)

                // 1) CheckUpdates — listing of laws not yet in DB. Since this tool has
                //    no DB, this returns whatever the connector's listing scrapes.
                docs, err := c.CheckUpdates(ctx)
                if err != nil {
                        if ctx.Err() != nil {
                                logger.Warn("cancelled during CheckUpdates", "connector", name)
                                break
                        }
                        logger.Error("CheckUpdates failed", "connector", name, "error", err)
                        appendManifest(&manifestMu, &manifest, manifestEntry{
                                Connector: name,
                                Error:     fmt.Sprintf("CheckUpdates: %v", err),
                                SavedAt:   time.Now(),
                        })
                        continue
                }
                logger.Info("CheckUpdates complete", "connector", name, "discovered", len(docs))

                if dryRun {
                        for _, d := range docs {
                                logger.Info("dry-run discovered", "connector", name, "law", d.LawNumber, "title", d.Title)
                        }
                        continue
                }

                // 2) Download each PDF, up to cap (cap counts only actual downloads, not skips).
                kept := 0
                skipped := 0
                for _, d := range docs {
                        if ctx.Err() != nil {
                                break
                        }
                        if perConnector > 0 && kept >= perConnector {
                                logger.Info("per-connector cap reached", "connector", name, "cap", perConnector)
                                break
                        }

                        entry := manifestEntry{
                                Connector:    name,
                                LawNumber:    d.LawNumber,
                                Title:        d.Title,
                                SourceURL:    d.SourceURL,
                                DocumentType: d.DocumentType,
                                Level:        d.Level,
                                SavedAt:      time.Now(),
                        }

                        // Predict filename (before downloading) so we can skip
                        // the HTTP request entirely if the file already exists
                        // locally or has already been uploaded to Drive.
                        // Connectors typically set Filename = LawNumber + ".pdf"
                        // (no sanitization), so we mirror that here. If a connector
                        // uses a different naming, the file will be re-downloaded
                        // (rare, acceptable).
                        predictedFilename := d.LawNumber + ".pdf"
                        if !strings.HasSuffix(strings.ToLower(predictedFilename), ".pdf") {
                                predictedFilename += ".pdf"
                        }
                        predictedPath := filepath.Join(connDir, predictedFilename)

                        // Skip if file already exists locally (resumable — don't re-download
                        // what we already have on disk).
                        if fi, err := os.Stat(predictedPath); err == nil && fi.Size() > 0 && isPDF(predictedPath) {
                                logger.Info("already exists locally, skipping download",
                                        "connector", name, "law", d.LawNumber, "path", predictedPath, "bytes", fi.Size())
                                entry.Filename = predictedFilename
                                entry.Bytes = fi.Size()
                                entry.MimeType = "application/pdf"
                                entry.Skipped = true
                                appendManifest(&manifestMu, &manifest, entry)
                                skipped++
                                continue
                        }

                        // Skip if already uploaded to Drive (per Drive manifest).
                        // This lets us resume after `upload_to_drive.py --delete-after-upload`
                        // removes the local copy.
                        if uploadedOnDrive[predictedFilename] {
                                logger.Info("already on Drive (per manifest), skipping download",
                                        "connector", name, "law", d.LawNumber, "filename", predictedFilename)
                                entry.Filename = predictedFilename
                                entry.MimeType = "application/pdf"
                                entry.Skipped = true
                                appendManifest(&manifestMu, &manifest, entry)
                                skipped++
                                continue
                        }

                        dlCtx, dlCancel := context.WithTimeout(ctx, perDownloadTO)
                        raw, err := c.Download(dlCtx, d)
                        if err != nil {
                                dlCancel()
                                logger.Warn("download failed", "connector", name, "law", d.LawNumber, "error", err)
                                entry.Error = fmt.Sprintf("Download: %v", err)
                                appendManifest(&manifestMu, &manifest, entry)
                                continue
                        }

                        // Filename: prefer connector-supplied, else use predicted.
                        filename := raw.Filename
                        if filename == "" {
                                filename = predictedFilename
                        }
                        // Ensure .pdf extension
                        if !strings.HasSuffix(strings.ToLower(filename), ".pdf") {
                                filename += ".pdf"
                        }
                        // Make unique — prefix law number if not already there
                        if d.LawNumber != "" && !strings.Contains(filename, d.LawNumber) {
                                filename = sanitizeFilename(d.LawNumber) + "_" + filename
                        }
                        outPath := filepath.Join(connDir, filename)

                        f, err := os.Create(outPath)
                        if err != nil {
                                logger.Error("create file failed", "connector", name, "path", outPath, "error", err)
                                raw.Content.Close()
                                dlCancel()
                                entry.Error = fmt.Sprintf("Create: %v", err)
                                appendManifest(&manifestMu, &manifest, entry)
                                continue
                        }

                        n, err := io.Copy(f, raw.Content)
                        f.Close()
                        raw.Content.Close()
                        dlCancel()
                        if err != nil {
                                logger.Warn("write failed", "connector", name, "path", outPath, "error", err)
                                entry.Error = fmt.Sprintf("Copy: %v", err)
                                appendManifest(&manifestMu, &manifest, entry)
                                _ = os.Remove(outPath)
                                continue
                        }

                        // Verify it's actually a PDF (some sources return HTML error pages).
                        if !isPDF(outPath) {
                                logger.Warn("downloaded file is not a PDF, removing",
                                        "connector", name, "path", outPath, "law", d.LawNumber)
                                _ = os.Remove(outPath)
                                entry.Error = "non-PDF response (likely HTML error/redirect page)"
                                appendManifest(&manifestMu, &manifest, entry)
                                continue
                        }

                        entry.Filename = filename
                        entry.Bytes = n
                        entry.MimeType = raw.MimeType
                        appendManifest(&manifestMu, &manifest, entry)
                        kept++
                        logger.Info("downloaded",
                                "connector", name,
                                "law", d.LawNumber,
                                "bytes", n,
                                "path", outPath,
                        )

                        // Polite delay between downloads (Komdigi adds its own
                        // 1.2s pre-delay in its Download() method for rate-limit safety).
                        select {
                        case <-ctx.Done():
                        case <-time.After(300 * time.Millisecond):
                        }
                }
                logger.Info("connector done", "name", name, "kept", kept, "skipped", skipped)
        }

        // Write manifest JSON
        manifestPath := filepath.Join(outDir, "_manifest.json")
        manifestMu.Lock()
        data, _ := json.MarshalIndent(manifest, "", "  ")
        manifestMu.Unlock()
        if err := os.WriteFile(manifestPath, data, 0o644); err != nil {
                logger.Error("write manifest failed", "error", err)
        } else {
                logger.Info("manifest written", "path", manifestPath, "entries", len(manifest))
        }

        // Per-connector summary
        totals := map[string]int{}
        for _, e := range manifest {
                if e.Error == "" {
                        totals[e.Connector]++
                }
        }
        logger.Info("collect_pdfs complete", "totals", totals)
}

func parseSet(csv string) map[string]bool {
        out := map[string]bool{}
        csv = strings.TrimSpace(csv)
        if csv == "" {
                return out
        }
        for _, s := range strings.Split(csv, ",") {
                s = strings.TrimSpace(s)
                if s != "" {
                        out[s] = true
                }
        }
        return out
}

func sortedNames(names []string) []string {
        out := make([]string, len(names))
        copy(out, names)
        for i := 1; i < len(out); i++ {
                for j := i; j > 0 && out[j-1] > out[j]; j-- {
                        out[j-1], out[j] = out[j], out[j-1]
                }
        }
        return out
}

// sanitizeDirName makes a connector name safe for use as a directory name.
func sanitizeDirName(name string) string {
        r := strings.NewReplacer(
                "/", "_",
                "\\", "_",
                ":", "_",
                " ", "_",
        )
        return r.Replace(name)
}

// sanitizeFilename strips characters that are problematic on the filesystem.
func sanitizeFilename(s string) string {
        r := strings.NewReplacer(
                "/", "-",
                "\\", "-",
                ":", "-",
                " ", "_",
                ".", "_",
                ",", "",
                "(", "",
                ")", "",
        )
        return r.Replace(s)
}

// isPDF reads the first 5 bytes and checks for the %PDF- magic.
func isPDF(path string) bool {
        f, err := os.Open(path)
        if err != nil {
                return false
        }
        defer f.Close()
        buf := make([]byte, 5)
        n, _ := f.Read(buf)
        if n < 5 {
                return false
        }
        return string(buf) == "%PDF-"
}

func appendManifest(mu *sync.Mutex, m *[]manifestEntry, e manifestEntry) {
        mu.Lock()
        *m = append(*m, e)
        mu.Unlock()
}

// loadDriveManifest reads <outDir>/_drive_manifest.json (written by
// upload_to_drive.py) and returns a set of filenames that have been
// successfully uploaded to Drive. Used to skip re-downloading PDFs whose
// local copies have been deleted post-upload.
func loadDriveManifest(outDir string) map[string]bool {
        out := map[string]bool{}
        path := filepath.Join(outDir, "_drive_manifest.json")
        data, err := os.ReadFile(path)
        if err != nil {
                return out
        }
        // The manifest is a JSON array of objects with at least "filename" and
        // "status" fields. Status "ok" or "skipped" means the file is on Drive.
        var entries []struct {
                Filename string `json:"filename"`
                Status   string `json:"status"`
        }
        if err := json.Unmarshal(data, &entries); err != nil {
                return out
        }
        for _, e := range entries {
                if e.Filename == "" {
                        continue
                }
                if e.Status == "ok" || e.Status == "skipped" {
                        out[e.Filename] = true
                }
        }
        return out
}

// init ensures python3 is on PATH (some sandboxes don't include it by default).
func init() {
        if _, err := exec.LookPath("python3"); err != nil {
                if _, err2 := os.Stat("/home/z/.venv/bin/python3"); err2 == nil {
                        os.Setenv("PATH", "/home/z/.venv/bin:"+os.Getenv("PATH"))
                }
        }
}
