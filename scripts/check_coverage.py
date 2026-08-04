"""Count and categorize all PDFs across law-related Drive folders."""

import json, re, os, io, time
from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build

CREDS_PATH = "/home/z/my-project/upload/token.json"
creds = Credentials.from_authorized_user_file(CREDS_PATH, ['https://www.googleapis.com/auth/drive'])
service = build('drive', 'v3', credentials=creds)

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
    'Putusan-MK': '18r2ZE9_R5zg72HToTBXwe3qH7Qr9Bidj',
}

def list_all_files(folder_id, folder_name):
    """List ALL files in a folder using pagination."""
    all_files = []
    page_token = None
    while True:
        results = service.files().list(
            q=f"'{folder_id}' in parents and trashed=false",
            fields='files(id, name, size, mimeType), nextPageToken',
            pageSize=1000,
            pageToken=page_token
        ).execute()
        files = results.get('files', [])
        all_files.extend(files)
        page_token = results.get('nextPageToken')
        if not page_token:
            break
        time.sleep(0.3)
    return all_files

def classify_filename(name):
    """Classify law type from filename."""
    name_lower = name.lower()
    
    # UUD
    if 'uud' in name_lower or 'undang-undang dasar' in name_lower or '1945' in name_lower:
        return 'UUD'
    
    # UU
    if re.search(r'uu[-_]?(no|nomor)?[-_]?', name_lower) or 'undang-undang' in name_lower:
        return 'UU'
    if re.search(r'\buu\b', name_lower):
        return 'UU'
    
    # PP
    if re.search(r'pp[-_]?(no|nomor)?[-_]?', name_lower) or 'peraturan-pemerintah' in name_lower:
        return 'PP'
    if 'peraturan pemerintah' in name_lower:
        return 'PP'
    
    # Perppu
    if 'perppu' in name_lower:
        return 'Perppu'
    
    # Perpres
    if 'perpres' in name_lower or 'peraturan-presiden' in name_lower or 'peraturan presiden' in name_lower:
        return 'Perpres'
    
    # Perda
    if 'perda' in name_lower or 'peraturan-daerah' in name_lower or 'peraturan daerah' in name_lower:
        return 'Perda'
    if 'peraturan daerah' in name_lower:
        return 'Perda'
    
    # PMK
    if 'pmk' in name_lower or 'peraturan menteri keuangan' in name_lower:
        return 'PMK'
    if 'permen' in name_lower or 'peraturan menteri' in name_lower:
        return 'Permen'
    
    # Kepmen
    if 'kepmen' in name_lower or 'keputusan menteri' in name_lower:
        return 'Kepmen'
    if 'keputusan' in name_lower and 'menteri' in name_lower:
        return 'Kepmen'
    
    # PKPU
    if 'pkpu' in name_lower or 'peraturan komisi pemilihan' in name_lower:
        return 'PKPU'
    if 'peraturan kp' in name_lower or 'kpu' in name_lower:
        return 'PKPU'
    
    # Keppres
    if 'keppres' in name_lower or 'keputusan presiden' in name_lower:
        return 'Keppres'
    
    # Inpres
    if 'inpres' in name_lower or 'instruksi presiden' in name_lower:
        return 'Inpres'
    
    # TAP MPR
    if 'tap mpr' in name_lower or 'ketetapan mpr' in name_lower:
        return 'TAP_MPR'
    
    # Perbup/Perwal
    if 'perbup' in name_lower or 'perwal' in name_lower:
        return 'Perbup/Perwal'
    
    return 'UNKNOWN'

def extract_law_number(name, law_type):
    """Try to extract law number and year from filename."""
    patterns = [
        r'(?:no|nomor)\s*(\d+)\s*(?:tahun|thn|th)\s*(\d{4})',
        r'(\d{4})\s*(?:no|nomor)?\s*(\d+)',
        r'_(\d+)\s*(?:tahun|thn)\s*(\d{4})',
        r'no[_.-]?(\d+)[_.-]?tahun[_.-]?(\d{4})',
    ]
    for pat in patterns:
        m = re.search(pat, name, re.IGNORECASE)
        if m:
            return int(m.group(1)), int(m.group(2))
    return None, None

results = {}

for folder_name, folder_id in FOLDERS.items():
    print(f"\n{'='*60}")
    print(f"FOLDER: {folder_name}")
    print(f"{'='*60}")
    
    files = list_all_files(folder_id, folder_name)
    print(f"  Total files: {len(files)}")
    
    # Separate PDFs from other files
    pdfs = [f for f in files if f.get('mimeType') == 'application/pdf']
    others = [f for f in files if f.get('mimeType') != 'application/pdf']
    subfolders = [f for f in files if f.get('mimeType') == 'application/vnd.google-apps.folder']
    
    print(f"  PDFs: {len(pdfs)}, Subfolders: {len(subfolders)}, Other: {len(others)}")
    
    if subfolders:
        print(f"  Subfolders: {[sf['name'] for sf in subfolders]}")
    
    # Classify PDFs
    type_counts = {}
    type_files = {}
    for f in pdfs:
        lt = classify_filename(f['name'])
        type_counts[lt] = type_counts.get(lt, 0) + 1
        if lt not in type_files:
            type_files[lt] = []
        num, year = extract_law_number(f['name'], lt)
        type_files[lt].append({
            'name': f['name'],
            'number': num,
            'year': year,
            'size': int(f.get('size', 0))
        })
    
    for lt, count in sorted(type_counts.items()):
        files_list = type_files[lt]
        years = sorted(set(f['year'] for f in files_list if f['year']))
        nums = sorted(set(f['number'] for f in files_list if f['number']))
        print(f"  {lt}: {count} files")
        print(f"    Years: {min(years) if years else '?'}-{max(years) if years else '?'} ({len(years)} years)")
        print(f"    Numbers: {len(nums)} unique")
        # Show a few samples
        for f in files_list[:3]:
            print(f"    - {f['name'][:100]}")
        if len(files_list) > 3:
            print(f"    ... and {len(files_list)-3} more")
    
    if others:
        print(f"  Non-PDF files: {[o['name'][:60] for o in others[:5]]}")
    
    results[folder_name] = {
        'total_files': len(files),
        'total_pdfs': len(pdfs),
        'subfolders': [sf['name'] for sf in subfolders],
        'type_counts': type_counts,
        'type_details': {lt: [{'name': f['name'], 'number': f['number'], 'year': f['year']} for f in flist] for lt, flist in type_files.items()}
    }

# Save
out_path = '/home/z/my-project/download/coverage_check.json'
with open(out_path, 'w') as f:
    json.dump(results, f, indent=2, ensure_ascii=False)
print(f"\nResults saved to: {out_path}")
