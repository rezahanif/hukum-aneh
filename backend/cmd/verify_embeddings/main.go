package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	repo, err := repository.NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	repos := repository.NewRepoSetFromFirestore(repo)

	laws, err := repos.LawRepo.ListAllLaws(ctx)
	if err != nil {
		log.Fatal(err)
	}
	embeddings, err := repos.EmbedRepo.ListAllEmbeddings(ctx)
	if err != nil {
		log.Fatal(err)
	}

	embByLaw := make(map[string]bool)
	mockCount, realCount := 0, 0
	for _, e := range embeddings {
		embByLaw[e.LawDocumentID] = true
		if e.IsMock {
			mockCount++
		} else {
			realCount++
		}
	}

	var missing, parsedMissing []string
	for _, law := range laws {
		if !embByLaw[law.ID] {
			missing = append(missing, law.ID+" ("+law.Status+")")
			if law.Status != "discovered" && law.Status != "downloaded" {
				parsedMissing = append(parsedMissing, law.ID)
			}
		}
	}

	fmt.Printf("Total laws: %d\n", len(laws))
	fmt.Printf("Total embeddings: %d (real: %d, mock: %d)\n", len(embeddings), realCount, mockCount)
	fmt.Printf("Laws missing any embedding: %d\n", len(missing))
	fmt.Printf("Parsed-or-later laws missing embedding (should be 0): %d\n", len(parsedMissing))
	for _, id := range parsedMissing {
		fmt.Println("  MISSING:", id)
	}
}
