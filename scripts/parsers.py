r"""Family A/B/C parsers for Indonesian legal documents.

v3 fixes (based on QA assessment round 3):
  - Fix #2: doc_type parsed from title block first, directory config as fallback only.
    Detected mismatches are flagged in metadata ("doc_type_override": true).
  - Fix #3: Section headings (BAB/Bagian/Paragraf titles) update path breadcrumb only,
    never emitted as standalone chunks. Eliminates the heading-as-chunk id collision.
  - Fix #4: PENJELASAN section detected and chunks get ":penjelasan" id suffix.
    Signature/promulgation block is also a hard boundary.
  - Fix #5: nomor/year extraction restricted to title block (first ~10 lines),
    never from body citations.
  - Fix #6: Glyph/noise filter expanded with garbled header patterns beyond
    "SK No" to catch variants like "FTRESIDEN R Ei:IUE...".
  - Fix #7: Inpres diktum numbering uses compound key (parent_ordinal:item_num)
    so items from different KESATU/KEDUA sections don't collide.
  - Fix #1: tap_mpr crash (NameError) fixed in clean_extractor.py.

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


# Minimum text length for a chunk to be emitted (prevents "." or "huruf c." chunks)
MIN_CHUNK_TEXT_LENGTH = 10


def build_id(doc_type, nomor, year, **parts):
    """Build chunk ID. Handles both hierarchical and flat structures."""
    if doc_type in ("Putusan_MK", "Putusan_MA"):
        nomor_perkara = parts.get("nomor_perkara", nomor or "?")
        chunk_type = parts.get("chunk_type", "amar")
        chunk_num = parts.get("chunk_num", "?")
        return f"{doc_type}:{nomor_perkara}:{chunk_type}:{chunk_num}"
    elif parts.get("chunk_type") in ("diktum", "pertimbangan"):
        parent_ordinal = parts.get("parent_ordinal", "")
        if parent_ordinal:
            return f"{doc_type}:{nomor or '?'}:{year or '?'}:{parts['chunk_type']}:{parent_ordinal}:{parts['chunk_num']}"
        return f"{doc_type}:{nomor or '?'}:{year or '?'}:{parts['chunk_type']}:{parts['chunk_num']}"
    else:
        pasal = parts.get("pasal", "?")
        ayat = parts.get("ayat", "?")
        section_suffix = parts.get("section_suffix", "")  # e.g. ":penjelasan"
        if section_suffix:
            if ayat and ayat != "?":
                return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}:{ayat}{section_suffix}"
            return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}{section_suffix}"
        if ayat and ayat != "?":
            return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}:{ayat}"
        return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}"


def build_parent_id(doc_type, nomor, year, pasal):
    """Build parent ID for ayat chunks."""
    if doc_type in ("Putusan_MK", "Putusan_MA"):
        return None
    return f"{doc_type}:{nomor or '?'}:{year or '?'}:{pasal}"


def build_path(headings):
    """Build path string from heading stack. Filters out None and duplicates."""
    seen = set()
    parts = []
    for h in headings:
        if h and h not in seen:
            seen.add(h)
            parts.append(h)
    return " > ".join(parts) if parts else ""


def _make_bab_headings(bab_num, bab_label, bagian_num, bagian_label,
                       paragraf_num, paragraf_label, pasal_str=None, ayat_str=None):
    """Build heading list for path, avoiding duplicate BAB labels.
    
    Fix for bug #2: when BAB has no descriptive label (e.g., just "BAB I"),
    the label defaults to "BAB I" which duplicates the number element.
    Now: if label equals the number element, omit it.
    """
    headings = []
    
    bab_num_str = f"BAB {bab_num}" if bab_num else None
    if bab_num_str:
        headings.append(bab_num_str)
    # Only add label if it's different from the number element and non-empty
    if bab_label and bab_label != bab_num_str:
        headings.append(bab_label)
    
    bagian_num_str = f"Bagian {bagian_num}" if bagian_num else None
    if bagian_num_str:
        headings.append(bagian_num_str)
    if bagian_label and bagian_label != bagian_num_str:
        headings.append(bagian_label)
    
    paragraf_num_str = f"Paragraf {paragraf_num}" if paragraf_num else None
    if paragraf_num_str:
        headings.append(paragraf_num_str)
    if paragraf_label and paragraf_label != paragraf_num_str:
        headings.append(paragraf_label)
    
    if pasal_str:
        headings.append(pasal_str)
    if ayat_str:
        headings.append(ayat_str)
    
    return headings


# ==========================================================================
# FAMILY A: HIERARCHICAL STATUTE PARSER
# ==========================================================================

class FamilyAParser:
    """
    Parses hierarchical legal documents: BAB > Bagian > Paragraf > Pasal > Ayat.
    Handles: UU, PP, PerPPU, Perpres, Perda, UUD1945, Permen, PKPU.
    
    v3 changes:
    - Headings (BAB/Bagian/Paragraf) update path only, never emit as chunks.
    - PENJELASAN section detected: chunks get ":penjelasan" suffix.
    - Signature/promulgation block is a hard boundary (flushes pasal without emitting).
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
    RE_QUOTE_START = re.compile(r'^\s*["\u201c]')
    RE_QUOTE_END = re.compile(r'["\u201d]\s*[.;]?\s*$')

    # Penjelasan (elucidation) section marker
    RE_PENJELASAN = re.compile(
        r'^\s*(?:PENJELASAN|Penjelasan)\s+(?:ATAS\s+)?(?:PERATURAN|Undang-Undang|Peraturan)',
        re.IGNORECASE
    )
    # Sub-sections within Penjelasan: "I. UMUM", "II. PASAL DEMI PASAL", etc.
    RE_PENJELASAN_SUB = re.compile(r'^\s*([IVXLCDM]+)\s+\.\s+(.+)', re.IGNORECASE)
    
    # Signature/promulgation block markers (hard boundaries)
    RE_SIGNATURE = re.compile(
        r'^(?:Ditetapkan\s+di|Di\s+tetapkan\s+di|Agar\s+setiap\s+orang\s+mengetahui|'
        r'TD\s*T|tt\s*d|\bTTD\b)',
        re.IGNORECASE
    )
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
        self.in_penjelasan = False
        self.penjelasan_sub = None  # e.g. "UMUM", "PASAL DEMI PASAL"
        self.in_quoted_amendment = False
        self.quoted_amendment_source = None
        self.current_ayat_lines = []
        self.current_pasal_lines = []
        self.quoted_amendment_lines = []

    def parse(self) -> List[Dict]:
        """Main parse loop. Returns list of chunks."""
        for i, line in enumerate(self.lines):
            text = line["text"].strip()
            if not text:
                continue

            text_upper = text.upper()

            # --- Penjelasan detection (can appear after penutup or signature) ---
            if self.RE_PENJELASAN.match(text):
                self._flush_ayat()
                self._flush_pasal()
                self._flush_quoted_amendment()
                self.in_penjelasan = True
                self.in_penutup = False
                self.current_pasal = None
                self.current_ayat = None
                continue

            # Inside Penjelasan: check for sub-sections ("I. UMUM", "II. PASAL DEMI PASAL")
            if self.in_penjelasan:
                m_sub = self.RE_PENJELASAN_SUB.match(text)
                if m_sub:
                    self._flush_ayat()
                    self._flush_pasal()
                    self.penjelasan_sub = m_sub.group(2).strip().upper()
                    self.current_pasal = None
                    self.current_ayat = None
                    continue
                
                # Check for pasal references within "PASAL DEMI PASAL" section
                if self.penjelasan_sub and "PASAL" in self.penjelasan_sub:
                    m_p = re.match(r'^\s*Pasal\s+(\d+[a-zA-Z]?)\s*$', text, re.IGNORECASE)
                    if m_p:
                        self._flush_ayat()
                        self._flush_pasal()
                        self.current_pasal = m_p.group(1)
                        self.current_ayat = None
                        continue
                    
                    m_a = self.RE_AYAT.match(text)
                    if m_a and self.current_pasal:
                        self._flush_ayat()
                        self.current_ayat = m_a.group(1)
                        ayat_text = self.RE_AYAT.sub('', text).strip()
                        if ayat_text:
                            self.current_ayat_lines.append(ayat_text)
                        continue
                
                # Accumulate penjelasan content
                if self.current_pasal:
                    if self.current_ayat:
                        self.current_ayat_lines.append(text)
                    else:
                        self.current_pasal_lines.append(text)
                # Non-pasal penjelasan text (e.g. UMUM section prose) — skip as unchunked context
                continue

            # --- Detect preamble end / body start ---
            if self.in_preamble:
                if self.RE_BAB.match(text) or self.RE_PASAL.match(text):
                    self.in_preamble = False
                elif text_upper.startswith("MENETAPKAN") or text_upper.startswith("MEMUTUSKAN"):
                    self.in_preamble = False
                    continue
                else:
                    continue

            # Detect penutup
            if self.RE_PENUTUP.match(text):
                self._flush_ayat()
                self._flush_pasal()
                self._flush_quoted_amendment()
                self.in_penutup = True
                continue

            if self.in_penutup:
                # Signature/promulgation block: hard boundary, flush but don't emit
                self._flush_pasal()
                continue

            # Signature block detection (outside penutup too)
            if self.RE_SIGNATURE.match(text):
                self._flush_ayat()
                self._flush_pasal()
                self._flush_quoted_amendment()
                self.in_penutup = True  # treat same as penutup
                continue

            # Check for quoted amendment TRIGGER ("Pasal X diubah sehingga berbunyi:")
            if not self.in_quoted_amendment:
                if self.RE_AMENDMENT_TRIGGER.search(text):
                    m_pasal = re.search(r'Pasal\s+(\d+|\w+)\s+(?:ayat\s*\(\d+\)\s+)?(?:diubah|diganti|dihapus)',
                                         text, re.IGNORECASE)
                    source_pasal = m_pasal.group(1) if m_pasal else "?"
                    self.in_quoted_amendment = True
                    self.quoted_amendment_source = source_pasal
                    self._flush_ayat()
                    self.quoted_amendment_lines = [text]
                    continue

            # Inside quoted amendment: accumulate text
            if self.in_quoted_amendment:
                self.quoted_amendment_lines.append(text)
                if self.RE_QUOTE_END.search(text):
                    self._flush_quoted_amendment()
                if (self.RE_BAB.match(text) or self.RE_PASAL.match(text) or
                    self.RE_PENUTUP.match(text)):
                    self._flush_quoted_amendment()
                    self.in_quoted_amendment = False
                elif self.RE_QUOTE_END.search(text):
                    pass
                else:
                    continue

            # === Match structural elements ===
            # FIX #3: Headings update path only, NEVER emit as chunks.

            m = self.RE_BAB.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_bab = m.group(1)
                label_match = re.match(r'^\s*BAB\s+[IVXLCDM]+\s+(.*)', text, re.IGNORECASE)
                raw_label = label_match.group(1).strip().rstrip('.') if label_match else ""
                if raw_label and raw_label.upper() != f"BAB {m.group(1).upper()}":
                    self.current_bab_label = raw_label
                else:
                    self.current_bab_label = None
                self.current_bagian = None
                self.current_bagian_label = None
                self.current_paragraf = None
                self.current_paragraf_label = None
                continue  # heading updates path, not a chunk

            m = self.RE_BAGIAN.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_bagian = m.group(1)
                label_match = re.match(r'^\s*BAGIAN\s+\w+\s+(.*)', text, re.IGNORECASE)
                raw_label = label_match.group(1).strip().rstrip('.') if label_match else ""
                if raw_label:
                    self.current_bagian_label = raw_label
                else:
                    self.current_bagian_label = None
                self.current_paragraf = None
                self.current_paragraf_label = None
                continue  # heading updates path, not a chunk

            m = self.RE_PARAGRAF.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_paragraf = m.group(1)
                label_match = re.match(r'^\s*PARAGRAF\s+\w+\s+(.*)', text, re.IGNORECASE)
                raw_label = label_match.group(1).strip().rstrip('.') if label_match else ""
                if raw_label:
                    self.current_paragraf_label = raw_label
                else:
                    self.current_paragraf_label = None
                continue  # heading updates path, not a chunk

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
        self._flush_quoted_amendment()

        return self.chunks

    def _flush_ayat(self):
        """Flush accumulated ayat lines as a chunk."""
        if not self.current_ayat or not self.current_ayat_lines:
            self.current_ayat_lines = []
            return

        text = ' '.join(self.current_ayat_lines).strip()
        if not text or len(text) < MIN_CHUNK_TEXT_LENGTH:
            self.current_ayat_lines = []
            return

        # Determine section suffix for Penjelasan
        section_suffix = ":penjelasan" if self.in_penjelasan else ""
        
        # Build path: include penjelasan sub-section if present
        if self.in_penjelasan and self.penjelasan_sub:
            path_prefix = f"PENJELASAN > {self.penjelasan_sub}"
        elif self.in_penjelasan:
            path_prefix = "PENJELASAN"
        else:
            path_prefix = None

        headings = _make_bab_headings(
            self.current_bab, self.current_bab_label,
            self.current_bagian, self.current_bagian_label,
            self.current_paragraf, self.current_paragraf_label,
            pasal_str=f"Pasal {self.current_pasal}",
            ayat_str=f"Ayat ({self.current_ayat})"
        )
        if path_prefix:
            headings = [path_prefix] + headings

        path = build_path(headings)
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            pasal=self.current_pasal, ayat=self.current_ayat,
                            section_suffix=section_suffix)
        parent_id = build_parent_id(self.doc_type, self.nomor, self.year, self.current_pasal)

        metadata = {
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
        if self.in_penjelasan:
            metadata["section"] = "penjelasan"
            if self.penjelasan_sub:
                metadata["penjelasan_sub"] = self.penjelasan_sub

        self.chunks.append({"id": chunk_id, "text": text, "metadata": metadata})
        self.current_ayat_lines = []

    def _flush_pasal(self):
        """Flush pasal-only content (no ayat detected) as a single chunk."""
        if not self.current_pasal or not self.current_pasal_lines:
            self.current_pasal_lines = []
            return

        text = ' '.join(self.current_pasal_lines).strip()
        if not text or len(text) < MIN_CHUNK_TEXT_LENGTH:
            self.current_pasal_lines = []
            return

        section_suffix = ":penjelasan" if self.in_penjelasan else ""
        
        if self.in_penjelasan and self.penjelasan_sub:
            path_prefix = f"PENJELASAN > {self.penjelasan_sub}"
        elif self.in_penjelasan:
            path_prefix = "PENJELASAN"
        else:
            path_prefix = None

        headings = _make_bab_headings(
            self.current_bab, self.current_bab_label,
            self.current_bagian, self.current_bagian_label,
            self.current_paragraf, self.current_paragraf_label,
            pasal_str=f"Pasal {self.current_pasal}",
        )
        if path_prefix:
            headings = [path_prefix] + headings
        path = build_path(headings)
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            pasal=self.current_pasal, section_suffix=section_suffix)

        metadata = {
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
        if self.in_penjelasan:
            metadata["section"] = "penjelasan"
            if self.penjelasan_sub:
                metadata["penjelasan_sub"] = self.penjelasan_sub

        self.chunks.append({"id": chunk_id, "text": text, "metadata": metadata})
        self.current_pasal_lines = []

    def _flush_quoted_amendment(self):
        """Flush quoted amendment text as a chunk with status='quoted_amendment'."""
        if not self.quoted_amendment_lines:
            self.in_quoted_amendment = False
            self.quoted_amendment_source = None
            return

        text = ' '.join(self.quoted_amendment_lines).strip()
        if not text or len(text) < MIN_CHUNK_TEXT_LENGTH:
            self.quoted_amendment_lines = []
            self.in_quoted_amendment = False
            self.quoted_amendment_source = None
            return

        source = self.quoted_amendment_source or "?"
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            pasal=source, ayat="quoted")

        self.chunks.append({
            "id": chunk_id,
            "text": text,
            "metadata": {
                "doc_type": self.doc_type,
                "issuer": self.issuer,
                "year": self.year,
                "bab": self.current_bab or None,
                "pasal": source,
                "ayat": "quoted",
                "parent": None,
                "path": f"[AMANDEMEN Pasal {source}]",
                "status": "quoted_amendment",
                "amends_pasal": source,
            }
        })
        self.quoted_amendment_lines = []
        self.in_quoted_amendment = False
        self.quoted_amendment_source = None


# ==========================================================================
# FAMILY B: DECIDE/DECISION PARSER
# ==========================================================================

class FamilyBParser:
    """
    Parses decree/decision documents: Menimbang > Mengingat > MEMUTUSKAN/MENETAPKAN > Diktum.
    Handles: Keppres, Inpres, Tap_MPR, Kepmen.
    
    v3 fix: Diktum item IDs use compound key (parent_ordinal:item_num)
    so items from different KESATU/KEDUA sections don't collide.
    """

    RE_MENIMBANG = re.compile(r'^\s*menimbang\s*:?\s*$', re.IGNORECASE)
    RE_MENGINGAT = re.compile(r'^\s*mengingat\s*:?\s*$', re.IGNORECASE)
    RE_MEMUTUSKAN = re.compile(r'memutuskan\s*:?\s*$', re.IGNORECASE)
    RE_MENETAPKAN = re.compile(r'menetapkan\s*:?\s*$', re.IGNORECASE)
    RE_MENINSTRUKSIKAN = re.compile(r'menginstruksikan\s*:?\s*$', re.IGNORECASE)
    RE_ITEM_LETTER = re.compile(r'^\s*([a-z])\s*\.\s*')
    RE_ITEM_NUMBER = re.compile(r'^\s*(\d+)\s*\.\s*')
    RE_ORDINAL = re.compile(
        r'^\s*(KE(?:SATU|DUA|TIGA|EMPAT|LIMA|NAM|TUJUH|DELAPAN|SEMBILAN|SEPULUH|'
        r'SEBELAS|DUA BELAS|TIGA BELAS|EMPAT BELAS|LIMA BELAS|ENAM BELAS|'
        r'TUJUH BELAS|DELAPAN BELAS|SEMBILAN BELAS|DUA PULUH))\b',
        re.IGNORECASE
    )
    RE_PASAL_DIKTUM = re.compile(r'^\s*Pasal\s+(\d+[a-zA-Z]?)\s*\.?\s*$', re.IGNORECASE)

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
        self.decision_keyword = None
        self.parent_ordinal = None  # FIX #7: track KESATU/KEDUA for compound key

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

            # Decision keyword triggers
            if self.section not in ("keputusan", "done"):
                if self.RE_MENINSTRUKSIKAN.search(text):
                    self._flush_item()
                    self.section = "keputusan"
                    self.decision_keyword = "MENGINSTRUKSIKAN"
                    self.current_item = None
                    self.parent_ordinal = None
                    continue

                if self.RE_MEMUTUSKAN.search(text):
                    self._flush_item()
                    self.section = "keputusan"
                    self.decision_keyword = "MEMUTUSKAN"
                    self.current_item = None
                    self.parent_ordinal = None
                    continue

                if self.RE_MENETAPKAN.search(text):
                    self._flush_item()
                    self.section = "keputusan"
                    self.decision_keyword = "MENETAPKAN"
                    self.current_item = None
                    self.parent_ordinal = None
                    continue

            if self.section == "pre":
                continue

            if self.section == "done":
                continue

            # Signature block detection
            text_lower = text.lower()
            if "ttd" in text_lower and len(text) < 20:
                self._flush_item()
                self.section = "done"
                continue

            # In keputusan section: check for item markers
            if self.section == "keputusan":
                m_letter = self.RE_ITEM_LETTER.match(text)
                m_number = self.RE_ITEM_NUMBER.match(text)
                m_ordinal = self.RE_ORDINAL.match(text)
                m_pasal = self.RE_PASAL_DIKTUM.match(text)

                if m_pasal:
                    self._flush_item()
                    self.current_item = f"Pasal {m_pasal.group(1)}"
                    continue

                if m_ordinal:
                    # FIX #7: ordinal = new parent section, reset item tracking
                    self._flush_item()
                    self.parent_ordinal = m_ordinal.group(1).upper()
                    self.current_item = self.parent_ordinal
                    continue

                if m_letter or m_number:
                    self._flush_item()
                    if m_letter:
                        self.current_item = m_letter.group(1)
                    else:
                        self.current_item = m_number.group(1)
                    m = m_letter or m_number
                    remaining = text[m.end():].strip()
                    if remaining:
                        self.current_item_lines.append(remaining)
                    continue

                # Sub-items: accumulate under current item
                if self.current_item:
                    self.current_item_lines.append(text)
                continue

        self._flush_item()
        return self.chunks

    def _flush_item(self):
        if not self.current_item or not self.current_item_lines:
            self.current_item_lines = []
            return

        text = ' '.join(self.current_item_lines).strip()
        if not text or len(text) < MIN_CHUNK_TEXT_LENGTH:
            self.current_item_lines = []
            return

        kw = self.decision_keyword or "MEMUTUSKAN"
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            chunk_type="diktum", chunk_num=self.current_item,
                            parent_ordinal=self.parent_ordinal or "")

        path = f"{kw} > Diktum {self.current_item}"
        if self.parent_ordinal and self.parent_ordinal != self.current_item:
            path = f"{kw} > {self.parent_ordinal} > Diktum {self.current_item}"

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
                "parent_ordinal": self.parent_ordinal,
                "parent": None,
                "path": path,
                "status": "active",
            }
        })
        self.current_item_lines = []

    def _flush_preamble(self):
        self.preamble_lines = []


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

        self.section = "pre"
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
        if not text or len(text) < MIN_CHUNK_TEXT_LENGTH:
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


def _extract_nomor_year_from_title_block(lines, max_lines=10):
    """Extract nomor/year from the first N lines ONLY (title block), never body text.
    
    FIX #5: Previous code searched the entire document and could pick up
    cited regulation numbers from body text. Now restricted to the title
    block where the document's own number is declared.
    """
    title_text = " ".join(l["text"].strip() for l in lines[:max_lines] if l["text"].strip())
    title_text = fix_glyph_corruption(title_text)
    
    # Pattern: "NOMOR X TAHUN YYYY"
    m = re.search(r'NOMOR\s+(\d+[a-zA-Z]*)\s+TAHUN\s+(\d{4})', title_text, re.IGNORECASE)
    if m:
        return m.group(1), m.group(2)
    
    # Pattern: "No. X Tahun YYYY"
    m = re.search(r'No\.?\s*(\d+[a-zA-Z]*)\s+Tahun\s+(\d{4})', title_text, re.IGNORECASE)
    if m:
        return m.group(1), m.group(2)
    
    # Case number pattern: "96/PUU-XVI/2018"
    m = re.search(r'(\d+/[A-Z]+-[A-Z]+/\d{4})', title_text)
    if m:
        return m.group(1), m.group(1).split("/")[-1]
    
    return None, None


def chunk_pdf(pdf_path: str, directory: str, family_map: Dict) -> List[Dict]:
    """
    Main entry point. Extract + parse a PDF into chunks.
    Returns list of chunk dicts with id, text, metadata.
    
    v3: doc_type resolved from title block first (with mismatch flagging),
    directory config as fallback only.
    """
    dir_config = family_map.get(directory, {})
    family = dir_config.get("family", "A")
    default_doc_type = dir_config.get("doc_type", "Unknown")
    default_issuer = dir_config.get("issuer", "Unknown")

    try:
        lines, meta = extract_clean(pdf_path)
    except Exception as e:
        return [{"id": f"ERROR:{directory}", "text": str(e), "metadata": {"error": True}}]

    if not lines:
        return [{"id": f"EMPTY:{directory}", "text": "No extractable text (possibly scanned PDF)", "metadata": {"error": True, "needs_ocr": True}}]

    # FIX #2: Detect doc_type from title block FIRST, use directory as fallback.
    title_block = extract_title_block(lines, max_first_n=10)
    detected_type, detected_issuer = detect_doc_type_from_title(title_block)
    
    doc_type_override = False
    if detected_type != "Unknown":
        doc_type = detected_type
        issuer = detected_issuer
        if default_doc_type != "Unknown" and detected_type != default_doc_type:
            doc_type_override = True
    elif default_doc_type != "Unknown":
        doc_type = default_doc_type
        issuer = default_issuer
    else:
        doc_type = "Unknown"
        issuer = "Unknown"

    # FIX #5: Use meta (extracted from raw first-20-lines before cleaning) as primary.
    # Title-block extraction as fallback only (for cases where PASS 0 missed it).
    nomor = meta.get("nomor")
    year = meta.get("year")
    if not nomor or not year:
        tb_nomor, tb_year = _extract_nomor_year_from_title_block(lines, max_lines=10)
        if not nomor and tb_nomor:
            nomor = tb_nomor
        if not year and tb_year:
            year = tb_year
    
    # UUD 1945 special case: use "?" for nomor (constitution has no regulation number)
    if doc_type == "UUD1945":
        nomor = "?"
        year = "1945"

    # JDIH_KPU special case: page 1 is image-only
    if doc_type == "Unknown" and directory == "JDIH_KPU":
        doc_type = "PKPU"
        issuer = "Komisi Pemilihan Umum"

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
        if doc_type_override:
            chunk["metadata"]["doc_type_override"] = True
            chunk["metadata"]["doc_type_directory"] = default_doc_type

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

    output_path = "/home/z/my-project/download/chunk_results_v3.json"
    with open(output_path, "w", encoding="utf-8") as f:
        json_lib.dump(all_results, f, ensure_ascii=False, indent=2)
    print(f"\nAll results saved to: {output_path}")
