#!/usr/bin/env python3
"""Quick verify: count files after dedup and check for remaining dupes."""

import re
from collections import defaultdict, Counter
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
ROOT_FOLDER = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def main():
    service = get_service()
    all_files = []
    folder_counts = Counter()
    folders = [ROOT_FOLDER]
    folder_map = {ROOT_FOLDER: "root"}

    while folders:
        fid = folders.pop(0)
        pt = None
        while True:
            r = service.files().list(
                q=f"'{fid}' in parents and trashed=false",
                fields="files(id, name, size, md5Checksum, mimeType, parents), nextPageToken",
                pageSize=1000, pageToken=pt
            ).execute()
            for f in r.get("files", []):
                if f["mimeType"] == "application/vnd.google-apps.folder":
                    folders.append(f["id"])
                    folder_map[f["id"]] = f["name"]
                else:
                    p = f.get("parents", [None])[0]
                    folder_name = folder_map.get(p, "?")
                    folder_counts[folder_name] += 1
                    f["_folder"] = folder_name
                    all_files.append(f)
            pt = r.get("nextPageToken")
            if not pt:
                break

    # Check remaining same-folder dupes
    by_folder_md5 = defaultdict(list)
    for f in all_files:
        md5 = f.get("md5Checksum")
        parent = f.get("parents", [None])[0]
        if md5 and parent:
            by_folder_md5[(parent, md5)].append(f)

    remaining_dupes = {k: v for k, v in by_folder_md5.items() if len(v) > 1}

    total_size = sum(int(f.get("size", 0)) for f in all_files)
    print(f"Total files: {len(all_files)} (was 6,787)")
    print(f"Total size: {total_size / (1024*1024*1024):.2f} GB")
    print(f"Remaining same-folder dupes: {len(remaining_dupes)}")
    print(f"\nFiles per folder:")
    for folder, count in sorted(folder_counts.items(), key=lambda x: -x[1]):
        print(f"  {folder}: {count}")

    if remaining_dupes:
        print(f"\nRemaining dupes (sample):")
        for i, (k, v) in enumerate(sorted(remaining_dupes.items())):
            if i >= 5:
                break
            folder_name = folder_map.get(k[0], "?")
            print(f"  [{folder_name}] MD5:{k[1][:8]}: {len(v)} copies")
            for f in v:
                print(f"    {f['name'][:70]}")


if __name__ == "__main__":
    main()