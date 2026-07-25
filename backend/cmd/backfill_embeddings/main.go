package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
)

type queuedEmbedding struct {
	Embedding *models.EmbeddingEntry `json:"embedding"`
	Error     string                 `json:"error"`
	QueuedAt  time.Time              `json:"queued_at"`
}

func main() {
	var (
		limit      int
		workers    int
		queueDir   string
		verbose    bool
	)
	flag.IntVar(&limit, "limit", 0, "maximum number of laws to backfill (0 for all)")
	flag.IntVar(&workers, "workers", 1, "number of concurrent worker goroutines")
	flag.StringVar(&queueDir, "queue", "backend/internal/storage/local_queue_embeddings", "directory for local queue fallback")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose debug logging")
	flag.Parse()

	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	repo, err := repository.NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath)
	if err != nil {
		log.Fatalf("firestore: %v", err)
	}
	defer repo.Close()

	ret, err := retrieval.New(ctx, cfg, repo)
	if err != nil {
		log.Fatalf("retrieval: %v", err)
	}

	fmt.Println("Loading laws and embeddings from Firestore...")
	laws, err := repo.ListAllLaws(ctx)
	if err != nil {
		log.Fatalf("list laws: %v", err)
	}
	embeddings, err := repo.ListAllEmbeddings(ctx)
	if err != nil {
		log.Fatalf("list embeddings: %v", err)
	}

	// Index valid (real + has CreatedAt) embeddings by LawDocumentID
	validEmbs := make(map[string]bool)
	for _, e := range embeddings {
		if !e.IsMock && !e.CreatedAt.IsZero() {
			validEmbs[e.LawDocumentID] = true
		}
	}

	// Filter laws needing backfill (Status >= "parsed" and no valid embedding)
	var targetLaws []models.LawDocument
	for _, law := range laws {
		if law.Status != "discovered" && law.Status != "downloaded" {
			if !validEmbs[law.ID] {
				targetLaws = append(targetLaws, law)
			}
		}
	}

	totalTargets := len(targetLaws)
	fmt.Printf("Total parsed laws: %d\n", len(laws))
	fmt.Printf("Total existing embeddings: %d (valid real: %d)\n", len(embeddings), len(validEmbs))
	fmt.Printf("Parsed laws needing embedding: %d\n", totalTargets)

	if limit > 0 && limit < totalTargets {
		targetLaws = targetLaws[:limit]
		totalTargets = limit
		fmt.Printf("Limiting run to first %d targets\n", limit)
	}

	if totalTargets == 0 {
		fmt.Println("All parsed laws have valid embeddings. Nothing to do!")
		return
	}

	// Channel to feed target laws to workers
	jobs := make(chan models.LawDocument, totalTargets)
	for _, l := range targetLaws {
		jobs <- l
	}
	close(jobs)

	var (
		processedCount int64
		successCount   int64
		failCount      int64
		localQueueCount int64
	)

	// Concurrency group
	var wg sync.WaitGroup
	wg.Add(workers)

	fmt.Printf("Starting backfill with %d workers...\n", workers)
	startTime := time.Now()

	for i := 0; i < workers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for law := range jobs {
				// Process single law
				currProcessed := atomic.AddInt64(&processedCount, 1)

				// Log progress every 100 docs
				if currProcessed%100 == 0 || currProcessed == 1 {
					fmt.Printf("[%s] Progress: %d/%d laws processed (success: %d, fail: %d, local_queued: %d)\n",
						time.Now().Format("15:04:05"), currProcessed, totalTargets,
						atomic.LoadInt64(&successCount), atomic.LoadInt64(&failCount), atomic.LoadInt64(&localQueueCount))
				}

				version, err := repo.GetLatestLawVersion(ctx, law.ID)
				if err != nil {
					atomic.AddInt64(&failCount, 1)
					if verbose {
						log.Printf("Worker %d: GetLatestLawVersion failed for law ID %s: %v", workerID, law.ID, err)
					}
					continue
				}

				textContent := strings.TrimSpace(version.TextContent)
				if textContent == "" {
					atomic.AddInt64(&failCount, 1)
					if verbose {
						log.Printf("Worker %d: Empty text content for law ID %s", workerID, law.ID)
					}
					continue
				}

				// Generate embedding via Gemini API
				vector, isMock, err := ret.GenerateEmbedding(ctx, textContent)
				if err != nil {
					atomic.AddInt64(&failCount, 1)
					log.Printf("Worker %d: GenerateEmbedding failed for law ID %s: %v", workerID, law.ID, err)
					continue
				}

				embEntry := &models.EmbeddingEntry{
					LawDocumentID: law.ID,
					Vector:        vector,
					IsMock:        isMock,
					CreatedAt:     time.Now(),
				}

				// Save embedding
				_, saveErr := repo.SaveEmbedding(ctx, embEntry)
				if saveErr != nil {
					// Firestore write failure — fallback to local queue
					log.Printf("Worker %d: SaveEmbedding failed for law ID %s (falling back to local queue): %v", workerID, law.ID, saveErr)
					if qErr := saveLocalEmbedding(queueDir, embEntry, saveErr.Error()); qErr != nil {
						log.Printf("Worker %d: Failed to save to local queue: %v", workerID, qErr)
						atomic.AddInt64(&failCount, 1)
					} else {
						atomic.AddInt64(&localQueueCount, 1)
					}
					continue
				}

				atomic.AddInt64(&successCount, 1)
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(startTime)

	fmt.Printf("\nBackfill complete! Duration: %v\n", duration)
	fmt.Printf("Total processed: %d\n", processedCount)
	fmt.Printf("Successfully saved to Firestore: %d\n", successCount)
	fmt.Printf("Saved to local fallback queue: %d\n", localQueueCount)
	fmt.Printf("Failed: %d\n", failCount)
}

func saveLocalEmbedding(dir string, emb *models.EmbeddingEntry, reason string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	q := queuedEmbedding{
		Embedding: emb,
		Error:     reason,
		QueuedAt:  time.Now(),
	}
	b, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return err
	}
	// Safe filename based on document ID
	name := strings.NewReplacer(" ", "_", "/", "_", "\\", "_", ":", "_").Replace(emb.LawDocumentID)
	return os.WriteFile(filepath.Join(dir, name+".json"), b, 0644)
}
