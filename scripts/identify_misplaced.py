#!/usr/bin/env python3
"""
identify_misplaced.py — Scan peraturan/ Drive folder for Inpres & Keppres files
that were accidentally uploaded there instead of their own folders.
"""

import json
import os
import re
from collections import Counter

from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
PERATURAN_FOLDER = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"

# Slug patterns
KEPPRES_RE = re.compile(r'^keppres-no-(\d+)-tahun-(\d+)', re.IGNORECASE)
INPRES_RE = re.compile(r'^inpres-no-(\d+)-tahun-(\d+)', re.IGNORECASE)


def main():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, [
        "https://www.googleapis.com/auth/drive"
    ])
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    drv = build("drive", "v3", credentials=creds)

    # List all files in peraturan/
    all_files = []
    page_token = None
    while True:
        results = drv.files().list(
            q=f"'{PERATURAN_FOLDER}' in parents and trashed=false",
            fields="files(id, name, size, md5Checksum), nextPageToken",
            pageSize=1000,
            pageToken=page_token
        ).execute()
        all_files.extend(results.get("files", []))
        page_token = results.get("nextPageToken")
        if not page_token:
            break

    print(f"Total files in peraturan/: {len(all_files)}\n")

    keppres_files = []
    inpres_files = []
    other_types = Counter()

    for f in all_files:
        name = f["name"].lower()
        m = KEPPRES_RE.match(name)
        if m:
            keppres_files.append(f)
            continue
        m = INPRES_RE.match(name)
        if m:
            inpres_files.append(f)
            continue
        # Count other types for context
        for prefix in ["uu-no-", "pp-no-", "perppu-no-", "perpres-no-", 
                        "uud-no-", "tap-mpr-no-"]:
            if prefix in name:
                other_types[prefix.replace("-no-", "").upper()] += 1
                break

    # === Keppres ===
    print(f"{'='*70}")
    print(f"  KEPPRES files found: {len(keppres_files)}")
    print(f"{'='*70}")
    keppres_years = Counter()
    for f in sorted(keppres_files, key=lambda x: x["name"]):
        m = KEPPRES_RE.match(f["name"].lower())
        num, year = m.group(1), m.group(2)
        keppres_years[year] += 1
        size_kb = int(f.get("size", 0)) // 1024
        print(f"  {f['id'][:12]}  Keppres No.{num}/Th.{year}  {size_kb:>6}KB  {f['name'][:70]}")
    print(f"\n  Year distribution: {dict(sorted(keppres_years.items()))}")

    # === Inpres ===
    print(f"\n{'='*70}")
    print(f"  INPRES files found: {len(inpres_files)}")
    print(f"{'='*70}")
    inpres_years = Counter()
    for f in sorted(inpres_files, key=lambda x: x["name"]):
        m = INPRES_RE.match(f["name"].lower())
        num, year = m.group(1), m.group(2)
        inpres_years[year] += 1
        size_kb = int(f.get("size", 0)) // 1024
        print(f"  {f['id'][:12]}  Inpres No.{num}/Th.{year}  {size_kb:>6}KB  {f['name'][:70]}")
    print(f"\n  Year distribution: {dict(sorted(inpres_years.items()))}")

    # Summary
    print(f"\n{'='*70}")
    print(f"  SUMMARY")
    print(f"{'='*70}")
    print(f"  Total peraturan/ files:    {len(all_files)}")
    print(f"  Keppres (misplaced):       {len(keppres_files)}")
    print(f"  Inpres (misplaced):        {len(inpres_files)}")
    print(f"  Other (UU/PP/etc):         {len(all_files) - len(keppres_files) - len(inpres_files)}")
    print(f"  Other breakdown:           {dict(other_types)}")

    # Save file IDs for move script
    if keppres_files or inpres_files:
        move_data = {
            "keppres": [{"id": f["id"], "name": f["name"]} for f in keppres_files],
            "inpres": [{"id": f["id"], "name": f["name"]} for f in inpres_files],
        }
        out_path = "/home/z/my-project/download/misplaced_keppres_inpres.json"
        with open(out_path, "w") as fp:
            json.dump(move_data, fp, indent=2)
        print(f"\n  File IDs saved to: {out_path}")


if __name__ == "__main__":
    main()
