import os, io, sys, json, re, traceback, fitz
from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseDownload

CREDS_PATH = "/home/z/my-project/upload/token.json"
creds = Credentials.from_authorized_user_file(CREDS_PATH, ['https://www.googleapis.com/auth/drive'])
service = build('drive', 'v3', credentials=creds)

SAMPLES_DIR = "/home/z/my-project/download/skeleton_samples"

# For each folder, download a DIFFERENT sample — pick one that's a regular law (not province)
TARGET_FILES = {
    'peraturan': None,  # Will pick non-provinsi UU
    'perpres': None,
    'perda': None,
    'uud-1945': None,
    'JDIH_Kemnaker': None,
    'JDIH_Kemenkeu': None,
    'JDIH_Kemendag': None,
    'JDIH_Komdigi': None,
    'JDIH_KPU': None,
}

FOLDERS = {
    'peraturan': '1ewGhmNJ0Oszc9lo7eZMemOaPRGrsyM4U',
    'perpres': '1RBhkXEH750LSjipVjAFrC8C9rcu0fTCw',
    'perda': '1aYuZPB07M9eINHfprNd22LyLt9kxleGM',
    'uud-1945': '1gMTVbgvAx_Kwvq-JfY6zJXOwSGnlbQmF',
    'JDIH_Kemnaker': '1-Q4-ITbhuvqlME3atWsESMKe1zhafKx0',
    'JDIH_Kemenkeu': '1fY5m3VJBR08Fv1HREUfk1_SuhKyrSqb0',
    'JDIH_Kemendag': '1KytH2-EORzsclM-1Z3rN4XwfVMhCeaK7',
    'JDIH_Komdigi': '1o_VyLh8oPToi-yEv42PNWVFOnRZuGlwE',
    'JDIH_KPU': '14YjUGuu5sS42De5VqqadpGXgZz-6Z-iE',
}

def find_and_download(folder_id, folder_name, prefer_keyword=None, exclude_keyword=None):
    """Find a suitable PDF and download it"""
    results = service.files().list(
        q=f"'{folder_id}' in parents and mimeType='application/pdf' and trashed=false",
        fields='files(id, name, size)',
        pageSize=20
    ).execute()
    files = results.get('files', [])
    if not files:
        return None, None
    
    # Filter by preference
    if prefer_keyword:
        preferred = [f for f in files if prefer_keyword.lower() in f['name'].lower()]
        if preferred:
            files = preferred
    if exclude_keyword:
        files = [f for f in files if exclude_keyword.lower() not in f['name'].lower()]
    if not files:
        # Reset and pick smallest
        results2 = service.files().list(
            q=f"'{folder_id}' in parents and mimeType='application/pdf' and trashed=false",
            fields='files(id, name, size)',
            pageSize=20
        ).execute()
        files = results2.get('files', [])
    
    # Pick a medium-sized file (not too big, not too small)
    files.sort(key=lambda x: int(x.get('size', 0)))
    mid = len(files) // 3
    f = files[mid]
    
    fname = f['name']
    local_path = os.path.join(SAMPLES_DIR, f"{folder_name}_deep_{fname}")
    
    if os.path.exists(local_path) and os.path.getsize(local_path) > 1000:
        print(f"  Cached: {fname}")
        return fname, local_path
    
    print(f"  Downloading: {fname} ({f.get('size','?')} bytes)")
    request = service.files().get_media(fileId=f['id'])
    fh = io.FileIO(local_path, 'wb')
    downloader = MediaIoBaseDownload(fh, request)
    done = False
    while not done:
        status, done = downloader.next_chunk()
    fh.close()
    return fname, local_path

def full_extract(pdf_path):
    """Extract ALL text from PDF"""
    doc = fitz.open(pdf_path)
    all_text = []
    for page in doc:
        text = page.get_text("text")
        for line in text.split('\n'):
            stripped = line.strip()
            if stripped:
                all_text.append(stripped)
    doc.close()
    return all_text

def tag_structure(lines):
    """Tag each line with its structural role"""
    tagged = []
    for i, line in enumerate(lines):
        tag = 'BODY'
        
        if re.match(r'^(UNDANG-UNDANG DASAR|UNDANG-UNDANG REPUBLIK|PERATURAN PRESIDEN|PERATURAN PEMERINTAH|PERATURAN DAERAH|KEPUTUSAN PRESIDEN|INSTRUKSI PRESIDEN|PERATURAN MENTERI|KEPUTUSAN MENTERI|PERATURAN KOMISI)', line, re.I):
            tag = 'DOC_TYPE'
        elif re.match(r'^SALINAN$', line):
            tag = 'MARKER_SALINAN'
        elif re.match(r'^(PRESIDEN|BUPATI|WALIKOTA|GUBERNUR|MENTERI)', line, re.I):
            tag = 'SIGNATORY'
        elif re.match(r'^REPUBLIK\s+INDONESIA', line, re.I):
            tag = 'REPUBLIK'
        elif re.match(r'^NEGARA\s+REPUBLIK\s+INDONESIA', line, re.I):
            tag = 'NEGARA'
        elif re.match(r'^(NOMOR|NO\.)\s+\d+\s+TAHUN\s+\d+', line, re.I):
            tag = 'NOMOR_TAHUN'
        elif re.match(r'^DENGAN\s+RAHMAT', line, re.I):
            tag = 'OPENING'
        elif re.match(r'^BAB\s+', line, re.I):
            tag = 'BAB'
        elif re.match(r'^BAGIAN\s+', line, re.I):
            tag = 'BAGIAN'
        elif re.match(r'^PARAGRAF\s+', line, re.I):
            tag = 'PARAGRAF'
        elif re.match(r'^Pasal\s+\d+', line, re.I):
            tag = 'PASAL'
        elif re.match(r'^\(\d+\)', line):
            tag = 'AYAT'
        elif re.match(r'^[a-z]\.', line):
            tag = 'HURUF'
        elif re.match(r'^\d+\.', line) and not re.match(r'^\d{4}', line):
            tag = 'ANGKA'
        elif re.match(r'^(Menimbang|Mengingat|Memperhatikan|Dengan Persetujuan)$', line):
            tag = 'CONSIDERANS_LABEL'
        elif re.match(r'^(a\.\s+bahwa|b\.\s+bahwa|c\.\s+bahwa|d\.\s+bahwa|e\.\s+bahwa)', line, re.I):
            tag = 'CONSIDERANS_ITEM'
        elif re.match(r'^(1\.\s|2\.\s|3\.\s|4\.\s|5\.\s|6\.\s|7\.\s|8\.\s|9\.\s)\s*Undang-Undang', line):
            tag = 'DASAR_HUKUM'
        elif re.match(r'^PENJELASAN', line, re.I):
            tag = 'PENJELASAN_HEADER'
        elif re.match(r'^Disahkan\s+di', line, re.I):
            tag = 'PENUTUP_LOKASI'
        elif re.match(r'^pada\s+tanggal', line, re.I):
            tag = 'PENUTUP_TANGGAL'
        elif re.match(r'^Agar', line, re.I):
            tag = 'PENUTUP_AGAR'
        elif re.match(r'^TENTANG\s*$', line):
            tag = 'TENTANG'
        elif len(line) < 60 and line == line.upper() and len(line) > 3:
            tag = 'HEADER_CAPS'
        
        tagged.append((i, tag, line))
    return tagged

# Process each folder
all_results = {}

for folder_name, folder_id in FOLDERS.items():
    print(f"\n{'='*70}")
    print(f"FOLDER: {folder_name}")
    print(f"{'='*70}")
    
    prefer = None
    exclude = None
    if folder_name == 'peraturan':
        exclude = 'provinsi'  # avoid province-formation UU
    
    fname, local_path = find_and_download(folder_id, folder_name, prefer, exclude)
    if not local_path:
        all_results[folder_name] = {'error': 'No PDFs'}
        continue
    
    try:
        lines = full_extract(local_path)
        tagged = tag_structure(lines)
        
        # Count tags
        tag_counts = {}
        for _, tag, _ in tagged:
            tag_counts[tag] = tag_counts.get(tag, 0) + 1
        
        # Find PENJELASAN location
        penjelasan_start = None
        for i, tag, line in tagged:
            if tag == 'PENJELASAN_HEADER':
                penjelasan_start = i
                break
        
        # Find penutup (closing)
        penutup_start = None
        for i, tag, line in tagged:
            if tag in ('PENUTUP_AGAR', 'PENUTUP_LOKASI'):
                penutup_start = i
                break
        
        total_pages = fitz.open(local_path).page_count
        
        print(f"  File: {fname}")
        print(f"  Pages: {total_pages}, Lines: {len(lines)}")
        print(f"  Tag counts: {json.dumps(tag_counts)}")
        print(f"  PENJELASAN starts at line: {penjelasan_start}")
        print(f"  Penutup starts at line: {penutup_start}")
        
        # Print tagged skeleton (first 100 + around penjelasan + last 30)
        print(f"\n  --- SKELETON (first 100 tagged lines) ---")
        for idx, tag, line in tagged[:100]:
            if tag != 'BODY':
                print(f"    [{idx:4d}] [{tag:20s}] {line[:120]}")
            else:
                # Only show first 2 words of body for skeleton
                words = line.split()[:3]
                print(f"    [{idx:4d}] [{tag:20s}] {' '.join(words)}..." )
        
        if penjelasan_start:
            start = max(0, penjelasan_start - 2)
            end = min(len(tagged), penjelasan_start + 30)
            print(f"\n  --- AROUND PENJELASAN (lines {start}-{end}) ---")
            for idx, tag, line in tagged[start:end]:
                if tag != 'BODY':
                    print(f"    [{idx:4d}] [{tag:20s}] {line[:120]}")
                else:
                    words = line.split()[:3]
                    print(f"    [{idx:4d}] [{tag:20s}] {' '.join(words)}..." )
        
        # Print last 30 lines
        print(f"\n  --- LAST 30 lines ---")
        for idx, tag, line in tagged[-30:]:
            if tag != 'BODY':
                print(f"    [{idx:4d}] [{tag:20s}] {line[:120]}")
            else:
                words = line.split()[:3]
                print(f"    [{idx:4d}] [{tag:20s}] {' '.join(words)}..." )
        
        all_results[folder_name] = {
            'filename': fname,
            'pages': total_pages,
            'total_lines': len(lines),
            'tag_counts': tag_counts,
            'penjelasan_start': penjelasan_start,
            'penutup_start': penutup_start,
        }
        
    except Exception as e:
        print(f"  ERROR: {e}")
        traceback.print_exc()
        all_results[folder_name] = {'error': str(e)}

# Save
out_path = os.path.join(SAMPLES_DIR, 'skeleton_deep_analysis.json')
with open(out_path, 'w') as f:
    json.dump(all_results, f, indent=2, ensure_ascii=False)
print(f"\nDeep analysis saved to: {out_path}")
