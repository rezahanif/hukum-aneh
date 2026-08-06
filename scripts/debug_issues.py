#!/usr/bin/env python3
"""Debug root causes for v4 QA issues."""
import sys, json, re
sys.path.insert(0, '/home/z/my-project/hukum-aneh/backend/python/chunker')
from clean_extractor import extract_clean, extract_title_block, detect_doc_type_from_title, fix_glyph_corruption

# ---- Issue 1: doc_type resolution ----
print("=" * 60)
print("ISSUE 1: DOC_TYPE RESOLUTION — title block content")
print("=" * 60)

for folder, path in [
    ('perppu', '/home/z/my-project/download/samples/perppu/perppu-no-148-tahun-2024.pdf'),
    ('JDIH_Kemnaker', '/home/z/my-project/download/samples/JDIH_Kemnaker/Permenaker No. 90 Tahun 2013.pdf'),
    ('keppres', '/home/z/my-project/download/samples/keppres/keppres-no-5-tahun-2015_Dewan Kawasan Kawasan Ekonomi Khusus Provinsi Kalimantan Timur.pdf'),
    ('JDIH_KPU', '/home/z/my-project/download/samples/JDIH_KPU/PKPU_8_2026.pdf'),
]:
    print(f"\n--- {folder} ---")
    lines, meta = extract_clean(path)
    
    # Show first 10 non-boilerplate lines (what title block sees)
    print(f"  Pages: {meta['total_pages']}, Active lines: {meta['total_lines_active']}")
    print(f"  Raw nomor/year: {meta['nomor']}/{meta['year']}")
    
    # Title block (first 3 non-boilerplate)
    title3 = extract_title_block(lines, max_first_n=3)
    print(f"  Title block (3 lines): {title3[:200]}")
    
    title5 = extract_title_block(lines, max_first_n=5)
    print(f"  Title block (5 lines): {title5[:300]}")
    
    dt3, iss3 = detect_doc_type_from_title(title3)
    dt5, iss5 = detect_doc_type_from_title(title5)
    print(f"  Detected (3 lines): {dt3} / {iss3}")
    print(f"  Detected (5 lines): {dt5} / {iss5}")

# ---- Issue 2: JDIH_KPU nomor/year ----
print("\n" + "=" * 60)
print("ISSUE 2: JDIH_KPU NOMOR/YEAR")
print("=" * 60)

lines_kpu, meta_kpu = extract_clean('/home/z/my-project/download/samples/JDIH_KPU/PKPU_8_2026.pdf')
print(f"  Page 1 lines: {[l['text'][:80] for l in lines_kpu if l['page'] == 1][:5]}")
print(f"  Page 2 lines: {[l['text'][:80] for l in lines_kpu if l['page'] == 2][:5]}")
print(f"  First 10 active: {[l['text'][:80] for l in lines_kpu[:10]]}")
raw = ' '.join(l['text'] for l in lines_kpu[:20])
print(f"  Raw first 500: {raw[:500]}")

# ---- Issue 3: pp penjelasan ----
print("\n" + "=" * 60)
print("ISSUE 3: PP PENJELASAN DETECTION")
print("=" * 60)

lines_pp, meta_pp = extract_clean('/home/z/my-project/download/samples/pp/PP_NO_70_TH_1991.pdf')
# Find PENJELASAN or penutup
for i, l in enumerate(lines_pp):
    t = l['text'].strip().upper()
    if 'PENJELASAN' in t or 'PENUTUP' in t or 'TETAPKAN' in t or 'DITETAPKAN' in t:
        print(f"  Line {i}: p{l['page']} | {l['text'][:100]}")

# ---- Issue 4: perppu pasal counter ----
print("\n" + "=" * 60)
print("ISSUE 4: PERPPU PASAL COUNTER")
print("=" * 60)

lines_perppu, meta_perppu = extract_clean('/home/z/my-project/download/samples/perppu/perppu-no-148-tahun-2024.pdf')
# Find all Pasal matches
pasal_pattern = re.compile(r'^\s*pasal\s+([IVXLCDM]+|\d+[a-zA-Z]?)\s*\.?(?:\s|$)', re.IGNORECASE)
for i, l in enumerate(lines_perppu):
    m = pasal_pattern.match(l['text'])
    if m:
        print(f"  Line {i}: p{l['page']} Pasal {m.group(1)} | {l['text'][:80]}")

# ---- Issue 5: perpres Pasal 3 ----
print("\n" + "=" * 60)
print("ISSUE 5: PERPRES PASAL 3 MISSING")
print("=" * 60)

lines_perpres, meta_perpres = extract_clean('/home/z/my-project/download/samples/perpres/perpres-no-127-tahun-2024.pdf')
for i, l in enumerate(lines_perpres):
    m = pasal_pattern.match(l['text'])
    if m:
        print(f"  Line {i}: p{l['page']} Pasal {m.group(1)} | {l['text'][:100]}")
    # Also check for stamp noise
    if 'SK No' in l['text'][:30]:
        print(f"  STAMP: Line {i}: p{l['page']} | {l['text'][:100]}")
    if 'Tunjangan' in l['text'] or 'PRESTDEN' in l['text'] or 'Tfrnjangan' in l['text'] or 'T\r' in l['text']:
        print(f"  GLYPH: Line {i}: p{l['page']} | {repr(l['text'][:100])}")

# ---- Issue 6: pp heading-as-chunk ----
print("\n" + "=" * 60)
print("ISSUE 6: PP HEADING-AS-CHUNK")
print("=" * 60)

# Show chunks with duplicate PP:70:1991:2
import json as j
with open('/home/z/my-project/download/chunk_results_v3.json') as f:
    v3data = j.load(f)
for c in v3data['pp']:
    if c['id'] in ('PP:70:1991:2', 'PP:70:1991:7:4'):
        print(f"  [{c['id']}] status={c['metadata'].get('status')} path={c['metadata'].get('path')}")
        print(f"    text: {c['text'][:150]}")
        print()