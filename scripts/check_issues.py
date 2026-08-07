#!/usr/bin/env python3
"""Check specific chunk issues in chunk_results_v5.json."""
import json
import sys
import re

DATA_PATH = "/home/z/my-project/download/chunk_results_v5.json"

with open(DATA_PATH, "r", encoding="utf-8") as f:
    data = json.load(f)

divider = "=" * 80

# ============================================================
# 1. UU chunks with ID containing ":13" or pasal 13
# ============================================================
print(divider)
print("SECTION 1: UU chunks with ':13' in ID or pasal 13")
print(divider)

uu_chunks = data.get("uu", [])
found_13 = []
for c in uu_chunks:
    cid = c.get("id", "")
    pasal = c.get("metadata", {}).get("pasal", "")
    if ":13" in cid or str(pasal) == "13":
        found_13.append(c)

if found_13:
    for c in found_13:
        print(f"\n--- ID: {c['id']} ---")
        print(f"    pasal={c['metadata'].get('pasal')}, ayat={c['metadata'].get('ayat')}, status={c['metadata'].get('status')}")
        print(f"    TEXT: {c['text']}")
else:
    print("  No chunks found with ':13' in ID or pasal==13.")
    # Fallback: search text for 'Pasal 13'
    for c in uu_chunks:
        if "Pasal 13" in c.get("text", ""):
            print(f"\n  [fallback text match] ID: {c['id']}")
            print(f"    TEXT: {c['text'][:300]}")

# ============================================================
# 2. Perpres chunks with 'Lampiran' or 'Rp' or Pasal 7
# ============================================================
print()
print(divider)
print("SECTION 2: Perpres chunks with 'Lampiran' / 'Rp' / Pasal 7")
print(divider)

pp_chunks = data.get("perpres", [])
found_pp = []
for c in pp_chunks:
    txt = c.get("text", "")
    pasal = c.get("metadata", {}).get("pasal", "")
    cid = c.get("id", "")
    if "Lampiran" in txt or re.search(r'Rp', txt) or str(pasal) == "7":
        found_pp.append(c)

# Filter to only those with Rp (currency)
rp_chunks = [c for c in found_pp if re.search(r'Rp', c.get("text", ""))]
lampiran_chunks = [c for c in found_pp if "Lampiran" in c.get("text", "")] if not rp_chunks else []
pasal7_chunks = [c for c in found_pp if str(c.get("metadata", {}).get("pasal", "")) == "7"] if not rp_chunks and not lampiran_chunks else []

if rp_chunks:
    print(f"\n  Found {len(rp_chunks)} chunk(s) with 'Rp' (currency):")
    for c in rp_chunks:
        print(f"\n--- ID: {c['id']} ---")
        print(f"    pasal={c['metadata'].get('pasal')}, ayat={c['metadata'].get('ayat')}, status={c['metadata'].get('status')}")
        print(f"    TEXT: {c['text']}")

if lampiran_chunks:
    print(f"\n  Found {len(lampiran_chunks)} chunk(s) with 'Lampiran':")
    for c in lampiran_chunks:
        print(f"\n--- ID: {c['id']} ---")
        print(f"    pasal={c['metadata'].get('pasal')}, ayat={c['metadata'].get('ayat')}, status={c['metadata'].get('status')}")
        print(f"    TEXT: {c['text']}")

if pasal7_chunks:
    print(f"\n  Found {len(pasal7_chunks)} chunk(s) with pasal 7:")
    for c in pasal7_chunks:
        print(f"\n--- ID: {c['id']} ---")
        print(f"    pasal={c['metadata'].get('pasal')}, ayat={c['metadata'].get('ayat')}, status={c['metadata'].get('status')}")
        print(f"    TEXT: {c['text']}")

if not rp_chunks and not lampiran_chunks and not pasal7_chunks:
    print("  No chunks found matching Lampiran/Rp/Pasal 7.")
    # Show all perpres chunk IDs for reference
    print(f"  Total perpres chunks: {len(pp_chunks)}")
    for c in pp_chunks:
        print(f"    ID: {c['id']}  pasal={c['metadata'].get('pasal')}  text[:80]={c['text'][:80]}")

# ============================================================
# 3. Inpres duplicate IDs
# ============================================================
print()
print(divider)
print("SECTION 3: Inpres duplicate IDs")
print(divider)

inpres_chunks = data.get("inpres", [])
print(f"  Total inpres chunks: {len(inpres_chunks)}")

id_counts = {}
for c in inpres_chunks:
    cid = c.get("id", "")
    id_counts.setdefault(cid, []).append(c)

dupes = {k: v for k, v in id_counts.items() if len(v) > 1}
if dupes:
    print(f"  Found {len(dupes)} duplicate ID(s):")
    for cid, chunks in dupes.items():
        print(f"\n  --- Duplicate ID: {cid} (count={len(chunks)}) ---")
        for i, c in enumerate(chunks):
            print(f"    [{i}] pasal={c['metadata'].get('pasal')}, ayat={c['metadata'].get('ayat')}, status={c['metadata'].get('status')}")
            print(f"        text[:150]: {c['text'][:150]}")
else:
    print("  No duplicate IDs found in inpres chunks.")
    # Show all IDs for reference
    for c in inpres_chunks:
        print(f"    ID: {c['id']}  text[:150]={c['text'][:150]}")

print()
print(divider)
print("DONE")
print(divider)
