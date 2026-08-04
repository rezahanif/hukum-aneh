#!/usr/bin/env python3
"""Move Keppres/Inpres in 3 steps: create folders, move inpres, move keppres chunk by chunk.
Usage: python3 move_in_3_chunks.py --step 1|2|3 [--chunk 0|1|2]
"""

import json, os, sys, time
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
PERATURAN_FOLDER = "1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U"
MISPLACED_PATH = "/home/z/my-project/download/misplaced_keppres_inpres.json"
STATE_PATH = "/home/z/my-project/download/move_state.json"

def get_drv():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, [
        "https://www.googleapis.com/auth/drive"
    ])
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build('drive', 'v3', credentials=creds)

def step1_create_folders():
    drv = get_drv()
    root = drv.files().get(fileId=PERATURAN_FOLDER, fields='parents').execute()
    root_parent = root['parents'][0]
    state = {'root_parent': root_parent}
    for name in ['keppres', 'inpres']:
        q = f"name='{name}' and '{root_parent}' in parents and trashed=false and mimeType='application/vnd.google-apps.folder'"
        existing = drv.files().list(q=q, fields='files(id)', pageSize=1).execute()
        if existing.get('files'):
            fid = existing['files'][0]['id']
            print(f"EXISTS {name}: {fid}")
        else:
            res = drv.files().create(body={
                'name': name,
                'mimeType': 'application/vnd.google-apps.folder',
                'parents': [root_parent]
            }, fields='id').execute()
            fid = res['id']
            print(f"CREATED {name}: {fid}")
        state[f'{name}_folder'] = fid
    with open(STATE_PATH, 'w') as f:
        json.dump(state, f, indent=2)
    print(f"State saved. {state}")

def step2_move_inpres():
    with open(STATE_PATH) as f:
        state = json.load(f)
    with open(MISPLACED_PATH) as f:
        data = json.load(f)
    drv = get_drv()
    ok = err = 0
    for f_info in data['inpres']:
        try:
            drv.files().update(
                fileId=f_info['id'],
                addParents=state['inpres_folder'],
                removeParents=PERATURAN_FOLDER,
                fields='id'
            ).execute()
            ok += 1
            print(f"  OK {f_info['name']}")
        except Exception as e:
            err += 1
            print(f"  ERR {f_info['name']}: {e}")
    print(f"Inpres: {ok} moved, {err} errors")

def step3_move_keppres(chunk=0, chunk_size=30):
    with open(STATE_PATH) as f:
        state = json.load(f)
    with open(MISPLACED_PATH) as f:
        data = json.load(f)
    drv = get_drv()
    all_keppres = data['keppres']
    start = chunk * chunk_size
    end = min(start + chunk_size, len(all_keppres))
    batch = all_keppres[start:end]
    print(f"Moving keppres chunk {chunk}: files {start}-{end-1} ({len(batch)} files)")
    ok = err = 0
    for f_info in batch:
        try:
            drv.files().update(
                fileId=f_info['id'],
                addParents=state['keppres_folder'],
                removeParents=PERATURAN_FOLDER,
                fields='id'
            ).execute()
            ok += 1
            print(f"  [{start+ok}] OK {f_info['name'][:60]}")
        except Exception as e:
            err += 1
            print(f"  ERR {f_info['name']}: {e}")
    print(f"Chunk {chunk}: {ok} moved, {err} errors")
    if chunk == 0 and err == 0:
        pass  # could auto-continue

import argparse
parser = argparse.ArgumentParser()
parser.add_argument('--step', type=int, required=True, help='1=create folders, 2=move inpres, 3=move keppres')
parser.add_argument('--chunk', type=int, default=0, help='keppres chunk number (0-based, 30 per chunk)')
args = parser.parse_args()

if args.step == 1:
    step1_create_folders()
elif args.step == 2:
    step2_move_inpres()
elif args.step == 3:
    step3_move_keppres(chunk=args.chunk)
