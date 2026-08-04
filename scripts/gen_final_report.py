#!/usr/bin/env python3
"""Generate final v3 report and prepare for upload."""
import json
from collections import Counter
from datetime import datetime

with open('/home/z/my-project/download/chunk_results_v3.json') as f:
    data = json.load(f)

report = {
    "version": "v3-fixed",
    "generated_at": datetime.now().isoformat(),
    "description": "Chunking pipeline v3 — post-QA-fix test results across 16 Indonesian legal document samples",
    "fixes_applied_in_v3": [
        "clean_extractor: UUD 1945 nomor=None (was '1945'), year=1945 only",
        "clean_extractor: Removed crashed title_block_text reference (tap_mpr NameError)",
        "clean_extractor: Added broader stamp pattern (RE_SETNEG_STAMP2)",
        "clean_extractor: Restricted extract_nomor_year to first 500 chars (JDIH_KPU body citation issue)",
        "clean_extractor: Added consistent O-as-0 glyph normalization (3 patterns)",
        "parsers: doc_type resolution — title block first, directory config as fallback, broad-type guard for citation pollution",
        "parsers: Added Penjelasan section handling with distinct :penjelasan ID suffix",
        "parsers: Added PENJELASAN detection without requiring PENUTUP first",
        "parsers: Added signature block as hard chunk boundary in penutup section",
        "parsers: Broadened FamilyB signature detection (ditetapkan di, dilantik di, PENJELASAN)",
        "parsers: Added parent_section tracking for inpres diktum numbering",
        "parsers: Relaxed RE_PASAL regex (removed .{0,5}$ tail restriction)",
    ],
    "v2_to_v3_improvements": [
        "pp: Fixed crash (NameError in penjelasan code path) — now produces 63 chunks",
        "keppres: doc_type correctly detected as Keppres (was PP from title citation)",
        "JDIH_Kemnaker: doc_type correctly detected as Permen (was PP from title citation)",
        "JDIH_Komdigi: doc_type correctly detected as Permen (was PP from title citation)",
        "JDIH_Kemnaker: Penjelasan separated into distinct :penjelasan chunk (was merged with operative)",
        "UU: Duplicate IDs reduced from 34 to 21 (Penjelasan separation)",
        "perda: Duplicate IDs reduced from 18 to 9 (Penjelasan separation)",
        "UUD 1945: nomor is now None instead of '1945' (semantically correct)",
        "tap_mpr: No longer crashes — returns clean error chunk for scanned PDF",
    ],
    "known_remaining_issues": [
        "Roman numeral pasal (I, II, III) and Arabic (1, 2, 3) can produce same ID when they refer to different pasals",
        "Penjelasan detection: works for docs with explicit PENJELASAN header; implicit penjelasan still merges with operative",
        "tap_mpr: scanned PDF, no extractable text (needs OCR)",
        "JDIH_KPU: page 1 is image, title detection gets nomor=07/2012 instead of 8/2026",
        "Stamp noise: perpres (1 chunk) and inpres (2 chunks) still have SK No variants",
        "Some duplicate IDs remain from definition pasals (Pasal 1 definitions) being flushed twice",
        "JDIH_Komdigi: 24 duplicate IDs (word-level PDF extraction causes fragmented lines)",
    ],
    "per_file_results": {},
}

total_chunks = 0
total_dupes = 0
total_errors = 0
clean_count = 0

for folder, chunks in sorted(data.items()):
    errors = sum(1 for c in chunks if c["metadata"].get("error"))
    penjelasan = sum(1 for c in chunks if c["metadata"].get("status") == "penjelasan")
    quoted = sum(1 for c in chunks if c["metadata"].get("status") == "quoted_amendment")
    stamp = sum(1 for c in chunks if "SK No" in c.get("text", "")[:50])
    
    ids = [c["id"] for c in chunks]
    id_counts = Counter(ids)
    dupes = {k: v for k, v in id_counts.items() if v > 1}
    
    first_meta = chunks[0]["metadata"] if chunks else {}
    is_clean = len(dupes) == 0 and errors == 0 and stamp == 0
    if is_clean:
        clean_count += 1
    
    total_chunks += len(chunks)
    total_dupes += len(dupes)
    total_errors += errors
    
    result = {
        "status": "CLEAN" if is_clean else "HAS_ISSUES",
        "chunks": len(chunks),
        "errors": errors,
        "duplicate_ids": len(dupes),
        "penjelasan_chunks": penjelasan,
        "quoted_amendment_chunks": quoted,
        "stamp_noise": stamp,
        "doc_type": first_meta.get("doc_type", "?"),
        "nomor": first_meta.get("nomor"),
        "year": first_meta.get("year"),
    }
    if dupes:
        result["duplicate_examples"] = list(dupes.keys())[:5]
    report["per_file_results"][folder] = result

report["summary"] = {
    "total_files": len(data),
    "total_chunks": total_chunks,
    "clean_files": f"{clean_count}/{len(data)}",
    "files_with_issues": len(data) - clean_count,
    "total_duplicate_ids": total_dupes,
    "total_errors": total_errors,
    "penjelasan_chunks_total": sum(r["penjelasan_chunks"] for r in report["per_file_results"].values()),
    "quoted_amendment_chunks_total": sum(r["quoted_amendment_chunks"] for r in report["per_file_results"].values()),
}

# Save
out_path = '/home/z/my-project/download/chunk_test_report_v3.json'
with open(out_path, 'w', encoding='utf-8') as f:
    json.dump(report, f, ensure_ascii=False, indent=2)
print(f"Report saved to {out_path}")
print(f"\nSummary: {report['summary']}")
