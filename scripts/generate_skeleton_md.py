"""
Generate comprehensive markdown document with full skeleton text from each sample PDF.
"""
import json
import os
import re

INPUT_JSON = "/home/z/my-project/download/skeletons_raw.json"
OUTPUT_MD = "/home/z/my-project/download/dokumen_skeleton_per_direktori.md"

DOC_TYPES = {
    "uu": {"type": "Undang-Undang (Statute)", "issuer": "DPR + Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "pp": {"type": "Peraturan Pemerintah", "issuer": "Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "perppu": {"type": "Perppu (Emergency Regulation)", "issuer": "Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "perpres": {"type": "Peraturan Presiden", "issuer": "Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat / clauses"},
    "perda": {"type": "Peraturan Daerah (Regional Reg)", "issuer": "Kepala Daerah", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "keppres": {"type": "Keputusan Presiden (Decision)", "issuer": "Presiden", "hierarchy": "Consideranda > MEMUTUSKAN > Items"},
    "inpres": {"type": "Instruksi Presiden (Instruction)", "issuer": "Presiden", "hierarchy": "Consideranda > INSTRUKSI > Numbered"},
    "tap_mpr": {"type": "Ketetapan MPR (MPR Decree)", "issuer": "MPR", "hierarchy": "Pertimbangan > MEMUTUSKAN > Items"},
    "uud-1945": {"type": "UUD 1945 (Constitution)", "issuer": "BPUPKI/PPKI", "hierarchy": "BAB > Pasal > Ayat"},
    "Putusan-MK": {"type": "Putusan MK (Court Ruling)", "issuer": "Mahkamah Konstitusi", "hierarchy": "Menimbang > MENGADILI > MEMUTUSKAN > Amar"},
    "JDIH_Kemnaker": {"type": "Peraturan Menteri", "issuer": "Menteri Ketenagakerjaan", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "JDIH_Kemenkeu": {"type": "Peraturan Menteri", "issuer": "Menteri Keuangan", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "JDIH_Kemendag": {"type": "Keputusan Menteri (Decision)", "issuer": "Menteri Perdagangan", "hierarchy": "Consideranda > MENETAPKAN > Items + Tables"},
    "JDIH_Komdigi": {"type": "Peraturan Menteri", "issuer": "Menteri Kominfo", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "JDIH_KPU": {"type": "Peraturan KPU", "issuer": "Komisi Pemilihan Umum", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "peraturan": {"type": "Peraturan Pemerintah (simple)", "issuer": "Presiden", "hierarchy": "Pasal > Ayat"},
}


def classify_line(text):
    t = text.strip().upper()
    for m in ["MENIMBANG", "MENGINGAT", "DENGAN RAHMAT", "MENETAPKAN"]:
        if t.startswith(m):
            return f"[PREAMBLE:{m}]"
    for m in ["MEMUTUSKAN", "MENGADILI", "MENGINSTRUKSIKAN", "MENUGASKAN"]:
        if t.startswith(m):
            return f"[KEPUTUSAN:{m}]"
    for m in ["BAB ", "BAGIAN ", "PARAGRAF "]:
        if t.startswith(m) and len(text.strip()) < 100:
            return f"[HEADING:{m.strip()}]"
    if re.match(r'^\s*pasal\s+\d+', t):
        return "[PASAL]"
    if re.match(r'^\s*\(\s*\d+', text):
        return "[AYAT]"
    if re.match(r'^\s*\d+\.', text) and len(text.strip()) < 150:
        return "[ITEM]"
    if re.match(r'^\s*[a-z]\.', text) and len(text.strip()) < 150:
        return "[SUB-ITEM]"
    return ""


def generate_section(folder, data):
    if "error" in data:
        return f"## {folder}\n\n**ERROR**: {data['error']}\n\n---\n\n"

    meta = DOC_TYPES.get(folder, {})
    stats = data["stats"]
    lines = data["all_lines"]
    file_name = data["file"]

    out = []
    out.append(f"## {folder}")
    out.append("")
    out.append(f"- **File**: `{file_name}`")
    out.append(f"- **Document Type**: {meta.get('type', 'Unknown')}")
    out.append(f"- **Issued by**: {meta.get('issuer', 'Unknown')}")
    out.append(f"- **Pages**: {stats['total_pages']} | **Lines**: {stats['total_lines']}")
    out.append(f"- **Font sizes**: {stats['unique_font_sizes']}")
    out.append(f"- **Most common font**: {stats['most_common_size']} ({stats['most_common_count']*100//max(stats['total_lines'],1)}% of lines)")
    bold_sz = stats['bold_font_sizes'] if stats['bold_font_sizes'] else 'None'
    out.append(f"- **Bold font sizes**: {bold_sz}")
    out.append(f"- **Indent clusters**: {stats['indent_clusters']}")
    out.append(f"- **Expected hierarchy**: {meta.get('hierarchy', 'Unknown')}")
    out.append("")
    out.append("### Full Text Skeleton")
    out.append("")
    out.append("```python")

    clusters = stats['indent_clusters']
    most_common = stats['most_common_size']
    prev_page = 0

    for i, line in enumerate(lines):
        text = line["text"]
        page = line["page"]
        fs = line["font_size"]
        bold = line["is_bold"]
        x_ind = line["x_indent"]

        if page != prev_page:
            out.append(f"{'='*20} PAGE {page} {'='*20}")
            prev_page = page

        bold_mark = "B" if bold else " "
        fs_mark = "" if fs == most_common else f"F{fs:.0f}"

        indent_level = 0
        if clusters:
            for ci, c in enumerate(clusters):
                if abs(x_ind - c) < 25:
                    indent_level = ci + 1
                    break
        indent_mark = "" if indent_level == 0 else f" I{indent_level}"

        classification = classify_line(text)
        prefix = f"p{page:02d} {bold_mark} {fs_mark:4s}{indent_mark}"

        if classification:
            out.append(f"{prefix} {classification}")

        display_text = text[:250] + "..." if len(text) > 250 else text
        out.append(f"{prefix} | {display_text}")

    out.append("```")
    out.append("")
    out.append("---")
    out.append("")
    return "\n".join(out) + "\n"


def main():
    with open(INPUT_JSON, "r", encoding="utf-8") as f:
        data = json.load(f)

    parts = []
    parts.append("# Dokumen Skeleton: 1 File Per Direktori")
    parts.append("")
    parts.append("Full text skeleton (dengan anotasi struktural) dari 1 file sampel per direktori.")
    parts.append("Tujuan: mempelajari pola penulisan dan struktur dokumen **per institusi penerbit**.")
    parts.append("")
    parts.append("## Legenda Anotasi")
    parts.append("")
    parts.append("| Marker | Arti |")
    parts.append("|--------|------|")
    parts.append("| `B` | Bold text |")
    parts.append("| `Fn` | Font size berbeda dari body |")
    parts.append("| `I1`, `I2`, etc | Indent level |")
    parts.append("| `[PREAMBLE:...]` | Bagian pembuka |")
    parts.append("| `[KEPUTUSAN:...]` | Bagian keputusan/amar |")
    parts.append("| `[HEADING:...]` | Judul struktural (BAB, Bagian, etc) |")
    parts.append("| `[PASAL]` | Baris Pasal |")
    parts.append("| `[AYAT]` | Baris Ayat |")
    parts.append("| `[ITEM]` | Item bernomor |")
    parts.append("| `[SUB-ITEM]` | Sub-item berhuruf |")
    parts.append("")
    parts.append("## Ringkasan Tipe Dokumen")
    parts.append("")
    parts.append("| Direktori | Tipe | Penerbit |")
    parts.append("|-----------|------|----------|")
    for folder, meta in DOC_TYPES.items():
        parts.append(f"| `{folder}` | {meta['type']} | {meta['issuer']} |")
    parts.append("")
    parts.append("---")
    parts.append("")

    for folder in DOC_TYPES:
        if folder in data:
            print(f"Generating: {folder}")
            parts.append(generate_section(folder, data[folder]))
        else:
            parts.append(f"## {folder}\n\n*No data*\n\n---\n\n")

    content = "\n".join(parts)
    with open(OUTPUT_MD, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"\nSaved: {OUTPUT_MD}")
    print(f"Size: {os.path.getsize(OUTPUT_MD) / 1024:.1f} KB")


if __name__ == "__main__":
    main()
