"""Deeper analysis: fix number/year extraction, count unique laws, detect dupes.""" 

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
    'uud-1945': '1gMTVbgvAx_Kwvq-JfY6zJXOwSGnlbQmF',
    'JDIH_Kemnaker': '1-Q4-ITbhuvqlME3atWsESMKe1zhafKx0',
    'JDIH_Kemenkeu': '1fY5m3VJBR08Fv1HREUfk1_SuhKyrSqb0',
    'JDIH_Kemendag': '1KytH2-EORzsclM-1Z3rN4XwfVMhCeaK7',
    'JDIH_Komdigi': '1o_VyLh8oPToi-yEv42PNWVFOnRZuGlwE',
    'JDIH_KPU': '14YjUGuu5sS42De5VqqadpGXgZz-6Z-iE',
    'Putusan-MK': '18r2ZE9_R5zg72HToTBXwe3qH7Qr9Bidj',
}

def list_all_files(folder_id):
    all_files = []
    page_token = None
    while True:
        results = service.files().list(
            q=f"'{folder_id}' in parents and mimeType='application/pdf' and trashed=false",
            fields='files(id, name, size), nextPageToken',
            pageSize=1000,
            pageToken=page_token
        ).execute()
        all_files.extend(results.get('files', []))
        page_token = results.get('nextPageToken')
        if not page_token:
            break
        time.sleep(0.3)
    return all_files

def extract_num_year(name):
    """Robust extraction of law number and year from any filename pattern."""
    # Pattern: No__123_Tahun_2025
    m = re.search(r'No[_]*(?:__)?(\d+)[_]+Tahun[_]+(\d{4})', name, re.I)
    if m: return int(m.group(1)), int(m.group(2))
    
    # Pattern: no-13-tahun-2023
    m = re.search(r'no[-_](\d+)[-_]tahun[-_](\d{4})', name, re.I)
    if m: return int(m.group(1)), int(m.group(2))
    
    # Pattern: No. 123 Tahun 2025
    m = re.search(r'No\.?\s*(\d+)\s*Tahun\s*(\d{4})', name, re.I)
    if m: return int(m.group(1)), int(m.group(2))
    
    # Pattern: PKPU_9_2026 (number_year)
    m = re.search(r'(?:PKPU|PKP|pmk|PMK)[_.-]?(\d+)[_.-](\d{4})', name, re.I)
    if m: return int(m.group(1)), int(m.group(2))
    
    # Pattern: putusan_mkri_5301
    m = re.search(r'putusan_mkri[_.-]?(\d+)', name, re.I)
    if m: return int(m.group(1)), None
    
    # Pattern: _2025 at end
    m = re.search(r'[_-](\d{4})\.pdf$', name, re.I)
    if m: return None, int(m.group(1))
    
    # Pattern: just 4-digit year anywhere
    m = re.search(r'(?:tahun|thn|th)[_. -]*(\d{4})', name, re.I)
    if m: return None, int(m.group(1))
    
    # Pattern: number before .pdf
    m = re.search(r'[_ -](\d{3,5})\.pdf$', name)
    if m: return int(m.group(1)), None
    
    return None, None

def classify_file(name, folder):
    """Classify law type more accurately using both name and folder context."""
    nl = name.lower()
    
    # Putusan MK
    if 'putusan' in nl and ('mkri' in nl or 'mk' in nl):
        return 'Putusan_MK'
    
    # UUD
    if 'uud' in nl and '1945' in nl:
        return 'UUD_1945'
    
    # Perppu
    if 'perppu' in nl:
        return 'Perppu'
    
    # PP  
    if 'pp-no-' in nl or 'pp_' in nl or 'peraturan pemerintah' in nl:
        return 'PP'
    
    # UU
    if 'uu-no-' in nl or 'uu_' in nl:
        return 'UU'
    
    # Perpres
    if 'perpres' in nl:
        return 'Perpres'
    
    # Perda
    if 'perda' in nl:
        return 'Perda'
    
    # PMK
    if 'pmk' in nl:
        return 'PMK'
    
    # Permenaker
    if 'permenaker' in nl:
        return 'Permenaker'
    
    # Permenkominfo / Kepmenkominfo
    if 'permenkominfo' in nl:
        return 'Permenkominfo'
    if 'kepmenkominfo' in nl:
        return 'Kepmenkominfo'
    
    # Permendag / Kepmendag
    if 'permendag' in nl:
        return 'Permendag'
    if 'kepmendag' in nl:
        return 'Kepmendag'
    
    # PKPU
    if 'pkpu' in nl:
        return 'PKPU'
    
    # Folder-based classification as fallback
    if folder == 'JDIH_Kemnaker':
        if 'permen' in nl: return 'Permenaker'
        if 'kepmen' in nl: return 'Kepmenaker'
        if 'pp' in nl: return 'PP'
        return 'Other_Kemnaker'
    
    if folder == 'JDIH_Kemenkeu':
        if 'pmk' in nl: return 'PMK'
        if 'permen' in nl: return 'PMK'  # PMK is Permen Keuangan
        if 'kepmen' in nl: return 'Kepmenkeu'
        return 'Other_Kemenkeu'
    
    if folder == 'JDIH_Kemendag':
        if 'kepmen' in nl: return 'Kepmendag'
        if 'permen' in nl: return 'Permendag'
        return 'Other_Kemendag'
    
    if folder == 'JDIH_Komdigi':
        if 'permen' in nl: return 'Permenkominfo'
        if 'kepmen' in nl: return 'Kepmenkominfo'
        return 'Other_Komdigi'
    
    if folder == 'JDIH_KPU':
        return 'PKPU'
    
    if folder == 'peraturan':
        if 'perppu' in nl: return 'Perppu'
        if 'pp-no-' in nl: return 'PP'
        return 'UU'
    
    if folder == 'perpres':
        return 'Perpres'
    
    if folder == 'perda':
        return 'Perda'
    
    return 'UNKNOWN'

all_results = {}

for folder_name, folder_id in FOLDERS.items():
    print(f"\n{'='*60}")
    print(f"FOLDER: {folder_name}")
    print(f"{'='*60}")
    
    files = list_all_files(folder_id)
    print(f"  Total PDFs: {len(files)}")
    
    type_data = {}
    for f in files:
        lt = classify_file(f['name'], folder_name)
        num, year = extract_num_year(f['name'])
        
        if lt not in type_data:
            type_data[lt] = {'files': [], 'unique_keys': set()}
        
        key = f"{num}_{year}" if num and year else (f"num={num}_year={year}" if num or year else f['name'][:80])
        type_data[lt]['files'].append({
            'name': f['name'],
            'number': num,
            'year': year,
            'size': int(f.get('size', 0)),
            'id': f['id']
        })
        type_data[lt]['unique_keys'].add(key)
    
    folder_summary = {}
    for lt, data in sorted(type_data.items()):
        flist = data['files']
        unique_count = len(data['unique_keys'])
        
        years_with_data = [f['year'] for f in flist if f['year'] and 1900 <= f['year'] <= 2030]
        years_range = f"{min(years_with_data)}-{max(years_with_data)}" if years_with_data else "?"
        num_years = len(set(years_with_data)) if years_with_data else 0
        
        # Detect duplicates (same number+year)
        num_year_pairs = [(f['number'], f['year']) for f in flist if f['number'] and f['year']]
        from collections import Counter
        pair_counts = Counter(num_year_pairs)
        dupes = {k: v for k, v in pair_counts.items() if v > 1}
        dupe_count = sum(v - 1 for v in dupes.values())
        
        print(f"  {lt}: {len(flist)} files, {unique_count} unique laws")
        print(f"    Year range: {years_range} ({num_years} years)")
        if dupe_count > 0:
            print(f"    DUPLICATES: {dupe_count} duplicate files")
            for (n, y), c in list(dupes.items())[:5]:
                matching = [f['name'][:70] for f in flist if f['number'] == n and f['year'] == y]
                print(f"      No.{n}/{y} appears {c}x: {matching}")
        
        # Show samples
        samples = flist[:3]
        for s in samples:
            print(f"    - {s['name'][:90]} ({s['size']//1024}KB)")
        if len(flist) > 3:
            print(f"    ... +{len(flist)-3} more")
        
        folder_summary[lt] = {
            'total_files': len(flist),
            'unique_laws': unique_count,
            'year_range': years_range,
            'duplicates': dupe_count,
            'sample_names': [f['name'][:100] for f in samples]
        }
    
    all_results[folder_name] = folder_summary

# Save
out_path = '/home/z/my-project/download/coverage_deep.json'
with open(out_path, 'w') as f:
    json.dump(all_results, f, indent=2, ensure_ascii=False)
print(f"\nSaved to: {out_path}")