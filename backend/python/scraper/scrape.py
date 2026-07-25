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
    law_number = req.get("law_number", "")

    log.info(f"action={action} source={source} url={url}")

    try:
        if action == "check_updates":
            docs = check_updates(url, source)
            print(json.dumps({"success": True, "documents": docs}))
        elif action == "download":
            doc = download(url, source)
            print(json.dumps({"success": True, "data": doc}))
        elif action == "search_bpk":
            result = search_bpk(url, law_number)
            print(json.dumps({"success": True, "data": result}))
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
    elif "LKPP" in source:
        return check_lkpp()
    elif "DPR" in source:
        return check_dpr()
    return []


def check_dpr():
    """Scrapes JDIH DPR RI latest regulations."""
    url = "https://jdih.dpr.go.id/"
    try:
        r = requests.get(url, impersonate="chrome120", timeout=15)
        if r.status_code != 200:
            log.warning(f"DPR returned status {r.status_code}")
            return []
    except Exception as e:
        log.warning(f"DPR connection failed: {e}")
        return []

    soup = BeautifulSoup(r.text, "html.parser")
    docs = []
    
    # Simple search for Keppres / Inpres on landing or basic links
    for a in soup.find_all("a", href=True):
        href = a["href"]
        title = a.text.strip()
        if not title:
            continue
            
        is_keppres = "keppres" in href.lower() or "keputusan presiden" in title.lower()
        is_inpres = "inpres" in href.lower() or "instruksi presiden" in title.lower()
        
        if is_keppres or is_inpres:
            doc_type = "Keputusan Presiden (Keppres)" if is_keppres else "Instruksi Presiden (Inpres)"
            
            # Simple number extraction
            num_match = re.search(r'(?:nomor|no\.?)\s+(\d+)\s+tahun\s+(\d+)', title.lower())
            if num_match:
                prefix = "KEPPRES" if is_keppres else "INPRES"
                law_num = f"{prefix} No. {num_match.group(1)} Tahun {num_match.group(2)}"
            else:
                law_num = title

            docs.append({
                "law_number": law_num,
                "title": title,
                "source_url": href if href.startswith("http") else f"https://jdih.dpr.go.id{href}",
                "source": "JDIH DPR RI",
                "level": "national",
                "document_type": doc_type,
                "published_date": ""
            })
    return docs


def check_lkpp():
    """Scrapes JDIH LKPP latest regulations."""
    url = "https://jdih.lkpp.go.id/regulation/index"
    try:
        r = requests.get(url, impersonate="chrome120", timeout=15, verify=False)
        if r.status_code != 200:
            log.warning(f"LKPP returned status {r.status_code}")
            return []
    except Exception as e:
        log.warning(f"LKPP connection failed: {e}")
        return []

    soup = BeautifulSoup(r.text, "html.parser")
    docs = []
    
    for a in soup.find_all("a", href=True):
        href = a["href"]
        if "/regulation/" in href and not any(x in href for x in ["/index", "/year", "/download", "/terjemahan"]):
            slug = href.split("/")[-1]
            title = a.text.strip()
            if not title or len(slug) < 5:
                continue
            
            # Form clean law number
            law_num = slug.replace("-", " ").title()
            num_match = re.search(r'nomor\s+(\d+)\s+tahun\s+(\d+)', slug.lower())
            if num_match:
                prefix = "Peraturan LKPP"
                if "keputusan" in slug.lower():
                    prefix = "Keputusan Kepala LKPP"
                law_num = f"{prefix} No. {num_match.group(1)} Tahun {num_match.group(2)}"
            
            docs.append({
                "law_number": law_num,
                "title": title,
                "source_url": f"https://jdih.lkpp.go.id/regulation/download-regulation?id={slug}",
                "source": "JDIH LKPP",
                "level": "sectoral",
                "document_type": "Peraturan LKPP",
                "published_date": ""
            })
    return docs


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
    if "BPK" in source:
        return download_bpk(url)

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


def download_bpk(url):
    """Download from BPK with Cloudflare bypass."""
    r = requests.get(url, impersonate="chrome120", timeout=45)
    if r.status_code != 200:
        raise Exception(f"BPK download failed: status {r.status_code}")

    # Check for Cloudflare challenge page
    content_type = r.headers.get("Content-Type", "")
    if "text/html" in content_type:
        text = r.text.lower()
        if "just a moment..." in text or "attention required!" in text:
            raise Exception("BPK blocked by Cloudflare challenge")

    # Return raw PDF bytes as base64
    import base64
    content_b64 = base64.b64encode(r.content).decode("utf-8")

    filename = url.split("/")[-1]
    # URL-decode the filename
    from urllib.parse import unquote
    filename = unquote(filename)

    return {
        "content": content_b64,
        "mime_type": r.headers.get("Content-Type", "application/pdf"),
        "filename": filename
    }


def search_bpk(url, law_number):
    """Search BPK for a specific law, parse results page."""
    r = requests.get(url, impersonate="chrome120", timeout=30)
    if r.status_code != 200:
        raise Exception(f"BPK search failed: status {r.status_code}")

    content_type = r.headers.get("Content-Type", "")
    if "text/html" in content_type:
        text = r.text.lower()
        if "just a moment..." in text or "attention required!" in text:
            raise Exception("BPK blocked by Cloudflare challenge")

    # Return raw HTML for Go to parse
    return {
        "content": r.text,
        "mime_type": "text/html",
        "filename": ""
    }


def extract_metadata(url, source):
    """Extract metadata from a detail page."""
    if "BPK" in source:
        return extract_metadata_bpk(url)
    return {"law_number": "", "title": "", "published_date": ""}


def extract_metadata_bpk(url):
    """Extract title and PDF link from BPK detail page."""
    r = requests.get(url, impersonate="chrome120", timeout=30)
    if r.status_code != 200:
        raise Exception(f"BPK detail fetch failed: status {r.status_code}")

    soup = BeautifulSoup(r.text, "html.parser")

    # Extract title from <title> tag
    title = ""
    title_tag = soup.find("title")
    if title_tag:
        title = title_tag.get_text(strip=True)
        # Remove " - JDIH BPK" suffix
        if " - " in title:
            title = title.split(" - ")[0]

    # Extract PDF download link
    pdf_url = ""
    for a in soup.find_all("a", href=True):
        href = a["href"]
        if "/Download/" in href and href.endswith(".pdf"):
            pdf_url = href
            if not href.startswith("http"):
                pdf_url = f"https://peraturan.bpk.go.id{href}"
            break

    return {
        "title": title,
        "pdf_url": pdf_url,
        "published_date": ""
    }


if __name__ == "__main__":
    main()
