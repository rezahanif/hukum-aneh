package retrieval

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
)

func TestCosineSimilarity(t *testing.T) {
	cases := []struct {
		name string
		a    []float32
		b    []float32
		want float32
	}{
		{
			name: "identical vectors",
			a:    []float32{1.0, 0.0, 0.0},
			b:    []float32{1.0, 0.0, 0.0},
			want: 1.0,
		},
		{
			name: "orthogonal vectors",
			a:    []float32{1.0, 0.0, 0.0},
			b:    []float32{0.0, 1.0, 0.0},
			want: 0.0,
		},
		{
			name: "opposite vectors",
			a:    []float32{1.0, 0.0, 0.0},
			b:    []float32{-1.0, 0.0, 0.0},
			want: -1.0,
		},
		{
			name: "empty inputs",
			a:    []float32{},
			b:    []float32{},
			want: 0.0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := CosineSimilarity(c.a, c.b)
			diff := got - c.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 1e-5 {
				t.Errorf("CosineSimilarity() = %f, want %f", got, c.want)
			}
		})
	}
}

func TestGenerateEmbedding_FallbackAndIntegration(t *testing.T) {
	ctx := context.Background()
	cfg := &config.Config{}

	// First test the Mock fallback flow with an invalid API Key (always runs, fast, deterministic)
	cfg.Gemini.APIKey = "invalid-api-key-to-trigger-401"
	service, err := New(ctx, cfg, nil, nil)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	vec, isMock, err := service.GenerateEmbedding(ctx, "test query")
	if err != nil {
		t.Fatalf("expected mock fallback, got error: %v", err)
	}
	if !isMock {
		t.Error("expected isMock to be true on invalid API key")
	}
	if len(vec) != embeddingDimensions {
		t.Errorf("expected mock vector of length %d, got %d", embeddingDimensions, len(vec))
	}

	// Live network call: gate with testing.Short()
	if testing.Short() {
		t.Skip("skipping live Gemini API call in short mode")
		return
	}

	// Try to get GEMINI_API_KEY from environment or by searching parent directories for .env
	realKey := os.Getenv("GEMINI_API_KEY")
	if realKey == "" {
		if envVars, err := loadDotEnvFromParents(); err == nil {
			realKey = envVars["GEMINI_API_KEY"]
		}
	}

	if realKey == "" {
		t.Skip("skipping real API integration test because GEMINI_API_KEY is not configured")
		return
	}

	t.Log("Running live Gemini API integration test...")
	cfg.Gemini.APIKey = realKey
	liveService, err := New(ctx, cfg, nil, nil)
	if err != nil {
		t.Fatalf("failed to create live service: %v", err)
	}

	liveVec, liveIsMock, err := liveService.GenerateEmbedding(ctx, "hello world")
	if err != nil {
		t.Fatalf("live embedding call failed: %v", err)
	}
	if liveIsMock {
		t.Error("expected liveIsMock to be false, but fell back to mock vector")
	}
	if len(liveVec) != embeddingDimensions {
		t.Errorf("expected live vector of length %d, got %d", embeddingDimensions, len(liveVec))
	}
	t.Logf("Successfully fetched real embedding with length: %d", len(liveVec))
}

func loadDotEnvFromParents() (map[string]string, error) {
	env := make(map[string]string)
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if data, err := os.ReadFile(envPath); err == nil {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					env[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
				}
			}
			return env, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return nil, os.ErrNotExist
}
