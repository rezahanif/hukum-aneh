package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"github.com/rezahanif/hukum-aneh/backend/internal/config"
)

// Quick tool to count documents in Firestore collections.
// Usage: go run ./backend/cmd/count
func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	var opts []option.ClientOption
	if cfg.Firebase.CredentialsPath != "" {
		opts = append(opts, option.WithCredentialsFile(cfg.Firebase.CredentialsPath))
	}
	client, err := firestore.NewClient(ctx, cfg.Firebase.ProjectID, opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "firestore: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	logger := slog.Default()

	collections := []string{
		"laws",
		"content_drafts",
		"image_assets",
		"approvals",
		"publishing_jobs",
		"embeddings",
	}

	for _, col := range collections {
		snapshots, err := client.Collection(col).Documents(ctx).GetAll()
		if err != nil {
			logger.Error("count failed", "collection", col, "error", err)
			continue
		}

		// Count by status for laws
		if col == "laws" {
			statusCount := make(map[string]int)
			for _, s := range snapshots {
				status, _ := s.Data()["status"].(string)
				statusCount[status]++
			}
			fmt.Printf("%-20s %d docs\n", col, len(snapshots))
			for status, count := range statusCount {
				fmt.Printf("  └─ status=%s: %d\n", status, count)
			}
		} else {
			fmt.Printf("%-20s %d docs\n", col, len(snapshots))
		}
	}
}
