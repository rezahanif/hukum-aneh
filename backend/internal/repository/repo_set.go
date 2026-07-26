package repository

// RepoSet bundles all 8 repository interfaces into a single struct.
//
// Why a wrapper instead of 8 separate fields in Engine?
//   - Engine struct stays compact (1 field instead of 8)
//   - NewEngine signature stays stable when interfaces change
//   - Factory (Phase 5.1) returns *RepoSet, so callers don't care which
//     concrete backend (Firestore / Postgres / DualWrite) is in use
//   - Forward-compatible: adding a 9th interface later only changes RepoSet,
//     not every consumer
//
// Construction:
//   - NewRepoSetFromFirestore(*FirestoreRepo) — current behavior, Phase 0
//   - NewRepoSetFromPostgres(*PostgresRepo) — Phase 3
//   - NewRepoSet(ctx, cfg) — Phase 5.1 factory, reads STORAGE_MODE env
type RepoSet struct {
	LawRepo      LawDocumentRepo
	VersionRepo  LawVersionRepo
	AnalysisRepo LawAnalysisRepo
	DraftRepo    ContentDraftRepo
	ImageRepo    ImageAssetRepo
	ApprovalRepo ApprovalRepo
	PublishRepo  PublishingJobRepo
	EmbedRepo    EmbeddingRepo

	// Closer holds the underlying concrete repo so callers can defer Close().
	// All 8 interfaces above are views into the same concrete struct.
	Closer Closer
}

// NewRepoSetFromFirestore wraps a *FirestoreRepo into a RepoSet.
// All 8 interface fields point to the same concrete repo.
func NewRepoSetFromFirestore(r *FirestoreRepo) *RepoSet {
	return &RepoSet{
		LawRepo:      r,
		VersionRepo:  r,
		AnalysisRepo: r,
		DraftRepo:    r,
		ImageRepo:    r,
		ApprovalRepo: r,
		PublishRepo:  r,
		EmbedRepo:    r,
		Closer:       r,
	}
}

// NewRepoSetFromPostgres will be added in Phase 3.1 (postgres.go).
// NewRepoSet (factory) will be added in Phase 5.1 (factory.go).
