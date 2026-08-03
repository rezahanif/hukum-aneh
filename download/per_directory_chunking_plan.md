# Per-Directory Chunking Plan

Based on 1 sample PDF per directory, skeleton analysis of wording patterns, numbering systems, and decision anchors.

---

## Structural Family Classification

After analyzing all 15 directories, the documents fall into **5 structural families** (not 15):

| Family | Directories | Section Divider | Decision Anchor | Chunk Boundary |
|--------|-------------|-----------------|-----------------|-----------------|
| **A. Pasal-Ayat Statutes** | uu, pp, perppu, perda, perpres, JDIH_Komdigi, JDIH_Kemnaker, JDIH_Kemenkeu | `BAB > Bagian > Paragraf > Pasal > (N)` | `Pasal N` heading + `(N)` ayat | Per-ayat |
| **B. Pasal-Prose (short)** | keppres, peraturan | `Pasal N` (no ayat, continuous prose) | `Pasal N` (whole pasal = 1 chunk) | Per-pasal |
| **C. Ordinal-Based** | inpres, JDIH_Kemendag | `KESATU:`, `KEDUA:`, `KETIGA:` ... | Each ordinal = 1 decision point | Per-ordinal |
| **D. Court Rulings** | Putusan-MK | `MENETAPKAN:` / `MEMUTUSKAN:` then numbered items | AMAR (verdict) section only | Per-verdict-point |
| **E. Amendment Regs** | JDIH_KPU, uu (Perubahan) | `Pasal I, II...` (Roman) with amendment list | Each numbered amendment item | Per-amendment-item |

---

## Universal Patterns

### Preamble (Families A, B, C, E)

All statute-type documents share this structure:

```
PRESIDEN REPUBLIK INDONESIA / MENTERI ...
DENGAN RAHMAT TUHAN YANG MAHA ESA     <- skip
PRESIDEN REPUBLIK INDONESIA,         <- skip
Menimbang                           <- SKIP ZONE START
  a. bahwa...
  b. bahwa...  (contains "Pasal X" references = false positives)
Mengingat                           <- still skip zone
  1. Pasal X ayat (Y)...             <- inline references (false positives)
  2. ...
MEMUTUSKAN:                        <- PREAMBLE END TRIGGER
Menetapkan: [DOC TYPE] TENTANG...   <- skip (title repeat)
```

**Rule**: Everything before `MEMUTUSKAN:` or `Menetapkan :` is preamble. Do NOT parse Pasal headings here.

### Closing (Families A, B, E)

```
KETENTUAN PERALIHAN                <- body content (keep)
  Pasal 125, 126...
KETENTUAN PENUTUP                  <- body content (keep)
  Pasal 127, 128...
Agar setiap orang mengetahuinya   <- SKIP: signing formula
Ditetapkan di Jakarta...            <- SKIP: signing block
[signature lines]                    <- SKIP
PENJELASAN                          <- SECONDARY ZONE (handle separately)
  PASAL DEMI PASAL                  <- per-pasal explanations
```

---

## Family A: Pasal-Ayat Statutes

**Directories**: uu, pp, perppu, perda, perpres, JDIH_Komdigi, JDIH_Kemnaker, JDIH_Kemenkeu

### Wording Pattern

```
BAB I [TITLE]
  Bagian Kesatu [TITLE]
    Paragraf 1 [TITLE]
      Pasal 1
        (1) Content of first ayat...
        (2) Content of second ayat...
            a. sub-item
            b. sub-item
        (3) ...
```

### The Decision = Pasal body content
Each ayat `(N)` is one operative rule. The hierarchy (BAB/Bagian/Paragraf) provides context.

### Chunking Strategy

**Parent chunk** = Full Pasal (all ayats joined)
- Purpose: Questions about the whole article
- Content: Pasal heading + all ayat text
- Metadata: BAB, Bagian, Paragraf, Pasal number

**Child chunk** = Per-Ayat
- Purpose: Precise RAG retrieval
- Content: Single ayat text (strip the `(N)` marker from content)
- Metadata: hierarchy path + ayat number

### Parser Rules for Family A

1. **Skip preamble**: Find `MEMUTUSKAN:` or `Menetapkan`, start parsing after it
2. **Detect body end**: `Agar setiap orang` or `PENJELASAN` or signing authority name
3. **Pasal heading**: `^\s*Pasal\s+\d+[A-Za-z]?\s*\.?\s*\.\.\.?\s*$`
4. **Ayat boundary**: `^\s*\(\s*\d+[A-Za-z]?\s*\)\s*`
5. **Sub-items**: `^[a-z]\)` or `^\d+\)` (absorb into current ayat, don't create new chunk)
6. **Definition pasals** (Pasal 1 often): If no ayat markers but has `1. 2. 3.` numbered items, treat each numbered item as a pseudo-ayat
7. **Strip ayat marker**: Remove leading `(N)` from ayat content text

### Per-Directory Notes

| Directory | Specific Notes |
|-----------|---------------|
| **uu** | May be "Perubahan" (amendment) — has Pasal I (Roman) + quoted Pasal N (Arabic) inside. Handle as Family E intersection. |
| **pp** | Standard pattern. PENJELASAN with "PASAL DEMI PASAL" present. Ayat markers sometimes on own line (merge with next). |
| **perppu** | Same as UU structure but issued by President. Very large (130+ pasals). No PENJELASAN in sample. |
| **perda** | Same as PP but regional. Has PENJELASAN. Font is uniform 12.0 — regex only. |
| **perpres** | Short (5-10 pasals). Many pasals have NO ayat — continuous prose. Treat whole pasal as 1 chunk. May have LAMPIRAN (table). |
| **JDIH_Komdigi** | Standard Permen pattern. 44 pasals. Uniform font. Has KETENTUAN PERALIHAN + PENUTUP. |
| **JDIH_Kemnaker** | Actually a PP (issued by President, filed under Kemnaker). 66 pasals. Very long. Has PENJELASAN. |
| **JDIH_Kemenkeu** | Short PMK (8 pasals). Mixed: some pasals have ayat, some are continuous prose. Has LAMPIRAN with table data. |

---

## Family B: Pasal-Prose (Short)

**Directories**: keppres, peraturan

### Wording Pattern

```
Pasal 1
Dewan Kawasan Kawasan Ekonomi Khusus Provinsi Kalimantan Timur
terdiri atas Ketua dan Anggota...

Pasal 2
Dewan Kawasan bertanggung jawab kepada Dewan Nasional...

Pasal 3
Biaya yang timbul dari pelaksanaan...

Pasal 4
Keputusan Presiden ini mulai berlaku...
```

### The Decision = Each Pasal as a whole (no ayat subdivision)
These are short documents (3-6 pasals) where each pasal is one continuous paragraph.

### Chunking Strategy

Each Pasal = 1 chunk (no parent/child split needed)
- If pasal has a numbered list inside (e.g., membership list `1. Ketua 2. Anggota...`), keep it all in one chunk
- Do NOT try to split by `1.` or `a.` — these are sub-items, not chunk boundaries

### Parser Rules

1. Same preamble skip as Family A
2. Pasal heading: `^\s*Pasal\s+\d+\s*\.?\s*$`
3. No ayat detection needed
4. Chunk = everything from `Pasal N` to next `Pasal M` or closing block
5. Old-style (peraturan/1962): Uses `Mendengar:` and `Memutuskan :` instead of `Menimbang`/`MEMUTUSKAN:`

---

## Family C: Ordinal-Based Decisions

**Directories**: inpres, JDIH_Kemendag

### Wording Pattern

```
INSTRUKSI PRESIDEN REPUBLIK INDONESIA
NOMOR 3 TAHUN 2023
TENTANG [TITLE]

PRESIDEN REPUBLIK INDONESIA,

Kepada:
1. Menteri PPN/Bappenas;
2. Menteri PUPR;
3. Menteri Keuangan;
...

KESATU:
[General instruction text with sub-items]

Khusus kepada:
1. Menteri PPN/Bappenas untuk:
   a. [specific task]
   b. [specific task]
2. Menteri PUPR untuk:
   a. [specific task]
   b. [specific task]
...

KEDUA:
[Effective date or additional instruction]

KETIGA:
[Closing provision]
```

### The Decision = Each ordinal section (KESATU, KEDUA...)
These documents have NO Pasal structure. The decisions are the instruction items.

### Chunking Strategy

**Parent chunk** = Per-ordinal section (KESATU, KEDUA...)
**Child chunks** = Per-instruction-item (each `a.` `b.` sub-item under each official)

For **Inpres**: The per-official instructions are the most granular actionable unit
For **Kemendag Kepmen**: Usually only 1-3 ordinals, very short. Each ordinal = 1 chunk.

### Parser Rules

1. No preamble skip needed (Inpres has no Menimbang; Kepmen has standard preamble)
2. Section heading: `^\s*(KESATU|KEDUA|KETIGA|KEEMPAT|KELIMA|KEENAM|KETUJUH|KEDELAPAN|KESEMBILAN|KESEPULUH)\s*:`
3. Sub-section: `Khusus kepada:` then numbered officials
4. Sub-items: `a.`, `b.`, `c.` under each official
5. Recipient list at top: skip or store as metadata

### Per-Directory Notes

| Directory | Specific Notes |
|-----------|---------------|
| **inpres** | No Menimbang preamble. Starts with title then recipient list. Instructions directed to specific officials. |
| **JDIH_Kemendag** | Standard Menimbang preamble. Very short (3 ordinals). Often just sets a price or quota. |

---

## Family D: Court Rulings (Putusan-MK)

**Directories**: Putusan-MK

### Wording Pattern

```
SALINAN
KETETAPAN / PUTUSAN
Nomor 96/PUU-XVI/2018

DEMI KEADILAN BERDASARKAN KETUHAN YANG MAHA ESA
MAHKAMAH KONSTITUSI REPUBLIK INDONESIA,

Menimbang:
1. Bahwa MK telah menerima permohonan...
2. Bahwa pemohon menguji...
...

Mengingat:
1. UUD 1945;
2. UU 24/2003 tentang MK;
...

MENETAPKAN:                          <- AMAR START
1. Mengabulkan permohonan...;
2. Menyatakan Undang-Undang...
3. Memerintahkan...

Demikian diputus dalam Rapat...     <- AMAR END
[Nama-nama Hakim Konstitusi]
[Tanggal, jam sidang]

KETUA,
[signature]
ANGGOTA-ANGGOTA,
[signatures]
PANITERA PENGGANTI,
[signature]
```

### The Decision = AMAR (verdict section only)

For RAG, only the AMAR matters — it contains the court's binding decision. The considerations (Menimbang) and legal basis (Mengingat) are secondary.

**For full Putusan** (not just Ketetapan), the structure would also include:
- `PERTIMBANGAN HUKUM` — legal reasoning (ratio decidendi)
- A more complex AMAR with multiple categories

### Chunking Strategy

**Primary chunk** = AMAR section (MENETAPKAN: to "Demikian diputus")
- Each numbered verdict point = 1 child chunk

**Secondary chunk** = PERTIMBANGAN HUKUM (if present)
- The legal reasoning — useful for understanding WHY the court decided

**Metadata chunk** = Case info (number, parties, provisions challenged)

### Parser Rules

1. No Pasal detection needed
2. AMAR trigger: `^\s*(MENETAPKAN|MEMUTUSKAN)\s*:`
3. AMAR end: `^\s*Demikian diputus`
4. Verdict points: numbered `1.`, `2.`, `3.` after AMAR trigger
5. Verdict action keywords: `Mengabulkan`, `Menyatakan`, `Memerintahkan`, `Menolak`, `Tidak Menerima`

---

## Family E: Amendment Regulations (Perubahan)

**Directories**: JDIH_KPU, uu (when it's a Perubahan UU)

### Wording Pattern

```
Pasal I
Beberapa ketentuan dalam [base regulation] diubah sebagai berikut:

1. Ketentuan Pasal 20 ayat (5) diubah, sehingga berbunyi:
   "Pasal 20
   (1) [replacement text for ayat 1]
   (2) [replacement text for ayat 2]
   (5) [replacement text for ayat 5]"

2. Ketentuan Pasal 21 ayat (1) diubah, sehingga berbunyi:
   "Pasal 21
   (1) [replacement text]"

3. Ketentuan Pasal 22 ayat (4) diubah...

Pasal II
Peraturan ini mulai berlaku...
```

### The Decision = Each numbered amendment item

Each `1.`, `2.`, `3.` under Pasal I is a separate amendment action. The quoted text is the NEW wording that replaces the old.

### Chunking Strategy

**Parent chunk** = Per-Roman-Pasal (Pasal I, Pasal II...)
**Child chunk** = Each numbered amendment item
- Contains: amendment instruction + full quoted replacement text
- Key metadata: which TARGET pasal/ayat is being amended

### Parser Rules

1. Pasal heading: `^\s*Pasal\s+[IVXLCDMivxlcdm]+\s*$` (Roman only)
2. Amendment items: `^\s*\d+\.` after Pasal I heading
3. Quoted text: Track `"` open/close to capture the replacement text
4. Inside quotes: `(N)` = TARGET regulation's ayat, NOT the amending regulation's ayat
5. Pasal II, III... = closing provisions (effective date, etc.) — treat as Family B

---

## Special Cases

### LAMPIRAN (Attachments)
Some documents (JDIH_Kemenkeu, perpres) reference a Lampiran containing tables or detailed data.
- Extract separately from the body
- Parse as structured data (table) rather than legal text
- Link to the pasal that references it (e.g., "Pasal 3: Besaran tercantum dalam Lampiran")

### PENJELASAN (Explanations)
Present in: uu, pp, perda, JDIH_Kemnaker (and sometimes others)
- Handle as SECONDARY chunks (lower priority for RAG)
- Trigger: standalone `PENJELASAN` or `PASAL DEMI PASAL` header
- Often has "Cukup jelas" (self-explanatory) for many pasals
- Structure: per-pasal explanation, linked to body pasal by number

### UUD 1945 (Constitution)
Special case: Contains original 1945 text + 4 amendments (Perubahan I-IV).
- Same pasal numbers appear in multiple versions
- Must tag each pasal with its version (original, Perubahan I, II, III, or IV)
- Chunking: Per-ayat, with version tag in metadata

---

## Summary: One Parser Per Family, Not Per Directory

| Family | Parser Class | Core Regex | Key Trigger |
|--------|-------------|------------|-------------|
| A | `StatuteParser` | `Pasal\s+\d+[A-Z]?` + `\(\d+\)` | `MEMUTUSKAN:` starts body |
| B | `ShortStatuteParser` | `Pasal\s+\d+` (no ayat) | Same preamble skip |
| C | `OrdinalParser` | `(KESATU|KEDUA|...)\s*:` | `KESATU:` starts decisions |
| D | `CourtRulingParser` | `(MENETAPKAN\|MEMUTUSKAN)\s*:` | AMAR section extraction |
| E | `AmendmentParser` | `Pasal\s+[IVXLCDM]+` | Roman numeral pasal detection |

Only **5 parser classes** needed for all 15 directories. Within each family, per-directory configuration handles minor variations (presence of PENJELASAN, LAMPIRAN, etc.).