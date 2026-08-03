#!/usr/bin/env python3
"""Analyze Pasal/Ayat pattern variations across all samples."""
import os, re, fitz, json
from collections import defaultdict

SAMPLE_DIR = "/home/z/my-project/download/samples"

pasal_variations = defaultdict(list)
ayat_variations = defaultdict(list)
bab_variations = defaultdict(list)

for folder in sorted(os.listdir(SAMPLE_DIR)):
    folder_path = os.path.join(SAMPLE_DIR, folder)
    if not os.path.isdir(folder_path):
        continue
    for fname in sorted(os.listdir(folder_path)):
        if not fname.endswith('.pdf'):
            continue
        fpath = os.path.join(folder_path, fname)
        try:
            doc = fitz.open(fpath)
            for page in doc:
                text = page.get_text()
                for line in text.split('\n'):
                    s = line.strip()
                    if not s:
                        continue
                    # Pasal patterns
                    if re.search(r'Pasal\s+\d+', s, re.IGNORECASE):
                        # Capture the exact raw string before the number
                        m = re.match(r'(.*?Pasal\s*\d+.*)', s, re.IGNORECASE)
                        if m:
                            raw = m.group(1)[:80]
                            key = re.sub(r'\s+', ' ', raw).lower()
                            pasal_variations[folder].append((raw, repr(raw)))
                    # Ayat patterns
                    if re.match(r'^\s*\(?\s*\d+\s*\)?', s):
                        ayat_variations[folder].append(repr(s[:60]))
                    # BAB patterns
                    if re.search(r'BAB\s+', s, re.IGNORECASE):
                        bab_variations[folder].append(repr(s[:80]))
            doc.close()
        except:
            pass

print("="*80)
print("PASAL VARIATIONS PER FOLDER")
print("="*80)
for folder, examples in pasal_variations.items():
    print(f"\n--- {folder} ({len(examples)} matches) ---")
    # Show unique patterns
    seen = set()
    for raw, rep in examples[:20]:
        normalized = re.sub(r'\s+', ' ', raw).lower()
        if normalized not in seen:
            seen.add(normalized)
            print(f"  {rep}")

print("\n" + "="*80)
print("AYAT VARIATIONS PER FOLDER")
print("="*80)
for folder, examples in ayat_variations.items():
    print(f"\n--- {folder} ({len(examples)} matches) ---")
    seen = set()
    for rep in examples[:10]:
        norm = re.sub(r'\s+', ' ', rep)
        if norm not in seen:
            seen.add(norm)
            print(f"  {rep}")

print("\n" + "="*80)
print("BAB VARIATIONS PER FOLDER")
print("="*80)
for folder, examples in bab_variations.items():
    print(f"\n--- {folder} ({len(examples)} matches) ---")
    seen = set()
    for rep in examples[:10]:
        norm = re.sub(r'\s+', ' ', rep)
        if norm not in seen:
            seen.add(norm)
            print(f"  {rep}")
