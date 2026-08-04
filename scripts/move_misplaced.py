#!/usr/bin/env python3
"""
move_misplaced.py — Create keppres/ and inpres/ Drive folders, move misplaced files.
"""

import json
import os
import time

from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
PERATURAN_FOLDER = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"
DRIVE_SCOPES = ["https://www.googleapis.com/auth/drive"]

MISPLACED_PATH = "/home/z/my-project/download/misplaced_keppres_inpres.json"


def authenticate():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, DRIVE_SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return creds


def create_folder(drv, name, parent_id):
    """Create folder if not exists, return folder ID."""
    # Check if already exists
    q = f"name='{name}' and '{parent_id}' in parents and trashed=false and mimeType='application/vnd.google-apps.folder'"
    existing = drv.files().list(q=q, fields="files(id)", pageSize=1).execute()
    if existing.get("files"):
        fid = existing["files"][0]["id"]
        print(f"  Folder '{name}' already exists: {fid}")
        return fid

    file_metadata = {
        "name": name,
        "mimeType": "application/vnd.google-apps.folder",
        "parents": [parent_id]
    }
    result = drv.files().create(
        body=file_metadata,
        fields="id"
    ).execute()
    print(f"  Created folder '{name}': {result['id']}")
    return result["id"]


def move_file(drv, file_id, old_parent, new_parent):
    """Move file from old_parent to new_parent."""
    drv.files().update(
        fileId=file_id,
        addParents=new_parent,
        removeParents=old_parent,
        fields="id, name"
    ).execute()


def main():
    creds = authenticate()
    drv = build("drive", "v3", credentials=creds)

    # Load misplaced file data
    with open(MISPLACED_PATH) as f:
        data = json.load(f)

    # Get root folder info (parent of peraturan/)
    root = drv.files().get(
        fileId=PERATURAN_FOLDER,
        fields="id, name, parents"
    ).execute()
    root_parent = root["parents"][0]
    print(f"Root Drive folder: {root_parent}")
    print(f"Source folder: {root['name']} ({PERATURAN_FOLDER})")

    # Create destination folders
    print("\n[1/3] Creating folders...")
    keppres_folder = create_folder(drv, "keppres", root_parent)
    inpres_folder = create_folder(drv, "inpres", root_parent)

    # Move Keppres files
    print(f"\n[2/3] Moving {len(data['keppres'])} Keppres files...")
    keppres_ok = 0
    keppres_err = 0
    for i, f in enumerate(data["keppres"]):
        try:
            move_file(drv, f["id"], PERATURAN_FOLDER, keppres_folder)
            keppres_ok += 1
            if (i + 1) % 50 == 0:
                print(f"  ... moved {i+1}/{len(data['keppres'])}")
        except Exception as e:
            print(f"  ERROR moving {f['name']}: {e}")
            keppres_err += 1
    print(f"  Keppres: {keppres_ok} moved, {keppres_err} errors")

    # Move Inpres files
    print(f"\n[3/3] Moving {len(data['inpres'])} Inpres files...")
    inpres_ok = 0
    inpres_err = 0
    for f in data["inpres"]:
        try:
            move_file(drv, f["id"], PERATURAN_FOLDER, inpres_folder)
            inpres_ok += 1
        except Exception as e:
            print(f"  ERROR moving {f['name']}: {e}")
            inpres_err += 1
    print(f"  Inpres: {inpres_ok} moved, {inpres_err} errors")

    # Verify
    print("\n--- Verification ---")
    for label, fid in [("keppres", keppres_folder), ("inpres", inpres_folder), ("peraturan", PERATURAN_FOLDER)]:
        count = 0
        page_token = None
        while True:
            res = drv.files().list(
                q=f"'{fid}' in parents and trashed=false",
                fields="files(id), nextPageToken",
                pageSize=1000,
                pageToken=page_token
            ).execute()
            count += len(res.get("files", []))
            page_token = res.get("nextPageToken")
            if not page_token:
                break
        print(f"  {label}/: {count} files")

    print(f"\nDone! New folder IDs:")
    print(f"  keppres: {keppres_folder}")
    print(f"  inpres:  {inpres_folder}")


if __name__ == "__main__":
    main()
