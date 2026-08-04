#!/usr/bin/env python3
"""Dedup Drive folders - two-pass: collect IDs first, then batch delete.

Pass 1: Scan & compute deletion list (saved to JSON).
Pass 2: Delete from JSON (resumable)."""

import json
import re
import sys
import time
from collections import defaultdict
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import BatchHttpRequest

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
ROOT_FOLDER = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"
DELETE_PLAN_PATH = "/home/z/my-project/download/dedup_plan.json"


def name_score(name):
    n = name.lower()
    if ".pdf" not in n:
        return 0
    stem = n.replace(".pdf", "")
    if re.match(r'^(uu|uud|perppu|pp|perpres|keppres|inpres|tap-mpr)-no-\d+-tahun-\d+_.+', stem):
        return 100
    if re.match(r'^(uu|uud|perppu|pp|perpres|keppres|inpres|tap-mpr)-no-\d+-tahun-\d+$', stem):
        return 50
    if re.match(r'^uu\d+-\d{4}\.pdf$', n):
        return 10
    if re.match(r'^uu\d+-\d{4}bt\.pdf$', n):
        return 5
    return 20


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def collect_files(service, root_id):
    all_files = []
    folder_map = {root_id: "root"}
    folders = [root_id]
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
                    f["_folder"] = folder_map.get(p, "?")
                    f["_parent"] = p
                    all_files.append(f)
            pt = r.get("nextPageToken")
            if not pt:
                break
    return all_files


def pass1_scan():
    """Scan and save deletion plan."""
    service = get_service()
    print("Scanning all files...")
    all_files = collect_files(service, ROOT_FOLDER)
    print(f"Total: {len(all_files)}")

    # Group by (parent, md5)
    groups = defaultdict(list)
    for f in all_files:
        md5 = f.get("md5Checksum")
        parent = f.get("_parent")
        if md5 and parent:
            groups[(parent, md5)].append(f)

    dup_groups = {k: v for k, v in groups.items() if len(v) > 1}
    print(f"Duplicate groups: {len(dup_groups)}")

    to_delete = []
    savings = 0
    for (parent_id, md5), files in dup_groups.items():
        ranked = sorted(files, key=lambda f: (name_score(f["name"]), len(f["name"])), reverse=True)
        keep = ranked[0]
        for d in ranked[1:]:
            to_delete.append({
                "id": d["id"],
                "name": d["name"],
                "folder": d["_folder"],
                "size": d.get("size", "0"),
                "keep_name": keep["name"],
            })
            savings += int(d.get("size", 0))

    plan = {
        "total_files": len(all_files),
        "dup_groups": len(dup_groups),
        "to_delete_count": len(to_delete),
        "savings_bytes": savings,
        "savings_mb": round(savings / (1024*1024), 1),
        "deletions": to_delete,
    }

    with open(DELETE_PLAN_PATH, "w") as fp:
        json.dump(plan, fp, indent=2, ensure_ascii=False)

    print(f"\nPlan saved: {len(to_delete)} files, ~{plan['savings_mb']} MB to free")
    print(f"Use: python3 {sys.argv[0]} --delete")

    # Show per-folder breakdown
    by_folder = defaultdict(int)
    for d in to_delete:
        by_folder[d["folder"]] += 1
    print(f"\nBy folder:")
    for folder, count in sorted(by_folder.items(), key=lambda x: -x[1]):
        print(f"  {folder}: {count}")


def pass2_delete():
    """Execute deletions from saved plan."""
    with open(DELETE_PLAN_PATH) as fp:
        plan = json.load(fp)

    deletions = plan["deletions"]
    total = len(deletions)
    print(f"Deleting {total} files (~{plan['savings_mb']} MB)...")

    # Check for already-done (resume support)
    done_file = DELETE_PLAN_PATH + ".done"
    done_ids = set()
    if __import__("os").path.exists(done_file):
        with open(done_file) as fp:
            done_ids = set(json.load(fp))
        print(f"Resuming - {len(done_ids)} already deleted")

    service = get_service()
    ok = 0
    fail = 0
    skip = 0

    for i, d in enumerate(deletions):
        fid = d["id"]
        if fid in done_ids:
            skip += 1
            continue

        try:
            service.files().delete(fileId=fid).execute()
            ok += 1
            done_ids.add(fid)
        except Exception as e:
            err = str(e)[:80]
            print(f"  FAIL {d['name'][:40]}: {err}")
            fail += 1

        if (ok + fail) % 30 == 0:
            print(f"  ... {ok+skip}/{total} (ok={ok}, fail={fail}, skip={skip})")
            time.sleep(1)
            # Save progress
            with open(done_file, "w") as fp:
                json.dump(list(done_ids), fp)

    # Final save
    with open(done_file, "w") as fp:
        json.dump(list(done_ids), fp)

    print(f"\nDONE: {ok} deleted, {fail} failed, {skip} skipped")
    print(f"Space freed: ~{plan['savings_mb']} MB")


if __name__ == "__main__":
    if "--delete" in sys.argv:
        pass2_delete()
    else:
        pass1_scan()
