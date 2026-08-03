"""Google Drive helpers - list folders, download files, refresh token."""
import json, os, sys
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseDownload
import io

TOKEN_PATH = "/home/z/my-project/upload/google_token.json"
SECRET_PATH = "/home/z/my-project/upload/client_secret_858768057989-7ql7keonq637uoejbmh8ptl7fl9cc5h3.apps.googleusercontent.com.json"


def get_credentials():
    with open(TOKEN_PATH) as f:
        token_data = json.load(f)
    creds = Credentials(
        token=token_data["token"],
        refresh_token=token_data["refresh_token"],
        token_uri=token_data["token_uri"],
        client_id=token_data["client_id"],
        client_secret=token_data["client_secret"],
        scopes=token_data["scopes"],
    )
    creds.refresh(Request())
    # persist refreshed token
    token_data["token"] = creds.token
    with open(TOKEN_PATH, "w") as f:
        json.dump(token_data, f, indent=2)
    return creds


def get_drive_service():
    creds = get_credentials()
    return build("drive", "v3", credentials=creds)


def list_folder(service, folder_id, max_files=5):
    """List files in a folder, return list of {id, name} dicts."""
    results = []
    page_token = None
    while len(results) < max_files:
        resp = (
            service.files()
            .list(
                q=f"'{folder_id}' in parents and mimeType='application/pdf'",
                pageSize=min(100, max_files - len(results)),
                fields="nextPageToken, files(id, name)",
                pageToken=page_token,
            )
            .execute()
        )
        results.extend(resp.get("files", []))
        page_token = resp.get("nextPageToken")
        if not page_token:
            break
    return results[:max_files]


def download_file(service, file_id, dest_path):
    """Download a single file by ID to dest_path."""
    request = service.files().get_media(fileId=file_id)
    fh = io.FileIO(dest_path, "wb")
    downloader = MediaIoBaseDownload(fh, request)
    done = False
    while not done:
        _, done = downloader.next_chunk()
    return dest_path


def list_subfolders(service, parent_id):
    """List immediate subfolders under parent_id."""
    results = []
    page_token = None
    while True:
        resp = (
            service.files()
            .list(
                q=f"'{parent_id}' in parents and mimeType='application/vnd.google-apps.folder'",
                pageSize=100,
                fields="nextPageToken, files(id, name)",
                pageToken=page_token,
            )
            .execute()
        )
        results.extend(resp.get("files", []))
        page_token = resp.get("nextPageToken")
        if not page_token:
            break
    return results


if __name__ == "__main__":
    service = get_drive_service()
    ROOT = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"
    folders = list_subfolders(service, ROOT)
    for f in folders:
        print(f"{f['name']}  ->  {f['id']}")
