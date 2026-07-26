const fs = require("fs");
const { Document, Packer, Paragraph, TextRun, Header, Footer, AlignmentType, HeadingLevel, PageNumber, Table, TableRow, TableCell, WidthType, ShadingType, BorderStyle, PageBreak, TableLayoutType, TableOfContents } = require("docx");

// ── GO-1 Graphite Orange palette (plan/proposal) ──
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

function calcTitleLayout(title, maxWidthTwips, preferredPt, minPt) {
  preferredPt = preferredPt || 40; minPt = minPt || 24;
  const charWidth = (pt) => pt * 20;
  const charsPerLine = (pt) => Math.floor(maxWidthTwips / charWidth(pt));
  let titlePt = preferredPt, lines;
  while (titlePt >= minPt) {
    const cpl = charsPerLine(titlePt);
    if (cpl < 2) { titlePt -= 2; continue; }
    lines = splitTitleLines(title, cpl);
    if (lines.length <= 3) break;
    titlePt -= 2;
  }
  if (!lines || lines.length > 3) {
    lines = splitTitleLines(title, charsPerLine(minPt));
    titlePt = minPt;
  }
  return { titlePt, titleLines: lines };
}

function splitTitleLines(title, charsPerLine) {
  if (title.length <= charsPerLine) return [title];
  const breakAfter = new Set([' ', '-', '/', ':', '(', ')', '.']);
  const lines = []; let remaining = title;
  while (remaining.length > charsPerLine) {
    let breakAt = -1;
    for (let i = charsPerLine; i >= Math.floor(charsPerLine * 0.6); i--) {
      if (i < remaining.length && breakAfter.has(remaining[i - 1])) { breakAt = i; break; }
    }
    if (breakAt === -1) breakAt = charsPerLine;
    lines.push(remaining.slice(0, breakAt));
    remaining = remaining.slice(breakAt);
  }
  if (remaining.length > 0) lines.push(remaining);
  if (lines.length > 0 && lines[lines.length - 1].length <= 2 && lines.length > 1) {
    lines[lines.length - 2] += " " + lines.pop();
  }
  return lines;
}

function buildCoverR4(config) {
  const P_ = config.palette;
  const padL = 1200, padR = 800;
  const availableWidth = 11906 - padL - padR;
  const { titlePt, titleLines } = calcTitleLayout(config.title, availableWidth, 36, 24);
  const titleSize = titlePt * 2;
  const titleBlockHeight = titleLines.length * (titlePt * 23 + 200);
  const englishLabelH = config.englishLabel ? (9 * 23 + 500) : 0;
  const subtitleH = config.subtitle ? (12 * 23 + 200) : 0;
  const upperContentH = englishLabelH + titleBlockHeight + subtitleH;
  const UPPER_MIN = 7500;
  const UPPER_H = Math.max(UPPER_MIN, upperContentH + 1500 + 800);
  const DIVIDER_H = 60;
  const contentEstimate = englishLabelH + titleLines.length * (titlePt * 23 + 200) + subtitleH;
  const spacerIntrinsic = 280;
  const topSpacing = Math.max(UPPER_H - contentEstimate - spacerIntrinsic - 800, 400);

  const upperBlock = new Table({
    width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED, borders: allNoBorders,
    rows: [new TableRow({ height: { value: UPPER_H, rule: "exact" }, children: [
      new TableCell({ shading: { fill: P_.bg }, borders: noBorders, verticalAlign: "top",
        margins: { left: padL, right: padR },
        children: [
          new Paragraph({ spacing: { before: topSpacing } }),
          config.englishLabel ? new Paragraph({ spacing: { after: 500 },
            children: [new TextRun({ text: config.englishLabel.split("").join(" "), size: 18, color: "D4875A", font: { ascii: "Calibri" }, characterSpacing: 60 })],
          }) : null,
          ...titleLines.map((line, i) => new Paragraph({
            spacing: { after: i < titleLines.length - 1 ? 100 : 200, line: Math.ceil(titlePt * 23), lineRule: "atLeast" },
            children: [new TextRun({ text: line, size: titleSize, bold: true, color: P_.titleColor, font: { ascii: "Arial" } })],
          })),
          config.subtitle ? new Paragraph({ spacing: { after: 100 },
            children: [new TextRun({ text: config.subtitle, size: 24, color: P_.subtitleColor, font: { ascii: "Arial" } })],
          }) : null,
        ].filter(Boolean),
      })
    ]})]
  });

  const divider = new Table({
    width: { size: 100, type: WidthType.PERCENTAGE }, borders: allNoBorders,
    rows: [new TableRow({ height: { value: DIVIDER_H, rule: "exact" }, children: [
      new TableCell({ borders: noBorders, shading: { fill: "D4875A" }, children: [new Paragraph({ children: [] })] }),
    ]})]
  });

  const lowerContent = [
    new Paragraph({ spacing: { before: 800 } }),
    ...(config.metaLines || []).map(line => new Paragraph({
      indent: { left: padL }, spacing: { after: 100 },
      children: [new TextRun({ text: line, size: 28, color: P_.metaColor, font: { ascii: "Arial" } })],
    })),
    new Paragraph({ spacing: { before: 2000 } }),
    new Paragraph({ indent: { left: padL },
      children: [
        new TextRun({ text: config.footerLeft || "", size: 22, color: "909090" }),
        new TextRun({ text: "          " }),
        new TextRun({ text: config.footerRight || "", size: 22, color: "909090" }),
      ],
    }),
  ];

  return [new Table({
    width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED, borders: allNoBorders,
    rows: [new TableRow({ height: { value: 16838, rule: "exact" }, children: [
      new TableCell({ shading: { fill: "FFFFFF" }, borders: noBorders, verticalAlign: "top",
        children: [upperBlock, divider, ...lowerContent],
      })
    ]})]
  })];
}

// ── Content helpers ──
function h1(text) {
  return new Paragraph({ heading: HeadingLevel.HEADING_1, spacing: { before: 360, after: 200 },
    children: [new TextRun({ text, bold: true, color: c(P.primary), font: { ascii: "Calibri" }, size: 32 })] });
}
function h2(text) {
  return new Paragraph({ heading: HeadingLevel.HEADING_2, spacing: { before: 280, after: 160 },
    children: [new TextRun({ text, bold: true, color: c(P.primary), font: { ascii: "Calibri" }, size: 28 })] });
}
function h3(text) {
  return new Paragraph({ heading: HeadingLevel.HEADING_3, spacing: { before: 200, after: 120 },
    children: [new TextRun({ text, bold: true, color: c(P.primary), font: { ascii: "Calibri" }, size: 26 })] });
}
function body(text) {
  return new Paragraph({ spacing: { after: 120, line: 312 },
    children: [new TextRun({ text, size: 22, color: c(P.body), font: { ascii: "Calibri" } })] });
}
function bodyBold(text) {
  return new Paragraph({ spacing: { after: 120, line: 312 },
    children: [new TextRun({ text, size: 22, color: c(P.body), font: { ascii: "Calibri" }, bold: true })] });
}
function code(text) {
  return new Paragraph({ spacing: { after: 80, line: 276 }, indent: { left: 400 },
    children: [new TextRun({ text, size: 20, color: "C7254E", font: { ascii: "Consolas" } })] });
}
function bullet(text) {
  return new Paragraph({ spacing: { after: 60, line: 312 }, indent: { left: 600, hanging: 300 },
    children: [new TextRun({ text: "\u2022 " + text, size: 22, color: c(P.body), font: { ascii: "Calibri" } })] });
}
function subBullet(text) {
  return new Paragraph({ spacing: { after: 40, line: 312 }, indent: { left: 1000, hanging: 300 },
    children: [new TextRun({ text: "- " + text, size: 21, color: c(P.secondary), font: { ascii: "Calibri" } })] });
}

function makeTable(headers, rows) {
  const hdrCells = headers.map(h => new TableCell({
    shading: { type: ShadingType.CLEAR, fill: T.headerBg }, borders: { top: NB, bottom: NB, left: NB, right: NB },
    children: [new Paragraph({ alignment: AlignmentType.CENTER, spacing: { before: 60, after: 60 },
      children: [new TextRun({ text: h, bold: true, size: 20, color: T.headerText, font: { ascii: "Calibri" } })] })],
  }));
  const dataRows = rows.map((row, ri) => new TableRow({
    children: row.map((cell, ci) => new TableCell({
      shading: ri % 2 === 1 ? { type: ShadingType.CLEAR, fill: T.surface } : undefined,
      borders: { top: thinBorder, bottom: thinBorder, left: NB, right: NB },
      children: [new Paragraph({ spacing: { before: 40, after: 40 },
        children: [new TextRun({ text: cell, size: 20, color: c(P.body), font: { ascii: "Calibri" } })] })],
    })),
  }));
  return new Table({ width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED, borders: allNoBorders,
    rows: [new TableRow({ tableHeader: true, children: hdrCells }), ...dataRows] });
}

// ── DOCUMENT CONTENT ──
const coverChildren = buildCoverR4({
  palette: COVER, title: "PostgreSQL + Qdrant Integration Plan", subtitle: "Hukum-Aneh Legal AI Pipeline Migration",
  englishLabel: "TECHNICAL MIGRATION PLAN",
  metaLines: ["Repository: github.com/rezahanif/hukum-aneh", "Target: PostgreSQL (legal data) + Firestore (user data) + Qdrant (vectors)", "Date: July 2026"],
  footerLeft: "Hukum-Aneh", footerRight: "v1.0"
});

const tocChildren = [
  new Paragraph({ spacing: { before: 200, after: 200 }, children: [new TextRun({ text: "Table of Contents", bold: true, size: 32, color: c(P.primary), font: { ascii: "Calibri" } })] }),
  new TableOfContents("Table of Contents", { hyperlink: true, headingStyleRange: "1-3" }),
  new Paragraph({ spacing: { before: 100 }, children: [new TextRun({ text: "(Right-click the TOC and select \"Update Field\" to refresh page numbers after opening in Word.)", italics: true, size: 18, color: c(P.secondary), font: { ascii: "Calibri" } })] }),
  new Paragraph({ children: [new PageBreak()] }),
];

// ── Build body sections ──
const bodyChildren = [

// ═══════════════════════════════════════════════════════
// SECTION 1: EXECUTIVE SUMMARY
// ═══════════════════════════════════════════════════════
h1("1. Executive Summary"),
body("This document provides a corrected, codebase-accurate PostgreSQL integration plan for the Hukum-Aneh Indonesian Legal AI pipeline. The plan migrates legal document storage from Google Firestore to PostgreSQL, adds Qdrant as a dedicated vector store to replace the current brute-force in-memory cosine similarity search, and retains Firestore exclusively for user/application data (Telegram user sessions, bot state). Each phase is broken into agent-executable subtasks with exact file paths, code structures, and acceptance criteria derived from a thorough audit of the actual codebase."),
body("The original plan contained 9 significant gaps when evaluated against the real codebase: it missed 7 of 10 data structs in its schema design, did not account for the 13-state pipeline lifecycle (including the pause-and-resume pattern at pending_prompt_approval), ignored the local JSON file queue fallback system, overlooked mock fallback strategies for AI/embedding services, listed TikTok as a publish target despite zero implementation existing, and mischaracterized the connector architecture. This corrected plan addresses all of these gaps and provides a complete, executable migration path."),
body("Target Architecture: PostgreSQL (all legal pipeline data: laws, versions, analyses, drafts, captions, image assets, approvals, publishing jobs, embeddings metadata) + Qdrant (1536-dim embedding vectors with cosine similarity) + Firestore (retained for user/bot state only) + Google Drive or S3 (PDF raw file storage, replacing local filesystem)."),

// ═══════════════════════════════════════════════════════
// SECTION 2: CURRENT ARCHITECTURE ANALYSIS
// ═══════════════════════════════════════════════════════
h1("2. Current Architecture Analysis"),

h2("2.1 Data Model (10 Structs)"),
body("The codebase defines 10 distinct data structs in backend/internal/models/models.go. These map to 6 top-level Firestore collections and 4 subcollections. The structs are: LawDocument (top-level), LawVersion (subcollection: laws/{id}/versions), LawRelationship (subcollection: laws/{id}/relationships), LawAnalysis (subcollection: laws/{id}/analyses), ContentDraft (top-level), Caption (subcollection: content_drafts/{id}/captions), ImageAsset (top-level), Approval (top-level), PublishingJob (top-level), and EmbeddingEntry (top-level). Collection name constants are centralized in backend/internal/models/collections.go."),
body("Critical fields that the original plan missed: LawAnalysis contains 4 scoring dimensions (controversy_score, economic_score, legal_consistency, overall_score) plus a confidence float and a raw_json string storing the full LLM response. ContentDraft has a 6-state status lifecycle (draft, pending_prompt_approval, pending_approval, approved, rejected, published). EmbeddingEntry has an is_mock boolean field that must be preserved in PostgreSQL to distinguish real Gemini embeddings from deterministic fallback vectors. LawVersion.TextContent can exceed 1MB for large Indonesian regulations, requiring PostgreSQL TEXT type (TOAST storage)."),

h2("2.2 Repository Layer (Monolithic)"),
body("The entire data access layer lives in a single 378-line file: backend/internal/repository/firestore.go. The FirestoreRepo struct wraps a single *firestore.Client and implements all CRUD operations for all 10 models as methods on one concrete struct. There are zero interfaces defined anywhere in the codebase. Every consumer (Engine, retrieval.Service, and 6 cmd/ entrypoints) imports and depends directly on *repository.FirestoreRepo. This makes it impossible to swap the storage backend without modifying every consumer. Interface extraction is the critical prerequisite before any PostgreSQL work can begin."),
body("The repository methods follow a consistent pattern: if ID is empty, call Collection.Add() (Firestore auto-generates ID) and set doc.ID = ref.ID; if ID is non-empty, call Doc(id).Set() (upsert). This pattern must be replicated in the PostgreSQL implementation. Notably, GetLawAnalysisByDraft() in firestore.go uses a CollectionGroup query to find an analysis across all law subcollections by its document ID, then filters in-memory. This is inefficient and must be replaced with a direct FK query in PostgreSQL."),

h2("2.3 Pipeline Lifecycle (13 States)"),
body("The Engine in backend/internal/workflow/engine.go implements a 13-state document lifecycle. The full state machine is: discovered -> downloaded -> parsed -> analyzed -> no_conflict (terminal stop) | pending_prompt_approval (pipeline pauses, waits for Telegram callback) -> pending_approval (image generated, waits for human review) -> approved -> published. Error terminal states: download_failed, parse_failed. Rejection paths lead to archived. The ContentDraft has its own parallel 6-state lifecycle: draft -> pending_prompt_approval -> pending_approval -> approved -> rejected -> published. The LawDocument status and ContentDraft status are updated in lockstep at gate transitions."),
body("The pause-and-resume pattern is critical: after the Prompt Builder Agent generates an image_prompt, the pipeline saves the draft with status pending_prompt_approval and sends a Telegram inline keyboard to the reviewer. The pipeline process exits. When the user taps a button on Telegram, the callback arrives via cmd/bot/main.go, which calls Engine.HandleApprovalAction() with the draftID, action, and reviewerID. This method re-fetches the draft, analysis, and law document from the repository, then resumes processing based on the action (prompt_approve, prompt_reject, prompt_regen, approve, reject, regen_img, regen_cap). The PostgreSQL schema must preserve this pause-and-resume capability with proper transaction isolation."),

h2("2.4 Vector Search (Brute-Force)"),
body("The retrieval.Service in backend/internal/retrieval/retrieval.go currently loads ALL embeddings from Firestore into memory via repo.ListAllEmbeddings(), then computes cosine similarity in a nested loop with O(n) bubble sort. This approach will not scale past approximately 5,000 records. The embedding dimension is 1536 (matching Gemini gemini-embedding-2 model). Mock fallback vectors (linear gradient) are generated when the Gemini API returns auth/quota errors. These mock embeddings are flagged with is_mock=true in the EmbeddingEntry struct. The migration to Qdrant must handle this is_mock flag: real embeddings go to Qdrant, mock embeddings stay in PostgreSQL only and are excluded from similarity search."),

h2("2.5 Connector Architecture (15 Connectors, Not 20+)"),
body("The codebase has 15 registered connectors in backend/internal/connectors/. These are: peraturan (handles 6 document types: UU, PP, Perpres, Permen, Kepmen, PBI), setneg, jdihn, bpk, mkri, kemenkeu, ma, kemnaker, kemendag, komdigi, kpu, bkn, lkpp, and dpr. The peraturan connector alone handles 6 of the commonly-cited document types, so the original plan's enumeration of separate connectors per type was incorrect. TikTok has zero implementation in the codebase despite being mentioned as a publish target. Instagram publishing exists but is broken because it requires a public image URL and the current code has no image hosting (S3/GCS/Imgur)."),

h2("2.6 Additional Systems"),
body("Local JSON file queue: cmd/flush_local/main.go implements a fallback queue at backend/internal/storage/local_queue/ that writes law documents as JSON files when Firestore is unavailable. The flush command reads these files, deduplicates via FindByLawNumber(), and pushes to Firestore. This system must be updated to target PostgreSQL. Mock fallbacks: Both the AI service (ai.go) and retrieval service (retrieval.go) return hardcoded mock responses on 401/429 API errors. These mock responses are saved to the database. The PostgreSQL schema must accommodate these. System dependencies: The parser requires poppler-utils (pdftotext, pdftoppm), tesseract (ind+eng OCR languages), and Python 3 with curl_cffi + beautifulsoup4 for the scraper bridge. Any Dockerfile must include these."),

// ═══════════════════════════════════════════════════════
// SECTION 3: TARGET ARCHITECTURE
// ═══════════════════════════════════════════════════════
h1("3. Target Architecture"),

h2("3.1 Storage Responsibilities"),
makeTable(
  ["Storage", "Responsibility", "Data Types"],
  [
    ["PostgreSQL", "All legal pipeline data", "laws, law_versions, law_relationships, law_analyses, content_drafts, captions, image_assets, approvals, publishing_jobs, embedding_metadata"],
    ["Qdrant", "Vector similarity search", "1536-dim float32 vectors with metadata (law_document_id, is_mock filter)"],
    ["Firestore (retained)", "User/app data only", "Telegram user sessions, bot state, approval callbacks (if needed for real-time)"],
    ["Google Drive / S3", "Raw PDF files", "Downloaded regulation PDFs (currently stored at backend/internal/storage/)"],
  ]
),

h2("3.2 Dependency Graph Overview"),
body("Phase 0 (Interface Abstraction) must complete before any other phase. Phase 1 (PostgreSQL Schema) and Phase 2 (Qdrant Setup) can run in parallel after Phase 0. Phase 3 (PostgreSQL Repository) depends on Phase 1. Phase 4 (Qdrant Retrieval) depends on Phase 2. Phase 5 (Dual-Write) depends on Phases 3 and 4. Phase 6 (Data Migration) depends on Phase 5. Phase 7 (Cleanup) depends on Phase 6."),

// ═══════════════════════════════════════════════════════
// SECTION 4: PHASE 0 - INTERFACE ABSTRACTION
// ═══════════════════════════════════════════════════════
h1("4. Phase 0: Repository Interface Abstraction (Prerequisite)"),
body("This phase is the critical prerequisite. The current codebase has zero interface abstraction for the repository layer. Every consumer (Engine, retrieval.Service, and 6 cmd/ files) depends directly on *repository.FirestoreRepo. Before any PostgreSQL implementation can begin, we must extract interfaces so that consumers depend on abstractions, not concretions."),

h2("4.1 Task 0.1: Define Repository Interfaces"),
bodyBold("File: backend/internal/repository/interfaces.go (NEW)"),
body("Create a new file defining granular repository interfaces. Each interface should cover one domain aggregate. This follows the Interface Segregation Principle and makes it possible to mock individual concerns in tests. The interfaces must exactly match the method signatures already used by consumers in the codebase."),
code('type LawDocumentRepo interface {'),
code('  SaveLawDocument(ctx context.Context, doc *models.LawDocument) (string, error)'),
code('  GetLawDocument(ctx context.Context, id string) (*models.LawDocument, error)'),
code('  FindByLawNumber(ctx context.Context, lawNumber string) (*models.LawDocument, error)'),
code('  ListLawsByStatus(ctx context.Context, status string) ([]models.LawDocument, error)'),
code('  ListAllLaws(ctx context.Context) ([]models.LawDocument, error)'),
code('  FindStuckDocuments(ctx context.Context, status string, before time.Time) ([]models.LawDocument, error)'),
code('}'),
body(""),
code('type LawVersionRepo interface {'),
code('  GetLatestLawVersion(ctx context.Context, lawID string) (*models.LawVersion, error)'),
code('  SaveLawVersion(ctx context.Context, lawID string, v *models.LawVersion) (string, error)'),
code('}'),
body(""),
code('type LawAnalysisRepo interface {'),
code('  SaveLawAnalysis(ctx context.Context, lawID string, a *models.LawAnalysis) (string, error)'),
code('  GetLawAnalysisByDraft(ctx context.Context, draftID string) (*models.LawAnalysis, error)'),
code('}'),
body(""),
code('type ContentDraftRepo interface {'),
code('  GetContentDraft(ctx context.Context, id string) (*models.ContentDraft, error)'),
code('  SaveContentDraft(ctx context.Context, draft *models.ContentDraft) (string, error)'),
code('}'),
body(""),
code('type ImageAssetRepo interface {'),
code('  GetImageAssetsByDraft(ctx context.Context, draftID string) ([]models.ImageAsset, error)'),
code('  SaveImageAsset(ctx context.Context, asset *models.ImageAsset) (string, error)'),
code('}'),
body(""),
code('type ApprovalRepo interface {'),
code('  SaveApproval(ctx context.Context, a *models.Approval) (string, error)'),
code('}'),
body(""),
code('type PublishingJobRepo interface {'),
code('  SavePublishingJob(ctx context.Context, j *models.PublishingJob) (string, error)'),
code('}'),
body(""),
code('type EmbeddingRepo interface {'),
code('  SaveEmbedding(ctx context.Context, emb *models.EmbeddingEntry) (string, error)'),
code('  ListAllEmbeddings(ctx context.Context) ([]models.EmbeddingEntry, error)'),
code('}'),
body("Acceptance Criteria: (1) File compiles with go build ./... (2) All method signatures exactly match existing FirestoreRepo method signatures (3) Each interface is in its own named type, not one mega-interface."),

h2("4.2 Task 0.2: Update FirestoreRepo to Implement Interfaces"),
bodyBold("File: backend/internal/repository/firestore.go (MODIFY)"),
body("Add compile-time interface assertions at the top of firestore.go to guarantee FirestoreRepo satisfies all interfaces. No logic changes needed, just assertions:"),
code('var _ repository.LawDocumentRepo = (*FirestoreRepo)(nil)'),
code('var _ repository.LawVersionRepo = (*FirestoreRepo)(nil)'),
code('var _ repository.LawAnalysisRepo = (*FirestoreRepo)(nil)'),
code('var _ repository.ContentDraftRepo = (*FirestoreRepo)(nil)'),
code('var _ repository.ImageAssetRepo = (*FirestoreRepo)(nil)'),
code('var _ repository.ApprovalRepo = (*FirestoreRepo)(nil)'),
code('var _ repository.PublishingJobRepo = (*FirestoreRepo)(nil)'),
code('var _ repository.EmbeddingRepo = (*FirestoreRepo)(nil)'),
body("Acceptance Criteria: (1) go build passes (2) If any interface method is missing or signature-mismatched, the compiler will fail at the assertion line, providing an immediate error."),

h2("4.3 Task 0.3: Refactor Engine to Accept Interfaces"),
bodyBold("File: backend/internal/workflow/engine.go (MODIFY)"),
body("Change the Engine struct field from repo *repository.FirestoreRepo to individual interface fields. This is the main refactoring step that unblocks all subsequent phases."),
code('// BEFORE:'),
code('type Engine struct {'),
code('  repo *repository.FirestoreRepo'),
code('  // ...'),
code('}'),
body(""),
code('// AFTER:'),
code('type Engine struct {'),
code('  lawRepo      repository.LawDocumentRepo'),
code('  versionRepo   repository.LawVersionRepo'),
code('  analysisRepo  repository.LawAnalysisRepo'),
code('  draftRepo     repository.ContentDraftRepo'),
code('  imageRepo     repository.ImageAssetRepo'),
code('  approvalRepo  repository.ApprovalRepo'),
code('  publishRepo   repository.PublishingJobRepo'),
code('  embedRepo     repository.EmbeddingRepo'),
code('  // ...'),
code('}'),
body("Update NewEngine() to accept individual interface parameters (or accept a single struct containing all interfaces for convenience). Then update every e.repo.XXX call in engine.go to use the appropriate interface field. There are approximately 30 call sites to update. The mapping is: e.repo.SaveLawDocument -> e.lawRepo.SaveLawDocument, e.repo.GetLawDocument -> e.lawRepo.GetLawDocument, e.repo.FindByLawNumber -> e.lawRepo.FindByLawNumber, e.repo.SaveLawVersion -> e.versionRepo.SaveLawVersion, e.repo.GetLatestLawVersion -> e.versionRepo.GetLatestLawVersion, e.repo.SaveLawAnalysis -> e.analysisRepo.SaveLawAnalysis, e.repo.GetLawAnalysisByDraft -> e.analysisRepo.GetLawAnalysisByDraft, e.repo.SaveContentDraft -> e.draftRepo.SaveContentDraft, e.repo.GetContentDraft -> e.draftRepo.GetContentDraft, e.repo.SaveImageAsset -> e.imageRepo.SaveImageAsset, e.repo.GetImageAssetsByDraft -> e.imageRepo.GetImageAssetsByDraft, e.repo.SaveApproval -> e.approvalRepo.SaveApproval, e.repo.SavePublishingJob -> e.publishRepo.SavePublishingJob, e.repo.SaveEmbedding -> e.embedRepo.SaveEmbedding, e.repo.FindStuckDocuments -> e.lawRepo.FindStuckDocuments, e.repo.ListLawsByStatus -> e.lawRepo.ListLawsByStatus."),
body("Acceptance Criteria: (1) go build passes (2) All 30+ e.repo.XXX calls updated to e.{specific}Repo.XXX (3) No remaining references to *repository.FirestoreRepo in engine.go (4) Engine pipeline still compiles and links correctly."),

h2("4.4 Task 0.4: Update retrieval.Service"),
bodyBold("File: backend/internal/retrieval/retrieval.go (MODIFY)"),
body("Change the Service struct to depend on EmbeddingRepo interface instead of *repository.FirestoreRepo. The retrieval.Service only uses repo.ListAllEmbeddings() and repo.SaveEmbedding(), so it only needs the EmbeddingRepo interface."),
code('// BEFORE:'),
code('type Service struct {'),
code('  cfg    *config.Config'),
code('  repo   *repository.FirestoreRepo'),
code('  client *genai.Client'),
code('  sem    chan struct{}'),
code('}'),
body(""),
code('// AFTER:'),
code('type Service struct {'),
code('  cfg    *config.Config'),
code('  repo   repository.EmbeddingRepo'),
code('  client *genai.Client'),
code('  sem    chan struct{}'),
code('}'),
body("Acceptance Criteria: (1) go build passes (2) retrieval.Service no longer imports Firestore client."),

h2("4.5 Task 0.5: Update All cmd/ Entry Points"),
bodyBold("Files to modify:"),
bullet("backend/cmd/pipeline/main.go - Pass FirestoreRepo as interfaces to NewEngine()"),
bullet("backend/cmd/bot/main.go - Same pattern"),
bullet("backend/cmd/backfill/main.go - Same pattern"),
bullet("backend/cmd/batch/main.go - Same pattern"),
bullet("backend/cmd/backfill_embeddings/main.go - Uses repo.SaveEmbedding and repo.ListAllEmbeddings"),
bullet("backend/cmd/flush_local/main.go - Uses repo.FindByLawNumber, repo.SaveLawDocument, repo.SaveLawVersion"),
body("Each cmd/ file currently creates a *repository.FirestoreRepo and passes it directly. After this change, each file continues creating FirestoreRepo (which implements all interfaces) but passes it through the interface typed parameters. No runtime behavior changes."),
body("Acceptance Criteria: (1) All 6 cmd/ files compile (2) go build ./... passes with zero errors (3) go test ./... still passes (if tests exist)."),

// ═══════════════════════════════════════════════════════
// SECTION 5: PHASE 1 - POSTGRESQL SCHEMA
// ═══════════════════════════════════════════════════════
h1("5. Phase 1: PostgreSQL Schema Design"),

h2("5.1 Task 1.1: Add PostgreSQL Dependencies"),
bodyBold("File: go.mod (MODIFY)"),
body("Add the following dependencies using go get:"),
code('go get github.com/jackc/pgx/v5'),
code('go get github.com/jackc/pgx/v5/pgxpool'),
code('go get github.com/golang-migrate/migrate/v4'),
code('go get github.com/golang-migrate/migrate/v4/database/pgx/v5'),
code('go get github.com/golang-migrate/migrate/v4/source/file'),
body("Also add Qdrant client:"),
code('go get github.com/qdrant/go-client/qdrant'),
body("Acceptance Criteria: (1) go mod tidy completes without errors (2) go build ./... still passes."),

h2("5.2 Task 1.2: Create Migration Files"),
bodyBold("Directory: backend/migrations/ (NEW)"),
body("Create migration files using golang-migrate naming convention. The initial migration must create all 10 tables matching the 10 Go structs. Critical design decisions for each table:"),

h3("Table: law_documents"),
body("Maps to: models.LawDocument. Primary key: id (TEXT, matching Firestore's auto-generated string IDs). The status column must have a CHECK constraint enforcing the 13 valid states. The raw_file_path column stores the local or remote path to the PDF file. Index on (law_number) UNIQUE for deduplication. Index on (status, updated_at) for the FindStuckDocuments query and ListLawsByStatus."),
code('CREATE TABLE law_documents ('),
code('  id              TEXT PRIMARY KEY,'),
code('  law_number      TEXT NOT NULL,'),
code('  title           TEXT NOT NULL,'),
code('  source_url      TEXT,'),
code('  source          TEXT NOT NULL,'),
code('  level           TEXT NOT NULL DEFAULT \'national\','),
code('  document_type   TEXT NOT NULL DEFAULT \'UU\','),
code('  raw_file_path   TEXT,'),
code('  published_date  TEXT,'),
code('  status          TEXT NOT NULL DEFAULT \'discovered\''),
code('    CHECK (status IN (\'discovered\',\'downloaded\',\'parsed\',\'analyzed\','),
code('      \'no_conflict\',\'pending_prompt_approval\',\'pending_approval\','),
code('      \'approved\',\'published\',\'archived\','),
code('      \'download_failed\',\'parse_failed\')),'),
code('  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),'),
code('  updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()'),
code(');'),
code('CREATE UNIQUE INDEX idx_law_documents_law_number ON law_documents(law_number);'),
code('CREATE INDEX idx_law_documents_status_updated ON law_documents(status, updated_at);'),

h3("Table: law_versions"),
body("Maps to: models.LawVersion. Critical: text_content uses TEXT type (PostgreSQL TOAST handles values >8KB automatically, supporting the 1MB+ regulation texts). The embedding field ([]float32) is NOT stored here - it moves to Qdrant. Only embedding_id (TEXT reference to Qdrant point ID) is kept. Foreign key to law_documents(id)."),
code('CREATE TABLE law_versions ('),
code('  id              TEXT PRIMARY KEY,'),
code('  law_document_id TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,'),
code('  version_number  INTEGER NOT NULL DEFAULT 1,'),
code('  text_content    TEXT NOT NULL,'),
code('  embedding_id    TEXT,'),
code('  parsed_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()'),
code(');'),
code('CREATE INDEX idx_law_versions_law_id ON law_versions(law_document_id, version_number DESC);'),

h3("Table: law_relationships"),
body("Maps to: models.LawRelationship. Currently a subcollection under each law. Flattened to top-level table with FK. The relationship_type must have a CHECK constraint for valid values (amends, repeals, supersedes, references)."),
code('CREATE TABLE law_relationships ('),
code('  id                TEXT PRIMARY KEY,'),
code('  law_document_id   TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,'),
code('  related_law_number TEXT NOT NULL,'),
code('  relationship_type TEXT NOT NULL CHECK (relationship_type IN (\'amends\',\'repeals\',\'supersedes\',\'references\')),'),
code('  article_ref       TEXT'),
code(');'),

h3("Table: law_analyses"),
body("Maps to: models.LawAnalysis. The affected_laws field ([]AffectedLaw) is stored as JSONB. The 4 score columns (controversy_score, economic_score, legal_consistency, overall_score) are INTEGER with CHECK 0-100. The raw_json field stores the complete LLM response string. Confidence is FLOAT8."),
code('CREATE TABLE law_analyses ('),
code('  id               TEXT PRIMARY KEY,'),
code('  law_document_id  TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,'),
code('  summary          TEXT NOT NULL,'),
code('  affected_laws    JSONB NOT NULL DEFAULT \'[]\','),
code('  overall_score    INTEGER NOT NULL CHECK (overall_score BETWEEN 0 AND 100),'),
code('  controversy_score INTEGER NOT NULL CHECK (controversy_score BETWEEN 0 AND 100),'),
code('  economic_score    INTEGER NOT NULL CHECK (economic_score BETWEEN 0 AND 100),'),
code('  legal_consistency INTEGER NOT NULL CHECK (legal_consistency BETWEEN 0 AND 100),'),
code('  confidence       FLOAT8 NOT NULL DEFAULT 0.0,'),
code('  raw_json         TEXT NOT NULL DEFAULT \'\','),
code('  created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()'),
code(');'),
code('CREATE INDEX idx_law_analyses_law_id ON law_analyses(law_document_id);'),

h3("Table: content_drafts"),
body("Maps to: models.ContentDraft. The hashtags field ([]string) is stored as JSONB. The status has a CHECK constraint for the 6 valid states. Foreign key to law_analyses (not law_documents directly, matching the current code where ContentDraft.LawAnalysisID references the analysis)."),
code('CREATE TABLE content_drafts ('),
code('  id               TEXT PRIMARY KEY,'),
code('  law_analysis_id  TEXT NOT NULL REFERENCES law_analyses(id),'),
code('  caption          TEXT NOT NULL DEFAULT \'\','),
code('  hashtags         JSONB NOT NULL DEFAULT \'[]\','),
code('  hook             TEXT NOT NULL DEFAULT \'\','),
code('  image_prompt     TEXT NOT NULL DEFAULT \'\','),
code('  status           TEXT NOT NULL DEFAULT \'draft\''),
code('    CHECK (status IN (\'draft\',\'pending_prompt_approval\',\'pending_approval\','),
code('      \'approved\',\'rejected\',\'published\',\'prompt_rejected\')),'),
code('  created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()'),
code(');'),

h3("Table: captions, image_assets, approvals, publishing_jobs"),
body("These are straightforward mappings from the remaining structs. Captions has FK to content_drafts. ImageAssets has FK to content_drafts. Approvals has FK to content_drafts and stores the reviewer_id (Telegram user ID). PublishingJobs has FK to content_drafts, with platform CHECK (instagram only currently - do NOT add tiktok), and posted_at as nullable TIMESTAMPTZ."),
code('CREATE TABLE captions ('),
code('  id               TEXT PRIMARY KEY,'),
code('  content_draft_id TEXT NOT NULL REFERENCES content_drafts(id) ON DELETE CASCADE,'),
code('  text             TEXT NOT NULL,'),
code('  variant_number   INTEGER NOT NULL DEFAULT 1'),
code(');'),
body(""),
code('CREATE TABLE image_assets ('),
code('  id                   TEXT PRIMARY KEY,'),
code('  content_draft_id     TEXT NOT NULL REFERENCES content_drafts(id) ON DELETE CASCADE,'),
code('  prompt_used          TEXT NOT NULL DEFAULT \'\','),
code('  file_path            TEXT NOT NULL,'),
code('  validated            BOOLEAN NOT NULL DEFAULT FALSE,'),
code('  design_guide_version TEXT NOT NULL DEFAULT \'\','),
code('  created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW()'),
code(');'),
body(""),
code('CREATE TABLE approvals ('),
code('  id               TEXT PRIMARY KEY,'),
code('  content_draft_id TEXT NOT NULL REFERENCES content_drafts(id),'),
code('  reviewer_id      TEXT NOT NULL,'),
code('  decision         TEXT NOT NULL CHECK (decision IN (\'approve\',\'reject\',\'regenerate_caption\',\'regenerate_image\',\'prompt_approve\',\'prompt_reject\',\'prompt_regen\')),'),
code('  reason           TEXT NOT NULL DEFAULT \'\','),
code('  timestamp        TIMESTAMPTZ NOT NULL DEFAULT NOW()'),
code(');'),
body(""),
code('CREATE TABLE publishing_jobs ('),
code('  id               TEXT PRIMARY KEY,'),
code('  content_draft_id TEXT NOT NULL REFERENCES content_drafts(id),'),
code('  platform         TEXT NOT NULL CHECK (platform IN (\'instagram\')),'),
code('  status           TEXT NOT NULL DEFAULT \'pending\' CHECK (status IN (\'pending\',\'pending_image_hosting\',\'published\',\'failed\')),'),
code('  posted_at        TIMESTAMPTZ,'),
code('  external_post_id TEXT NOT NULL DEFAULT \'\''),
code(');'),

h3("Table: embedding_metadata"),
body("Maps to: models.EmbeddingEntry. The vector field is NOT stored here - it moves to Qdrant. Only metadata (law_document_id, is_mock, created_at) is kept in PostgreSQL. The is_mock boolean is critical: it determines whether the vector was a real Gemini embedding or a deterministic fallback. Mock embeddings must NOT be upserted to Qdrant."),
code('CREATE TABLE embedding_metadata ('),
code('  id              TEXT PRIMARY KEY,'),
code('  law_document_id TEXT NOT NULL REFERENCES law_documents(id) ON DELETE CASCADE,'),
code('  is_mock         BOOLEAN NOT NULL DEFAULT FALSE,'),
code('  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()'),
code('  qdrant_point_id TEXT'),
code(');'),
code('CREATE INDEX idx_embedding_metadata_law_id ON embedding_metadata(law_document_id);'),
code('CREATE INDEX idx_embedding_metadata_mock ON embedding_metadata(is_mock);'),

body("Acceptance Criteria: (1) Migration file is named 000001_init_schema.up.sql and 000001_init_schema.down.sql (2) All 10 tables created (3) All foreign keys match the struct relationships (4) CHECK constraints cover all enum-like status fields (5) Indexes cover all query patterns used by existing repository methods (6) migrate up succeeds on a fresh PostgreSQL instance."),

h2("5.3 Task 1.3: Add PostgreSQL Config"),
bodyBold("File: backend/internal/config/config.go (MODIFY)"),
body("Add a PostgreSQL section and a Qdrant section to the Config struct:"),
code('Postgres struct {'),
code('  DSN             string  // json tag: dsn'),
code('  MaxOpenConns    int     // json tag: max_open_conns'),
code('  MaxIdleConns    int     // json tag: max_idle_conns'),
code('} // json tag: postgres'),
body(""),
code('Qdrant struct {'),
code('  URL       string  // json tag: url'),
code('  Collection string  // json tag: collection'),
code('  APIKey    string  // json tag: api_key'),
code('} // json tag: qdrant'),
body("Add environment variable loading in Load(): POSTGRES_DSN, QDRANT_URL, QDRANT_COLLECTION, QDRANT_API_KEY. Provide sensible defaults (localhost:5432, localhost:6333, collection name 'law_embeddings')."),
body("Acceptance Criteria: (1) Config compiles (2) Environment variables are loaded correctly (3) Backward compatible - existing FIREBASE_* vars still work."),

// ═══════════════════════════════════════════════════════
// SECTION 6: PHASE 2 - QDRANT SETUP
// ═══════════════════════════════════════════════════════
h1("6. Phase 2: Qdrant Vector Store Setup"),

h2("6.1 Task 2.1: Create Qdrant Client Wrapper"),
bodyBold("File: backend/internal/retrieval/qdrant.go (NEW)"),
body("Create a QdrantClient wrapper that handles collection creation, upsert, and search. The wrapper must use the qdrant-go-client library. Key requirements: (1) Collection creation with 1536 dimensions, cosine similarity metric, and HNSW index. (2) Upsert method that stores vectors with payload containing law_document_id and is_mock. (3) Search method that accepts a query vector, topN, and optional is_mock=false filter. (4) The wrapper should auto-create the collection on first use if it does not exist."),
code('type QdrantClient struct {'),
code('  client    *qdrant.Client'),
code('  collection string'),
code('  dims      uint64'),
code('}'),
body(""),
code('func (q *QdrantClient) Upsert(ctx context.Context, id string, vector []float32, lawDocumentID string, isMock bool) error'),
code('func (q *QdrantClient) Search(ctx context.Context, vector []float32, topN int) ([]SearchResult, error)'),
code('func (q *QdrantClient) EnsureCollection(ctx context.Context) error'),
body("The Search method must filter out is_mock=true vectors using a Qdrant filter condition: must NOT match { key: 'is_mock', match: { value: true } }. This ensures mock embeddings never appear in similarity results."),
body("Acceptance Criteria: (1) File compiles (2) Collection is created with correct dimensions and metric (3) Upsert stores vectors with correct payload (4) Search returns results excluding mock vectors (5) Integration test passes against a running Qdrant instance."),

h2("6.2 Task 2.2: Create Docker Compose for Qdrant"),
bodyBold("File: docker-compose.yml (NEW at project root)"),
body("Create a docker-compose.yml that defines Qdrant and PostgreSQL services. This is the foundation for the development and deployment environment."),
code('services:'),
code('  postgres:'),
code('    image: postgres:16-alpine'),
code('    environment:'),
code('      POSTGRES_DB: hukumaneh'),
code('      POSTGRES_USER: hukumaneh'),
code('      POSTGRES_PASSWORD: \${POSTGRES_PASSWORD}'),
code('    ports: [\"5432:5432\"]"),
code('    volumes: [pgdata:/var/lib/postgresql/data]'),
code('  qdrant:'),
code('    image: qdrant/qdrant:latest'),
code('    ports: [\"6333:6333\", \"6334:6334\"]"),
code('    volumes: [qdrant_storage:/qdrant/storage]'),
code('volumes:'),
code('  pgdata:'),
code('  qdrant_storage:'),
body("Acceptance Criteria: (1) docker compose up -d starts both services (2) PostgreSQL is accessible on port 5432 (3) Qdrant dashboard is accessible on port 6333."),

// ═══════════════════════════════════════════════════════
// SECTION 7: PHASE 3 - POSTGRESQL REPOSITORY
// ═══════════════════════════════════════════════════════
h1("7. Phase 3: PostgreSQL Repository Implementation"),

h2("7.1 Task 3.1: Create PostgresRepo Struct"),
bodyBold("File: backend/internal/repository/postgres.go (NEW)"),
body("Create a PostgresRepo struct that implements all 8 repository interfaces. The struct wraps a *pgxpool.Pool. Use a single struct (like FirestoreRepo) that satisfies all interfaces, keeping implementation simple. Use compile-time assertions just like in Task 0.2."),
code('type PostgresRepo struct {'),
code('  pool *pgxpool.Pool'),
code('}'),
body(""),
code('func NewPostgresRepo(ctx context.Context, dsn string) (*PostgresRepo, error) {'),
code('  pool, err := pgxpool.New(ctx, dsn)'),
code('  if err != nil { return nil, err }'),
code('  if err := pool.Ping(ctx); err != nil { return nil, err }'),
code('  return &PostgresRepo{pool: pool}, nil'),
code('}'),
body(""),
code('func (r *PostgresRepo) Close() { r.pool.Close() }'),
body("Implement each interface method using pgx SQL queries. Key implementation notes: (1) For SaveLawDocument: use INSERT ... ON CONFLICT (id) DO UPDATE (PostgreSQL upsert). If ID is empty, generate a UUID or nanoid string. (2) For FindByLawNumber: SELECT ... WHERE law_number = $1 LIMIT 1. (3) For ListLawsByStatus: SELECT ... WHERE status = $1. (4) For FindStuckDocuments: SELECT ... WHERE status = $1 AND updated_at < $2. (5) For GetLatestLawVersion: SELECT ... WHERE law_document_id = $1 ORDER BY version_number DESC LIMIT 1. (6) For GetLawAnalysisByDraft: JOIN content_drafts and law_analyses (no more CollectionGroup scan)."),
body("For JSONB fields (affected_laws, hashtags): use pgx's native JSON/JSONB support. Go structs will need json.Marshal/json.Unmarshal at the boundary. For the affected_laws field, define a custom scanner/valuer or handle conversion in the repository methods."),
body("Acceptance Criteria: (1) All 8 interface assertions compile (2) Each method implements the exact same behavior as the Firestore equivalent (3) Save with empty ID auto-generates, Save with existing ID upserts (4) All queries use parameterized SQL ($1, $2, etc.) - no string interpolation."),

h2("7.2 Task 3.2: Implement LawDocument Methods"),
bodyBold("File: backend/internal/repository/postgres.go (CONTINUE)"),
body("Implement the 6 LawDocumentRepo methods. Pay special attention to SaveLawDocument which must handle both insert and update patterns. The Firestore version uses Collection.Add() for new docs (auto-ID) and Doc(id).Set() for updates. In PostgreSQL, use INSERT ... ON CONFLICT (id) DO UPDATE. For auto-ID generation when ID is empty, use a UUID v4 or a nanoid-like short string generator to maintain compatibility with existing Firestore IDs (which are random alphanumeric strings)."),
body("The FindStuckDocuments query is critical for the scheduler. The current Firestore query uses two WHERE clauses (status == and updated_at <). In PostgreSQL, this is a simple compound WHERE with the existing composite index."),

h2("7.3 Task 3.3: Implement All Remaining Methods"),
bodyBold("File: backend/internal/repository/postgres.go (CONTINUE)"),
body("Implement all remaining interface methods for LawVersionRepo (2 methods), LawAnalysisRepo (2 methods), ContentDraftRepo (2 methods), ImageAssetRepo (2 methods), ApprovalRepo (1 method), PublishingJobRepo (1 method), and EmbeddingRepo (2 methods). Total: approximately 16 methods. The most complex is GetLawAnalysisByDraft, which currently scans all analyses across all law subcollections. In PostgreSQL, this becomes a simple JOIN: SELECT a.* FROM law_analyses a JOIN content_drafts d ON a.id = d.law_analysis_id WHERE d.id = $1."),
body("For EmbeddingRepo: SaveEmbedding should insert into embedding_metadata table only (the vector itself goes to Qdrant, handled by the caller). ListAllEmbeddings should select from embedding_metadata (no vector data). This is a deliberate design choice: PostgreSQL only stores metadata about embeddings, while Qdrant stores the actual vectors."),
body("Acceptance Criteria: (1) All 16+ methods implemented (2) go build ./... passes (3) Manual test: create PostgresRepo, call Save/Get cycles for each entity type."),

// ═══════════════════════════════════════════════════════
// SECTION 8: PHASE 4 - QDRANT RETRIEVAL INTEGRATION
// ═══════════════════════════════════════════════════════
h1("8. Phase 4: Qdrant Retrieval Integration"),

h2("8.1 Task 4.1: Update retrieval.Service for Qdrant"),
bodyBold("File: backend/internal/retrieval/retrieval.go (MODIFY)"),
body("Modify the Service struct to include a QdrantClient alongside the EmbeddingRepo. The GenerateEmbedding method remains unchanged (still calls Gemini API with mock fallback). The Search method must be rewritten to use Qdrant instead of the brute-force in-memory approach."),
code('// BEFORE (brute-force):'),
code('func (s *Service) Search(ctx context.Context, queryVector []float32, topN int) ([]SearchResult, error) {'),
code('  allEmbs, err := s.repo.ListAllEmbeddings(ctx)  // loads ALL vectors into memory!'),
code('  // ... O(n) cosine similarity + bubble sort ...'),
code('}'),
body(""),
code('// AFTER (Qdrant):'),
code('func (s *Service) Search(ctx context.Context, queryVector []float32, topN int) ([]SearchResult, error) {'),
code('  if s.qdrant == nil {'),
code('    // Fallback to brute-force if Qdrant not configured'),
code('    return s.bruteForceSearch(ctx, queryVector, topN)'),
code('  }'),
code('  return s.qdrant.Search(ctx, queryVector, topN)'),
code('}'),
body("Keep the brute-force implementation as a fallback (renamed to bruteForceSearch) for environments without Qdrant. This preserves the current behavior as a degradation path."),
body("Acceptance Criteria: (1) Search uses Qdrant when qdrant client is configured (2) Falls back to brute-force when qdrant is nil (3) Mock embeddings (is_mock=true) are excluded from Qdrant search results via filter."),

h2("8.2 Task 4.2: Upsert to Qdrant on Embed"),
bodyBold("File: backend/internal/workflow/engine.go (MODIFY)"),
body("In the ProcessParsedDocument method, after generating an embedding and saving it to the database via embedRepo, also upsert the vector to Qdrant if it is not a mock. Add a new interface or method to handle this. The simplest approach is to add the QdrantClient directly to the Engine struct and call it after embedding save."),
code('// In ProcessParsedDocument, after saving embedding:'),
code('if !isMock && e.qdrant != nil {'),
code('  if err := e.qdrant.Upsert(ctx, embID, vector, doc.ID, false); err != nil {'),
code('    e.logger.Warn("qdrant upsert failed, embedding only in PG", "error", err)"),
code('  }'),
code('}'),
body("Acceptance Criteria: (1) Real embeddings are upserted to Qdrant (2) Mock embeddings are NOT upserted (3) Pipeline continues even if Qdrant upsert fails (log warning, do not halt)."),

// ═══════════════════════════════════════════════════════
// SECTION 9: PHASE 5 - DUAL-WRITE & WIRING
// ═══════════════════════════════════════════════════════
h1("9. Phase 5: Dual-Write Configuration and Wiring"),

h2("9.1 Task 5.1: Create Repository Factory"),
bodyBold("File: backend/internal/repository/factory.go (NEW)"),
body("Create a factory function that instantiates the correct repository based on configuration. During migration, both repositories can be active. The factory supports three modes: 'firestore' (original), 'postgres' (new), and 'dual' (writes to both, reads from postgres)."),
code('type RepoSet struct {'),
code('  LawDocumentRepo  LawDocumentRepo'),
code('  LawVersionRepo   LawVersionRepo'),
code('  LawAnalysisRepo   LawAnalysisRepo'),
code('  ContentDraftRepo  ContentDraftRepo'),
code('  ImageAssetRepo    ImageAssetRepo'),
code('  ApprovalRepo     ApprovalRepo'),
code('  PublishingJobRepo PublishingJobRepo'),
code('  EmbeddingRepo     EmbeddingRepo'),
code('}'),
body(""),
code('func NewRepoSet(ctx context.Context, cfg *config.Config) (*RepoSet, func(), error) {'),
code('  mode := cfg.Storage.Mode  // \"firestore\", \"postgres\", or \"dual\""),
code('  // ... instantiate based on mode ...'),
code('}'),
body("For 'dual' mode, create a DualWriteRepo wrapper that writes to both backends and reads from the primary (postgres). This ensures no data is lost during migration while allowing validation that PostgreSQL writes are correct."),
body("Acceptance Criteria: (1) Factory creates correct repo based on config (2) 'firestore' mode works identically to current behavior (3) 'dual' mode writes to both backends (4) Close function cleans up all connections."),

h2("9.2 Task 5.2: Update All cmd/ Files for New Factory"),
bodyBold("Files: All 6 cmd/*/main.go files (MODIFY)"),
body("Replace the direct FirestoreRepo instantiation with the factory call. Each cmd/ file currently has code like: repo, err := repository.NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath). Replace with: repos, cleanup, err := repository.NewRepoSet(ctx, cfg). Then pass the individual interface fields from repos to NewEngine() and retrieval.New()."),
body("Acceptance Criteria: (1) All 6 cmd/ files use the factory (2) STORAGE_MODE=firestore makes everything work identically to before (3) STORAGE_MODE=postgres initializes PostgreSQL connections."),

// ═══════════════════════════════════════════════════════
// SECTION 10: PHASE 6 - DATA MIGRATION
// ═══════════════════════════════════════════════════════
h1("10. Phase 6: Data Migration Script"),

h2("10.1 Task 6.1: Create Migration Tool"),
bodyBold("File: backend/cmd/migrate_to_pg/main.go (NEW)"),
body("Create a one-time migration tool that reads all data from Firestore and writes it to PostgreSQL and Qdrant. The tool must: (1) Read all law_documents from Firestore, insert into PostgreSQL. (2) For each law, read all subcollections (versions, analyses, relationships) and insert into PostgreSQL with correct FK references. (3) Read all content_drafts, captions, image_assets, approvals, publishing_jobs from Firestore and insert into PostgreSQL. (4) Read all embeddings, insert metadata into PostgreSQL, and upsert non-mock vectors into Qdrant. (5) Preserve all original Firestore document IDs as PostgreSQL primary keys (they are string IDs, not UUIDs)."),
body("The migration must handle: (1) Subcollection flattening - Firestore subcollections (laws/{id}/versions, laws/{id}/analyses) become top-level PostgreSQL tables with FK. (2) The GetLawAnalysisByDraft workaround - since the current code uses a CollectionGroup scan, the migration must correctly link analyses to their law documents. (3) Mock embedding filtering - only non-mock embeddings are upserted to Qdrant. (4) Large text_content - PostgreSQL TOAST handles this automatically, but the migration must not truncate. (5) JSONB conversion - affected_laws and hashtags arrays must be properly JSON-encoded."),
body("Acceptance Criteria: (1) Migration tool reads all existing Firestore data (2) All data is written to PostgreSQL with correct relationships (3) All non-mock embeddings are in Qdrant (4) Row counts match between source (Firestore) and target (PostgreSQL) (5) Tool can be re-run idempotently (ON CONFLICT DO NOTHING)."),

h2("10.2 Task 6.2: Update Local Queue Flush"),
bodyBold("File: backend/cmd/flush_local/main.go (MODIFY)"),
body("The local queue flush command currently writes queued JSON files to Firestore. Update it to use the repository factory, so it writes to whatever backend is configured (PostgreSQL in the new setup). The dedup check (FindByLawNumber) and SaveLawDocument/SaveLawVersion calls remain the same - they just go through the interface now. Since Task 5.2 already updated this file to use the factory, this task may already be complete. Verify and test."),
body("Acceptance Criteria: (1) flush_local writes to PostgreSQL when STORAGE_MODE=postgres (2) Dedup check works against PostgreSQL (3) Queue files are deleted after successful push."),

// ═══════════════════════════════════════════════════════
// SECTION 11: PHASE 7 - CLEANUP & DOCKERFILE
// ═══════════════════════════════════════════════════════
h1("11. Phase 7: Cleanup, Dockerfile, and Testing"),

h2("11.1 Task 7.1: Create Multi-Stage Dockerfile"),
bodyBold("File: Dockerfile (NEW at project root)"),
body("Create a multi-stage Dockerfile for the pipeline binary. The final runtime image must include all system dependencies that the parser and scraper bridge require: poppler-utils (pdftotext for PDF text extraction, pdftoppm for image conversion), tesseract-ocr (with ind and eng language packs for Indonesian document OCR), and Python 3 with pip-installed curl_cffi and beautifulsoup4. The Go build stage should produce a static binary. The final image should be based on debian:bookworm-slim for minimum size while providing the required system libraries."),
code('FROM golang:1.25-bookworm AS builder'),
code('WORKDIR /app'),
code('COPY go.mod go.sum ./'),
code('RUN go mod download'),
code('COPY . .'),
code('RUN CGO_ENABLED=0 GOOS=linux go build -o /pipeline ./backend/cmd/pipeline/'),
code('FROM debian:bookworm-slim'),
code('RUN apt-get update && apt-get install -y --no-install-recommends \\\'),
code('  poppler-utils tesseract-ocr tesseract-ocr-ind tesseract-ocr-eng \\\'),
code('  python3 python3-pip && \\\'),
code('  pip3 install --break-system-packages curl-cffi beautifulsoup4 && \\\'),
code('  rm -rf /var/lib/apt/lists/*'),
code('COPY --from=builder /pipeline /usr/local/bin/pipeline'),
code('COPY --from=builder /app/backend/internal/prompts /prompts'),
code('COPY --from=builder /app/backend/python /python'),
body("Create a separate Dockerfile or docker-compose service for the Telegram bot (cmd/bot/main.go) with the same base image."),
body("Acceptance Criteria: (1) docker build produces a working image (2) Pipeline binary starts and connects to PostgreSQL and Qdrant (3) pdftotext, pdftoppm, and tesseract commands are available in the container (4) Python scraper bridge works (curl_cffi importable)."),

h2("11.2 Task 7.2: Remove Firestore Dependency from Legal Pipeline"),
bodyBold("Files: Multiple (MODIFY after validation)"),
body("Once the PostgreSQL migration is validated and all data is confirmed correct, remove the Firestore dependency from the legal pipeline code paths. Keep the Firebase Admin SDK import only for user/bot state if needed. Specifically: (1) Remove cloud.google.com/go/firestore from go.mod if no longer used. (2) Remove Firebase config from config.go if not needed for user state. (3) Update docker-compose.yml to remove any Firebase-related environment variables. (4) Remove the FirestoreRepo implementation or move it to a legacy/ directory."),
body("Acceptance Criteria: (1) go build ./... passes without Firestore import (2) STORAGE_MODE env var can be removed (PostgreSQL is the only backend) (3) All pipeline operations work with PostgreSQL only."),

h2("11.3 Task 7.3: Integration Testing"),
bodyBold("File: backend/internal/repository/postgres_test.go (NEW)"),
body("Create integration tests that verify the PostgreSQL repository against a real database (using testcontainers-go or a docker-compose test fixture). Tests must cover: (1) Full pipeline lifecycle: discover -> download -> parse -> analyze -> no_conflict (stopping) and discover -> ... -> published (full path). (2) Pause-and-resume: save draft with pending_prompt_approval, fetch it in a separate test, simulate approval callback. (3) Vector search: upsert embeddings to Qdrant, search and verify results. (4) Mock embedding handling: verify is_mock=true embeddings are excluded from search. (5) Concurrent access: multiple goroutines writing to the same law document (simulating the scheduler + Telegram callback race)."),
body("Acceptance Criteria: (1) All integration tests pass (2) Test coverage for all 8 interfaces (3) Race condition tests pass with -race flag."),

// ═══════════════════════════════════════════════════════
// SECTION 12: KNOWN ISSUES & BLOCKERS
// ═══════════════════════════════════════════════════════
h1("12. Known Issues and Blockers"),

h2("12.1 Instagram Publishing is Broken"),
body("The current HandleApprovalAction method in engine.go (line ~438) has a placeholder for the public image URL: publicImageURL := e.cfg.Instagram.AccessToken. This means Instagram publishing cannot work without an image hosting service (S3, GCS, or Imgur). The publishing job is saved with status pending_image_hosting and the pipeline stops. This is NOT a PostgreSQL migration issue, but it must be addressed before the system can fully function. The recommended fix is to add an image upload step in the generateAndApprove method, after image generation and validation, that uploads the generated image to S3/GCS and stores the public URL in the ImageAsset record."),

h2("12.2 TikTok Has Zero Implementation"),
body("The original plan listed TikTok as a publish target. The codebase has zero TikTok implementation. The PublishingJob.Platform field has a CHECK constraint for only 'instagram'. Do not add TikTok support in this migration. If TikTok support is desired, it should be a separate feature effort with its own connector implementation, content format adaptation (video vs image), and API integration."),

h2("12.3 GetLawAnalysisByDraft Inefficiency"),
body("The current implementation in firestore.go (line ~150-173) uses a CollectionGroup query to scan ALL analyses across ALL laws, then filters in-memory by matching the analysis document ID to the draft's law_analysis_id. This is O(n) where n is the total number of analyses. The PostgreSQL implementation fixes this with a JOIN query, which is O(1) with proper indexing. This is a significant performance improvement that comes for free with the migration."),

h2("12.4 Connector Architecture Clarification"),
body("The original plan implied many separate connectors per document type. In reality, the peraturan connector (backend/internal/connectors/peraturan/peraturan.go) handles 6 document types (UU, PP, Perpres, Permen, Kepmen, PBI) in a single connector. The connector priority list should reflect this: peraturan.go.id (6 types), JDIHN, JDIH Setneg, JDIH BPK, Mahkamah Konstitusi, JDIH Kemenkeu, Mahkamah Agung, JDIH Kemnaker, JDIH Kemendag, JDIH Komdigi, JDIH KPU, JDIH BKN, JDIH LKPP, JDIH DPR RI. That is 15 registered connectors, not 20+ separate ones."),

// ═══════════════════════════════════════════════════════
// SECTION 13: TASK DEPENDENCY & EXECUTION ORDER
// ═══════════════════════════════════════════════════════
h1("13. Task Dependency and Execution Order"),
body("The following table summarizes all tasks, their dependencies, and execution order. Tasks in the same phase can be parallelized unless one depends on another within the same phase."),

makeTable(
  ["Task ID", "Task", "Depends On", "Phase"],
  [
    ["0.1", "Define repository interfaces", "None", "0"],
    ["0.2", "FirestoreRepo implements interfaces", "0.1", "0"],
    ["0.3", "Refactor Engine to interfaces", "0.1", "0"],
    ["0.4", "Update retrieval.Service", "0.1", "0"],
    ["0.5", "Update all cmd/ files", "0.2, 0.3, 0.4", "0"],
    ["1.1", "Add PostgreSQL + Qdrant deps", "None (can run parallel with Phase 0)", "1"],
    ["1.2", "Create migration files", "1.1", "1"],
    ["1.3", "Add PostgreSQL + Qdrant config", "None", "1"],
    ["2.1", "Create Qdrant client wrapper", "1.1", "2"],
    ["2.2", "Docker Compose for Qdrant + PG", "None", "2"],
    ["3.1", "Create PostgresRepo struct", "1.2", "3"],
    ["3.2", "Implement LawDocument methods", "3.1", "3"],
    ["3.3", "Implement all remaining methods", "3.2", "3"],
    ["4.1", "Update retrieval.Service for Qdrant", "2.1, 3.1", "4"],
    ["4.2", "Upsert to Qdrant on embed", "2.1", "4"],
    ["5.1", "Create repository factory", "3.3, 4.1", "5"],
    ["5.2", "Update cmd/ files for factory", "5.1", "5"],
    ["6.1", "Create Firestore-to-PG migration tool", "5.2", "6"],
    ["6.2", "Update local queue flush", "5.2", "6"],
    ["7.1", "Create Dockerfile", "2.2, 5.2", "7"],
    ["7.2", "Remove Firestore dependency", "6.1", "7"],
    ["7.3", "Integration testing", "5.2, 2.2", "7"],
  ]
),

]; // end bodyChildren

// ── Assemble document ──
const doc = new Document({
  styles: { default: { document: {
    run: { font: { ascii: "Calibri" }, size: 22, color: c(P.body) },
    paragraph: { spacing: { line: 312 } },
  }}},
  sections: [
    // Section 1: Cover (no page numbers)
    { properties: { page: { margin: { top: 0, bottom: 0, left: 0, right: 0 }, size: { width: 11906, height: 16838 } } },
      children: coverChildren,
    },
    // Section 2: TOC (Roman numerals)
    { properties: { page: { margin: { top: 1440, bottom: 1440, left: 1701, right: 1417 }, pageNumbers: { start: 1, formatType: "upperRoman" } } },
      footers: { default: new Footer({ children: [new Paragraph({ alignment: AlignmentType.CENTER, children: [new TextRun({ children: ["PAGE \\\\* ROMAN \\\\* MERGEFORMAT"], size: 18, color: c(P.secondary) })] })] }) },
      children: tocChildren,
    },
    // Section 3: Body (Arabic, reset to 1)
    { properties: { page: { margin: { top: 1440, bottom: 1440, left: 1701, right: 1417 }, pageNumbers: { start: 1, formatType: "decimal" } } },
      headers: { default: new Header({ children: [new Paragraph({ alignment: AlignmentType.RIGHT, spacing: { after: 0 }, children: [new TextRun({ text: "PostgreSQL + Qdrant Integration Plan", size: 16, color: c(P.secondary), font: { ascii: "Calibri" }, italics: true })] })] }) },
      footers: { default: new Footer({ children: [new Paragraph({ alignment: AlignmentType.CENTER, children: [new TextRun({ children: ["PAGE \\\\* arabic \\\\* MERGEFORMAT"], size: 18, color: c(P.secondary) })] })] }) },
      children: bodyChildren,
    },
  ],
});

const outPath = "/home/z/my-project/download/Hukum-Aneh_PostgreSQL_Qdrant_Integration_Plan.docx";
Packer.toBuffer(doc).then(buf => {
  fs.writeFileSync(outPath, buf);
  console.log("Written to: " + outPath);
});
