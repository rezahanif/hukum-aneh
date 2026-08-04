#!/usr/bin/env python3
"""Move misplaced Keppres and Inpres files from peraturan/ to their own Drive folders.

Uses batch requests for efficiency.
"""

import json
import time
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import BatchHttpRequest

# ── Config ──────────────────────────────────────────────────────────────────────
PERATURAN_FOLDER_ID = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"
TOKEN_PATH = "/home/z/my-project/upload/token.json"
MISPLACED_FILE = "/home/z/my-project/download/misplaced_keppres_inpres.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
BATCH_SIZE = 50  # Google batch limit


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def create_folder(service, name, parent_id):
    file_metadata = {
        "name": name,
        "mimeType": "application/vnd.google-apps.folder",
        "parents": [parent_id],
    }
    file = service.files().create(body=file_metadata, fields="id").execute()
    print(f"Created folder: {name} -> {file['id']}")
    return file["id"]


def move_files_batch(service, file_ids, source_folder_id, target_folder_id, label):
    """Move files using batch requests (much faster than individual calls)."""
    ok = 0
    fail = 0
    errors = []
    
    for batch_start in range(0, len(file_ids), BATCH_SIZE):
        batch = service.new_batch_http_request()
        batch_callbacks = {}
        
        chunk = file_ids[batch_start:batch_start + BATCH_SIZE]
        for fid in chunk:
            def callback(req_id, resp, exc, fid=fid):
                nonlocal ok, fail
                if exc:
                    fail += 1
                    errors.append(f"{fid[:12]}: {exc}")
                else:
                    ok += 1
            
            req_id = f"move_{fid}"
            batch_callbacks[req_id] = callback
            batch.add(service.files().update(
                fileId=fid,
                addParents=target_folder_id,
                removeParents=source_folder_id,
                fields="id",
            ), request_id=req_id, callback=callback)
        
        batch.execute()
        print(f"  {label} batch {batch_start//BATCH_SIZE + 1}: {min(batch_start + BATCH_SIZE, len(file_ids))}/{len(file_ids)} done (ok={ok}, fail={fail})")
    
    if errors:
        print(f"  Errors ({len(errors)}):")
        for e in errors[:10]:
            print(f"    {e}")
        if len(errors) > 10:
            print(f"    ... and {len(errors) - 10} more")
    
    return ok, fail


def main():
    service = get_service()

    with open(MISPLACED_FILE, "r") as f:
        data = json.load(f)
    keppres_files = data["keppres"]
    inpres_files = data["inpres"]
    print(f"Loaded {len(keppres_files)} Keppres + {len(inpres_files)} Inpres to move")

    # Step 1: Create folders
    print("\n=== Creating folders ===")
    keppres_folder_id = create_folder(service, "keppres", PERATURAN_FOLDER_ID)
    inpres_folder_id = create_folder(service, "inpres", PERATURAN_FOLDER_ID)

    # Step 2: Move Keppres
    print(f"\n=== Moving {len(keppres_files)} Keppres files ===")
    keppres_ids = [f["id"] for f in keppres_files]
    k_ok, k_fail = move_files_batch(service, keppres_ids, PERATURAN_FOLDER_ID, keppres_folder_id, "Keppres")

    # Step 3: Move Inpres
    print(f"\n=== Moving {len(inpres_files)} Inpres files ===")
    inpres_ids = [f["id"] for f in inpres_files]
    i_ok, i_fail = move_files_batch(service, inpres_ids, PERATURAN_FOLDER_ID, inpres_folder_id, "Inpres")

    # Summary
    total_ok = k_ok + i_ok
    total_fail = k_fail + i_fail
    total = len(keppres_files) + len(inpres_files)
    print(f"\n=== DONE ===")
    print(f"keppres folder: {keppres_folder_id}")
    print(f"inpres folder:  {inpres_folder_id}")
    print(f"Total: {total_ok}/{total} moved ({total_fail} failed)")

    folder_ids = {"keppres": keppres_folder_id, "inpres": inpres_folder_id}
    with open("/home/z/my-project/download/new_folder_ids.json", "w") as f:
        json.dump(folder_ids, f, indent=2)
    print("Folder IDs saved to download/new_folder_ids.json")


if __name__ == "__main__":
    main()
