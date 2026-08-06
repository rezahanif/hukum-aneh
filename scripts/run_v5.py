#!/usr/bin/env python3
"""v5 chunking pipeline — run all 16 sample PDFs and produce chunk results + QA report."""
import json
import os
import sys
from collections import Counter, defaultdict

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
total_chunks = 0

for folder, rel_path in SAMPLES.items():
    full_path = os.path.join(SAMPLES_DIR, rel_path)
    if not os.path.exists(full_path):
        print(f"SKIP {folder}: file not found")
        continue

    print(f"\n{'='*60}")
    print(f"{folder} (Family {family_map.get(folder, {}).get('family', '?')})")
    print(f"{'='*60}")

    chunks = chunk_pdf(full_path, folder, family_map)
    all_results[folder] = chunks
    total_chunks += len(chunks)

    print(f"Chunks produced: {len(chunks)}")
    for c in chunks[:5]:
        txt_preview = c['text'][:80] + "..." if len(c['text']) > 80 else c['text']
        print(f"  [{c['id']}] {txt_preview}")
    if len(chunks) > 5:
        print(f"  ... and {len(chunks)-5} more")

# Save full chunk results
output_path = "/home/z/my-project/download/chunk_results_v5.json"
with open(output_path, "w", encoding="utf-8") as f:
    json.dump(all_results, f, ensure_ascii=False, indent=2)
print(f"\nAll results saved to: {output_path}")
print(f"Total chunks: {total_chunks}")

# === QA Report ===
print(f"\n{'='*60}")
print("QA REPORT")
print(f"{'='*60}")

qa = {
    "version": "v5",
    "total_files": len(all_results),
    "total_chunks": total_chunks,
    "per_file": {},
    "issues": []
}

dup_stats = defaultdict(lambda: {"count": 0, "ids": set()})
global_dup_chunks = 0

for folder, chunks in all_results.items():
    file_qa = {"chunks": len(chunks), "doc_type": None, "nomor": None, "year": None, "pasal_ids": [], "dup_ids": 0, "issues": []}
    
    if chunks and not chunks[0].get("metadata", {}).get("error"):
        m = chunks[0]["metadata"]
        file_qa["doc_type"] = m.get("doc_type")
        file_qa["nomor"] = m.get("nomor")
        file_qa["year"] = m.get("year")
    
    # Collect pasal IDs and check duplicates
    id_counts = Counter(c["id"] for c in chunks)
    dups = {id_: cnt for id_, cnt in id_counts.items() if cnt > 1}
    file_qa["dup_ids"] = len(dups)
    file_qa["dup_detail"] = {id_: cnt for id_, cnt in sorted(dups.items(), key=lambda x: -x[1])[:10]}
    
    dup_chunk_count = sum(cnt - 1 for cnt in dups.values())
    global_dup_chunks += dup_chunk_count
    
    # Collect pasal numbers
    pasal_nums = set()
    for c in chunks:
        p = c.get("metadata", {}).get("pasal")
        if p and p != "?":
            pasal_nums.add(p)
    file_qa["pasal_ids"] = sorted(pasal_nums, key=lambda x: (len(x), x))
    
    # Check for known issues
    # perpres: Pasal 3 missing?
    if folder == "perpres":
        if "3" not in pasal_nums:
            file_qa["issues"].append("Pasal 3 still missing")
        else:
            file_qa["issues"].append("Pasal 3 PRESENT (fixed)")
        # Check for glyph corruption
        for c in chunks:
            t = c["text"]
            if "Trnjangan" in t or "Tfrnjangan" in t or "T\\rnjangan" in t:
                file_qa["issues"].append(f"Glyph corruption still present: {t[:60]}")
                break
            if "sglagaimana" in t or "sslagaimana" in t:
                file_qa["issues"].append(f"sebagaimana corruption: {t[:60]}")
                break
            if "SK No" in t and "Agar" in t:
                file_qa["issues"].append(f"Stamp noise still present: {t[:60]}")
                break
        else:
            file_qa["issues"].append("Glyph/stamp checks passed")
    
    # uu: check 68A-E
    if folder == "uu":
        alpha_pasals = [p for p in pasal_nums if any(c.isalpha() for c in p)]
        if alpha_pasals:
            file_qa["issues"].append(f"Alpha pasals found: {alpha_pasals}")
        else:
            file_qa["issues"].append("WARNING: No alpha-suffixed pasals found (68A-E may still collapse)")
    
    # perppu: check duplicate count
    if folder == "perppu":
        file_qa["issues"].append(f"Dup chunks: {dup_chunk_count}")
    
    qa["per_file"][folder] = file_qa
    print(f"\n{folder}: {len(chunks)} chunks, {len(dups)} dup IDs ({dup_chunk_count} dup chunks)")
    for issue in file_qa["issues"]:
        print(f"  - {issue}")

qa["total_dup_ids"] = sum(f["dup_ids"] for f in qa["per_file"].values())
qa["total_dup_chunks"] = global_dup_chunks

# Save QA report
qa_path = "/home/z/my-project/download/chunk_qa_report_v5.json"
with open(qa_path, "w", encoding="utf-8") as f:
    json.dump(qa, f, ensure_ascii=False, indent=2)
print(f"\nQA report saved to: {qa_path}")
print(f"\nSUMMARY: {qa['total_files']} files, {qa['total_chunks']} chunks, {qa['total_dup_ids']} dup IDs, {qa['total_dup_chunks']} dup chunks")
