import json
import os
import re
import sys
import fitz
from collections import defaultdict

SAMPLES_DIR = "/home/z/my-project/download/samples"
FOLDERS = {
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

# Common Indonesian legal words (for dictionary-match confidence scoring)
INDO_WORDS = set([
    "yang", "dan", "dengan", "untuk", "dalam", "pada", "adalah", "ini", "itu", "atau",
    "dari", "ke", "di", "sebuah", "sebagai", "telah", "akan", "dapat", "oleh", "tidak",
    "atas", "antara", "juga", "serta", "bagi", "sementara", "setiap", "tentang",
    "nomor", "tahun", "republik", "indonesia", "presiden", "menteri", "undang",
    "peraturan", "pemerintah", "daerah", "pasal", "ayat", "bab", "bagian",
    "menimbang", "mengingat", "menetapkan", "memutuskan", "dengan", "raamat",
    "ketentuan", "sebagaimana", "dimaksud", "huruf", "ayat", " Pasal",
    "negara", "warga", "hak", "kewajiban", "hukum", "pengadilan", "mahkamah",
    "konstitusi", "putusan", "keputusan", "instruksi", "ketetapan",
    "pelaksanaan", "tata", "cara", "usaha", "pajak", "bea", "cukai",
    "impor", "ekspor", "perdagangan", "keuangan", "ketenagakerjaan",
    "kesehatan", "pendidikan", "lingkungan", "pertanian", "perikanan",
    "karena", "apabila", "tersebut", "berdasarkan", "melalui", "sebelum",
    "sesudah", "selain", "baik", "maka", "agar", "supaya", "agar",
    "diperlukan", "ditetapkan", "ditunjuk", "dibentuk", "diatur",
    "seharusnya", "sebagaimana", "sebelumnya", "selanjutnya",
    "dengan", "mempunyai", "memiliki", "melakukan", "memberikan",
])

# Known corruption patterns from our earlier analysis
CORRUPTION_PATTERNS = [
    (r'PRESIOEN', 'PRESIDEN'),
    (r'REPLJBLIK', 'REPUBLIK'),
    (r'2OO9', '2009'),
    (r'2O2O', '2020'),
    (r'2O24', '2024'),
    (r'2O25', '2025'),
    (r'RAAMAT', 'RAHMAT'),
    (r'\bl\.', '1.'),  # lowercase L as digit 1
    (r'\bO\b', '0'),  # letter O as digit 0 (in numeric contexts)
]


def extract_raw_text(pdf_path):
    """Extract raw text from PDF using PyMuPDF."""
    doc = fitz.open(pdf_path)
    full_text = ""
    pages_text = []
    for page_num, page in enumerate(doc):
        text = page.get_text("text")
        pages_text.append(text)
        full_text += text
    doc.close()
    return full_text, pages_text


def check_font_cmap(pdf_path):
    """Check for embedded font CMap issues."""
    doc = fitz.open(pdf_path)
    font_issues = []
    for page_num, page in enumerate(doc):
        try:
            fonts = page.get_fonts()
            for fname, ftype, ext, _, _, encoding in fonts:
                if ftype == "Type1" or encoding not in ("WinAnsiEncoding", "MacRomanEncoding", "Identity-H", "UniJIS-UCS2-H"):
                    font_issues.append({
                        "page": page_num + 1,
                        "font": fname,
                        "type": ftype,
                        "encoding": encoding,
                    })
        except Exception as e:
            font_issues.append({"page": page_num + 1, "error": str(e)})
    doc.close()
    return font_issues


def compute_confidence_score(text):
    """Compute ratio of dictionary-word matches vs total tokens."""
    # Tokenize: split on whitespace and punctuation
    tokens = re.findall(r'[a-zA-Z]+', text.lower())
    if not tokens:
        return 0.0, 0, 0
    
    matches = sum(1 for t in tokens if t in INDO_WORDS)
    # Also check for common English legal words
    eng_legal = {"the", "of", "and", "in", "to", "for", "is", "by", "with", "from"}
    matches += sum(1 for t in tokens if t in eng_legal)
    
    return matches / len(tokens), matches, len(tokens)


def detect_corruption_patterns(text):
    """Find known corruption patterns in text."""
    findings = []
    for pattern, correction in CORRUPTION_PATTERNS:
        matches = re.findall(pattern, text)
        if matches:
            findings.append({
                "pattern": pattern,
                "correction": correction,
                "count": len(matches),
                "samples": matches[:3],
            })
    return findings


def detect_header_footer_boilerplate(pages_text):
    """Detect lines that repeat across pages (header/footer)."""
    line_pages = defaultdict(list)
    for page_idx, page_text in enumerate(pages_text):
        lines = page_text.strip().split('\n')
        for line in lines:
            stripped = line.strip()
            if len(stripped) > 3 and len(stripped) < 100:
                line_pages[stripped].append(page_idx)
    
    total_pages = len(pages_text)
    boilerplate = []
    for text, pages in line_pages.items():
        if len(pages) >= max(3, total_pages * 0.4):
            boilerplate.append({
                "text": text,
                "appears_on_pages": len(pages),
                "total_pages": total_pages,
                "ratio": round(len(pages) / total_pages, 2),
            })
    
    return sorted(boilerplate, key=lambda x: -x["ratio"])


def check_multi_column(pdf_path):
    """Heuristic: check if blocks on a page suggest multi-column layout."""
    doc = fitz.open(pdf_path)
    multi_col_pages = []
    
    for page_num, page in enumerate(doc):
        blocks = page.get_text("blocks")
        if len(blocks) < 3:
            continue
        
        # Get x-centers of text blocks
        x_centers = sorted(set(round((b[0] + b[2]) / 2) for b in blocks if b[6] == 0))
        
        # If there are 2+ distinct x-center clusters, likely multi-column
        if len(x_centers) >= 2:
            gaps = []
            for i in range(1, len(x_centers)):
                gaps.append(x_centers[i] - x_centers[i-1])
            
            page_width = page.rect.width
            # Check if there's a significant gap that's not just margin
            big_gaps = [g for g in gaps if g > page_width * 0.1]
            if big_gaps:
                multi_col_pages.append({
                    "page": page_num + 1,
                    "x_centers": x_centers,
                    "gaps": gaps,
                    "page_width": round(page_width, 1),
                })
    
    doc.close()
    return multi_col_pages


def diagnose_pdf(folder, rel_path):
    """Run all diagnostics on a single PDF."""
    full_path = os.path.join(SAMPLES_DIR, rel_path)
    if not os.path.exists(full_path):
        return {"folder": folder, "error": "File not found", "file": rel_path}
    
    try:
        full_text, pages_text = extract_raw_text(full_path)
    except Exception as e:
        return {"folder": folder, "error": str(e), "file": rel_path}
    
    # Encoding / corruption
    confidence, dict_matches, total_tokens = compute_confidence_score(full_text)
    corruption = detect_corruption_patterns(full_text)
    font_issues = check_font_cmap(full_path)
    
    # Reading order
    boilerplate = detect_header_footer_boilerplate(pages_text)
    multi_col = check_multi_column(full_path)
    
    return {
        "folder": folder,
        "file": rel_path,
        "extraction_method": "text_layer",
        "confidence_score": round(confidence, 4),
        "dict_matches": dict_matches,
        "total_tokens": total_tokens,
        "needs_ocr": confidence < 0.90,
        "corruption_patterns": corruption,
        "font_cmap_issues": font_issues,
        "header_footer_boilerplate": boilerplate,
        "multi_column_pages": multi_col,
    }


def main():
    results = []
    for folder, rel_path in FOLDERS.items():
        print(f"Diagnosing: {folder}...", end=" ")
        result = diagnose_pdf(folder, rel_path)
        results.append(result)
        if "error" in result:
            print(f"ERROR: {result['error']}")
        else:
            ocr_flag = " [NEEDS OCR]" if result["needs_ocr"] else ""
            print(f"conf={result['confidence_score']:.2%}{ocr_flag}, corrupt={len(result['corruption_patterns'])}, boilerplate={len(result['header_footer_boilerplate'])}, multicols={len(result['multi_column_pages'])}")
    
    output_path = "/home/z/my-project/download/phase0_diagnosis.json"
    with open(output_path, "w", encoding="utf-8") as f:
        json.dump(results, f, ensure_ascii=False, indent=2, default=str)
    
    # Print summary
    print("\n" + "="*80)
    print("PHASE 0 DIAGNOSIS SUMMARY")
    print("="*80)
    
    needs_ocr = [r for r in results if r.get("needs_ocr")]
    has_corruption = [r for r in results if r.get("corruption_patterns")]
    has_multicol = [r for r in results if r.get("multi_column_pages")]
    has_font_issues = [r for r in results if r.get("font_cmap_issues")]
    
    print(f"\nFiles needing OCR (conf < 90%): {len(needs_ocr)}")
    for r in needs_ocr:
        print(f"  - {r['folder']}: {r['confidence_score']:.2%} ({r['dict_matches']}/{r['total_tokens']} tokens)")
    
    print(f"\nFiles with corruption patterns: {len(has_corruption)}")
    for r in has_corruption:
        for c in r["corruption_patterns"]:
            print(f"  - {r['folder']}: '{c['pattern']}' -> '{c['correction']}' ({c['count']}x, e.g. {c['samples']})")
    
    print(f"\nFiles with header/footer boilerplate: {len([r for r in results if r.get('header_footer_boilerplate')])}")
    for r in results:
        if r.get("header_footer_boilerplate"):
            for b in r["header_footer_boilerplate"][:3]:
                print(f"  - {r['folder']}: \"{b['text'][:60]}\" on {b['appears_on_pages']}/{b['total_pages']} pages ({b['ratio']:.0%})")
    
    print(f"\nFiles with multi-column pages: {len(has_multicol)}")
    for r in has_multicol:
        for mc in r["multi_column_pages"][:3]:
            print(f"  - {r['folder']} p{mc['page']}: x_centers={mc['x_centers']}, gaps={mc['gaps']}")
    
    print(f"\nFiles with font CMap issues: {len(has_font_issues)}")
    for r in has_font_issues:
        for fi in r["font_cmap_issues"][:3]:
            print(f"  - {r['folder']} p{fi['page']}: font={fi.get('font','?')}, type={fi.get('type','?')}, enc={fi.get('encoding','?')}")
    
    print(f"\nSaved to: {output_path}")


if __name__ == "__main__":
    main()
