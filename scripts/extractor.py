"""Font-size-aware PDF extractor using PyMuPDF. Returns structured lines with font metadata."""
import fitz
import re
import json
import os
from typing import List, Dict, Optional, Tuple


def extract_with_font_info(pdf_path: str) -> List[Dict]:
    """
    Extract text from PDF with font size, bold, and positional info per line.
    Returns list of dicts: { page, line_idx, text, font_size, is_bold, font_name, bbox }
    """
    doc = fitz.open(pdf_path)
    results = []
    for page_num, page in enumerate(doc):
        blocks = page.get_text("dict")["blocks"]
        for block in blocks:
            if block["type"] != 0:  # skip image blocks
                continue
            for line in block["lines"]:
                spans = line["spans"]
                if not spans:
                    continue
                text = "".join(span["text"] for span in spans).strip()
                if not text:
                    continue
                # Use the largest font in the line (heading often has larger)
                font_size = max(span["size"] for span in spans)
                is_bold = any(
                    "bold" in span["font"].lower() or "black" in span["font"].lower()
                    for span in spans
                )
                font_name = spans[0]["font"]
                bbox = line["bbox"]  # (x0, y0, x1, y1)
                results.append({
                    "page": page_num + 1,
                    "line_idx": len(results),
                    "text": text,
                    "font_size": round(font_size, 1),
                    "is_bold": is_bold,
                    "font_name": font_name,
                    "bbox": [round(v, 1) for v in bbox],
                })
    doc.close()
    return results


def compute_font_stats(lines: List[Dict]) -> Dict:
    """Compute font statistics to identify heading vs body font sizes."""
    if not lines:
        return {}
    sizes = [l["font_size"] for l in lines]
    from collections import Counter
    size_counts = Counter(sizes)
    most_common_size, most_common_count = size_counts.most_common(1)[0]

    bold_sizes = set(l["font_size"] for l in lines if l["is_bold"])
    body_sizes = {s for s, c in size_counts.items() if c > len(lines) * 0.05}
    heading_sizes = sorted(sizes, reverse=True)[:5]

    return {
        "total_lines": len(lines),
        "unique_sizes": sorted(set(sizes)),
        "most_common_size": most_common_size,
        "most_common_count": most_common_count,
        "body_sizes": sorted(body_sizes),
        "bold_sizes": sorted(bold_sizes),
        "heading_sizes": heading_sizes,
        "size_distribution": dict(size_counts.most_common(10)),
    }


def normalize_text(text: str) -> str:
    """Normalize common Indonesian legal doc formatting quirks."""
    t = re.sub(r'[ \t]+', ' ', text)
    t = re.sub(r'\.\.\.+$', '', t)
    t = re.sub(r'^\*+\s*', '', t)
    t = t.replace('\u2013', '-').replace('\u2014', '-').replace('\u2015', '-')
    t = t.strip()
    return t


def extract_lines_for_parsing(pdf_path: str) -> Tuple[List[Dict], Dict]:
    """Full extraction pipeline: extract + compute stats. Returns (lines, stats)."""
    raw_lines = extract_with_font_info(pdf_path)
    for line in raw_lines:
        line["normalized"] = normalize_text(line["text"])
    stats = compute_font_stats(raw_lines)
    return raw_lines, stats


if __name__ == "__main__":
    import sys
    path = sys.argv[1] if len(sys.argv) > 1 else "/home/z/my-project/download/samples/uud-1945/uud_1945.pdf"
    lines, stats = extract_lines_for_parsing(path)
    print(f"File: {os.path.basename(path)}")
    print(f"Total lines: {stats['total_lines']}")
    print(f"Unique font sizes: {stats['unique_sizes']}")
    print(f"Body sizes: {stats['body_sizes']}")
    print(f"Bold sizes: {stats['bold_sizes']}")
    print(f"Most common: {stats['most_common_size']} ({stats['most_common_count']} lines)")
    print(f"Size distribution (top 10): {json.dumps(stats['size_distribution'], indent=2)}")
    print(f"\nFirst 30 lines:")
    for l in lines[:30]:
        print(f"  p{l['page']:02d} fs={l['font_size']:5.1f} B={str(l['is_bold']):5s} | {l['normalized'][:80]}")
