#!/usr/bin/env python3
"""Upload v3 chunk test results to Google Drive."""
import json
import os

from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build
from googleapiclient.http import MediaFileUpload
from googleapiclient.errors import HttpError

# Load token
with open('/home/z/my-project/upload/google_token.json') as f:
    token_data = json.load(f)

creds = Credentials(
    token=token_data['token'],
    refresh_token=token_data['refresh_token'],
    token_uri=token_data['token_uri'],
    client_id=token_data['client_id'],
    client_secret=token_data['client_secret'],
    scopes=token_data['scopes'],
)

service = build('drive', 'v3', credentials=creds)

# Files to upload
FILES = [
    {
        'local': '/home/z/my-project/download/chunk_results_v3.json',
        'remote': 'chunk_results_v3.json',
        'mime': 'application/json',
        'desc': 'Full chunk data for all 16 samples (974 chunks)',
    },
    {
        'local': '/home/z/my-project/download/chunk_test_report_v3.json',
        'remote': 'chunk_test_report_v3.json',
        'mime': 'application/json',
        'desc': 'Per-file QA summary with fixes applied and known issues',
    },
    {
        'local': '/home/z/my-project/download/chunk_qa_report_v3.json',
        'remote': 'chunk_qa_report_v3.json',
        'mime': 'application/json',
        'desc': 'Raw QA metrics (duplicate IDs, stamp noise, errors)',
    },
]

uploaded = []

for f_info in FILES:
    local_path = f_info['local']
    if not os.path.exists(local_path):
        print(f'SKIP {f_info["remote"]}: file not found')
        continue

    file_metadata = {
        'name': f_info['remote'],
        'description': f_info['desc'],
    }
    media = MediaFileUpload(local_path, mimetype=f_info['mime'], resumable=True)

    try:
        file = service.files().create(
            body=file_metadata,
            media_body=media,
            fields='id,name,size,webViewLink'
        ).execute()
        
        size_kb = os.path.getsize(local_path) / 1024
        print(f'UPLOADED: {file["name"]} ({size_kb:.1f} KB)')
        print(f'  ID: {file["id"]}')
        print(f'  Link: {file["webViewLink"]}')
        uploaded.append(file)
    except HttpError as e:
        print(f'ERROR uploading {f_info["remote"]}: {e}')

print(f'\nDone: {len(uploaded)}/{len(FILES)} files uploaded to Google Drive')
