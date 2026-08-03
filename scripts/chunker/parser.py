import re
from dataclasses import dataclass, field
from typing import Optional


# ─── Data Classes ─────────────────────────────────────────────────────────


@dataclass
class Ayat:
    """A single clause (ayat) within a Pasal."""
    number: int
    text: str
    huruf_items: list[dict] = field(default_factory=list)  # [{letter, text}]
    angka_items: list[dict] = field(default_factory=list)  # [{number, text}]


@dataclass
class Pasal:
    """An article (Pasal) containing one or more Ayat."""
    number: str  # str to handle "6A", "7A" etc.
    title: Optional[str] = None  # Some pasal have inline titles
    ayat_list: list[Ayat] = field(default_factory=list)
    raw_text: str = ""  # Full text of this pasal (for parent chunk)
    start_line: int = -1
    end_line: int = -1


@dataclass
class Bab:
    """A chapter (BAB)."""
    number: str  # "I", "II", "IIA", etc.
    title: str = ""
    pasal_list: list[Pasal] = field(default_factory=list)


@dataclass
class Bagian:
    """A part/section within a BAB."""
    number: str
    title: str = ""
    pasal_list: list[Pasal] = field(default_factory=list)


@dataclass
class ParsedDocument:
    """Full parsed structure of a statute PDF."""
    filename: str
    title: str = ""
    reg_number: str = ""
    preamble_text: str = ""  # Menimbang + Mengingat + enacting clause
    babs: list[Bab] = field(default_factory=list)
    loose_pasals: list[Pasal] = field(default_factory=list)  # Pasals outside any BAB
    all_pasals: list[Pasal] = field(default_factory=list)  # Flat list of all
    issues: list[dict] = field(default_factory=list)  # Parsing issues


# ─── Regex Patterns ────────────────────────────────────────────────────────

# STRICT: Only matches lines that are actual Pasal HEADINGS
# Handles: Pasal 1, *Pasal 3, * Pasal 6A, Pasal I, * Pasal II
RE_PASAL_HEADING = re.compile(
    r'^\*?\s*Pasal\s+([0-9]+[a-zA-Z]?|[IVXLCDM]+[a-zA-Z]?)\s*\.?\s*(?:\.\.\.)?\s*$',
    re.IGNORECASE
)

# LOOSE: Matches any mention of Pasal (for detecting inline references)
RE_PASAL_MENTION = re.compile(
    r'pasal\s+([0-9]+[a-zA-Z]?|[IVXLCDM]+[a-zA-Z]?)', re.IGNORECASE
)

# Ayat (clause): (1), (2), etc. at line start
RE_AYAT = re.compile(r'^\(\s*(\d+)\s*\)\s*(.*)', re.DOTALL)

# BAB heading: "BAB I", "BAB II", "* BAB IV", "BAB VIIA", "BAB IX A"
RE_BAB = re.compile(
    r'^\*?\s*BAB\s+([IVXLCDM]+\s*[a-zA-Z]?)\s*\.?\s*(?:\.\.\.)?\s*$', re.IGNORECASE
)

# BAB heading with title: "BAB I KETENTUAN UMUM"
RE_BAB_TITLE = re.compile(
    r'^\*?\s*BAB\s+([IVXLCDM]+\s*[a-zA-Z]?)\s+(.+)$', re.IGNORECASE
)

# Bagian heading
RE_BAGIAN = re.compile(
    r'^\*?\s*Bagian\s+(Ke-\d+|Pertama|Kedua|Ketiga|Keempat|Kelima|Keenam|Ketujuh|Kedelapan|Kesembilan|Kesepuluh|[IVXLCDM]+)\s+(.*)',
    re.IGNORECASE
)

# Paragraf heading
RE_PARAGRAF = re.compile(
    r'^\*?\s*Paragraf\s+(\d+|[IVXLCDM]+)\s+(.*)', re.IGNORECASE
)

# Preamble markers
RE_MENIMBANG = re.compile(r'^Menimbang\s*:', re.IGNORECASE)
RE_MENGINGAT = re.compile(r'^Mengingat\s*:', re.IGNORECASE)
RE_MEMUTUSKAN = re.compile(r'^Memutuskan\s*:', re.IGNORECASE)
RE_MENETAPKAN = re.compile(r'^Menetapkan\s*:', re.IGNORECASE)

# ─── Helper Functions ──────────────────────────────────────────────────────


def _get_next_nonempty(lines, idx):
    """Return next non-empty line text, or None."""
    for i in range(idx + 1, len(lines)):
        if lines[i]["text"].strip():
            return lines[i]["text"].strip()
    return None


def _get_prev_nonempty(lines, idx):
    """Return previous non-empty line text, or None."""
    for i in range(idx - 1, -1, -1):
        if lines[i]["text"].strip():
            return lines[i]["text"].strip()
    return None


def _is_larger_font(lines, idx, body_size, threshold=1.05):
    """Check if line at idx has a larger font than body text."""
    if idx >= len(lines):
        return False
    return lines[idx]["font_size"] >= body_size * threshold


def _is_likely_ayat(text):
    """Check if text looks like ayat content (starts with parenthesis number)."""
    return bool(re.match(r'^\(\s*\d+\s*\)', text))


# ─── Main Parser ────────────────────────────────────────────────────────────


def parse_statute(lines, body_font_size=None, filename=""):
    """Parse font-annotated lines into structured Pasal/Ayat/BAB hierarchy.
    
    Args:
        lines: List of dicts from extractor.extract_with_fonts()
        body_font_size: Detected body font size (auto-detected if None)
        filename: Source filename for issue tracking
    
    Returns:
        ParsedDocument
    """
    if body_font_size is None:
        from extractor import detect_body_font_size
        body_font_size = detect_body_font_size(lines)
    
    doc = ParsedDocument(filename=filename)
    
    # ── Phase 1: Identify structural boundaries ──
    # Classify each line: pasal_heading, ayat, bab, bagian, preamble, body
    
    line_roles = []  # [(role, match_data), ...]
    
    preamble_started = False
    preamble_ended = False
    
    for i, line_info in enumerate(lines):
        text = line_info["text"].strip()
        if not text:
            line_roles.append(("empty", None))
            continue
        
        # Check preamble markers
        if RE_MENIMBANG.match(text) or RE_MENGINGAT.match(text):
            preamble_started = True
            line_roles.append(("preamble", None))
            continue
        if RE_MEMUTUSKAN.match(text) or RE_MENETAPKAN.match(text):
            preamble_ended = True
            preamble_started = False
            line_roles.append(("preamble_end", None))
            continue
        
        # Check BAB heading
        m_bab = RE_BAB.match(text) or RE_BAB_TITLE.match(text)
        if m_bab and not _is_inline_reference(lines, i, keyword="BAB"):
            preamble_ended = True
            bab_num = m_bab.group(1).strip()  # remove spaces like "IX A" -> "IXA"
            line_roles.append(("bab", bab_num))
            continue
        
        # Check Bagian
        m_bag = RE_BAGIAN.match(text)
        if m_bag and not _is_inline_reference(lines, i, keyword="Bagian"):
            line_roles.append(("bagian", (m_bag.group(1), m_bag.group(2).strip())))
            continue
        
        # Check Paragraf
        m_par = RE_PARAGRAF.match(text)
        if m_par and not _is_inline_reference(lines, i, keyword="Paragraf"):
            line_roles.append(("paragraf", (m_par.group(1), m_par.group(2).strip())))
            continue
        
        # Check Pasal heading (STRICT + font validation)
        m_pasal = RE_PASAL_HEADING.match(text)
        if m_pasal and _is_real_pasal_heading(lines, i, body_font_size):
            preamble_ended = True
            line_roles.append(("pasal_heading", m_pasal.group(1)))
            continue
            
        # Check Ayat
        m_ayat = RE_AYAT.match(text)
        if m_ayat:
            line_roles.append(("ayat", (int(m_ayat.group(1)), m_ayat.group(2).strip())))
            continue
        
        # Default
        if preamble_started and not preamble_ended:
            line_roles.append(("preamble", None))
        else:
            line_roles.append(("body", None))
    
    # ── Phase 2: Group lines into Pasals ──
    # Find pasal boundaries and group content between them
    
    pasal_boundaries = []  # (line_idx, pasal_number)
    for i, (role, data) in enumerate(line_roles):
        if role == "pasal_heading":
            pasal_boundaries.append((i, data))
    
    # Build Pasal objects
    for p_idx, (start, pasal_num) in enumerate(pasal_boundaries):
        end = pasal_boundaries[p_idx + 1][0] if p_idx + 1 < len(pasal_boundaries) else len(lines)
        
        pasal = Pasal(
            number=pasal_num,
            start_line=start,
            end_line=end,
        )
        
        # Collect raw text and parse ayat
        pasal_text_lines = []
        current_ayat = None
        current_ayat_num = None
        
        for i in range(start + 1, end):
            role, data = line_roles[i]
            text = lines[i]["text"].strip()
            
            if role in ("empty", "bab", "bagian", "paragraf", "pasal_heading"):
                # Structural marker — end current ayat
                if current_ayat is not None:
                    pasal.ayat_list.append(current_ayat)
                    current_ayat = None
                if role != "empty":
                    pasal_text_lines.append("")
                continue
            
            if role == "ayat":
                # Save previous ayat
                if current_ayat is not None:
                    pasal.ayat_list.append(current_ayat)
                
                ayat_num, ayat_text = data
                # Check for skipped ayat numbers
                if current_ayat_num is not None and ayat_num != current_ayat_num + 1:
                    if ayat_num > current_ayat_num + 1:
                        doc.issues.append({
                            "type": "skipped_ayat",
                            "pasal": pasal_num,
                            "expected": current_ayat_num + 1,
                            "found": ayat_num,
                            "line": i,
                        })
                
                current_ayat = Ayat(number=ayat_num, text=ayat_text)
                current_ayat_num = ayat_num
                pasal_text_lines.append(f"({ayat_num}) {ayat_text}")
                continue
            
            # Body text — append to current ayat or pasal-level text
            if role == "body" or role == "preamble":
                if current_ayat is not None:
                    current_ayat.text += " " + text
                    pasal_text_lines[-1] += " " + text  # append to last line
                else:
                    pasal_text_lines.append(text)
        
        # Don't forget the last ayat
        if current_ayat is not None:
            pasal.ayat_list.append(current_ayat)
        
        pasal.raw_text = "\n".join(pasal_text_lines).strip()
        
        # Handle single-ayat pasals (no explicit (1) marker)
        if not pasal.ayat_list and pasal.raw_text:
            pasal.ayat_list.append(Ayat(number=1, text=pasal.raw_text))
        
        doc.all_pasals.append(pasal)
    
    # ── Phase 2.5: Deduplicate pasals with same number ──
    # UUD has *Pasal N (amended) and Pasal N (original) — keep first
    seen = set()
    deduped = []
    for p in doc.all_pasals:
        normalized = p.number.strip().upper().lstrip('*').strip()
        if normalized in seen:
            doc.issues.append({
                "type": "duplicate_pasal",
                "pasal": p.number,
                "detail": f"Duplicate Pasal {p.number}, keeping first version",
            })
            continue
        seen.add(normalized)
        deduped.append(p)
    doc.all_pasals = deduped
    
    # ── Phase 3: Extract preamble ──
    preamble_lines = []
    for i, (role, data) in enumerate(line_roles):
        if role == "preamble":
            preamble_lines.append(lines[i]["text"].strip())
        elif role == "preamble_end":
            preamble_lines.append(lines[i]["text"].strip())
            break
    doc.preamble_text = "\n".join(preamble_lines).strip()
    
    # ── Phase 4: Build BAB hierarchy ──
    current_bab = None
    current_bagian = None
    
    for i, (role, data) in enumerate(line_roles):
        if role == "bab":
            current_bab = Bab(number=data)
            current_bagian = None
            doc.babs.append(current_bab)
        elif role == "bagian" and current_bab is not None:
            current_bagian = Bagian(number=data[0], title=data[1])
            # (we track bagian but keep pasals flat in all_pasals)
        elif role == "pasal_heading":
            # Find the pasal in all_pasals
            for p in doc.all_pasals:
                if p.number == data and p.start_line == i:
                    if current_bab is not None:
                        current_bab.pasal_list.append(p)
                    else:
                        doc.loose_pasals.append(p)
                    break
    
    return doc


def _is_real_pasal_heading(lines, idx, body_font_size):
    """3-layer validation: pattern + font + context."""
    text = lines[idx]["text"].strip()
    
    # Layer 1: Must match strict pattern
    if not RE_PASAL_HEADING.match(text):
        return False
    
    # Layer 2: Font size check (heading usually >= body size)
    # Be lenient — only reject if clearly smaller
    if lines[idx]["font_size"] < body_font_size * 0.95:
        return False
    
    # Layer 3: Context check
    # If next non-empty line is an ayat → definitely a heading
    next_text = _get_next_nonempty(lines, idx)
    if next_text and _is_likely_ayat(next_text):
        return True
    
    # If prev non-empty line is MENETAPKAN/MEMUTUSKAN → first pasal
    prev_text = _get_prev_nonempty(lines, idx)
    if prev_text and (RE_MENETAPKAN.match(prev_text) or RE_MEMUTUSKAN.match(prev_text)):
        return True
    
    # If prev line is a BAB/Bagian → heading
    if prev_text and (RE_BAB.match(prev_text) or RE_BAGIAN.match(prev_text)):
        return True
    
    # If the line is bold or larger font → likely heading
    if lines[idx]["is_bold"] or lines[idx]["font_size"] > body_font_size:
        return True
    
    # Default: accept if strict pattern matched (most cases)
    return True


def _is_inline_reference(lines, idx, keyword="pasal"):
    """Check if a line containing BAB/Pasal is an inline reference, not a heading.
    
    An inline reference has substantial text BEFORE the keyword,
    e.g. 'sebagaimana dimaksud dalam Pasal 5 ayat (1)'
    """
    text = lines[idx]["text"].strip()
    
    kw_lower = keyword.lower()
    kw_pos = text.lower().find(kw_lower)
    
    if kw_pos < 0:
        return False
    
    # If keyword is not near the start, likely inline reference
    # Allow up to 2 chars before (for asterisk)
    if kw_pos > 2:
        prefix = text[:kw_pos].strip()
        if len(prefix) > 1 and not prefix.endswith(':'):
            return True
    
    return False


# ─── Debug / Inspection ─────────────────────────────────────────────────────


def debug_line_roles(lines, body_font_size=None):
    """Print line-by-line role classification for debugging."""
    if body_font_size is None:
        from extractor import detect_body_font_size
        body_font_size = detect_body_font_size(lines)
    
    for i, line_info in enumerate(lines):
        text = line_info["text"].strip()
        if not text:
            continue
        fs = line_info["font_size"]
        bold = "B" if line_info["is_bold"] else " "
        
        # Determine role
        role = "???"
        if RE_BAB.match(text) or RE_BAB_TITLE.match(text):
            role = "BAB" if not _is_inline_reference(lines, i) else "ref"
        elif RE_PASAL_HEADING.match(text):
            role = "PASAL" if _is_real_pasal_heading(lines, i, body_font_size) else "ref"
        elif RE_AYAT.match(text):
            role = "AYAT"
        elif RE_BAGIAN.match(text):
            role = "BAGIAN"
        elif RE_PARAGRAF.match(text):
            role = "PARAGRAF"
        elif RE_MENIMBANG.match(text):
            role = "MENIMBANG"
        elif RE_MENGINGAT.match(text):
            role = "MENGINGAT"
        elif RE_MEMUTUSKAN.match(text):
            role = "MEMUTUSKAN"
        elif RE_MENETAPKAN.match(text):
            role = "MENETAPKAN"
        
        marker = " → " if fs > body_font_size else "   "
        print(f"{i:4d} [{bold}] fs={fs:5.1f}{marker} {role:14s} | {text[:80]}")
