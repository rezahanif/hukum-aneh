#!/usr/bin/env python3
"""
Python scraper bridge for hukum-aneh.
Handles TLS-fingerprinted scraping for Indonesian government sources.

Receives JSON on stdin, returns JSON on stdout. Never writes to DB.

Input:  {"url": "...", "action": "check_updates|download|extract_metadata", "source": "..."}
Output: {"success": true, "documents": [...]} or {"success": false, "error": "..."}
"""
import sys
import json
import logging

logging.basicConfig(level=logging.INFO, format="%(asctime)s [scraper] %(message)s")
log = logging.getLogger(__name__)


def main():
    try:
        req = json.loads(sys.stdin.read())
    except Exception as e:
        print(json.dumps({"success": False, "error": f"invalid input: {e}"}))
        sys.exit(1)

    action = req.get("action", "")
    url = req.get("url", "")
    source = req.get("source", "")

    log.info(f"action={action} source={source} url={url}")

    # ponytail: stub — replace with curl_cffi/patchright when anti-bot detected
    # upgrade: add per-source fingerprint profiles
    try:
        if action == "check_updates":
            docs = check_updates(url, source)
            print(json.dumps({"success": True, "documents": docs}))
        elif action == "download":
            doc = download(url, source)
            print(json.dumps({"success": True, "data": doc}))
        elif action == "extract_metadata":
            meta = extract_metadata(url, source)
            print(json.dumps({"success": True, "data": meta}))
        else:
            print(json.dumps({"success": False, "error": f"unknown action: {action}"}))
            sys.exit(1)
    except Exception as e:
        log.error(f"failed: {e}")
        print(json.dumps({"success": False, "error": str(e)}))
        sys.exit(1)


def check_updates(url, source):
    """Poll source for new/changed laws. Returns list of document metadata."""
    # ponytail: stub — implement per-source parsing logic
    # each source has different HTML structure
    return []


def download(url, source):
    """Download raw document file (PDF/HTML). Returns content + metadata."""
    # ponytail: stub — use curl_cffi for TLS fingerprint when needed
    return {"content": "", "mime_type": "application/pdf", "filename": ""}


def extract_metadata(url, source):
    """Extract structured metadata from source listing page."""
    # ponytail: stub — per-source parser
    return {"law_number": "", "title": "", "published_date": ""}


if __name__ == "__main__":
    main()
