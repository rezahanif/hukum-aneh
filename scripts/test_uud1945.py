#!/usr/bin/env python3
"""End-to-end test: extract → parse → chunk → store for uud-1945.pdf"""

import sys, os
sys.path.insert(0, os.path.dirname(__file__))

from chunker.extractor import extract_with_fonts, detect_body_font_size, extract_raw_text, file_hash
from chunker.parser import parse_statute, debug_line_roles
from chunker.store import init_db, store_parsed_document, print_summary

PDF_PATH = "/home/z/my-project/download/samples/uud-1945/uud_1945.pdf"

print("=" * 60)
print("TEST: uud-1945.pdf")
print("=" * 60)

# Step 1: Extract with font info
print("\n[1] Extracting PDF with font metadata...")
lines = extract_with_fonts(PDF_PATH)
print(f"    {len(lines)} lines extracted")

body_size = detect_body_font_size(lines)
print(f"    Body font size: {body_size}")

# Show font size distribution
from collections import Counter
sizes = Counter(l["font_size"] for l in lines)
print(f"    Font sizes: {dict(sizes.most_common())}")

# Step 2: Debug — show line roles
print("\n[2] Line-by-line classification (first 80 lines):")
print("-" * 60)
display_lines = lines[:80]
for i, line_info in enumerate(display_lines):
    text = line_info["text"].strip()
    if not text:
        continue
    fs = line_info["font_size"]
    bold = "B" if line_info["is_bold"] else " "
    marker = " >>" if fs > body_size else "   "
    print(f"  {i:4d} [{bold}] fs={fs:5.1f}{marker} | {text[:90]}")
print(f"    ... ({len(lines) - 80} more lines)")

# Step 2b: Debug with parser roles
print("\n[2b] Parser role classification:")
print("-" * 60)
debug_line_roles(lines, body_size)

# Step 3: Parse
print("\n[3] Parsing document structure...")
from chunker.parser import ParsedDocument
parsed = parse_statute(lines, body_size, filename="uud_1945.pdf")
print(f"    BABs found: {len(parsed.babs)}")
print(f"    Pasals found: {len(parsed.all_pasals)}")
print(f"    Loose Pasals (no BAB): {len(parsed.loose_pasals)}")
total_ayat = sum(len(p.ayat_list) for p in parsed.all_pasals)
print(f"    Total Ayat: {total_ayat}")
print(f"    Preamble length: {len(parsed.preamble_text)} chars")
print(f"    Issues: {len(parsed.issues)}")
if parsed.issues:
    for iss in parsed.issues:
        print(f"      ! {iss}")

# Step 4: Show parsed Pasals
print("\n[4] Parsed Pasals (showing first 10):")
print("-" * 60)
for p in parsed.all_pasals[:10]:
    ayat_info = f"{len(p.ayat_list)} ayat" if len(p.ayat_list) > 1 else f"{len(p.ayat_list)} ayat"
    preview = p.raw_text[:100].replace("\n", " ")
    print(f"    Pasal {p.number:>4s} ({ayat_info}): {preview}...")
if len(parsed.all_pasals) > 10:
    print(f"    ... and {len(parsed.all_pasals) - 10} more")

# Step 5: Show Pasal 1 in detail (definitions)
print("\n[5] Pasal 1 — full detail:")
print("-" * 60)
if parsed.all_pasals:
    p1 = parsed.all_pasals[0]
    print(f"    Pasal {p1.number} ({len(p1.ayat_list)} ayat)")
    for a in p1.ayat_list[:5]:
        print(f"      ({a.number}) {a.text[:120]}...")
    if len(p1.ayat_list) > 5:
        print(f"      ... +{len(p1.ayat_list) - 5} more ayat")

# Step 6: Show Pasal with ayat sub-chunks
print("\n[6] Pasal 5 — pasal vs ayat chunks:")
print("-" * 60)
for p in parsed.all_pasals:
    if p.number == "5":
        print(f"    [PASAL CHUNK] tokens~{len(p.raw_text)//4}:")
        print(f"    {p.raw_text[:200]}")
        print(f"")
        for a in p.ayat_list:
            print(f"    [AYAT ({a.number})] tokens~{len(a.text)//4}:")
            print(f"    {a.text[:200]}")
            print(f"")
        break

# Step 7: Store in SQLite
print("\n[7] Storing in SQLite...")
conn = init_db()
raw_text = extract_raw_text(PDF_PATH)
fhash = file_hash(PDF_PATH)
doc_id = store_parsed_document(conn, parsed, "uud-1945", raw_text, fhash, reg_number="1945")
print(f"    Document ID: {doc_id}")

# Step 8: Print summary
print_summary(conn)

# Step 9: Verify — query a few chunks
print("\n[8] Sample queries:")
print("-" * 60)
c = conn.cursor()

# Show first 3 pasal chunks
rows = c.execute("SELECT id, hierarchy_path, substr(content, 1, 80), token_count FROM chunks WHERE chunk_level='pasal' LIMIT 3").fetchall()
for r in rows:
    print(f"    [{r[0]}] {r[1]}")
    print(f"        {r[2]}... ({r[3]} tokens)")
    print()

# Show ayat sub-chunks for Pasal 5
rows = c.execute("""SELECT id, hierarchy_path, substr(content, 1, 80), token_count 
                     FROM chunks WHERE chunk_level='ayat' AND pasal_num='5'
                     ORDER BY ayat_num""").fetchall()
for r in rows:
    print(f"    [{r[0]}] {r[1]}")
    print(f"        {r[2]}... ({r[3]} tokens)")
    print()

conn.close()
print("\n[DONE] Test complete.")
