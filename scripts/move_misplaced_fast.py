#!/usr/bin/env python3
"""move_misplaced_fast.py — Move misplaced Keppres/Inpres from peraturan/ to own folders.
Uses batch approach for speed.
"""

import json, os, sys

from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
PERATURAN_FOLDER = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"
MISPLACED_PATH = "/home/z/my-project/download/misplaced_keppres_inpres.json"


def main():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, [
        "https://www.googleapis.com/auth/drive"
    ])
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    drv = build('drive', 'v3', credentials=creds)

    with open(MISPLACED_PATH) as f:
        data = json.load(f)

    # Get root parent of peraturan/
    root = drv.files().get(fileId=PERATURAN_FOLDER, fields='parents').execute()
    root_parent = root['parents'][0]

    # Create folders
    for name in ['keppres', 'inpres']:
        q = f"name='{name}' and '{root_parent}' in parents and trashed=false and mimeType='application/vnd.google-apps.folder'"
        existing = drv.files().list(q=q, fields='files(id)', pageSize=1).execute()
        if existing.get('files'):
            fid = existing['files'][0]['id']
            print(f"  Folder '{name}' exists: {fid}")
        else:
            res = drv.files().create(body={
                'name': name,
                'mimeType': 'application/vnd.google-apps.folder',
                'parents': [root_parent]
            }, fields='id').execute()
            fid = res['id']
            print(f"  Created '{name}': {fid}")
        if name == 'keppres':
            keppres_folder = fid
        else:
            inpres_folder = fid

    # Move Keppres (157 files)
    print(f"\nMoving {len(data['keppres'])} Keppres...")
    ok = err = 0
    for i, f in enumerate(data['keppres']):
        try:
            drv.files().update(
                fileId=f['id'],
                addParents=keppres_folder,
                removeParents=PERATURAN_FOLDER,
                fields='id'
            ).execute()
            ok += 1
        except Exception as e:
            err += 1
            if err <= 3:
                print(f"  ERR {f['name']}: {e}")
        if (i+1) % 50 == 0:
            print(f"  ... {i+1}/{len(data['keppres'])} (ok={ok}, err={err})")
    print(f"  Keppres done: {ok} moved, {err} errors")

    # Move Inpres (3 files)
    print(f"\nMoving {len(data['inpres'])} Inpres...")
    ok2 = err2 = 0
    for f in data['inpres']:
        try:
            drv.files().update(
                fileId=f['id'],
                addParents=inpres_folder,
                removeParents=PERATURAN_FOLDER,
                fields='id'
            ).execute()
            ok2 += 1
        except Exception as e:
            err2 += 1
            print(f"  ERR {f['name']}: {e}")
    print(f"  Inpres done: {ok2} moved, {err2} errors")

    # Quick verify
    print(f"\n--- Verify ---")
    for label, fid in [('peraturan', PERATURAN_FOLDER), ('keppres', keppres_folder), ('inpres', inpres_folder)]:
        count = 0
        pt = None
        while True:
            r = drv.files().list(q=f"'{fid}' in parents and trashed=false", fields='files(id),nextPageToken', pageSize=1000, pageToken=pt).execute()
            count += len(r.get('files', []))
            pt = r.get('nextPageToken')
            if not pt: break
        print(f"  {label}/: {count} files")

    print(f"\nNew folder IDs: keppres={keppres_folder} inpres={inpres_folder}")


if __name__ == '__main__':
    main()
