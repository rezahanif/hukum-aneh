package retrieval

import (
        "context"
        "fmt"
        "log/slog"
        "strconv"

        "github.com/qdrant/go-client/qdrant"
)

const defaultQdrantCollection = "law_embeddings"

// QdrantClient wraps the qdrant-go-client for the hukum-aneh vector store.
//
// Responsibilities:
//   - EnsureCollection: idempotent creation of the law_embeddings collection
//     with cosine distance + HNSW index, plus payload indexes on
//     law_document_id (keyword) and is_mock (bool) for fast filtering.
//   - Upsert: store a single vector with payload {law_document_id, is_mock}.
//   - Search: top-N similarity search, excluding is_mock=true vectors.
//   - Delete: remove a vector by point ID (used when re-embedding).
type QdrantClient struct {
        client     *qdrant.Client
        collection string
        vectorSize uint64 // matches the embedding model output dim (e.g. 1536 for gemini-embedding-2, 768 for text-embedding-004)
        logger     *slog.Logger
}

// NewQdrantClient creates the client and ensures the collection exists.
// Call this once at app startup. apiKey may be empty for local dev.
// host and port map to the Qdrant gRPC endpoint (default: localhost:6334).
// vectorSize must match the embedding model output dimension.
func NewQdrantClient(ctx context.Context, host string, port int, collection string, apiKey string, vectorSize int, logger *slog.Logger) (*QdrantClient, error) {
        if collection == "" {
                collection = defaultQdrantCollection
        }
        if vectorSize <= 0 {
                vectorSize = defaultEmbeddingDimensions // fallback to 1536 (gemini-embedding-2)
        }

        client, err := qdrant.NewClient(&qdrant.Config{
                Host:   host,
                Port:   port,
                APIKey: apiKey,
                UseTLS: apiKey != "",
        })
        if err != nil {
                return nil, fmt.Errorf("create qdrant client: %w", err)
        }

        qc := &QdrantClient{
                client:     client,
                collection: collection,
                vectorSize: uint64(vectorSize),
                logger:     logger,
        }

        if err := qc.EnsureCollection(ctx); err != nil {
                client.Close()
                return nil, fmt.Errorf("ensure collection: %w", err)
        }

        logger.Info("qdrant client ready",
                "host", host,
                "port", port,
                "collection", collection,
                "vector_size", vectorSize,
        )
        return qc, nil
}

// EnsureCollection creates the collection + payload indexes if they don't exist.
// Idempotent — safe to call on every startup.
func (q *QdrantClient) EnsureCollection(ctx context.Context) error {
        exists, err := q.client.CollectionExists(ctx, q.collection)
        if err != nil {
                return fmt.Errorf("check collection exists: %w", err)
        }
        if exists {
                return nil
        }

        // Create collection with cosine distance + HNSW index
        vectorsConfig := qdrant.NewVectorsConfig(&qdrant.VectorParams{
                Size:     q.vectorSize,
                Distance: qdrant.Distance_Cosine,
        })

        err = q.client.CreateCollection(ctx, &qdrant.CreateCollection{
                CollectionName: q.collection,
                VectorsConfig:  vectorsConfig,
        })
        if err != nil {
                return fmt.Errorf("create collection: %w", err)
        }

        // Payload index on law_document_id (keyword type) for fast lookup
        if _, err := q.client.CreateFieldIndex(ctx, &qdrant.CreateFieldIndexCollection{
                CollectionName: q.collection,
                FieldName:      "law_document_id",
                FieldType:      qdrant.FieldType_FieldTypeKeyword.Enum(),
        }); err != nil {
                return fmt.Errorf("create law_document_id index: %w", err)
        }

        // Payload index on is_mock (bool type) for fast filtering
        if _, err := q.client.CreateFieldIndex(ctx, &qdrant.CreateFieldIndexCollection{
                CollectionName: q.collection,
                FieldName:      "is_mock",
                FieldType:      qdrant.FieldType_FieldTypeBool.Enum(),
        }); err != nil {
                return fmt.Errorf("create is_mock index: %w", err)
        }

        q.logger.Info("qdrant collection created",
                "name", q.collection,
                "vector_size", q.vectorSize,
                "distance", "cosine",
        )
        return nil
}

// Upsert stores a single vector with payload.
// pointID should be a UUID string (matches embedding_metadata.id in PG).
// isMock=true vectors are stored but excluded from search results via Search filter.
func (q *QdrantClient) Upsert(ctx context.Context, pointID string, vector []float32, lawDocumentID string, isMock bool) error {
        if len(vector) != int(q.vectorSize) {
                return fmt.Errorf("vector dim mismatch: got %d, want %d", len(vector), q.vectorSize)
        }

        points := []*qdrant.PointStruct{
                {
                        Id:      qdrant.NewIDUUID(pointID),
                        Vectors: qdrant.NewVectorsDense(vector),
                        Payload: qdrant.NewValueMap(map[string]interface{}{
                                "law_document_id": lawDocumentID,
                                "is_mock":         isMock,
                        }),
                },
        }

        _, err := q.client.Upsert(ctx, &qdrant.UpsertPoints{
                CollectionName: q.collection,
                Points:         points,
        })
        if err != nil {
                return fmt.Errorf("upsert point %s: %w", pointID, err)
        }
        return nil
}

// Search returns top-N similar vectors, excluding is_mock=true payloads.
// Score field is cosine similarity in range [-1, 1] (higher = more similar).
func (q *QdrantClient) Search(ctx context.Context, queryVector []float32, topN int) ([]SearchResult, error) {
        if len(queryVector) != int(q.vectorSize) {
                return nil, fmt.Errorf("query vector dim mismatch: got %d, want %d", len(queryVector), q.vectorSize)
        }
        if topN <= 0 {
                topN = 5
        }

        // Filter: exclude is_mock=true vectors
        filter := &qdrant.Filter{
                MustNot: []*qdrant.Condition{
                        qdrant.NewMatchBool("is_mock", true),
                },
        }

        scoredPoints, err := q.client.Query(ctx, &qdrant.QueryPoints{
                CollectionName: q.collection,
                Query:          qdrant.NewQueryDense(queryVector),
                Limit:          qdrant.PtrOf(uint64(topN)),
                Filter:         filter,
                WithPayload:    qdrant.NewWithPayload(true),
        })
        if err != nil {
                return nil, fmt.Errorf("qdrant query: %w", err)
        }

        results := make([]SearchResult, 0, len(scoredPoints))
        for _, point := range scoredPoints {
                lawDocID := ""
                if val, ok := point.Payload["law_document_id"]; ok {
                        lawDocID = val.GetStringValue()
                }
                results = append(results, SearchResult{
                        LawDocumentID: lawDocID,
                        Score:         point.Score,
                })
        }
        return results, nil
}

// Delete removes a single vector by point ID.
// Used when re-embedding a law (old vector deleted, new one upserted).
func (q *QdrantClient) Delete(ctx context.Context, pointID string) error {
        _, err := q.client.Delete(ctx, &qdrant.DeletePoints{
                CollectionName: q.collection,
                Points: &qdrant.PointsSelector{
                        PointsSelectorOneOf: &qdrant.PointsSelector_Points{
                                Points: &qdrant.PointsIdsList{
                                        Ids: []*qdrant.PointId{qdrant.NewIDUUID(pointID)},
                                },
                        },
                },
        })
        if err != nil {
                return fmt.Errorf("delete point %s: %w", pointID, err)
        }
        return nil
}

// Close cleans up the underlying gRPC connection.
func (q *QdrantClient) Close() error {
        return q.client.Close()
}

// CollectionName returns the configured collection name.
func (q *QdrantClient) CollectionName() string {
        return q.collection
}

// VectorSize returns the configured vector dimension size.
func (q *QdrantClient) VectorSize() int {
        return int(q.vectorSize)
}

// Helper: safely parse score from ScoredPoint payload.
// Not currently used externally but useful for debugging.
func payloadStringValue(payload map[string]*qdrant.Value, key string) string {
        if val, ok := payload[key]; ok {
                return val.GetStringValue()
        }
        return ""
}

func payloadBoolValue(payload map[string]*qdrant.Value, key string) bool {
        if val, ok := payload[key]; ok {
                // qdrant Value stores bools as integer (0/1) or as bool type
                switch v := val.GetKind().(type) {
                case *qdrant.Value_BoolValue:
                        return v.BoolValue
                case *qdrant.Value_IntegerValue:
                        return v.IntegerValue != 0
                case *qdrant.Value_StringValue:
                        b, _ := strconv.ParseBool(v.StringValue)
                        return b
                default:
                        return false
                }
        }
        return false
}
