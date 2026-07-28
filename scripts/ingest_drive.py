#!/usr/bin/env python3
"""
ingest_drive.py — Ingest Indonesian law PDFs from Google Drive into PostgreSQL.

Pipeline:
  1. Auth with Google Drive (OAuth2)
  2. List all PDFs in the target folder
  3. Dedup by MD5 (Drive-level) + law_number (semantic, PG-level)
  4. Download PDF, parse (pdftotext → OCR fallback)
  5. Extract law_number + metadata from parsed text (regex)
  6. Insert into law_documents + law_versions (idempotent)

Usage:
  python3 ingest_drive.py                    # first run: opens browser for OAuth
  python3 ingest_drive.py --dry-run           # list what would be ingested
  python3 ingest_drive.py --force             # re-parse and update existing
  python3 ingest_drive.py --limit 10          # only process first 10 PDFs

Env vars (or defaults):
  POSTGRES_HOST     default: localhost
  POSTGRES_PORT     default: 5432
  POSTGRES_DB       default: hukum_aneh
  POSTGRES_USER     default: hukum
  POSTGRES_PASSWORD (required)
  DRIVE_FOLDER_ID   (required) Google Drive folder ID
  DRIVE_CREDENTIALS_PATH  default: client_secret.json
  DRIVE_TOKEN_PATH        default: token.json
"""

import argparse
import hashlib
import io
import json
import os
import re
import sqlite3
import subprocess
import sys
import tempfile
import time
import uuid
from datetime import datetime
from pathlib import Path

import psycopg2
import psycopg2.extras
from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseDownload


# ============================================================================
# Constants
# ============================================================================

DRIVE_SCOPES = ["https://www.googleapis.com/auth/drive.readonly"]

# Regex patterns for Indonesian law metadata extraction
# These cover the standard header format found in 90%+ of Indonesian regulations

DOC_TYPE_PATTERNS = {
    # Order matters — more specific first
    "PERATURAN PEMERINTAH PENGGANTI UNDANG-UNDANG": "Perppu",
    "UNDANG-UNDANG DASAR NEGARA REPUBLIK INDONESIA TAHUN 1945": "UUD 1945",
    "UNDANG-UNDANG REPUBLIK INDONESIA": "UU",
    "PERATURAN PEMERINTAH REPUBLIK INDONESIA": "PP",
    "PERATURAN PRESIDEN REPUBLIK INDONESIA": "Perpres",
    "KEPUTUSAN PRESIDEN REPUBLIK INDONESIA": "Keppres",
    "PERATURAN MENTERI": "Permen",
    "KEPUTUSAN MENTERI": "Kepmen",
    "PERATURAN BANK INDONESIA": "PBI",
    "PERATURAN DAERAH": "Perda",
    "TRITURA": "TAP MPR",
    "KETETAPAN MAJELIS PERMUSYAWARATAN RAKYAT": "TAP MPR",
    "KETETAPAN DEWAN PERWAKILAN RAKYAT": "TAP DPR",
}

# Matches: "NOMOR 11 TAHUN 2020" or "No. 11 Tahun 2020" or "Nomor 11/2020"
LAW_NUMBER_RE = re.compile(
    r"(?:NOMOR|NO\.?|Nomor)\s+(\d+)\s+(?:TAHUN|THN|Tahun|th\.?|/)\s*(\d{4})",
    re.IGNORECASE,
)

# Matches: "TENTANG" followed by the subject
TENTANG_RE = re.compile(
    r"TENTANG\s+(.+?)(?:\n|\r|DENGAN|DENGAN\s+RAHMAT|BAB\s|PASAL\s|Dengan\s|dengan\s)",
    re.IGNORECASE | re.DOTALL,
)

# Header/footer noise to strip from OCR text
NOISE_LINES = [
    "DISTRIBUSI II",
    "DISTRIBUSI III",
    "LEMBAGA NEGARA REPUBLIK INDONESIA",
    "TAHUN 2020",
    "BERITA NEGARA REPUBLIK INDONESIA",
    "PENJELASAN ATAS",
    "JAKARTA",
    "SALINAN",
    "SEKRETARIAT NEGARA",
    "RI",
]

# TESSDATA_PREFIX for Indonesian OCR
TESSDATA_PREFIX = os.environ.get("TESSDATA_PREFIX", "")


# ============================================================================
# Metadata Extraction
# ============================================================================


def extract_metadata(text: str, filename: str) -> dict:
    """
    Extract law_number, title, document_type, published_date from parsed text.

    law_number format: "{DocType} No. {Number} Tahun {Year}"
    Example: "UU No. 11 Tahun 2020"

    Returns dict with keys: law_number, title, document_type, published_date, level
    """
    # Take first 3000 chars for header analysis (first 2-3 pages)
    header = text[:3000].upper()
    header_orig = text[:3000]

    # 1. Detect document type
    doc_type = None
    doc_type_long = None
    for pattern, dtype in DOC_TYPE_PATTERNS.items():
        if pattern in header:
            doc_type = dtype
            doc_type_long = pattern
            break

    # 2. Extract number and year
    number = None
    year = None
    match = LAW_NUMBER_RE.search(text[:2000])
    if match:
        number = match.group(1)
        year = match.group(2)

    # 3. Build law_number
    law_number = None
    if doc_type and number and year:
        law_number = f"{doc_type} No. {number} Tahun {year}"
    elif number and year:
        # Fallback: use filename to infer doc type
        doc_type = guess_doc_type_from_filename(filename)
        if doc_type:
            law_number = f"{doc_type} No. {number} Tahun {year}"

    # 4. Extract title (subject after TENTANG)
    title = None
    tentang_match = TENTANG_RE.search(header_orig)
    if tentang_match:
        title = tentang_match.group(1).strip()
        # Clean up multi-line titles
        title = re.sub(r"\s+", " ", title)
        # Remove trailing punctuation
        title = title.rstrip(",.;")
    elif law_number:
        title = law_number

    # 5. Published date
    published_date = f"{year}-01-01" if year else None

    # 6. Level
    level = "national"  # default for Drive-imported laws

    # 7. If still no law_number, try harder with LLM-style fallback patterns
    if not law_number:
        law_number = try_fallback_extraction(text, filename)

    return {
        "law_number": law_number,
        "title": title or filename.replace(".pdf", "").replace(".PDF", ""),
        "document_type": doc_type or "Unknown",
        "published_date": published_date,
        "level": level,
        "source": "google_drive",
        "source_url": f"https://drive.google.com/drive/folders/" + os.environ.get(
            "DRIVE_FOLDER_ID", ""
        ),
    }


def guess_doc_type_from_filename(filename: str) -> str | None:
    """Infer document type from filename patterns."""
    fn = filename.upper()
    patterns = [
        (r"UU\s*No", "UU"),
        (r"UNDANG.?UNDANG", "UU"),
        (r"PP\s*No", "PP"),
        (r"PERATURAN.?PEMERINTAH", "PP"),
        (r"PERPRES\s*No", "Perpres"),
        (r"PERATURAN.?PRESIDEN", "Perpres"),
        (r"KEPPRES\s*No", "Keppres"),
        (r"PERMEN\s*No", "Permen"),
        (r"KEPMEN\s*No", "Kepmen"),
        (r"PBI\s*No", "PBI"),
        (r"PERDA\s*No", "Perda"),
    ]
    for pat, dtype in patterns:
        if re.search(pat, fn):
            return dtype
    return None


def try_fallback_extraction(text: str, filename: str) -> str | None:
    """Try alternative patterns when standard header extraction fails."""
    # Pattern: "UU/2020/11" or "PP/2021/56"
    alt_re = re.compile(
        r"(UU|PP|Perpres|Perppu|Keppres|Permen|Kepmen|PBI|Perda)\s*/\s*(\d{4})\s*/\s*(\d+)",
        re.IGNORECASE,
    )
    m = alt_re.search(text[:3000])
    if m:
        return f"{m.group(1)} No. {m.group(3)} Tahun {m.group(2)}"

    # Pattern: "(UU 11/2020)" or "(PP 56/2021)"
    paren_re = re.compile(
        r"\((UU|PP|Perpres|Perppu)\s+(\d+)\s*/\s*(\d{4})\)", re.IGNORECASE
    )
    m = paren_re.search(text[:3000])
    if m:
        return f"{m.group(1)} No. {m.group(2)} Tahun {m.group(3)}"

    return None


# ============================================================================
# PDF Parsing
# ============================================================================


def parse_pdf(pdf_path: str) -> tuple[str, str]:
    """
    Parse a PDF file. Returns (text, source) where source is 'pdftotext' or 'ocr'.
    Tries pdftotext first, falls back to tesseract OCR.
    """
    # Phase 1: Try pdftotext (fast, works for text-embedded PDFs)
    try:
        result = subprocess.run(
            ["pdftotext", "-layout", pdf_path, "-"],
            capture_output=True,
            text=True,
            timeout=60,
        )
        text = result.stdout
        if len(text.strip()) > 100:
            return clean_text(text), "pdftotext"
    except (subprocess.TimeoutExpired, FileNotFoundError) as e:
        print(f"  pdftotext failed: {e}")

    # Phase 2: OCR fallback
    print("  Text extraction empty/failed, falling back to OCR...")
    return ocr_pdf(pdf_path), "ocr"


def ocr_pdf(pdf_path: str) -> str:
    """Convert PDF to images via pdftoppm, then OCR each with tesseract."""
    with tempfile.TemporaryDirectory() as img_dir:
        img_prefix = os.path.join(img_dir, "page")

        # Convert to images: 300 DPI, up to 200 pages
        try:
            subprocess.run(
                ["pdftoppm", "-png", "-r", "300", "-l", "200", pdf_path, img_prefix],
                capture_output=True,
                text=True,
                timeout=300,
                check=True,
            )
        except subprocess.CalledProcessError as e:
            print(f"  pdftoppm failed: {e.stderr}")
            return ""

        # Find generated images
        images = sorted([f for f in os.listdir(img_dir) if f.endswith(".png")])
        if not images:
            return ""

        print(f"  OCR processing {len(images)} pages...")

        # Build tesseract command
        env = os.environ.copy()
        if TESSDATA_PREFIX:
            env["TESSDATA_PREFIX"] = TESSDATA_PREFIX

        all_text = []
        for i, img_name in enumerate(images):
            img_path = os.path.join(img_dir, img_name)
            try:
                result = subprocess.run(
                    [
                        "tesseract",
                        img_path,
                        "stdout",
                        "-l",
                        "ind+eng",
                        "--psm",
                        "6",
                    ],
                    capture_output=True,
                    text=True,
                    timeout=60,
                    env=env,
                )
                all_text.append(result.stdout)
                os.remove(img_path)  # free disk immediately
            except subprocess.TimeoutExpired:
                print(f"  OCR timeout on page {i+1}, skipping")
                continue

        return clean_text("\n".join(all_text))


def clean_text(text: str) -> str:
    """Clean OCR/pdftotext output."""
    lines = text.split("\n")
    cleaned = []
    for line in lines:
        stripped = line.strip()
        if not stripped:
            continue
        # Skip known noise lines
        upper = stripped.upper()
        if any(noise in upper for noise in NOISE_LINES):
            continue
        # Skip page numbers (standalone numbers)
        if re.match(r"^\d{1,4}$", stripped):
            continue
        cleaned.append(stripped)

    result = "\n".join(cleaned)
    # Collapse multiple blank lines
    while "\n\n\n" in result:
        result = result.replace("\n\n\n", "\n\n")
    return result.strip()


# ============================================================================
# Google Drive
# ============================================================================


def authenticate(creds_path: str, token_path: str) -> Credentials:
    """Authenticate with Google Drive. Generates token.json on first run."""
    creds = None
    if os.path.exists(token_path):
        creds = Credentials.from_authorized_user_file(token_path, DRIVE_SCOPES)

    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(creds_path, DRIVE_SCOPES)
            creds = flow.run_local_server(port=0)

        # Save token for future runs
        with open(token_path, "w") as f:
            json.dump(
                {
                    "token": creds.token,
                    "refresh_token": creds.refresh_token,
                    "token_uri": creds.token_uri,
                    "client_id": creds.client_id,
                    "client_secret": creds.client_secret,
                    "scopes": creds.scopes,
                },
                f,
            )
        print(f"OAuth token saved to {token_path}")

    return creds


def list_drive_pdfs(service, folder_id: str) -> list[dict]:
    """List all PDFs in a Drive folder with name, ID, and MD5."""
    files = []
    page_token = None

    while True:
        query = f"mimeType='application/pdf' and trashed=false"
        if folder_id:
            query += f" and '{folder_id}' in parents"

        results = (
            service.files()
            .list(
                q=query,
                spaces="drive",
                fields="nextPageToken, files(id, name, md5Checksum, size)",
                pageToken=page_token,
                pageSize=100,
            )
            .execute()
        )

        for f in results.get("files", []):
            files.append(
                {
                    "id": f["id"],
                    "name": f["name"],
                    "md5": f.get("md5Checksum", ""),
                    "size": int(f.get("size", 0)),
                }
            )

        page_token = results.get("nextPageToken")
        if not page_token:
            break

    return files


def download_pdf(service, file_id: str) -> io.BytesIO:
    """Download a PDF from Drive, returns BytesIO."""
    request = service.files().get_media(fileId=file_id)
    fh = io.BytesIO()
    downloader = MediaIoBaseDownload(fh, request)
    done = False
    while not done:
        _, done = downloader.next_chunk()
    fh.seek(0)
    return fh


# ============================================================================
# PostgreSQL
# ============================================================================


def get_pg_conn():
    """Create a PostgreSQL connection from env vars."""
    return psycopg2.connect(
        host=os.environ.get("POSTGRES_HOST", "localhost"),
        port=int(os.environ.get("POSTGRES_PORT", "5432")),
        dbname=os.environ.get("POSTGRES_DB", "hukum_aneh"),
        user=os.environ.get("POSTGRES_USER", "hukum"),
        password=os.environ.get("POSTGRES_PASSWORD", ""),
        sslmode=os.environ.get("POSTGRES_SSL_MODE", "prefer"),
    )


def insert_law_document(cur, metadata: dict, drive_file: dict, force: bool = False) -> str | None:
    """
    Insert a law document into PG. Returns doc_id on success, None on skip.
    Uses ON CONFLICT DO NOTHING for idempotency (unless --force).
    """
    law_number = metadata["law_number"]
    if not law_number:
        print(f"    SKIP: no law_number extracted from {drive_file['name']}")
        return None

    doc_id = str(uuid.uuid4())
    now = datetime.utcnow().isoformat()

    if force:
        # Delete existing and re-insert
        cur.execute(
            "DELETE FROM law_versions WHERE law_document_id IN (SELECT id FROM law_documents WHERE law_number = %s)",
            (law_number,),
        )
        cur.execute(
            "DELETE FROM law_documents WHERE law_number = %s",
            (law_number,),
        )

    # Check if law_number already exists
    cur.execute(
        "SELECT id FROM law_documents WHERE law_number = %s",
        (law_number,),
    )
    existing = cur.fetchone()
    if existing and not force:
        print(f"    SKIP: {law_number} already exists (id={existing[0]})")
        return None

    if existing and force:
        doc_id = existing[0]

    raw_file_path = f"drive://{drive_file['id']}"

    try:
        cur.execute(
            """
            INSERT INTO law_documents
                (id, law_number, title, source_url, source, level,
                 document_type, raw_file_path, published_date, status,
                 created_at, updated_at)
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
            ON CONFLICT (id) DO UPDATE SET
                law_number = EXCLUDED.law_number,
                title = EXCLUDED.title,
                source_url = EXCLUDED.source_url,
                source = EXCLUDED.source,
                level = EXCLUDED.level,
                document_type = EXCLUDED.document_type,
                raw_file_path = EXCLUDED.raw_file_path,
                published_date = EXCLUDED.published_date,
                status = EXCLUDED.status,
                updated_at = EXCLUDED.updated_at
            RETURNING id
            """,
            (
                doc_id,
                law_number,
                metadata["title"],
                metadata["source_url"],
                metadata["source"],
                metadata["level"],
                metadata["document_type"],
                raw_file_path,
                metadata["published_date"],
                "parsed",
                now,
                now,
            ),
        )
        result = cur.fetchone()
        if result:
            doc_id = result[0]
            print(f"    INSERTED: {law_number} (id={doc_id})")
            return doc_id
        return None
    except psycopg2.errors.UniqueViolation:
        print(f"    SKIP: {law_number} violates unique constraint")
        return None


def insert_law_version(cur, doc_id: str, text_content: str, parse_source: str):
    """Insert a parsed law version."""
    version_id = str(uuid.uuid4())
    now = datetime.utcnow().isoformat()
    version_number = int(time.time())

    cur.execute(
        """
        INSERT INTO law_versions
            (id, law_document_id, version_number, text_content, parsed_at)
        VALUES (%s, %s, %s, %s, %s)
        ON CONFLICT (id) DO UPDATE SET
            text_content = EXCLUDED.text_content,
            parsed_at = EXCLUDED.parsed_at
        """,
        (version_id, doc_id, version_number, text_content, now),
    )
    print(f"    VERSION: saved ({parse_source}, {len(text_content)} chars)")


# ============================================================================
# MD5 Dedup Cache
# ============================================================================


def get_md5_cache_path() -> str:
    return os.path.join(os.path.dirname(__file__), ".ingest_md5_cache.db")


def init_md5_cache() -> sqlite3.Connection:
    conn = sqlite3.connect(get_md5_cache_path())
    conn.execute(
        "CREATE TABLE IF NOT EXISTS processed (drive_file_id TEXT PRIMARY KEY, md5 TEXT, law_number TEXT)"
    )
    conn.commit()
    return conn


def was_processed(cache: sqlite3.Connection, file_id: str, md5: str) -> bool:
    row = cache.execute(
        "SELECT law_number FROM processed WHERE drive_file_id = ?", (file_id,)
    ).fetchone()
    if row:
        cached_md5 = cache.execute(
            "SELECT md5 FROM processed WHERE drive_file_id = ?", (file_id,)
        ).fetchone()
        if cached_md5 and cached_md5[0] == md5:
            return True
    return False


def mark_processed(
    cache: sqlite3.Connection, file_id: str, md5: str, law_number: str
):
    cache.execute(
        "INSERT OR REPLACE INTO processed (drive_file_id, md5, law_number) VALUES (?, ?, ?)",
        (file_id, md5, law_number),
    )
    cache.commit()


# ============================================================================
# Main
# ============================================================================


def main():
    parser = argparse.ArgumentParser(description="Ingest law PDFs from Google Drive to PostgreSQL")
    parser.add_argument("--dry-run", action="store_true", help="List files without ingesting")
    parser.add_argument("--force", action="store_true", help="Re-parse and update existing laws")
    parser.add_argument("--limit", type=int, default=0, help="Max files to process (0=all)")
    parser.add_argument(
        "--folder-id",
        default=os.environ.get("DRIVE_FOLDER_ID", ""),
        help="Google Drive folder ID",
    )
    parser.add_argument(
        "--credentials",
        default=os.environ.get(
            "DRIVE_CREDENTIALS_PATH", "/home/z/my-project/upload/client_secret_858768057989-7ql7keonq637uoejbmh8ptl7fl9cc5h3.apps.googleusercontent.com.json"
        ),
        help="Path to OAuth client secret JSON",
    )
    parser.add_argument(
        "--token",
        default=os.environ.get("DRIVE_TOKEN_PATH", "/home/z/my-project/upload/token.json"),
        help="Path to OAuth token JSON",
    )
    args = parser.parse_args()

    # Required env vars
    if not os.environ.get("POSTGRES_PASSWORD"):
        print("ERROR: POSTGRES_PASSWORD env var is required")
        sys.exit(1)
    if not args.folder_id:
        print("ERROR: --folder-id or DRIVE_FOLDER_ID env var is required")
        sys.exit(1)

    # Set TESSDATA_PREFIX for Indonesian OCR
    global TESSDATA_PREFIX
    home_tessdata = os.path.expanduser("~/.tessdata")
    if os.path.exists(os.path.join(home_tessdata, "ind.traineddata")):
        TESSDATA_PREFIX = home_tessdata

    print(f"TESSDATA_PREFIX={TESSDATA_PREFIX}")
    print(f"Drive folder: {args.folder_id}")
    print(f"PG: {os.environ.get('POSTGRES_HOST','localhost')}:{os.environ.get('POSTGRES_PORT','5432')}/{os.environ.get('POSTGRES_DB','hukum_aneh')}")
    print()

    # 1. Authenticate with Google Drive
    print("[1/4] Authenticating with Google Drive...")
    creds = authenticate(args.credentials, args.token)
    service = build("drive", "v3", credentials=creds)
    print("  Authenticated!")

    # 2. List PDFs in folder
    print(f"\n[2/4] Listing PDFs in folder...")
    drive_files = list_drive_pdfs(service, args.folder_id)
    print(f"  Found {len(drive_files)} PDF(s)")

    if not drive_files:
        print("No PDFs found. Exiting.")
        return

    if args.limit > 0:
        drive_files = drive_files[: args.limit]
        print(f"  Limited to first {args.limit} file(s)")

    if args.dry_run:
        print(f"\n[DRY RUN] Would process these {len(drive_files)} files:")
        for f in drive_files:
            size_mb = f["size"] / (1024 * 1024)
            print(f"  - {f['name']} ({size_mb:.1f} MB, md5={f['md5'][:12]}...)")
        return

    # 3. Connect to PG
    print(f"\n[3/4] Connecting to PostgreSQL...")
    conn = get_pg_conn()
    conn.autocommit = False
    print("  Connected!")

    # 4. Initialize MD5 dedup cache
    md5_cache = init_md5_cache()

    # 5. Process each PDF
    print(f"\n[4/4] Processing {len(drive_files)} PDF(s)...")
    stats = {"inserted": 0, "skipped": 0, "failed": 0, "no_law_number": 0}

    for i, drive_file in enumerate(drive_files):
        fname = drive_file["name"]
        print(f"\n[{i+1}/{len(drive_files)}] {fname}")

        # MD5 dedup check
        if was_processed(md5_cache, drive_file["id"], drive_file["md5"]):
            print(f"    SKIP: already processed (MD5 match)")
            stats["skipped"] += 1
            continue

        try:
            # Download PDF
            print(f"    Downloading ({drive_file['size']//(1024*1024)} MB)...")
            pdf_data = download_pdf(service, drive_file["id"])

            # Save to temp file for parsing
            with tempfile.NamedTemporaryFile(suffix=".pdf", delete=False) as tmp:
                tmp.write(pdf_data.read())
                tmp_path = tmp.name

            try:
                # Parse PDF
                print(f"    Parsing...")
                text, source = parse_pdf(tmp_path)
                print(f"    Parsed: {len(text)} chars via {source}")

                if len(text) < 50:
                    print(f"    SKIP: parsed text too short ({len(text)} chars)")
                    stats["failed"] += 1
                    continue

                # Extract metadata
                metadata = extract_metadata(text, fname)
                print(f"    Metadata: law_number={metadata['law_number']}, doc_type={metadata['document_type']}, title={metadata['title'][:60]}...")

                # Insert into PG
                with conn.cursor() as cur:
                    doc_id = insert_law_document(cur, metadata, drive_file, args.force)
                    if doc_id:
                        insert_law_version(cur, doc_id, text, source)
                        conn.commit()
                        stats["inserted"] += 1
                        mark_processed(md5_cache, drive_file["id"], drive_file["md5"], metadata["law_number"] or "")
                    else:
                        conn.rollback()
                        if not metadata["law_number"]:
                            stats["no_law_number"] += 1
                        else:
                            stats["skipped"] += 1

            finally:
                os.unlink(tmp_path)

        except Exception as e:
            print(f"    ERROR: {e}")
            conn.rollback()
            stats["failed"] += 1

    conn.close()
    md5_cache.close()

    # Summary
    print(f"\n{'='*60}")
    print(f"DONE. Results:")
    print(f"  Inserted:     {stats['inserted']}")
    print(f"  Skipped:      {stats['skipped']}")
    print(f"  No law_number:{stats['no_law_number']}")
    print(f"  Failed:       {stats['failed']}")
    print(f"{'='*60}")


if __name__ == "__main__":
    main()
