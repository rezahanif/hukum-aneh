# Indonesian Legal Document Library — PDF Structure Skeleton for Chunking

> Source: Google Drive folder `1vneHF9YxwgSnBh3ashORK0cYPo16vmQS`
> Purpose: Clause/Article-level chunking reference

---

## 1. uud-1945

**Full name:** Undang-Undang Dasar Negara Republik Indonesia Tahun 1945 (1945 Constitution, with Amendments I–IV)

**Files:** 1 PDF

### Hierarchy

```
PEMBUKAAN (Preamble)
├── BATANG TUBUH
│   ├── BAB I — Bentuk dan Kedaulatan (Form & Sovereignty)
│   ├── BAB II — Majelis Permusyawaratan Rakyat (MPR)
│   ├── BAB III — Kekuasaan Pemerintahan Negara (State Power)
│   ├── BAB IIIA — Komisi Yudisial
│   ├── BAB IV — Dewan Pertimbangan Agung [Repealed]
│   ├── BAB IV — Lembaga Negara Lainnya [Amendment III]
│   ├── BAB V — Hak dan Kewajiban Warga Negara (Citizen Rights & Duties)
│   ├── BAB VI — Agama
│   ├── BAB VII — Negara
│   ├── BAB VIIA — Wilayah Negara
│   ├── BAB VIIIA — Warga Negara dan Penduduk
│   ├── BAB IX — Kekuasaan Kehakiman (Judiciary)
│   ├── BAB IXA — Negara Kesatuan
│   ├── BAB X — Perubahan UUD (Amendments)
│   ├── BAB XA — Ketentuan Peralihan (Transitional Provisions)
│   └── BAB XB — Ketentuan Tambahan (Additional Provisions)
└── PENUTUP (Closing)
```

### Structural Elements per Pasal

| Element | Pattern | Example |
|---------|---------|---------|
| **Pasal** (Article) | `Pasal {N}` | Pasal 1 |
| **Ayat** (Clause) | `(1)`, `(2)`, `(3)` | (1) Negara Indonesia ialah... |
| **Huruf** (Letter point) | `a)`, `b)`, `c)` | a) bentuk republik |

**Chunking unit:** Each `Pasal` with all its `Ayat` → one chunk. Pasal 1 (definitions) may be larger and could be split per-Ayat.

---

## 2. perpres

**Full name:** Peraturan Presiden (Presidential Regulations)

**Files:** ~75 PDFs, 2022–2025

**Naming:** `perpres-no-{N}-tahun-{YYYY}.pdf`

### Hierarchy

```
JUDUL (Title)
├── DENGAN RAHMAT TUHAN YANG MAHA ESA
├── PRESIDEN REPUBLIK INDONESIA
├── Menimbang (Considering)
│   ├── a), b), c) ...
├── Mengingat (Bearing in mind)
│   ├── 1., 2., 3. ...
├── MEMUTUSKAN (Decides)
│   ├── MENETAPKAN (Enacts): PERATURAN PRESIDEN ...
├── Pasal 1
│   ├── (1), (2), (3) ...
├── Pasal 2
│   ├── (1), (2) ...
├── ... (typically 5–15 Pasal)
├── Pasal {N} — Penutup (Closing)
│   ├── (1) Peraturan ini mulai berlaku...
└── TTD Presiden (Presidential signature)
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **Pasal** | `^Pasal\s+\d+` | Article |
| **Ayat** | `^\(\d+\)\s*` | Clause within an article |
| **Angka** | `^\d+\.\s*` | Numbered list (legal references, items) |
| **Huruf** | `^[a-z]\)\s*` | Letter sub-point |

**Chunking unit:** One chunk per `Pasal` (including all `Ayat`). For Pasal 1 (definisi/definitions), consider per-Ayat chunks since definitions can be lengthy.

**Characteristics:**
- No `BAB` (Chapter) divisions — flat article sequence
- `Menimbang` + `Mengingat` preamble → can be a single “considering” chunk
- Shorter than UU (typically 5–20 Pasal)

---

## 3. perda

**Full name:** Peraturan Daerah (Regional/Municipal Regulations)

**Files:** ~105 PDFs

**Naming:** `perda-{kabupaten/kota}-{type}-no-{N}-tahun-{YYYY}.pdf`

### Hierarchy

```
JUDUL (Title: BUPATI/WALI KOTA + PROVINSI)
├── PERATURAN DAERAH KABUPATEN/KOTA {NAME}
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── DENGAN RAHMAT TUHAN YANG MAHA ESA
├── BUPATI/WALI KOTA ...
├── Menimbang
│   ├── a), b), c) ...
├── Mengingat
│   ├── 1. Undang-Undang ...
│   ├── 2. ...
├── Dengan Persetujuan Bersama DEWAN PERWAKILAN RAKYAT DAERAH ...
│   └── MEMUTUSKAN
│       └── MENETAPKAN: PERATURAN DAERAH ...
├── BAB I — KETENTUAN UMUM (General Provisions)
│   ├── Pasal 1 (Definisi — Definitions)
│   │   ├── 1. {term} adalah {definition}
│   │   ├── 2. ...
│   ├── Pasal 2, Pasal 3 ...
├── BAB II — {Subject-specific chapter}
│   ├── Pasal 4, Pasal 5 ...
├── ... (typically 8–15 BAB)
├── BAB {N} — KETENTUAN PENUTUP (Closing Provisions)
├── TTD Bupati/Wali Kota
└── PENJELASAN (Explanatory Notes — optional appendix)
    ├── UMUM (General explanation)
    └── PASAL DEMI PASAL (Article-by-article explanation)
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **BAB** | `^BAB\s+[IVXLCDM]+\s+` | Chapter |
| **Pasal** | `^Pasal\s+\d+` | Article |
| **Ayat** | `^\(\d+\)\s*` | Clause |
| **Huruf** | `^[a-z]\)\s*` | Letter sub-point |
| **Angka** | `^\d+\.\s*` | Numbered list |

**Chunking unit:** One chunk per `Pasal` (all `Ayat`). Pasal 1 (definitions) → per-Ayat. `Menimbang` + `Mengingat` → single preamble chunk. `PENJELASAN PASAL DEMI PASAL` → per-Pasal explanation chunks linked to main text Pasal.

**Characteristics:**
- Full `BAB` (Chapter) structure, similar to UU (Undang-Undang)
- Often includes `PENJELASAN` (explanatory memorandum) — valuable for RAG, should be chunked and linked to corresponding articles
- More varied quality of PDF formatting (regional scanners)

---

## 4. Putusan-MK

**Full name:** Putusan Mahkamah Konstitusi Republik Indonesia (Constitutional Court Rulings)

**Files:** ~2,394 PDFs

**Naming:** `putusan_mkri_{N}.pdf` where N = sequential ruling number

### Hierarchy

```
PUTUSAN
├── Nomor {N}/PUU-{XX}/YYYY (Case number)
├── DEMI KEADILAN BERDASARKAN KETUHANAN YANG MAHA ESA
├── MAHKAMAH KONSTITUSI REPUBLIK INDONESIA
├── [1.1] Yang mengadili perkara ... (Court introduction)
│   ├── Pemohon (Applicant/Petitioner)
│   ├── Termohon / Pihak Terkait (Respondent / Related Party)
│   └── Wakil / Advokat (Legal counsel)
├── [1.2] Procedural summary (reading, hearing, examining)
├── 2. DUDUK PERKARA (Case Position / Statement of Facts)
│   ├── [2.1] Timeline of filing, registration
│   ├── I. KEWENANGAN MK (Jurisdiction)
│   │   ├── A. DASAR HUKUM (Legal Basis)
│   │   │   ├── 1., 2., 3. ... (cited laws)
│   │   └── B. OBJEK PENGUJIAN (Subject of Review)
│   ├── II. DASAR PERMOHONAN (Basis of Application)
│   │   ├── A. POSITA (Factual Allegations)
│   │   └── B. PETITUM (Prayers / Requests)
├── 3. PERTIMBANGAN HUKUM (Legal Considerations / Reasoning)
│   ├── [3.1] Analysis of jurisdiction/admissibility
│   ├── [3.2] Merits analysis
│   │   ├── a. Constitutionality analysis per article
│   │   ├── b. Comparative law references
│   │   └── c. Precedent (putusan MK sebelumnya)
├── 4. AMAR (Ruling / Order)
│   ├── 1. Mengabulkan/Menolak/Menyatakan... (Grant/Reject/Declare)
│   ├── 2. ...
└── TTD Ketua & Hakim (Signatures of Chief Justice & Justices)
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **Section** | `^\d+\.\s+[A-Z ]+$` | Major section (DUDUK PERKARA, PERTIMBANGAN HUKUM, etc.) |
| **Sub-section** | `^\[\d+\.\d+\]` | Numbered sub-section |
| **Angka** | `^\d+\.\s*` | Numbered list |
| **Huruf** | `^[a-z]\.\s*` | Letter sub-point |

**Chunking unit:** One chunk per major section:
- **Metadata chunk:** Case number, parties, counsel
- **Duduk Perkara chunk:** Full facts/position statement
- **Dasar Hukum + Petita chunk:** Legal basis and prayers
- **Pertimbangan Hukum:** This is the core — split by sub-section `[N.M]` if long, otherwise keep as one chunk
- **Amar chunk:** The actual ruling order — critical for RAG retrieval

**⚠️ Important for RAG:** The `AMAR` (ruling/disposition) is the most检索-valuable section — it states what the court actually decided. The `PERTIMBANGAN HUKUM` (legal reasoning) provides the `why`. Consider metadata-enriching each chunk with case number, parties, and provisions reviewed.

---

## 5. peraturan

**Full name:** Peraturan Pemerintah (Government Regulations)

**Files:** ~224+ PDFs (may contain nested sub-folders: `inpres`, `keppres`)

**Naming:** `PP{N}{YYYY}.pdf` or `PP_NO_{N}_TH_{YYYY}.pdf`

### Hierarchy

```
JUDUL: PERATURAN PEMERINTAH REPUBLIK INDONESIA
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── PRESIDEN REPUBLIK INDONESIA
├── Menimbang
├── Mengingat
│   ├── 1. Undang-Undang ...
├── MEMUTUSKAN
│   └── MENETAPKAN: PERATURAN PEMERINTAH ...
├── BAB I — KETENTUAN UMUM
│   ├── Pasal 1 (Definitions)
│   ├── Pasal 2, 3 ...
├── BAB II, III, ... (Subject chapters)
│   ├── Pasal {N}
│   │   ├── (1), (2), (3) ... (Ayat)
│   │   ├── a), b), c) ... (Huruf)
│   │   └── 1., 2., 3. ... (Angka)
├── BAB {N} — KETENTUAN PERALIHAN (Transitional)
├── BAB {N+1} — KETENTUAN PENUTUP (Closing)
│   ├── Pasal {N}: mulai berlaku pada tanggal ...
└── Agar ...
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **BAB** | `^BAB\s+` | Chapter |
| **Bagian** | `^Bagian\s+` | Part (sub-chapter, in longer PP) |
| **Paragraf** | `^Paragraf\s+` | Paragraph section (in some docs) |
| **Pasal** | `^Pasal\s+\d+` | Article |
| **Ayat** | `^\(\d+\)\s*` | Clause |
| **Huruf** | `^[a-z]\)\s*` | Letter sub-point |
| **Angka** | `^\d+\.\s*` | Numbered list |

**Chunking unit:** Per `Pasal` (all `Ayat`). Pasal 1 definitions → per-Ayat. May have `Bagian` (Part) between `BAB` and `Pasal` — include Part metadata in chunk header.

---

## 6. JDIH_Kemnaker

**Full name:** Jaringan Dokumentasi dan Informasi Hukum — Kementerian Ketenagakerjaan (Ministry of Manpower Legal Database)

**Files:** ~381 PDFs

**Naming:** `Permenaker No. {N} Tahun {YYYY}.pdf` / `PP No. {N} Tahun {YYYY}.pdf`

### Hierarchy

```
PERATURAN MENTERI / PERATURAN PEMERINTAH
├── REPUBLIK INDONESIA
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── DENGAN RAHMAT TUHAN YANG MAHA ESA
├── MENTERI / PRESIDEN REPUBLIK INDONESIA
├── Menimbang
├── Mengingat
│   ├── 1., 2., 3. ...
├── MEMUTUSKAN / MENETAPKAN
├── BAB I — KETENTUAN UMUM
│   ├── Pasal 1 (Definitions — may have 20+ definitions)
│   │   ├── 1. Term A adalah...
│   │   ├── 2. Term B adalah...
│   ├── Pasal 2, 3 ...
├── BAB II — {Subject}
│   ├── BAGIAN KE-1 / BAGIAN PERTAMA
│   │   ├── Paragraf 1
│   │   │   ├── Pasal {N}
│   │   │   │   ├── (1), (2) ...
│   │   │   │   │   ├── a), b) ...
│   ├── BAGIAN KE-2
│   │   ├── Paragraf 1, 2 ...
├── ... (up to 12+ BAB in complex regulations)
├── BAB {N} — KETENTUAN PERALIHAN
├── BAB {N+1} — KETENTUAN PENUTUP
└── TTD Menteri
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **BAB** | `^BAB\s+` | Chapter |
| **Bagian** | `^Bagian\s+` | Part |
| **Paragraf** | `^Paragraf\s+` | Paragraph section |
| **Pasal** | `^Pasal\s+\d+` | Article |
| **Ayat** | `^\(\d+\)\s*` | Clause |
| **Huruf** | `^[a-z]\)\s*` | Letter sub-point |
| **Angka** | `^\d+\.\s*` | Numbered list |

**Chunking unit:** Per `Pasal` with all `Ayat`. Pasal 1 (definitions) → **per-numbered-item** since Kemnaker definitions can be very extensive (20+ definitions in a single article). Include `BAB > Bagian > Paragraf > Pasal` path in chunk metadata.

---

## 7. JDIH_Kemenkeu

**Full name:** JDIH — Kementerian Keuangan (Ministry of Finance)

**Files:** ~231 PDFs

**Naming:** `PMK_No__{N}_Tahun_{YYYY}_{internal_ref}.pdf`

### Hierarchy

```
PERATURAN MENTERI KEUANGAN REPUBLIK INDONESIA
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── Menimbang
├── Mengingat
│   ├── 1., 2., ... (cited regulations)
├── Menetapkan: PERATURAN MENTERI KEUANGAN ...
├── Pasal 1 (Definitions)
│   ├── 1. ..., 2. ...
├── Pasal 2
│   ├── (1), (2) ...
├── Pasal 3, 4, 5 ... (typically compact, 5–15 Pasal)
├── Pasal {N} — Penutup
└── TTD Menteri Keuangan
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **Pasal** | `^Pasal\s+\d+` | Article |
| **Ayat** | `^\(\d+\)\s*` | Clause |
| **Angka** | `^\d+\.\s*` | Numbered list |

**Chunking unit:** Per `Pasal`. No `BAB` in most PMK (flat structure). Often tax/tariff tables embedded — consider extracting tables separately.

**Characteristics:**
- Typically **no BAB/Chapter** divisions — flat Pasal sequence
- Often contains **tariff tables** (HS codes, tax rates) that need special handling for chunking
- More compact than Kemnaker regulations

---

## 8. JDIH_Kemendag

**Full name:** JDIH — Kementerian Perdagangan (Ministry of Trade)

**Files:** ~191 PDFs

**Naming:** `Kepmendag_No__{N}_Tahun_{YYYY}_*.pdf` / `Permendag_{N}_Tahun_{YYYY}.pdf`

### Hierarchy

Two document types in this folder:

**A. Keputusan Menteri (Ministerial Decrees) — e.g., Kepmendag 123/2025**
```
KEPUTUSAN MENTERI PERDAGANGAN
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── DASAR HUKUM
│   ├── 1. UUD 1945 ...
│   ├── 2. ...
├── Pasal 1
├── Pasal 2
│   ├── (1), (2) ...
├── ... (few Pasal, decree-style)
└── TTD Menteri
```

**B. Peraturan Menteri (Ministerial Regulations) — e.g., Permendag**
```
PERATURAN MENTERI PERDAGANGAN
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── Menimbang / Mengingat
├── BAB I — KETENTUAN UMUM
│   ├── Pasal 1 (Definitions)
│   ├── Pasal 2, 3 ...
├── BAB II, III ... (Subject chapters)
├── KETENTUAN PENUTUP
└── TTD Menteri
```

### Structural Elements

Same as standard Kemendag patterns. **Kepmendag** (decrees) are shorter and flatter; **Permendag** (regulations) have full BAB structure.

**Chunking unit:** Per `Pasal`. Decrees are short — can be chunked as single documents if < 10 Pasal.

---

## 9. JDIH_Komdigi

**Full name:** JDIH — Kementerian Komunikasi dan Digital (Ministry of Digital & Communication)

**Files:** ~69 PDFs

**Naming:** `Permenkominfo No. {N} Tahun {YYYY}.pdf`

### Hierarchy

```
PERATURAN MENTERI KOMUNIKASI DAN INFORMATIKA
├── REPUBLIK INDONESIA
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── DENGAN RAHMAT TUHAN YANG MAHA ESA
├── MENTERI ...
├── Menimbang
├── Mengingat
│   ├── 1., 2., ... (extensive legal basis — often 20+ references)
├── MEMUTUSKAN
│   └── MENETAPKAN: PERATURAN MENTERI ...
├── BAB I — KETENTUAN UMUM
│   ├── Pasal 1 (Definitions — extensive, 20+ terms)
│   │   ├── 1. ..., 2. ..., 3. ...
│   ├── Pasal 2, 3 ...
├── BAB II — {Subject} (may have BAGIAN/Paragraf)
│   ├── BAGIAN PERTAMA
│   │   ├── Pasal {N}
│   │   │   ├── (1)
│   │   │   │   ├── a), b), c) ...
│   │   │   │   │   ├── 1., 2. ...
│   ├── BAGIAN KEDUA
├── ... (typically 7+ BAB)
├── BAB {N} — KETENTUAN PERALIHAN
├── BAB {N+1} — KETENTUAN PENUTUP
└── TTD Menteri
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **BAB** | `^BAB\s+` | Chapter |
| **Bagian** | `^Bagian\s+` | Part |
| **Pasal** | `^Pasal\s+\d+` | Article |
| **Ayat** | `^\(\d+\)\s*` | Clause |
| **Huruf** | `^[a-z]\)\s*` | Letter sub-point |
| **Angka** | `^\d+\.\s*` | Numbered list |

**Chunking unit:** Per `Pasal`. Deeply nested (BAB > Bagian > Pasal > Ayat > Huruf > Angka) — preserve full path in chunk metadata. Pasal 1 definitions → per-item.

---

## 10. JDIH_KPU

**Full name:** JDIH — Komisi Pemilihan Umum (Electoral Commission)

**Files:** ~264 PDFs

**Naming:** `PKPU_{N}_{YYYY}.pdf`

### Hierarchy

```
PERATURAN KOMISI PEMILIHAN UMUM
├── REPUBLIK INDONESIA
├── NOMOR {N} TAHUN {YYYY}
├── TENTANG {SUBJECT}
├── Menimbang
├── Mengingat
│   ├── 1. Undang-Undang ...
│   ├── 2. Peraturan KPU sebelumnya ...
├── MEMUTUSKAN
│   └── MENETAPKAN: PERATURAN KPU ...
├── BAB I — KETENTUAN UMUM
│   ├── Pasal 1 (Definitions)
│   ├── Pasal 2, 3 ...
├── BAB II, III ... (Election procedure chapters)
│   ├── Pasal {N}
│   │   ├── (1), (2), (3) ...
│   │   ├── a), b), c) ...
│   │   ├── 1., 2., 3. ...
├── BAB {N} — KETENTUAN PERALIHAN
├── BAB {N+1} — KETENTUAN PENUTUP
└── TTD Ketua KPU
```

### Structural Elements

| Element | Regex Pattern | Description |
|---------|---------------|-------------|
| **BAB** | `^BAB\s+` | Chapter |
| **Pasal** | `^Pasal\s+\d+` | Article |
| **Ayat** | `^\(\d+\)\s*` | Clause |
| **Angka** | `^\d+\.\s*` | Numbered list |

**Chunking unit:** Per `Pasal`. PKPU regulations reference prior PKPU versions extensively in `Mengingat` — keep preamble as one chunk. Election procedures are highly procedural — preserve sequence order in chunk metadata.

---

## Additional Folders (detected in Drive)

### tap_mpr
Ketetapan MPR (MPR Decrees) — ~66 files. Constitutional decrees from the People's Consultative Assembly. Structure similar to UU but shorter. Often contains `SIDANG` (session) references.

### inpres
Instruksi Presiden (Presidential Instructions) — directive orders to specific ministers/officials. **No Pasal structure** — uses numbered instruction points. Flat list format: `1. {directive}`, `2. {directive}`. Chunk as whole document or per-instruction-point.

### keppres
Keputusan Presiden (Presidential Decrees) — ~57 files. Appointment/organizational decrees. Has Pasal but very few (typically 1–7). Short documents — often chunkable as single document.

### pp
Peraturan Pemerintah (Government Regulations) — separate collection from `peraturan` folder. ~137 files. Structure identical to `peraturan` folder.

### perppu
Peraturan Pemerintah Pengganti Undang-Undang (Government Regulations in Lieu of Law) — emergency regulations with UU-level force. Full BAB + Pasal structure. Treat like UU for chunking.

---

## Universal Regex Patterns for Chunking

```python
import re

PATTERNS = {
    "bab":       r"^(BAB|Bab)\s+(?:[IVXLCDM]+|\d+)\b(.*)",
    "bagian":    r"^(BAGIAN|Bagian)\s+(?:KE-|PERTAMA|KEDUA|KETIGA|[IVXLCDM]+|\d+)\b(.*)",
    "paragraf":  r"^(PARAGRAF|Paragraf)\s+(?:[IVXLCDM]+|\d+)\b(.*)",
    "pasal":     r"^Pasal\s+(\d+)(.*)",
    "ayat":      r"^\((\d+)\)\s*(.*)",
    "huruf":     r"^([a-z])\)\s*(.*)",
    "angka":     r"^(\d+)\.\s+(.*)",
}
```

## Recommended Chunking Strategy

| Level | Strategy | Metadata to Include |
|-------|----------|---------------------|
| **Document** | Preamble (`Menimbang` + `Mengingat` + enacting clause) | doc_type, number, year, subject |
| **Chapter** | BAB title + Pasal range | bab_number, bab_title |
| **Article** | **Primary chunk unit** — full Pasal with all Ayat | pasal_number, bab, bagian, paragraf |
| **Clause** | Only split Pasal 1 (definitions) or very long articles | pasal_number, ayat_number |
| **Putusan-MK** | Section-based: Duduk Perkara, Pertimbangan, Amar | case_number, parties, provisions_reviewed |
