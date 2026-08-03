"""
Extract full text skeleton from 1 sample PDF per directory.
Outputs raw text with structural markers (font size, bold, indentation) to JSON.
"""
import fitz
import json
import os
import re
from collections import Counter

SAMPLES_DIR = "/home/z/my-project/download/samples"
OUTPUT_PATH = "/home/z/my-project/download/skeletons_raw.json"

# Pick 1 file per directory
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


def extract_skeleton(pdf_path):
    """Extract full text with font/structural info, page by page."""
    doc = fitz.open(pdf_path)
    pages_data = []
    all_lines = []

    for page_num, page in enumerate(doc):
        blocks = page.get_text("dict")["blocks"]
        page_lines = []

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
                is_bold = any("bold" in s["font"].lower() or "black" in s["font"].lower() for s in spans)
                font_name = spans[0]["font"]
                bbox = line["bbox"]
                x0 = round(bbox[0], 1)

                page_lines.append({
                    "text": text,
                    "font_size": font_size,
                    "is_bold": is_bold,
                    "font_name": font_name,
                    "x_indent": x0,
                })
                all_lines.append({
                    "page": page_num + 1,
                    "text": text,
                    "font_size": font_size,
                    "is_bold": is_bold,
                    "font_name": font_name,
                    "x_indent": x0,
                })

        pages_data.append({
            "page_num": page_num + 1,
            "lines": page_lines,
        })

    doc.close()

    # Compute font stats
    sizes = [l["font_size"] for l in all_lines]
    size_counts = Counter(sizes)
    most_common_size, most_common_count = size_counts.most_common(1)[0]
    bold_sizes = sorted(set(l["font_size"] for l in all_lines if l["is_bold"]))

    # Find indent clusters
    indents = sorted(set(l["x_indent"] for l in all_lines if l["x_indent"] > 10))
    indent_clusters = []
    if indents:
        cluster_start = indents[0]
        for i in range(1, len(indents)):
            if indents[i] - indents[i-1] > 15:
                indent_clusters.append(round(cluster_start, 0))
                cluster_start = indents[i]
        indent_clusters.append(round(cluster_start, 0))

    stats = {
        "total_lines": len(all_lines),
        "total_pages": len(pages_data),
        "unique_font_sizes": sorted(set(sizes)),
        "size_distribution": dict(size_counts.most_common()),
        "most_common_size": most_common_size,
        "most_common_count": most_common_count,
        "bold_font_sizes": bold_sizes,
        "indent_clusters": indent_clusters,
    }

    return {"stats": stats, "pages": pages_data, "all_lines": all_lines}


def get_truncated_text(text, max_len=200):
    """Truncate text for display."""
    if len(text) <= max_len:
        return text
    return text[:max_len] + "..."


def main():
    results = {}

    for folder, rel_path in SAMPLES.items():
        full_path = os.path.join(SAMPLES_DIR, rel_path)
        if not os.path.exists(full_path):
            print(f"MISSING: {folder} -> {rel_path}")
            continue

        print(f"Processing: {folder}...")
        try:
            skeleton = extract_skeleton(full_path)
            results[folder] = {
                "file": rel_path,
                "stats": skeleton["stats"],
                "all_lines": skeleton["all_lines"],
            }
            print(f"  OK: {skeleton['stats']['total_lines']} lines, {skeleton['stats']['total_pages']} pages, "
                  f"fonts: {skeleton['stats']['unique_font_sizes']}")
        except Exception as e:
            print(f"  ERROR: {e}")
            results[folder] = {"file": rel_path, "error": str(e)}

    # Save raw JSON
    with open(OUTPUT_PATH, "w", encoding="utf-8") as f:
        json.dump(results, f, ensure_ascii=False, indent=2)
    print(f"\nSaved raw JSON to {OUTPUT_PATH}")
    print(f"Total folders processed: {len(results)}")

    # Print summary
    print("\n" + "="*80)
    print("SUMMARY PER FOLDER")
    print("="*80)
    for folder, data in results.items():
        if "error" in data:
            print(f"\n[{folder}] ERROR: {data['error']}")
            continue
        s = data["stats"]
        print(f"\n[{folder}] {data['file']}")
        print(f"  Pages: {s['total_pages']}, Lines: {s['total_lines']}")
        print(f"  Font sizes: {s['unique_font_sizes']}")
        print(f"  Most common: {s['most_common_size']} ({s['most_common_count']} lines, {s['most_common_count']*100//max(s['total_lines'],1)}%)")
        print(f"  Bold sizes: {s['bold_font_sizes']}")
        print(f"  Indent clusters: {s['indent_clusters']}")
        # Show first 15 lines
        print(f"  --- First 15 lines ---")
        for line in data["all_lines"][:15]:
            bold_mark = "[B]" if line["is_bold"] else "   "
            print(f"    p{line['page']:02d} fs={line['font_size']:5.1f} ind={line['x_indent']:6.1f} {bold_mark} | {get_truncated_text(line['text'])}")


if __name__ == "__main__":
    main()
