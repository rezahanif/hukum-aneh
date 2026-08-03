"""
Family A/B/C parsers for Indonesian legal documents.

Family A (Hierarchical Statute): BAB > Bagian > Paragraf > Pasal > Ayat > huruf/angka
Family B (Decree/Decision): Menimbang > Mengingat > MEMUTUSKAN/MENETAPKAN > Diktum
Family C (Court Ruling): Duduk Perkara > Menimbang > MENGADILI/MENETAPKAN > Amar

All parsers return a list of chunks in the target JSON format.
"""
import re
import json
import sys
import os
from typing import List, Dict, Optional, Tuple

# Add parent dir for imports
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
from scripts.clean_extractor import extract_clean, extract_title_block, detect_doc_type_from_title, fix_glyph_corruption


# ==========================================================================
# SHARED UTILITIES
# ==========================================================================

ROMAN_MAP = {'I':1,'II':2,'III':3,'IV':4,'V':5,'VI':6,'VII':7,'VIII':8,'IX':9,'X':10,
             'XI':11,'XII':12,'XIII':13,'XIV':14,'XV':15,'XVI':16,'XVII':17,'XVIII':18,
             'XIX':19,'XX':20}

def roman_to_int(s):
    return ROMAN_MAP.get(s.upper(), None)


def build_id(doc_type, nomor, year, **parts):
    """Build chunk ID. Handles both hierarchical and flat structures."""
    if doc_type in ("Putusan_MK", "Putusan_MA"):
        # Court rulings use nomor_perkara
        nomor_perkara = parts.get("nomor_perkara", nomor or "?")
        chunk_type = parts.get("chunk_type", "amar")
        chunk_num = parts.get("chunk_num", "?")
        return f"{doc_type}:{nomor_perkara}:{chunk_type}:{chunk_num}"
    elif parts.get("chunk_type") in ("diktum", "pertimbangan"):
        # Family B: flat numbered
        return f"{doc_type}:{nomor or '?'}:{year or '?'}:{parts['chunk_type']}:{parts['chunk_num']}"
    else:
        # Family A: hierarchical
        pasal = parts.get("pasal", "?")
        ayat = parts.get("ayat", "?")
        if ayat and ayat != "?":
            return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}:{ayat}"
        return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}"


def build_parent_id(doc_type, nomor, year, pasal):
    """Build parent ID for ayat chunks."""
    if doc_type in ("Putusan_MK", "Putusan_MA"):
        return None
    return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}"


def build_path(headings):
    """Build path string from heading stack."""
    parts = [h for h in headings if h]
    return " > ".join(parts) if parts else ""


# ==========================================================================
# FAMILY A: HIERARCHICAL STATUTE PARSER
# ==========================================================================

class FamilyAParser:
    """
    Parses hierarchical legal documents: BAB > Bagian > Paragraf > Pasal > Ayat.
    Handles: UU, PP, PerPPU, Perpres, Perda, UUD1945, Permen, PKPU.
    """

    # Heading patterns (order matters: higher levels first)
    RE_BAB = re.compile(r'^\s*BAB\s+([IVXLCDM]+)\b', re.IGNORECASE)
    RE_BAGIAN = re.compile(r'^\s*BAGIAN\s+(\w+)', re.IGNORECASE)
    RE_PARAGRAF = re.compile(r'^\s*PARAGRAF\s+(\w+)', re.IGNORECASE)
    RE_PASAL = re.compile(r'^\s*pasal\s+([IVXLCDM]+|\d+[a-zA-Z]?)\s*\.?\s*.{0,5}$', re.IGNORECASE)
    RE_AYAT = re.compile(r'^\s*\(\s*(\d+[a-zA-Z]?)\s*\)')
    RE_HURUF = re.compile(r'^\s*([a-z])\s*\.\s')
    RE_ANGKA = re.compile(r'^\s*(\d+)\s*\.\s')

    # Quoted amendment detection
    RE_AMENDMENT_TRIGGER = re.compile(
        r'(?:diubah|diganti|dihapus)\s+sehingga\s+berbunyi\s+(?:sebagai\s+berikut|:\s*")',
        re.IGNORECASE
    )
    RE_QUOTE_START = re.compile(r'^\s*["“]')
    RE_QUOTE_END = re.compile(r'["”]\s*[.;]?\s*$')

    # Penutup (closing section) marker
    RE_PENUTUP = re.compile(r'^(?:PENUTUP|penutup)\b')

    def __init__(self, lines, doc_type, nomor, year, issuer):
        self.lines = lines
        self.doc_type = doc_type
        self.nomor = nomor
        self.year = year
        self.issuer = issuer
        self.chunks = []

        # Heading stack
        self.current_bab = None
        self.current_bab_label = None
        self.current_bagian = None
        self.current_bagian_label = None
        self.current_paragraf = None
        self.current_paragraf_label = None
        self.current_pasal = None
        self.current_ayat = None

        # State
        self.in_preamble = True
        self.in_penutup = False
        self.in_quoted_amendment = False
        self.quoted_amendment_note = None
        self.current_ayat_lines = []
        self.current_pasal_lines = []

    def parse(self) -> List[Dict]:
        """Main parse loop. Returns list of chunks."""
        for i, line in enumerate(self.lines):
            text = line["text"].strip()
            if not text:
                continue

            text_upper = text.upper()

            # Detect preamble end / body start
            if self.in_preamble:
                if self.RE_BAB.match(text) or self.RE_PASAL.match(text):
                    self.in_preamble = False
                elif text_upper.startswith("MENETAPKAN") or text_upper.startswith("MEMUTUSKAN"):
                    self.in_preamble = False
                    continue  # skip the decision keyword line itself
                else:
                    continue  # still in preamble

            # Detect penutup
            if self.RE_PENUTUP.match(text):
                self.in_penutup = True
                # Flush any pending ayat
                self._flush_ayat()
                self._flush_pasal()
                continue

            if self.in_penutup:
                continue  # skip penutup

            # Check for quoted amendment
            if not self.in_quoted_amendment:
                if self.RE_AMENDMENT_TRIGGER.search(text):
                    self.in_quoted_amendment = True
                    self.quoted_amendment_note = text[:100]
                    self._flush_ayat()
                    continue
                if self.RE_QUOTE_START.match(text):
                    # Could be inline quote
                    self.in_quoted_amendment = True
                    self.quoted_amendment_note = "Quoted passage (inline)"

            if self.in_quoted_amendment:
                if self.RE_QUOTE_END.search(text):
                    self.in_quoted_amendment = False
                    self.quoted_amendment_note = None
                continue  # skip all quoted text

            # Match structural elements (order: BAB > Bagian > Paragraf > Pasal > Ayat > huruf > angka)
            m = self.RE_BAB.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_bab = m.group(1)
                # Get label text after BAB X
                label_match = re.match(r'^\s*BAB\s+[IVXLCDM]+\s+(.*)', text, re.IGNORECASE)
                self.current_bab_label = label_match.group(1).strip().rstrip('.') if label_match else f"BAB {m.group(1)}"
                self.current_bagian = None
                self.current_bagian_label = None
                self.current_paragraf = None
                self.current_paragraf_label = None
                continue

            m = self.RE_BAGIAN.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_bagian = m.group(1)
                label_match = re.match(r'^\s*BAGIAN\s+\w+\s+(.*)', text, re.IGNORECASE)
                self.current_bagian_label = label_match.group(1).strip().rstrip('.') if label_match else f"Bagian {m.group(1)}"
                self.current_paragraf = None
                self.current_paragraf_label = None
                continue

            m = self.RE_PARAGRAF.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_paragraf = m.group(1)
                label_match = re.match(r'^\s*PARAGRAF\s+\w+\s+(.*)', text, re.IGNORECASE)
                self.current_paragraf_label = label_match.group(1).strip().rstrip('.') if label_match else f"Paragraf {m.group(1)}"
                continue

            m = self.RE_PASAL.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_pasal = m.group(1)
                self.current_ayat = None
                continue

            m = self.RE_AYAT.match(text)
            if m and self.current_pasal:
                self._flush_ayat()
                self.current_ayat = m.group(1)
                # The rest of the line after (N) is the start of ayat text
                ayat_text = self.RE_AYAT.sub('', text).strip()
                if ayat_text:
                    self.current_ayat_lines.append(ayat_text)
                continue

            # If we're inside a pasal but no ayat detected, accumulate lines
            if self.current_pasal:
                if self.current_ayat:
                    self.current_ayat_lines.append(text)
                else:
                    self.current_pasal_lines.append(text)

        # Flush remaining
        self._flush_ayat()
        self._flush_pasal()

        return self.chunks

    def _flush_ayat(self):
        """Flush accumulated ayat lines as a chunk."""
        if not self.current_ayat or not self.current_ayat_lines:
            self.current_ayat_lines = []
            return

        text = ' '.join(self.current_ayat_lines).strip()
        if not text:
            self.current_ayat_lines = []
            return

        headings = [
            f"BAB {self.current_bab}" if self.current_bab else None,
            self.current_bab_label,
            f"Bagian {self.current_bagian}" if self.current_bagian else None,
            self.current_bagian_label,
            f"Paragraf {self.current_paragraf}" if self.current_paragraf else None,
            self.current_paragraf_label,
            f"Pasal {self.current_pasal}" if self.current_pasal else None,
            f"Ayat ({self.current_ayat})",
        ]
        clean_headings = [h for h in headings if h]

        path = build_path(clean_headings)
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            pasal=self.current_pasal, ayat=self.current_ayat)
        parent_id = build_parent_id(self.doc_type, self.nomor, self.year, self.current_pasal)

        self.chunks.append({
            "id": chunk_id,
            "text": text,
            "metadata": {
                "doc_type": self.doc_type,
                "issuer": self.issuer,
                "year": self.year,
                "bab": self.current_bab or None,
                "pasal": self.current_pasal,
                "ayat": self.current_ayat,
                "parent": parent_id,
                "path": path,
                "status": "active",
            }
        })
        self.current_ayat_lines = []

    def _flush_pasal(self):
        """Flush pasal-only content (no ayat detected) as a single chunk."""
        if not self.current_pasal or not self.current_pasal_lines:
            self.current_pasal_lines = []
            return

        text = ' '.join(self.current_pasal_lines).strip()
        if not text:
            self.current_pasal_lines = []
            return

        headings = [
            f"BAB {self.current_bab}" if self.current_bab else None,
            self.current_bab_label,
            f"Bagian {self.current_bagian}" if self.current_bagian else None,
            self.current_bagian_label,
            f"Paragraf {self.current_paragraf}" if self.current_paragraf else None,
            self.current_paragraf_label,
            f"Pasal {self.current_pasal}",
        ]
        clean_headings = [h for h in headings if h]
        path = build_path(clean_headings)
        chunk_id = build_id(self.doc_type, self.nomor, self.year, pasal=self.current_pasal)

        self.chunks.append({
            "id": chunk_id,
            "text": text,
            "metadata": {
                "doc_type": self.doc_type,
                "issuer": self.issuer,
                "year": self.year,
                "bab": self.current_bab or None,
                "pasal": self.current_pasal,
                "ayat": None,
                "parent": None,
                "path": path,
                "status": "active",
            }
        })
        self.current_pasal_lines = []


# ==========================================================================
# FAMILY B: DECREE/DECISION PARSER
# ==========================================================================

class FamilyBParser:
    """
    Parses decree/decision documents: Menimbang > Mengingat > MEMUTUSKAN/MENETAPKAN > Diktum.
    Handles: Keppres, Inpres, Tap_MPR, Kepmen.
    """

    RE_MENIMBANG = re.compile(r'^\s*menimbang\s*:?\s*$', re.IGNORECASE)
    RE_MENGINGAT = re.compile(r'^\s*mengingat\s*:?\s*$', re.IGNORECASE)
    # Decision keywords: match keyword at/near end of line, allowing a short prefix
    # like "dengan ini" before the keyword. This handles both standalone "MEMUTUSKAN:"
    # and inline "dengan ini menginstruksikan:" patterns.
    RE_MEMUTUSKAN = re.compile(r'^(?:.{0,30}\s+)?memutuskan\s*:?\s*$', re.IGNORECASE)
    RE_MENETAPKAN = re.compile(r'^(?:.{0,30}\s+)?menetapkan\s*:?\s*$', re.IGNORECASE)
    RE_MENINSTRUKSIKAN = re.compile(r'^(?:.{0,30}\s+)?menginstruksikan\s*:?\s*$', re.IGNORECASE)
    # Item markers: trailing space is optional (handles bare "1." or "a.")
    RE_ITEM_LETTER = re.compile(r'^\s*([a-z])\s*\.\s*')
    RE_ITEM_NUMBER = re.compile(r'^\s*(\d+)\s*\.\s*')
    # Indonesian ordinal words used as diktum markers (KESATU, KEDUA, KETIGA, ...)
    RE_ORDINAL = re.compile(
        r'^\s*(KE(?:SATU|DUA|TIGA|EMPAT|LIMA|NAM|TUJUH|DELAPAN|SEMBILAN|SEPULUH|'
        r'SEBELAS|DUA BELAS|TIGA BELAS|EMPAT BELAS|LIMA BELAS|ENAM BELAS|'
        r'TUJUH BELAS|DELAPAN BELAS|SEMBILAN BELAS|DUA PULUH))\b',
        re.IGNORECASE
    )

    def __init__(self, lines, doc_type, nomor, year, issuer):
        self.lines = lines
        self.doc_type = doc_type
        self.nomor = nomor
        self.year = year
        self.issuer = issuer
        self.chunks = []

        self.section = "pre"  # pre, menimbang, mengingat, keputusan, done
        self.current_item = None
        self.current_item_lines = []
        self.preamble_lines = []
        self.decision_keyword = None  # MEMUTUSKAN, MENETAPKAN, etc.

    def parse(self) -> List[Dict]:
        for line in self.lines:
            text = line["text"].strip()
            if not text:
                continue

            text_upper = text.upper()

            # Section transitions
            if self.RE_MENIMBANG.match(text):
                self._flush_preamble()
                self.section = "menimbang"
                self.current_item = None
                continue

            if self.RE_MENGINGAT.match(text):
                self._flush_preamble()
                self.section = "mengingat"
                self.current_item = None
                continue

            # Decision keyword triggers (regex allows short prefix like "dengan ini").
            # Guard: only trigger if not already in keputusan, to avoid
            # re-triggering on "menetapkan" appearing inside diktum item text.
            if self.section not in ("keputusan", "done"):
                if self.RE_MENINSTRUKSIKAN.match(text):
                    self._flush_item()
                    self.section = "keputusan"
                    self.decision_keyword = "MENGINSTRUKSIKAN"
                    self.current_item = None
                    continue

                if self.RE_MEMUTUSKAN.match(text):
                    self._flush_item()
                    self.section = "keputusan"
                    self.decision_keyword = "MEMUTUSKAN"
                    self.current_item = None
                    continue

                if self.RE_MENETAPKAN.match(text):
                    self._flush_item()
                    self.section = "keputusan"
                    self.decision_keyword = "MENETAPKAN"
                    self.current_item = None
                    continue

            if self.section == "pre":
                continue  # skip title block

            if self.section == "done":
                continue

            # Check for signature block (end of substantive content)
            # Must come AFTER pre/done guards to avoid false triggers in preamble.
            text_lower = text.lower()
            if "ttd" in text_lower and len(text) < 20:
                self._flush_item()
                self.section = "done"
                continue

            # Check for new item in keputusan section
            if self.section == "keputusan":
                m_letter = self.RE_ITEM_LETTER.match(text)
                m_number = self.RE_ITEM_NUMBER.match(text)
                m_ordinal = self.RE_ORDINAL.match(text)

                if m_letter or m_number or m_ordinal:
                    self._flush_item()
                    if m_ordinal:
                        self.current_item = m_ordinal.group(1).upper()
                    elif m_letter:
                        self.current_item = m_letter.group(1)
                    else:
                        self.current_item = m_number.group(1)
                    # Get text after the item marker
                    m = m_letter or m_number or m_ordinal
                    remaining = text[m.end():].strip()
                    if remaining:
                        self.current_item_lines.append(remaining)
                    continue

                # Sub-items: continue accumulating under current item
                if self.current_item:
                    self.current_item_lines.append(text)
                continue

            # menimbang/mengingat: accumulate as preamble (skip for chunking or tag separately)
            # Per the plan, optionally chunk these with section tag

        self._flush_item()
        return self.chunks

    def _flush_item(self):
        if not self.current_item or not self.current_item_lines:
            self.current_item_lines = []
            return

        text = ' '.join(self.current_item_lines).strip()
        if not text:
            self.current_item_lines = []
            return

        kw = self.decision_keyword or "MEMUTUSKAN"
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            chunk_type="diktum", chunk_num=self.current_item)

        self.chunks.append({
            "id": chunk_id,
            "text": text,
            "metadata": {
                "doc_type": self.doc_type,
                "issuer": self.issuer,
                "year": self.year,
                "nomor": self.nomor,
                "section": "keputusan",
                "decision_keyword": kw,
                "diktum": self.current_item,
                "parent": None,
                "path": f"{kw} > Diktum {self.current_item}",
                "status": "active",
            }
        })
        self.current_item_lines = []

    def _flush_preamble(self):
        self.preamble_lines = []  # Could optionally save


# ==========================================================================
# FAMILY C: COURT RULING PARSER
# ==========================================================================

class FamilyCParser:
    """
    Parses court rulings: Duduk Perkara > Menimbang > MENGADILI/MENETAPKAN > Amar.
    Handles: Putusan MK, Putusan MA.
    """

    RE_MENGADILI = re.compile(r'^\s*mengadili\s*:?\s*$', re.IGNORECASE)
    RE_MENETAPKAN = re.compile(r'^\s*menetapkan\s*:?\s*$', re.IGNORECASE)
    RE_AMAR = re.compile(r'^\s*(\d+)\s*\.\s+')
    RE_PUTUSAN = re.compile(r'^\s*putusan\s*:?\s*$', re.IGNORECASE)

    def __init__(self, lines, doc_type, nomor, year, issuer, nomor_perkara=None):
        self.lines = lines
        self.doc_type = doc_type
        self.nomor = nomor
        self.year = year
        self.issuer = issuer
        self.nomor_perkara = nomor_perkara
        self.chunks = []

        self.section = "pre"  # pre, menimbang, mengadili, amar, done
        self.current_amar = None
        self.current_amar_lines = []
        self.pertimbangan_lines = []

    def parse(self) -> List[Dict]:
        for line in self.lines:
            text = line["text"].strip()
            if not text:
                continue

            text_upper = text.upper()

            if self.RE_MENGADILI.match(text):
                self._flush_amar()
                self.section = "amar"
                self.decision_keyword = "MENGADILI"
                continue

            if self.RE_MENETAPKAN.match(text):
                self._flush_amar()
                self.section = "amar"
                self.decision_keyword = "MENETAPKAN"
                continue

            if self.section == "pre":
                continue

            if self.section == "done":
                continue

            if self.section == "amar":
                m = self.RE_AMAR.match(text)
                if m:
                    self._flush_amar()
                    self.current_amar = m.group(1)
                    remaining = text[m.end():].strip()
                    if remaining:
                        self.current_amar_lines.append(remaining)
                    continue

                if self.current_amar:
                    self.current_amar_lines.append(text)

                # Detect end: signature block
                if re.search(r'\bttd\b|\bHAKIM\b|\bKETUA\b', text_upper) and len(text) < 30:
                    self._flush_amar()
                    self.section = "done"

        self._flush_amar()
        return self.chunks

    def _flush_amar(self):
        if not self.current_amar or not self.current_amar_lines:
            self.current_amar_lines = []
            return

        text = ' '.join(self.current_amar_lines).strip()
        if not text:
            self.current_amar_lines = []
            return

        kw = getattr(self, 'decision_keyword', 'MENGADILI')
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            nomor_perkara=self.nomor_perkara,
                            chunk_type="amar", chunk_num=self.current_amar)

        self.chunks.append({
            "id": chunk_id,
            "text": text,
            "metadata": {
                "doc_type": self.doc_type,
                "issuer": self.issuer,
                "year": self.year,
                "nomor_perkara": self.nomor_perkara,
                "section": "amar",
                "decision_keyword": kw,
                "amar": self.current_amar,
                "parent": None,
                "path": f"{kw} > Amar {self.current_amar}",
                "status": "active",
            }
        })
        self.current_amar_lines = []


# ==========================================================================
# UNIFIED CHUNKER
# ==========================================================================

def extract_nomor_perkara(lines):
    """Extract case number (e.g. 96/PUU-XVI/2018) from court ruling lines."""
    for line in lines[:30]:
        m = re.search(r'(\d+/[A-Z]+-[A-Z]+/\d{4})', line["text"])
        if m:
            return m.group(1)
    return None


def chunk_pdf(pdf_path: str, directory: str, family_map: Dict) -> List[Dict]:
    """
    Main entry point. Extract + parse a PDF into chunks.
    Returns list of chunk dicts with id, text, metadata.
    """
    # Get family config
    dir_config = family_map.get(directory, {})
    family = dir_config.get("family", "A")
    default_doc_type = dir_config.get("doc_type", "Unknown")
    default_issuer = dir_config.get("issuer", "Unknown")

    # Extract clean text
    try:
        lines, meta = extract_clean(pdf_path)
    except Exception as e:
        return [{"id": f"ERROR:{directory}", "text": str(e), "metadata": {"error": True}}]

    if not lines:
        return [{"id": f"EMPTY:{directory}", "text": "No extractable text (possibly scanned PDF)", "metadata": {"error": True, "needs_ocr": True}}]

    # Detect doc type from title block (override directory default)
    title_block = extract_title_block(lines)
    detected_type, detected_issuer = detect_doc_type_from_title(title_block)
    doc_type = detected_type if detected_type != "Unknown" else default_doc_type
    issuer = detected_issuer if detected_issuer != "Unknown" else default_issuer
    nomor = meta.get("nomor")
    year = meta.get("year")

    # Parse based on family
    if family == "A":
        parser = FamilyAParser(lines, doc_type, nomor, year, issuer)
        chunks = parser.parse()
    elif family == "B":
        parser = FamilyBParser(lines, doc_type, nomor, year, issuer)
        chunks = parser.parse()
    elif family == "C":
        nomor_perkara = extract_nomor_perkara(lines)
        parser = FamilyCParser(lines, doc_type, nomor, year, issuer, nomor_perkara)
        chunks = parser.parse()
    else:
        chunks = [{"id": f"UNKNOWN_FAMILY:{directory}", "text": "", "metadata": {"error": True}}]

    # Add extraction metadata to all chunks
    for chunk in chunks:
        chunk["metadata"]["source_file"] = meta["file"]
        chunk["metadata"]["directory"] = directory
        chunk["metadata"]["family"] = family
        if nomor:
            chunk["metadata"]["nomor"] = nomor
        if year:
            chunk["metadata"]["year"] = year

    return chunks


if __name__ == "__main__":
    import json as json_lib

    CONFIG_PATH = "/home/z/my-project/config/directory_family_map.json"
    SAMPLES_DIR = "/home/z/my-project/download/samples"

    SAMPLES = {
        "uu": "uu/uunomor41tahun2014.pdf",
        "pp": "pp/PP_NO_70_TH_1991.pdf",
        "perppu": "perppu/perppu-no-148-tahun-2024.pdf",
        "perpres": "perpres/perpres-no-127-tahun-2024.pdf",
        "perda": "perda/perda-kabupaten-sukoharjo-no-1-tahun-2025.pdf",
        "keppres": "keppres/keppres-no-5-tahun-2015_Dewan Kawasan Kawasan Ekonomi Khusus Provinsi Kalimantan Timur.pdf",
        "inpres": "inpres/inpres-no-3-tahun-2023_Percepatan Peningkatan Konektivitas Jalan Daerah.pdf",
        "tap_mpr": "tap_mpr/Ketetapan Majelis Permusyawaratan Rakyat Republik Indonesia Nomor IIIMPR2002 ten.pdf",
        "uud-1945": "uud-1945/uud_1945.pdf",
        "Putusan-MK": "Putusan-MK/putusan_mkri_5301.pdf",
        "JDIH_Kemnaker": "JDIH_Kemnaker/Permenaker No. 90 Tahun 2013.pdf",
        "JDIH_Kemenkeu": "JDIH_Kemenkeu/PMK_No__9_Tahun_2025_2024pmkeuangan009.pdf",
        "JDIH_Kemendag": "JDIH_Kemendag/Kepmendag_No__123_Tahun_2025_download_3142_2.pdf",
        "JDIH_Komdigi": "JDIH_Komdigi/Permenkominfo No. 5 Tahun 2024.pdf",
        "JDIH_KPU": "JDIH_KPU/PKPU_8_2026.pdf",
        "peraturan": "peraturan/PP0201962.pdf",
    }

    with open(CONFIG_PATH) as f:
        family_map = json_lib.load(f)

    all_results = {}
    for folder, rel_path in SAMPLES.items():
        full_path = os.path.join(SAMPLES_DIR, rel_path)
        if not os.path.exists(full_path):
            print(f"SKIP {folder}: file not found")
            continue

        print(f"\n{'='*60}")
        print(f"{folder} (Family {family_map.get(folder, {}).get('family', '?')})")
        print(f"{'='*60}")

        chunks = chunk_pdf(full_path, folder, family_map)
        all_results[folder] = chunks

        print(f"Chunks produced: {len(chunks)}")
        for c in chunks[:5]:
            txt_preview = c['text'][:80] + "..." if len(c['text']) > 80 else c['text']
            print(f"  [{c['id']}] {txt_preview}")
        if len(chunks) > 5:
            print(f"  ... and {len(chunks)-5} more")

    # Save results
    output_path = "/home/z/my-project/download/chunk_results_v2.json"
    with open(output_path, "w", encoding="utf-8") as f:
        json_lib.dump(all_results, f, ensure_ascii=False, indent=2)
    print(f"\nAll results saved to: {output_path}")
