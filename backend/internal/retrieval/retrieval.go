package retrieval

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

// Service handles embedding generation and vector similarity search.
// Spec §5.4: embed text, search related laws.
type Service struct {
	cfg    *config.Config
	repo   *repository.FirestoreRepo
	client *http.Client
}

func New(cfg *config.Config, repo *repository.FirestoreRepo) *Service {
	return &Service{
		cfg:    cfg,
		repo:   repo,
		client: &http.Client{},
	}
}

type EmbeddingRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type EmbeddingResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

// GenerateEmbedding calls 9Router to generate embedding vector for text.
func (s *Service) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	reqBody := EmbeddingRequest{
		Input: text,
		Model: "text-embedding-3-small", // standard OpenAI embedding model
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/embeddings", s.cfg.Router9.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.cfg.Router9.APIKey))

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 status: %d, body: %s", resp.StatusCode, string(respBytes))
	}

	var embeddingResp EmbeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&embeddingResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(embeddingResp.Data) == 0 {
		return nil, fmt.Errorf("empty embedding data returned")
	}

	return embeddingResp.Data[0].Embedding, nil
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
