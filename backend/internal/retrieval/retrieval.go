package retrieval

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strings"

	"google.golang.org/genai"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

const embeddingDimensions = 1536 // matches existing mock fallback + any prior stored data

type Service struct {
	cfg    *config.Config
	repo   *repository.FirestoreRepo
	client *genai.Client
	sem    chan struct{}
}

func New(ctx context.Context, cfg *config.Config, repo *repository.FirestoreRepo) (*Service, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: cfg.Gemini.APIKey})
	if err != nil {
		return nil, fmt.Errorf("create genai client: %w", err)
	}
	return &Service{
		cfg:    cfg,
		repo:   repo,
		client: client,
		sem:    make(chan struct{}, 2),
	}, nil
}

// GenerateEmbedding calls Gemini API (gemini-embedding-2) to generate embedding vector for text.
// Returns the embedding vector, a boolean indicating if a mock fallback was used, and any error.
func (s *Service) GenerateEmbedding(ctx context.Context, text string) ([]float32, bool, error) {
	select {
	case <-ctx.Done():
		return nil, false, ctx.Err()
	case s.sem <- struct{}{}:
	}
	defer func() { <-s.sem }()

	dims := int32(embeddingDimensions)
	res, err := s.client.Models.EmbedContent(ctx, "gemini-embedding-2", genai.Text(text), &genai.EmbedContentConfig{
		OutputDimensionality: &dims,
	})
	if err != nil {
		var apiErr genai.APIError
		if errors.As(err, &apiErr) {
			// Trigger mock fallback only on auth/quota/billing/rate-limit errors.
			// Google API returns HTTP 400 (INVALID_ARGUMENT) with "API key not valid" message for bad keys.
			if apiErr.Code == 401 || apiErr.Code == 402 || apiErr.Code == 403 || apiErr.Code == 429 || apiErr.Code == 503 ||
				(apiErr.Code == 400 && (strings.Contains(apiErr.Message, "API key not valid") || strings.Contains(err.Error(), "API key not valid"))) {
				slog.Warn("gemini embedding API quota/auth error, falling back to mock vector", "status", apiErr.Code, "error", err)
				return s.getMockEmbedding(), true, nil
			}
		}
		// Treat other errors (like network/context timeouts, bad requests) as hard failures
		return nil, false, fmt.Errorf("gemini embed content: %w", err)
	}

	if len(res.Embeddings) == 0 || res.Embeddings[0] == nil {
		return nil, false, fmt.Errorf("empty embedding returned")
	}

	values := res.Embeddings[0].Values
	if len(values) != embeddingDimensions {
		return nil, false, fmt.Errorf("unexpected embedding dimension: got %d, want %d", len(values), embeddingDimensions)
	}

	return values, false, nil
}

func (s *Service) getMockEmbedding() []float32 {
	mockVec := make([]float32, embeddingDimensions)
	for i := range mockVec {
		mockVec[i] = float32(i) / float32(embeddingDimensions)
	}
	return mockVec
}

// CosineSimilarity computes similarity score between two vectors.
func CosineSimilarity(a, b []float32) float32 {
	if len(a) != len(b) || len(a) == 0 {
		return 0
	}
	var dotProduct, normA, normB float64
	for i := range a {
		dotProduct += float64(a[i]) * float64(b[i])
		normA += float64(a[i]) * float64(a[i])
		normB += float64(b[i]) * float64(b[i])
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return float32(dotProduct / (math.Sqrt(normA) * math.Sqrt(normB)))
}

type SearchResult struct {
	LawDocumentID string
	Score         float32
}

// Search retrieves all embeddings from Firestore, computes similarity scores against
// the query vector, and returns the top-N candidate matches.
func (s *Service) Search(ctx context.Context, queryVector []float32, topN int) ([]SearchResult, error) {
	allEmbs, err := s.repo.ListAllEmbeddings(ctx)
	if err != nil {
		return nil, fmt.Errorf("list all embeddings: %w", err)
	}

	var results []SearchResult
	for _, emb := range allEmbs {
		score := CosineSimilarity(queryVector, emb.Vector)
		results = append(results, SearchResult{
			LawDocumentID: emb.LawDocumentID,
			Score:         score,
		})
	}

	// Sort results descending by score
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Score > results[i].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	if len(results) > topN {
		results = results[:topN]
	}

	return results, nil
}
