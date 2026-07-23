package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Config holds all runtime configuration for the pipeline.
type Config struct {
	Firebase struct {
		ProjectID      string `json:"project_id"`
		CredentialsPath string `json:"credentials_path"`
	} `json:"firebase"`

	Router9 struct {
		BaseURL string `json:"base_url"`
		APIKey  string `json:"api_key"`
	} `json:"router9"`

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
		DiscoveryInterval string `json:"discovery_interval"` // e.g. "1h", "30m"
		StuckJobThreshold string `json:"stuck_job_threshold"` // e.g. "6h"
	} `json:"scheduler"`

	Scraper struct {
		PythonPath string `json:"python_path"`
		ScriptPath string `json:"script_path"`
	} `json:"scraper"`

	SourcesPath string `json:"sources_path"`
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
	}

	cfg.Firebase.ProjectID = os.Getenv("FIREBASE_PROJECT_ID")
	cfg.Firebase.CredentialsPath = envOrDefault("FIREBASE_CREDENTIALS_PATH", "backend/configs/firebase-service-account.json")

	cfg.Router9.BaseURL = envOrDefault("ROUTER9_BASE_URL", "http://localhost:4000/v1")
	cfg.Router9.APIKey = os.Getenv("ROUTER9_API_KEY")

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
