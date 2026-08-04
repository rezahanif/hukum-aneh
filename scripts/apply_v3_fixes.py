#!/usr/bin/env python3
"""Apply v3 QA fixes to clean_extractor.py and parsers.py."""

import re

# ============================================================
# FIX 1: clean_extractor.py — Add broader stamp pattern + update is_noise_line
# ============================================================
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/clean_extractor.py', 'r') as f:
    content = f.read()

# 1a: Add RE_SETNEG_STAMP2 after RE_SETNEG_STAMP
old = "RE_SETNEG_STAMP = re.compile(r'^SK\\s+No\\s+\\d+\\s*A$', re.IGNORECASE)\nRE_PAGE_NUMBER"
new = """RE_SETNEG_STAMP = re.compile(r'^SK\\s+No\\s+\\d+\\s*A$', re.IGNORECASE)
RE_SETNEG_STAMP2 = re.compile(r'^SK\\s+No\\.?\\s*\\d+\\s*$', re.IGNORECASE)  # broader stamp pattern
RE_PAGE_NUMBER"""
if 'RE_SETNEG_STAMP2' not in content:
    content = content.replace(old, new, 1)
    print("[clean_extractor] Added RE_SETNEG_STAMP2")
else:
    print("[clean_extractor] RE_SETNEG_STAMP2 already exists")

# 1b: Update is_noise_line to use RE_SETNEG_STAMP2
old_noise = """    if RE_SETNEG_STAMP.match(t):
        return True
    if RE_PAGE_NUMBER.match(t):"""
new_noise = """    if RE_SETNEG_STAMP.match(t):
        return True
    if RE_SETNEG_STAMP2.match(t):
        return True
    if RE_PAGE_NUMBER.match(t):"""
if 'RE_SETNEG_STAMP2.match' not in content:
    content = content.replace(old_noise, new_noise, 1)
    print("[clean_extractor] Updated is_noise_line with RE_SETNEG_STAMP2")
else:
    print("[clean_extractor] is_noise_line already has RE_SETNEG_STAMP2")

# 1c: Restrict extract_nomor_year to title block (first ~500 chars)
old_nomor = 'def extract_nomor_year(full_text: str) -> Tuple[Optional[str], Optional[str]]:\n    """Extract document number and year from title block text."""\n    # Apply glyph fixes first for better matching\n    fixed = fix_glyph_corruption(full_text)'
new_nomor = 'def extract_nomor_year(full_text: str) -> Tuple[Optional[str], Optional[str]]:\n    """Extract document number and year from title block text.\n    \n    v3 fix: Only search within the first ~500 chars (title block) to avoid\n    grabbing body citations as the document\'s own nomor/year.\n    """\n    # Restrict search to title block region (first ~500 chars)\n    search_text = full_text[:500] if len(full_text) > 500 else full_text\n    # Apply glyph fixes first for better matching\n    fixed = fix_glyph_corruption(search_text)'
if 'search_text = full_text[:500]' not in content:
    content = content.replace(old_nomor, new_nomor, 1)
    print("[clean_extractor] Restricted extract_nomor_year to title block")
else:
    print("[clean_extractor] extract_nomor_year already restricted")

# 1d: Add glyph consistency — normalize O-as-0 more broadly
old_glyph = "GLYPH_FIXES = ["
new_glyph = """GLYPH_FIXES = [
    # v3: Consistent O-as-0 normalization — handle all positions
    (re.compile(r'(?<=\d)O(?=\d)'), '0'),  # between digits: 2O24 -> 2024
    (re.compile(r'(?<=\d)O(?=\s|$)'), '0'),  # digit-O-boundary: 2O -> 20
    (re.compile(r'(?<=\s)O(?=\d)'), '0'),  # boundary-O-digit: O24 -> 024"""
# Only replace the first line (GLYPH_FIXES opening) but keep the rest
if 'v3: Consistent O-as-0' not in content:
    # Find the opening of GLYPH_FIXES list and insert after it
    content = content.replace(
        "GLYPH_FIXES = [\n    # O-as-zero in year/number contexts: 2OO9, 2O2O, 2O24, etc.",
        new_glyph,
        1
    )
    print("[clean_extractor] Added consistent O-as-0 glyph fixes")
else:
    print("[clean_extractor] O-as-0 consistency already added")

with open('/home/z/my-project/hukum-aneh/backend/python/chunker/clean_extractor.py', 'w') as f:
    f.write(content)

print("\n--- clean_extractor.py fixes applied ---\n")

# ============================================================
# FIX 2: parsers.py — Multiple critical fixes
# ============================================================
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'r') as f:
    pcontent = f.read()

# 2a: doc_type resolution — title block first, directory as fallback
# Change from "if default_doc_type != Unknown -> use it" to
# "if detected_type != Unknown -> use it, else fallback to default"
old_doc_type = '''    # Detect doc type: prefer directory config, use detection only as fallback.
    # The directory mapping is authoritative (set by the pipeline operator).
    title_block = extract_title_block(lines, max_first_n=5)
    detected_type, detected_issuer = detect_doc_type_from_title(title_block)
    
    if default_doc_type != "Unknown":
        doc_type = default_doc_type
        issuer = default_issuer
    elif detected_type != "Unknown":
        doc_type = detected_type
        issuer = detected_issuer
    else:
        doc_type = "Unknown"
        issuer = "Unknown"'''

new_doc_type = '''    # v3 fix: Detect doc type from title block FIRST (more reliable),
    # use directory config as fallback. Flag mismatches for review.
    title_block = extract_title_block(lines, max_first_n=5)
    detected_type, detected_issuer = detect_doc_type_from_title(title_block)
    
    # If title block detection succeeds, use it (it reads the actual document)
    if detected_type != "Unknown":
        doc_type = detected_type
        issuer = detected_issuer
        # Warn if directory config disagrees (indicates possible misfiling)
        if default_doc_type != "Unknown" and detected_type != default_doc_type:
            print(f"  [WARN] doc_type mismatch: title_block={detected_type} vs dir_config={default_doc_type} for {directory}")
    elif default_doc_type != "Unknown":
        doc_type = default_doc_type
        issuer = default_issuer
    else:
        doc_type = "Unknown"
        issuer = "Unknown"'''

if 'v3 fix: Detect doc type from title block FIRST' not in pcontent:
    pcontent = pcontent.replace(old_doc_type, new_doc_type, 1)
    print("[parsers] Fixed doc_type resolution: title block first, dir fallback")
else:
    print("[parsers] doc_type resolution already fixed")

# 2b: Fix Penjelasan handling — detect PENJELASAN as hard boundary,
#     give it distinct ID suffix, prevent ID collision with operative articles

# First, add PENJELASAN to Family A preamble/end detection.
# In FamilyAParser.__init__, add in_penjelasan state
old_init = '''        self.in_penutup = False
        self.in_quoted_amendment = False
        self.quoted_amendment_source = None  # tracks which pasal is being amended
        self.current_ayat_lines = []
        self.current_pasal_lines = []
        self.quoted_amendment_lines = []  # accumulate quoted text for chunking'''

new_init = '''        self.in_penutup = False
        self.in_penjelasan = False  # v3: track elucidation section
        self.in_quoted_amendment = False
        self.quoted_amendment_source = None  # tracks which pasal is being amended
        self.current_ayat_lines = []
        self.current_pasal_lines = []
        self.quoted_amendment_lines = []  # accumulate quoted text for chunking'''

if 'in_penjelasan' not in pcontent:
    pcontent = pcontent.replace(old_init, new_init, 1)
    print("[parsers] Added in_penjelasan state to FamilyAParser")
else:
    print("[parsers] in_penjelasan already exists")

# 2c: Add PENJELASAN + signature block detection in FamilyA parse loop
# Insert after the penutup detection block
old_penutup_block = '''            # Detect penutup
            if self.RE_PENUTUP.match(text):
                self._flush_ayat()
                self._flush_pasal()
                self._flush_quoted_amendment()
                self.in_penutup = True
                continue

            if self.in_penutup:
                continue'''

new_penutup_block = '''            # Detect penutup
            if self.RE_PENUTUP.match(text):
                self._flush_ayat()
                self._flush_pasal()
                self._flush_quoted_amendment()
                self.in_penutup = True
                continue

            if self.in_penutup:
                # v3: Detect signature block and Penjelasan as hard boundaries
                text_lower = text.lower()
                # Signature block: lines like "Ditetapkan di Jakarta", "pada tanggal", or ttd/
                # Also detect by very short lines after penutup start
                if (re.search(r'\bttd\b|\bdisahkan\b|\bditetapkan\b', text_lower) or
                    (len(text) < 15 and text_upper == text and len(text) > 2)):
                    self._flush_penjelasan()  # flush any accumulated penjelasan
                    continue
                # PENJELASAN section header
                if text_upper == "PENJELASAN" or text_upper.startswith("PENJELASAN "):
                    self._flush_penjelasan()
                    self.in_penjelasan = True
                    self._penjelasan_pasal = None
                    continue
                # Inside Penjelasan: track pasal references for IDs
                if self.in_penjelasan:
                    m_pj = re.match(r'^\s*Pasal\s+(\d+[a-zA-Z]?)', text, re.IGNORECASE)
                    if m_pj:
                        self._flush_penjelasan()
                        self._penjelasan_pasal = m_pj.group(1)
                    else:
                        self._penjelasan_lines.append(text)
                continue'''

if 'in_penjelasan' not in pcontent or '_flush_penjelasan' not in pcontent:
    pcontent = pcontent.replace(old_penutup_block, new_penutup_block, 1)
    print("[parsers] Added Penjelasan + signature block handling")
else:
    print("[parsers] Penjelasan handling already exists")

# 2d: Add penjelasan accumulation init and flush method to FamilyAParser
# Add after _flush_quoted_amendment method
old_flush_quoted_end = '''        self.quoted_amendment_lines = []
        self.in_quoted_amendment = False
        self.quoted_amendment_source = None


# ==========================================================================
# FAMILY B: DECREE/DECISION PARSER'''

new_flush_quoted_end = '''        self.quoted_amendment_lines = []
        self.in_quoted_amendment = False
        self.quoted_amendment_source = None

    def _init_penjelasan(self):
        """Initialize penjelasan accumulation (called when PENJELASAN section starts)."""
        if not hasattr(self, '_penjelasan_lines'):
            self._penjelasan_lines = []
            self._penjelasan_pasal = None

    def _flush_penjelasan(self):
        """Flush accumulated Penjelasan text as chunk(s) with distinct :penjelasan ID suffix.
        
        v3 fix: Penjelasan (elucidation) gets distinct ID suffix to prevent
        collision with operative article IDs.
        """
        self._init_penjelasan()
        if not self._penjelasan_lines:
            return

        text = ' '.join(self._penjelasan_lines).strip()
        if not text or len(text) < MIN_CHUNK_TEXT_LENGTH:
            self._penjelasan_lines = []
            return

        pasal_ref = self._penjelasan_pasal or "umum"
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            pasal=pasal_ref, ayat="penjelasan")

        self.chunks.append({
            "id": chunk_id,
            "text": text,
            "metadata": {
                "doc_type": self.doc_type,
                "issuer": self.issuer,
                "year": self.year,
                "bab": self.current_bab or None,
                "pasal": pasal_ref,
                "ayat": "penjelasan",
                "parent": None,
                "path": f"Penjelasan > Pasal {pasal_ref}",
                "status": "penjelasan",
            }
        })
        self._penjelasan_lines = []
        self._penjelasan_pasal = None


# ==========================================================================
# FAMILY B: DECREE/DECISION PARSER'''

if '_flush_penjelasan' not in pcontent:
    pcontent = pcontent.replace(old_flush_quoted_end, new_flush_quoted_end, 1)
    print("[parsers] Added _init_penjelasan and _flush_penjelasan methods")
else:
    print("[parsers] _flush_penjelasan already exists")

# 2e: Initialize penjelasan state in __init__
old_init_check = 'self.quoted_amendment_lines = []  # accumulate quoted text for chunking'
new_init_check = '''self.quoted_amendment_lines = []  # accumulate quoted text for chunking
        self._penjelasan_lines = []  # v3: accumulate Penjelasan text
        self._penjelasan_pasal = None'''

if 'self._penjelasan_lines' not in pcontent:
    pcontent = pcontent.replace(old_init_check, new_init_check, 1)
    print("[parsers] Initialized _penjelasan_lines in __init__")
else:
    print("[parsers] _penjelasan_lines already initialized")

# 2f: Fix FamilyB — add signature/Penjelasan as hard boundary, add parent section to diktum ID
# For inpres: detect lettered sections (a., b., c.) as parent keys for diktum numbering
old_familyb_done = '''            if self.section == "done":
                continue

            # Signature block detection
            text_lower = text.lower()
            if "ttd" in text_lower and len(text) < 20:
                self._flush_item()
                self.section = "done"
                continue'''

new_familyb_done = '''            if self.section == "done":
                continue

            # v3: Signature block detection (broader)
            text_lower = text.lower()
            if ("ttd" in text_lower and len(text) < 30):
                self._flush_item()
                self.section = "done"
                continue
            # Also stop at formal signing language
            if re.search(r'\bditetapkan\s+di\b|\bdilantik\s+di\b', text_lower):
                self._flush_item()
                self.section = "done"
                continue
            # Stop at Penjelasan in Family B documents
            if text_upper == "PENJELASAN":
                self._flush_item()
                self.section = "done"
                continue'''

if 'v3: Signature block detection (broader)' not in pcontent:
    pcontent = pcontent.replace(old_familyb_done, new_familyb_done, 1)
    print("[parsers] Broadened FamilyB signature + Penjelasan boundary")
else:
    print("[parsers] FamilyB boundary already broadened")

# 2g: Fix inpres diktum numbering — track parent lettered section
old_familyb_init = '''        self.section = "pre"  # pre, menimbang, mengingat, keputusan, done
        self.current_item = None
        self.current_item_lines = []
        self.preamble_lines = []
        self.decision_keyword = None'''

new_familyb_init = '''        self.section = "pre"  # pre, menimbang, mengingat, keputusan, done
        self.current_item = None
        self.current_item_lines = []
        self.preamble_lines = []
        self.decision_keyword = None
        self.parent_section = None  # v3: track parent lettered section for inpres diktum IDs'''

if 'self.parent_section' not in pcontent:
    pcontent = pcontent.replace(old_familyb_init, new_familyb_init, 1)
    print("[parsers] Added parent_section tracking to FamilyB")
else:
    print("[parsers] parent_section already exists")

# 2h: Update FamilyB item matching to track parent section and use it in IDs
old_item_match = '''                if m_letter or m_number or m_ordinal:
                    self._flush_item()
                    if m_ordinal:
                        self.current_item = m_ordinal.group(1).upper()
                    elif m_letter:
                        self.current_item = m_letter.group(1)
                    else:
                        self.current_item = m_number.group(1)'''

new_item_match = '''                if m_letter or m_number or m_ordinal:
                    self._flush_item()
                    if m_ordinal:
                        self.current_item = m_ordinal.group(1).upper()
                    elif m_letter:
                        # v3: lettered items (a., b., c.) are parent sections for inpres
                        self.parent_section = m_letter.group(1)
                        self.current_item = m_letter.group(1)
                    else:
                        # v3: number items under a lettered parent get compound ID
                        if self.parent_section and self.current_item and self.current_item == self.parent_section:
                            pass  # already set as parent
                        self.current_item = m_number.group(1)'''

if 'self.parent_section' not in pcontent.split('def _flush_item')[0] if 'def _flush_item' in pcontent else True:
    # Only replace if parent_section is NOT already in the item matching block
    test_area = pcontent[pcontent.find('if m_letter or m_number or m_ordinal'):pcontent.find('if m_letter or m_number or m_ordinal')+600]
    if 'self.parent_section = m_letter' not in test_area:
        pcontent = pcontent.replace(old_item_match, new_item_match, 1)
        print("[parsers] Added parent section tracking in item matching")
    else:
        print("[parsers] Parent section tracking in item matching already exists")
else:
    print("[parsers] Parent section tracking already in init area")

# 2i: Update build_id for FamilyB diktum to include parent_section
old_flush_item_id = '''        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            chunk_type="diktum", chunk_num=self.current_item)'''

new_flush_item_id = '''        # v3: Include parent section in diktum ID for inpres (prevents numbering collision)
        diktum_num = self.current_item
        if self.parent_section and diktum != self.parent_section:
            diktum_num = f"{self.parent_section}.{diktum_num}"
        chunk_id = build_id(self.doc_type, self.nomor, self.year,
                            chunk_type="diktum", chunk_num=diktum_num)'''

if 'parent_section' not in pcontent.split('_flush_item')[1][:500] if '_flush_item' in pcontent else True:
    test_flush = pcontent[pcontent.find('def _flush_item'):pcontent.find('def _flush_item')+800]
    if 'self.parent_section' not in test_flush:
        pcontent = pcontent.replace(old_flush_item_id, new_flush_item_id, 1)
        print("[parsers] Updated _flush_item to include parent_section in diktum ID")
    else:
        print("[parsers] _flush_item already has parent_section")
else:
    print("[parsers] _flush_item already has parent_section")

# 2j: Fix perppu pasal counter — the counter uses RE_PASAL which requires trailing "." or whitespace.
# The issue is the regex is too strict. Make it more permissive.
old_re_pasal = "RE_PASAL = re.compile(r'^\\s*pasal\\s+([IVXLCDM]+|\\d+[a-zA-Z]?)\\s*\\.?\\s*.{0,5}$', re.IGNORECASE)"
new_re_pasal = "RE_PASAL = re.compile(r'^\\s*pasal\\s+([IVXLCDM]+|\\d+[a-zA-Z]?)\\s*\\.?(?:\\s|$)', re.IGNORECASE)"

if '.{0,5}$' in pcontent:
    pcontent = pcontent.replace(old_re_pasal, new_re_pasal, 1)
    print("[parsers] Relaxed RE_PASAL regex (removed .{0,5}$ tail restriction)")
else:
    print("[parsers] RE_PASAL already relaxed")

# 2k: Also flush penjelasan at end of parse loop (before return)
old_flush_remaining = '''        # Flush remaining
        self._flush_ayat()
        self._flush_pasal()
        self._flush_quoted_amendment()'''

new_flush_remaining = '''        # Flush remaining
        self._flush_ayat()
        self._flush_pasal()
        self._flush_quoted_amendment()
        self._flush_penjelasan()  # v3: flush any remaining penjelasan text'''

if '_flush_penjelasan()' not in pcontent:
    pcontent = pcontent.replace(old_flush_remaining, new_flush_remaining, 1)
    print("[parsers] Added _flush_penjelasan to end-of-parse flush")
else:
    print("[parsers] End-of-parse penjelasan flush already exists")

# 2l: Update version header
old_header = 'v2 fixes (based on QA assessment):'
new_header = 'v2+v3 fixes (based on QA assessment):'
if 'v3' not in pcontent.split('\n')[0]:
    pcontent = pcontent.replace(old_header, new_header, 1)
    print("[parsers] Updated version header")

# Also update clean_extractor version
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/clean_extractor.py', 'r') as f:
    ce_content = f.read()
old_ce_header = 'v2 fixes (based on QA assessment):'
new_ce_header = 'v2+v3 fixes (based on QA assessment):'
if 'v3' not in ce_content.split('\n')[0]:
    ce_content = ce_content.replace(old_ce_header, new_ce_header, 1)
    print("[clean_extractor] Updated version header")
with open('/home/z/my-project/hukum-aneh/backend/python/chunker/clean_extractor.py', 'w') as f:
    f.write(ce_content)

with open('/home/z/my-project/hukum-aneh/backend/python/chunker/parsers.py', 'w') as f:
    f.write(pcontent)

print("\n--- All parsers.py fixes applied ---")
print("Done!")