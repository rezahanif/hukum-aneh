#!/usr/bin/env python3
"""Apply v6 fixes to clean_extractor.py and parsers.py."""

path_ce = '/home/z/my-project/hukum-aneh/backend/python/chunker/clean_extractor.py'
with open(path_ce, 'r') as f:
    ce_lines = f.readlines()

ce_text = ''.join(ce_lines)

# ============================================================
# FIX 1: Currency glyph normalization in GLYPH_FIXES
# ============================================================
old1 = """    (re.compile(r'(\\d)O(\\s|T|$)'), r'\\g<1>0\\g<2>'),
    # Fix known garbled REPUBLIK patterns"""

new1 = """    (re.compile(r'(\\d)O(\\s|T|$)'), r'\\g<1>0\\g<2>'),
    # v6: Currency (Rp) digit normalization — fix O/a/o inside Rp amounts
    (re.compile(r'(Rp[\\s.]*)[Oo](?=\\d)'), r'\\g<1>0'),
    (re.compile(r'(Rp[\\s.]*\\d[\\d.]*)[Oo](?=\\d)'), r'\\g<1>0'),
    (re.compile(r'(Rp[\\s.]*\\d[\\d.]*)[Oo](?=\\d)'), r'\\g<1>0'),
    (re.compile(r'(Rp[\\s.]*\\d[\\d.]*)[aA](?=\\d)'), r'\\g<1>0'),
    (re.compile(r'(Rp[\\s.]*\\d[\\d.]*)[oO](?=,|[^0-9Oo])'), r'\\g<1>0'),
    # Fix known garbled REPUBLIK patterns"""

assert old1 in ce_text, "FIX 1: pattern not found"
ce_text = ce_text.replace(old1, new1, 1)
print("[FIX 1] Currency glyph normalization added")

# ============================================================
# FIX 2: Add FTRESIDEN to RE_GARBLED_HEADER
# ============================================================
ce_text = ce_text.replace(
    "r'|FRESIDEN'",
    "r'|FRESIDEN'\n    r'|FTRESIDEN'",
    1
)
print("[FIX 2a] Added FTRESIDEN to RE_GARBLED_HEADER")

# ============================================================
# FIX 2b: Add fuzzy noise module after RE_SHORT_NOISE
# ============================================================
marker = "RE_SHORT_NOISE = re.compile(r'^[A-Z]\\.?$')"
fuzzy_code = '''

# v6: Fuzzy noise detection
_FUZZY_BOILERPLATE = [
    'PRESIDEN REPUBLIK INDONESIA',
    'REPUBLIK INDONESIA',
    'SALINAN SESUAI DENGAN ASLINYA',
    'LEMBARAN NEGARA REPUBLIK INDONESIA',
    'MENTERI SEKRETARIAT NEGARA',
]

def _edit_distance(s1, s2):
    if abs(len(s1) - len(s2)) > 10:
        return 999
    m, n = len(s1), len(s2)
    if m < n:
        s1, s2 = s2, s1
        m, n = n, m
    prev = list(range(n + 1))
    for i in range(1, m + 1):
        curr = [i] + [0] * n
        for j in range(1, n + 1):
            cost = 0 if s1[i-1] == s2[j-1] else 1
            curr[j] = min(curr[j-1] + 1, prev[j] + 1, prev[j-1] + cost)
        prev = curr
    return prev[n]

def _is_fuzzy_noise(text):
    t = text.strip().upper()
    if len(t) < 10:
        return False
    alpha_count = sum(1 for c in t if c.isalpha())
    if alpha_count < len(t) * 0.5:
        return False
    for canonical in _FUZZY_BOILERPLATE:
        if _edit_distance(t, canonical) <= max(8, len(canonical) // 4):
            return True
    return False
'''

assert marker in ce_text, "FIX 2b: RE_SHORT_NOISE marker not found"
ce_text = ce_text.replace(marker, marker + fuzzy_code, 1)
print("[FIX 2b] Added fuzzy noise detection module")

# ============================================================
# FIX 2c: Add fuzzy noise check in is_noise_line
# ============================================================
ce_text = ce_text.replace(
    '    if RE_GARBLED_HEADER.match(t):\n        return True\n    # Standalone page numbers',
    '    if RE_GARBLED_HEADER.match(t):\n        return True\n    # v6: fuzzy noise detection\n    if _is_fuzzy_noise(t):\n        return True\n    # Standalone page numbers',
    1
)
print("[FIX 2c] Added _is_fuzzy_noise() to is_noise_line()")

with open(path_ce, 'w') as f:
    f.write(ce_text)
print(f"Wrote {path_ce}")

# ============================================================
# FIX 3: parsers.py — UUD '* Pasal XXx' and glued patterns
# ============================================================
path_pa = '/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py'
with open(path_pa, 'r') as f:
    pa_text = f.read()

# 3a: Extend RE_PASAL to match '* Pasal XXx'
old_re = "RE_PASAL = re.compile(r'^\\s*pasal\\s+([IVXLCDM]+|\\d+(?:\\s+\\d+)*[a-zA-Z]?)\\s*[.}\\s]*$', re.IGNORECASE)"
new_re = "RE_PASAL = re.compile(r'^\\*?\\s*pasal\\s+([IVXLCDM]+|\\d+(?:\\s+\\d+)*[a-zA-Z]?)\\s*[.}\\s]*$', re.IGNORECASE)\n    RE_PASAL_EMBEDDED = re.compile(r'\\*\\s*Pasal\\s+(\\d+[a-zA-Z]*)', re.IGNORECASE)"

assert old_re in pa_text, f"FIX 3a: RE_PASAL not found"
pa_text = pa_text.replace(old_re, new_re, 1)
print("[FIX 3a] Extended RE_PASAL and added RE_PASAL_EMBEDDED")

# 3b: After BAB match, check for glued pasal
old_bab = "                # v4: reset pasal when BAB changes to prevent heading-as-chunk\n                self.current_pasal = None\n                continue\n\n            m = self.RE_BAGIAN.match(text)"
new_bab = """                # v4: reset pasal when BAB changes to prevent heading-as-chunk
                self.current_pasal = None
                # v6: Check if a Pasal reference is glued to the BAB line
                m_glued = self.RE_PASAL_EMBEDDED.search(text)
                if m_glued:
                    self.current_pasal = m_glued.group(1)
                    self.current_ayat = None
                continue

            m = self.RE_BAGIAN.match(text)"""

assert old_bab in pa_text, "FIX 3b: BAB block not found"
pa_text = pa_text.replace(old_bab, new_bab, 1)
print("[FIX 3b] Added glued-pasal extraction after BAB match")

with open(path_pa, 'w') as f:
    f.write(pa_text)
print(f"Wrote {path_pa}")

print("\n=== All v6 core fixes applied ===")
