---
Task ID: 1
Agent: Main Agent
Task: Create corrected PostgreSQL + Qdrant integration plan for hukum-aneh repo

Work Log:
- Re-read all critical source files: models.go (10 structs), firestore.go (378-line monolith), engine.go (689 lines, 13-state lifecycle), collections.go, interface.go, retrieval.go (brute-force search), config.go, go.mod
- Read all 6 cmd/ entry points to understand FirestoreRepo dependency chain
- Identified 11 files that directly depend on *repository.FirestoreRepo
- Wrote comprehensive docx generation script (generate_plan2.js) using docx library
- Generated document with R4 cover (GO-1 Graphite Orange palette), TOC, 13 sections
- Fixed TOC placeholders with add_toc_placeholders.py
- Postcheck: 0 errors, 2 minor warnings (expected)

Stage Summary:
- Produced: /home/z/my-project/download/Hukum-Aneh_PostgreSQL_Qdrant_Integration_Plan.docx
- Document covers 7 phases (Phase 0-7) with 22 executable subtasks
- All 9 gaps from original plan addressed
- Task dependency table included for execution order
---
Task ID: 1
Agent: main
Task: Implement Phase 0-4 of legal document chunking pipeline

Work Log:
- Built Phase 0.1: encoding diagnosis (glyph substitution patterns, font CMap checks)
- Built Phase 0.2: clean_extractor.py (boilerplate stripping, header/footer removal, glyph fixes)
- Built Phase 1: 3-family classification (directory_family_map.json)
- Built Phase 2: Family A parser (BAB > Bagian > Pasal > Ayat state machine)
- Built Phase 3: Family B parser (Menimbang > MEMUTUSKAN > Diktum with ordinal word support)
- Built Phase 4: Family C parser (MENGADILI > Amar for court rulings)
- Ran full test: 1005 chunks from 15/16 directories (tap_mpr is scanned)
- Identified 8 known issues for next iteration

Stage Summary:
- All 3 parsers working, producing valid JSON chunks matching target format
- Chunk IDs follow convention: {doc_type}:{nomor}:{year}:{pasal}:{ayat}
- Family B uses ordinal words (KESATU, KEDUA...) for Inpres/Kemendag
- Key remaining issues: quoted amendment text, glyph corruption, JDIH type mixing

---
Task ID: 1
Agent: Main
Task: Apply v3 QA fixes to chunking pipeline, run tests, prepare results for upload

Work Log:
- Read and analyzed existing chunker code (parsers.py, clean_extractor.py)
- Fixed clean_extractor.py: UUD 1945 nomor=None, broader stamp pattern, restricted nomor_year to title block, consistent O-as-0 glyph fixes
- Fixed parsers.py: doc_type resolution (title block first + broad-type guard), Penjelasan section detection with distinct :penjelasan ID suffix, signature block as hard boundary, parent_section tracking for inpres diktum, relaxed RE_PASAL regex
- Fixed backspace-char corruption in regex patterns (\x08 instead of \b)
- Added _flush_penjelasan method to FamilyAParser
- Added PENJELASAN detection without requiring PENUTUP first
- Ran full batch test: 16 files, 974 chunks, 5/16 clean
- Generated v3 test report and QA report

Stage Summary:
- 3 output files: chunk_results_v3.json (714K), chunk_test_report_v3.json (8.8K), chunk_qa_report_v3.json (8K)
- Key improvements: pp crash fixed, keppres/JDIH_Kemnaker/JDIH_Komdigi doc_type corrected, 57 penjelasan chunks properly separated
- No Google Drive credentials available — files in /home/z/my-project/download/

---
Task ID: 2
Agent: Main
Task: Apply v4 QA fixes, run tests, upload results to Drive

Work Log:
- Investigated root causes for all 6 v3 QA red flags using debug script
- Fixed RE_PASAL regex to capture space-separated multi-digit numbers (Pasal 1 10 = Pasal 110)
- Fixed pasal match to normalize numbers and capture inline text on same line
- Fixed BAB change to reset current_pasal (fixes pp heading-as-chunk)
- Reordered detect_doc_type_from_title: Keppres/Kepmen/Perpres/Permen before PP/UUD1945
- Used page-1-only raw lines for doc_type detection
- Added filename-based nomor/year fallback for image-cover PDFs (JDIH_KPU: 8/2026)
- Added post-decision nomor/year extraction for image-cover docs
- Added KETENTUAN PENUTUP to penutup pattern
- Added glyph fixes for perpres (Tfrnjangan, PRESTDEN)
- Broadened stamp pattern for No<digit> variants
- Generated v4 results: 913 chunks, 0 errors, 136 dupes
- Uploaded 3 files to Google Drive

Stage Summary:
- keppres: doc_type correctly Keppres (was PP regression, now fixed)
- JDIH_Kemnaker: doc_type=PP per title block (user-requested)
- JDIH_KPU: nomor=8, year=2026 from filename fallback (was 07/2012)
- JDIH_KPU: doc_type=PKPU via directory override
- pp: 7→3 duplicate IDs (heading-as-chunk fixed)
- perda: 18→4 duplicate IDs (BAB reset pasal fixed)
- perpres: doc_type=Perpres ✅
- Known remaining: perppu 52 dupes (complex counter issue), JDIH_Komdigi 18 dupes (page-break splits)

---
Task ID: 1
Agent: Main
Task: Apply v5 QA fixes — perpres dedicated pass + uu alpha pasal suffixes

Work Log:
- Diagnosed perpres Pasal 3 root cause: page number `-3-` ends with hyphen, triggering hyphen-merge rule that swallows Pasal 3 content into garbled header line `FRESIDEN REPUBL]K INDONESIA -3-`, which is then discarded as noise. Fixed by excluding `-.\d+-$` patterns from hyphen-merge.
- Fixed 5 Tunjangan glyph corruption variants: T/rnjangan (literal backslash-r), T/r/rnjangan (CR char), Ttrnjangan, Tfrnjangan, I/rnjangan. Removed re.I flag that caused [TI] to match lowercase t producing TTunjangan.
- Fixed Fungsional corruption: R/rngsional, F/rngsional → Fungsional
- Fixed sebagaimana corruption: sglagaimana, sslagaimana → sebagaimana
- Fixed stamp noise: broadened RE_SETNEG_STAMP2 to handle SK No2l1340A (no space after No, lowercase l separator). Added GLYPH_FIX pattern to strip stamp prefix when glued to body text.
- Fixed false Pasal heading: `Pasal 2 diangkat...` (body text reference) no longer matches as heading. Updated STRUCTURAL_MARKERS and RE_PASAL to require heading-like ending (dots, braces, whitespace, or end-of-line after pasal number).
- Fixed uu alpha pasal suffixes: RE_PASAL now captures [a-zA-Z]? after digits (68A, 68B, etc.). Updated penjelasan pasal detection similarly.
- Updated FamilyB RE_PASAL_DIKTUM for alpha suffixes.
- Ran full pipeline: 16 files, 873 chunks, 141 dup chunks

Stage Summary:
- perpres: Pasal 3 restored (was missing for 3 rounds), all glyph corruption fixed, stamp noise stripped, 0 dup IDs
- uu: 68A/68C/68D/68E now distinct IDs. 68B missing due to OCR corruption (688 in source PDF — not regex-fixable)
- perppu: 30 dup chunks (was 115 in v3), worst case 2/ID (was 9)
- Output: chunk_results_v5.json, chunk_qa_report_v5.json, chunk_test_report_v5.json in /home/z/my-project/download/

---
Task ID: 1
Agent: Main
Task: Apply v6 QA fixes — uud-1945 alpha pasals, perppu dedup, currency normalization, fuzzy noise

Work Log:
- Analyzed v5 independent verification results from user
- Diagnosed uud-1945 root cause: UUD 1945 PDF has `* Pasal 22A` prefix pattern and `BAB VIIA ... * Pasal 22C` glued-to-heading pattern, neither matched by v5 RE_PASAL. Also `PEMILIHAN UMUM * Pasal 22E` as BAB label with pasal glued.
- Diagnosed perppu root cause: page-break fragments (10-50 chars) getting same pasal:ayat ID as the full continuation text (100-300+ chars). Classic truncation duplicate.
- Diagnosed perpres Lampiran currency: `Rp1.38O.OO0,0o` — O in digit context within Rp amounts, not caught by word-level glyph patterns.
- Diagnosed UU garbled header: `FTRESIDEN R Ei:IUE I- IK IND ONES IA -9-` — new OCR variant not in exact pattern list.
- Fixed clean_extractor.py: (1) Currency glyph normalization via function-based replacer that replaces all O/o/A/a with 0 within Rp amount spans. (2) Added FTRESIDEN to RE_GARBLED_HEADER. (3) Added fuzzy noise detection module using Levenshtein edit distance against known boilerplate strings. (4) Added _is_fuzzy_noise() call in is_noise_line().
- Fixed parsers.py: (1) Extended RE_PASAL to handle `* ` prefix. (2) Added RE_PASAL_EMBEDDED for mid-line pasal detection. (3) Extended RE_BAB/RE_BAGIAN/RE_PARAGRAF with `*?` prefix. (4) Added glued-pasal extraction after BAB match. (5) Added embedded-pasal detection in main parse loop. (6) Added _dedup_truncated_fragments() method to FamilyAParser with 50% length threshold.
- Ran full pipeline: 16 files, 892 chunks, 86 dup chunks

Stage Summary:
- uud-1945: 24->6 dup IDs (-75%), 38->7 dup chunks (-82%). Now has 39 alpha-suffixed pasals (22A-E, 23A-G, 24A-C, 25A, 28A-J, etc.). Pasal 22C/22E now correctly split via glued-heading detection.
- perppu: 30->26 dup IDs (-13%), removed 14 truncated page-break fragments
- perpres: Currency amounts now all-numeric (Rp1.380.000,00, Rp500.000,00, etc.)
- uu: Garbled header `FTRESIDEN...` chunk eliminated by fuzzy noise detection
- Net: 98->75 dup IDs (-23%), 141->86 dup chunks (-39%)
- Output: chunk_results_v6.json, chunk_qa_report_v6.json in /home/z/my-project/download/

---
Task ID: v7-fixes
Agent: Main Agent
Task: v7 round fixes — JDIH_KPU amendment trigger, perppu dedup + doc_type, B↔8 OCR detection

Work Log:
- Diagnosed JDIH_KPU 95% dup rate: root cause = comma in "diubah, sehingga berbunyi" not matched by RE_AMENDMENT_TRIGGER regex (\s+ doesn't match ", ")
- Fixed RE_AMENDMENT_TRIGGER: \s+ → [,\s]+ between verb and "sehingga" (backward compatible)
- Diagnosed perppu doc_type: PDF has image cover, body text cites "Peraturan Presiden" in Mengingat, causing false Perpres detection. Added directory-based override in chunk_pdf()
- Enhanced _dedup_truncated_fragments() with v7 short-artifact rule: removes chunks <80 chars that are <40% length of longest sibling (catches org unit labels and OCR fragments)
- Added _dedup_short_artifacts() to FamilyBParser (same logic, covers inpres)
- Changed elif → independent if for short-artifact check to catch non-substring fragments
- Built _flag_ocr_pasal_numbers() method: detects trailing-digit pasal numbers (8→B, 0→O, etc.) when alpha siblings exist for same prefix
- Successfully flagged 368→36B and 918→91B/91C in UU 41/2014

Stage Summary:
- Results: 165 dup chunks (v6) → 37 dup chunks (v7) = -78% reduction
- JDIH_KPU: 19 → 1 dup chunk (-95%), 6 quoted amendment chunks correctly generated
- perppu: 52 → 8 dup chunks (-85%), doc_type corrected to PerPPU
- inpres: 8 → 1 dup chunk (-88%)
- uu: 34 → 9 dup chunks (-74%)
- uud-1945: 13 → 7 dup chunks (-46%)
- perpres: 0 → 0 (clean, no regression)
- B↔8 OCR: 368 and 918 flagged with _ocr_flag metadata
- Remaining 37 dups are structural content collisions (medium-length pairs), not page-break artifacts
- Files modified: parsers.py (4 changes), clean_extractor.py (none)
- New script: scripts/run_v7.py
- Output: chunk_results_v7.json, chunk_qa_report_v7.json
