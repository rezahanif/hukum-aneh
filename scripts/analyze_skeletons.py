import os, io, sys, json, re, textwrap, traceback, fitz  # PyMuPDF is fitz
from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseDownload

CREDS_PATH = "/home/z/my-project/upload/token.json"
creds = Credentials.from_authorized_user_file(CREDS_PATH, ['https://www.googleapis.com/auth/drive'])
service = build('drive', 'v3', credentials=creds)

# Relevant law folders (excluding Putusan-MK as requested)
FOLDERS = {
    'peraturan': '1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U',
    'perpres': '1RBhkXEH750LSjipVjAFrC8C9rcu0fTCw',
    'perda': '1aYuZPB07M9eINHfprNd22LyLt9kxleGM',
    'Undang-Undang': '1vneHF9YxwgSnBh3ashORK0cYPo16vmQS',
    'uud-1945': '1gMTVbgvAx_Kwvq-JfY6zJXOwSGnlbQmF',
    'JDIH_Kemnaker': '1-Q4-ITbhuvqlME3atWsESMKe1zhafKx0',
    'JDIH_Kemenkeu': '1fY5m3VJBR08Fv1HREUfk1_SuhKyrSqb0',
    'JDIH_Kemendag': '1KytH2-EORzsclM-1Z3rN4XwfVMhCeaK7',
    'JDIH_Komdigi': '1o_VyLh8oPToi-yEv42PNWVFOnRZuGlwE',
    'JDIH_KPU': '14YjUGuu5sS42De5VqqadpGXgZz-6Z-iE',
}

OUT_DIR = "/home/z/my-project/download/skeleton_samples"
os.makedirs(OUT_DIR, exist_ok=True)

def download_one_pdf(folder_id, folder_name):
    """Download first PDF found in folder, return (filename, local_path)"""
    results = service.files().list(
        q=f"'{folder_id}' in parents and mimeType='application/pdf' and trashed=false",
        fields='files(id, name, size)',
        pageSize=5
    ).execute()
    files = results.get('files', [])
    if not files:
        print(f"  [WARN] No PDFs in {folder_name}")
        return None, None
    
    # Pick smallest file (faster download)
    f = sorted(files, key=lambda x: int(x.get('size', 999999999)))[0]
    fid, fname = f['id'], f['name']
    print(f"  Downloading: {fname} ({f.get('size','?')} bytes)")
    
    local_path = os.path.join(OUT_DIR, f"{folder_name}_{fname}")
    if os.path.exists(local_path) and os.path.getsize(local_path) > 1000:
        print(f"  Already cached: {local_path}")
        return fname, local_path
    
    request = service.files().get_media(fileId=fid)
    fh = io.FileIO(local_path, 'wb')
    downloader = MediaIoBaseDownload(fh, request)
    done = False
    while not done:
        status, done = downloader.next_chunk()
    fh.close()
    print(f"  Saved: {local_path}")
    return fname, local_path

def extract_skeleton(pdf_path, max_pages=5):
    """Extract text skeleton showing structural elements (headers, pasal, ayat, huruf, angka)"""
    doc = fitz.open(pdf_path)
    lines = []
    for page_num in range(min(len(doc), max_pages)):
        page = doc[page_num]
        text = page.get_text("text")
        for line in text.split('\n'):
            stripped = line.strip()
            if stripped:
                lines.append(stripped)
    doc.close()
    return lines

def analyze_structure(lines):
    """Identify structural patterns in the extracted lines"""
    patterns = {
        'judul': [], 'bab': [], 'bagian': [], 'pasal': [], 'ayat': [], 
        'huruf': [], 'angka': [], 'penjelasan': [], 'other_header': []
    }
    
    for i, line in enumerate(lines):
        # Judul / title
        if re.match(r'^(UNDANG-UNDANG|PERATURAN PRESIDEN|PERATURAN PEMERINTAH|PERATURAN DAERAH|KEPUTUSAN PRESIDEN|INSTRUKSI PRESIDEN|PERATURAN MENTERI|KEPUTUSAN MENTERI)', line, re.I):
            patterns['judul'].append((i, line))
        # BAB
        elif re.match(r'^BAB\s+(X{0,3}(IX|IV|V?I{0,3})|\d+)', line, re.I):
            patterns['bab'].append((i, line))
        # BAGIAN
        elif re.match(r'^BAGIAN\s+(KE-|KESATU|KEDUA|KETIGA|KEEMPAT|KELIMA|KEENAM|KETUJUH|KEDELAPAN|KESEMBILAN|KESEPULUH|\d+)', line, re.I):
            patterns['bagian'].append((i, line))
        # PASAL
        elif re.match(r'^Pasal\s+\d+', line, re.I):
            patterns['pasal'].append((i, line))
        # AYAT (numbered paragraphs: (1), (2), etc.)
        elif re.match('^\(\d+\)', line):
            patterns['ayat'].append((i, line))
        # HURUF (lettered: a., b., c. etc.)
        elif re.match('^[a-z]\.', line):
            patterns['huruf'].append((i, line))
        # ANGKA (numbered: 1., 2., 3. etc.)
        elif re.match('^\d+\.', line) and not re.match('^\d{4}', line):
            patterns['angka'].append((i, line))
        # PENJELASAN
        elif re.match('^PENJELASAN', line, re.I):
            patterns['penjelasan'].append((i, line))
        # Other potential headers (all caps short lines)
        elif len(line) < 80 and line == line.upper() and len(line) > 3 and not line.startswith('('):
            patterns['other_header'].append((i, line))
    
    return patterns

results = {}
for folder_name, folder_id in FOLDERS.items():
    print(f"\n{'='*60}")
    print(f"FOLDER: {folder_name}")
    print(f"{'='*60}")
    
    fname, local_path = download_one_pdf(folder_id, folder_name)
    if not local_path:
        results[folder_name] = {'error': 'No PDFs found'}
        continue
    
    try:
        lines = extract_skeleton(local_path, max_pages=8)
        patterns = analyze_structure(lines)
        
        results[folder_name] = {
            'filename': fname,
            'total_lines_in_sample': len(lines),
            'structure': {k: [(i, l[:120]) for i, l in v] for k, v in patterns.items()}
        }
        
        # Print first 80 lines to see raw skeleton
        print(f"FILE: {fname}")
        print(f"Total lines (first 8 pages): {len(lines)}")
        print(f"--- RAW TEXT (first 80 lines) ---")
        for i, line in enumerate(lines[:80]):
            print(f"  [{i:3d}] {line[:150]}")
        
        # Print pattern summary
        print(f"--- STRUCTURE SUMMARY ---")
        for k, v in patterns.items():
            if v:
                print(f"  {k}: {len(v)} matches")
                for idx, txt in v[:5]:
                    print(f"    [{idx:3d}] {txt[:100]}")
                if len(v) > 5:
                    print(f"    ... and {len(v)-5} more")
    except Exception as e:
        print(f"  ERROR: {e}")
        traceback.print_exc()
        results[folder_name] = {'error': str(e)}

# Save results
out_path = os.path.join(OUT_DIR, 'skeleton_analysis.json')
with open(out_path, 'w') as f:
    json.dump(results, f, indent=2, ensure_ascii=False)
print(f"\nResults saved to: {out_path}")
