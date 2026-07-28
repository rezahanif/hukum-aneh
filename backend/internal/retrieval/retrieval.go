package retrieval

import (
        "bytes"
        "container/heap"
        "context"
        "encoding/json"
        "errors"
        "fmt"
        "log/slog"
        "math"
        "net/http"
        "strings"

        "cloud.google.com/go/firestore"

        "github.com/rezahanif/hukum-aneh/backend/internal/config"
        "github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

// embeddingDimensions is the fallback vector size when no Qdrant config is available
// (e.g. STORAGE_MODE=firestore where cfg.Qdrant.VectorSize may be unset).
// The actual vector size is determined by s.vectorSize (config-driven) at runtime.
const defaultEmbeddingDimensions = 1536 // matches gemini-embedding-2 output

type Service struct {
        cfg        *config.Config
        repo       repository.EmbeddingRepo // changed from *repository.FirestoreRepo in Phase 0.4
        qdrant     *QdrantClient            // nil = brute-force fallback; non-nil = use Qdrant
        sem        chan struct{}
        logger     *slog.Logger
        vectorSize int // embedding dimension, from cfg.Qdrant.VectorSize or default 1536
}

func New(ctx context.Context, cfg *config.Config, repo repository.EmbeddingRepo, qdrantClient *QdrantClient) (*Service, error) {
        var logger *slog.Logger
        if qdrantClient != nil {
                logger = qdrantClient.logger
        } else {
                logger = slog.Default()
        }

        // Determine embedding dimension: use config when Qdrant is in play,
        // otherwise fall back to gemini-embedding-2's 1536 output.
        vectorSize := defaultEmbeddingDimensions
        if cfg.Qdrant.VectorSize > 0 {
                vectorSize = cfg.Qdrant.VectorSize
        }

        return &Service{
                cfg:        cfg,
                repo:       repo,
                qdrant:     qdrantClient,
                sem:        make(chan struct{}, 2),
                logger:     logger,
                vectorSize: vectorSize,
        }, nil
}

type Router9EmbeddingRequest struct {
        Model      string `json:"model"`
        Input      string `json:"input"`
        Dimensions int    `json:"dimensions,omitempty"`
}

type Router9EmbeddingResponse struct {
        Data []struct {
                Embedding []float32 `json:"embedding"`
        } `json:"data"`
}

// GenerateEmbedding calls Router9 embedding API to generate embedding vector for text.
// Returns the embedding vector, a boolean indicating if a mock fallback was used, and any error.
func (s *Service) GenerateEmbedding(ctx context.Context, text string) ([]float32, bool, error) {
        select {
        case <-ctx.Done():
                return nil, false, ctx.Err()
        case s.sem <- struct{}{}:
        }
        defer func() { <-s.sem }()

        modelName := s.cfg.Router9.Model
        if modelName == "" || modelName == "gpt-4o" {
                modelName = "openrouter/openai/text-embedding-3-large"
        }

        reqBody := Router9EmbeddingRequest{
                Model:      modelName,
                Input:      text,
                Dimensions: s.vectorSize,
        }
        bodyBytes, err := json.Marshal(reqBody)
        if err != nil {
                return nil, false, fmt.Errorf("marshal request: %w", err)
        }

        url := fmt.Sprintf("%s/embeddings", s.cfg.Router9.BaseURL)
        req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(bodyBytes))
        if err != nil {
                return nil, false, fmt.Errorf("create http request: %w", err)
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.cfg.Router9.APIKey))

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                s.logger.Warn("router9 embedding API error, falling back to mock vector", "error", err)
                return s.getMockEmbedding(), true, nil
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
                s.logger.Warn("router9 embedding API non-200 status, falling back to mock vector", "status", resp.StatusCode)
                return s.getMockEmbedding(), true, nil
        }

        var resBody Router9EmbeddingResponse
        if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
                return nil, false, fmt.Errorf("decode response: %w", err)
        }

        if len(resBody.Data) == 0 {
                return nil, false, fmt.Errorf("empty embedding returned from router9")
        }

        values := resBody.Data[0].Embedding
        if len(values) != s.vectorSize {
                return nil, false, fmt.Errorf("unexpected embedding dimension from router9: got %d, want %d", len(values), s.vectorSize)
        }

        return values, false, nil
}

func (s *Service) getMockEmbedding() []float32 {
        mockVec := make([]float32, s.vectorSize)
        for i := range mockVec {
                mockVec[i] = float32(i) / float32(s.vectorSize)
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
