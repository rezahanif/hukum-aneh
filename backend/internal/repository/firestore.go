package repository

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"github.com/rezahanif/hukum-aneh/backend/internal/models"
)

// FirestoreRepo wraps Firestore client for all model CRUD.
// Each method touches only the collections it owns.
type FirestoreRepo struct {
	client *firestore.Client
}

func NewFirestoreRepo(ctx context.Context, projectID string, credsPath string) (*FirestoreRepo, error) {
	var opts []option.ClientOption
	if credsPath != "" {
		opts = append(opts, option.WithCredentialsFile(credsPath))
	}
	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("firestore client: %w", err)
	}
	return &FirestoreRepo{client: client}, nil
}

func (r *FirestoreRepo) Close() error {
	return r.client.Close()
}

// --- LawDocument ---

func (r *FirestoreRepo) SaveLawDocument(ctx context.Context, doc *models.LawDocument) (string, error) {
	if doc.ID == "" {
		ref, _, err := r.client.Collection(models.ColLaws).Add(ctx, doc)
		if err != nil {
			return "", fmt.Errorf("add law doc: %w", err)
		}
		doc.ID = ref.ID
		return ref.ID, nil
	}
	_, err := r.client.Collection(models.ColLaws).Doc(doc.ID).Set(ctx, doc)
	if err != nil {
		return "", fmt.Errorf("set law doc: %w", err)
	}
	return doc.ID, nil
}

func (r *FirestoreRepo) GetLawDocument(ctx context.Context, id string) (*models.LawDocument, error) {
	ds, err := r.client.Collection(models.ColLaws).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("get law doc: %w", err)
	}
	var doc models.LawDocument
	if err := ds.DataTo(&doc); err != nil {
		return nil, fmt.Errorf("decode law doc: %w", err)
	}
	doc.ID = ds.Ref.ID
	return &doc, nil
}

// FindByLawNumber checks if a law already exists by law_number.
// Returns nil if not found.
func (r *FirestoreRepo) FindByLawNumber(ctx context.Context, lawNumber string) (*models.LawDocument, error) {
	q := r.client.Collection(models.ColLaws).Where("law_number", "==", lawNumber).Limit(1)
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("query law_number: %w", err)
	}
	if len(docs) == 0 {
		return nil, nil
	}
	var doc models.LawDocument
	if err := docs[0].DataTo(&doc); err != nil {
		return nil, fmt.Errorf("decode law doc: %w", err)
	}
	doc.ID = docs[0].Ref.ID
	return &doc, nil
}

func (r *FirestoreRepo) ListLawsByStatus(ctx context.Context, status string) ([]models.LawDocument, error) {
	q := r.client.Collection(models.ColLaws).Where("status", "==", status)
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("query laws by status: %w", err)
	}
	var result []models.LawDocument
	for _, d := range docs {
		var doc models.LawDocument
		if err := d.DataTo(&doc); err != nil {
			continue
		}
		doc.ID = d.Ref.ID
		result = append(result, doc)
	}
	return result, nil
}

// --- LawVersion ---

func (r *FirestoreRepo) SaveLawVersion(ctx context.Context, lawID string, v *models.LawVersion) (string, error) {
	if v.VersionNumber == 0 {
		v.VersionNumber = int(time.Now().Unix())
	}
	if v.ID == "" {
		ref, _, err := r.client.Collection(models.ColLaws).Doc(lawID).
			Collection(models.SubVersions).Add(ctx, v)
		if err != nil {
			return "", fmt.Errorf("add version: %w", err)
		}
		v.ID = ref.ID
		return ref.ID, nil
	}
	_, err := r.client.Collection(models.ColLaws).Doc(lawID).
		Collection(models.SubVersions).Doc(v.ID).Set(ctx, v)
	if err != nil {
		return "", fmt.Errorf("set version: %w", err)
	}
	return v.ID, nil
}

// --- LawAnalysis ---

func (r *FirestoreRepo) SaveLawAnalysis(ctx context.Context, lawID string, a *models.LawAnalysis) (string, error) {
	a.CreatedAt = time.Now()
	if a.ID == "" {
		ref, _, err := r.client.Collection(models.ColLaws).Doc(lawID).
			Collection(models.SubAnalyses).Add(ctx, a)
		if err != nil {
			return "", fmt.Errorf("add analysis: %w", err)
		}
		a.ID = ref.ID
		return ref.ID, nil
	}
	_, err := r.client.Collection(models.ColLaws).Doc(lawID).
		Collection(models.SubAnalyses).Doc(a.ID).Set(ctx, a)
	return a.ID, fmt.Errorf("set analysis: %w", err)
}

// --- ContentDraft ---

func (r *FirestoreRepo) SaveContentDraft(ctx context.Context, draft *models.ContentDraft) (string, error) {
	draft.CreatedAt = time.Now()
	if draft.ID == "" {
		ref, _, err := r.client.Collection(models.ColContentDrafts).Add(ctx, draft)
		if err != nil {
			return "", fmt.Errorf("add draft: %w", err)
		}
		draft.ID = ref.ID
		return ref.ID, nil
	}
	_, err := r.client.Collection(models.ColContentDrafts).Doc(draft.ID).Set(ctx, draft)
	return draft.ID, fmt.Errorf("set draft: %w", err)
}

// --- ImageAsset ---

func (r *FirestoreRepo) SaveImageAsset(ctx context.Context, asset *models.ImageAsset) (string, error) {
	asset.CreatedAt = time.Now()
	if asset.ID == "" {
		ref, _, err := r.client.Collection(models.ColImageAssets).Add(ctx, asset)
		if err != nil {
			return "", fmt.Errorf("add image asset: %w", err)
		}
		asset.ID = ref.ID
		return ref.ID, nil
	}
	_, err := r.client.Collection(models.ColImageAssets).Doc(asset.ID).Set(ctx, asset)
	return asset.ID, fmt.Errorf("set image asset: %w", err)
}

// --- Approval ---

func (r *FirestoreRepo) SaveApproval(ctx context.Context, a *models.Approval) (string, error) {
	a.Timestamp = time.Now()
	if a.ID == "" {
		ref, _, err := r.client.Collection(models.ColApprovals).Add(ctx, a)
		if err != nil {
			return "", fmt.Errorf("add approval: %w", err)
		}
		a.ID = ref.ID
		return ref.ID, nil
	}
	_, err := r.client.Collection(models.ColApprovals).Doc(a.ID).Set(ctx, a)
	return a.ID, fmt.Errorf("set approval: %w", err)
}

// --- PublishingJob ---

func (r *FirestoreRepo) SavePublishingJob(ctx context.Context, j *models.PublishingJob) (string, error) {
	if j.ID == "" {
		ref, _, err := r.client.Collection(models.ColPublishingJobs).Add(ctx, j)
		if err != nil {
			return "", fmt.Errorf("add pub job: %w", err)
		}
		j.ID = ref.ID
		return ref.ID, nil
	}
	_, err := r.client.Collection(models.ColPublishingJobs).Doc(j.ID).Set(ctx, j)
	return j.ID, fmt.Errorf("set pub job: %w", err)
}

// --- Stuck job detection ---

// FindStuckDocuments returns laws whose status hasn't advanced past the given
// stage within the threshold. Used by Scheduler §5.1 for re-driving stalled work.
func (r *FirestoreRepo) FindStuckDocuments(ctx context.Context, status string, before time.Time) ([]models.LawDocument, error) {
	q := r.client.Collection(models.ColLaws).
		Where("status", "==", status).
		Where("updated_at", "<", before)
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("query stuck docs: %w", err)
	}
	var result []models.LawDocument
	for _, d := range docs {
		var doc models.LawDocument
		if err := d.DataTo(&doc); err != nil {
			continue
		}
		doc.ID = d.Ref.ID
		result = append(result, doc)
	}
	return result, nil
}
