import sys, os
sys.path.insert(0, '/home/z/my-project/scripts')
from extractor import extract_lines_for_parsing
from parser_statute import StatuteParser

# Diagnose EMPTY_PASAL - check pp (PP_NO_70_TH_1991.pdf)
path = '/home/z/my-project/download/samples/pp/PP_NO_70_TH_1991.pdf'
lines, stats = extract_lines_for_parsing(path)
parser = StatuteParser(lines, stats)
result = parser.parse()

print('=== EMPTY_PASAL diagnosis (PP_NO_70_TH_1991.pdf) ===')
for p in result['pasals'][:5]:
    print(f'\nPasal {p["pasal_num"]} (BAB {p["bab"]})')
    print(f'  Ayats: {len(p["ayats"])}')
    if not p['ayats']:
        # Find lines between this pasal and next
        idx = result['pasals'].index(p)
        # Find line index of this pasal heading
        pasal_text = f'Pasal {p["pasal_num"]}'
        start = None
        end = None
        for li, l in enumerate(lines):
            if pasal_text.lower() in l['normalized'].lower() and l['normalized'].lower().strip().startswith('pasal'):
                if start is None:
                    start = li
                else:
                    end = li
                    break
        if start is not None:
            end = end or start + 15
            print(f'  Lines after pasal heading ({start} to {end}):')
            for l in lines[start:end]:
                print(f'    p{l["page"]:02d} fs={l["font_size"]:5.1f} B={str(l["is_bold"]):5s} | {l["normalized"][:90]}')

# Diagnose DUPLICATE_PASAL in pp
print('\n\n=== DUPLICATE_PASAL diagnosis (PP_NO_70_TH_1991.pdf) ===')
seen = {}
for p in result['pasals']:
    num = p['pasal_num']
    if num in seen:
        print(f'  DUP: Pasal {num} in BAB {p["bab"]} (prev was BAB {seen[num]})')
    else:
        seen[num] = p['bab']

# Check a known problematic: Kemnaker
print('\n\n=== EMPTY_PASAL diagnosis (Kemnaker) ===')
path2 = '/home/z/my-project/download/samples/JDIH_Kemnaker/Permenaker No. 99 Tahun 2013.pdf'
lines2, stats2 = extract_lines_for_parsing(path2)
parser2 = StatuteParser(lines2, stats2)
result2 = parser2.parse()
empty_count = sum(1 for p in result2['pasals'] if len(p['ayats']) == 0)
print(f'Total pasals: {len(result2["pasals"])}, empty: {empty_count}')
# Show lines around Pasal 1 (likely empty)
for li, l in enumerate(lines2):
    if l['normalized'].lower().strip() == 'pasal 1':
        print(f'  Found "Pasal 1" at line {li}')
        for l2 in lines2[li:li+15]:
            print(f'    p{l2["page"]:02d} fs={l2["font_size"]:5.1f} B={str(l2["is_bold"]):5s} | {l2["normalized"][:90]}')
        break

# Show lines around ayat (1) that works
found_ayat = False
for li, l in enumerate(lines2):
    if l['normalized'].strip().startswith('(1)') and not found_ayat:
        found_ayat = True
        print(f'\n  First ayat (1) at line {li}')
        for l2 in lines2[max(0,li-2):li+8]:
            print(f'    p{l2["page"]:02d} fs={l2["font_size"]:5.1f} B={str(l2["is_bold"]):5s} | {l2["normalized"][:90]}')