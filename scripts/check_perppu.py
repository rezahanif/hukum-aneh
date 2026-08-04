#!/usr/bin/env python3
"""Check what Perppu files exist on Drive and what peraturan.go.id lists."""

import re
import json
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
PERATURAN_FOLDER = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"

SLUG_RE = re.compile(r'^(uu|uud|perppu|pp|perpres|keppres|inpres)-no-(\d+)-tahun-(\d+)$')


def main():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    drv = build("drive", "v3", credentials=creds)

    # List all files still in peraturan/
    all_files = []
    page_token = None
    while True:
        results = drv.files().list(
            q=f"'{PERATURAN_FOLDER}' in parents and trashed=false",
            fields="files(id, name, size), nextPageToken",
            pageSize=1000,
            pageToken=page_token
        ).execute()
        all_files.extend(results.get("files", []))
        page_token = results.get("nextPageToken")
        if not page_token:
            break

    print(f"Total files in peraturan/ folder: {len(all_files)}")

    perppu_files = []
    type_counts = {}
    for f in all_files:
        name = f["name"].lower()
        m = SLUG_RE.match(name)
        if not m:
            # Check non-standard names
            if 'perppu' in name:
                perppu_files.append(f)
            continue
        dtype = m.group(1)
        type_counts[dtype] = type_counts.get(dtype, 0) + 1
        if dtype == "perppu":
            perppu_files.append(f)

    print(f"\nFile type breakdown in peraturan/:")
    for t, c in sorted(type_counts.items()):
        print(f"  {t}: {c}")

    print(f"\n{'='*70}")
    print(f"  PERPPU files on Drive: {len(perppu_files)}")
    print(f"{'='*70}")
    for f in sorted(perppu_files, key=lambda x: x["name"]):
        size_kb = int(f.get("size", 0)) // 1024
        m = SLUG_RE.match(f["name"].lower())
        if m:
            print(f"  perppu-no-{m.group(2)}/th.{m.group(3)}  {size_kb:>6}KB  {f['name'][:80]}")
        else:
            print(f"  (non-standard slug)  {size_kb:>6}KB  {f['name'][:80]}")

    # Save perppu info
    perppu_data = []
    for f in perppu_files:
        m = SLUG_RE.match(f["name"].lower())
        if m:
            perppu_data.append({
                "slug": m.group(0),
                "number": int(m.group(2)),
                "year": int(m.group(3)),
                "file_id": f["id"],
                "name": f["name"]
            })
    with open("/home/z/my-project/download/perppu_on_drive.json", "w") as fp:
        json.dump(perppu_data, fp, indent=2)
    print(f"\nPerppu data saved to download/perppu_on_drive.json")


if __name__ == "__main__":
    main()