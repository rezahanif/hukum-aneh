import sys
sys.path.insert(0, '/home/z/my-project/scripts')
from extractor import extract_lines_for_parsing
from parser_statute import StatuteParser

path = 'download/samples/JDIH_Kemnaker/Permenaker No. 99 Tahun 2013.pdf'
lines, stats = extract_lines_for_parsing(path)
parser = StatuteParser(lines, stats)
result = parser.parse()

print(f'Total pasals: {len(result["pasals"])}')

# Count dups
from collections import Counter
cnts = Counter(p['pasal_num'] for p in result['pasals'])
dups = {k:v for k,v in cnts.items() if v > 1}
print(f'Unique pasal nums: {len(cnts)}')
print(f'Duplicated pasal nums: {len(dups)}')

# Show lines around 'Pasal 4' in the PDF
print('\nLines containing "Pasal 4":')
for li, l in enumerate(lines):
    n = l['normalized'].lower()
    if 'pasal 4' in n:
        print(f'  L{li}: p{l["page"]:02d} fs={l["font_size"]:5.1f} B={str(l["is_bold"]):5s} | {l["normalized"][:100]}')

# Show first 30 lines
print('\nFirst 30 lines:')
for l in lines[:30]:
    print(f'  p{l["page"]:02d} fs={l["font_size"]:5.1f} B={str(l["is_bold"]):5s} | {l["normalized"][:100]}')

# Show lines 100-120
print('\nLines 100-120:')
for l in lines[100:120]:
    print(f'  p{l["page"]:02d} fs={l["font_size"]:5.1f} B={str(l["is_bold"]):5s} | {l["normalized"][:100]}')
