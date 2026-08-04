#!/usr/bin/env python3
"""Resume dedup using batch HTTP requests (much faster)."""

import json
import os
import time
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build
from googleapiclient.http import BatchHttpRequest

TOKEN_PATH = "/home/z/my-project/upload/token.json"
PLAN_PATH = "/home/z/my-project/download/dedup_plan.json"
DONE_PATH = PLAN_PATH + ".done"
SCOPES = ["https://www.googleapis.com/auth/drive"]
BATCH_SIZE = 50


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def main():
    with open(PLAN_PATH) as f:
        plan = json.load(f)
    deletions = plan["deletions"]
    total = len(deletions)

    # Load progress
    done_ids = set()
    if os.path.exists(DONE_PATH):
        with open(DONE_PATH) as f:
            done_ids = set(json.load(f))
    remaining = [d for d in deletions if d["id"] not in done_ids]
    print(f"Total: {total}, Done: {len(done_ids)}, Remaining: {len(remaining)}")

    if not remaining:
        print("All done!")
        return

    service = get_service()
    ok = 0
    fail = 0
    errors = []

    for batch_start in range(0, len(remaining), BATCH_SIZE):
        batch = service.new_batch_http_request()
        chunk = remaining[batch_start:batch_start + BATCH_SIZE]
        batch_ok = 0
        batch_fail = 0

        for d in chunk:
            fid = d["id"]
            def callback(req_id, resp, exc, fid=fid, d=d):
                nonlocal ok, fail
                if exc:
                    fail += 1
                    if len(errors) < 10:
                        errors.append(f"{d['name'][:40]}: {str(exc)[:60]}")
                else:
                    ok += 1
                    done_ids.add(fid)
                    batch_ok += 1

            batch.add(service.files().delete(fileId=fid), request_id=fid, callback=callback)

        batch.execute()
        print(f"  Batch {batch_start//BATCH_SIZE+1}: +{batch_ok} ok (total ok={ok}, fail={fail}, left={len(remaining)-batch_start-BATCH_SIZE})")

        # Save progress after each batch
        with open(DONE_PATH, "w") as f:
            json.dump(list(done_ids), f)

        time.sleep(0.5)

    print(f"\nDONE: {ok} deleted, {fail} failed")
    print(f"Space freed: ~{plan['savings_mb']} MB")
    if errors:
        print("Errors:")
        for e in errors:
            print(f"  {e}")


if __name__ == "__main__":
    main()
