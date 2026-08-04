#!/usr/bin/env python3
"""
download_to_folder.py - Download berlaku perpres/inpres from peraturan.go.id to specific Drive folders.

Usage:
  python3 download_to_folder.py perpres
  python3 download_to_folder.py inpres
  python3 download_to_folder.py perpres --resume
  python3 download_to_folder.py perpres --limit 50
  python3 download_to_folder.py perpres --year 2024

Requires:
  - pip install requests google-api-python-client google-auth-httplib2 google-auth-oauthlib
  - GOOGLE_APPLICATION_CREDENTIALS or token.json for Drive access
"""

import argparse
import hashlib
import io
import json
import os
import re
import sys
import time
from pathlib import Path

import requests
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseUpload

BASE_URL = "https://www.peraturan.go.id"

FOLDER_IDS = {
    "perpres": "1RBhkXEH750LSjipVjAFrC8C9rcu0fTCw",
    "inpres":  "1KnQsVfO06J6IGudhBbpsX47dC26QeW3q",
}

# Regex for listing page links
LINK_RE = re.compile(
    r'href="/id/((?:perpres|inpres)-no-\d+-tahun-\d+)"'
    r'[^>]*title="lihat detail"[^>]*>([^<]*)</a>'
)
SLUG_RE = re.compile(r'^(perpres|inpres)-no-(\d+)-tahun-(\d+)$')

SESSION = requests.Session()
SESSION.headers.update({
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
                  "(KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
})


# ── Drive helpers ──────────────────────────────────────────────────────────


def auth(token_path="token.json", creds_path="client_secret.json"):
    scopes = ["https://www.googleapis.com/auth/drive"]
    creds = None
    if os.path.exists(token_path):
        creds = Credentials.from_authorized_user_file(token_path, scopes)
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            from google_auth_oauthlib.flow import InstalledAppFlow
            flow = InstalledAppFlow.from_client_secrets_file(creds_path, scopes)
            creds = flow.run_local_server(port=0)
        with open(token_path, "w") as f:
            json.dump({
                "token": creds.token, "refresh_token": creds.refresh_token,
                "token_uri": creds.token_uri, "client_id": creds.client_id,
                "client_secret": creds.client_secret, "scopes": creds.scopes,
            }, f)
    return creds


def upload_pdf(drv, folder_id, filename, pdf_bytes):
    new_md5 = hashlib.md5(pdf_bytes).hexdigest()
    q = f"name='{filename}' and '{folder_id}' in parents and trashed=false"
    existing = drv.files().list(q=q, fields="files(id,md5Checksum)", pageSize=1).execute()
    if existing.get("files"):
        if existing["files"][0].get("md5Checksum") == new_md5:
            return None  # same file
        drv.files().delete(fileId=existing["files"][0]["id"]).execute()
    media = MediaIoBaseUpload(io.BytesIO(pdf_bytes), mimetype="application/pdf", resumable=True)
    result = drv.files().create(
        body={"name": filename, "parents": [folder_id]},
        media_body=media, fields="id,size"
    ).execute()
    return result["id"]


def get_existing_slugs(drv, folder_id):
    slugs = set()
    pt = None
    while True:
        r = drv.files().list(
            q=f"'{folder_id}' in parents and trashed=false",
            fields="files(name), nextPageToken", pageSize=1000, pageToken=pt
        ).execute()
        for f in r.get("files", []):
            name = f["name"].lower().replace(".pdf", "")
            m = SLUG_RE.match(name)
            if m:
                slugs.add(m.group(0))
        pt = r.get("nextPageToken")
        if not pt:
            break
    return slugs


# ── Main download logic ────────────────────────────────────────────────────


def download_type(law_type, folder_id, resume=False, limit=0, year_filter=0):
    creds = auth()
    drv = build("drive", "v3", credentials=creds)

    # Get existing slugs to skip
    print(f"Checking existing files on Drive...")
    existing = get_existing_slugs(drv, folder_id)
    print(f"Already have: {len(existing)} {law_type}")

    # Determine max pages (generous)
    max_pages = 300 if law_type == "perpres" else 50

    # Resume cursor
    cursor_file = f".cursor_{law_type}.json"
    start_page = 1
    if resume and os.path.exists(cursor_file):
        with open(cursor_file) as f:
            start_page = json.load(f).get("page", 1)
        print(f"Resuming from page {start_page}")

    downloaded = 0
    skipped = 0
    no_pdf_streak = 0
    MAX_STREAK = 15  # tolerate more missing PDFs

    for page in range(start_page, max_pages + 1):
        if 0 < limit <= downloaded:
            print(f"Limit ({limit}) reached.")
            break

        url = f"{BASE_URL}/{law_type}?PeraturanSearch%5Bstatus%5D=Berlaku&page={page}"
        print(f"Page {page}... ", end="", flush=True)

        try:
            resp = SESSION.get(url, timeout=30)
            if resp.status_code != 200:
                print(f"HTTP {resp.status_code}")
                break
        except Exception as e:
            print(f"Error: {e}")
            break

        matches = LINK_RE.findall(resp.text)
        if not matches:
            print("empty, done.")
            break

        page_dl = 0
        page_skip = 0
        for slug, title in matches:
            if 0 < limit <= downloaded:
                break

            m = SLUG_RE.match(slug)
            if not m:
                continue
            doc_year = int(m.group(3))
            if year_filter and doc_year != year_filter:
                continue

            if slug in existing:
                page_skip += 1
                continue

            # Download PDF
            pdf_url = f"{BASE_URL}/files/{slug}.pdf"
            try:
                r = SESSION.get(pdf_url, timeout=60)
                if r.status_code != 200 or "text/html" in r.headers.get("content-type", ""):
                    no_pdf_streak += 1
                    if no_pdf_streak >= MAX_STREAK:
                        print(f"\n  {MAX_STREAK} consecutive missing PDFs, stopping.")
                        with open(cursor_file, "w") as f:
                            json.dump({"page": page, "downloaded": downloaded}, f)
                        return downloaded
                    continue
            except:
                no_pdf_streak += 1
                continue

            no_pdf_streak = 0
            clean_title = re.sub(r'[^\w\s-]', '', title.strip())[:80]
            filename = f"{slug}_{clean_title}.pdf"

            try:
                file_id = upload_pdf(drv, folder_id, filename, r.content)
                if file_id:
                    size_kb = len(r.content) // 1024
                    print(f"\n    + {slug} ({size_kb}KB)", end="")
                    downloaded += 1
                    page_dl += 1
                    existing.add(slug)  # avoid re-downloading in this session
                else:
                    page_skip += 1
            except Exception as e:
                print(f"\n    ! {slug} upload failed: {e}", end="")

            time.sleep(0.3)

        print(f" +{page_dl} skip:{page_skip} (total: {downloaded})")

        # Save cursor
        with open(cursor_file, "w") as f:
            json.dump({"page": page + 1, "downloaded": downloaded}, f)
        time.sleep(0.5)

    # Clean up cursor on completion
    if os.path.exists(cursor_file):
        os.remove(cursor_file)

    return downloaded


def main():
    parser = argparse.ArgumentParser(description="Download berlaku laws from peraturan.go.id")
    parser.add_argument("type", choices=["perpres", "inpres"], help="Law type")
    parser.add_argument("--resume", action="store_true", help="Resume from last page")
    parser.add_argument("--limit", type=int, default=0, help="Max downloads (0=unlimited)")
    parser.add_argument("--year", type=int, default=0, help="Filter by year")
    parser.add_argument("--token", default="token.json", help="Drive token path")
    parser.add_argument("--creds", default="client_secret.json", help="Drive credentials path")
    args = parser.parse_args()

    folder_id = FOLDER_IDS[args.type]
    print(f"{'='*60}")
    print(f"  Downloading {args.type.upper()} (berlaku)")
    print(f"  Target folder: {folder_id}")
    print(f"{'='*60}")

    total = download_type(
        args.type, folder_id,
        resume=args.resume,
        limit=args.limit,
        year_filter=args.year,
    )
    print(f"\nDone. {total} new {args.type} uploaded.")


if __name__ == "__main__":
    main()
