#!/usr/bin/env python3
"""
investigate_perppu.py — Full investigation of Perppu coverage.

Run this on the user's server (182.8.250.2) where peraturan.go.id is reachable.
It will:
1. Scrape ALL berlaku Perppu listing pages from peraturan.go.id
2. Check which have downloadable PDFs
3. Compare with what's already on Drive
4. Generate a download plan for missing files

Usage:
  python3 investigate_perppu.py              # investigate only (no downloads)
  python3 investigate_perppu.py --download     # also download missing Perppu to Drive
"""

import argparse
import hashlib
import io
import json
import os
import re
import time
from collections import Counter

import requests
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseUpload

# ── Config ──────────────────────────────────────────────────────────────────
PERATURAN_BASE = "https://www.peraturan.go.id"
TOKEN_PATH = "/home/z/my-project/upload/token.json"
DRIVE_FOLDER_ID = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"
SCOPES = ["https://www.googleapis.com/auth/drive"]

SESSION = requests.Session()
SESSION.headers.update({
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
                  "(KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
})

LAW_LINK_RE = re.compile(
    r'href="/id/(perppu-no-\d+-tahun-\d+)"[^>]*title="lihat detail"[^>]*>([^<]*)</a>'
)
SLUG_RE = re.compile(r'^perppu-no-(\d+)-tahun-(\d+)$')


def get_drive_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def get_perppu_on_drive(drv):
    """Get set of (number, year) tuples for Perppu on Drive."""
    perppu = set()
    all_files = []
    page_token = None
    while True:
        results = drv.files().list(
            q=f"'{DRIVE_FOLDER_ID}' in parents and trashed=false and name contains 'perppu'",
            fields="files(id, name), nextPageToken",
            pageSize=1000,
            pageToken=page_token
        ).execute()
        all_files.extend(results.get("files", []))
        page_token = results.get("nextPageToken")
        if not page_token:
            break

    for f in all_files:
        name = f["name"].lower().replace(".pdf", "")
        m = SLUG_RE.match(name)
        if m:
            perppu.add((int(m.group(1)), int(m.group(2))))
    return perppu, all_files


def scrape_perppu_listing(max_pages=50):
    """Scrape ALL berlaku Perppu from peraturan.go.id listing pages.
    Returns list of (slug, title) tuples."""
    all_perppu = []

    for page in range(1, max_pages + 1):
        url = f"{PERATURAN_BASE}/perppu?PeraturanSearch%5Bstatus%5D=Berlaku&page={page}"
        print(f"  Scraping page {page}... ", end="", flush=True)

        try:
            resp = SESSION.get(url, timeout=30)
            if resp.status_code != 200:
                print(f"HTTP {resp.status_code}")
                break
        except Exception as e:
            print(f"Error: {e}")
            break

        matches = LAW_LINK_RE.findall(resp.text)
        if not matches:
            print("empty page, done.")
            break

        for slug, title in matches:
            m = SLUG_RE.match(slug)
            if m:
                all_perppu.append({
                    "slug": slug,
                    "number": int(m.group(1)),
                    "year": int(m.group(2)),
                    "title": title.strip(),
                })

        print(f"found {len(matches)} items (total: {len(all_perppu)})")
        time.sleep(0.5)

    return all_perppu


def check_pdf_availability(perppu_list):
    """Check which Perppu have downloadable PDFs."""
    available = []
    unavailable = []

    for item in perppu_list:
        pdf_url = f"{PERATURAN_BASE}/files/{item['slug']}.pdf"
        try:
            resp = SESSION.head(pdf_url, timeout=15, allow_redirects=True)
            if resp.status_code == 200 and 'pdf' in resp.headers.get('content-type', '').lower():
                item['pdf_url'] = pdf_url
                available.append(item)
            else:
                unavailable.append(item)
        except:
            unavailable.append(item)
        time.sleep(0.2)

    return available, unavailable


def upload_pdf(drv, folder_id, filename, pdf_bytes):
    """Upload PDF to Drive."""
    new_md5 = hashlib.md5(pdf_bytes).hexdigest()
    query = f"name='{filename}' and '{folder_id}' in parents and trashed=false"
    existing = drv.files().list(q=query, fields="files(id,md5Checksum,size)", pageSize=1).execute()
    if existing.get("files"):
        ef = existing["files"][0]
        if ef.get("md5Checksum") == new_md5:
            return None  # identical
        drv.files().delete(fileId=ef["id"]).execute()

    media = MediaIoBaseUpload(io.BytesIO(pdf_bytes), mimetype="application/pdf", resumable=True)
    result = drv.files().create(
        body={"name": filename, "parents": [folder_id]},
        media_body=media, fields="id,size"
    ).execute()
    return result["id"]


def download_missing(drv, missing_items):
    """Download and upload missing Perppu to Drive."""
    downloaded = 0
    failed = 0

    for item in missing_items:
        pdf_url = f"{PERATURAN_BASE}/files/{item['slug']}.pdf"
        try:
            resp = SESSION.get(pdf_url, timeout=60)
            if resp.status_code != 200 or 'text/html' in resp.headers.get('content-type', ''):
                print(f"  FAIL {item['slug']}: no PDF")
                failed += 1
                continue

            filename = f"{item['slug']}_{re.sub(r'[^\w\s\-]', '', item['title'][:80])}.pdf"
            file_id = upload_pdf(drv, DRIVE_FOLDER_ID, filename, resp.content)
            if file_id:
                size_kb = len(resp.content) // 1024
                print(f"  OK   {item['slug']} ({size_kb}KB)")
                downloaded += 1
            else:
                print(f"  SKIP {item['slug']} (duplicate)")
        except Exception as e:
            print(f"  FAIL {item['slug']}: {e}")
            failed += 1
        time.sleep(0.3)

    return downloaded, failed


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--download", action="store_true", help="Download missing Perppu")
    args = parser.parse_args()

    # Step 1: Get Perppu on Drive
    print("=== Step 1: Checking Perppu on Drive ===")
    drv = get_drive_service()
    drive_perppu, drive_files = get_perppu_on_drive(drv)
    print(f"  Perppu files on Drive: {len(drive_files)}")
    for f in sorted(drive_files, key=lambda x: x["name"]):
        print(f"    {f['name']}")
    print(f"  Unique (number, year): {len(drive_perppu)}")
    for num, year in sorted(drive_perppu):
        print(f"    Perppu No.{num}/Th.{year}")

    # Step 2: Scrape peraturan.go.id listing
    print(f"\n=== Step 2: Scraping peraturan.go.id/perppu (berlaku) ===")
    all_perppu = scrape_perppu_listing(max_pages=50)
    print(f"  Total berlaku Perppu listed: {len(all_perppu)}")

    if not all_perppu:
        print("  WARNING: No Perppu found! Site may be down or structure changed.")
        return

    # Year distribution
    years = Counter(p["year"] for p in all_perppu)
    print(f"  Year distribution: {dict(sorted(years.items()))}")

    # Step 3: Compare
    site_perppu = set((p["number"], p["year"]) for p in all_perppu)
    on_drive = drive_perppu
    missing = site_perppu - on_drive
    extra = on_drive - site_perppu

    print(f"\n=== Step 3: Comparison ===")
    print(f"  On peraturan.go.id (berlaku): {len(site_perppu)}")
    print(f"  On Drive:                      {len(on_drive)}")
    print(f"  Missing from Drive:            {len(missing)}")
    print(f"  Extra on Drive (not berlaku):  {len(extra)}")

    if missing:
        print(f"\n  Missing Perppu:")
        missing_items = [p for p in all_perppu if (p["number"], p["year"]) in missing]
        for item in sorted(missing_items, key=lambda x: (x["year"], x["number"])):
            print(f"    Perppu No.{item['number']}/Th.{item['year']} - {item['title'][:60]}")

    if extra:
        print(f"\n  Extra on Drive (not in berlaku listing):")
        for num, year in sorted(extra):
            print(f"    Perppu No.{num}/Th.{year}")

    # Step 4: Check PDF availability for missing
    if missing:
        print(f"\n=== Step 4: Checking PDF availability for {len(missing)} missing ===")
        missing_items = [p for p in all_perppu if (p["number"], p["year"]) in missing]
        available, unavailable = check_pdf_availability(missing_items)
        print(f"  Have PDF:      {len(available)}")
        print(f"  No PDF found:  {len(unavailable)}")

        if unavailable:
            print(f"\n  Unavailable (no PDF on site):")
            for item in sorted(unavailable, key=lambda x: (x["year"], x["number"])):
                print(f"    Perppu No.{item['number']}/Th.{item['year']}")

        # Save report
        report = {
            "on_drive": sorted([f"perppu-no-{n}-tahun-{y}" for n, y in on_drive]),
            "on_site_berlaku": sorted([p["slug"] for p in all_perppu]),
            "missing_with_pdf": [p for p in available],
            "missing_no_pdf": [p for p in unavailable],
            "extra_on_drive": sorted([f"perppu-no-{n}-tahun-{y}" for n, y in extra]),
        }
        report_path = "/home/z/my-project/download/perppu_gap_report.json"
        with open(report_path, "w") as f:
            json.dump(report, f, indent=2, ensure_ascii=False)
        print(f"\n  Report saved to {report_path}")

        # Step 5: Download if requested
        if args.download and available:
            print(f"\n=== Step 5: Downloading {len(available)} missing Perppu ===")
            dl_ok, dl_fail = download_missing(drv, available)
            print(f"\n  Downloaded: {dl_ok}, Failed: {dl_fail}")
        elif args.download:
            print(f"\n  No downloadable PDFs found for missing Perppu.")
        else:
            print(f"\n  Use --download to fetch {len(available)} missing Perppu with PDFs.")
    else:
        print(f"\n  All berlaku Perppu are already on Drive!")


if __name__ == "__main__":
    main()
