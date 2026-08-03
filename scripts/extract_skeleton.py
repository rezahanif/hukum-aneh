#!/usr/bin/env python3
"""Extract clause/article-level skeleton from sample PDFs."""
import os, re, json, glob
import fitz  # PyMuPDF

SAMPLE_DIR = "/home/z/my-project/download/samples"

def extract_structure(pdf_path):
    """Extract headings, articles, clauses, chapters from a PDF."""
    doc = fitz.open(pdf_path)
    text = ""
    for page in doc:
        text += page.get_text()
    doc.close()
    
    lines = text.split('\n')
    structure = []
    
    # Pattern detection for Indonesian legal documents
    patterns = [
        ('BAB', r'^BAB\s+(?:[IVXLCDM]+|\d+)\b.*'),           # Chapter
        ('BAGIAN', r'^BAGIAN\s+(?:[IVXLCDM]+|\d+)\b.*'),   # Part
        ('PARAGRAPH', r'^PARAGRAF\s+(?:[IVXLCDM]+|\d+)\b.*'), # Paragraph (in some docs)
        ('PASAL', r'^Pasal\s+\d+.*'),                        # Article
        ('AYAT', r'^\(\d+\)\s*'),                           # Clause (ayat)
        ('HURUF', r'^[a-z]\)\s*'),                            # Letter point (huruf)
        ('ANGKA', r'^\d+\.\s*'),                             # Numbered point
    ]
    
    for line in lines:
        line_stripped = line.strip()
        if not line_stripped:
            continue
        for label, pattern in patterns:
            if re.match(pattern, line_stripped, re.IGNORECASE):
                structure.append((label, line_stripped[:120]))
                break
    
    return structure, text[:500]

def analyze_folder(folder_path):
    pdfs = sorted(glob.glob(os.path.join(folder_path, "*.pdf")))
    results = []
    for pdf in pdfs:
        try:
            structure, preview = extract_structure(pdf)
            results.append({
                'file': os.path.basename(pdf),
                'structure': structure,
                'preview': preview.replace('\n', ' ')[:300]
            })
        except Exception as e:
            results.append({'file': os.path.basename(pdf), 'error': str(e)})
    return results

all_results = {}
for folder in sorted(os.listdir(SAMPLE_DIR)):
    folder_path = os.path.join(SAMPLE_DIR, folder)
    if os.path.isdir(folder_path):
        all_results[folder] = analyze_folder(folder_path)

print(json.dumps(all_results, indent=2, ensure_ascii=False))
