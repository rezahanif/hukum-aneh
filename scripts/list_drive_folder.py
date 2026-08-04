#!/usr/bin/env python3
"""List all files in a Drive folder with dedup analysis."""

import json
import re
from collections import Counter
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]

# All folder IDs to check
FOLDERS = {
    "shared_root": "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS",
}

# Slug patterns for all law types
SLUG_RE = re.compile(
    r'^(uu|uud|perppu|pp|perpres|keppres|inpres|tap-mpr)-no-(\d+)-tahun-(\d+)',
    re.IGNORECASE
)
SLUG_TAP_MPR_RE = re.compile(
    r'^tap-mpr-no-([ivxlcdm]+)', re.IGNORECASE
)


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def list_folder(service, folder_id, label=""):
    """Recursively list all files in a folder and subfolders."""
    all_files = []
    subfolders = []
    page_token = None

    while True:
        results = service.files().list(
            q=f"'{folder_id}' in parents and trashed=false",
            fields="files(id, name, size, md5Checksum, mimeType), nextPageToken",
            pageSize=1000,
            pageToken=page_token
        ).execute()
        for f in results.get("files", []):
            if f["mimeType"] == "application/vnd.google-apps.folder":
                subfolders.append(f)
            else:
                all_files.append(f)
        page_token = results.get("nextPageToken")
        if not page_token:
            break

    # Recurse into subfolders
    for sf in subfolders:
        sf_label = f"{label}/{sf['name']}" if label else sf["name"]
        child_files, child_folders = list_folder(service, sf["id"], sf_label)
        all_files.extend(child_files)
        subfolders.extend(child_folders)

    return all_files, subfolders


def classify_file(name):
    """Extract (type, number, year, slug_key) from filename."""
    name_lower = name.lower().replace(".pdf", "")
    
    m = SLUG_TAP_MPR_RE.match(name_lower)
    if m:
        return {"type": "TAP MPR", "slug": name_lower.split("_")[0] if "_" in name_lower else name_lower}
    
    m = SLUG_RE.match(name_lower)
    if m:
        dtype = m.group(1).upper()
        num = m.group(2)
        year = m.group(3)
        slug = f"{m.group(1)}-no-{num}-tahun-{year}"
        return {"type": dtype, "number": int(num), "year": int(year), "slug": slug}
    
    return {"type": "UNKNOWN", "slug": None}


def main():
    service = get_service()

    for folder_label, folder_id in FOLDERS.items():
        print(f"{'='*70}")
        print(f"  Scanning: {folder_label} ({folder_id})")
        print(f"{'='*70}")

        files, subfolders = list_folder(service, folder_id, folder_label)
        print(f"  Subfolders found: {len(subfolders)}")
        for sf in subfolders:
            print(f"    📁 {sf['name']}")
        print(f"  Total PDF files: {len(files)}")

        # Classify
        type_counts = Counter()
        type_years = {}  # {type: Counter({year: count})}
        all_slugs = {}  # {slug: [file_info]}
        unknown_files = []
        duplicates = []

        for f in files:
            info = classify_file(f["name"])
            if info["type"] == "UNKNOWN":
                unknown_files.append(f)
                continue

            type_counts[info["type"]] += 1
            slug = info["slug"]

            if slug:
                if slug not in all_slugs:
                    all_slugs[slug] = []
                all_slugs[slug].append(f)

                if info["type"] not in type_years:
                    type_years[info["type"]] = Counter()
                if "year" in info:
                    type_years[info["type"]][info["year"]] += 1

        # Duplicates (same slug, multiple files)
        dup_slugs = {s: files for s, files in all_slugs.items() if len(files) > 1}

        print(f"\n  ── Type breakdown ──")
        for t, c in sorted(type_counts.items()):
            print(f"    {t}: {c}")

        print(f"\n  ── By year ──")
        for dtype in sorted(type_years.keys()):
            yc = type_years[dtype]
            print(f"    {dtype}:")
            for y in sorted(yc.keys()):
                print(f"      {y}: {yc[y]}")

        if dup_slugs:
            print(f"\n  ── DUPLICATES ({len(dup_slugs)} slugs have multiple copies) ──")
            for slug, dups in sorted(dup_slugs.items()):
                print(f"    {slug}: {len(dups)} copies")
                for d in sorted(dups, key=lambda x: int(x.get("size", 0)), reverse=True):
                    size_kb = int(d.get("size", 0)) // 1024
                    md5 = d.get("md5Checksum", "?")[:8]
                    print(f"      {size_kb:>6}KB  md5:{md5}  {d['name'][:70]}")
        else:
            print(f"\n  ── No duplicates found ──")

        if unknown_files:
            print(f"\n  ── UNKNOWN/UNCLASSIFIED ({len(unknown_files)} files) ──")
            for f in sorted(unknown_files, key=lambda x: x["name"])[:30]:
                size_kb = int(f.get("size", 0)) // 1024
                print(f"    {size_kb:>6}KB  {f['name'][:80]}")
            if len(unknown_files) > 30:
                print(f"    ... and {len(unknown_files) - 30} more")

        # Save full inventory
        inventory = {
            "folder": folder_label,
            "folder_id": folder_id,
            "total_files": len(files),
            "type_counts": dict(type_counts),
            "duplicates": {s: len(f) for s, f in dup_slugs.items()},
            "unique_slugs": sorted(all_slugs.keys()),
            "unknown_count": len(unknown_files),
            "unknown_files": [f["name"] for f in unknown_files[:50]],
        }
        out_path = f"/home/z/my-project/download/drive_inventory_{folder_label}.json"
        with open(out_path, "w") as fp:
            json.dump(inventory, fp, indent=2, ensure_ascii=False)
        print(f"\n  Inventory saved to {out_path}")


if __name__ == "__main__":
    main()