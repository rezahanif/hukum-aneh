package retrieval

import (
	"container/heap"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/genai"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

const embeddingDimensions = 1536 // matches existing mock fallback + any prior stored data

type Service struct {
	cfg    *config.Config
	repo   repository.EmbeddingRepo // changed from *repository.FirestoreRepo in Phase 0.4
	qdrant *QdrantClient            // nil = brute-force fallback; non-nil = use Qdrant
	client *genai.Client
	sem    chan struct{}
	logger *slog.Logger
}

func New(ctx context.Context, cfg *config.Config, repo repository.EmbeddingRepo, qdrantClient *QdrantClient) (*Service, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: cfg.Gemini.APIKey})
	if err != nil {
		return nil, fmt.Errorf("create genai client: %w", err)
	}

	var logger *slog.Logger
	if qdrantClient != nil {
		logger = qdrantClient.logger
	} else {
		logger = slog.Default()
	}

	return &Service{
		cfg:    cfg,
		repo:   repo,
		qdrant: qdrantClient,
		client: client,
		sem:    make(chan struct{}, 2),
		logger: logger,
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

// scoreHeap implements a min-heap of SearchResult.
// The smallest (worst) score is at the top, so we can evict it
// when we find a better candidate, keeping only the top-N.
type scoreHeap []SearchResult

func (h scoreHeap) Len() int           { return len(h) }
func (h scoreHeap) Less(i, j int) bool { return h[i].Score < h[j].Score } // min-heap
func (h scoreHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *scoreHeap) Push(x interface{}) {
	*h = append(*h, x.(SearchResult))
}

func (h *scoreHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const defaultBatchSize = 500

// Search retrieves similar embeddings, using Qdrant when available
// and falling back to brute-force when Qdrant is nil or returns an error.
//
// CRITICAL: brute-force fallback only works during dual-write period (Phase 5-6),
// when Firestore still has the embeddings. After Phase 7.2 (Firestore removal),
// brute-force will return 0 results. Do not rely on fallback in production
// post-migration.
func (s *Service) Search(ctx context.Context, queryVector []float32, topN int) ([]SearchResult, error) {
	if topN <= 0 {
		topN = 5
	}

	// Qdrant path
	if s.qdrant != nil {
		results, err := s.qdrant.Search(ctx, queryVector, topN)
		if err != nil {
			s.logger.Warn("qdrant search failed, falling back to brute-force", "error", err)
			// fall through to brute-force
		} else {
			return results, nil
		}
	}

	// Brute-force fallback (legacy path)
	return s.bruteForceSearch(ctx, queryVector, topN)
}

// bruteForceSearch is the legacy in-memory cosine similarity search.
// Kept as fallback during dual-write migration period.
// TODO: remove after Phase 7.2 (Firestore removal) completes successfully.
func (s *Service) bruteForceSearch(ctx context.Context, queryVector []float32, topN int) ([]SearchResult, error) {
	h := &scoreHeap{}
	heap.Init(h)

	limit := defaultBatchSize
	var cursor *firestore.DocumentSnapshot

	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		batch, next, err := s.repo.ListEmbeddingsBatch(ctx, cursor, limit)
		if err != nil {
			return nil, fmt.Errorf("list embeddings batch failed: %w", err)
		}

		for _, emb := range batch {
			score := CosineSimilarity(queryVector, emb.Vector)
			candidate := SearchResult{
				LawDocumentID: emb.LawDocumentID,
				Score:         score,
			}

			if h.Len() < topN {
				heap.Push(h, candidate)
			} else if score > (*h)[0].Score {
				heap.Pop(h)
				heap.Push(h, candidate)
			}
		}

		if len(batch) < limit || next == nil {
			break
		}
		cursor = next
	}

	results := make([]SearchResult, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		results[i] = heap.Pop(h).(SearchResult)
	}

	return results, nil
}

// SearchByDocumentType pre-filters by document type, then computes
// similarity and returns top-N via min-heap. Uses brute-force because
// Qdrant doesn't know about doc_type unless we add it to payload — future enhancement.
func (s *Service) SearchByDocumentType(ctx context.Context, queryVector []float32, topN int, docType string) ([]SearchResult, error) {
	if topN <= 0 {
		topN = 5
	}

	embs, err := s.repo.ListEmbeddingsByDocType(ctx, docType)
	if err != nil {
		return nil, fmt.Errorf("list embeddings by doc_type failed: %w", err)
	}

	h := &scoreHeap{}
	heap.Init(h)

	for _, emb := range embs {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		score := CosineSimilarity(queryVector, emb.Vector)
		candidate := SearchResult{
			LawDocumentID: emb.LawDocumentID,
			Score:         score,
		}

		if h.Len() < topN {
			heap.Push(h, candidate)
		} else if score > (*h)[0].Score {
			heap.Pop(h)
			heap.Push(h, candidate)
		}
	}

	results := make([]SearchResult, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		results[i] = heap.Pop(h).(SearchResult)
	}

	return results, nil
}
