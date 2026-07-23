package workflow

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
	"github.com/rezahanif/hukum-aneh/backend/internal/ai"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/imagegen"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/telegram"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/publishing"
	"github.com/rezahanif/hukum-aneh/backend/internal/validator"
)

// Engine orchestrates the full pipeline. Owns all control flow.
// AI agents are workers the engine calls — they never orchestrate. Spec §2.
type Engine struct {
	cfg        *config.Config
	repo       *repository.FirestoreRepo
	registry   *connectors.Registry
	parser     *parser.Parser
	retrieval  *retrieval.Service
	ai         *ai.Service
	imagegen   *imagegen.Service
	tg         *telegram.Service
	publishing *publishing.Service
	validator  *validator.ImageValidator
	logger     *slog.Logger
}

func NewEngine(
	cfg *config.Config,
	repo *repository.FirestoreRepo,
	registry *connectors.Registry,
	p *parser.Parser,
	ret *retrieval.Service,
	aiSvc *ai.Service,
	imgGen *imagegen.Service,
	tgSvc *telegram.Service,
	pubSvc *publishing.Service,
	val *validator.ImageValidator,
	logger *slog.Logger,
) *Engine {
	return &Engine{
		cfg:        cfg,
		repo:       repo,
		registry:   registry,
		parser:     p,
		retrieval:  ret,
		ai:         aiSvc,
		imagegen:   imgGen,
		tg:         tgSvc,
		publishing: pubSvc,
		validator:  val,
		logger:     logger,
	}
}

// RunDiscovery is the entry point triggered by the Scheduler.
// Iterates all registered connectors, checks for updates, and writes new
// LawDocuments to Firestore. Does NOT parse or analyze — that's event-driven
// off subsequent steps. Spec §3 pipeline.
func (e *Engine) RunDiscovery(ctx context.Context) error {
	e.logger.Info("discovery run started")
	var wg sync.WaitGroup

	for name, conn := range e.registry.All() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e.logger.Info("checking source", "connector", name)
		docs, err := conn.CheckUpdates(ctx)
		if err != nil {
			e.logger.Error("connector check failed", "connector", name, "error", err)
			continue
		}

		for _, meta := range docs {
			existing, err := e.repo.FindByLawNumber(ctx, meta.LawNumber)
			if err != nil {
				e.logger.Error("dup check failed", "law_number", meta.LawNumber, "error", err)
				continue
			}
			if existing != nil {
				continue
			}

			doc := &models.LawDocument{
				LawNumber:     meta.LawNumber,
				Title:         meta.Title,
				SourceURL:     meta.SourceURL,
				Source:        meta.Source,
				Level:         meta.Level,
				DocumentType:  meta.DocumentType,
				PublishedDate: meta.PublishedDate,
				Status:        "discovered",
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
			id, err := e.repo.SaveLawDocument(ctx, doc)
			if err != nil {
				e.logger.Error("save law doc failed", "law_number", meta.LawNumber, "error", err)
				continue
			}
			e.logger.Info("discovered new law", "id", id, "law_number", meta.LawNumber, "title", meta.Title)

			// Trigger download → parse → analyze pipeline for this law
			wg.Add(1)
			go func(lawDoc *models.LawDocument) {
				defer wg.Done()
				procCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
				defer cancel()
				if err := e.ProcessDocument(procCtx, lawDoc); err != nil {
					e.logger.Error("process document failed", "id", lawDoc.ID, "law_number", lawDoc.LawNumber, "error", err)
				}
			}(doc)
		}
	}

	// Wait for all discovered documents to finish processing
	e.logger.Info("waiting for in-flight document processing to complete")
	wg.Wait()

	e.logger.Info("discovery run complete")
	return nil
}

// ProcessDocument handles the download → parse → save pipeline for a single law.
// Called after discovery finds a new law. Spec §3: Document Download → Document Parsing → Save to Database.
func (e *Engine) ProcessDocument(ctx context.Context, doc *models.LawDocument) error {
	conn, ok := e.registry.Get(doc.Source)
	if !ok {
		return fmt.Errorf("no connector for source: %s", doc.Source)
	}

	meta := connectors.DocumentMeta{
		LawNumber:     doc.LawNumber,
		Title:         doc.Title,
		SourceURL:     doc.SourceURL,
		Source:        doc.Source,
		Level:         doc.Level,
		DocumentType:  doc.DocumentType,
		PublishedDate: doc.PublishedDate,
	}

	// Download
	raw, err := conn.Download(ctx, meta)
	if err != nil {
		doc.Status = "download_failed"
		doc.UpdatedAt = time.Now()
		e.repo.SaveLawDocument(ctx, doc)
		return fmt.Errorf("download: %w", err)
	}
	defer raw.Content.Close()

	// Save raw file to storage
	storageDir := "backend/internal/storage"
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return fmt.Errorf("mkdir storage: %w", err)
	}
	rawPath := filepath.Join(storageDir, doc.ID+"_"+raw.Filename)
	rawFile, err := os.Create(rawPath)
	if err != nil {
		return fmt.Errorf("create raw file: %w", err)
	}
	defer rawFile.Close()

	// Tee reader — write to file while parsing
	rawBytes, err := io.ReadAll(raw.Content)
	if err != nil {
		return fmt.Errorf("read raw content: %w", err)
	}
	raw.Content.Close()

	if _, err := rawFile.Write(rawBytes); err != nil {
		return fmt.Errorf("write raw file: %w", err)
	}

	doc.RawFilePath = rawPath
	doc.Status = "downloaded"
	doc.UpdatedAt = time.Now()
	if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
		return fmt.Errorf("save doc status: %w", err)
	}

	// Parse
	result, err := e.parser.Parse(ctx, bytes.NewReader(rawBytes), raw.MimeType, raw.Filename)
	if err != nil {
		doc.Status = "parse_failed"
		doc.UpdatedAt = time.Now()
		e.repo.SaveLawDocument(ctx, doc)
		return fmt.Errorf("parse: %w", err)
	}

	// Save version to Firestore
	version := &models.LawVersion{
		LawDocumentID: doc.ID,
		VersionNumber: int(time.Now().Unix()),
		TextContent:   result.TextContent,
		ParsedAt:      time.Now(),
	}
	if _, err := e.repo.SaveLawVersion(ctx, doc.ID, version); err != nil {
		return fmt.Errorf("save version: %w", err)
	}

	doc.Status = "parsed"
	doc.UpdatedAt = time.Now()
	if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
		return fmt.Errorf("update doc status: %w", err)
	}

	e.logger.Info("document processed", "id", doc.ID, "law_number", doc.LawNumber, "sections", len(result.Sections))

	// Trigger next pipeline steps: Embed → Similarity Search → Analyze → Strategy → Prompt → ImageGen → Approval
	if err := e.ProcessParsedDocument(ctx, doc, version); err != nil {
		e.logger.Error("process parsed document failed", "id", doc.ID, "error", err)
		return fmt.Errorf("parsed pipeline: %w", err)
	}

	return nil
}

// ProcessParsedDocument implements the downstream modular pipeline steps.
// Spec §3: Embed → Search → LawAnalysis → ContentStrategy → PromptBuilder → ImageGen → Validator → Telegram.
func (e *Engine) ProcessParsedDocument(ctx context.Context, doc *models.LawDocument, version *models.LawVersion) error {
	e.logger.Info("generating embedding for law version", "doc_id", doc.ID)

	// Step 1: Generate & Save Embedding
	vector, err := e.retrieval.GenerateEmbedding(ctx, version.TextContent)
	if err != nil {
		return fmt.Errorf("generate embedding: %w", err)
	}

	embEntry := &models.EmbeddingEntry{
		LawDocumentID: doc.ID,
		Vector:        vector,
	}
	embID, err := e.repo.SaveEmbedding(ctx, embEntry)
	if err != nil {
		return fmt.Errorf("save embedding: %w", err)
	}

	version.EmbeddingID = embID
	if _, err := e.repo.SaveLawVersion(ctx, doc.ID, version); err != nil {
		return fmt.Errorf("update version embedding_id: %w", err)
	}

	// Step 2: Similarity Search
	e.logger.Info("running similarity search against other laws")
	candidates, err := e.retrieval.Search(ctx, vector, 3)
	if err != nil {
		e.logger.Warn("similarity search failed, continuing without related laws", "error", err)
	}

	var relatedTexts []string
	for _, c := range candidates {
		// Do not compare against self
		if c.LawDocumentID == doc.ID {
			continue
		}
		// Try to fetch related law text
		relDoc, err := e.repo.GetLawDocument(ctx, c.LawDocumentID)
		if err != nil {
			continue
		}
		relatedTexts = append(relatedTexts, fmt.Sprintf("Law: %s, Title: %s", relDoc.LawNumber, relDoc.Title))
	}

	// Step 3: Law Analysis Agent (Hermes LLM worker)
	e.logger.Info("calling Law Analysis Agent")
	analysis, err := e.ai.AnalyzeLaw(ctx, version.TextContent, relatedTexts)
	if err != nil {
		return fmt.Errorf("law analysis agent: %w", err)
	}
	analysis.LawDocumentID = doc.ID

	analysisID, err := e.repo.SaveLawAnalysis(ctx, doc.ID, analysis)
	if err != nil {
		return fmt.Errorf("save law analysis: %w", err)
	}
	analysis.ID = analysisID

	// GATE: Only proceed to content generation if the law is "suspicious"
	// (i.e. has high controversy, low consistency, or high conflict severity)
	if !e.isSuspicious(analysis) {
		doc.Status = "no_conflict"
		doc.UpdatedAt = time.Now()
		if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
			return fmt.Errorf("update doc status: %w", err)
		}
		e.logger.Info("law has no significant conflict or controversy, stopping pipeline", "doc_id", doc.ID, "law_number", doc.LawNumber)
		return nil
	}

	doc.Status = "analyzed"
	doc.UpdatedAt = time.Now()
	if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
		return fmt.Errorf("update doc status: %w", err)
	}

	// Step 4: Content Strategy Agent
	e.logger.Info("calling Content Strategy Agent")
	draft, err := e.ai.CreateContentStrategy(ctx, analysis)
	if err != nil {
		return fmt.Errorf("content strategy agent: %w", err)
	}
	draft.Status = "draft"

	draftID, err := e.repo.SaveContentDraft(ctx, draft)
	if err != nil {
		return fmt.Errorf("save content draft: %w", err)
	}
	draft.ID = draftID

	// Step 5: Prompt Builder Agent
	e.logger.Info("calling Prompt Builder Agent")
	designGuide, err := os.ReadFile("backend/internal/prompts/design_guide.json")
	if err != nil {
		return fmt.Errorf("read design guide: %w", err)
	}
	characterSheet, err := os.ReadFile("backend/internal/prompts/character_sheet.json")
	if err != nil {
		return fmt.Errorf("read character sheet: %w", err)
	}

	imagePrompt, err := e.ai.BuildImagePrompt(ctx, draft, string(designGuide), string(characterSheet))
	if err != nil {
		return fmt.Errorf("prompt builder agent: %w", err)
	}

	// Step 6: Image Generation Service
	e.logger.Info("calling Image Generation Service", "prompt", imagePrompt)
	asset, err := e.imagegen.GenerateImage(ctx, draft.ID, imagePrompt)
	if err != nil {
		return fmt.Errorf("image generation: %w", err)
	}

	// Step 7: Image Validation
	e.logger.Info("validating generated image")
	valid, err := e.validator.Validate(asset.FilePath)
	if err != nil {
		e.logger.Error("image validation returned error", "error", err)
	}
	asset.Validated = valid

	assetID, err := e.repo.SaveImageAsset(ctx, asset)
	if err != nil {
		return fmt.Errorf("save image asset: %w", err)
	}
	asset.ID = assetID

	// Step 8: Send Telegram Approval Request
	e.logger.Info("sending Telegram approval request")
	_, err = e.tg.SendApprovalRequest(ctx, draft, asset, analysis, doc.Title)
	if err != nil {
		return fmt.Errorf("send telegram approval: %w", err)
	}

	draft.Status = "pending_approval"
	if _, err := e.repo.SaveContentDraft(ctx, draft); err != nil {
		return fmt.Errorf("update draft status: %w", err)
	}

	doc.Status = "pending_approval"
	doc.UpdatedAt = time.Now()
	if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
		return fmt.Errorf("update doc status: %w", err)
	}

	e.logger.Info("pipeline completed up to approval stage", "doc_id", doc.ID)
	return nil
}

// HandleApprovalAction processes the callback query from Telegram.
// Decisions: approve, reject, regen_img, regen_cap.
func (e *Engine) HandleApprovalAction(ctx context.Context, draftID string, action string, reviewerID string) error {
	e.logger.Info("processing approval action", "draft_id", draftID, "action", action, "reviewer_id", reviewerID)

	draft, err := e.repo.GetContentDraft(ctx, draftID)
	if err != nil {
		return fmt.Errorf("get content draft: %w", err)
	}

	analysis, err := e.repo.GetLawAnalysisByDraft(ctx, draftID)
	if err != nil {
		return fmt.Errorf("get law analysis: %w", err)
	}

	doc, err := e.repo.GetLawDocument(ctx, analysis.LawDocumentID)
	if err != nil {
		return fmt.Errorf("get law document: %w", err)
	}

	// Save Approval record
	approval := &models.Approval{
		ContentDraftID: draftID,
		ReviewerID:     reviewerID,
		Decision:       action,
		Reason:         fmt.Sprintf("Action triggered via Telegram inline keyboard"),
		Timestamp:      time.Now(),
	}
	if _, err := e.repo.SaveApproval(ctx, approval); err != nil {
		e.logger.Error("failed to save approval log", "error", err)
	}

	switch action {
	case "approve":
		draft.Status = "approved"
		if _, err := e.repo.SaveContentDraft(ctx, draft); err != nil {
			return err
		}

		doc.Status = "approved"
		if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
			return err
		}

		// Spec §5.11: Trigger Publishing Engine
		e.logger.Info("publishing approved content to Instagram", "draft_id", draftID)

		// Fetch asset for the image
		assets, err := e.repo.GetImageAssetsByDraft(ctx, draftID)
		if err != nil || len(assets) == 0 {
			return fmt.Errorf("no image asset for draft: %w", err)
		}
		asset := &assets[0]

		// ponytail: Instagram Graph API requires public image URL.
		// upgrade: upload to S3/GCS/Imgur first, then pass public URL here.
		publicImageURL := e.cfg.Instagram.AccessToken // placeholder — real impl needs image hosting
		if publicImageURL == "" {
			return fmt.Errorf("instagram credentials or public image URL missing")
		}

		postID, err := e.publishing.PublishToInstagram(ctx, draft, asset, publicImageURL)
		if err != nil {
			e.logger.Error("failed to publish to instagram", "error", err)
			pubJob := &models.PublishingJob{
				ContentDraftID: draftID,
				Platform:       "instagram",
				Status:         "failed",
			}
			e.repo.SavePublishingJob(ctx, pubJob)
			return fmt.Errorf("instagram publish: %w", err)
		}

		now := time.Now()
		pubJob := &models.PublishingJob{
			ContentDraftID: draftID,
			Platform:       "instagram",
			Status:         "published",
			PostedAt:       &now,
			ExternalPostID: postID,
		}
		e.repo.SavePublishingJob(ctx, pubJob)
		e.logger.Info("instagram publication success", "post_id", postID)

	case "reject":
		draft.Status = "rejected"
		if _, err := e.repo.SaveContentDraft(ctx, draft); err != nil {
			return err
		}

		doc.Status = "archived"
		if _, err := e.repo.SaveLawDocument(ctx, doc); err != nil {
			return err
		}

	case "regen_img":
		designGuide, err := os.ReadFile("backend/internal/prompts/design_guide.json")
		if err != nil {
			return err
		}
		characterSheet, err := os.ReadFile("backend/internal/prompts/character_sheet.json")
		if err != nil {
			return err
		}

		imagePrompt, err := e.ai.BuildImagePrompt(ctx, draft, string(designGuide), string(characterSheet))
		if err != nil {
			return err
		}

		asset, err := e.imagegen.GenerateImage(ctx, draftID, imagePrompt)
		if err != nil {
			return err
		}

		valid, err := e.validator.Validate(asset.FilePath)
		if err != nil {
			e.logger.Error("validation error", "error", err)
		}
		asset.Validated = valid

		if _, err := e.repo.SaveImageAsset(ctx, asset); err != nil {
			return err
		}

		// Re-send approval request
		if _, err := e.tg.SendApprovalRequest(ctx, draft, asset, analysis, doc.Title); err != nil {
			return err
		}

	case "regen_cap":
		// Re-run content strategy with analysis
		newDraft, err := e.ai.CreateContentStrategy(ctx, analysis)
		if err != nil {
			return err
		}
		// Overwrite the existing draft text
		draft.Caption = newDraft.Caption
		draft.Hook = newDraft.Hook
		draft.Hashtags = newDraft.Hashtags
		draft.Status = "pending_approval"
		if _, err := e.repo.SaveContentDraft(ctx, draft); err != nil {
			return err
		}

		// Fetch existing asset
		assets, err := e.repo.GetImageAssetsByDraft(ctx, draftID)
		if err != nil || len(assets) == 0 {
			return fmt.Errorf("no existing image asset for draft: %w", err)
		}

		// Re-send approval request
		if _, err := e.tg.SendApprovalRequest(ctx, draft, &assets[0], analysis, doc.Title); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown approval action: %s", action)
	}

	return nil
}

// isSuspicious checks if the law has significant conflict, low consistency, or controversy.
func (e *Engine) isSuspicious(analysis *models.LawAnalysis) bool {
	// Check standard patterns for years
	reYear := regexp.MustCompile(`\b(19\d{2}|20\d{2})\b`)
	m := reYear.FindString(analysis.RawJSON) // Check overall JSON payload
	if m == "" {
		m = reYear.FindString(analysis.LawDocumentID)
	}
	
	if m != "" {
		var yr int
		_, _ = fmt.Sscanf(m, "%d", &yr)
		if yr > 0 && yr < 2019 {
			e.logger.Info("law year is before 2019, skipping", "year", yr)
			return false
		}
	}

	// 1. High Controversy
	if analysis.ControversyScore >= 60 {
		return true
	}
	// 2. Low Consistency (high legal conflict/opposite)
	if analysis.LegalConsistency <= 65 {
		return true
	}
	// 3. High Severity on any affected law
	for _, aff := range analysis.AffectedLaws {
		if aff.Severity >= 0.70 {
			return true
		}
	}

	return false
}

// CheckStuckJobs finds documents stuck in a stage too long and re-drives them.
// Spec §5.1 responsibility.
func (e *Engine) CheckStuckJobs(ctx context.Context) error {
	threshold, err := time.ParseDuration(e.cfg.Scheduler.StuckJobThreshold)
	if err != nil {
		threshold = 6 * time.Hour
	}
	cutoff := time.Now().Add(-threshold)

	stuck, err := e.repo.FindStuckDocuments(ctx, "discovered", cutoff)
	if err != nil {
		return fmt.Errorf("query stuck: %w", err)
	}

	for _, doc := range stuck {
		e.logger.Warn("stuck document detected", "id", doc.ID, "law_number", doc.LawNumber, "status", doc.Status)
		// Re-trigger: mark for download retry
		doc.Status = "discovered"
		doc.UpdatedAt = time.Now()
		if _, err := e.repo.SaveLawDocument(ctx, &doc); err != nil {
			e.logger.Error("re-queue stuck doc failed", "id", doc.ID, "error", err)
		}
	}

	return nil
}
