# Chunking Pipeline Phase 0–4 Summary

## What was built

### Phase 0: Extraction Quality

**0.1 Character-encoding corruption**
- Diagnosed all 16 samples. Confirmed glyph substitution patterns: `O`→`0` (2OO9→2009), `l`→`1`, `REPLJBLIK`→REPUBLIK, `PRESIOEN`→PRESIDEN
- Built `fix_glyph_corruption()` with 7 regex-based fixes in `clean_extractor.py`
- tap_mpr: confirmed scanned/image PDF (0 extractable tokens), needs OCR

**0.2 Reading-order + boilerplate**
- Built `detect_and_strip_boilerplate()`: detects lines repeating across ≥40% of pages
- Built `is_page_number_line()` and `is_likely_heading_or_preamble()` for header/footer stripping
- Multi-column detection tested — all 15 samples are single-column (x_center variance is from indent levels, not columns)
- Font size explicitly NOT used as primary structural signal (per plan requirement)

**0.3 Validation harness**
- 16 golden-set samples in `/download/samples/`
- `chunk_results_v2.json` has full output for hand-checking

### Phase 1: Family Classification

`config/directory_family_map.json` — 3 families:

| Family | Directories | Chunk unit |
|--------|-----------|------------|
| **A: Hierarchical Statute** | uu, pp, perppu, perpres, perda, uud-1945, JDIH_Kemnaker, JDIH_Kemenkeu, JDIH_Komdigi, JDIH_KPU, peraturan | Ayat (fallback: Pasal) |
| **B: Decree/Decision** | keppres, inpres, tap_mpr, JDIH_Kemendag | Diktum point |
| **C: Court Ruling** | Putusan-MK | Amar point |

### Phase 2: Family A Parser

State machine with heading stack: `BAB > Bagian > Paragraf > Pasal > Ayat > huruf/angka`

Key features:
- Preamble auto-detection (exits on first BAB/Pasal/MENETAPKAN/MEMUTUSKAN)
- Penutup section detection (skips closing/repeal listings)
- Quoted amendment detection (skips old pasal text quoted in amendment clauses)
- Partial hierarchy support (Pasal > Ayat without BAB)
- Roman-numeral Pasal support (Pasal I, II — used in UUD-1945)

### Phase 3: Family B Parser

Section-based: `Menimbang > Mengingat > MEMUTUSKAN/MENETAPKAN/MENGINSTRUKSIKAN > Diktum`

Key features:
- Decision keyword regex allows short prefix (e.g. `dengan ini menginstruksikan:`)
- Item markers: `1.`, `a.`, AND Indonesian ordinals (KESATU, KEDUA, KETIGA, ...)
- Guard against re-triggering on keywords inside diktum text
- Signature-block detection to end substantive content

### Phase 4: Family C Parser

Court ruling: `Pre > MENGADILI/MENETAPKAN > numbered Amar > signature block`

Key features:
- Extracts `nomor_perkara` (e.g. `96/PUU-XVI/2018`) as primary ID
- Amar chunks with full legal-reasoning text

### Phase 5: Metadata Enrichment

- `doc_type` + `issuer`: detected from title block (page 1, first 10 non-boilerplate lines)
- `nomor` + `year`: extracted from title block only (not full text, to avoid reference matches)
- `family`: from directory config (not auto-detected, per plan)
- `path`: full breadcrumb trail (e.g. `BAB III > Ketentuan Umum > Pasal 5 > Ayat (2)`)
- `parent`: set for ayat chunks (e.g. `UU:41:2014:5`)
- `status`: default `active`; amendment/repeal cross-referencing noted as known gap

## Results

| Directory | Family | Chunks | Doc Type | Year |
|-----------|--------|--------|----------|------|
| uu | A | 135 | UU | 2014 |
| pp | A | 63 | PP | 1991 |
| perppu | A | 292 | Perpres* | 2024 |
| perpres | A | 7 | Perpres | ? |
| perda | A | 93 | Perda | 2025 |
| keppres | B | 9 | Keppres | 2015 |
| inpres | B | 41 | Inpres | 2023 |
| tap_mpr | B | 0 | (scanned) | - |
| uud-1945 | A | 175 | UUD1945 | - |
| Putusan-MK | C | 3 | Putusan_MK | 2018 |
| JDIH_Kemnaker | A | 4 | PP* | 2013 |
| JDIH_Kemenkeu | A | 10 | Permen | 2025 |
| JDIH_Kemendag | B | 6 | Kepmen | 2025 |
| JDIH_Komdigi | A | 162 | Permen | 2024 |
| JDIH_KPU | A | 1 | PKPU | ? |
| peraturan | A | 3 | PP | 1962 |
| **TOTAL** | | **1005** | | |

*perppu: file title says "PERATURAN PRESIDEN" but is actually a Perppu — title-based detection overrides directory config
*Kemnaker: filename says "Permenaker" but content says "PERATURAN PEMERINTAH" — title-based detection correct

## Known Issues / Next Steps

1. **Quoted amendment text** — some amendment clauses (e.g. UU pasal 13 quote) still leak through. The RE_AMENDMENT_TRIGGER needs the quote delimiters to appear on the SAME LINE as the trigger phrase, but PDF line-breaking often splits them.

2. **Garbled text in some chunks** — glyph corruption not fully caught. The `FTRESIDEN R Ei:IUE I- IK IND ONES IA` line (UU pasal 13) is a severe corruption case that needs OCR fallback (Phase 0.1 confidence scoring).

3. **JDIH_KPU: page 1 missing** — the PDF starts extraction at page 2. Only 1 chunk produced. Likely a cover page or image-based page 1.

4. **perpres: `?:?` nomor/year** — `NOMOR I27` has `I` (letter) instead of `1` (digit). The glyph fix only handles `O`→`0`, not `I`→`1`.

5. **Kemendag: signature block leaked** — "Ditetapkan di Jakarta..." appears as diktum. Need stronger signature-block detection.

6. **Duplicate pasal chunks** — definition pasals (Pasal 1) sometimes produce multiple chunks when the definition list (1., 2., 3.) gets split incorrectly.

7. **JDIH directories may mix types** — some folders contain both Peraturan (Family A) and Keputusan (Family B). Per-file detection needed (not just per-directory).

8. **tap_mpr needs OCR** — 0 text extractable, requires image-based OCR pipeline.

## Files

| File | Purpose |
|------|---------|
| `scripts/clean_extractor.py` | PDF text extraction with glyph fix, boilerplate removal, doc_type detection |
| `scripts/parsers.py` | Family A/B/C parsers + unified `chunk_pdf()` entry point |
| `config/directory_family_map.json` | Directory → family/type/issuer mapping |
| `download/skeletons_raw.json` | Raw skeleton data (all lines with font info) |
| `download/phase0_diagnosis.json` | Encoding/boilerplate/multicolumn diagnosis per file |
| `download/chunk_results_v2.json` | Full chunk output for all 16 samples |
| `download/dokumen_skeleton_per_direktori.md` | Full text skeleton per directory (510KB) |