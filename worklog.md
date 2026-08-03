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
