# Dockerfile for hukum-aneh
# Multi-stage build:
#   1. go-builder: builds all cmd/ binaries with CGO enabled (gosseract requires it)
#   2. runtime: debian-slim + system deps (poppler, tesseract, python) + binaries
#
# Build: docker build -t hukum-aneh .
# Run pipeline: docker run --rm hukum-aneh pipeline
# Run bot:      docker run --rm hukum-aneh bot

# ============================================================================
# Stage 1: Go builder
# ============================================================================
FROM golang:1.25-bookworm AS go-builder

# CGO deps for gosseract (OCR in parser package)
RUN apt-get update && apt-get install -y --no-install-recommends \
    libtesseract-dev \
    libleptonica-dev \
    clang \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Cache deps before copying source
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build all cmd binaries
# CGO_ENABLED=1 required for gosseract
# -ldflags="-s -w" strips debug info for smaller images
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/pipeline     ./backend/cmd/pipeline && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/bot          ./backend/cmd/bot && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/backfill     ./backend/cmd/backfill && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/batch        ./backend/cmd/batch && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/backfill_emb ./backend/cmd/backfill_embeddings && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/verify_emb   ./backend/cmd/verify_embeddings && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/flush_local  ./backend/cmd/flush_local && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/count        ./backend/cmd/count && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/parse_pdf    ./backend/cmd/parse_pdf && \
    CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/backfill_bpk ./backend/cmd/backfill_bpk

# NOTE: migrate_to_pg binary is NOT included yet.
# Stream A will create backend/cmd/migrate_to_pg in Phase 6.1.
# After Stream A completes Phase 6.1, add the following line to this Dockerfile:
#   CGO_ENABLED=1 go build -ldflags="-s -w" -o /bin/migrate_to_pg ./backend/cmd/migrate_to_pg
#   COPY --from=go-builder /bin/migrate_to_pg /usr/local/bin/migrate_to_pg

# ============================================================================
# Stage 2: Runtime
# ============================================================================
FROM debian:bookworm-slim

# System deps:
#   - poppler-utils: pdftotext, pdftoppm (PDF parsing in parser package)
#   - tesseract-ocr + ind + eng: OCR (used by gosseract via CGO)
#   - libtesseract5, libleptonica6: shared libs for gosseract at runtime
#   - python3 + pip: scraper bridge (curl_cffi for BPK TLS bypass)
#   - ca-certificates: HTTPS
RUN apt-get update && apt-get install -y --no-install-recommends \
    poppler-utils \
    tesseract-ocr \
    tesseract-ocr-ind \
    tesseract-ocr-eng \
    libtesseract5 \
    libleptonica6 \
    python3 \
    python3-pip \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Python deps for scraper bridge
# --break-system-packages needed for Debian 12+ PEP 668
RUN pip3 install --no-cache-dir --break-system-packages \
    curl_cffi \
    beautifulsoup4

# Copy Go binaries from builder
COPY --from=go-builder /bin/pipeline     /usr/local/bin/pipeline
COPY --from=go-builder /bin/bot          /usr/local/bin/bot
COPY --from=go-builder /bin/backfill     /usr/local/bin/backfill
COPY --from=go-builder /bin/batch        /usr/local/bin/batch
COPY --from=go-builder /bin/backfill_emb /usr/local/bin/backfill_embeddings
COPY --from=go-builder /bin/verify_emb   /usr/local/bin/verify_embeddings
COPY --from=go-builder /bin/flush_local  /usr/local/bin/flush_local
COPY --from=go-builder /bin/count        /usr/local/bin/count
COPY --from=go-builder /bin/parse_pdf    /usr/local/bin/parse_pdf
COPY --from=go-builder /bin/backfill_bpk /usr/local/bin/backfill_bpk

# Copy Python scraper script
COPY backend/python/scraper/scrape.py /app/scraper/scrape.py

# Copy migration files (used by migrate_to_pg tool once Stream A creates it)
COPY backend/migrations /app/migrations

# Copy config files
COPY backend/configs /app/configs

WORKDIR /app

# Default command — override via docker-compose or `docker run hukum-aneh <cmd>`
CMD ["pipeline"]
