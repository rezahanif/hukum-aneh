import json, os, time, sys
from googleapiclient.http import MediaFileUpload
from googleapiclient.discovery import build
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request

TOKEN_PATH = "/home/z/my-project/upload/google_token.json"
DEST_FOLDER_ID = "1Y14_wbozLWkV6V5yBGCkaNsPOWYQK0QZ"
CHUNKS_PATH = "/home/z/my-project/download/chunk_results_v2.json"
OUTPUT_DIR = "/home/z/my-project/download/chunks_for_upload"


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
    """Upload a single file, return file ID."""
    file_name = name or os.path.basename(file_path)
    file_metadata = {"name": file_name, "parents": [dest_folder_id]}
    mime = "application/json"
    media = MediaFileUpload(file_path, mimetype=mime, resumable=True)
    request = service.files().create(
        body=file_metadata, media_body=media, fields="id, name, size"
    )
    result = None
    while result is None:
        _, result = request.next_chunk()
    print(f"  Uploaded: {result['name']} ({result.get('size', '?')} bytes) -> {result['id']}")
    return result


def main():
    os.makedirs(OUTPUT_DIR, exist_ok=True)
    service = get_drive_service()

    # Load chunks
    with open(CHUNKS_PATH) as f:
        all_chunks = json.load(f)

    dirs = list(all_chunks.keys())
    print(f"Directories: {len(dirs)}, Total chunks: {sum(len(v) for v in all_chunks.values())}")

    # Phase 1: Upload per-directory JSON files (small, avoid rate limit)
    print("\n=== PHASE 1: Per-directory JSON files ===")
    for i, dirname in enumerate(dirs):
        chunks = all_chunks[dirname]
        out_path = os.path.join(OUTPUT_DIR, f"chunks_{dirname}.json")
        with open(out_path, "w", encoding="utf-8") as f:
            json.dump({
                "directory": dirname,
                "chunk_count": len(chunks),
                "chunks": chunks
            }, f, ensure_ascii=False, indent=2)
        file_size = os.path.getsize(out_path)
        print(f"\n[{i+1}/{len(dirs)}] {dirname}: {len(chunks)} chunks ({file_size:,} bytes)")
        try:
            upload_file(service, out_path, DEST_FOLDER_ID, f"chunks_{dirname}.json")
        except Exception as e:
            print(f"  ERROR: {e}")
            # retry once after 5s
            time.sleep(5)
            try:
                upload_file(service, out_path, DEST_FOLDER_ID, f"chunks_{dirname}.json")
            except Exception as e2:
                print(f"  RETRY FAILED: {e2}")
        # Small delay to avoid rate limit
        time.sleep(1)

    # Phase 2: Upload full combined file
    print("\n=== PHASE 2: Full combined JSON ===")
    full_path = os.path.join(OUTPUT_DIR, "chunks_ALL_combined.json")
    with open(full_path, "w", encoding="utf-8") as f:
        json.dump({
            "total_chunks": sum(len(v) for v in all_chunks.values()),
            "directories": {k: len(v) for k, v in all_chunks.items()},
            "chunks": all_chunks
        }, f, ensure_ascii=False, indent=2)
    file_size = os.path.getsize(full_path)
    print(f"Full file: {file_size:,} bytes")
    try:
        upload_file(service, full_path, DEST_FOLDER_ID, "chunks_ALL_combined.json")
    except Exception as e:
        print(f"ERROR uploading full file: {e}")
        time.sleep(5)
        try:
            upload_file(service, full_path, DEST_FOLDER_ID, "chunks_ALL_combined.json")
        except Exception as e2:
            print(f"RETRY FAILED: {e2}")

    # Phase 3: Upload a summary TXT for quick reference
    print("\n=== PHASE 3: Summary text ===")
    summary_path = os.path.join(OUTPUT_DIR, "chunks_summary.txt")
    with open(summary_path, "w", encoding="utf-8") as f:
        f.write(f"Chunk Results Summary\n")
        f.write(f"Generated: {time.strftime('%Y-%m-%d %H:%M:%S')}\n")
        f.write(f"\nTotal chunks: {sum(len(v) for v in all_chunks.values())}\n")
        f.write(f"Directories: {len(dirs)}\n")
        f.write(f"\n")
        for dirname in dirs:
            chunks = all_chunks[dirname]
            f.write(f"--- {dirname}: {len(chunks)} chunks ---\n")
            for j, c in enumerate(chunks):
                text_preview = c.get('text', '')[:120].replace('\n', ' ')
                f.write(f"  [{j+1}] id={c.get('id','?')}")
                f.write(f" | {text_preview}...\n")
    file_size = os.path.getsize(summary_path)
    print(f"Summary: {file_size:,} bytes")
    try:
        upload_file(service, summary_path, DEST_FOLDER_ID, "chunks_summary.txt")
    except Exception as e:
        print(f"ERROR: {e}")

    print("\n=== DONE ===")


if __name__ == "__main__":
    main()
