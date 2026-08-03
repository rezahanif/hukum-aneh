r"""Extract full text skeleton from 1 sample per directory.Show line-by-line with font info so we can see the structural pattern per institution."""
import os, sys, json, re
from collections import Counter

sys.path.insert(0, '/home/z/my-project/scripts')
from extractor import extract_lines_for_parsing

SAMPLES_DIR = '/home/z/my-project/download/samples'
OUTPUT_DIR = '/home/z/my-project/download/skeletons'
os.makedirs(OUTPUT_DIR, exist_ok=True)

# Section anchor keywords by document type
SECTION_KEYWORDS = [
    'menimbang', 'mengingat', 'menetapkan', 'memutuskan',
    'amar', 'putusan', 'keputusan', 'instruksi',
    'bab ', 'bagian', 'paragraf', 'pasal',
    'ketentuan peralihan', 'ketentuan penutup', 'penutup',
    'dengan rahmat', 'presiden', 'menteri',
    'considering', 'decides', 'ruling',
    'penjelasan', 'elaan', 'dasar hukum',
    'pendahuluan', 'ruang lingkup', 'tujuan',
    'tugas', 'wewenang', 'kewajiban', 'larangan',
    'sanksi', 'pembinaan', 'pengawasan',
]


def extract_skeleton(folder, pdf_path):
    """Extract structured skeleton from a PDF."""
    lines, stats = extract_lines_for_parsing(pdf_path)
    
    result = {
        'folder': folder,
        'filename': os.path.basename(pdf_path),
        'total_lines': len(lines),
        'font_sizes': stats['unique_sizes'],
        'most_common_size': stats['most_common_size'],
        'sections': [],  # detected section boundaries
        'numbering_patterns': [],  # what numbering systems are used
        'key_anchors': [],  # lines that are structural markers
        'full_text_lines': [],  # all lines for reference
    }
    
    # Detect numbering patterns used in the document
    numbering_found = set()
    for l in lines:
        n = l['normalized']
        # Ayat style: (1), (2), (3)
        if re.match(r'^\s*\(\s*\d+\s*\)', n):
            numbering_found.add('ayat_(N)')
        # Letter style: a., b., c. or a), b), c)
        if re.match(r'^\s*[a-z]\)|^\s*[a-z]\.', n):
            numbering_found.add('letter_a.')
        # Number style: 1., 2., 3.
        if re.match(r'^\s*\d+\.', n) and not re.match(r'^\s*\d+\s*\.', n):
            numbering_found.add('number_1.')
        # Roman numeral pasal: Pasal I, Pasal IV
        if re.match(r'^\s*pasal\s+[ivxlcdmIVXLCDM]+', n, re.I):
            numbering_found.add('pasal_roman')
        # Arabic pasal: Pasal 1, Pasal 2
        if re.match(r'^\s*pasal\s+\d+', n, re.I):
            numbering_found.add('pasal_arabic')
        # BAB: BAB I, BAB 1
        if re.match(r'^\s*bab\s+', n, re.I):
            numbering_found.add('bab')
        # Angka style: 1), 2), 3)
        if re.match(r'^\s*\d+\)', n):
            numbering_found.add('number_1)')
    
    result['numbering_patterns'] = sorted(numbering_found)
    
    # Detect key structural anchors
    for li, l in enumerate(lines):
        n_lower = l['normalized'].lower().strip()
        for kw in SECTION_KEYWORDS:
            if n_lower == kw or n_lower.startswith(kw):
                result['key_anchors'].append({
                    'line_idx': li,
                    'page': l['page'],
                    'text': l['normalized'][:120],
                    'keyword': kw,
                    'font_size': l['font_size'],
                    'is_bold': l['is_bold'],
                })
                break
    
    # Full text (first 200 lines for skeleton)
    result['full_text_lines'] = [
        {
            'line_idx': l['line_idx'],
            'page': l['page'],
            'font_size': l['font_size'],
            'is_bold': l['is_bold'],
            'text': l['normalized'][:150],
        }
        for l in lines[:200]
    ]
    
    # If document is long, also get last 50 lines to see closing pattern
    if len(lines) > 200:
        result['closing_lines'] = [
            {
                'line_idx': l['line_idx'],
                'page': l['page'],
                'font_size': l['font_size'],
                'is_bold': l['is_bold'],
                'text': l['normalized'][:150],
            }
            for l in lines[-50:]
        ]
    
    return result


def find_samples():
    samples = {}
    for folder in sorted(os.listdir(SAMPLES_DIR)):
        folder_path = os.path.join(SAMPLES_DIR, folder)
        if not os.path.isdir(folder_path):
            continue
        pdfs = [f for f in os.listdir(folder_path) if f.lower().endswith('.pdf')]
        if pdfs:
            samples[folder] = os.path.join(folder_path, pdfs[0])
    return samples


def main():
    samples = find_samples()
    print(f'Found {len(samples)} folders to analyze')
    
    all_skeletons = []
    for folder, path in samples.items():
        print(f'\nProcessing: {folder} ({os.path.basename(path)})')
        try:
            skeleton = extract_skeleton(folder, path)
            all_skeletons.append(skeleton)
            
            out_path = os.path.join(OUTPUT_DIR, f'{folder}_skeleton.json')
            with open(out_path, 'w') as f:
                json.dump(skeleton, f, indent=2, ensure_ascii=False, default=str)
            print(f'  Lines: {skeleton["total_lines"]} | Font sizes: {skeleton["font_sizes"]}')
            print(f'  Numbering patterns: {skeleton["numbering_patterns"]}')
            print(f'  Key anchors ({len(skeleton["key_anchors"])}):')
            for a in skeleton['key_anchors'][:15]:
                print(f'    L{a["line_idx"]:4d} p{a["page"]:02d} fs={a["font_size"]:5.1f} B={str(a["is_bold"]):5s} [{a["keyword"]}] {a["text"][:80]}')
            if len(skeleton['key_anchors']) > 15:
                print(f'    ... and {len(skeleton["key_anchors"]) - 15} more')
        except Exception as e:
            print(f'  ERROR: {e}')
            import traceback
            traceback.print_exc()
    
    # Save combined
    combined_path = os.path.join(OUTPUT_DIR, 'all_skeletons.json')
    with open(combined_path, 'w') as f:
        json.dump(all_skeletons, f, indent=2, ensure_ascii=False, default=str)
    print(f'\nAll skeletons saved to {OUTPUT_DIR}/')
    print(f'Combined: {combined_path}')


if __name__ == '__main__':
    main()