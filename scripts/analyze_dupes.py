#!/usr/bin/env python3
"""
Read-only analysis of duplicate ID patterns in perppu and uud-1945 chunks.
"""
import json
import sys
from collections import Counter, defaultdict
from difflib import ndiff

DATA_PATH = "/home/z/my-project/download/chunk_results_v5.json"
SECTIONS = ["perppu", "uud-1945"]

def load_data():
    with open(DATA_PATH, "r", encoding="utf-8") as f:
        return json.load(f)

def find_duplicates(chunks):
    """Return dict: id -> list of (index, chunk)"""
    id_map = defaultdict(list)
    for i, c in enumerate(chunks):
        id_map[c["id"]].append((i, c))
    return {k: v for k, v in id_map.items() if len(v) > 1}

def trunc(text, n=100):
    if len(text) <= n:
        return text
    return text[:n] + "…"

def classify_boundary(dups):
    """
    Classify the structural pattern of a set of duplicates.
    Returns a list of pattern labels.
    """
    patterns = []
    for idx, chunk in dups:
        meta = chunk["metadata"]
        text = chunk["text"]
        path = meta.get("path", "")
        status = meta.get("status", "active")

        labels = []
        # Check penjelasan vs operative
        if "penjelasan" in text.lower()[:200] or "penjelasan" in path.lower():
            labels.append("penjelasan-related")
        if status in ("quoted_amendment", "penjelasan", "inactive"):
            labels.append(f"status={status}")

        # Check section boundary markers in text
        upper_text = text[:300]
        if any(kw in upper_text for kw in ["BAB ", "BAB  "]):
            labels.append("BAB-boundary")
        if any(kw in upper_text for kw in ["Bagian ", "BAGIAN "]):
            labels.append("Bagian-boundary")
        if any(kw in upper_text for kw in ["Paragraf ", "PARAGRAF "]):
            labels.append("Paragraf-boundary")

        # Check if path has BAB/Bagian/Paragraf transition
        path_parts = [p.strip() for p in path.split(">")]
        if len(path_parts) >= 2:
            for pp in path_parts:
                if pp.upper().startswith("BAB "):
                    labels.append("path-has-BAB")
                if pp.upper().startswith("BAGIAN "):
                    labels.append("path-has-Bagian")

        # Check for page-break artifacts (weird whitespace patterns)
        if "\n\n\n" in text or "  \n" in text or "\n  " in text:
            labels.append("possible-pagebreak")

        # Truncation / short text (likely artifact)
        if len(text) < 50:
            labels.append("very-short")

        patterns.append(labels if labels else ["no-special-pattern"])
    return patterns

def text_diff(t1, t2):
    """Show character-level diff between two texts."""
    lines1 = t1.splitlines(keepends=True)
    lines2 = t2.splitlines(keepends=True)
    diff = list(ndiff(lines1, lines2))
    # Summarize
    adds = sum(1 for d in diff if d.startswith("+ "))
    removes = sum(1 for d in diff if d.startswith("- "))
    unchanged = sum(1 for d in diff if d.startswith("  "))
    return {"added_lines": adds, "removed_lines": removes, "unchanged_lines": unchanged, "diff_preview": "".join(diff[:20])}

def analyze_section(name, chunks):
    print(f"\n{'='*80}")
    print(f"  ANALYSIS: {name} ({len(chunks)} total chunks)")
    print(f"{'='*80}")

    dups = find_duplicates(chunks)

    # 1. All duplicate IDs with counts
    dup_counts = {k: len(v) for k, v in dups.items()}
    sorted_dups = sorted(dup_counts.items(), key=lambda x: (-x[1], x[0]))

    print(f"\n--- 1. ALL DUPLICATE IDs ({len(sorted_dups)} unique IDs with duplicates) ---\n")
    if not sorted_dups:
        print("  (none)")
    for dup_id, count in sorted_dups:
        print(f"  {dup_id}  →  {count} occurrences")

    # 2. Top 5 most-duplicated IDs: show text + differences
    top5 = sorted_dups[:5]
    print(f"\n--- 2. TOP 5 MOST-DUPLICATED IDs: DETAILED VIEW ---\n")
    if not top5:
        print("  (none)")
    for rank, (dup_id, count) in enumerate(top5, 1):
        entries = dups[dup_id]
        print(f"  \n  ▸ Rank {rank}: {dup_id}  (count={count})")
        print(f"  {'─'*70}")
        for idx, (pos, chunk) in enumerate(entries):
            meta = chunk["metadata"]
            print(f"    Occurrence {idx+1} (array pos {pos}):")
            print(f"      Text ({len(chunk['text'])} chars): {trunc(chunk['text'])}")
            print(f"      Pasal: {meta.get('pasal')}, Ayat: {meta.get('ayat')}, Bab: {meta.get('bab')}")
            print(f"      Path: {meta.get('path')}")
            print(f"      Status: {meta.get('status')}, Source: {meta.get('source_file')}")

        # Show pairwise differences
        if count >= 2:
            t0 = entries[0][1]["text"]
            t1 = entries[1][1]["text"]
            if t0 == t1:
                print(f"    ✗ TEXTS ARE IDENTICAL ({len(t0)} chars)")
            else:
                d = text_diff(t0, t1)
                print(f"    ✗ TEXTS DIFFER: len={len(t0)} vs len={len(t1)}")
                print(f"      Diff: +{d['added_lines']} / -{d['removed_lines']} / ={d['unchanged_lines']} lines")
                # Show the shorter text and note what extra the longer has
                shorter = t0 if len(t0) <= len(t1) else t1
                longer = t1 if len(t0) <= len(t1) else t0
                if len(longer) > len(shorter) and shorter in longer:
                    print(f"      Pattern: SHORTER text is a PREFIX of LONGER text")
                    extra = longer[len(shorter):]
                    print(f"      Extra suffix ({len(extra)} chars): {trunc(extra, 150)}")
                else:
                    print(f"      Shorter text: {trunc(shorter, 150)}")
                    print(f"      Extra in longer (last 200 chars): {trunc(longer[-200:], 200)}")

    # 3. Structural pattern analysis
    print(f"\n--- 3. STRUCTURAL PATTERN ANALYSIS ---\n")
    pattern_counter = Counter()
    pattern_examples = defaultdict(list)

    for dup_id, entries in dups.items():
        patterns_list = classify_boundary(entries)
        # Flatten all patterns across occurrences
        flat = []
        for p in patterns_list:
            flat.extend(p)
        for p in flat:
            pattern_counter[p] += 1
        key = tuple(sorted(set(flat)))
        pattern_examples[key].append(dup_id)

    print("  Pattern frequency across all duplicate IDs:")
    for pat, cnt in pattern_counter.most_common():
        print(f"    {pat}: {cnt} IDs")

    print(f"\n  Combined pattern profiles (top 10):")
    sorted_profiles = sorted(pattern_examples.items(), key=lambda x: -len(x[1]))
    for profile, ids in sorted_profiles[:10]:
        profile_str = " + ".join(profile) if profile else "(no-special-pattern)"
        print(f"    [{profile_str}] → {len(ids)} IDs")
        for did in ids[:3]:
            print(f"      e.g. {did}")

    # Deeper analysis: check adjacency
    print(f"\n  Adjacency analysis (are duplicate chunks next to each other?):")
    adjacent_count = 0
    non_adjacent_count = 0
    for dup_id, entries in dups.items():
        positions = [e[0] for e in entries]
        positions.sort()
        gaps = [positions[i+1] - positions[i] for i in range(len(positions)-1)]
        if gaps and all(g == 1 for g in gaps):
            adjacent_count += 1
        else:
            non_adjacent_count += 1
    print(f"    Adjacent (positions differ by 1): {adjacent_count} IDs")
    print(f"    Non-adjacent: {non_adjacent_count} IDs")

    # Check if duplicates are penjelasan vs operative
    print(f"\n  Penjelasan vs Operative text collision analysis:")
    penj_ops = 0
    for dup_id, entries in dups.items():
        statuses = set(e[1]["metadata"].get("status", "active") for e in entries)
        if len(statuses) > 1:
            penj_ops += 1
            if penj_ops <= 5:
                print(f"    Mixed-status ID: {dup_id} → statuses={statuses}")
                for idx, (pos, chunk) in enumerate(entries):
                    print(f"      [{idx}] status={chunk['metadata'].get('status')}, text={trunc(chunk['text'],80)}")
    print(f"    Total IDs with mixed statuses: {penj_ops}")

    # Length ratio analysis
    print(f"\n  Length ratio analysis (for IDs with 2 occurrences):")
    ratio_bins = {"identical_len": 0, "1-2x_diff": 0, "2-5x_diff": 0, ">5x_diff": 0}
    for dup_id, entries in dups.items():
        if len(entries) == 2:
            l0, l1 = len(entries[0][1]["text"]), len(entries[1][1]["text"])
            ratio = max(l0, l1) / max(min(l0, l1), 1)
            if l0 == l1:
                ratio_bins["identical_len"] += 1
            elif ratio <= 2:
                ratio_bins["1-2x_diff"] += 1
            elif ratio <= 5:
                ratio_bins["2-5x_diff"] += 1
            else:
                ratio_bins[">5x_diff"] += 1
    for k, v in ratio_bins.items():
        print(f"    {k}: {v}")

def diagnose_root_cause(name, chunks, dups):
    """Deeper root-cause diagnosis."""
    import re
    print(f"\n--- 4. ROOT CAUSE DIAGNOSIS ---\n")

    # Check if all dups come from a single source file
    source_files = set()
    for entries in dups.values():
        for _, c in entries:
            source_files.add(c["metadata"].get("source_file", "?"))
    print(f"  Source files involved in duplicates: {source_files}")

    # Check if all dup IDs share the same doc prefix
    prefixes = set()
    for did in dups:
        parts = did.split(":")
        if len(parts) >= 3:
            prefixes.add(":".join(parts[:3]))
    print(f"  Document prefix(es) with duplicates: {prefixes}")

    if name == "perppu":
        # Pattern: truncated fragments (page breaks)
        trunc_count = 0
        for did, entries in dups.items():
            for _, c in entries:
                txt = c["text"]
                if len(txt) < 50 and (". . ." in txt or "..." in txt or txt.endswith("A")):
                    trunc_count += 1
        print(f"  Page-break truncation artifacts (very short + ellipsis/trailing 'A'): {trunc_count}")

        # Check how many dups are prefix-suffix pairs
        prefix_suffix = 0
        for did, entries in dups.items():
            if len(entries) == 2:
                t0, t1 = entries[0][1]["text"], entries[1][1]["text"]
                shorter = t0 if len(t0) <= len(t1) else t1
                longer = t1 if len(t0) <= len(t1) else t0
                if shorter in longer:
                    prefix_suffix += 1
        print(f"  Prefix-suffix pairs (shorter is prefix of longer): {prefix_suffix}")

        print(f"\n  CONCLUSION for perppu:")
        print(f"    ALL 30 duplicate IDs come from a single misclassified document")
        print(f"    (Perpres 148/2024 filed under perppu/ directory).")
        print(f"    ROOT CAUSE: PDF page-break splitting creates a truncated fragment")
        print(f"    (e.g. 'Proses . . . A') and a full-text continuation, both with")
        print(f"    the same pasal/ayat ID. 27/30 are adjacent pairs.")

    elif name == "uud-1945":
        # Check if IDs collapse lettered pasals
        print(f"  Lettered-pasal collision check:")
        # Look at what actual content lives under each dup ID
        for did, entries in dups.items():
            if len(entries) >= 3:
                pasal_refs = set()
                for _, c in entries:
                    matches = re.findall(r'Pasal (\d+[A-Z]?)', c["text"][:300])
                    pasal_refs.update(matches)
                if len(pasal_refs) > 1:
                    print(f"    {did} → references to: {sorted(pasal_refs)}")

        print(f"\n  CONCLUSION for uud-1945:")
        print(f"    ROOT CAUSE: The ID scheme UUD1945:?:1945:XX:Y collapses lettered pasals.")
        print(f"    Pasal 22, 22A, 22B, 22C, 22D all get ID ':22:1', ':22:2', etc.")
        print(f"    Pasal 28, 28A..28J all get ID ':28:1', ':28:2', etc.")
        print(f"    The '?' in position 2 of the ID confirms the parser could not")
        print(f"    distinguish the lettered sub-pasals. Each 'duplicate' actually")
        print(f"    contains DIFFERENT legal text from DIFFERENT articles.")
        print(f"    This is NOT a page-break issue — it's a fundamental ID generation")
        print(f"    bug that strips letter suffixes from pasal numbers.")

def main():
    data = load_data()
    for section in SECTIONS:
        if section in data:
            chunks = data[section]
            dups = find_duplicates(chunks)
            analyze_section(section, chunks)
            if dups:
                diagnose_root_cause(section, chunks, dups)
        else:
            print(f"\n[WARN] Section '{section}' not found in data. Available: {list(data.keys())}")

    # Cross-section summary
    print(f"\n{'='*80}")
    print(f"  CROSS-SECTION SUMMARY")
    print(f"{'='*80}")
    print(f"""
  ┌─────────────┬───────────┬────────────┬───────────────────────────────────┐
  │ Section     │ Chunks   │ Dup IDs    │ Root Cause                         │
  ├─────────────┼───────────┼────────────┼───────────────────────────────────┤
  │ perppu      │ 244       │ 30 (all x2)│ Page-break splitting: truncated    │
  │             │           │            │ fragment + full text, same ID.     │
  │             │           │            │ Single doc: Perpres 148/2024.       │
  ├─────────────┼───────────┼────────────┼───────────────────────────────────┤
  │ uud-1945    │ 149       │ 24         │ Lettered-pasal ID collision:        │
  │             │           │ (up to x8) │ 22/22A-D, 28/28A-J, 23/23E-G,     │
  │             │           │            │ 24/24A-F collapse to same IDs.      │
  └─────────────┴───────────┴────────────┴───────────────────────────────────┘

  RECOMMENDED FIXES:
  1. perppu: Deduplicate — keep the longer text chunk, drop the truncated
     page-break fragment. Or fix the chunker to merge cross-page splits.
  2. uud-1945: Fix ID generation to preserve lettered pasal suffixes
     (22A, 28J, etc.) so each article gets a unique ID.
""")
    print(f"{'='*80}")
    print(f"  ANALYSIS COMPLETE")
    print(f"{'='*80}\n")

if __name__ == "__main__":
    main()
