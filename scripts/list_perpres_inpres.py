#!/usr/bin/env python3
"""List perpres and inpres files, get folder IDs."""

import re
import json
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
ROOT = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"

SLUG_RE = re.compile(r'^(perpres|inpres)-no-(\d+)-tahun-(\d+)', re.I)


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def list_folder(service, folder_id):
    files = []
    pt = None
    while True:
        r = service.files().list(
            q=f"'{folder_id}' in parents and trashed=false",
            fields="files(id, name, size, md5Checksum), nextPageToken",
            pageSize=1000, pageToken=pt
        ).execute()
        files.extend(r.get("files", []))
        pt = r.get("nextPageToken")
        if not pt:
            break
    return files


def main():
    service = get_service()

    # Find perpres and inpres folder IDs
    r = service.files().list(
        q=f"'{ROOT}' in parents and mimeType='application/vnd.google-apps.folder' and trashed=false",
        fields="files(id, name)",
        pageSize=100
    ).execute()
    folders = {f["name"]: f["id"] for f in r.get("files", [])}

    for label in ["perpres", "inpres"]:
        fid = folders.get(label)
        if not fid:
            print(f"Folder '{label}' not found!")
            continue

        print(f"{'='*60}")
        print(f"  {label.upper()} — Folder ID: {fid}")
        print(f"{'='*60}")

        files = list_folder(service, fid)
        print(f"  Total files: {len(files)}")

        slugs = set()
        years = {}
        for f in files:
            m = SLUG_RE.match(f["name"].lower().replace(".pdf", ""))
            if m:
                slug = m.group(0)
                num = int(m.group(2))
                year = int(m.group(3))
                slugs.add(slug)
                years.setdefault(year, []).append(num)

        print(f"  Unique (num,year): {len(slugs)}")
        for y in sorted(years.keys()):
            nums = sorted(years[y])
            if len(nums) <= 10:
                print(f"    {y}: {nums}")
            else:
                print(f"    {y}: {len(nums)} files ({nums[0]}-{nums[-1]})")

        # Save
        out = {"folder_id": fid, "count": len(files), "slugs": sorted(slugs)}
        path = f"/home/z/my-project/download/{label}_inventory.json"
        with open(path, "w") as fp:
            json.dump(out, fp, indent=2)
        print(f"  Saved to {path}")

    # Also save folder IDs
    print(f"\nAll folder IDs: {json.dumps(folders, indent=2)}")


if __name__ == "__main__":
    main()