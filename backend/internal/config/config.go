package config

import (
        "encoding/json"
        "fmt"
        "os"
        "strconv"
        "strings"
)

// StorageMode controls which storage backend the pipeline uses.
//   - "firestore"  : Firestore only (current production behavior)
//   - "postgres"   : PostgreSQL only (target end state after migration)
//   - "dual_write" : Write to both Firestore and Postgres; reads prefer Firestore
//     (migration cutover mode — Phase 5 dual-write pattern)
//
// The repository factory (Phase 5.1) reads this to decide which concrete
// RepoSet to construct. Empty defaults to "firestore" for backward compat.
type StorageMode string

const (
        StorageModeFirestore string = "firestore"
        StorageModePostgres  string = "postgres"
        StorageModeDualWrite string = "dual_write"
)

// Config holds all runtime configuration for the pipeline.
type Config struct {
        // StorageMode selects the storage backend. See StorageMode docs above.
        StorageMode string `json:"storage_mode"`

        Firebase struct {
                ProjectID       string `json:"project_id"`
                CredentialsPath string `json:"credentials_path"`
        } `json:"firebase"`

        // Postgres holds PostgreSQL connection settings. Required when
        // StorageMode is "postgres" or "dual_write". Ignored otherwise.
        Postgres struct {
                Host     string `json:"host"`
                Port     int    `json:"port"`
                Database string `json:"database"`
                Username string `json:"username"`
                Password string `json:"password"`
                SSLMode  string `json:"ssl_mode"` // disable, require, verify-ca, verify-full
                // MaxConns caps the pgxpool size. Default 10 if unset.
                MaxConns int `json:"max_conns"`
                // MigrationsPath is the filesystem path to .sql files used by
                // golang-migrate. Defaults to "backend/migrations".
                MigrationsPath string `json:"migrations_path"`
        } `json:"postgres"`

        // Qdrant holds Qdrant vector store settings. Required when
        // StorageMode is "postgres" (vector data lives in Qdrant, not PG).
        // Ignored when StorageMode is "firestore".
        Qdrant struct {
                Host string `json:"host"` // gRPC host (default localhost)
                Port int    `json:"port"` // gRPC port (default 6334)
                // Collection is the Qdrant collection name. Default "hukum_aneh_laws".
                Collection string `json:"collection"`
                // VectorSize must match the embedding model output dim.
                // gemini-embedding-2 outputs 1536; text-embedding-004 outputs 768.
                // Defaults to 1536 (current model: gemini-embedding-2).
                VectorSize int `json:"vector_size"`
                // APIKey optional; set if Qdrant is behind an auth gateway.
                APIKey string `json:"api_key"`
        } `json:"qdrant"`

        Router9 struct {
                BaseURL string `json:"base_url"`
                APIKey  string `json:"api_key"`
                Model   string `json:"model"`
        } `json:"router9"`

        Gemini struct {
                APIKey string `json:"api_key"`
        } `json:"gemini"`

        Telegram struct {
                BotToken string `json:"bot_token"`
                ChatID   string `json:"chat_id"`
        } `json:"telegram"`

        Instagram struct {
                AccessToken string `json:"access_token"`
                AccountID   string `json:"account_id"`
        } `json:"instagram"`

        ImageGen struct {
                Provider string `json:"provider"` // openai, flux, sdxl, ideogram
                APIKey   string `json:"api_key"`
                BaseURL  string `json:"base_url"`
        } `json:"image_gen"`

        Scheduler struct {
                DiscoveryInterval string `json:"discovery_interval"`  // e.g. "1h", "30m"
                StuckJobThreshold string `json:"stuck_job_threshold"` // e.g. "6h"
        } `json:"scheduler"`

        Scraper struct {
                PythonPath string `json:"python_path"`
                ScriptPath string `json:"script_path"`
        } `json:"scraper"`

        WorkerPoolSize int `json:"worker_pool_size"`

        SourcesPath string `json:"sources_path"`

        Google struct {
                CredentialsPath string `json:"credentials_path"`
                TokenPath       string `json:"token_path"`
                FolderID        string `json:"folder_id"`
        } `json:"google"`
}

// Load reads config from environment variables.
// Secrets never hardcoded — all from env.
func Load() (*Config, error) {
        // Load .env if present
        if data, err := os.ReadFile(".env"); err == nil {
                lines := strings.Split(string(data), "\n")
                for _, line := range lines {
                        line = strings.TrimSpace(line)
                        if line == "" || strings.HasPrefix(line, "#") {
                                continue
                        }
                        parts := strings.SplitN(line, "=", 2)
                        if len(parts) == 2 {
                                key := strings.TrimSpace(parts[0])
                                val := strings.TrimSpace(parts[1])
                                os.Setenv(key, val)
                        }
                }
        }

        cfg := &Config{
                SourcesPath: envOrDefault("SOURCES_PATH", "backend/configs/sources.json"),
                StorageMode: envOrDefault("STORAGE_MODE", StorageModeFirestore),
        }

        cfg.Firebase.ProjectID = os.Getenv("FIREBASE_PROJECT_ID")
        cfg.Firebase.CredentialsPath = envOrDefault("FIREBASE_CREDENTIALS_PATH", "backend/configs/firebase-service-account.json")

        // PostgreSQL config (Phase 1.3)
        cfg.Postgres.Host = envOrDefault("POSTGRES_HOST", "localhost")
        cfg.Postgres.Port = envIntOrDefault("POSTGRES_PORT", 5432)
        cfg.Postgres.Database = envOrDefault("POSTGRES_DB", "hukum_aneh")
        cfg.Postgres.Username = envOrDefault("POSTGRES_USER", "hukum")
        cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
        cfg.Postgres.SSLMode = envOrDefault("POSTGRES_SSL_MODE", "disable")
        cfg.Postgres.MaxConns = envIntOrDefault("POSTGRES_MAX_CONNS", 10)
        cfg.Postgres.MigrationsPath = envOrDefault("POSTGRES_MIGRATIONS_PATH", "backend/migrations")

        // Qdrant config (Phase 1.3) — used by Stream B B-2.1 + Phase 4 of Stream A
        cfg.Qdrant.Host = envOrDefault("QDRANT_HOST", "localhost")
        cfg.Qdrant.Port = envIntOrDefault("QDRANT_PORT", 6334)
        cfg.Qdrant.Collection = envOrDefault("QDRANT_COLLECTION", "hukum_aneh_laws")
        // VectorSize must match the embedding model output dimension.
        // gemini-embedding-2 outputs 1536 by default; text-embedding-004 outputs 768.
        // Set via QDRANT_VECTOR_SIZE env or defaults to 1536 (gemini-embedding-2).
        cfg.Qdrant.VectorSize = envIntOrDefault("QDRANT_VECTOR_SIZE", 1536)
        cfg.Qdrant.APIKey = os.Getenv("QDRANT_API_KEY")

        cfg.Router9.BaseURL = envOrDefault("ROUTER9_BASE_URL", "http://localhost:4000/v1")
        cfg.Router9.APIKey = os.Getenv("ROUTER9_API_KEY")
        cfg.Router9.Model = envOrDefault("ROUTER9_MODEL", "gpt-4o")

        cfg.Gemini.APIKey = os.Getenv("GEMINI_API_KEY")

        cfg.Telegram.BotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
        cfg.Telegram.ChatID = os.Getenv("TELEGRAM_CHAT_ID")

        cfg.Instagram.AccessToken = os.Getenv("IG_ACCESS_TOKEN")
        cfg.Instagram.AccountID = os.Getenv("IG_ACCOUNT_ID")

        cfg.ImageGen.Provider = envOrDefault("IMAGE_GEN_PROVIDER", "openai")
        cfg.ImageGen.APIKey = os.Getenv("IMAGE_GEN_API_KEY")
        cfg.ImageGen.BaseURL = envOrDefault("IMAGE_GEN_BASE_URL", "https://api.openai.com/v1")

        cfg.Scheduler.DiscoveryInterval = envOrDefault("DISCOVERY_INTERVAL", "1h")
        cfg.Scheduler.StuckJobThreshold = envOrDefault("STUCK_JOB_THRESHOLD", "6h")

        cfg.Scraper.PythonPath = envOrDefault("PYTHON_PATH", "python3")
        cfg.Scraper.ScriptPath = envOrDefault("SCRAPER_SCRIPT_PATH", "backend/python/scraper/scrape.py")

        cfg.WorkerPoolSize = 3
        if wp := os.Getenv("WORKER_POOL_SIZE"); wp != "" {
                if val, err := strconv.Atoi(wp); err == nil && val > 0 {
                        cfg.WorkerPoolSize = val
                }
        }

        cfg.Google.CredentialsPath = envOrDefault("GOOGLE_CREDENTIALS_PATH", "/project/google-credentials.json")
        cfg.Google.TokenPath = envOrDefault("GOOGLE_TOKEN_PATH", "/project/token.json")
        cfg.Google.FolderID = os.Getenv("GOOGLE_DRIVE_FOLDER_ID")

        return cfg, nil
}

// LoadSources reads the sources.json file.
func LoadSources(path string) ([]SourceConfig, error) {
        data, err := os.ReadFile(path)
        if err != nil {
                return nil, fmt.Errorf("read sources: %w", err)
        }
        var sources []SourceConfig
        if err := json.Unmarshal(data, &sources); err != nil {
                return nil, fmt.Errorf("parse sources: %w", err)
        }
        return sources, nil
}

type SourceConfig struct {
        Name           string `json:"name"`
        Level          string `json:"level"`
        DocumentType   string `json:"document_type"`
        OfficialSource string `json:"official_source"`
        OfficialURL    string `json:"official_url"`
        Notes          string `json:"notes"`
}

func envOrDefault(key, defaultVal string) string {
        if v := os.Getenv(key); v != "" {
                return v
        }
        return defaultVal
}

// envIntOrDefault parses an env var as int; falls back to defaultVal on missing/invalid.
func envIntOrDefault(key string, defaultVal int) int {
        if v := os.Getenv(key); v != "" {
                if val, err := strconv.Atoi(v); err == nil {
                        return val
                }
        }
        return defaultVal
}

// PostgresDSN builds a libpq-style connection string from Postgres config.
// Uses URL format (postgres://user:pass@host:port/db?sslmode=...) because
// pgx's keyword=value parser mishandles empty password values (treats the
// next key=value as the password value).
//
// Used by pgxpool.ParseConfig and golang-migrate.
func (c *Config) PostgresDSN() string {
        return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
                c.Postgres.Username,
                c.Postgres.Password,
                c.Postgres.Host,
                c.Postgres.Port,
                c.Postgres.Database,
                c.Postgres.SSLMode,
        )
}

// IsPostgres returns true when StorageMode is postgres OR dual_write.
// Convenience for code paths that need to know if PG is in play at all.
func (c *Config) IsPostgres() bool {
        return c.StorageMode == StorageModePostgres || c.StorageMode == StorageModeDualWrite
}
