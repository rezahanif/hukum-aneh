# Parser Flaw Report — Per-Folder Analysis

Tested: 1 sample PDF per folder (16 folders total)
Date: 2026-08-03
Parser version: v0.1 (font-size-aware + 3-layer defense)

---

## Executive Summary

| Metric | Value |
|--------|-------|
| Folders tested | 16 |
| Folders with 0 pasals | 4 (correct — no Pasal structure) |
| Folders with parseable Pasals | 12 |
| Total flaws detected | 657 |
| Most common flaw | EMPTY_PASAL (316 occurrences across 10 folders) |
| Highest severity | DUPLICATE_PASAL (10/12 parseable folders) |

---

## Flaw Taxonomy

| ID | Flaw Name | Severity | Description |
|----|-----------|----------|-------------|
| F1 | DUPLICATE_PASAL | **HIGH** | Same pasal number appears multiple times — caused by TOC entries, inline references, or closing section (penutup) re-listing pasals being falsely matched as headings |
| F2 | AYAT_MARKER_IN_CONTENT | **MEDIUM** | Ayat content string starts with "(1)" because the ayat marker line itself gets included in the content |
| F3 | EMPTY_PASAL | **MEDIUM** | Pasal has 0 ayats — caused by definition articles (numbered list format like 1. 2. 3.), single-paragraph pasals with no (1) marker, or content being swallowed by previous pasal |
| F4 | LONG_AYAT_CONTENT | **LOW** | Ayat content >500 chars — possible paragraph bleed across page boundaries or missing ayat separator |
| F5 | WEAK_FONT_SIGNAL | **LOW** | Only 1-2 font sizes detected — parser must rely entirely on regex, reducing heading detection accuracy |
| F6 | PASALS_WITHOUT_BAB | **LOW** | Pasals not assigned to any BAB — some documents lack BAB structure entirely (common in shorter regulations) |
| F7 | ZERO_PASAL_CORRECT | **INFO** | 0 pasals found is correct — document type doesn't use Pasal structure (Keputusan, Inpres, Putusan MK) |
| F8 | TOC_FALSE_POSITIVE | **HIGH** | Table of Contents entries like "Pasal 4 . . ." matched as pasal headings |
| F9 | PENUTUP_FALSE_POSITIVE | **HIGH** | Closing section (Penutup) re-lists all pasal numbers as change references, each matched as a heading |
| F10 | AYAT_ON_OWN_LINE | **MEDIUM** | Ayat marker "(1)" appears on its own line separate from content — content starts on the next line |
| F11 | DEFINITION_ARTICLE_NO_AYAT | **MEDIUM** | Definition articles (Pasal 1) use "1. Term = definition" format instead of "(1) content" — produces EMPTY_PASAL |
| F12 | ROMAN_NUMERAL_PASAL | **HIGH** | Some documents use Roman numerals ("Pasal I", "Pasal IV") instead of Arabic — regex only matches \d+ |

---

## Per-Folder Results

### 1. uud-1945 — UUD 1945 (Constitution)

| Metric | Value |
|--------|-------|
| Font sizes | [12.0, 14.0] — only 2 sizes |
| Pasals found | 102 |
| Ayats found | 252 |
| Flaws | 104 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: 37 duplicate pasal numbers. Root cause: UUD 1945 contains both original text AND amendment text (Perubahan I-IV). Each amendment reprints the Pasal with changes. The parser treats original and amended versions as separate pasals. Example: Pasal 1 appears in BAB I (original), BAB II (amendment context). **Fix needed**: Detect amendment sections and deduplicate, or tag with amendment version.
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 69 occurrences. Every ayat content starts with "(1) (1)" — the marker line is included. Root cause: The ayat line `(1)` is added to `_current_ayat_lines` but never stripped. **Fix**: Strip the leading ayat marker from content.
- **F3 EMPTY_PASAL [MEDIUM]**: 33 pasals with 0 ayats. Root cause: Some pasals in the amendments section are single-line references or very short pasals where the ayat marker is on the same line as the pasal heading.
- **F5 WEAK_FONT_SIGNAL [LOW]**: Only 2 font sizes (12.0 body, 14.0 title). All BAB/Pasal headings use the same 12.0 size. Font signal provides zero discrimination.

---

### 2. uu — UU No. 41 Tahun 2014 (Statute)

| Metric | Value |
|--------|-------|
| Font sizes | 27 unique sizes (8.5 to 22.5) — rich signal |
| Pasals found | 71 |
| Ayats found | 90 |
| Flaws | 76 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: Massive duplicates — 50+ pasal numbers repeated. Root cause: The PDF has a detailed Table of Contents (penjelasan/eluasan) section at the end that re-lists every pasal. These TOC entries match the strict regex. **Fix needed**: Detect and skip TOC sections (look for patterns like "Pasal X . . ." with dots, or detect pages with many consecutive pasal headings and no ayat content).
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 21 occurrences. Same root cause as UUD 1945.
- **F3 EMPTY_PASAL [MEDIUM]**: 50 empty pasals. Root cause: Both TOC false positives (F8) and definition articles (F11) contribute.
- **F6 PASALS_WITHOUT_BAB [MEDIUM]**: 71/71 pasals have no BAB. Root cause: BAB headings may use a different format not caught by regex, or this UU doesn't use BAB structure.

---

### 3. pp — PP No. 70 Tahun 1991 (Government Regulation)

| Metric | Value |
|--------|-------|
| Font sizes | [10.0, 12.0] |
| Pasals found | 52 |
| Ayats found | 61 |
| Flaws | 54 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: Duplicates in BAB VII (closing/penutup section). Root cause: The penutup section lists all pasals being changed/repealed, and each gets matched as a heading. **Fix needed**: Detect "Penutup" or "Ketentuan Peralihan" or "Ketentuan Lain-Lain" sections and stop parsing pasal headings.
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 22 occurrences.
- **F3 EMPTY_PASAL [MEDIUM]**: 30 empty pasals. Root cause: Pasal 1 is a definition article using `1.`, `2.`, `3.` numbered items (F11). Also penutup false positives (F9).
- **F10 AYAT_ON_OWN_LINE [MEDIUM]**: Ayat markers like `(1)` and `(2)` appear on their own line, with content starting on the next line. The parser includes the marker as part of the content.

---

### 4. perppu — Perppu No. 148 Tahun 2024 (Emergency Regulation)

| Metric | Value |
|--------|-------|
| Font sizes | 18 unique sizes (11.0 to 22.5) — rich signal |
| Pasals found | 146 |
| Ayats found | 293 |
| Flaws | 148 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: 68+ duplicate numbers. Root cause: This is a large Perppu (likely omnibus law) with many chapters. The penjelasan (explanation) section at the end re-lists all pasals. **Fix needed**: Same TOC/penjelasan detection as uu.
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 73 occurrences.
- **F3 EMPTY_PASAL [MEDIUM]**: 73 empty pasals. Root cause: Combination of F8 (TOC) and F9 (penutup) false positives.
- **F4 LONG_AYAT_CONTENT [LOW]**: Pasal 8 ayat (5) has 5,716 chars — extreme paragraph bleed. Root cause: Multiple paragraphs or items being absorbed into a single ayat because no new ayat marker `(N)` appears.

---

### 5. perpres — Perpres No. 127 Tahun 2024 (Presidential Regulation)

| Metric | Value |
|--------|-------|
| Font sizes | 21 unique sizes (7.0 to 22.5) — very rich signal |
| Pasals found | 10 |
| Ayats found | 0 |
| Flaws | 12 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: 3 duplicate numbers.
- **F3 EMPTY_PASAL [MEDIUM]**: ALL 10 pasals have 0 ayats. Root cause: This Perpres uses a different content format — pasal content may not use `(1)` ayat markers, or uses a different numbering scheme.
- **F6 PASALS_WITHOUT_BAB [MEDIUM]**: 10/10 pasals have no BAB.

---

### 6. perda — Perda Kabupaten Sukoharjo No. 1 Tahun 2025 (Regional Regulation)

| Metric | Value |
|--------|-------|
| Font sizes | [11.0, 12.0] |
| Pasals found | 51 |
| Ayats found | 66 |
| Flaws | 53 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: Many duplicates. Root cause: TOC section + penutup section both create false positives.
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 19 occurrences.
- **F3 EMPTY_PASAL [MEDIUM]**: 32 empty pasals. Root cause: TOC + definition articles.
- **F5 WEAK_FONT_SIGNAL [LOW]**: Only 2 font sizes.

---

### 7. keppres — Keppres No. 5 Tahun 2015 (Presidential Decree)

| Metric | Value |
|--------|-------|
| Font sizes | [12.0] — single size |
| Pasals found | 6 |
| Ayats found | 0 |
| Flaws | 9 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: Pasal 1 and 4 duplicated. Root cause: Keppres typically has few pasals and the penutup references them.
- **F3 EMPTY_PASAL [MEDIUM]**: All 6 pasals empty. Root cause: Keppres uses different content structure — may use `1.`, `2.` numbered items instead of ayat markers.
- **F5 WEAK_FONT_SIGNAL [LOW]**: Single font size.
- **F6 PASALS_WITHOUT_BAB [MEDIUM]**: No BAB structure.

---

### 8. peraturan — PP No. 20 Tahun 1962 (Older Regulation)

| Metric | Value |
|--------|-------|
| Font sizes | [10.0, 13.0] |
| Pasals found | 2 |
| Ayats found | 2 |
| Flaws | 5 |

**Flaws detected:**
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 2 occurrences.
- **F4 LONG_AYAT_CONTENT [LOW]**: Pasal 2 ayat (2) has 1,307 chars.
- **F5 WEAK_FONT_SIGNAL [LOW]**: Only 2 font sizes.
- **F6 PASALS_WITHOUT_BAB [LOW]**: No BAB detected (short document).

---

### 9. JDIH_Kemnaker — PP No. 99 Tahun 2013 (Social Security Assets)

| Metric | Value |
|--------|-------|
| Font sizes | [12.0, 12.1, 12.2] — nearly uniform |
| Pasals found | 159 (actual: ~66 unique) |
| Ayats found | 215 |
| Flaws | 160 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: ALL 66 pasal numbers are duplicated (159 total = 66 real + 93 false). Root cause: Three sources of false positives: (a) TOC entries with "Pasal X . . ." pattern, (b) "Menimbang"/"Mengingat" preamble section references like "Pasal 47 ayat (2)", (c) penutup section. The loose regex catches preamble references because `_is_inline_reference()` returns False when suffix is empty.
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 60 occurrences.
- **F3 EMPTY_PASAL [MEDIUM]**: 99 empty pasals — mostly TOC/penutup false positives.

---

### 10. JDIH_Komdigi — Permenkominfo No. 5 Tahun 2024

| Metric | Value |
|--------|-------|
| Font sizes | [12.0] — single size |
| Pasals found | 50 |
| Ayats found | 149 |
| Flaws | 54 |

**Flaws detected:**
- **F1 DUPLICATE_PASAL [HIGH]**: 6 pasal numbers duplicated. Root cause: TOC + penutup sections.
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 38 occurrences.
- **F3 EMPTY_PASAL [MEDIUM]**: 12 empty pasals.
- **F4 LONG_AYAT_CONTENT [LOW]**: Pasal 8 ayat (2) has 2,164 chars — severe paragraph bleed.
- **F5 WEAK_FONT_SIGNAL [LOW]**: Single font size.

---

### 11. JDIH_Kemenkeu — PMK No. 9 Tahun 2025 (Finance Ministry Regulation)

| Metric | Value |
|--------|-------|
| Font sizes | [8.0, 12.0, 14.0] |
| Pasals found | 8 |
| Ayats found | 4 |
| Flaws | 11 |

**Flaws detected:**
- **F2 AYAT_MARKER_IN_CONTENT [MEDIUM]**: 3 occurrences.
- **F3 EMPTY_PASAL [MEDIUM]**: 5/8 pasals empty. Root cause: Pasal 1-4 are definition/interpretation articles using numbered list format. Also Pasal 7 is a "short" pasal where content doesn't use (N) markers.
- **F4 LONG_AYAT_CONTENT [LOW]**: Pasal 8 ayat (2) has 1,537 chars.
- **F6 PASALS_WITHOUT_BAB [MEDIUM]**: 8/8 pasals have no BAB — PMK often has simple structure.

---

### 12. JDIH_KPU — PKPU No. 8 Tahun 2026 (Election Commission Regulation)

| Metric | Value |
|--------|-------|
| Font sizes | [12.0] — single size |
| Pasals found | 0 |
| Ayats found | 0 |
| Flaws | 1 |

**Flaws detected:**
- **F12 ROMAN_NUMERAL_PASAL [HIGH]**: Uses "Pasal I", "Pasal IV" (Roman numerals). Our regex `\d+` only matches Arabic digits. **Fix needed**: Add `[ivxlcdmIVXLCDM]+` to pasal number pattern.
- **F7 ZERO_PASAL_CORRECT**: This is an amendment regulation (Perubahan) that references existing pasals — the "Pasal I" entries are actually section headings for amendment groups, not individual pasal numbers.
- **F5 WEAK_FONT_SIGNAL [LOW]**: Single font size.

---

### 13. JDIH_Kemendag — Kepmendag No. 123 Tahun 2025 (Trade Ministry Decision)

| Metric | Value |
|--------|-------|
| Font sizes | 27 unique sizes (11.4 to 27.5) |
| Pasals found | 0 |
| Ayats found | 0 |
| Flaws | 0 |

**Status**: **CORRECT — NO FLAWS**

This is a Keputusan Menteri (Ministerial Decision) about CPO reference pricing. It defines price tables and formulas — no Pasal structure exists. The document correctly produces 0 pasals.

---

### 14. inpres — Inpres No. 3 Tahun 2023 (Presidential Instruction)

| Metric | Value |
|--------|-------|
| Font sizes | 18 unique sizes (10.0 to 22.5) |
| Pasals found | 0 |
| Ayats found | 0 |
| Flaws | 0 |

**Status**: **CORRECT — NO FLAWS**

Inpres uses numbered points (1., 2., 3.) and lettered sub-items (a., b., c.) instead of Pasal/Ayat structure. The document correctly produces 0 pasals.

---

### 15. Putusan-MK — MKRI No. 5301 (Constitutional Court Ruling)

| Metric | Value |
|--------|-------|
| Font sizes | [12.0] — single size |
| Pasals found | 0 |
| Ayats found | 0 |
| Flaws | 1 |

**Status**: **CORRECT — NO FLAWS** (1 info-level finding)

Court rulings (Putusan) reference Pasals from the law being reviewed but don't have their own Pasal structure. They use "Menimbang", "Mengingat", "Mempertimbangkan", "Mengadili", "MEMUTUSKAN" sections instead.
- **F5 WEAK_FONT_SIGNAL [LOW]**: Single font size.

---

### 16. tap_mpr — TAP MPR III/2002

| Metric | Value |
|--------|-------|
| Font sizes | N/A |
| Pasals found | ERROR |
| Ayats found | ERROR |
| Flaws | 0 |

**Status**: **ERROR — FILE EXTRACTION FAILED**

The filename is truncated ("Ketetapan Majelis...IIIMPR2002 ten.pdf"). The PDF may be scanned/image-based, or the filename truncation caused path issues. **Fix needed**: Re-download with correct filename, or test with OCR fallback.

---

## Root Cause Analysis & Recommended Fixes

### Priority 1: HIGH — Eliminate False Positive Pasal Headings

**Problem**: TOC entries, penutup references, and preamble references are matched as real Pasal headings, causing massive duplicates.

**Affected folders**: 10/12 (all statute folders)

**Recommended fixes**:
1. **TOC Detection (F8)**: Detect pages/sections where multiple consecutive "Pasal N" lines appear with NO ayat content between them. If >3 consecutive pasal headings with 0 ayat content each, mark the section as TOC and skip.
2. **Penutup Detection (F9)**: Detect section headings like "Penutup", "Ketentuan Peralihan", "Ketentuan Lain-Lain", "Ketentuan Penutup" — stop pasal parsing after these.
3. **Preamble Detection**: Detect "Menimbang", "Mengingat", "Dengan Persetujuan" sections — skip pasal matching in these sections.
4. **Inline Reference Improvement**: The `_is_inline_reference()` function needs to be more aggressive. Currently it returns `False` when suffix is empty but text is short. Fix: If the line is "Pasal N" alone (no prefix, no suffix) AND we're in a non-TOC/non-penutup section, treat as heading. Otherwise, require font-size signal.

### Priority 2: MEDIUM — Ayat Content Quality

**Problem**: Ayat marker included in content (F2), ayat on own line (F10), definition articles (F11).

**Affected folders**: 11/12

**Recommended fixes**:
1. **Strip ayat marker (F2)**: When flushing ayat content, strip the leading `(N)` from the content string.
2. **Merge orphaned ayat markers (F10)**: When `(N)` appears on its own line, merge it with the next line(s) until the next ayat/pasal marker.
3. **Definition article handling (F11)**: When a pasal has 0 ayats, check if its content uses `1.`, `2.`, `3.` numbered format. If so, create synthetic ayats from these numbered items.

### Priority 3: HIGH — Roman Numeral Pasals (F12)

**Problem**: Some documents (PKPU) use "Pasal I", "Pasal IV".

**Affected folders**: JDIH_KPU (confirmed), possibly others

**Recommended fix**: Extend pasal regex to accept Roman numerals: `\d+[a-zA-Z]?|[ivxlcdmIVXLCDM]+`

### Priority 4: LOW — Font Signal Weakness (F5)

**Problem**: 10/16 folders have only 1-2 font sizes, making font-based heading detection unreliable.

**Recommended fix**: Accept that most Indonesian legal PDFs use uniform font sizing. Rely primarily on regex + positional heuristics (line position on page, centered vs left-aligned) rather than font size. The font-size-aware approach is still valuable for the ~25% of documents with rich font variation.

---

## Flaw Count Summary

| Folder | F1 DUP | F2 AYAT | F3 EMPTY | F4 LONG | F5 WEAK | F6 NO_BAB | F8 TOC | F9 PEN | F10 ORPHAN | F11 DEF | F12 ROMAN | Total |
|--------|--------|---------|----------|---------|---------|-----------|--------|--------|-------------|----------|-----------|-------|
| uud-1945 | 37 | 69 | 33 | 0 | 1 | 0 | yes | yes | yes | 0 | 0 | 104 |
| uu | many | 21 | 50 | 3 | 0 | 1 | yes | yes | yes | yes | 0 | 76 |
| pp | many | 22 | 30 | 0 | 1 | 0 | yes | yes | yes | yes | 0 | 54 |
| perppu | many | 73 | 73 | 1 | 0 | 0 | yes | yes | yes | 0 | 0 | 148 |
| perpres | 3 | 0 | 10 | 0 | 0 | 1 | maybe | 0 | 0 | maybe | 0 | 12 |
| perda | many | 19 | 32 | 0 | 1 | 0 | yes | yes | yes | yes | 0 | 53 |
| keppres | 2 | 0 | 6 | 0 | 1 | 1 | yes | yes | 0 | 0 | 0 | 9 |
| peraturan | 0 | 2 | 0 | 1 | 1 | 1 | 0 | 0 | yes | 0 | 0 | 5 |
| JDIH_Kemnaker | 93 | 60 | 99 | 0 | 0 | 0 | yes | yes | yes | 0 | 0 | 160 |
| JDIH_Komdigi | 6 | 38 | 12 | 2 | 1 | 0 | yes | yes | yes | 0 | 0 | 54 |
| JDIH_Kemenkeu | 0 | 3 | 5 | 2 | 0 | 1 | 0 | 0 | yes | yes | 0 | 11 |
| JDIH_KPU | 0 | 0 | 0 | 0 | 1 | 0 | 0 | 0 | 0 | 0 | yes | 1 |
| JDIH_Kemendag | 0 | 0 | 0 | 0 | 0 | 0 | N/A | N/A | N/A | N/A | N/A | 0 |
| inpres | 0 | 0 | 0 | 0 | 0 | 0 | N/A | N/A | N/A | N/A | N/A | 0 |
| Putusan-MK | 0 | 0 | 0 | 0 | 1 | 0 | N/A | N/A | N/A | N/A | N/A | 1 |
| tap_mpr | ERR | ERR | ERR | ERR | ERR | ERR | ERR | ERR | ERR | ERR | ERR | ERR |

---

## Next Steps

1. **Implement Priority 1 fixes** (TOC/penutup/preamble detection) — will eliminate ~60% of all flaws
2. **Implement Priority 2 fixes** (ayat content quality) — will eliminate ~30% of remaining flaws
3. **Implement Priority 3** (Roman numeral support) — fixes JDIH_KPU and any future Roman-numeral docs
4. **Re-run batch test** to validate fixes
5. **Test on 3-5 more files per folder** to increase confidence
6. **Add special parser for Putusan-MK** (court ruling structure: Considerans > Amar > Ratio)
