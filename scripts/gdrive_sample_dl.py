import json, os, subprocess, sys, time
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_FILE = "/home/z/my-project/upload/google_token.json"
ROOT_ID = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"
DL_DIR = "/home/z/my-project/download/samples"

with open(TOKEN_FILE) as f: d=json.load(f)
c=Credentials(token=d["token"],refresh_token=d.get("refresh_token"),token_uri=d["token_uri"],client_id=d["client_id"],client_secret=d["client_secret"],scopes=d["scopes"])
if c.expired: c.refresh(Request())
token = c.token
svc = build("drive", "v3", credentials=c)

folders = svc.files().list(q=f"'{ROOT_ID}' in parents and mimeType='application/vnd.google-apps.folder'",fields="files(id,name)",pageSize=100).execute().get("files",[])

for folder in folders:
    fname = folder["name"]
    files = svc.files().list(q=f"'{folder['id']}' in parents and mimeType != 'application/vnd.google-apps.folder'",fields="files(id,name,size)",pageSize=5).execute().get("files",[])
    if not files: continue
    # Pick a small-ish file (first one)
    pick = min(files, key=lambda f: int(f.get("size",999999999)))
    dest = os.path.join(DL_DIR, fname, pick["name"])
    os.makedirs(os.path.dirname(dest), exist_ok=True)
    print(f"{fname}: {pick['name']} ({int(pick.get('size',0))/1024:.0f} KB)", flush=True)
    r = subprocess.run(["curl","-sf","-H",f"Authorization: Bearer {token}","-o",dest,f"https://www.googleapis.com/drive/v3/files/{pick['id']}?alt=media"],capture_output=True,timeout=60)
    if r.returncode==0: print(f"  OK")
    else: print(f"  FAIL: {r.returncode}")
print("Done sampling")
