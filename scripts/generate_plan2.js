const fs = require("fs");
const { Document, Packer, Paragraph, TextRun, Header, Footer, AlignmentType, HeadingLevel, PageNumber, Table, TableRow, TableCell, WidthType, ShadingType, BorderStyle, PageBreak, TableLayoutType, TableOfContents } = require("docx");

const COVER = {
  bg: "1A2330", titleColor: "FFFFFF", subtitleColor: "B0B8C0", metaColor: "90989F", footerColor: "687078"
};
const P = { primary: "2C3E50", body: "000000", secondary: "607080", accent: "D4875A", surface: "F8F0EB" };
const T = { headerBg: "D4875A", headerText: "FFFFFF", accentLine: "D4875A", innerLine: "DDD0C8", surface: "F8F0EB" };
const c = (hex) => hex.replace("#", "");
const NB = { style: BorderStyle.NONE, size: 0, color: "FFFFFF" };
const noBorders = { top: NB, bottom: NB, left: NB, right: NB };
const allNoBorders = { top: NB, bottom: NB, left: NB, right: NB, insideHorizontal: NB, insideVertical: NB };
const thinBorder = { style: BorderStyle.SINGLE, size: 4, color: T.innerLine };

function h1(t) { return new Paragraph({ heading: HeadingLevel.HEADING_1, spacing: { before: 360, after: 200 }, children: [new TextRun({ text: t, bold: true, color: c(P.primary), font: { ascii: "Calibri" }, size: 32 })] }); }
function h2(t) { return new Paragraph({ heading: HeadingLevel.HEADING_2, spacing: { before: 280, after: 160 }, children: [new TextRun({ text: t, bold: true, color: c(P.primary), font: { ascii: "Calibri" }, size: 28 })] }); }
function h3(t) { return new Paragraph({ heading: HeadingLevel.HEADING_3, spacing: { before: 200, after: 120 }, children: [new TextRun({ text: t, bold: true, color: c(P.primary), font: { ascii: "Calibri" }, size: 26 })] }); }
function body(t) { return new Paragraph({ spacing: { after: 120, line: 312 }, children: [new TextRun({ text: t, size: 22, color: c(P.body), font: { ascii: "Calibri" } })] }); }
function bodyBold(t) { return new Paragraph({ spacing: { after: 120, line: 312 }, children: [new TextRun({ text: t, size: 22, color: c(P.body), font: { ascii: "Calibri" }, bold: true })] }); }
function code(t) { return new Paragraph({ spacing: { after: 60, line: 276 }, indent: { left: 400 }, children: [new TextRun({ text: t, size: 20, color: "C7254E", font: { ascii: "Consolas" } })] }); }
function bullet(t) { return new Paragraph({ spacing: { after: 60, line: 312 }, indent: { left: 600, hanging: 300 }, children: [new TextRun({ text: "\u2022 " + t, size: 22, color: c(P.body), font: { ascii: "Calibri" } })] }); }

function makeTable(headers, rows) {
  const hdrCells = headers.map(h => new TableCell({
    shading: { type: ShadingType.CLEAR, fill: T.headerBg }, borders: noBorders,
    children: [new Paragraph({ alignment: AlignmentType.CENTER, spacing: { before: 60, after: 60 }, children: [new TextRun({ text: h, bold: true, size: 20, color: T.headerText, font: { ascii: "Calibri" } })] })],
  }));
  const dataRows = rows.map((row, ri) => new TableRow({
    children: row.map(cell => new TableCell({
      shading: ri % 2 === 1 ? { type: ShadingType.CLEAR, fill: T.surface } : undefined,
      borders: { top: thinBorder, bottom: thinBorder, left: NB, right: NB },
      children: [new Paragraph({ spacing: { before: 40, after: 40 }, children: [new TextRun({ text: cell, size: 20, color: c(P.body), font: { ascii: "Calibri" } })] })],
    })),
  }));
  return new Table({ width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED, borders: allNoBorders, rows: [new TableRow({ tableHeader: true, children: hdrCells }), ...dataRows] });
}

// ── Cover (R4 Top Color Block) ──
const coverChildren = [
  new Table({
    width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED, borders: allNoBorders,
    rows: [new TableRow({ height: { value: 16838, rule: "exact" }, children: [
      new TableCell({ shading: { fill: "FFFFFF" }, borders: noBorders, verticalAlign: "top", children: [
        new Table({ width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED, borders: allNoBorders,
          rows: [new TableRow({ height: { value: 8000, rule: "exact" }, children: [
            new TableCell({ shading: { fill: COVER.bg }, borders: noBorders, verticalAlign: "top", margins: { left: 1200, right: 800 }, children: [
              new Paragraph({ spacing: { before: 2500 } }),
              new Paragraph({ spacing: { after: 400 }, children: [new TextRun({ text: "T E C H N I C A L   M I G R A T I O N   P L A N", size: 18, color: "D4875A", font: { ascii: "Calibri" }, characterSpacing: 60 })] }),
              new Paragraph({ spacing: { after: 200, line: 828, lineRule: "atLeast" }, children: [new TextRun({ text: "PostgreSQL + Qdrant Integration Plan", size: 56, bold: true, color: COVER.titleColor, font: { ascii: "Arial" } })] }),
              new Paragraph({ spacing: { after: 100, line: 828, lineRule: "atLeast" }, children: [new TextRun({ text: "Hukum-Aneh Legal AI Pipeline Migration", size: 28, color: COVER.subtitleColor, font: { ascii: "Arial" } })] }),
            ] })
          ] })]
        }),
        new Table({ width: { size: 100, type: WidthType.PERCENTAGE }, borders: allNoBorders,
          rows: [new TableRow({ height: { value: 60, rule: "exact" }, children: [
            new TableCell({ borders: noBorders, shading: { fill: "D4875A" }, children: [new Paragraph({ children: [] })] }),
          ] })]
        }),
        new Paragraph({ spacing: { before: 600 } }),
        new Paragraph({ indent: { left: 1200 }, spacing: { after: 80 }, children: [new TextRun({ text: "Repository: github.com/rezahanif/hukum-aneh", size: 24, color: COVER.metaColor, font: { ascii: "Arial" } })] }),
        new Paragraph({ indent: { left: 1200 }, spacing: { after: 80 }, children: [new TextRun({ text: "Target: PostgreSQL + Qdrant + Firestore (user data only)", size: 24, color: COVER.metaColor, font: { ascii: "Arial" } })] }),
        new Paragraph({ indent: { left: 1200 }, spacing: { after: 80 }, children: [new TextRun({ text: "Date: July 2026", size: 24, color: COVER.metaColor, font: { ascii: "Arial" } })] }),
      ] })
    ] })]
  })
];

const tocChildren = [
  new Paragraph({ spacing: { before: 200, after: 200 }, children: [new TextRun({ text: "Table of Contents", bold: true, size: 32, color: c(P.primary), font: { ascii: "Calibri" } })] }),
  new TableOfContents("TOC", { hyperlink: true, headingStyleRange: "1-3" }),
  new Paragraph({ spacing: { before: 60 }, children: [new TextRun({ text: "(Right-click TOC and select Update Field to refresh page numbers.)", italics: true, size: 18, color: c(P.secondary), font: { ascii: "Calibri" } })] }),
  new Paragraph({ children: [new PageBreak()] }),
];

// ── Body ──
const b = [
h1("1. Executive Summary"),
body("This document provides a corrected, codebase-accurate PostgreSQL integration plan for the Hukum-Aneh Indonesian Legal AI pipeline. The plan migrates legal document storage from Google Firestore to PostgreSQL, adds Qdrant as a dedicated vector store to replace the current brute-force in-memory cosine similarity search, and retains Firestore exclusively for user/application data. Each phase is broken into agent-executable subtasks with exact file paths, code structures, and acceptance criteria derived from a thorough audit of the actual codebase."),
body("The original plan contained 9 significant gaps when evaluated against the real codebase: it missed 7 of 10 data structs in its schema design, did not account for the 13-state pipeline lifecycle including the pause-and-resume pattern at pending_prompt_approval, ignored the local JSON file queue fallback system, overlooked mock fallback strategies for AI/embedding services, listed TikTok as a publish target despite zero implementation existing, and mischaracterized the connector architecture. This corrected plan addresses all of these gaps."),

h1("2. Current Architecture Analysis"),
h2("2.1 Data Model (10 Structs)"),
body("The codebase defines 10 distinct data structs in backend/internal/models/models.go. These map to 6 top-level Firestore collections and 4 subcollections. The structs are: LawDocument (top-level), LawVersion (subcollection: laws/{id}/versions), LawRelationship (subcollection: laws/{id}/relationships), LawAnalysis (subcollection: laws/{id}/analyses), ContentDraft (top-level), Caption (subcollection: content_drafts/{id}/captions), ImageAsset (top-level), Approval (top-level), PublishingJob (top-level), and EmbeddingEntry (top-level). Collection name constants are centralized in backend/internal/models/collections.go."),
body("Critical fields missed by the original plan: LawAnalysis contains 4 scoring dimensions (controversy_score, economic_score, legal_consistency, overall_score) plus confidence and raw_json. ContentDraft has a 6-state status lifecycle (draft, pending_prompt_approval, pending_approval, approved, rejected, published). EmbeddingEntry has an is_mock boolean that must be preserved to distinguish real Gemini embeddings from deterministic fallback vectors. LawVersion.TextContent can exceed 1MB for large Indonesian regulations, requiring PostgreSQL TEXT type with TOAST storage."),

h2("2.2 Repository Layer (Monolithic)"),
body("The entire data access layer lives in a single 378-line file: backend/internal/repository/firestore.go. The FirestoreRepo struct wraps a single *firestore.Client and implements all CRUD for all 10 models as methods on one concrete struct. There are zero interfaces defined anywhere. Every consumer (Engine, retrieval.Service, and 6 cmd/ entrypoints) imports and depends directly on *repository.FirestoreRepo. This makes it impossible to swap storage backends without modifying every consumer. Interface extraction is the critical prerequisite before any PostgreSQL work."),
body("Notably, GetLawAnalysisByDraft() uses a CollectionGroup query to scan all analyses across all law subcollections, then filters in-memory by document ID. This O(n) approach will be replaced with a direct JOIN in PostgreSQL. The SaveLawDocument method uses the pattern: if ID is empty, call Collection.Add() (auto-ID); if ID is non-empty, call Doc(id).Set() (upsert). This pattern must be replicated with INSERT ... ON CONFLICT in PostgreSQL."),

h2("2.3 Pipeline Lifecycle (13 States)"),
body("The Engine in backend/internal/workflow/engine.go implements a 13-state document lifecycle: discovered -> downloaded -> parsed -> analyzed -> no_conflict (terminal) | pending_prompt_approval (pipeline pauses, awaits Telegram callback) -> pending_approval (image generated, awaits human review) -> approved -> published. Error terminals: download_failed, parse_failed. Rejection paths lead to archived. ContentDraft has its own parallel 6-state lifecycle updated in lockstep."),
body("The pause-and-resume pattern is critical: after generating an image_prompt, the pipeline saves the draft with status pending_prompt_approval and sends a Telegram inline keyboard. The pipeline exits. When the user taps a button, cmd/bot/main.go calls Engine.HandleApprovalAction(draftID, action, reviewerID), which re-fetches all data and resumes. PostgreSQL must preserve this with proper transaction isolation."),

h2("2.4 Vector Search (Brute-Force)"),
body("The retrieval.Service in backend/internal/retrieval/retrieval.go loads ALL embeddings from Firestore into memory via repo.ListAllEmbeddings(), then computes cosine similarity in a nested O(n) loop with bubble sort. This will not scale past ~5,000 records. Embedding dimension is 1536 (Gemini gemini-embedding-2). Mock fallback vectors are generated on API auth/quota errors, flagged with is_mock=true. Migration to Qdrant must exclude mock embeddings from search."),

h2("2.5 Connector Architecture (15 Connectors)"),
body("The codebase has 15 registered connectors in backend/internal/connectors/. The peraturan connector alone handles 6 document types (UU, PP, Perpres, Permen, Kepmen, PBI). TikTok has zero implementation despite being mentioned as a publish target. Instagram publishing exists but is broken - it requires a public image URL and has no image hosting (S3/GCS/Imgur)."),

h2("2.6 Additional Systems"),
body("Local JSON queue: cmd/flush_local/main.go implements a fallback at backend/internal/storage/local_queue/ writing law documents as JSON when Firestore is unavailable. Mock fallbacks: AI service (ai.go) and retrieval service return hardcoded mock responses on 401/429 errors, saved to DB. System dependencies: poppler-utils (pdftotext/pdftoppm), tesseract (ind+eng), Python 3 + curl_cffi + beautifulsoup4."),

h1("3. Target Architecture"),
h2("3.1 Storage Responsibilities"),
makeTable(["Storage", "Responsibility", "Data Types"], [
  ["PostgreSQL", "All legal pipeline data", "10 tables: laws, versions, relationships, analyses, drafts, captions, images, approvals, pub_jobs, embed_meta"],
  ["Qdrant", "Vector similarity search", "1536-dim float32 vectors with metadata (law_doc_id, is_mock filter)"],
  ["Firestore (retained)", "User/app data only", "Telegram sessions, bot state"],
  ["Google Drive / S3", "Raw PDF files", "Downloaded regulation PDFs"],
]),
h2("3.2 Phase Dependency Order"),
body("Phase 0 (Interface Abstraction) is the prerequisite for everything. Phase 1 (PG Schema) and Phase 2 (Qdrant Setup) run in parallel after Phase 0. Phase 3 (PG Repo) depends on Phase 1. Phase 4 (Qdrant Retrieval) depends on Phase 2. Phase 5 (Dual-Write Wiring) depends on Phases 3+4. Phase 6 (Data Migration) depends on Phase 5. Phase 7 (Cleanup) depends on Phase 6."),

h1("4. Phase 0: Repository Interface Abstraction (Prerequisite)"),
body("This phase extracts interfaces from the monolithic FirestoreRepo so that consumers depend on abstractions, not concretions. This unblocks all subsequent phases."),

h2("4.1 Task 0.1: Define Repository Interfaces"),
bodyBold("File: backend/internal/repository/interfaces.go (NEW)"),
body("Create 8 granular interfaces covering one domain aggregate each. Each interface must exactly match the method signatures already used by existing consumers."),
body("LawDocumentRepo: SaveLawDocument, GetLawDocument, FindByLawNumber, ListLawsByStatus, ListAllLaws, FindStuckDocuments. LawVersionRepo: GetLatestLawVersion, SaveLawVersion. LawAnalysisRepo: SaveLawAnalysis, GetLawAnalysisByDraft. ContentDraftRepo: GetContentDraft, SaveContentDraft. ImageAssetRepo: GetImageAssetsByDraft, SaveImageAsset. ApprovalRepo: SaveApproval. PublishingJobRepo: SavePublishingJob. EmbeddingRepo: SaveEmbedding, ListAllEmbeddings."),
body("Acceptance: (1) go build ./... passes (2) Signatures exactly match FirestoreRepo methods (3) Each interface is a separate named type."),

h2("4.2 Task 0.2: Compile-Time Interface Assertions"),
bodyBold("File: backend/internal/repository/firestore.go (MODIFY)"),
body("Add 8 compile-time assertions: var _ LawDocumentRepo = (*FirestoreRepo)(nil) etc. If any method is missing or mismatched, the compiler fails immediately at the assertion line."),

h2("4.3 Task 0.3: Refactor Engine Struct"),
bodyBold("File: backend/internal/workflow/engine.go (MODIFY)"),
body("Replace the single repo *repository.FirestoreRepo field with 8 individual interface fields: lawRepo, versionRepo, analysisRepo, draftRepo, imageRepo, approvalRepo, publishRepo, embedRepo. Update NewEngine() to accept these individually (or as a wrapper struct). Update all ~30 call sites: e.repo.SaveLawDocument -> e.lawRepo.SaveLawDocument, e.repo.GetLawAnalysisByDraft -> e.analysisRepo.GetLawAnalysisByDraft, etc. See the full mapping in the original analysis."),

h2("4.4 Task 0.4: Update retrieval.Service"),
bodyBold("File: backend/internal/retrieval/retrieval.go (MODIFY)"),
body("Change repo field from *repository.FirestoreRepo to repository.EmbeddingRepo (only needs ListAllEmbeddings and SaveEmbedding)."),

h2("4.5 Task 0.5: Update All cmd/ Entry Points"),
bodyBold("Files: 6 cmd/*/main.go files (MODIFY)"),
body("Each file creates FirestoreRepo and passes it through interface-typed parameters. Files: cmd/pipeline/main.go, cmd/bot/main.go, cmd/backfill/main.go, cmd/batch/main.go, cmd/backfill_embeddings/main.go, cmd/flush_local/main.go. No runtime behavior changes."),

h1("5. Phase 1: PostgreSQL Schema Design"),

h2("5.1 Task 1.1: Add Dependencies"),
bodyBold("File: go.mod (MODIFY)"),
body("Run: go get github.com/jackc/pgx/v5, github.com/jackc/pgx/v5/pgxpool, github.com/golang-migrate/migrate/v4, github.com/golang-migrate/migrate/v4/database/pgx/v5, github.com/golang-migrate/migrate/v4/source/file, github.com/qdrant/go-client/qdrant. Then go mod tidy."),

h2("5.2 Task 1.2: Create Migration Files"),
bodyBold("Directory: backend/migrations/ (NEW)"),
body("Create 000001_init_schema.up.sql with all 10 tables. Key design decisions per table:"),

h3("law_documents"),
body("PK: id (TEXT, matching Firestore string IDs). Status CHECK constraint for all 13 valid states. UNIQUE index on law_number. Composite index on (status, updated_at) for FindStuckDocuments and ListLawsByStatus."),

h3("law_versions"),
body("FK to law_documents(id). text_content uses TEXT (TOAST handles 1MB+ regulation texts). embedding field is NOT stored here (moves to Qdrant). Only embedding_id (TEXT ref to Qdrant point ID) kept. Index on (law_document_id, version_number DESC)."),

h3("law_analyses"),
body("FK to law_documents(id). affected_laws stored as JSONB. Four score columns as INTEGER CHECK 0-100. raw_json as TEXT. confidence as FLOAT8."),

h3("content_drafts"),
body("FK to law_analyses(id). hashtags as JSONB. Status CHECK for 6 valid states including pending_prompt_approval and prompt_rejected."),

h3("captions, image_assets, approvals, publishing_jobs"),
body("Straightforward FK mappings. Approvals stores reviewer_id (Telegram user ID) and decision with CHECK (approve, reject, regenerate_caption, regenerate_image, prompt_approve, prompt_reject, prompt_regen). PublishingJobs platform CHECK only allows 'instagram' - do NOT add tiktok."),

h3("embedding_metadata"),
body("Replaces the embeddings collection. Vector NOT stored here - only metadata (law_document_id, is_mock, created_at, qdrant_point_id). The is_mock boolean is critical for filtering. Index on is_mock and law_document_id."),

h2("5.3 Task 1.3: Add PostgreSQL + Qdrant Config"),
bodyBold("File: backend/internal/config/config.go (MODIFY)"),
body("Add Postgres struct (DSN, MaxOpenConns, MaxIdleConns) and Qdrant struct (URL, Collection, APIKey) to Config. Load from env: POSTGRES_DSN, QDRANT_URL, QDRANT_COLLECTION, QDRANT_API_KEY. Defaults: localhost:5432, localhost:6333, law_embeddings."),

h1("6. Phase 2: Qdrant Vector Store Setup"),

h2("6.1 Task 2.1: Create Qdrant Client Wrapper"),
bodyBold("File: backend/internal/retrieval/qdrant.go (NEW)"),
body("Create QdrantClient with: (1) EnsureCollection - creates collection with 1536 dims, cosine metric, HNSW index. (2) Upsert - stores vectors with payload {law_document_id, is_mock}. (3) Search - returns topN results, filtering out is_mock=true using a Qdrant must_not filter. Use qdrant-go-client library."),

h2("6.2 Task 2.2: Docker Compose"),
bodyBold("File: docker-compose.yml (NEW at project root)"),
body("Define postgres:16-alpine and qdrant:latest services with volumes. PostgreSQL on 5432, Qdrant on 6333/6334."),

h1("7. Phase 3: PostgreSQL Repository Implementation"),

h2("7.1 Task 3.1: Create PostgresRepo"),
bodyBold("File: backend/internal/repository/postgres.go (NEW)"),
body("Create PostgresRepo wrapping *pgxpool.Pool implementing all 8 interfaces. Use INSERT ... ON CONFLICT (id) DO UPDATE for upsert pattern. Auto-generate ID (UUID v4 or nanoid) when empty. For GetLawAnalysisByDraft, use JOIN: SELECT a.* FROM law_analyses a JOIN content_drafts d ON a.id = d.law_analysis_id WHERE d.id = $1."),

h2("7.2 Task 3.2-3.3: Implement All Methods"),
body("Implement ~16 methods across all 8 interfaces. Use parameterized SQL ($1, $2). JSONB fields use pgx native JSON support with json.Marshal/Unmarshal at boundaries. EmbeddingRepo.SaveEmbedding inserts to embedding_metadata only (vector goes to Qdrant). EmbeddingRepo.ListAllEmbeddings selects from embedding_metadata (no vector data)."),

h1("8. Phase 4: Qdrant Retrieval Integration"),

h2("8.1 Task 4.1: Update retrieval.Service"),
bodyBold("File: backend/internal/retrieval/retrieval.go (MODIFY)"),
body("Add QdrantClient to Service struct. Rewrite Search() to use Qdrant when available, fallback to brute-force (renamed bruteForceSearch) when nil. The GenerateEmbedding method (Gemini API + mock fallback) remains unchanged."),

h2("8.2 Task 4.2: Upsert on Embed"),
bodyBold("File: backend/internal/workflow/engine.go (MODIFY)"),
body("In ProcessParsedDocument, after saving embedding metadata, upsert non-mock vectors to Qdrant. If Qdrant is unavailable, log warning and continue (non-blocking). Mock embeddings (is_mock=true) are never upserted."),

h1("9. Phase 5: Dual-Write Configuration and Wiring"),

h2("9.1 Task 5.1: Repository Factory"),
bodyBold("File: backend/internal/repository/factory.go (NEW)"),
body("Create NewRepoSet(ctx, cfg) returning a RepoSet struct with all 8 interface fields. Support modes: 'firestore' (original), 'postgres' (new), 'dual' (writes both, reads from PG). Dual mode uses a DualWriteRepo wrapper for validation during migration."),

h2("9.2 Task 5.2: Wire All cmd/ Files"),
bodyBold("Files: All 6 cmd/*/main.go (MODIFY)"),
body("Replace direct FirestoreRepo instantiation with NewRepoSet factory call. STORAGE_MODE env var selects backend. 'firestore' mode is 100% backward compatible."),

h1("10. Phase 6: Data Migration"),

h2("10.1 Task 6.1: Migration Tool"),
bodyBold("File: backend/cmd/migrate_to_pg/main.go (NEW)"),
body("One-time tool: read all Firestore data, write to PostgreSQL + Qdrant. Handle: (1) Subcollection flattening (laws/{id}/versions -> top-level with FK). (2) Preserve all Firestore string IDs as PG primary keys. (3) Only upsert non-mock embeddings to Qdrant. (4) Large text_content handled by TOAST. (5) JSONB conversion for affected_laws and hashtags. (6) Idempotent: ON CONFLICT DO NOTHING."),

h2("10.2 Task 6.2: Update Local Queue Flush"),
bodyBold("File: backend/cmd/flush_local/main.go (VERIFY)"),
body("Should already work via factory. Verify it writes to PostgreSQL when STORAGE_MODE=postgres."),

h1("11. Phase 7: Cleanup, Dockerfile, Testing"),

h2("11.1 Task 7.1: Dockerfile"),
bodyBold("File: Dockerfile (NEW at project root)"),
body("Multi-stage: Go 1.25 builder -> debian:bookworm-slim runtime with poppler-utils, tesseract-ocr (ind+eng), python3 + pip (curl_cffi, beautifulsoup4). Separate Dockerfile or compose service for cmd/bot."),

h2("11.2 Task 7.2: Remove Firestore from Pipeline"),
body("After validation: remove cloud.google.com/go/firestore from go.mod, remove Firebase config from config.go, update docker-compose. Keep Firebase only if user/bot state requires it."),

h2("11.3 Task 7.3: Integration Testing"),
bodyBold("File: backend/internal/repository/postgres_test.go (NEW)"),
body("Cover: (1) Full pipeline lifecycle through PG. (2) Pause-and-resume via pending_prompt_approval. (3) Qdrant search with mock exclusion. (4) Concurrent access with -race flag. (5) All 8 interface methods."),

h1("12. Known Issues and Blockers"),
h2("12.1 Instagram Publishing Broken"),
body("HandleApprovalAction (line ~438) uses cfg.Instagram.AccessToken as placeholder for public image URL. Publishing cannot work without S3/GCS/Imgur for image hosting. Job saved with status pending_image_hosting. This must be addressed separately."),
h2("12.2 TikTok: Zero Implementation"),
body("Do NOT add TikTok. PublishingJob platform CHECK allows only 'instagram'. TikTok support is a separate feature effort."),
h2("12.3 Connector Architecture"),
body("peraturan.go handles 6 doc types (UU, PP, Perpres, Permen, Kepmen, PBI). Total: 15 registered connectors, not 20+."),

h1("13. Task Dependency Table"),
makeTable(["Task ID", "Task", "Depends On", "Phase"], [
  ["0.1", "Define repository interfaces", "None", "0"],
  ["0.2", "FirestoreRepo implements interfaces", "0.1", "0"],
  ["0.3", "Refactor Engine to interfaces", "0.1", "0"],
  ["0.4", "Update retrieval.Service", "0.1", "0"],
  ["0.5", "Update all cmd/ files", "0.2, 0.3, 0.4", "0"],
  ["1.1", "Add PostgreSQL + Qdrant deps", "None (parallel Phase 0)", "1"],
  ["1.2", "Create migration files (10 tables)", "1.1", "1"],
  ["1.3", "Add PostgreSQL + Qdrant config", "None", "1"],
  ["2.1", "Create Qdrant client wrapper", "1.1", "2"],
  ["2.2", "Docker Compose for PG + Qdrant", "None", "2"],
  ["3.1", "Create PostgresRepo struct", "1.2", "3"],
  ["3.2", "Implement LawDocument methods (6)", "3.1", "3"],
  ["3.3", "Implement remaining methods (~10)", "3.2", "3"],
  ["4.1", "Update retrieval.Service for Qdrant", "2.1, 3.1", "4"],
  ["4.2", "Upsert to Qdrant on embed", "2.1", "4"],
  ["5.1", "Create repository factory", "3.3, 4.1", "5"],
  ["5.2", "Wire cmd/ files for factory", "5.1", "5"],
  ["6.1", "Create FS-to-PG migration tool", "5.2", "6"],
  ["6.2", "Update local queue flush", "5.2", "6"],
  ["7.1", "Create Dockerfile", "2.2, 5.2", "7"],
  ["7.2", "Remove Firestore dependency", "6.1", "7"],
  ["7.3", "Integration testing", "5.2, 2.2", "7"],
]),

]; // end body

const doc = new Document({
  styles: { default: { document: {
    run: { font: { ascii: "Calibri" }, size: 22, color: c(P.body) },
    paragraph: { spacing: { line: 312 } },
  }}},
  sections: [
    { properties: { page: { margin: { top: 0, bottom: 0, left: 0, right: 0 }, size: { width: 11906, height: 16838 } } }, children: coverChildren },
    { properties: { page: { margin: { top: 1440, bottom: 1440, left: 1701, right: 1417 }, pageNumbers: { start: 1, formatType: "upperRoman" } } },
      footers: { default: new Footer({ children: [new Paragraph({ alignment: AlignmentType.CENTER, children: [new TextRun({ children: ["PAGE \\* ROMAN \\* MERGEFORMAT"], size: 18, color: c(P.secondary) })] })] }) },
      children: tocChildren },
    { properties: { page: { margin: { top: 1440, bottom: 1440, left: 1701, right: 1417 }, pageNumbers: { start: 1, formatType: "decimal" } } },
      headers: { default: new Header({ children: [new Paragraph({ alignment: AlignmentType.RIGHT, children: [new TextRun({ text: "PostgreSQL + Qdrant Integration Plan", size: 16, color: c(P.secondary), font: { ascii: "Calibri" }, italics: true })] })] }) },
      footers: { default: new Footer({ children: [new Paragraph({ alignment: AlignmentType.CENTER, children: [new TextRun({ children: ["PAGE \\* arabic \\* MERGEFORMAT"], size: 18, color: c(P.secondary) })] })] }) },
      children: b },
  ],
});

const outPath = "/home/z/my-project/download/Hukum-Aneh_PostgreSQL_Qdrant_Integration_Plan.docx";
Packer.toBuffer(doc).then(buf => { fs.writeFileSync(outPath, buf); console.log("Written: " + outPath); });
