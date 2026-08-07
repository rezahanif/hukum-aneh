#!/usr/bin/env python3
import json, re
from collections import Counter

data = json.load(open('/home/z/my-project/download/chunk_results_v6.json'))

# 1. UUD 1945 alpha pasals
uud = data['uud-1945']
pasals = set()
for c in uud:
    p = c.get('metadata',{}).get('pasal')
    if p: pasals.add(str(p))
alpha = [p for p in sorted(pasals) if any(ch.isalpha() for ch in p)]
print('=== UUD 1945 ALPHA PASALS ===')
print(f'Found {len(alpha)} alpha pasals: {alpha}')
print()
for base in ['22', '23', '24', '25', '26', '27', '28']:
    variants = [p for p in pasals if p.startswith(base)]
    print(f'Pasal {base} variants: {variants}')
print()

# 2. perppu remaining dups
perppu = data['perppu']
idc = Counter(c['id'] for c in perppu)
dups = {k:v for k,v in idc.items() if v > 1}
print(f'=== PERPPU REMAINING DUPS: {len(dups)} IDs ===')
for did, cnt in sorted(dups.items(), key=lambda x:-x[1])[:5]:
    chunks = [c for c in perppu if c['id'] == did]
    for c in chunks:
        t = c["text"]
        print(f'  [{did}] len={len(t):4d} text={t[:80]}')
    print()

# 3. perpres currency
perpres = data['perpres']
print('=== PERPRES: Currency check ===')
for c in perpres:
    if 'Rp' in c['text']:
        amounts = re.findall(r'Rp[\d.,]+', c['text'])
        print(f'  ID={c["id"]}')
        for a in amounts:
            print(f'    {a}')
        bad = re.findall(r'Rp[\d.]*[Oo][\d.,]*', c['text'])
        if bad:
            print(f'  REMAINING BAD: {bad}')
print()

# 4. UU garbled header
print('=== UU: Pasal 13 check ===')
uu = data['uu']
for c in uu:
    if ':13' in c['id'] or (c.get('metadata',{}).get('pasal') == '13'):
        t = c['text']
        print(f'  [{c["id"]}] len={len(t):4d} text={t[:100]}')
print()
print('=== UU: Garbled header check ===')
found_garbled = False
for c in uu:
    t = c['text']
    if 'FTRESIDEN' in t or 'Ei:IUE' in t or 'IND ONES IA' in t:
        print(f'  GARBLED: [{c["id"]}] {t[:100]}')
        found_garbled = True
if not found_garbled:
    print('  No garbled header chunks found (FIXED)')