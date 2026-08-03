)
"""
Generate comprehensive markdown document with full skeleton text from each sample PDF.
Annotated with structural markers to reveal each institution's document pattern.
"""
import json
import os

INPUT_JSON = "/home/z/my-project/download/skeletons_raw.json"
OUTPUT_MD = "/home/z/my-project/download/dokumen_skeleton_per_direktori.md"

# Document type classification
DOC_TYPES = {
    "uu": {"type": "Undang-Undang (Statute)", "issuer": "DPR + Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "pp": {"type": "Peraturan Pemerintah (Government Regulation)", "issuer": "Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "perppu": {"type": "Peraturan Pemerintah Pengganti UU (Emergency Regulation)", "issuer": "Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "perpres": {"type": "Peraturan Presiden (Presidential Regulation)", "issuer": "Presiden", "hierarchy": "BAB > Bagian > Pasal > Ayat / Numbered clauses"},
    "perda": {"type": "Peraturan Daerah (Regional Regulation)", "issuer": "Kepala Daerah", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "keppres": {"type": "Keputusan Presiden (Presidential Decision)", "issuer": "Presiden", "hierarchy": "Consideranda > MEMUTUSKAN > Decision items"},
    "inpres": {"type": "Instruksi Presiden (Presidential Instruction)", "issuer": "Presiden", "hierarchy": "Consideranda > INSTRUKSI > Numbered instructions"},
    "tap_mpr": {"type": "Ketetapan MPR (MPR Decree)", "issuer": "MPR", "hierarchy": "Pertimbangan > MEMUTUSKAN > Decision items"},
    "uud-1945": {"type": "Undang-Undang Dasar 1945 (Constitution)", "issuer": "BPUPKI/PPKI", "hierarchy": "BAB > Pasal > Ayat"},
    "Putusan-MK": {"type": "Putusan Mahkamah Konstitusi (Constitutional Court Ruling)", "issuer": "MK", "hierarchy": "Menimbang > Mengadili > MEMUTUSKAN > Amar"},
    "JDIH_Kemnaker": {"type": "Peraturan Menteri (Ministerial Regulation)", "issuer": "Menteri Ketenagakerjaan", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "JDIH_Kemenkeu": {"type": "Peraturan Menteri (Ministerial Regulation)", "issuer": "Menteri Keuangan", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "JDIH_Kemendag": {"type": "Keputusan Menteri (Ministerial Decision)", "issuer": "Menteri Perdagangan", "hierarchy": "Consideranda > MENETAPKAN > Decision items + Tables"},
    "JDIH_Komdigi": {"type": "Peraturan Menteri (Ministerial Regulation)", "issuer": "Menteri Kominfo", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "JDIH_KPU": {"type": "Peraturan KPU (KPU Regulation)", "issuer": "Komisi Pemilihan Umum", "hierarchy": "BAB > Bagian > Pasal > Ayat"},
    "peraturan": {"type": "Peraturan Pemerintah (Government Regulation)", "issuer": "Presiden", "hierarchy": "Pasal > Ayat (simple/short)"},
}


def escape_md(text):
    """Escape markdown special chars but preserve readability."""
    # Don't escape # since we use them for headers
    t = text.replace('|', '\|')
    t = t.replace('*', '\*')
    return t


def get_indent_marker(x_indent, clusters):
    """Return visual indent marker based on position."""
    if not clusters or x_indent <= 50:
        return ""
    # Find which cluster this is closest to
    for i, c in enumerate(clusters):
        if abs(x_indent - c) < 20:
            return f"  [{'I'*(i+1)}]"
    return ""


def classify_line(text, is_bold, font_size, most_common_size):
    """Classify line type for annotation."""
    t = text.strip().upper()
    
    # Structural keywords
    preamble_markers = ["MENIMBANG", "MENGINGAT", "DENGAN RAHMAT", "MENETAPKAN"]
    heading_markers = ["BAB ", "BAGIAN ", "PARAGRAF ", "PASAL ", "UNTUK "]
    decision_markers = ["MEMUTUSKAN", "MENGADILI", "MENETAPKAN", "MENGINTRUKSIKAN", "MENUGASKAN"]
    
    for m in preamble_markers:
        if t.startswith(m):
            return f"**[PREAMBLE:{m}]**"
    
    for m in decision_markers:
        if t.startswith(m):
            return f"**[KEPUTUSAN:{m}]**"
    
    for m in heading_markers:
        if t.startswith(m) and len(text.strip()) < 80:
            return f"**[HEADING:{m.strip()}]**"
    
    # Check for ayat pattern
    import re
    if re.match(r'^\s*\(\s*\d+\s*\)', text):
        return "**[AYAT]**"
    
    # Check for numbered items (1., 2., etc.)
    if re.match(r'^\s*\d+\.', text) and len(text.strip()) < 120:
        return "**[ITEM]**"
    
    # Check for lettered items (a., b., etc.)
    if re.match(r'^\s*[a-z]\.', text) and len(text.strip()) < 120:
        return "**[SUB-ITEM]**"
    
    return ""


def generate_section(folder, data):
    """Generate markdown section for one folder."""
    if "error" in data:
        return f"## {folder}\n\n**ERROR**: {data['error']}\n\n"
    
    meta = DOC_TYPES.get(folder, {})
    stats = data["stats"]
    lines = data["all_lines"]
    file = data["file"]
    
    md = f"## {folder}\n\n"
    md += f"- **File**: `{file}`\n"
    md += f"- **Document Type**: {meta.get('type', 'Unknown')}\n"
    md += f"- **Issued by**: {meta.get('issuer', 'Unknown')}\n"
    md += f"- **Pages**: {stats['total_pages']} | **Lines**: {stats['total_lines']}\n"
    md += f"- **Font sizes**: {stats['unique_font_sizes']}\n"
    md += f"- **Most common font**: {stats['most_common_size']} ({stats['most_common_count']*100//max(stats['total_lines'],1)}% of lines)\n"
    md += f"- **Bold font sizes**: {stats['bold_font_sizes'] if stats['bold_font_sizes'] else 'None detected'}\n"
    md += f"- **Indent clusters**: {stats['indent_clusters']}\n"
    md += f"- **Expected hierarchy**: {meta.get('hierarchy', 'Unknown')}\n\n"
    
    md += f"### Full Text Skeleton\n\n"
    md += "```
"
    
    clusters = stats['indent_clusters']
    most_common = stats['most_common_size']
    
    for line in lines:
        text = line["text"]
        page = line["page"]
        fs = line["font_size"]
        bold = line["is_bold"]
        x_ind = line["x_indent"]
        
        # Build prefix annotation
        bold_mark = "B" if bold else " "
        indent_mark = get_indent_marker(x_ind, clusters)
        fs_mark = "" if fs == most_common else f"F{fs:.0f}"
        
        # Classify
        classification = classify_line(text, bold, fs, most_common)
        
        # Page break indicator
        if line == lines[0] or (line != lines[0] and lines[lines.index(line)-1]["page"] != page):
            md += f"\n{'='*60} PAGE {page} {'='*60}\n"
        
        # Format the line
        prefix = f"p{page:02d} {bold_mark} {fs_mark:4s} {indent_mark}"
        
        if classification:
            md += f"{prefix} {classification}\n"
        
        # Truncate very long lines for readability
        display_text = text[:200] + "..." if len(text) > 200 else text
        md += f"{prefix} | {display_text}\n"
    
    md += "```\n\n"
    md += "---\n\n"
    return md


def main():
    with open(INPUT_JSON, "r", encoding="utf-8") as f:
        data = json.load(f)
    
    md = "# Dokumen Skeleton: 1 File Per Direktori\n\n"
    md += "Dokumen ini berisi full text skeleton (dengan anotasi struktural) dari 1 file sampel per direktori.\n"
    md += "Tujuannya untuk mempelajari pola penulisan dan struktur dokumen **per institusi penerbit**.\n\n"
    md += "## Legenda Anotasi\n\n"
    md += "| Marker | Arti |\n"
    md += "|--------|------||\n"
    md += "| `B` | Bold text |\n"
    md += "| `Fn` | Font size berbeda dari yang paling umum |\n"
    md += "| `[I]`, `[II]`, etc. | Indent level berdasarkan posisi x |\n"
    md += "| `[PREAMBLE:...]` | Bagian pembuka (Menimbang, Mengingat, dll) |\n"
    md += "| `[KEPUTUSAN:...]` | Bagian keputusan/amar (MEMUTUSKAN, MENETAPKAN, dll) |\n"
    md += "| `[HEADING:...]` | Judul struktural (BAB, Bagian, Pasal, dll) |\n"
    md += "| `[AYAT]` | Baris ayat dimulai dengan angka dalam kurung |\n"
    md += "| `[ITEM]` | Item bernomor |\n"
    md += "| `[SUB-ITEM]` | Sub-item berhuruf |\n\n"
    md += "## Ringkasan Tipe Dokumen\n\n"
    md += "| Direktori | Tipe Dokumen | Penerbit |\n"
    md += "|-----------|-------------|---------|\n"
    for folder, meta in DOC_TYPES.items():
        md += f"| `{folder}` | {meta['type']} | {meta['issuer']} |\n"
    md += "\n---\n\n"
    
    # Generate each section
    for folder in DOC_TYPES:
        if folder in data:
            print(f"Generating section: {folder}")
            md += generate_section(folder, data[folder])
        else:
            md += f"## {folder}\n\n*No data available*\n\n---\n\n"
    
    with open(OUTPUT_MD, "w", encoding="utf-8") as f:
        f.write(md)
    
    print(f"\nSaved to: {OUTPUT_MD}")
    print(f"Total size: {os.path.getsize(OUTPUT_MD) / 1024:.1f} KB")


if __name__ == "__main__":
    main()
