"""
Clean extraction module for Indonesian legal PDFs.

v2 fixes (based on QA assessment):
  - Cross-page text joining (bug #1/#4: orphan chunks from page breaks)
  - Fragment line merging (bug #4: word-level extraction artifacts)
  - Stamp pattern stripping (bug #3: SK No ###### A)
  - Garbled header detection (bug #3: R EPI'FILIK INOONESIA)
  - Title extraction: I/l-as-1 in nomor (bug #7: perpres I27)
  - Title block fallback to page 2+ when page 1 is empty (bug #7: JDIH_KPU)
"""
import fitz
import re
import os
from collections import Counter, defaultdict
from typing import List, Dict, Tuple, Optional


# Glyph substitution fixes (letter -> digit in numeric context)
GLYPH_FIXES = [
    # O-as-zero in year/number contexts: 2OO9, 2O2O, 2O24, etc.
    (re.compile(r'(\d)O(\d)'), r'\g<1>0\g<2>'),
    (re.compile(r'(\d)O(\d)'), r'\g<1>0\g<2>'),       # second pass
    # I/l-as-1 in number context (perpres I27 = 127)
    (re.compile(r'(\d|[Nn]omor)\s+I(\d{2,})', re.IGNORECASE), r'\g<1> 1\g<2>'),
    # Specific known corruptions
    (re.compile(r'PRESIOEN', re.I), 'PRESIDEN'),
    (re.compile(r'REPLJBLIK', re.I), 'REPUBLIK'),
    (re.compile(r'REPUBUK', re.I), 'REPUBLIK'),
    (re.compile(r'RAAMAT', re.I), 'RAHMAT'),
    (re.compile(r'RAKHMAT', re.I), 'RAHMAT'),
    (re.compile(r'Cukuo', re.I), 'Cukup'),
    # l (lowercase L) as 1 in numbered list context: "l." at line start
    (re.compile(r'^l\.\s'), '1. '),
    # O-as-zero not just between digits but also at word boundary: "2O " or "2O23" 
    (re.compile(r'(\d)O(\s|T|$)'), r'\g<1>0\g<2>'),
    # Fix known garbled REPUBLIK patterns
    (re.compile(r'R\s*E\s*PI[\s\',\-]*[IL]I[\s\',\-]*K\s*INOONESIA', re.I), 'REPUBLIK INDONESIA'),
    (re.compile(r'REPI\s*,\s*IEILIK\s*INDONESIA', re.I), 'REPUBLIK INDONESIA'),
    (re.compile(r'REPUBLIK\s+INDOONESIA', re.I), 'REPUBLIK INDONESIA'),
]


# Structural markers that must NEVER be merged with adjacent lines
STRUCTURAL_MARKERS = re.compile(
    r'^(?:'
    r'BAB\s+[IVXLCDM]+|'
    r'BAGIAN\s+\w+|'
    r'PARAGRAF\s+\w+|'
    r'Pasal\s+[\dIVXLCDM]+[a-zA-Z]?|'
    r'\(\s*\d+[a-zA-Z]?\s*\)|'
    r'MEMUTUSI?\s*AN\s*:?$|'
    r'MENETAPKAN\s*:?$|'
    r'MENGINSTRUKSIKAN\s*:?$|'
    r'MENIMBANG\s*:?$|'
    r'MENGINGAT\s*:?$|'
    r'MENGADILI\s*:?$|'
    r'DENGAN RAHMAT|'
    r'KE(?:SATU|DUA|TIGA|EMPAT|LIMA|NAM|TUJUH|DELAPAN|SEMBILAN|SEPULUH' 
    r'|SEBELAS|DUA\s+BELAS|TIGA\s+BELAS|EMPAT\s+BELAS|LIMA\s+BELAS' 
    r'|ENAM\s+BELAS|TUJUH\s+BELAS|DELAPAN\s+BELAS|SEMBILAN\s+BELAS'
    r'|DUA\s+PULUH)\b|'
    r'PENUTUP\b|'
    r'Putusan\s*:?$'
    r')',
    re.IGNORECASE
)


# Noise patterns (header stamps, tracking numbers, garbled text)
RE_SETNEG_STAMP = re.compile(r'^SK\s+No\s+\d+\s*A$', re.IGNORECASE)
RE_PAGE_NUMBER = re.compile(r'^-\s*\d+\s*-$')
RE_GARBLED_HEADER = re.compile(
    r'^(?:'
    r'R\s*E\s*P(?:I|L|1|J)[\s\',\-]*[IL1J][\s\',\-]*(?:K|L|I|E)[\s\',\-]*(?:I|K|L)'
    r'|FRESIDEN'
    r'|MEMUTUSI?\s*\(?\s*AN'
    r')',
    re.IGNORECASE
)
RE_SHORT_NOISE = re.compile(r'^[A-Z]\.?$')  # Single letter like "A." or "b"


def fix_glyph_corruption(text: str) -> str:
    """Apply all known glyph substitution fixes."""
    for pattern, replacement in GLYPH_FIXES:
        text = pattern.sub(replacement, text)
    return text


def is_noise_line(text: str) -> bool:
    """Check if a line is known noise (stamp, garbled header, page number)."""
    t = text.strip()
    if not t:
        return True
    if RE_SETNEG_STAMP.match(t):
        return True
    if RE_PAGE_NUMBER.match(t):
        return True
    if RE_GARBLED_HEADER.match(t):
        return True
    # Standalone page numbers: just a digit, 1-4 chars
    if re.match(r'^\d{1,4}$', t):
        return True
    return False


def is_structural_marker(text: str) -> bool:
    """Check if a line is a structural marker that should not be merged."""
    return bool(STRUCTURAL_MARKERS.match(text.strip()))


def join_fragment_lines(all_lines: List[Dict]) -> List[Dict]:
    """
    Merge word-level fragments into proper sentences.
    
    Handles two fragmentation patterns:
    1. PDF with word-per-line extraction (JDIH_Komdigi: 462 lines <= 5 chars)
    2. Page-break splits where a sentence continues on the next page
    
    Strategy: merge current line into previous if:
    - Current line is NOT a structural marker
    - AND (current line starts with lowercase
    -      OR current line is short (<50 chars) and previous line doesn't end with sentence punctuation)
    """
    if not all_lines:
        return all_lines
    
    SENTENCE_END = re.compile(r'[.?!;:]\s*$')
    merged = [all_lines[0].copy()]
    
    for i in range(1, len(all_lines)):
        curr = all_lines[i]
        curr_text = curr["text"].strip()
        
        if not curr_text or is_structural_marker(curr_text):
            merged.append(curr.copy())
            continue
        
        prev = merged[-1]
        prev_text = prev["text"].strip()
        
        if not prev_text:
            merged.append(curr.copy())
            continue
        
        should_merge = False
        
        # NEVER merge into a structural marker (prev line is BAB, Menimbang, etc.)
        if is_structural_marker(prev_text):
            merged.append(curr.copy())
            continue
        
        # NEVER merge standalone list markers (a. b. 1. 2.)
        if re.match(r'^[a-z]\.$', curr_text) or re.match(r'^\d+\.$', curr_text):
            merged.append(curr.copy())
            continue
        
        # Case 1: current line starts with lowercase -> continuation
        if curr_text[0].islower():
            should_merge = True
        
        # Case 2: short line after non-sentence-ending previous line
        elif len(curr_text) < 50 and not SENTENCE_END.search(prev_text):
            # But don't merge if current is a single uppercase word (likely a heading)
            if not (curr_text.isupper() and ' ' not in curr_text and len(curr_text) > 2):
                should_merge = True
        
        # Case 3: previous line ends with hyphen (word split across lines)
        if prev_text.endswith('-') and not prev_text.endswith('--'):
            should_merge = True
        
        if should_merge:
            # Merge: join with space, keep metadata from the earlier line
            prev["text"] = prev_text + " " + curr_text
            # Keep the later page number if it changed (cross-page join)
            if curr["page"] != prev["page"]:
                prev["_page_joined"] = True
            # Keep max font size
            prev["font_size"] = round(max(prev["font_size"], curr["font_size"]), 1)
            prev["is_bold"] = prev["is_bold"] or curr["is_bold"]
        else:
            merged.append(curr.copy())
    
    return merged


def detect_and_strip_boilerplate(pages_lines: List[List[Dict]], threshold_ratio=0.4) -> List[str]:
    """
    Detect lines that repeat across many pages (header/footer).
    Returns set of boilerplate strings to strip.
    """
    if not pages_lines:
        return set()
    
    total_pages = len(pages_lines)
    line_occurrences = defaultdict(list)
    
    for page_idx, lines in enumerate(pages_lines):
        for line in lines:
            stripped = line["text"].strip()
            if 3 < len(stripped) < 120:
                line_occurrences[stripped].append(page_idx)
    
    min_pages = max(3, int(total_pages * threshold_ratio))
    boilerplate = set()
    for text, pages in line_occurrences.items():
        if len(set(pages)) >= min_pages:
            boilerplate.add(text)
    
    return boilerplate


def is_page_number_line(text: str) -> bool:
    """Check if line is a page number indicator (e.g. '- 2 -')."""
    t = text.strip()
    if re.match(r'^-\s*\d+\s*-$', t):
        return True
    if re.match(r'^\d+\s*$', t) and len(t) <= 4:
        return True
    return False


def is_likely_heading_or_preamble(text: str) -> bool:
    """Lines that are structural markers or letterhead, not body text."""
    t = text.strip().upper()
    headers = [
        "PRESIDEN REPUBLIK INDONESIA",
        "REPUBLIK INDONESIA",
        "MENTERI",
        "DENGAN RAHMAT TUHAN YANG MAHA ESA",
        "SALINAN",
        "LEMBARAN NEGARA",
        "TAMBAHAN LEMBARAN NEGARA",
    ]
    for h in headers:
        if t == h or (len(t) > 10 and t.startswith(h[:10])):
            return True
    return False


def extract_clean(pdf_path: str) -> Tuple[List[Dict], Dict]:
    """
    Main extraction function. Returns (lines, metadata).
    Each line: { page, text, font_size, is_bold, font_name, x_indent, is_boilerplate }
    metadata: { file, doc_type_hint, issuer_hint, nomor, year, total_pages, total_lines, extraction_issues }
    """
    doc = fitz.open(pdf_path)
    pages_lines = []  # List of lists, per page (for boilerplate detection)
    all_lines = []
    
    for page_num, page in enumerate(doc):
        page_lines = []
        blocks = page.get_text("dict")["blocks"]
        for block in blocks:
            if block["type"] != 0:
                continue
            for line in block["lines"]:
                spans = line["spans"]
                if not spans:
                    continue
                text = "".join(s["text"] for s in spans).strip()
                if not text:
                    continue
                font_size = round(max(s["size"] for s in spans), 1)
                is_bold = any("bold" in s["font"].lower() for s in spans)
                font_name = spans[0]["font"]
                x0 = round(line["bbox"][0], 1)
                
                line_data = {
                    "page": page_num + 1,
                    "text": text,
                    "font_size": font_size,
                    "is_bold": is_bold,
                    "font_name": font_name,
                    "x_indent": x0,
                }
                page_lines.append(line_data)
                all_lines.append(line_data)
        pages_lines.append(page_lines)
    
    doc.close()
    
    issues = []
    
    # === PASS 0: Extract nomor/year from RAW text before any cleaning ===
    # This prevents the nomor/year line from being lost to boilerplate stripping
    # or aggressive fragment joining.
    raw_title = " ".join(
        l["text"].strip() for l in all_lines[:20]
    )
    raw_title = fix_glyph_corruption(raw_title)
    nomor, year = extract_nomor_year(raw_title)
    
    # Special case: UUD 1945 has no nomor/year
    if not nomor and not year:
        tb_upper = raw_title.upper()
        if "UNDANG-UNDANG DASAR" in tb_upper or "UUD 1945" in tb_upper:
            nomor = "1945"
            year = "1945"
    
    # === PASS 1: Join fragment lines (word-level splits, page-break splits) ===
    all_lines = join_fragment_lines(all_lines)
    
    # === PASS 2: Apply glyph fixes ===
    for line in all_lines:
        line["text"] = fix_glyph_corruption(line["text"])
    
    # === PASS 3: Strip noise lines (stamps, garbled headers, page numbers) ===
    noise_count = 0
    for line in all_lines:
        if is_noise_line(line["text"]):
            line["_noise"] = True
            noise_count += 1
    if noise_count > 0:
        issues.append(f"noise_lines_stripped={noise_count}")
    
    # === PASS 4: Detect and strip boilerplate (repeating headers/footers) ===
    boilerplate = detect_and_strip_boilerplate(pages_lines)
    
    # Mark boilerplate and noise lines
    for line in all_lines:
        stripped = line["text"].strip()
        line["is_boilerplate"] = (
            stripped in boilerplate or
            is_page_number_line(stripped) or
            is_likely_heading_or_preamble(stripped) or
            line.get("_noise", False)
        )
    
    # Remove boilerplate/noise lines from active parsing set
    active_lines = [l for l in all_lines if not l["is_boilerplate"]]
    
    # Special case: UUD 1945 has no nomor/year
    if not nomor and not year:
        # Check if this is UUD 1945 from title
        tb_upper = title_block_text.upper()
        if "UNDANG-UNDANG DASAR" in tb_upper or "UUD 1945" in tb_upper:
            nomor = "1945"
            year = "1945"
    
    # Font stats
    sizes = [l["font_size"] for l in active_lines]
    size_counts = Counter(sizes) if sizes else Counter()
    most_common_size = size_counts.most_common(1)[0][0] if size_counts else 12.0
    
    metadata = {
        "file": os.path.basename(pdf_path),
        "total_pages": len(pages_lines),
        "total_lines_raw": len(all_lines),
        "total_lines_active": len(active_lines),
        "boilerplate_lines_removed": len(all_lines) - len(active_lines),
        "boilerplate_samples": list(boilerplate)[:10],
        "nomor": nomor,
        "year": year,
        "most_common_font_size": most_common_size,
        "font_sizes": sorted(set(sizes)),
        "extraction_issues": issues,
    }
    
    return active_lines, metadata


def extract_nomor_year(full_text: str) -> Tuple[Optional[str], Optional[str]]:
    """Extract document number and year from title block text."""
    # Apply glyph fixes first for better matching
    fixed = fix_glyph_corruption(full_text)
    
    # Pattern: "NOMOR X TAHUN YYYY" (handles I/l-as-1 after glyph fix)
    m = re.search(r'NOMOR\s+(\d+[a-zA-Z]*)\s+TAHUN\s+(\d{4})', fixed, re.IGNORECASE)
    if m:
        return m.group(1), m.group(2)
    
    # Try "No. X Tahun YYYY"
    m = re.search(r'No\.?\s*(\d+[a-zA-Z]*)\s+Tahun\s+(\d{4})', fixed, re.IGNORECASE)
    if m:
        return m.group(1), m.group(2)
    
    # Try case number pattern for court rulings: "96/PUU-XVI/2018"
    m = re.search(r'(\d+/[A-Z]+-[A-Z]+/\d{4})', fixed)
    if m:
        return m.group(1), m.group(1).split("/")[-1]
    
    # Try "Nomor X Tahun YYYY" with possible multi-word gap
    m = re.search(r'Nomor\s+(\d+)\s+Tahun\s+(\d{4})', fixed, re.IGNORECASE)
    if m:
        return m.group(1), m.group(2)
    
    return None, None


def extract_title_block(lines: List[Dict], max_first_n=15) -> str:
    """
    Get the title block for doc_type detection and nomor/year extraction.
    Looks at page 1 first. If page 1 is empty (image cover), falls back to page 2+.
    Only uses non-boilerplate lines.
    """
    # Check if page 1 has any lines at all
    page1_lines = [l for l in lines if l.get("page") == 1 and not l.get("is_boilerplate", False)]
    
    # If page 1 is empty, start from page 2
    start_page = 1
    if not page1_lines:
        start_page = 2
    
    title_lines = []
    for line in lines:
        if line["page"] > start_page:
            break
        if line["page"] < start_page:
            continue
        if not line.get("is_boilerplate", False):
            title_lines.append(line["text"].strip())
        if len(title_lines) >= max_first_n:
            break
    
    # If still empty (very short doc), use all available lines from earliest page
    if not title_lines and lines:
        earliest_page = lines[0]["page"]
        for line in lines:
            if line["page"] > earliest_page:
                break
            if not line.get("is_boilerplate", False):
                title_lines.append(line["text"].strip())
            if len(title_lines) >= max_first_n:
                break
    
    return " ".join(title_lines)


def detect_doc_type_from_title(title_block: str) -> Tuple[str, str]:
    """
    Detect document type and issuer from title block text.
    Order: most specific patterns first to avoid misclassification.
    """
    t = title_block.upper()
    if "MAHKAMAH KONSTITUSI" in t:
        return "Putusan_MK", "Mahkamah Konstitusi"
    if "KETETAPAN" in t and "MAJELIS PERMUSYAWARATAN" in t:
        return "Tap_MPR", "MPR"
    if "PERATURAN PEMERINTAH PENGGANTI" in t:
        return "PerPPU", "Presiden"
    if "PERATURAN PEMERINTAH" in t:
        return "PP", "Presiden"
    if "PERATURAN PRESIDEN" in t:
        return "Perpres", "Presiden"
    if "PERATURAN DAERAH" in t:
        return "Perda", "Kepala Daerah"
    if "KEPUTUSAN PRESIDEN" in t:
        return "Keppres", "Presiden"
    if "INSTRUKSI PRESIDEN" in t:
        return "Inpres", "Presiden"
    if "UNDANG-UNDANG DASAR" in t or "UUD 1945" in t:
        return "UUD1945", "BPUPKI/PPKI"
    if "UNDANG-UNDANG" in t:
        return "UU", "Presiden/DPR"
    if "KEPUTUSAN MENTERI" in t:
        return "Kepmen", "Menteri"
    if "PERATURAN MENTERI" in t:
        return "Permen", "Menteri"
    if "PERATURAN KPU" in t or "PKPU" in t:
        return "PKPU", "KPU"
    if "PERATURAN" in t:
        return "Peraturan", "Unknown"
    return "Unknown", "Unknown"


if __name__ == "__main__":
    import sys
    import json
    
    path = sys.argv[1] if len(sys.argv) > 1 else "/home/z/my-project/download/samples/uu/uunomor41tahun2014.pdf"
    lines, meta = extract_clean(path)
    
    print(f"File: {meta['file']}")
    print(f"Pages: {meta['total_pages']}, Raw lines: {meta['total_lines_raw']}, Active: {meta['total_lines_active']}")
    print(f"Boilerplate removed: {meta['boilerplate_lines_removed']}")
    print(f"Boilerplate samples: {meta['boilerplate_samples'][:5]}")
    print(f"Nomor: {meta['nomor']}, Year: {meta['year']}")
    print(f"Font sizes: {meta['font_sizes']}")
    print(f"Issues: {meta['extraction_issues']}")
    
    title = extract_title_block(lines)
    doc_type, issuer = detect_doc_type_from_title(title)
    print(f"Detected: doc_type={doc_type}, issuer={issuer}")
    print(f"\nFirst 30 active lines:")
    for l in lines[:30]:
        print(f"  p{l['page']:02d} fs={l['font_size']:5.1f} B={str(l['is_bold']):5s} | {l['text'][:100]}")
