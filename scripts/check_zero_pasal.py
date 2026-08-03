import sys, os
sys.path.insert(0, '/home/z/my-project/scripts')
from extractor import extract_lines_for_parsing

files = [
    ('JDIH_KPU', 'download/samples/JDIH_KPU/PKPU_8_2026.pdf'),
    ('JDIH_Kemendag', 'download/samples/JDIH_Kemendag/Kepmendag_No__123_Tahun_2025_download_3142_2.pdf'),
    ('Putusan-MK', 'download/samples/Putusan-MK/putusan_mkri_5301.pdf'),
    ('inpres', 'download/samples/inpres/inpres-no-3-tahun-2023_Percepatan Peningkatan Konektivitas Jalan Daerah.pdf'),
    ('tap_mpr', 'download/samples/tap_mpr/Ketetapan Majelis Permusyawaratan Rakyat Republik Indonesia Nomor IIIMPR2002 ten.pdf'),
]

for label, path in files:
    print(f'\n=== {label}: {os.path.basename(path)} ===')
    try:
        lines, stats = extract_lines_for_parsing(path)
        print(f'  Total lines: {stats["total_lines"]}')
        print(f'  --- Lines with pasal/bab/bagian keywords ---')
        for l in lines:
            n = l['normalized'].lower()
            if any(kw in n for kw in ['pasal', 'bab ', 'bagian', 'paragraf', 'ayat']):
                print(f'  p{l["page"]:02d} fs={l["font_size"]:5.1f} B={str(l["is_bold"]):5s} | {l["normalized"][:100]}')
        print(f'  --- First 10 lines ---')
        for l in lines[:10]:
            print(f'  p{l["page"]:02d} fs={l["font_size"]:5.1f} B={str(l["is_bold"]):5s} | {l["normalized"][:100]}')
    except Exception as e:
        print(f'  ERROR: {e}')