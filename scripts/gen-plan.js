const fs = require("fs");
const {
  Document, Packer, Paragraph, TextRun, Table, TableRow, TableCell,
  Header, Footer, PageNumber, NumberFormat, AlignmentType,
  HeadingLevel, WidthType, BorderStyle, ShadingType, TableLayoutType,
  TableOfContents, PageBreak,
} = require("docx");

// Palette: GO-1 (Graphite Orange) with R1 cover
const PAL = {
  bg: "1A2330", accent: "D4875A",
  titleColor: "FFFFFF", subtitleColor: "B0B8C0",
  metaColor: "90989F", footerColor: "687078",
  table: { headerBg: "D4875A", headerText: "FFFFFF", accentLine: "D4875A", innerLine: "DDD0C8", surface: "F8F0EB" },
};

const NB = { style: BorderStyle.NONE, size: 0, color: "FFFFFF" };
const noBorders = { top: NB, bottom: NB, left: NB, right: NB };
const allNoBorders = { top: NB, bottom: NB, left: NB, right: NB, insideHorizontal: NB, insideVertical: NB };

// calcTitleLayout for English
function calcTitleLayout(title, maxWidthTwips, preferredPt = 40, minPt = 24) {
  const charWidth = (pt) => pt * 11;
  const charsPerLine = (pt) => Math.floor(maxWidthTwips / charWidth(pt));
  let titlePt = preferredPt;
  let lines;
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
  const breakAfter = new Set([' ', '-', '_', '/', ':', '\u2014']);
  const lines = [];
  let remaining = title;
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
  return lines;
}

function calcCoverSpacing(params) {
  const { titleLineCount = 1, titlePt = 36, hasSubtitle = false, hasEnglishLabel = false, metaLineCount = 0, fixedHeight = 400, pageHeight = 16838, marginTop = 0, marginBottom = 0 } = params;
  const SAFETY = 1200;
  const usableHeight = pageHeight - marginTop - marginBottom - SAFETY;
  const titleHeight = titleLineCount * (titlePt * 23 + 200);
  const subtitleHeight = hasSubtitle ? (12 * 23 + 600) : 0;
  const englishLabelHeight = hasEnglishLabel ? (9 * 23 + 600) : 0;
  const metaHeight = metaLineCount * (10 * 23 + 100);
  const implicitParaHeight = 3 * 300;
  const contentHeight = titleHeight + subtitleHeight + englishLabelHeight + metaHeight + fixedHeight + implicitParaHeight;
  const remainingSpace = Math.max(usableHeight - contentHeight, 400);
  const FOOTER_MIN = 800;
  const rawTop = Math.floor(remainingSpace * 0.45);
  const rawBottom = Math.floor(remainingSpace * 0.45);
  const bottomSpacing = Math.max(rawBottom, FOOTER_MIN);
  return { topSpacing: rawTop, midSpacing: Math.floor(remainingSpace * 0.1), bottomSpacing };
}

// Cover R1
function buildCoverR1(config) {
  const P = config.palette;
  const padL = 1200, padR = 800;
  const availableWidth = 11906 - padL - padR - 300;
  const { titlePt, titleLines } = calcTitleLayout(config.title, availableWidth, 38, 24);
  const titleSize = titlePt * 2;
  const spacing = calcCoverSpacing({
    titleLineCount: titleLines.length, titlePt,
    hasSubtitle: !!config.subtitle, hasEnglishLabel: !!config.englishLabel,
    metaLineCount: (config.metaLines || []).length, fixedHeight: 400,
  });
  const accentLeft = { style: BorderStyle.SINGLE, size: 8, color: P.accent, space: 12 };
  const children = [];
  children.push(new Paragraph({ spacing: { before: spacing.topSpacing } }));
  if (config.englishLabel) {
    children.push(new Paragraph({
      indent: { left: padL, right: padR }, spacing: { after: 500 },
      border: { bottom: { style: BorderStyle.SINGLE, size: 6, color: P.accent, space: 8 } },
      children: [new TextRun({ text: config.englishLabel.split("").join("  "), size: 18, color: P.accent, font: { ascii: "Calibri" }, characterSpacing: 40 })],
    }));
  }
  for (let i = 0; i < titleLines.length; i++) {
    children.push(new Paragraph({
      indent: { left: padL },
      spacing: { after: i < titleLines.length - 1 ? 100 : 300, line: Math.ceil(titlePt * 23), lineRule: "atLeast" },
      children: [new TextRun({ text: titleLines[i], size: titleSize, bold: true, color: P.titleColor, font: { ascii: "Arial" } })],
    }));
  }
  if (config.subtitle) {
    children.push(new Paragraph({
      indent: { left: padL }, spacing: { after: 800 },
      children: [new TextRun({ text: config.subtitle, size: 24, color: P.subtitleColor, font: { ascii: "Arial" } })],
    }));
  }
  for (const line of (config.metaLines || [])) {
    children.push(new Paragraph({
      indent: { left: padL + 200 }, spacing: { after: 80 }, border: { left: accentLeft },
      children: [new TextRun({ text: line, size: 24, color: P.metaColor, font: { ascii: "Arial" } })],
    }));
  }
  children.push(new Paragraph({ spacing: { before: spacing.bottomSpacing } }));
  children.push(new Paragraph({
    indent: { left: padL, right: padR },
    border: { top: { style: BorderStyle.SINGLE, size: 2, color: P.accent, space: 8 } },
    spacing: { before: 200 },
    children: [
      new TextRun({ text: config.footerLeft || "", size: 16, color: P.footerColor, font: { ascii: "Arial" } }),
      new TextRun({ text: "                                        " }),
      new TextRun({ text: config.footerRight || "", size: 16, color: P.footerColor, font: { ascii: "Arial" } }),
    ],
  }));
  return [new Table({
    width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED,
    borders: allNoBorders,
    rows: [new TableRow({ height: { value: 16838, rule: "exact" }, children: [
      new TableCell({ shading: { type: ShadingType.CLEAR, fill: P.bg }, borders: noBorders, verticalAlign: "top", children }),
    ]})],
  })];
}

// Body helpers
const FONT = { ascii: "Times New Roman" };
const FONT_H = { ascii: "Arial" };

function h1(text) {
  return new Paragraph({ heading: HeadingLevel.HEADING_1, spacing: { before: 360, after: 200 }, keepNext: true,
    children: [new TextRun({ text, bold: true, size: 32, font: FONT_H, color: "000000" })] });
}
function h2(text) {
  return new Paragraph({ heading: HeadingLevel.HEADING_2, spacing: { before: 280, after: 140 }, keepNext: true,
    children: [new TextRun({ text, bold: true, size: 28, font: FONT_H, color: "000000" })] });
}
function h3(text) {
  return new Paragraph({ heading: HeadingLevel.HEADING_3, spacing: { before: 200, after: 100 }, keepNext: true,
    children: [new TextRun({ text, bold: true, size: 24, font: FONT_H, color: "000000" })] });
}
function p(text) {
  return new Paragraph({ spacing: { after: 120, line: 312 }, alignment: AlignmentType.LEFT,
    children: [new TextRun({ text, size: 24, font: FONT, color: "000000" })] });
}
function boldP(label, text) {
  return new Paragraph({ spacing: { after: 120, line: 312 }, alignment: AlignmentType.LEFT,
    children: [
      new TextRun({ text: label, bold: true, size: 24, font: FONT, color: "000000" }),
      new TextRun({ text, size: 24, font: FONT, color: "000000" }),
    ] });
}
function bullet(text) {
  return new Paragraph({ spacing: { after: 60, line: 312 }, indent: { left: 480 }, alignment: AlignmentType.LEFT,
    children: [new TextRun({ text: "\u2022 " + text, size: 24, font: FONT, color: "000000" })] });
}
function codeBlock(lines) {
  return lines.map(line => new Paragraph({ spacing: { after: 0, line: 280 }, indent: { left: 480 }, alignment: AlignmentType.LEFT,
    shading: { type: ShadingType.CLEAR, fill: "F5F5F5" },
    children: [new TextRun({ text: line, size: 20, font: { ascii: "Courier New" }, color: "333333" })] }));
}

function makeTable(headers, rows) {
  const colCount = headers.length;
  const colW = Math.floor(100 / colCount);
  const thinBorder = { style: BorderStyle.SINGLE, size: 1, color: PAL.table.innerLine };
  const headerRow = new TableRow({ tableHeader: true, cantSplit: true, children:
    headers.map(h => new TableCell({ width: { size: colW, type: WidthType.PERCENTAGE },
      shading: { type: ShadingType.CLEAR, fill: PAL.table.headerBg },
      borders: { top: { style: BorderStyle.SINGLE, size: 1, color: PAL.table.accentLine }, bottom: { style: BorderStyle.SINGLE, size: 1, color: PAL.table.accentLine }, left: NB, right: NB },
      margins: { top: 60, bottom: 60, left: 100, right: 100 },
      children: [new Paragraph({ alignment: AlignmentType.LEFT, children: [new TextRun({ text: h, bold: true, size: 20, color: PAL.table.headerText, font: FONT_H })] })],
    }))
  });
  const dataRows = rows.map((row, ri) => new TableRow({ cantSplit: true, children:
    row.map(cell => new TableCell({ width: { size: colW, type: WidthType.PERCENTAGE },
      shading: ri % 2 === 1 ? { type: ShadingType.CLEAR, fill: PAL.table.surface } : undefined,
      borders: { top: thinBorder, bottom: thinBorder, left: NB, right: NB },
      margins: { top: 50, bottom: 50, left: 100, right: 100 },
      children: [new Paragraph({ alignment: AlignmentType.LEFT, children: [new TextRun({ text: cell, size: 20, font: FONT, color: "000000" })] })],
    }))
  }));
  return new Table({ width: { size: 100, type: WidthType.PERCENTAGE }, layout: TableLayoutType.FIXED, rows: [headerRow, ...dataRows] });
}

// ─── CONTENT ───
const coverChildren = buildCoverR1({
  title: "PostgreSQL Integration Plan",
  subtitle: "Indonesian Legal AI \u2014 hukum-aneh",
  englishLabel: "EXECUTION PLAN",
  metaLines: ["Project: hukum-aneh", "Status: Planning", "Date: July 26, 2026"],
  footerLeft: "Version 1.0",
  footerRight: "Confidential",
  palette: PAL,
});

const body = [];

// 1. Executive Summary
body.push(h1("1. Executive Summary"));
body.push(p("This document defines the complete execution plan for integrating PostgreSQL as the primary legal document database for the hukum-aneh project. The project is a Go-based pipeline that scrapes Indonesian legal documents, analyzes them for unusual content using AI, generates social media content, and publishes to Instagram via Telegram-based human approval gates. The current system stores all data in Google Cloud Firestore through a monolithic FirestoreRepo struct with no interface abstraction."));
body.push(p("PostgreSQL will become the source of truth for all legal corpus data including laws, parsed versions, analyses, content drafts, image assets, approval records, publishing jobs, and embedding mappings. Firebase will be retained exclusively for user-facing application features such as authentication, user profiles, chat history, bookmarks, and analytics. The vector search component will migrate from the current brute-force cosine similarity (which loads all embeddings into memory) to Qdrant, with PostgreSQL storing only the article-to-vector-ID mapping."));
body.push(p("The plan is structured into 11 phases, each containing detailed subtasks with specific file paths, code changes, and acceptance criteria. The migration is designed to be incremental: first extract repository interfaces, then implement PostgreSQL alongside Firestore, then switch connectors one by one, and finally remove Firestore dependency for legal data. No phase requires a big-bang cutover."));

// 2. Target Architecture
body.push(h1("2. Target Architecture"));
body.push(p("The architecture follows a clear separation of concerns. The frontend communicates with Firebase for authentication and user data. The backend API layer sits behind Firebase Auth and routes legal data queries to PostgreSQL and vector searches to Qdrant. Raw legal documents (PDFs) remain stored on Google Drive as they are today, with PostgreSQL holding metadata and parsed text content. The following table summarizes the responsibility split between the databases."));
body.push(makeTable(
  ["Layer", "Technology", "Responsibility"],
  [
    ["User Features", "Firestore", "Auth, profiles, chat, bookmarks, analytics"],
    ["Legal Corpus", "PostgreSQL", "Laws, versions, analyses, drafts, approvals"],
    ["Vector Search", "Qdrant", "1536-dim embeddings, semantic similarity"],
    ["File Storage", "Google Drive", "Raw PDFs, generated images (2 TB)"],
    ["Application", "Go Backend", "Pipeline engine, connectors, services"],
  ]
));
body.push(p("The critical architectural change is the introduction of repository interfaces between the Engine and the data layer. Currently, Engine directly references the concrete FirestoreRepo type. After migration, Engine will depend on interfaces (LawRepository, DraftRepository, etc.), and either a Firestore adapter or a PostgreSQL adapter can be injected. This allows incremental migration and makes testing possible without a live Firestore connection."));

// 3. Phase 1
body.push(h1("3. Phase 1 \u2014 Environment Setup"));
body.push(p("This phase establishes the local development environment with Docker Compose, including PostgreSQL, Qdrant, and the Go backend service. The goal is to have a reproducible, containerized development setup that any developer can spin up with a single command. This phase does not change any application code beyond adding new environment variables and the PostgreSQL driver dependency."));

body.push(h2("3.1 Task: Create docker-compose.yml"));
body.push(boldP("Files to create: ", "docker-compose.yml (project root)"));
body.push(boldP("Description: ", "Define three services: postgres, qdrant, and backend. The postgres service uses the official PostgreSQL 16 image with a dedicated database and user. The qdrant service uses the official Qdrant image for vector storage. The backend service builds from a Dockerfile and depends on both postgres and qdrant being healthy."));
body.push(boldP("Specific actions: ", ""));
body.push(bullet("Create docker-compose.yml at project root with version '3.8'"));
body.push(bullet("postgres service: image postgres:16-alpine, port 5432, env vars POSTGRES_DB=hukum_anneh, POSTGRES_USER=hukum, POSTGRES_PASSWORD=dev_password, volume pgdata:/var/lib/postgresql/data, healthcheck pg_isready"));
body.push(bullet("qdrant service: image qdrant/qdrant:latest, ports 6333 (HTTP) and 6334 (gRPC), volume qdrant_storage:/qdrant/storage"));
body.push(bullet("backend service: build from ./Dockerfile, port 8080, depends_on postgres (condition: service_healthy) and qdrant, env_file .env"));
body.push(bullet("Define named volumes pgdata and qdrant_storage"));
body.push(boldP("Acceptance: ", "docker-compose up starts all three services. postgres is accessible on localhost:5432. qdrant dashboard on localhost:6333. Backend connects to both without errors."));

body.push(h2("3.2 Task: Create Dockerfile"));
body.push(boldP("Files to create: ", "Dockerfile (project root)"));
body.push(boldP("Description: ", "The backend has hard dependencies on system tools: pdftotext and pdftoppm (from poppler-utils) for PDF text extraction, and tesseract with Indonesian and English language packs for OCR fallback. Additionally, 4 connectors (MKRI, DPR, JDIHN, LKPP) require Python 3 with curl_cffi and beautifulsoup4 for TLS-fingerprinted web scraping. The Dockerfile must install all of these system-level dependencies that the Go binary needs at runtime."));
body.push(boldP("Specific actions: ", ""));
body.push(bullet("Base image: golang:1.25-bookworm"));
body.push(bullet("Install system deps via apt-get: poppler-utils, tesseract-ocr, tesseract-ocr-ind, tesseract-ocr-eng, python3, python3-pip"));
body.push(bullet("Install Python packages: curl_cffi, beautifulsoup4 via pip3"));
body.push(bullet("Copy go.mod/go.sum, run go mod download"));
body.push(bullet("Copy full source, run go build ./backend/cmd/pipeline/"));
body.push(boldP("Acceptance: ", "docker-compose build backend completes. The built image can parse a PDF using pdftotext and fall back to tesseract OCR."));

body.push(h2("3.3 Task: Add PostgreSQL Environment Variables"));
body.push(boldP("Files to modify: ", ".env.example, backend/internal/config/config.go"));
body.push(boldP("Description: ", "Add new environment variables for PostgreSQL connection and Qdrant URL to .env.example. Update the manual .env parser in config.go to load these new variables. The config struct needs new fields for the PostgreSQL DSN and Qdrant endpoint. All existing env vars must remain unchanged to maintain backward compatibility during the transition period."));
body.push(boldP("Specific actions: ", ""));
body.push(bullet("Add to .env.example: DATABASE_URL=postgres://hukum:dev_password@localhost:5432/hukum_anneh?sslmode=disable"));
body.push(bullet("Add to .env.example: QDRANT_URL=http://localhost:6333"));
body.push(bullet("In config.go: add DatabaseURL string and QdrantURL string fields to Config struct"));
body.push(bullet("In the .env parser loop: read DATABASE_URL and QDRANT_URL with empty string defaults"));
body.push(boldP("Acceptance: ", "Config loads without errors when DATABASE_URL and QDRANT_URL are set. Existing config fields still work. No existing test breaks."));

body.push(h2("3.4 Task: Initialize Go PostgreSQL Dependency"));
body.push(boldP("Files to modify: ", "go.mod (via go get)"));
body.push(boldP("Description: ", "Add the pgx PostgreSQL driver. The recommended library is github.com/jackc/pgx/v5 with its connection pool (pgxpool). This is the most performant and feature-rich PostgreSQL driver for Go, supporting prepared statements, connection pooling, and automatic type scanning. It is the de facto standard for production Go-PostgreSQL applications."));
body.push(boldP("Specific actions: ", ""));
body.push(bullet("Run: go get github.com/jackc/pgx/v5 github.com/jackc/pgx/v5/pgxpool"));
body.push(bullet("Verify go.mod and go.sum are updated with the new dependency"));
body.push(bullet("Verify go build ./... still succeeds"));
body.push(boldP("Acceptance: ", "go build ./... succeeds with the new dependency. No import errors."));

// 4. Phase 2
body.push(h1("4. Phase 2 \u2014 Database Schema Design"));
body.push(p("This phase designs and implements the PostgreSQL schema. The schema must reflect the actual data model used by the codebase, not an idealized version. The current codebase has 10 Go structs across 6 top-level collections and 4 subcollections in Firestore. The schema must support the 13-state status lifecycle, multi-version parsed text (via TOAST), JSONB for nested structures like AffectedLaws, and proper indexing for the query patterns used by the engine."));

body.push(h2("4.1 Task: Create Migration Framework"));
body.push(boldP("Files to create: ", "backend/migrations/001_initial_schema.up.sql, backend/migrations/001_initial_schema.down.sql, backend/cmd/migrate/main.go"));
body.push(boldP("Description: ", "Establish a migration framework using plain SQL files with a Go migration runner. Avoid heavy migration libraries; a simple Go script that reads .sql files from a directory and executes them in order, tracking applied migrations in a schema_migrations table, is sufficient for this project's needs. Each migration has an up file and a down file for rollback support."));
body.push(boldP("Specific actions: ", ""));
body.push(bullet("Create backend/migrations/ directory for SQL files"));
body.push(bullet("Write cmd/migrate/main.go: connect via pgxpool, create schema_migrations table (version TEXT PRIMARY KEY, applied_at TIMESTAMPTZ), read .up.sql files in sorted order, execute those not yet applied, support --down flag to reverse last migration"));
body.push(bullet("Migration 001 contains all CREATE TABLE statements (Tasks 4.2 through 4.7)"));
body.push(boldP("Acceptance: ", "Running go run backend/cmd/migrate/ applies all pending migrations. Running again is a no-op (idempotent). Running with --down reverses the last migration."));

body.push(h2("4.2 Task: Create laws Table"));
body.push(boldP("Files to modify: ", "backend/migrations/001_initial_schema.up.sql"));
body.push(boldP("Description: ", "The laws table maps directly to the LawDocument struct in models.go. It is the core table that the engine's discovery and processing pipeline reads and writes. The status field must support the full 13-state lifecycle defined in the engine. The published_date field remains TEXT (not a parsed date) to match the existing string format in the codebase and avoid data transformation issues during migration."));
body.push(boldP("Schema: ", ""));
body.push(...codeBlock([
  "CREATE TABLE laws (",
  "  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),",
  "  law_number      TEXT NOT NULL,",
  "  title           TEXT NOT NULL,",
  "  source_url      TEXT,",
  "  source          TEXT NOT NULL,",
  "  level           TEXT NOT NULL DEFAULT 'national',",
  "  document_type   TEXT NOT NULL,",
  "  raw_file_path   TEXT,",
  "  published_date  TEXT,",
  "  status          TEXT NOT NULL DEFAULT 'discovered',",
  "  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),",
  "  updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()",
  ");",
])));
body.push(bullet("UNIQUE INDEX idx_laws_law_number ON laws(law_number)"));
body.push(bullet("INDEX idx_laws_status ON laws(status)"));
body.push(bullet("INDEX idx_laws_document_type ON laws(document_type)"));
body.push(bullet("INDEX idx_laws_status_updated ON laws(status, updated_at) -- for FindStuckDocuments"));
body.push(bullet("CHECK constraint on status for the 13 valid states"));
body.push(boldP("Acceptance: ", "Table created with all fields matching LawDocument struct. Unique index on law_number prevents duplicates. Status index supports ListLawsByStatus and FindStuckDocuments queries."));

body.push(h2("4.3 Task: Create law_versions Table"));
body.push(boldP("Files to modify: ", "backend/migrations/001_initial_schema.up.sql"));
body.push(boldP("Description: ", "Maps to the LawVersion struct and the laws/{lawId}/versions Firestore subcollection. A single law can have multiple versions (re-parses). The text_content field can exceed 1MB for large Indonesian regulations. PostgreSQL handles this automatically via TOAST (The Oversized-Attribute Storage Technique), which transparently compresses and stores large field values in a separate TOAST table. No special handling is needed in application code."));
body.push(boldP("Schema: ", ""));
body.push(...codeBlock([
  "CREATE TABLE law_versions (",
  "  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),",
  "  law_id          UUID NOT NULL REFERENCES laws(id) ON DELETE CASCADE,",
  "  version_number  INTEGER NOT NULL,",
  "  text_content    TEXT NOT NULL,",
  "  embedding_id    TEXT,",
  "  parsed_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()",
  ");",
  "CREATE INDEX idx_versions_law_version ON law_versions(law_id, version_number DESC);",
])));
body.push(boldP("Acceptance: ", "Foreign key to laws with CASCADE delete supports the Firestore subcollection deletion pattern. Composite index supports GetLatestLawVersion query (ORDER BY version_number DESC LIMIT 1)."));

body.push(h2("4.4 Task: Create law_relationships Table"));
body.push(boldP("Files to modify: ", "backend/migrations/001_initial_schema.up.sql"));
body.push(boldP("Description: ", "Maps to LawRelationship struct and the laws/{lawId}/relationships subcollection. Stores citation links between regulations. The relationship_type field uses the same values as the codebase: amends, repeals, supersedes, references. The article_ref field stores optional article-level references (e.g., 'Pasal 33')."));
body.push(...codeBlock([
  "CREATE TABLE law_relationships (",
  "  id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),",
  "  law_id             UUID NOT NULL REFERENCES laws(id) ON DELETE CASCADE,",
  "  related_law_number TEXT NOT NULL,",
  "  relationship_type  TEXT NOT NULL,",
  "  article_ref        TEXT,",
  "  created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW()",
  ");",
  "CREATE INDEX idx_relationships_law_id ON law_relationships(law_id);",
  "CREATE INDEX idx_relationships_related ON law_relationships(related_law_number);",
])));
body.push(boldP("Acceptance: ", "Table supports both forward lookups (all relationships for a law) and reverse lookups (all laws that reference a given law number)."));

body.push(h2("4.5 Task: Create law_analyses Table"));
body.push(boldP("Files to modify: ", "backend/migrations/001_initial_schema.up.sql"));
body.push(boldP("Description: ", "Maps to the LawAnalysis struct and the laws/{lawId}/analyses subcollection. Contains AI-generated analysis with 4 score dimensions (overall, controversy, economic, legal_consistency, all 0-100) and a nested array of affected laws. The affected_laws array is stored as JSONB for efficient querying using PostgreSQL's JSONB containment operators. The raw_json field preserves the original LLM JSON response for debugging."));
body.push(...codeBlock([
  "CREATE TABLE law_analyses (",
  "  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),",
  "  law_id            UUID NOT NULL REFERENCES laws(id) ON DELETE CASCADE,",
  "  summary           TEXT,",
  "  affected_laws     JSONB DEFAULT '[]',",
  "  overall_score     INTEGER NOT NULL DEFAULT 0,",
  "  controversy_score INTEGER NOT NULL DEFAULT 0,",
  "  economic_score    INTEGER NOT NULL DEFAULT 0,",
  "  legal_consistency INTEGER NOT NULL DEFAULT 0,",
  "  confidence        DOUBLE PRECISION NOT NULL DEFAULT 0.0,",
  "  raw_json          TEXT,",
  "  created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()",
  ");",
  "CREATE INDEX idx_analyses_law_id ON law_analyses(law_id);",
  "CREATE INDEX idx_analyses_scores ON law_analyses(overall_score, controversy_score);",
])));
body.push(boldP("Acceptance: ", "JSONB column supports containment queries like WHERE affected_laws @> '[{"severity": 0.7}]'. Score indexes support sorting by oddness level."));

body.push(h2("4.6 Task: Create Content Pipeline Tables"));
body.push(boldP("Files to modify: ", "backend/migrations/001_initial_schema.up.sql"));
body.push(boldP("Description: ", "Four tables that power the social media content pipeline: content_drafts, captions, image_assets, and approvals. Additionally, publishing_jobs tracks social media publish status. These map directly to the ContentDraft, Caption, ImageAsset, Approval, and PublishingJob structs in models.go. The approval table records the full audit trail of the 6 decision types used in the Telegram approval flow."));
body.push(boldP("Tables: content_drafts, captions, image_assets, approvals, publishing_jobs. ", "Each table maps 1:1 to the corresponding Go struct. Key design decisions: hashtags stored as TEXT[] (PostgreSQL array), content_drafts.law_analysis_id is a foreign key to law_analyses (replacing the expensive collection group query in Firestore), all tables use UUID primary keys, and CASCADE deletes propagate from content_drafts down to captions and image_assets."));
body.push(boldP("Acceptance: ", "All 5 tables created with correct foreign keys, indexes, and defaults. CASCADE deletes work from laws through the full chain. The JOIN between content_drafts and law_analyses replaces the Firestore collection group scan."));

body.push(h2("4.7 Task: Create Embedding Mapping Table"));
body.push(boldP("Files to modify: ", "backend/migrations/001_initial_schema.up.sql"));
body.push(boldP("Description: ", "This table stores ONLY the mapping between a law document and its vector ID in Qdrant. The actual 1536-dimensional float vectors are stored in Qdrant, not in PostgreSQL. This is a critical design decision: the current codebase stores full vectors in Firestore EmbeddingEntry.Vector field, and the brute-force Search() in retrieval.go loads ALL of them into memory. Storing vectors in PostgreSQL would recreate this exact problem. The is_mock column is preserved to allow filtering out mock/fallback embeddings from search results."));
body.push(...codeBlock([
  "CREATE TABLE embeddings (",
  "  id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),",
  "  law_id     UUID NOT NULL REFERENCES laws(id) ON DELETE CASCADE,",
  "  vector_id  TEXT NOT NULL,",
  "  is_mock    BOOLEAN NOT NULL DEFAULT FALSE,",
  "  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()",
  ");",
  "CREATE UNIQUE INDEX idx_embeddings_vector_id ON embeddings(vector_id);",
])));
body.push(boldP("Acceptance: ", "Table stores mapping only, no vector data. Unique index on vector_id prevents duplicates. is_mock filter allows excluding fallback data from search results."));

body.push(h2("4.8 Task: Add Full-Text Search Indexes"));
body.push(boldP("Files to create: ", "backend/migrations/002_fulltext_search.up.sql"));
body.push(boldP("Description: ", "Add GIN indexes for PostgreSQL full-text search on the most commonly queried text fields. This enables natural language queries like 'UU 11 Tahun 2020' or 'Pasal 33' without requiring vector search. The tsvector columns are generated columns that automatically extract searchable tokens from the source text columns using PostgreSQL's built-in text search configuration."));
body.push(boldP("Specific actions: ", ""));
body.push(bullet("ALTER TABLE laws ADD COLUMN search_vector tsvector GENERATED ALWAYS AS (to_tsvector('english', coalesce(title,'') || ' ' || coalesce(law_number,'') || ' ' || coalesce(document_type,''))) STORED"));
body.push(bullet("CREATE INDEX idx_laws_search ON laws USING GIN(search_vector)"));
body.push(bullet("ALTER TABLE law_versions ADD COLUMN text_search_vector tsvector GENERATED ALWAYS AS (to_tsvector('english', text_content)) STORED"));
body.push(bullet("CREATE INDEX idx_versions_text_search ON law_versions USING GIN(text_search_vector)"));
body.push(boldP("Acceptance: ", "SELECT * FROM laws WHERE search_vector @@ plainto_tsquery('english', 'Undang-Undang 11 2020') returns matching laws. EXPLAIN ANALYZE confirms GIN index usage."));

// 5. Phase 3
body.push(h1("5. Phase 3 \u2014 Repository Interface Extraction"));
body.push(p("This is the most critical phase of the entire migration. Currently, Engine and all 9 cmd/ entry points directly depend on the concrete FirestoreRepo type (defined as a struct field in engine.go line 33). Without interface extraction, it is impossible to swap database implementations incrementally. This phase defines interfaces for each domain, verifies that FirestoreRepo satisfies them, and refactors Engine to depend on interfaces instead of the concrete type. No PostgreSQL code is written in this phase."));

body.push(h2("5.1 Task: Define Repository Interfaces"));
body.push(boldP("Files to create: ", "backend/internal/repository/interfaces.go"));
body.push(boldP("Description: ", "Extract all repository methods from the monolithic FirestoreRepo into domain-specific interfaces. Based on the actual method signatures in firestore.go (378 lines), three sub-interfaces cover all operations, plus a composite Repository interface that combines them. The Engine's NewEngine function currently accepts *repository.FirestoreRepo; after this change, it will accept repository.Repository (the composite interface)."));
body.push(boldP("Interface definitions: ", ""));
body.push(...codeBlock([
  "// LawRepository: laws, versions, analyses",
  "type LawRepository interface {",
  "    SaveLawDocument(ctx context.Context, doc *models.LawDocument) (string, error)",
  "    GetLawDocument(ctx context.Context, id string) (*models.LawDocument, error)",
  "    FindByLawNumber(ctx context.Context, lawNumber string) (*models.LawDocument, error)",
  "    ListLawsByStatus(ctx context.Context, status string) ([]models.LawDocument, error)",
  "    ListAllLaws(ctx context.Context) ([]models.LawDocument, error)",
  "    FindStuckDocuments(ctx context.Context, status string, before time.Time) ([]models.LawDocument, error)",
  "    SaveLawVersion(ctx context.Context, lawID string, v *models.LawVersion) (string, error)",
  "    GetLatestLawVersion(ctx context.Context, lawID string) (*models.LawVersion, error)",
  "    SaveLawAnalysis(ctx context.Context, lawID string, a *models.LawAnalysis) (string, error)",
  "    GetLawAnalysisByDraft(ctx context.Context, draftID string) (*models.LawAnalysis, error)",
  "}",
  "",
  "// DraftRepository: content drafts, assets, approvals, publishing",
  "type DraftRepository interface {",
  "    GetContentDraft(ctx context.Context, id string) (*models.ContentDraft, error)",
  "    SaveContentDraft(ctx context.Context, draft *models.ContentDraft) (string, error)",
  "    GetImageAssetsByDraft(ctx context.Context, draftID string) ([]models.ImageAsset, error)",
  "    SaveImageAsset(ctx context.Context, asset *models.ImageAsset) (string, error)",
  "    SaveApproval(ctx context.Context, a *models.Approval) (string, error)",
  "    SavePublishingJob(ctx context.Context, j *models.PublishingJob) (string, error)",
  "}",
  "",
  "// EmbeddingRepository: embedding mapping only (vectors in Qdrant)",
  "type EmbeddingRepository interface {",
  "    SaveEmbedding(ctx context.Context, emb *models.EmbeddingEntry) (string, error)",
  "    ListAllEmbeddings(ctx context.Context) ([]models.EmbeddingEntry, error)",
  "}",
  "",
  "// Repository: composite interface for Engine injection",
  "type Repository interface {",
  "    LawRepository",
  