#!/usr/bin/env python3
"""Run v3 chunker on all 16 sample PDFs and save results."""
import os
import sys
import json
from datetime import datetime

sys.path.insert(0, '/home/z/my-project/hukum-aneh/backend/python/chunker')
from parsers import chunk_pdf

CONFIG_PATH = "/home/z/my-project/config/directory_family_map.json"
SAMPLES_DIR = "/home/z/my-project/download/samples"

SAMPLES = {
    "uu": "uu/uunomor41tahun2014.pdf",
    "pp": "pp/PP_NO_70_TH_1991.pdf",
    "perppu": "perppu/perppu-no-148-tahun-2024.pdf",
    "perpres": "perpres/perpres-no-127-tahun-2024.pdf",
    "perda": "perda/perda-kabupaten-sukoharjo-no-1-tahun-2025.pdf",
    "keppres": "keppres/keppres-no-5-tahun-2015_Dewan Kawasan Kawasan Ekonomi Khusus Provinsi Kalimantan Timur.pdf",
    "inpres": "inpres/inpres-no-3-tahun-2023_Percepatan Peningkatan Konektivitas Jalan Daerah.pdf",
    "tap_mpr": "tap_mpr/Ketetapan Majelis Permusyawaratan Rakyat Republik Indonesia Nomor IIIMPR2002 ten.pdf",
    "uud-1945": "uud-1945/uud_1945.pdf",
    "Putusan-MK": "Putusan-MK/putusan_mkri_5301.pdf",
    "JDIH_Kemnaker": "JDIH_Kemnaker/Permenaker No. 90 Tahun 2013.pdf",
    "JDIH_Kemenkeu": "JDIH_Kemenkeu/PMK_No__9_Tahun_2025_2024pmkeuangan009.pdf",
    "JDIH_Kemendag": "JDIH_Kemendag/Kepmendag_No__123_Tahun_2025_download_3142_2.pdf",
    "JDIH_Komdigi": "JDIH_Komdigi/Permenkominfo No. 5 Tahun 2024.pdf",
    "JDIH_KPU": "JDIH_KPU/PKPU_8_2026.pdf",
    "peraturan": "peraturan/PP0201962.pdf",
}

with open(CONFIG_PATH) as f:
    family_map = json.load(f)

all_results = {}
all_ids = {}  # track ALL IDs for duplicate detection
qa_report = {
    "version": "v3-fixed",
    "timestamp": datetime.now().isoformat(),
    "files": {},
    "cross_cutting": {
        "duplicate_ids": [],
        "doc_type_mismatches": [],
        "missing_nomor_year": [],
        "stamp_noise_remaining": [],
    }
}

total_chunks = 0
total_errors = 0

for folder, rel_path in SAMPLES.items():
    full_path = os.path.join(SAMPLES_DIR, rel_path)
    if not os.path.exists(full_path):
        print(f"SKIP {folder}: file not found at {full_path}")
        continue

    print(f"\n{'='*60}")
    print(f"{folder} (Family {family_map.get(folder, {}).get('family', '?')})")
    print(f"{'='*60}")

    try:
        chunks = chunk_pdf(full_path, folder, family_map)
    except Exception as e:
        print(f"  ERROR: {e}")
        import traceback
        traceback.print_exc()
        all_results[folder] = [{"id": f"ERROR:{folder}", "text": str(e), "metadata": {"error": True}}]
        total_errors += 1
        continue

    all_results[folder] = chunks

    # QA checks
    file_qa = {"folder": folder, "filename": rel_path, "chunks": len(chunks), "issues": []}
    ids_seen = {}
    stamp_noise = 0

    for c in chunks:
        cid = c["id"]
        # Track duplicates
        if cid in ids_seen:
            ids_seen[cid] += 1
        else:
            ids_seen[cid] = 1
        # Track globally
        if cid in all_ids:
            all_ids[cid] += 1
        else:
            all_ids[cid] = 1
        # Check for stamp noise in text
        if "SK No" in c.get("text", "")[:50]:
            stamp_noise += 1

    dupes = {k: v for k, v in ids_seen.items() if v > 1}
    if dupes:
        file_qa["issues"].append({"type": "DUPLICATE_IDS", "count": len(dupes), "examples": list(dupes.keys())[:5]})
        qa_report["cross_cutting"]["duplicate_ids"].append({"folder": folder, "count": len(dupes), "examples": list(dupes.keys())[:5]})

    if stamp_noise > 0:
        file_qa["issues"].append({"type": "STAMP_NOISE", "count": stamp_noise})
        qa_report["cross_cutting"]["stamp_noise_remaining"].append({"folder": folder, "count": stamp_noise})

    # Check metadata
    first_meta = chunks[0]["metadata"] if chunks else {}
    doc_type = first_meta.get("doc_type", "?")
    nomor = first_meta.get("nomor")
    year = first_meta.get("year")
    
    if not nomor and doc_type != "UUD1945":
        file_qa["issues"].append({"type": "MISSING_NOMOR"})
        qa_report["cross_cutting"]["missing_nomor_year"].append(folder)
    if not year:
        file_qa["issues"].append({"type": "MISSING_YEAR"})

    # Check for error chunks
    error_chunks = [c for c in chunks if c["metadata"].get("error")]
    if error_chunks:
        file_qa["issues"].append({"type": "ERROR_CHUNKS", "count": len(error_chunks)})

    # Check for penjelasan status
    penjelasan_chunks = [c for c in chunks if c["metadata"].get("status") == "penjelasan"]
    if penjelasan_chunks:
        file_qa["penjelasan_chunks"] = len(penjelasan_chunks)

    qa_report["files"][folder] = file_qa

    total_chunks += len(chunks)

    # Print summary
    print(f"  Chunks: {len(chunks)}")
    print(f"  doc_type={doc_type}, nomor={nomor}, year={year}")
    if file_qa["issues"]:
        for iss in file_qa["issues"]:
            print(f"  [QA] {iss['type']}: {iss.get('count', '')} {iss.get('examples', '')}")
    else:
        print(f"  [QA] Clean")
    
    # Show first 3 chunk IDs
    for c in chunks[:3]:
        txt_preview = c['text'][:80] + "..." if len(c['text']) > 80 else c['text']
        print(f"  [{c['id']}] {txt_preview}")
    if len(chunks) > 3:
        print(f"  ... and {len(chunks)-3} more")

# Save results
output_path = "/home/z/my-project/download/chunk_results_v3.json"
with open(output_path, "w", encoding="utf-8") as f:
    json.dump(all_results, f, ensure_ascii=False, indent=2)
print(f"\nAll chunk results saved to: {output_path}")

# Save QA report
qa_path = "/home/z/my-project/download/chunk_qa_report_v3.json"
with open(qa_path, "w", encoding="utf-8") as f:
    json.dump(qa_report, f, ensure_ascii=False, indent=2)
print(f"QA report saved to: {qa_path}")

# Summary
print(f"\n{'='*60}")
print(f"SUMMARY: {len(all_results)} files, {total_chunks} total chunks, {total_errors} errors")
dupes_global = {k: v for k, v in all_ids.items() if v > 1}
if dupes_global:
    print(f"GLOBAL DUPLICATE IDs: {len(dupes_global)}")
    for k, v in list(dupes_global.items())[:10]:
        print(f"  {k}: appears {v}x")
else:
    print(f"GLOBAL DUPLICATE IDs: 0 (all unique!)")
