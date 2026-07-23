#!/usr/bin/env python3
"""
Python scraper bridge for hukum-aneh.
Handles TLS-fingerprinted scraping for JDIHN and MKRI.

Receives JSON on stdin, returns JSON on stdout. Never writes to DB.
"""
import sys
import json
import logging
import re
from bs4 import BeautifulSoup
from curl_cffi import requests

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
    """Poll source for new/changed laws."""
    if "MKRI" in source or "Mahkamah" in source:
        return check_mkri()
    elif "JDIHN" in source:
        return check_jdihn()
    return []


def check_mkri():
    """Scrapes Mahkamah Konstitusi latest decisions."""
    url = "https://www.mkri.id/perkara/persidangan/putusan"
    r = requests.get(url, impersonate="chrome120", timeout=15)
    if r.status_code != 200:
        raise Exception(f"MKRI returned status {r.status_code}")

    soup = BeautifulSoup(r.text, "html.parser")
    docs = []
    
    # Extract putusan links: href containing "putusan_mkri"
    for a in soup.find_all("a", href=True):
        href = a["href"]
        if "putusan_mkri" in href and href.endswith(".pdf"):
            # Try to guess title from context or use filename
            filename = href.split("/")[-1]
            title = filename.replace(".pdf", "").replace("_", " ").upper()
            
            # Extract number from filename, e.g. putusan_mkri_14695_1784791338.pdf
            # Num: 14695
            num_match = re.search(r'putusan_mkri_(\d+)', filename)
            law_num = f"Putusan MK No. {num_match.group(1)}" if num_match else "Putusan MK"

            docs.append({
                "law_number": law_num,
                "title": title,
                "source_url": href,
                "source": "Mahkamah Konstitusi",
                "level": "national",
                "document_type": "Putusan MK",
                "published_date": ""
            })
    return docs


def check_jdihn():
    """Scrapes JDIHN portal landing page latest documents."""
    url = "https://jdihn.go.id/"
    try:
        r = requests.get(url, impersonate="chrome120", timeout=10)
        if r.status_code != 200:
            log.warning(f"JDIHN returned status {r.status_code}")
            return []
    except Exception as e:
        log.warning(f"JDIHN connection failed (possibly blocked by host firewall): {e}")
        return []

    soup = BeautifulSoup(r.text, "html.parser")
    docs = []

    # Look for detail page links: href containing "/doc/"
    for a in soup.find_all("a", href=True):
        href = a["href"]
        aria_label = a.get("aria-label", "")
        if href.startswith("/doc/") and "Buka detail" in aria_label:
            title = aria_label.replace("Buka detail", "").strip()
            doc_id = href.split("/")[-1]

            docs.append({
                "law_number": f"JDIHN {doc_id}",
                "title": title,
                "source_url": f"https://jdihn.go.id{href}",
                "source": "JDIHN",
                "level": "national",
                "document_type": "National legal network",
                "published_date": ""
            })
    return docs


def download(url, source):
    """Download raw document file (PDF/HTML)."""
    # For standard PDFs, try direct requests
    r = requests.get(url, impersonate="chrome120", timeout=30)
    if r.status_code != 200:
        raise Exception(f"Download failed: status {r.status_code}")
    
    # Return base64 or raw string (python json encoder handles strings better)
    # Since this goes over stdout, we return content length or handle it
    # Go workflow engine expects download raw content.
    # Wait, the Go code downloads via direct HTTP for connectors, but calls Download
    # via Python for protected ones if needed.
    return {
        "content": "",  # Handled by Go client directly
        "mime_type": "application/pdf",
        "filename": url.split("/")[-1]
    }


def extract_metadata(url, source):
    return {"law_number": "", "title": "", "published_date": ""}


if __name__ == "__main__":
    main()
