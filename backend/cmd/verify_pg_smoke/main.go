// verify_pg_smoke — Phase 6.2 + 7.3 verification
//
// Verifies that PostgresRepo correctly implements all 8 interfaces against
// a real PostgreSQL instance with the schema applied.
//
// Run via:
//
//	source /home/z/my-project/scripts/build-env.sh
//	cd /home/z/my-project/hukum-aneh
//	STORAGE_MODE=postgres \
//	POSTGRES_HOST=localhost POSTGRES_PORT=15432 \
//	POSTGRES_DB=hukum_aneh POSTGRES_USER=hukum \
//	go run ./backend/cmd/verify_pg_smoke
//
// Exit codes:
//
//	0 = all assertions passed
//	1 = setup error (config, PG connection)
//	2 = assertion failure (data integrity issue)
package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		logger.Error("config", "error", err)
		os.Exit(1)
	}
	cfg.StorageMode = config.StorageModePostgres

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	repo, err := repository.NewPostgresRepo(ctx, cfg)
	if err != nil {
		logger.Error("postgres open", "error", err)
		os.Exit(1)
	}
	defer repo.Close()
	logger.Info("connected to PG", "host", cfg.Postgres.Host, "port", cfg.Postgres.Port)

	repos := repository.NewRepoSetFromPostgres(repo)
	var failures int

	// === LawDocumentRepo ===
	logger.Info("=== LawDocumentRepo ===")
	seedLaw := &models.LawDocument{
		ID:            "smoke-test-uu-13-2024",
		LawNumber:     "UU 13/2024",
		Title:         "Smoke Test Law",
		SourceURL:     "https://example.test/uu-13-2024",
		Source:        "smoke-test",
		Level:         "national",
		DocumentType:  "UU",
		PublishedDate: "2024-08-15",
		Status:        "discovered",
	}
	if _, err := repos.LawRepo.SaveLawDocument(ctx, seedLaw); err != nil {
		logger.Error("SaveLawDocument failed", "error", err)
		failures++
	}

	got, err := repos.LawRepo.GetLawDocument(ctx, seedLaw.ID)
	if err != nil || got.LawNumber != seedLaw.LawNumber {
		logger.Error("GetLawDocument mismatch", "err", err, "got", got)
		failures++
	} else {
		logger.Info("GetLawDocument OK", "id", got.ID, "law_number", got.LawNumber)
	}

	found, err := repos.LawRepo.FindByLawNumber(ctx, seedLaw.LawNumber)
	if err != nil || found == nil {
		logger.Error("FindByLawNumber failed", "err", err, "found", found)
		failures++
	} else {
		logger.Info("FindByLawNumber OK", "id", found.ID)
	}

	none, err := repos.LawRepo.FindByLawNumber(ctx, "DOES-NOT-EXIST-9999")
	if err != nil || none != nil {
		logger.Error("FindByLawNumber (not found) should return nil", "err", err, "none", none)
		failures++
	} else {
		logger.Info("FindByLawNumber (not found) correctly returned nil")
	}

	// Update status, then list
	seedLaw.Status = "downloaded"
	if _, err := repos.LawRepo.SaveLawDocument(ctx, seedLaw); err != nil {
		logger.Error("SaveLawDocument (update) failed", "error", err)
		failures++
	}
	byStatus, err := repos.LawRepo.ListLawsByStatus(ctx, "downloaded")
	if err != nil || len(byStatus) < 1 {
		logger.Error("ListLawsByStatus failed", "err", err, "count", len(byStatus))
		failures++
	} else {
		logger.Info("ListLawsByStatus OK", "count", len(byStatus))
	}

	stuck, err := repos.LawRepo.FindStuckDocuments(ctx, "downloaded", time.Now().Add(time.Hour))
	if err != nil || len(stuck) < 1 {
		logger.Error("FindStuckDocuments failed", "err", err, "count", len(stuck))
		failures++
	} else {
		logger.Info("FindStuckDocuments OK", "count", len(stuck))
	}

	// === LawVersionRepo ===
	logger.Info("=== LawVersionRepo ===")
	seedVersion := &models.LawVersion{
		ID:            "smoke-test-version-1",
		LawDocumentID: seedLaw.ID,
		VersionNumber: 1,
		TextContent:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}
	if _, err := repos.VersionRepo.SaveLawVersion(ctx, seedLaw.ID, seedVersion); err != nil {
		logger.Error("SaveLawVersion failed", "error", err)
		failures++
	}
	latestV, err := repos.VersionRepo.GetLatestLawVersion(ctx, seedLaw.ID)
	if err != nil || latestV.VersionNumber != 1 {
		logger.Error("GetLatestLawVersion mismatch", "err", err, "v", latestV)
		failures++
	} else {
		logger.Info("GetLatestLawVersion OK", "version", latestV.VersionNumber)
	}

	// === LawAnalysisRepo + ContentDraftRepo (chained — analysis needs draft for retrieval) ===
	logger.Info("=== LawAnalysisRepo ===")
	seedAnalysis := &models.LawAnalysis{
		ID:               "smoke-test-analysis-1",
		LawDocumentID:    seedLaw.ID,
		Summary:          "Test analysis summary",
		AffectedLaws:     []models.AffectedLaw{{Law: "UU 12/2024", Article: "5", Reason: "Test", Severity: 0.7}},
		OverallScore:     75,
		ControversyScore: 60,
		EconomicScore:    80,
		LegalConsistency: 85,
		Confidence:       0.92,
		RawJSON:          `{"test": true}`,
	}
	if _, err := repos.AnalysisRepo.SaveLawAnalysis(ctx, seedLaw.ID, seedAnalysis); err != nil {
		logger.Error("SaveLawAnalysis failed", "error", err)
		failures++
	}

	logger.Info("=== ContentDraftRepo ===")
	seedDraft := &models.ContentDraft{
		ID:            "smoke-test-draft-1",
		LawAnalysisID: seedAnalysis.ID,
		Caption:       "Test caption with #hashtags",
		Hashtags:      []string{"#hukum", "#test"},
		Hook:          "Hook line",
		ImagePrompt:   "image prompt",
		Status:        "draft",
	}
	if _, err := repos.DraftRepo.SaveContentDraft(ctx, seedDraft); err != nil {
		logger.Error("SaveContentDraft failed", "error", err)
		failures++
	}
	gotDraft, err := repos.DraftRepo.GetContentDraft(ctx, seedDraft.ID)
	if err != nil || len(gotDraft.Hashtags) != 2 {
		logger.Error("GetContentDraft mismatch", "err", err, "hashtags", gotDraft)
		failures++
	} else {
		logger.Info("GetContentDraft OK", "hashtags", gotDraft.Hashtags)
	}

	gotA, err := repos.AnalysisRepo.GetLawAnalysisByDraft(ctx, seedDraft.ID)
	if err != nil || gotA == nil || len(gotA.AffectedLaws) != 1 {
		logger.Error("GetLawAnalysisByDraft mismatch", "err", err, "got", gotA)
		failures++
	} else {
		logger.Info("GetLawAnalysisByDraft OK", "affected_count", len(gotA.AffectedLaws))
	}

	// === ImageAssetRepo ===
	logger.Info("=== ImageAssetRepo ===")
	seedAsset := &models.ImageAsset{
		ID:                 "smoke-test-asset-1",
		ContentDraftID:     seedDraft.ID,
		PromptUsed:         "test prompt",
		FilePath:           "/tmp/test.png",
		Validated:          true,
		DesignGuideVersion: "v1",
	}
	if _, err := repos.ImageRepo.SaveImageAsset(ctx, seedAsset); err != nil {
		logger.Error("SaveImageAsset failed", "error", err)
		failures++
	}
	assets, err := repos.ImageRepo.GetImageAssetsByDraft(ctx, seedDraft.ID)
	if err != nil || len(assets) != 1 {
		logger.Error("GetImageAssetsByDraft mismatch", "err", err, "count", len(assets))
		failures++
	} else {
		logger.Info("GetImageAssetsByDraft OK", "count", len(assets))
	}

	// === ApprovalRepo ===
	logger.Info("=== ApprovalRepo ===")
	seedApproval := &models.Approval{
		ID:             "smoke-test-approval-1",
		ContentDraftID: seedDraft.ID,
		ReviewerID:     "telegram-user-123",
		Decision:       "approve",
		Reason:         "LGTM",
	}
	if _, err := repos.ApprovalRepo.SaveApproval(ctx, seedApproval); err != nil {
		logger.Error("SaveApproval failed", "error", err)
		failures++
	} else {
		logger.Info("SaveApproval OK")
	}

	// === PublishingJobRepo (with CHECK constraint verification) ===
	logger.Info("=== PublishingJobRepo ===")
	seedJob := &models.PublishingJob{
		ID:             "smoke-test-job-1",
		ContentDraftID: seedDraft.ID,
		Platform:       "instagram",
		Status:         "pending",
	}
	if _, err := repos.PublishRepo.SavePublishingJob(ctx, seedJob); err != nil {
		logger.Error("SavePublishingJob failed", "error", err)
		failures++
	} else {
		logger.Info("SavePublishingJob OK (instagram)")
	}

	tiktokJob := &models.PublishingJob{
		ID:             "smoke-test-job-tiktok",
		ContentDraftID: seedDraft.ID,
		Platform:       "tiktok",
		Status:         "pending",
	}
	if _, err := repos.PublishRepo.SavePublishingJob(ctx, tiktokJob); err == nil {
		logger.Error("SavePublishingJob allowed tiktok — CHECK constraint not enforced!")
		failures++
	} else {
		logger.Info("SavePublishingJob correctly rejected tiktok platform")
	}

	// === EmbeddingRepo ===
	logger.Info("=== EmbeddingRepo ===")
	seedEmb := &models.EmbeddingEntry{
		ID:            "smoke-test-emb-1",
		LawDocumentID: seedLaw.ID,
		Vector:        []float32{0.1, 0.2, 0.3}, // dropped in PG — vector lives in Qdrant
		IsMock:        false,
	}
	if _, err := repos.EmbedRepo.SaveEmbedding(ctx, seedEmb); err != nil {
		logger.Error("SaveEmbedding failed", "error", err)
		failures++
	}
	allEmbs, err := repos.EmbedRepo.ListAllEmbeddings(ctx)
	if err != nil || len(allEmbs) != 1 || allEmbs[0].Vector != nil {
		logger.Error("ListAllEmbeddings mismatch", "err", err, "count", len(allEmbs))
		failures++
	} else {
		logger.Info("ListAllEmbeddings OK", "count", len(allEmbs), "vector_nil", allEmbs[0].Vector == nil)
	}

	byDocType, err := repos.EmbedRepo.ListEmbeddingsByDocType(ctx, "UU")
	if err != nil || len(byDocType) != 1 {
		logger.Error("ListEmbeddingsByDocType mismatch", "err", err, "count", len(byDocType))
		failures++
	} else {
		logger.Info("ListEmbeddingsByDocType OK", "count", len(byDocType))
	}

	// ListEmbeddingsBatch should error in PG mode (Firestore-specific signature)
	if _, _, err := repos.EmbedRepo.ListEmbeddingsBatch(ctx, nil, 10); err == nil {
		logger.Error("ListEmbeddingsBatch should have returned error in PG mode")
		failures++
	} else {
		logger.Info("ListEmbeddingsBatch correctly returned error")
	}

	// === ListAllLaws (final count check) ===
	allLaws, err := repos.LawRepo.ListAllLaws(ctx)
	if err != nil || len(allLaws) < 1 {
		logger.Error("ListAllLaws failed", "err", err, "count", len(allLaws))
		failures++
	} else {
		logger.Info("ListAllLaws OK", "count", len(allLaws))
	}

	// === Cleanup ===
	logger.Info("=== Cleanup ===")
	_, err = repo.Pool().Exec(ctx, `
		TRUNCATE embedding_metadata, publishing_jobs, approvals, image_assets,
		        captions, content_drafts, law_analyses, law_relationships,
		        law_versions, law_documents
		CASCADE
	`)
	if err != nil {
		logger.Warn("cleanup failed (manual TRUNCATE may be needed)", "error", err)
	} else {
		logger.Info("cleanup OK — all test rows removed")
	}

	if failures > 0 {
		logger.Error("SMOKE TEST FAILED", "failures", failures)
		os.Exit(2)
	}
	logger.Info("SMOKE TEST PASSED — all assertions OK")
}
