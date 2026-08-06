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
