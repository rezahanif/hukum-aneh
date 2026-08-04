#!/usr/bin/env python3
"""Spot-check: download a few duplicate files and verify MD5 matches."""

import hashlib
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseDownload
import io

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
ROOT_FOLDER = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def download_and_hash(service, file_id, name, max_bytes=1024*1024):
    """Download first max_bytes of a file and compute partial MD5."""
    request = service.files().get_media(fileId=file_id)
    fh = io.BytesIO()
    downloader = MediaIoBaseDownload(fh, request, chunksize=256*1024)
    done = False
    h = hashlib.md5()
    while not done:
        status, done = downloader.next_chunk()
    content = fh.getvalue()
    # Full MD5
    full_md5 = hashlib.md5(content).hexdigest()
    # Partial MD5 (first 64KB)
    h.update(content[:65536])
    partial_md5 = h.hexdigest()
    return full_md5, len(content)


def main():
    service = get_service()

    # Find files that are claimed duplicates — same folder, same claimed MD5
    # Check a few specific cases
    test_cases = [
        # (folder_name, name_pattern)
        ("uu", "uu-no-7-tahun-2023"),
        ("uu", "uu-no-5-tahun-2022"),
        ("uu", "uu-no-4-tahun-2018"),
    ]

    for folder_name, name_pattern in test_cases:
        print(f"\n{'='*60}")
        print(f"  Checking: {name_pattern} in {folder_name}/")
        print(f"{'='*60}")

        # Find the folder
        results = service.files().list(
            q=f"name='{folder_name}' and '{ROOT_FOLDER}' in parents and mimeType='application/vnd.google-apps.folder' and trashed=false",
            fields="files(id, name)"
        ).execute()
        folders = results.get("files", [])
        if not folders:
            print(f"  Folder '{folder_name}' not found")
            continue
        folder_id = folders[0]["id"]

        # Find matching files
        results = service.files().list(
            q=f"name contains '{name_pattern}' and '{folder_id}' in parents and trashed=false",
            fields="files(id, name, size, md5Checksum)"
        ).execute()
        files = results.get("files", [])

        print(f"  Found {len(files)} files:")
        for f in sorted(files, key=lambda x: x["name"]):
            api_md5 = f.get("md5Checksum", "N/A")
            size = int(f.get("size", 0))
            print(f"    API MD5: {api_md5}  Size: {size:>8}  {f['name'][:70]}")

        # Download and verify first 2 files
        if len(files) >= 2:
            print(f"  \n  Downloading & verifying first 2 files...")
            for f in files[:2]:
                real_md5, dl_size = download_and_hash(service, f["id"], f["name"])
                match = "✓ MATCH" if real_md5 == api_md5 else "✗ MISMATCH!"
                print(f"    {f['name'][:50]}")
                print(f"      Downloaded: {dl_size} bytes, MD5: {real_md5}")
                print(f"      API claim:  {api_md5}  {match}")

    # Also check the perppu-no-148 in both folders
    print(f"\n{'='*60}")
    print(f"  Checking: perppu-no-148-tahun-2024 (PP vs Perppu folder)")
    print(f"{'='*60}")
    for folder_name in ["pp", "perppu"]:
        results = service.files().list(
            q=f"name contains 'perppu-no-148' and '{ROOT_FOLDER}' in parents and trashed=false",
            fields="files(id, name, size, md5Checksum)"
        ).execute()
        for f in results.get("files", []):
            # Check if it's in the right folder by getting parents
            parent_info = service.files().get(fileId=f["id"], fields="parents").execute()
            parent_id = parent_info.get("parents", ["?"])[0]
            parent_name = service.files().get(fileId=parent_id, fields="name").execute().get("name", "?")
            print(f"  [{parent_name}] MD5:{f.get('md5Checksum','?')[:12]}  {int(f.get('size',0))//1024}KB  {f['name']}")


if __name__ == "__main__":
    main()