import json, os, time
from googleapiclient.http import MediaFileUpload
from googleapiclient.discovery import build
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request

TOKEN_PATH = "/home/z/my-project/upload/google_token.json"
DEST_FOLDER_ID = "1Y14_wbozLWkV6V5yBGCkaNsPOWYQK0QZ"
SAMPLES_DIR = "/home/z/my-project/download/samples"


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
    token_data["token"] = creds.token
    with open(TOKEN_PATH, "w") as f:
        json.dump(token_data, f, indent=2)
    return creds


def get_drive_service():
    creds = get_credentials()
    return build("drive", "v3", credentials=creds)


def upload_file(service, file_path, dest_folder_id, name=None):
    file_name = name or os.path.basename(file_path)
    file_metadata = {"name": file_name, "parents": [dest_folder_id]}
    mime = "application/pdf"
    media = MediaFileUpload(file_path, mimetype=mime, resumable=True)
    request = service.files().create(
        body=file_metadata, media_body=media, fields="id, name, size"
    )
    result = None
    while result is None:
        _, result = request.next_chunk()
    print(f"  Uploaded: {result['name']} ({result.get('size', '?')} bytes)")
    return result


def main():
    service = get_drive_service()

    # Find all PDFs
    pdfs = []
    for root, dirs, files in os.walk(SAMPLES_DIR):
        for fn in files:
            if fn.lower().endswith('.pdf'):
                pdfs.append(os.path.join(root, fn))
    pdfs.sort()
    print(f"Found {len(pdfs)} PDFs to upload\n")

    for i, pdf_path in enumerate(pdfs):
        rel = os.path.relpath(pdf_path, SAMPLES_DIR)
        dirname = rel.split(os.sep)[0] if os.sep in rel else "root"
        safe_name = f"source_{dirname}_{os.path.basename(pdf_path)}"
        sz = os.path.getsize(pdf_path)
        print(f"[{i+1}/{len(pdfs)}] {safe_name} ({sz:,} bytes)")
        try:
            upload_file(service, pdf_path, DEST_FOLDER_ID, safe_name)
        except Exception as e:
            print(f"  ERROR: {e}")
            time.sleep(3)
            try:
                upload_file(service, pdf_path, DEST_FOLDER_ID, safe_name)
            except Exception as e2:
                print(f"  RETRY FAILED: {e2}")
        time.sleep(0.5)

    print(f"\n=== DONE: {len(pdfs)} PDFs uploaded ===")


if __name__ == "__main__":
    main()
