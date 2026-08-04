#!/usr/bin/env python3
"""Deep analysis v2 — broader pattern matching for all file types."""

import json
import re
from collections import Counter, defaultdict
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
ROOT_FOLDER = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"

SLUG_STD = re.compile(r'^(uu|uud|perppu|pp|perpres|keppres|inpres)-no-(\d+)-tahun-(\d+)', re.IGNORECASE)

# Broader patterns — order matters (more specific first)
PATTERNS = [
    # TAP MPR
    ("TAP MPR",  re.compile(r'KETETAPAN\s+MAJELIS\s+PERMUSYAWARATAN\s+RAKYAT.*?NOMOR\s+([IVXLCDM]+)[^\d]*?TAHUN\s+(\d{4})', re.I)),
    ("TAP MPR",  re.compile(r'TAP\s*[_-]?MPR\s*[_-]?NO\s*[_-]?([IVXLCDM]+)[^\d]*(\d{4})', re.I)),
    # UU
    ("UU",       re.compile(r'UNDANG[.\s-]*UNDANG\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    # Perppu
    ("Perppu",   re.compile(r'PERATURAN\s+PEMERINTAH\s+PENGGANTI\s+UNDANG[.\s-]*UNDANG\s+NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    ("Perppu",   re.compile(r'PERPPU\s*(?:NO)?\s*(?:_)?\s*(\d+)\s*(?:TAHUN)?\s*(\d{4})', re.I)),
    # PP
    ("PP",       re.compile(r'PERATURAN\s+PEMERINTAH\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    ("PP",       re.compile(r'^PP[_\s]+(?:NO[_\s]*)?(\d+)[_\s]+TAHUN[_\s]+(\d{4})', re.I)),
    ("PP",       re.compile(r'^PP[_\s]+(?:NO[_\s]*)?(\d+)\s*(?:TH)?\s*(?:/)?\s*(\d{4})', re.I)),
    # Perpres
    ("Perpres",  re.compile(r'PERATURAN\s+PRESIDEN\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    ("Perpres",  re.compile(r'^PERPRES[_\s]+(?:NO[_\s]*)?(\d+)[_\s]+TAHUN[_\s]+(\d{4})', re.I)),
    ("Perpres",  re.compile(r'^PERPRES\s*(?:NO)?\s*(\d+)\s*(?:TAHUN|TH)?\s*(?:/)?\s*(\d{4})', re.I)),
    # Keppres
    ("Keppres",  re.compile(r'KEPUTUSAN\s+PRESIDEN\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    ("Keppres",  re.compile(r'^KEPPRES[_\s]+(?:NO[_\s]*)?(\d+)[_\s]+TAHUN[_\s]+(\d{4})', re.I)),
    # Inpres
    ("Inpres",   re.compile(r'INSTRUKSI\s+PRESIDEN\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    ("Inpres",   re.compile(r'^INPRES[_\s]+(?:NO[_\s]*)?(\d+)[_\s]+TAHUN[_\s]+(\d{4})', re.I)),
    # Kepmen (department ministerial decrees) — many variants
    ("Kepmen",   re.compile(r'^(Kep\w+)\s*[_]*No\s*_*\s*(\d+)\s*_*Tahun\s*_*(\d{4})', re.I)),
    ("Kepmen",   re.compile(r'^Keputusan\s+(\w+)\s+NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    # PMK (Peraturan Menteri Keuangan)
    ("PMK",      re.compile(r'^PMK[_\s]*(?:NO)?\s*(?:_)?\s*(\d+)[/\s]*PMK[/\s]*(?:\d{4}[/-])?(\d{4})', re.I)),
    ("PMK",      re.compile(r'^PMK[_\s]+(?:NO[_\s]*)?(\d+)[_\s]+(?:TAHUN[_\s]+)?(\d{4})', re.I)),
    # PBI (Peraturan Bank Indonesia)
    ("PBI",      re.compile(r'^PBI[_\s]+(?:NO[_\s]*)?(\d+)[/_\s]+(?:TAHUN[_\s]+)?(\d{4})', re.I)),
    # Peraturan BPK
    ("PerBPK",   re.compile(r'^Peraturan\s+BPK\s+NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    # Peraturan Daerah
    ("Perda",    re.compile(r'PERATURAN\s+(?:DAERAH|PROVINSI|KABUPATEN|KOTA|WALI)\s+.*?NOMOR\s+(\d+)\s+TAHUN\s+(\d{4})', re.I)),
    ("Perda",    re.compile(r'^PERDA[_\s]+(?:NO)?\s*(\d+)\s*(?:TAHUN|TH)?\s*(?:/)?\s*(\d{4})', re.I)),
    # Putusan MK
    ("Putusan MK", re.compile(r'PUTUSAN\s+(?:MK|MAHKAMAH\s+KONSTITUSI)', re.I)),
    ("Putusan MK", re.compile(r'^Putusan[_\s]+MK', re.I)),
    # SE (Surat Edaran)
    ("SE",       re.compile(r'^SE[_\s]+.*?(\d{4})', re.I)),
    ("SE",       re.compile(r'^SURAT\s+EDARAN', re.I)),
    # Peraturan LKPP
    ("PerLKPP",  re.compile(r'^Peraturan[_\s]+LKPP', re.I)),
]

# Department patterns
DEPT_PATTERNS = [
    ("Kemnaker",  re.compile(r'kemnaker|kementerian\s+ketenagakerjaan', re.I)),
    ("Kemenkeu",  re.compile(r'kemenkeu|kementerian\s+keuangan', re.I)),
    ("Kemendag",  re.compile(r'kemendag|kepmendag', re.I)),
    ("Komdigi",   re.compile(r'komdigi|kementerian\s+komunikasi\s+dan\s+digital', re.I)),
    ("KPU",       re.compile(r'\bkpu\b|komisi\s+pemilihan\s+umum', re.I)),
    ("BPK",       re.compile(r'\bbpk\b|badan\s+pemeriksa\s+keuangan', re.I)),
    ("OJK",       re.compile(r'\bojk\b', re.I)),
    ("BI",        re.compile(r'bank\s+indonesia|\bpbi\b', re.I)),
    ("LKPP",      re.compile(r'\blkpp\b', re.I)),
    ("BNPT",      re.compile(r'\bbnpt\b', re.I)),
    ("BSSN",      re.compile(r'\bbssn\b', re.I)),
    ("BKN",       re.compile(r'\bbkn\b', re.I)),
    ("Kemendikbud", re.compile(r'kemendikbud|kemdikbud', re.I)),
    ("Kemenkes",  re.compile(r'kemenkes', re.I)),
    ("Kemenag",   re.compile(r'kemenag', re.I)),
    ("Kemenkumham", re.compile(r'kemenkumham|kementerian\s+hukum', re.I)),
    ("DPR",       re.compile(r'dewan\s+perwakilan\s+rakyat', re.I)),
    ("Setneg",    re.compile(r'setneg|sekretariat\s+negara', re.I)),
]


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def collect_all_files(service, folder_id):
    all_files = []
    folders_to_scan = [folder_id]
    while folders_to_scan:
        fid = folders_to_scan.pop(0)
        page_token = None
        while True:
            results = service.files().list(
                q=f"'{fid}' in parents and trashed=false",
                fields="files(id, name, size, md5Checksum, mimeType, parents), nextPageToken",
                pageSize=1000,
                pageToken=page_token
            ).execute()
            for f in results.get("files", []):
                if f["mimeType"] == "application/vnd.google-apps.folder":
                    folders_to_scan.append(f["id"])
                else:
                    all_files.append(f)
            page_token = results.get("nextPageToken")
            if not page_token:
                break
    return all_files


def classify(name):
    for label, pattern in PATTERNS:
        m = pattern.search(name)
        if m:
            groups = m.groups()
            return label, groups
    return "OTHER", None


def detect_dept(name):
    for dept, pattern in DEPT_PATTERNS:
        if pattern.search(name):
            return dept
    return None


def main():
    service = get_service()
    print("Collecting all files...")
    all_files = collect_all_files(service, ROOT_FOLDER)
    print(f"Total files: {len(all_files)}")

    # Count unique and duplicate (by md5)
    by_md5 = defaultdict(list)
    for f in all_files:
        md5 = f.get("md5Checksum") or "no_md5"
        by_md5[md5].append(f)

    dup_md5 = {m: files for m, files in by_md5.items() if len(files) > 1}
    dup_count = sum(len(v) - 1 for v in dup_md5.values())
    print(f"Unique MD5s: {len(by_md5)}, duplicate files (same content): {dup_count}")

    # Classify all
    results = defaultdict(lambda: {"count": 0, "by_year": Counter(), "by_dept": Counter(), "samples": []})
    std_count = 0
    unknown_count = 0
    unknown_samples = []
    dept_counts = Counter()

    for f in all_files:
        name = f["name"]
        name_lower = name.lower().replace(".pdf", "")
        dept = detect_dept(name)
        if dept:
            dept_counts[dept] += 1

        # Standard slug
        m = SLUG_STD.match(name_lower)
        if m:
            dtype = m.group(1).upper()
            year = int(m.group(3))
            key = dtype
            results[key]["count"] += 1
            results[key]["by_year"][year] += 1
            if dept:
                results[key]["by_dept"][dept] += 1
            std_count += 1
            continue

        # Try broader patterns
        label, groups = classify(name)
        if label != "OTHER":
            year = None
            if groups:
                for g in groups:
                    if g and g.isdigit() and len(g) == 4:
                        year = int(g)
                        break
            results[label]["count"] += 1
            if year:
                results[label]["by_year"][year] += 1
            if dept:
                results[label]["by_dept"][dept] += 1
            if len(results[label]["samples"]) < 5:
                results[label]["samples"].append(name[:80])
        else:
            unknown_count += 1
            if len(unknown_samples) < 50:
                unknown_samples.append({"name": name[:90], "size_kb": int(f.get("size", 0)) // 1024, "dept": dept})

    # Output
    total_classified = sum(r["count"] for r in results.values())
    print(f"\nStandard slug files: {std_count}")
    print(f"Reclassified: {total_classified - std_count}")
    print(f"Still unknown: {unknown_count}")

    print(f"\n{'='*70}")
    print(f"  FULL TYPE BREAKDOWN")
    print(f"{'='*70}")
    for key in sorted(results.keys(), key=lambda k: -results[k]["count"]):
        r = results[key]
        print(f"\n  {key}: {r['count']} files")
        if r["by_dept"]:
            dept_str = ", ".join(f"{d}:{c}" for d, c in sorted(r["by_dept"].items(), key=lambda x: -x[1]))
            print(f"    Depts: {dept_str}")
        if r["by_year"]:
            years_str = ", ".join(f"{y}:{c}" for y, c in sorted(r["by_year"].items())[-8:])
            total_yrs = len(r["by_year"])
            print(f"    Years ({total_yrs} unique): ...{years_str}")
        if r["samples"]:
            for s in r["samples"][:3]:
                print(f"    e.g. {s}")

    print(f"\n{'='*70}")
    print(f"  DEPARTMENT OVERVIEW")
    print(f"{'='*70}")
    for d, c in sorted(dept_counts.items(), key=lambda x: -x[1]):
        print(f"  {d}: {c}")

    print(f"\n{'='*70}")
    print(f"  DUPLICATES (same MD5, different files): {dup_count}")
    print(f"{'='*70}")
    for md5, files in sorted(dup_md5.items(), key=lambda x: -len(x[1]))[:15]:
        print(f"  MD5:{md5[:8]} — {len(files)} copies:")
        for f in files[:3]:
            print(f"    {f['name'][:80]}")

    if unknown_samples:
        print(f"\n{'='*70}")
        print(f"  STILL UNKNOWN: {unknown_count} (sample)")
        print(f"{'='*70}")
        for s in unknown_samples[:30]:
            dept_str = f" [{s['dept']}]" if s['dept'] else ""
            print(f"    {s['size_kb']:>6}KB{dept_str}  {s['name']}")


if __name__ == "__main__":
    main()