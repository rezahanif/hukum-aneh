# Hukum Aneh — Indonesian Law Content Pipeline

Automated pipeline: discovers, parses, and analyzes Indonesian laws, then turns analysis into social-media-ready content.

## Stack

- **Go** — workflow engine, connectors, services, scheduler
- **Python** — TLS-fingerprinted scrapers (subprocess, invoked by Go)
- **Firebase Firestore** — primary database (NoSQL)
- **OpenAI via 9router** — LLM + embeddings
- **Instagram Meta Graph API** — publishing

## Quick start

```bash
# Set required env vars
export FIREBASE_PROJECT_ID=your-project
export FIREBASE_CREDENTIALS_PATH=path/to/service-account.json
export ROUTER9_BASE_URL=http://localhost:4000/v1
export ROUTER9_API_KEY=your-key
export TELEGRAM_BOT_TOKEN=your-token
export TELEGRAM_CHAT_ID=your-chat-id

# Run discovery once
go run ./backend/cmd/pipeline -once

# Run scheduler (default: hourly discovery)
go run ./backend/cmd/pipeline
```

## Architecture

See `ENGINEERING_SPEC.md` for the authoritative build spec.

```
backend/
  cmd/pipeline/          # entrypoint
  internal/
    workflow/            # orchestration engine
    connectors/          # gov source interfaces (peraturan, jdihn, bpk, ...)
    parser/              # PDF/OCR → text/markdown
    retrieval/           # similarity search (embeddings)
    ai/                  # Hermes agent wrappers (analysis, strategy, prompt builder)
    models/              # data models (Firestore collections)
    repository/          # Firestore CRUD
    services/
      imagegen/          # image generation
      telegram/          # approval bot
      publishing/        # Instagram
      analytics/         # post metrics
    scheduler/           # interval triggers
    validator/           # image QA gate
    config/              # env-based config
    prompts/             # design_guide.json, character sheets, templates
    storage/             # downloaded files, generated images
  pkg/scraper/           # Go → Python subprocess bridge
  python/scraper/        # Python TLS-fingerprinted scraper
  configs/               # sources.json, firebase creds
```

## Pipeline

```
Scheduler → Discovery → Download → Parse → Similarity Search
  → Law Analysis (AI) → Content Strategy (AI) → Prompt Builder (AI)
  → Image Gen → Image Validator → Telegram Approval → Publish → Analytics
```

All AI calls return structured JSON only. No free text downstream.
