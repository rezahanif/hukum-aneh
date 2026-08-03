import fitz, re, os, json, hashlib
from collections import Counter


def extract_with_fonts(pdf_path):
    """Extract PDF text with font size and bold metadata per line.
    
    Returns list of dicts:
      {text, font_size, is_bold, page, span_fonts}
    """
    doc = fitz.open(pdf_path)
    lines = []
    
    for page in doc:
        blocks = page.get_text("dict").get("blocks", [])
        for block in blocks:
            if "lines" not in block:
                continue
            for line in block["lines"]:
                spans = line.get("spans", [])
                if not spans:
                    continue
                text = "".join(s["text"] for s in spans)
                text = text.strip()
                if not text:
                    continue
                # Dominant font size
                sizes = [s["size"] for s in spans]
                font_size = round(Counter(sizes).most_common(1)[0][0], 1)
                # Check if any span is bold
                is_bold = any(
                    "bold" in s.get("font", "").lower()
                    or "demibold" in s.get("font", "").lower()
                    for s in spans
                )
                # Collect all font names
                font_names = list(set(s.get("font", "") for s in spans))
                lines.append({
                    "text": text,
                    "font_size": font_size,
                    "is_bold": is_bold,
                    "page": page.number,
                    "fonts": font_names,
                })
    doc.close()
    return lines


def detect_body_font_size(lines):
    """Detect the most common font size (body text)."""
    if not lines:
        return 12.0
    sizes = [l["font_size"] for l in lines]
    return round(Counter(sizes).most_common(1)[0][0], 1)


def normalize_text(text):
    """Clean common PDF extraction artifacts."""
    # Form feed -> newline
    text = re.sub(r"\f", "\n", text)
    # Collapse multiple spaces (but preserve newlines)
    text = re.sub(r"[ \t]+", " ", text)
    # Fix "Pasal 1 1" -> "Pasal 11" (broken numbers after Pasal)
    text = re.sub(r"(Pasal\s+)\d\s+(\d)", r"\1\2", text, flags=re.IGNORECASE)
    # Fix broken ayat: "(21" -> "(21)"
    text = re.sub(r"\((\d+)\s*$", r"(\1)", text, flags=re.MULTILINE)
    # Remove * prefix from Pasal/BAB (amendment markers in UUD)
    # Don't remove — keep as-is for now, parser handles it
    return text.strip()


def extract_raw_text(pdf_path):
    """Extract full raw text from PDF (fallback/simple mode)."""
    doc = fitz.open(pdf_path)
    text = ""
    for page in doc:
        text += page.get_text()
    doc.close()
    return normalize_text(text)


def file_hash(pdf_path):
    """Compute SHA256 hash of file for dedup."""
    h = hashlib.sha256()
    with open(pdf_path, "rb") as f:
        for chunk in iter(lambda: f.read(8192), b""):
            h.update(chunk)
    return h.hexdigest()
