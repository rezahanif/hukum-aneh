#!/usr/bin/env python3"""Deep analysis of unclassified files in the Drive folder."""

import json
import re
from collections import Counter
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request
from googleapiclient.discovery import build

TOKEN_PATH = "/home/z/my-project/upload/token.json"
SCOPES = ["https://www.googleapis.com/auth/drive"]
ROOT_FOLDER = "1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"

# Standard slug pattern
SLUG_STD = re.compile(r'^(uu|uud|perppu|pp|perpres|keppres|inpres)-no-(\d+)-tahun-(\d+)', re.IGNORECASE)

# Indonesian-language patterns for unknown files
PATTERNS = [
    ("PP",       re.compile(r'PERATURAN\s+PEMERINTAH\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("PP",       re.compile(r'^PP\s+(?:No\.?\s*)?(\d+)\s+(?:Tahun\s+)?(\d{4})', re.I)),
    ("Perpres",  re.compile(r'PERATURAN\s+PRESIDEN\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("Perpres",  re.compile(r'^PERPRES\s+(?:No\.?\s*)?(\d+)\s+(?:Tahun\s+)?(\d{4})', re.I)),
    ("Keppres",  re.compile(r'KEPUTUSAN\s+PRESIDEN\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("Keppres",  re.compile(r'^KEPPRES\s+(?:No\.?\s*)?(\d+)\s+(?:Tahun\s+)?(\d{4})', re.I)),
    ("Inpres",   re.compile(r'INSTRUKSI\s+PRESIDEN\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("Inpres",   re.compile(r'^INPRES\s+(?:No\.?\s*)?(\d+)\s+(?:Tahun\s+)?(\d{4})', re.I)),
    ("Perppu",   re.compile(r'PERATURAN\s+PEMERINTAH\s+PENGGANTI\s+UNDANG.?UNDANG\s+NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("TAP MPR",  re.compile(r'KETETAPAN\s+MAJELIS\s+PERMUSYAWARATAN\s+RAKYAT\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+([IVXLCDM]+).*?TAHUN\s+(\d{4})', re.I)),
    ("TAP MPR",  re.compile(r'TAP\s+MPR\s+(?:No\.?\s*)?([IVXLCDM]+).*?(\d{4})', re.I)),
    ("UU",       re.compile(r'UNDANG.?UNDANG\s+(?:REPUBLIK\s+INDONESIA\s+)?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("UUD",      re.compile(r'UNDANG.?UNDANG\s+DASAR\s+(?:NEGARA\s+)?REPUBLIK\s+INDONESIA\s*(\d{4})?', re.I)),
    ("Kepmen",   re.compile(r'^KEPMEN\w*\s+(?:No\.?\s*)?(\d+)\s+(?:Tahun\s+)?(\d{4})', re.I)),
    ("PMK",      re.compile(r'^PMK\s+(?:No\.?\s*)?(\d+)[/.]?(?:PMK)?\s*[/]?\s*(?:Tahun\s+)?(\d{4})', re.I)),
    ("Perbup",   re.compile(r'PERATURAN\s+BUPATI\s+.*?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("Perwal",   re.compile(r'PERATURAN\s+WALIKOTA\s+.*?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("Perwali",  re.compile(r'PERATURAN\s+WALI\s+KOTA\s+.*?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("Perda",    re.compile(r'PERATURAN\s+DAERAH\s+.*?NOMOR\s+(\d+)\s+TAHUN\s+(\d+)', re.I)),
    ("Putusan",  re.compile(r'PUTUSAN\s+(?:MK|MA|pengadilan)', re.I)),
    ("SE",       re.compile(r'^SE\s+.*?(\d{4})', re.I)),
    ("PBI",      re.compile(r'^PBI\s+(?:No\.?\s*)?(\d+)\s+(?:Tahun\s+)?(\d{4})', re.I)),
]

# JDIH department patterns
DEPT_PATTERNS = [
    ("Kemnaker",  re.compile(r'kemnaker|kementerian\s+ketenagakerjaan', re.I)),
    ("Kemenkeu",  re.compile(r'kemenkeu|kementerian\s+keuangan', re.I)),
    ("Kemendag",  re.compile(r'kemendag|kepmendag|kementerian\s+perdagangan', re.I)),
    ("Komdigi",   re.compile(r'komdigi|kementerian\s+komunikasi\s+dan\s+digital', re.I)),
    ("KPU",       re.compile(r'kpu|komisi\s+pemilihan\s+umum', re.I)),
    ("BPK",       re.compile(r'bpk|badan\s+pemeriksa\s+keuangan', re.I)),
    ("BKPM",      re.compile(r'bkpm|badan\s+koordinasi\s+penanaman\s+modal', re.I)),
    ("OJK",       re.compile(r'ojk|otoritas\s+jasa\s+keuangan', re.I)),
    ("BI",        re.compile(r'bank\s+indonesia|pbi\b', re.I)),
    ("LKPP",      re.compile(r'lkpp|lembaga\s+kebijakan\s+pengadaan', re.I)),
    ("BNPT",      re.compile(r'bnpt', re.I)),
    ("BSSN",      re.compile(r'bssn', re.I)),
    ("Kemendikbud", re.compile(r'kemendikbud|kementerian\s+pendidikan', re.I)),
    ("Kemenkes",  re.compile(r'kemenkes|kementerian\s+kesehatan', re.I)),
    ("Kemenag",   re.compile(r'kemenag|kementerian\s+agama', re.I)),
    ("Kemenkumham", re.compile(r'kemenkumham|kementerian\s+hukum\s+dan', re.I)),
    ("BNP2TKI",   re.compile(r'bnp2tki', re.I)),
    ("BKN",       re.compile(r'\bbkn\b', re.I)),
    ("DPR",       re.compile(r'dewan\s+perwakilan\s+rakyat', re.I)),
    ("Setneg",    re.compile(r'setneg|sekretariat\s+negara', re.I)),
]


def get_service():
    creds = Credentials.from_authorized_user_file(TOKEN_PATH, SCOPES)
    if not creds.valid and creds.expired and creds.refresh_token:
        creds.refresh(Request())
    return build("drive", "v3", credentials=creds)


def collect_all_files(service, folder_id):
    """Get all files recursively."""
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


def classify_unknown(name):
    """Try to classify unknown filenames."""
    name_clean = name.replace(".pdf", "")
    
    # Try each pattern
    for label, pattern in PATTERNS:
        m = pattern.search(name_clean)
        if m:
            groups = m.groups()
            return {"type": label, "num": groups[0] if len(groups) > 0 else None, "year": groups[1] if len(groups) > 1 else None}
    
    return {"type": "OTHER", "num": None, "year": None}


def detect_department(name):
    """Detect which JDIH department a file belongs to."""
    for dept, pattern in DEPT_PATTERNS:
        if pattern.search(name):
            return dept
    return None


def main():
    service = get_service()
    print("Collecting all files (this may take a while)...")
    all_files = collect_all_files(service, ROOT_FOLDER)
    print(f"Total files: {len(all_files)}")

    # Separate classified vs unknown
    classified = []
    unknown = []
    for f in all_files:
        name = f["name"].lower().replace(".pdf", "")
        if SLUG_STD.match(name):
            classified.append(f)
        else:
            unknown.append(f)

    print(f"Standard slug files: {len(classified)}")
    print(f"Non-standard files:  {len(unknown)}")

    # Classify unknown files
    print(f"\n{'='*70}")
    print("  CLASSIFYING UNKNOWN FILES")
    print(f"{'='*70}")

    type_counts = Counter()
    dept_counts = Counter()
    type_dept = {}  # {(type, dept): count}
    year_by_type = {}  # {type: Counter({year: count})}
    reclassified = []
    still_unknown = []

    for f in unknown:
        info = classify_unknown(f["name"])
        dept = detect_department(f["name"])

        if dept:
            dept_counts[dept] += 1

        if info["type"] != "OTHER":
            type_counts[info["type"]] += 1
            key = (info["type"], dept or "unspecified")
            type_dept[key] = type_dept.get(key, 0) + 1

            if info["year"] and info["year"].isdigit():
                yr = int(info["year"])
                if info["type"] not in year_by_type:
                    year_by_type[info["type"]] = Counter()
                year_by_type[info["type"]][yr] += 1

            reclassified.append({"file": f["name"], **info, "dept": dept})
        else:
            still_unknown.append(f)

    print(f"\n  Reclassified: {len(reclassified)}")
    print(f"  Still unknown: {len(still_unknown)}")

    print(f"\n  ── By law type (reclassified) ──")
    for t, c in sorted(type_counts.items(), key=lambda x: -x[1]):
        print(f"    {t}: {c}")

    print(f"\n  ── By department ──")
    for d, c in sorted(dept_counts.items(), key=lambda x: -x[1]):
        print(f"    {d}: {c}")

    print(f"\n  ── By (type, department) ──")
    for (t, d), c in sorted(type_dept.items(), key=lambda x: -x[1])[:30]:
        print(f"    {t:12} | {d:15}: {c}")

    if year_by_type:
        print(f"\n  ── Reclassified by year ──")
        for dtype in sorted(year_by_type.keys()):
            yc = year_by_type[dtype]
            years_str = ", ".join(f"{y}:{c}" for y, c in sorted(yc.items())[-10:])
            print(f"    {dtype}: {sum(yc.values())} total | recent: {years_str}")

    # Show sample of still-unknown
    if still_unknown:
        print(f"\n  ── STILL UNKNOWN (sample, first 40) ──")
        for f in sorted(still_unknown, key=lambda x: x["name"])[:40]:
            size_kb = int(f.get("size", 0)) // 1024
            print(f"    {size_kb:>6}KB  {f['name'][:80]}")
        if len(still_unknown) > 40:
            print(f"    ... {len(still_unknown) - 40} more")

    # Save full report
    report = {
        "total_files": len(all_files),
        "classified_standard": len(classified),
        "reclassified": len(reclassified),
        "still_unknown": len(still_unknown),
        "type_breakdown": dict(type_counts),
        "dept_breakdown": dict(dept_counts),
        "type_dept_breakdown": {f"{t}/{d}": c for (t, d), c in type_dept.items()},
        "unknown_samples": [f["name"] for f in sorted(still_unknown, key=lambda x: x["name"])[:100]],
    }
    with open("/home/z/my-project/download/drive_full_analysis.json", "w") as fp:
        json.dump(report, fp, indent=2, ensure_ascii=False)
    print(f"\n  Full report saved to download/drive_full_analysis.json")


if __name__ == "__main__":
    main()