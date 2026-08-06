#!/usr/bin/env python3
"""Apply v4 QA fixes — all root causes addressed."""
import re

# ============================================================
# FIX 1: parsers.py — RE_PASAL to capture space-separated multi-digit
# ============================================================
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'r') as f:
    pcontent = f.read()

# 1a: Fix RE_PASAL regex to capture space-separated multi-digit numbers
old_re_pasal = "RE_PASAL = re.compile(r'^\\s*pasal\\s+([IVXLCDM]+|\\d+[a-zA-Z]?)\\s*\\.?(?:\\s|$)', re.IGNORECASE)"
new_re_pasal = "RE_PASAL = re.compile(r'^\\s*pasal\\s+([IVXLCDM]+|\\d+(?:\\s+\\d+)*)', re.IGNORECASE)"
if old_re_pasal in pcontent:
    pcontent = pcontent.replace(old_re_pasal, new_re_pasal, 1)
    print('[parsers] Fixed RE_PASAL to capture space-separated multi-digit numbers')
else:
    print('[parsers] RE_PASAL already fixed or pattern not found')

# 1b: After matching pasal, normalize the number (remove spaces) and capture remaining text
old_pasal_match = '''            m = self.RE_PASAL.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                self.current_pasal = m.group(1)
                self.current_ayat = None
                continue'''
new_pasal_match = '''            m = self.RE_PASAL.match(text)
            if m:
                self._flush_ayat()
                self._flush_pasal()
                raw_num = m.group(1)
                # v4: normalize space-separated digits ("1 1" -> "11", "1 10" -> "110")
                self.current_pasal = re.sub(r'\s+', '', raw_num) if re.match(r'^\d', raw_num) else raw_num
                self.current_ayat = None
                # v4: capture any remaining text on the same line as pasal heading
                remaining = text[m.end():].strip().lstrip('.').strip()
                if remaining and len(remaining) > 5:
                    self.current_pasal_lines.append(remaining)
                continue'''
if old_pasal_match in pcontent:
    pcontent = pcontent.replace(old_pasal_match, new_pasal_match, 1)
    print('[parsers] Pasal match now normalizes multi-digit numbers and captures inline text')
else:
    print('[parsers] Pasal match block not found')

# 1c: Reset current_pasal when BAB changes (fixes heading-as-chunk bug)
old_bab_reset = '''                self.current_bagian = None
                self.current_bagian_label = None
                self.current_paragraf = None
                self.current_paragraf_label = None
                continue

            m = self.RE_BAGIAN.match(text)'''
new_bab_reset = '''                self.current_bagian = None
                self.current_bagian_label = None
                self.current_paragraf = None
                self.current_paragraf_label = None
                # v4: reset pasal when BAB changes to prevent heading-as-chunk
                self.current_pasal = None
                continue

            m = self.RE_BAGIAN.match(text)'''
if old_bab_reset in pcontent:
    pcontent = pcontent.replace(old_bab_reset, new_bab_reset, 1)
    print('[parsers] BAB change now resets current_pasal')
else:
    print('[parsers] BAB reset block not found')

# 1d: Remove broad-type guard — title block first, directory as fallback only when Unknown
old_doc_type = '''    # v3 fix: Detect doc type from title block FIRST (more reliable),
    # use directory config as fallback. Flag mismatches for review.
    # Use only first 3 lines for detection to avoid citation pollution.
    title_block = extract_title_block(lines, max_first_n=3)
    detected_type, detected_issuer = detect_doc_type_from_title(title_block)
    
    # If title block detection succeeds, use it (it reads the actual document)
    # BUT: if directory config exists and is more specific, prefer it when
    # title detection returns a generic/broad match from a citation.
    BROAD_TYPES = {"PP", "Peraturan", "Unknown"}
    if detected_type != "Unknown":
        if (default_doc_type != "Unknown" and 
            detected_type != default_doc_type and
            detected_type in BROAD_TYPES and
            default_doc_type not in BROAD_TYPES):
            # Directory config is more specific (e.g. Permen vs PP from citation)
            doc_type = default_doc_type
            issuer = default_issuer
        else:
            doc_type = detected_type
            issuer = detected_issuer
            if default_doc_type != "Unknown" and detected_type != default_doc_type:
                print(f"  [WARN] doc_type mismatch: title={detected_type} vs dir={default_doc_type} for {directory}")
    elif default_doc_type != "Unknown":
        doc_type = default_doc_type
        issuer = default_issuer
    else:
        doc_type = "Unknown"
        issuer = "Unknown"'''

new_doc_type = '''    # v4 fix: Detect doc_type from title block FIRST, directory as fallback.
    # Use RAW lines (pre-boilerplate) for detection since boilerplate stripping
    # removes the actual document title (e.g. "PERATURAN PRESIDEN" = preamble header).
    raw_title_lines = [l["text"].strip() for l in lines[:25] if l.get("page", 1) <= 2]
    raw_title_block = " ".join(raw_title_lines[:10])
    detected_type, detected_issuer = detect_doc_type_from_title(raw_title_block)
    
    # Also try with cleaned (post-boilerplate) title block as second source
    if detected_type == "Unknown":
        clean_title = extract_title_block(lines, max_first_n=5)
        detected_type, detected_issuer = detect_doc_type_from_title(clean_title)
    
    # Title block wins when it returns a real type; directory is fallback only
    if detected_type != "Unknown":
        doc_type = detected_type
        issuer = detected_issuer
        if default_doc_type != "Unknown" and detected_type != default_doc_type:
            print(f"  [WARN] doc_type: title={detected_type} overrides dir={default_doc_type} for {directory}")
    elif default_doc_type != "Unknown":
        doc_type = default_doc_type
        issuer = default_issuer
    else:
        doc_type = "Unknown"
        issuer = "Unknown"'''

if 'v4 fix: Detect doc_type from title block' not in pcontent:
    if old_doc_type in pcontent:
        pcontent = pcontent.replace(old_doc_type, new_doc_type, 1)
        print('[parsers] doc_type: raw lines first, clean fallback, directory last')
    else:
        print('[parsers] doc_type block not found (check for partial match)')
else:
    print('[parsers] doc_type v4 fix already applied')

with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'w') as f:
    f.write(pcontent)

print('--- parsers.py v4 fixes applied ---')


# ============================================================
# FIX 2: clean_extractor.py — JDIH_KPU nomor/year, glyph fixes, stamp
# ============================================================
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/clean_extractor.py', 'r') as f:
    cecontent = f.read()

# 2a: Add glyph fixes for perpres corruption patterns
glyph_insert_point = "    (re.compile(r'Cukuo', re.I), 'Cukup'),"
new_glyphs = '''    (re.compile(r'Cukuo', re.I), 'Cukup'),
    # v4: perpres glyph corruption variants
    (re.compile(r'T\r\n', re.I), 'Tun'),     # carriage-return corruption: T\r\njangan -> Tunjangan
    (re.compile(r'Tfrnjangan', re.I), 'Tunjangan'),
    (re.compile(r'PRESTDEN', re.I), 'PRESIDEN'),
    (re.compile(r'PRESIOEN', re.I), 'PRESIDEN'),  # already existed, keep as safety'''

if 'Tfrnjangan' not in cecontent:
    cecontent = cecontent.replace(glyph_insert_point, new_glyphs, 1)
    print('[clean_extractor] Added perpres glyph fixes (T\\r\\n, Tfrnjangan, PRESTDEN)')
else:
    print('[clean_extractor] Perpres glyph fixes already present')

# 2b: Broader stamp pattern — catch "SK No243676A" (no space after No)
old_stamp2 = "RE_SETNEG_STAMP2 = re.compile(r'^SK\\s+No\\.?\\s*\\d+\\s*$', re.IGNORECASE)  # broader stamp pattern"
new_stamp2 = "RE_SETNEG_STAMP2 = re.compile(r'^SK\s+No\.?\s*\d+\s*[A-Z]?$', re.IGNORECASE)  # broader: SK No243676A, SK No 123 A"

if old_stamp2 in cecontent:
    cecontent = cecontent.replace(old_stamp2, new_stamp2, 1)
    print('[clean_extractor] Broadened RE_SETNEG_STAMP2 to catch No<digit> variants')
else:
    print('[clean_extractor] Stamp pattern check needed')

# 2c: For JDIH_KPU — add post-menetapkan nomor/year extraction
old_nomor_year_end = "    return None, None\n\n\ndef extract_title_block"
new_nomor_year_end = '''    return None, None


def extract_nomor_year_post_decision(lines: List[Dict]) -> Tuple[Optional[str], Optional[str]]:
    """v4: Extract nomor/year from after the MENETAPKAN/MEMUTUSKAN keyword.
    
    For documents where page 1 is an image and the first text page starts
    with "Menimbang" considerations (citations), the document's own
    nomor/year appears AFTER the decision keyword.
    """
    in_decision = False
    text_buffer = []
    
    for line in lines:
        text = line["text"].strip()
        t_upper = text.upper()
        
        if re.match(r'^(?:MENETAPKAN|MEMUTUSKAN|MENGINSTRUKSIKAN)\b', t_upper):
            in_decision = True
            continue
        
        if in_decision:
            text_buffer.append(text)
            if len(text_buffer) >= 15:
                break
    
    combined = " ".join(text_buffer)
    return extract_nomor_year(combined)


def extract_title_block'''

if 'extract_nomor_year_post_decision' not in cecontent:
    if old_nomor_year_end in cecontent:
        cecontent = cecontent.replace(old_nomor_year_end, new_nomor_year_end, 1)
        print('[clean_extractor] Added extract_nomor_year_post_decision for JDIH_KPU-type docs')
    else:
        print('[clean_extractor] Insertion point not found')
else:
    print('[clean_extractor] extract_nomor_year_post_decision already exists')

with open('/home/z/my-project/hukum-aneh/backend/python/chunker/clean_extractor.py', 'w') as f:
    f.write(cecontent)

print('--- clean_extractor.py v4 fixes applied ---')


# ============================================================
# FIX 3: parsers.py — Use post-decision nomor/year for JDIH_KPU
# ============================================================
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'r') as f:
    pcontent = f.read()

# Update import to include new function
old_import = 'from clean_extractor import extract_clean, extract_title_block, detect_doc_type_from_title, fix_glyph_corruption'
new_import = 'from clean_extractor import extract_clean, extract_title_block, detect_doc_type_from_title, fix_glyph_corruption, extract_nomor_year_post_decision'

if 'extract_nomor_year_post_decision' not in pcontent.split('\n')[0:30]:
    if old_import in pcontent:
        pcontent = pcontent.replace(old_import, new_import, 1)
        print('[parsers] Added extract_nomor_year_post_decision import')
    else:
        print('[parsers] Import line not found')
else:
    print('[parsers] Import already present')

# Add JDIH_KPU special case: if page 1 is empty and nomor looks wrong, re-extract
old_jdih_kpu = '''    # Special case for JDIH_KPU: page 1 is image, detect from page 2+ content
    if doc_type == "Unknown" and directory == "JDIH_KPU":
        doc_type = "PKPU"
        issuer = "Komisi Pemilihan Umum"'''
new_jdih_kpu = '''    # v4: If page 1 was empty (image), re-extract nomor/year from post-decision text
    if meta.get("total_pages", 0) > 1:
        page1_has_text = any(l.get("page") == 1 for l in lines)
        if not page1_has_text and (not nomor or not year):
            post_nomor, post_year = extract_nomor_year_post_decision(lines)
            if post_nomor:
                nomor = post_nomor
                print(f"  [INFO] Post-decision extraction: nomor={nomor}, year={year} (was {meta.get('nomor')}/{meta.get('year')})")
            if post_year and not year:
                year = post_year
    
    # Special case for JDIH_KPU: page 1 is image, detect from page 2+ content
    if doc_type == "Unknown" and directory == "JDIH_KPU":
        doc_type = "PKPU"
        issuer = "Komisi Pemilihan Umum"'''

if 'Post-decision extraction' not in pcontent:
    if old_jdih_kpu in pcontent:
        pcontent = pcontent.replace(old_jdih_kpu, new_jdih_kpu, 1)
        print('[parsers] Added post-decision nomor/year re-extraction for image-cover docs')
    else:
        print('[parsers] JDIH_KPU block not found')
else:
    print('[parsers] Post-decision extraction already present')

with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'w') as f:
    f.write(pcontent)

print('--- parsers.py v4 JDIH_KPU fix applied ---')


# ============================================================
# FIX 4: parsers.py — penjelasan for pp (KETENTUAN PENUTUP path)
# ============================================================
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'r') as f:
    pcontent = f.read()

# Add KETENTUAN PENUTUP as a penjelasan/penutup marker
old_re_penutup = "RE_PENUTUP = re.compile(r'^(?:PENUTUP|penutup)\\b')"
new_re_penutup = "RE_PENUTUP = re.compile(r'^(?:KETENTUAN\s+PENUTUP|PENUTUP|penutup)\\b')"

if 'KETENTUAN' not in pcontent.split('RE_PENUTUP')[1][:60] if 'RE_PENUTUP' in pcontent else True:
    if old_re_penutup in pcontent:
        pcontent = pcontent.replace(old_re_penutup, new_re_penutup, 1)
        print('[parsers] Added KETENTUAN PENUTUP to RE_PENUTUP pattern')
    else:
        print('[parsers] RE_PENUTUP not found')
else:
    print('[parsers] KETENTUAN PENUTUP already in RE_PENUTUP')

with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'w') as f:
    f.write(pcontent)

print('--- parsers.py v4 penutup pattern fix applied ---')
print('\nAll v4 fixes applied!')
