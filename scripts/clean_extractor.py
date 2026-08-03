"""
Clean extraction module for Indonesian legal PDFs.
Fixes: glyph corruption (O/0, l/1), strips header/footer boilerplate,
fixes line-break mid-word, returns clean lines with metadata.
"""
import fitz
import re
import os
from collections import Counter, defaultdict
from typing import List, Dict, Tuple, Optional


# Glyph substitution fixes (letter -> digit in numeric context)
GLYPH_FIXES = [
    # O-as-zero in year/number contexts: 2OO9, 2O2O, 2O24, etc.
    (re.compile(r'(\d)O(\d)'), r'\g<1>0\g<2>'),       # 2OO9 -> 2009
    (re.compile(r'(\d)O(\d)'), r'\g<1>0\g<2>'),       # second pass for 2O24
    # Specific known corruptions
    (re.compile(r'PRESIOEN', re.I), 'PRESIDEN'),
    (re.compile(r'REPLJBLIK', re.I), 'REPUBLIK'),
    (re.compile(r'REPUBUK', re.I), 'REPUBLIK'),
    (re.compile(r'RAAMAT', re.I), 'RAHMAT'),
    # l (lowercase L) as 1 in numbered list context: "l." at line start
    (re.compile(r'^l\.\s'), '1. '),
    # RAKHMAT -> RAHMAT
    (re.compile(r'RAKHMAT', re.I), 'RAHMAT'),
    # Cukuo -> Cukup (seen in some docs)
    (re.compile(r'Cukuo', re.I), 'Cukup'),
]


def fix_glyph_corruption(text: str) -> str:
    """Apply all known glyph substitution fixes."""
    for pattern, replacement in GLYPH_FIXES:
        text = pattern.sub(replacement, text)
    return text


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
    """Lines that are structural markers, not body text."""
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
    pages_lines = []  # List of lists, per page
    all_lines = []
    all_text_raw = ""
    
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
                all_text_raw += text + " "
        pages_lines.append(page_lines)
    
    doc.close()
    
    # Detect boilerplate
    boilerplate = detect_and_strip_boilerplate(pages_lines)
    
    # Mark boilerplate lines and fix glyph corruption
    for line in all_lines:
        stripped = line["text"].strip()
        line["is_boilerplate"] = (
            stripped in boilerplate or
            is_page_number_line(stripped) or
            is_likely_heading_or_preamble(stripped)
        )
        if not line["is_boilerplate"]:
            line["text"] = fix_glyph_corruption(line["text"])
    
    # Remove boilerplate lines from active parsing set
    active_lines = [l for l in all_lines if not l["is_boilerplate"]]
    
    # Try to extract nomor and year from TITLE BLOCK only (first 15 non-boilerplate lines)
    title_block_text = " ".join(
        l["text"].strip() for l in active_lines[:15]
    )
    nomor, year = extract_nomor_year(title_block_text)
    
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
        "extraction_issues": [],
    }
    
    return active_lines, metadata


def extract_nomor_year(full_text: str) -> Tuple[Optional[str], Optional[str]]:
    """Extract document number and year from title block text."""
    # Pattern: "NOMOR X TAHUN YYYY" or "Nomor X Tahun YYYY"
    m = re.search(r'NOMOR\s+(\d+[\w]*)\s+TAHUN\s+(\d{4})', full_text, re.IGNORECASE)
    if m:
        return m.group(1), m.group(2)
    
    # Try "No. X Tahun YYYY"
    m = re.search(r'No\.?\s*(\d+[\w]*)\s+Tahun\s+(\d{4})', full_text, re.IGNORECASE)
    if m:
        return m.group(1), m.group(2)
    
    # Try case number pattern for court rulings: "96/PUU-XVI/2018"
    m = re.search(r'(\d+/[A-Z]+-[A-Z]+/\d{4})', full_text)
    if m:
        return m.group(1), m.group(1).split("/")[-1]
    
    return None, None


def extract_title_block(lines: List[Dict], max_first_n=10) -> str:
    """Get the title block (first ~10 non-boilerplate lines) for doc_type detection.
    Only use lines from page 1 to avoid picking up preamble references."""
    title_lines = []
    for line in lines:
        if line["page"] > 1:
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
    
    title = extract_title_block(lines)
    doc_type, issuer = detect_doc_type_from_title(title)
    print(f"Detected: doc_type={doc_type}, issuer={issuer}")
    print(f"\nFirst 30 active lines:")
    for l in lines[:30]:
        print(f"  p{l['page']:02d} fs={l['font_size']:5.1f} B={str(l['is_bold']):5s} | {l['text'][:100]}")
