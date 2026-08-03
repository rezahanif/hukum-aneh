import json,os,subprocess
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

with open('/home/z/my-project/upload/google_token.json') as f: d=json.load(f)
c=Credentials(token=d['token'],refresh_token=d.get('refresh_token'),token_uri=d['token_uri'],client_id=d['client_id'],client_secret=d['client_secret'],scopes=d['scopes'])
if c.expired: c.refresh(Request())
token=c.token
svc=build('drive','v3',credentials=c)

# Get a Putusan-MK with more substance (larger file)
folder_id='18r2ZE9_R5zg72HToTBXwe3qH7Qr9Bidj'
files=svc.files().list(q=f"'{folder_id}' in parents and mimeType != 'application/vnd.google-apps.folder'",fields='files(id,name,size)',pageSize=100).execute().get('files',[])
# Pick one that's 500KB-2MB for substance
pick = [f for f in files if 500000 < int(f.get('size',0)) < 2000000][:1]
if not pick: pick = files[:1]
if pick:
    f=pick[0]
    dest='/home/z/my-project/download/samples/Putusan-MK/mk_substantive.pdf'
    os.makedirs(os.path.dirname(dest),exist_ok=True)
    print(f"Downloading {f['name']} ({int(f.get('size',0))/1024:.0f} KB)")
    r=subprocess.run(['curl','-sf','-H',f'Authorization: Bearer {token}','-o',dest,f"https://www.googleapis.com/drive/v3/files/{f['id']}?alt=media"],capture_output=True,timeout=60)
    print('OK' if r.returncode==0 else f'FAIL {r.returncode}')
