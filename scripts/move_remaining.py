#!/usr/bin/env python3"""Continue moving remaining Keppres/Inpres files (retry script)."""

import json
import time
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import BatchHttpRequest

TOKEN_PATH = "/home/z/my-project/upload/token.json"
MISPLACED_FILE = "/home/z/my-project/download/misplaced_keppres_inpres.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]

KEPPRES_FOLDER = "14i_b17iwUGN7tDB6AhIN2nbxZZjAZxRi"
INPRES_FOLDER = "1KnQsVfO06J6IGudhBbpsX47dC26QeW3q"
PERATURAN_FOLDER = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"
BATCH_SIZE = 20  # smaller batches to avoid timeout


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def is_already_moved(service, file_id, expected_parent):
    """Check if a file is already in the target folder."""
    try:
        f = service.files().get(fileId=file_id, fields="parents").execute()
        return expected_parent in f.get("parents", [])
    except:
        return False


def move_files_batch(service, file_ids, source_id, target_id, label):
    ok = 0
    fail = 0
    skip = 0
    errors = []

    for batch_start in range(0, len(file_ids), BATCH_SIZE):
        batch = service.new_batch_http_request()
        chunk = file_ids[batch_start:batch_start + BATCH_SIZE]

        # Skip already-moved files
        to_move = []
        for fid in chunk:
            if is_already_moved(service, fid, target_id):
                skip += 1
                print(f"  skip {fid[:12]} (already in target)")
            else:
                to_move.append(fid)

        if not to_move:
            print(f"  {label} batch {batch_start//BATCH_SIZE + 1}: all {len(chunk)} already moved")
            continue

        for fid in to_move:
            def callback(req_id, resp, exc, fid=fid):
                nonlocal ok, fail
                if exc:
                    fail += 1
                    errors.append(f"{fid[:12]}: {exc}")
                else:
                    ok += 1

            batch.add(service.files().update(
                fileId=fid,
                addParents=target_id,
                removeParents=source_id,
                fields="id",
            ), request_id=f"move_{fid}", callback=callback)

        batch.execute()
        print(f"  {label} batch {batch_start//BATCH_SIZE + 1}: +{ok} moved (total ok={ok}, fail={fail}, skip={skip})")
        time.sleep(1)  # pause between batches

    return ok, fail, skip


def main():
    service = get_service()

    with open(MISPLACED_FILE, "r") as f:
        data = json.load(f)
    keppres_ids = [f["id"] for f in data["keppres"]]
    inpres_ids = [f["id"] for f in data["inpres"]]

    print(f"=== Moving remaining Keppres ({len(keppres_ids)} total) ===")
    k_ok, k_fail, k_skip = move_files_batch(service, keppres_ids, PERATURAN_FOLDER, KEPPRES_FOLDER, "Keppres")

    print(f"\n=== Moving Inpres ({len(inpres_ids)} total) ===")
    i_ok, i_fail, i_skip = move_files_batch(service, inpres_ids, PERATURAN_FOLDER, INPRES_FOLDER, "Inpres")

    print(f"\n=== RESULT ===")
    print(f"Keppres: {k_ok} moved, {k_skip} skipped, {k_fail} failed")
    print(f"Inpres:  {i_ok} moved, {i_skip} skipped, {i_fail} failed")


if __name__ == "__main__":
    main()
