"""Statute parser with 3-layer defense for Indonesian legal documents.

Layer 1: Text normalization (done in extractor)
Layer 2: Pattern matching (strict regex + font size signals)
Layer 3: Context validation (heuristic checks)

Handles the hierarchy: BAB > Bagian > Paragraf > Pasal > Ayat > Huruf/Angka
"""
import re
from typing import List, Dict, Optional, Tuple

# ======================== REGEX PATTERNS ========================

# PASAL patterns
RE_PASAL_STRICT = re.compile(
    r'^\s*\*?\s*pasal\s+(\d+[a-zA-Z]?)\s*\.?\s*\.{0,5}\s*$',
    re.IGNORECASE
)
RE_PASAL_LOOSE = re.compile(
    r'pasal\s+(\d+[a-zA-Z]?)',
    re.IGNORECASE
)

# BAB patterns
RE_BAB_STRICT = re.compile(
    r'^\s*\*?\s*bab\s+([ivxlcdm]+|\d+)\s*$',
    re.IGNORECASE
)
RE_BAB_LOOSE = re.compile(
    r'bab\s+([ivxlcdm]+|\d+)',
    re.IGNORECASE
)

# BAGIAN patterns
RE_BAGIAN_STRICT = re.compile(
    r'^\s*\*?\s*bagian\s+([ivxlcdm]+|\d+)',
    re.IGNORECASE
)
RE_BAGIAN_LOOSE = re.compile(
    r'bagian\s+([ivxlcdm]+|\d+)',
    re.IGNORECASE
)

# PARAGRAF patterns
RE_PARAGRAF_STRICT = re.compile(
    r'^\s*\*?\s*paragraf\s+([ivxlcdm]+|\d+)',
    re.IGNORECASE
)
RE_PARAGRAF_LOOSE = re.compile(
    r'paragraf\s+([ivxlcdm]+|\d+)',
    re.IGNORECASE
)

# AYAT patterns (1), (2), (2a), etc.
RE_AYAT = re.compile(r'^\s*\(\s*(\d+[a-zA-Z]?)\s*\)')

# HURUF patterns a), b), c)
RE_HURUF = re.compile(r'^\s*([a-z])\)')

# ANGKA patterns 1., 2., 3.
RE_ANGKA = re.compile(r'^\s*(\d+)\.')


# ======================== PARSER ========================

class StatuteParser:
    def __init__(self, lines: List[Dict], stats: Dict):
        self.lines = lines
        self.stats = stats
        self.body_size = stats.get('most_common_size', 12.0)
        self.heading_threshold = self.body_size + 1.5
        self._reset_state()

    def _reset_state(self):
        self.current_bab = None
        self.current_bab_title = None
        self.current_bagian = None
        self.current_bagian_title = None
        self.current_paragraf = None
        self.current_paragraf_title = None
        self.current_pasal = None
        self.current_pasal_title = None
        self.current_ayat = None
        self.pasals = []
        self.issues = []
        self._current_pasal_lines = []
        self._current_ayat_lines = []

    def _add_issue(self, issue_type: str, line: Dict, detail: str):
        self.issues.append({
            'issue_type': issue_type,
            'page': line['page'],
            'line_text': line['normalized'][:100],
            'detail': detail,
        })

    def _is_likely_heading(self, line: Dict) -> bool:
        """Check if a line is likely a heading based on font size and bold."""
        if line['font_size'] > self.heading_threshold:
            return True
        if line['is_bold'] and line['font_size'] >= self.body_size:
            return True
        if len(line['normalized']) < 60 and line['font_size'] > self.body_size:
            return True
        return False

    def _is_inline_reference(self, line: Dict, pattern_match) -> bool:
        """Check if a Pasal/BAB match is an inline reference (not a heading)."""
        text = line['normalized']
        match_start = pattern_match.start()
        match_end = pattern_match.end()
        prefix = text[:match_start].strip()
        suffix = text[match_end:].strip()
        if prefix and len(prefix) > 3:
            return True
        if suffix and len(suffix) > 3:
            return False
        if len(text) > 80:
            return True
        return False

    def parse(self) -> Dict:
        self._reset_state()
        for i, line in enumerate(self.lines):
            norm = line['normalized']
            if not norm:
                continue
            matched = False

            # 1. BAB
            if RE_BAB_STRICT.match(norm):
                m = RE_BAB_STRICT.match(norm)
                self._flush_pasal()
                self.current_bab = m.group(1).upper()
                self.current_bab_title = None
                matched = True
            elif RE_BAB_LOOSE.search(norm) and self._is_likely_heading(line):
                m = RE_BAB_LOOSE.search(norm)
                if not self._is_inline_reference(line, m):
                    self._flush_pasal()
                    self.current_bab = m.group(1).upper()
                    self.current_bab_title = norm
                    self._add_issue('BAB_LOOSE_MATCH', line, 'BAB matched via loose regex')
                    matched = True

            # 2. BAGIAN
            if not matched and RE_BAGIAN_STRICT.match(norm):
                m = RE_BAGIAN_STRICT.match(norm)
                self._flush_pasal()
                self.current_bagian = m.group(1).upper()
                self.current_bagian_title = None
                matched = True
            elif not matched and RE_BAGIAN_LOOSE.search(norm) and self._is_likely_heading(line):
                m = RE_BAGIAN_LOOSE.search(norm)
                if not self._is_inline_reference(line, m):
                    self._flush_pasal()
                    self.current_bagian = m.group(1).upper()
                    self.current_bagian_title = norm
                    self._add_issue('BAGIAN_LOOSE_MATCH', line, 'Bagian matched via loose regex')
                    matched = True

            # 3. PARAGRAF
            if not matched and RE_PARAGRAF_STRICT.match(norm):
                m = RE_PARAGRAF_STRICT.match(norm)
                self._flush_pasal()
                self.current_paragraf = m.group(1).upper()
                self.current_paragraf_title = None
                matched = True
            elif not matched and RE_PARAGRAF_LOOSE.search(norm) and self._is_likely_heading(line):
                m = RE_PARAGRAF_LOOSE.search(norm)
                if not self._is_inline_reference(line, m):
                    self._flush_pasal()
                    self.current_paragraf = m.group(1).upper()
                    self.current_paragraf_title = norm
                    self._add_issue('PARAGRAF_LOOSE_MATCH', line, 'Paragraf matched via loose regex')
                    matched = True

            # 4. PASAL
            if not matched:
                pasal_match = RE_PASAL_STRICT.match(norm)
                if pasal_match:
                    self._flush_pasal()
                    self.current_pasal = pasal_match.group(1)
                    self.current_pasal_title = None
                    self._current_pasal_lines = []
                    matched = True
                else:
                    pasal_loose = RE_PASAL_LOOSE.search(norm)
                    if pasal_loose and self._is_likely_heading(line):
                        if not self._is_inline_reference(line, pasal_loose):
                            self._flush_pasal()
                            self.current_pasal = pasal_loose.group(1)
                            self.current_pasal_title = norm
                            self._current_pasal_lines = []
                            self._add_issue('PASAL_LOOSE_MATCH', line, 'Pasal matched via loose regex')
                            matched = True

            # 5. AYAT
            if not matched and self.current_pasal is not None:
                ayat_match = RE_AYAT.match(norm)
                if ayat_match:
                    self._flush_ayat()
                    self.current_ayat = ayat_match.group(1)
                    self._current_ayat_lines = [norm]
                    matched = True

            # 6. Accumulate content
            if not matched and self.current_pasal is not None:
                self._current_ayat_lines.append(norm)

        self._flush_ayat()
        self._flush_pasal()
        return {
            'pasals': self.pasals,
            'issues': self.issues,
            'stats': {
                'total_pasals': len(self.pasals),
                'total_ayats': sum(len(p['ayats']) for p in self.pasals),
                'total_issues': len(self.issues),
            }
        }

    def _flush_ayat(self):
        if self.current_ayat is not None and self._current_ayat_lines:
            pasal_content = '\n'.join(self._current_ayat_lines).strip()
            self._current_pasal_lines.append({
                'ayat_num': self.current_ayat,
                'content': pasal_content,
            })

    def _flush_pasal(self):
        if self.current_pasal is not None:
            self.pasals.append({
                'pasal_num': self.current_pasal,
                'bab': self.current_bab,
                'bab_title': self.current_bab_title,
                'bagian': self.current_bagian,
                'bagian_title': self.current_bagian_title,
                'paragraf': self.current_paragraf,
                'paragraf_title': self.current_paragraf_title,
                'title': self.current_pasal_title,
                'ayats': self._current_pasal_lines,
            })

    def get_hierarchy_path(self, pasal: Dict) -> str:
        parts = []
        if pasal.get('bab'):
            parts.append(f'BAB {pasal["bab"]}')
        if pasal.get('bagian'):
            parts.append(f'Bagian {pasal["bagian"]}')
        if pasal.get('paragraf'):
            parts.append(f'Paragraf {pasal["paragraf"]}')
        parts.append(f'Pasal {pasal["pasal_num"]}')
        return ' > '.join(parts) if parts else ''


if __name__ == '__main__':
    import sys
    import json
    from extractor import extract_lines_for_parsing
    path = sys.argv[1] if len(sys.argv) > 1 else '/home/z/my-project/download/samples/uud-1945/uud_1945.pdf'
    lines, stats = extract_lines_for_parsing(path)
    parser = StatuteParser(lines, stats)
    result = parser.parse()
    print(f'File: {path.split("/")[-1]}')
    print(f'Pasals found: {result["stats"]["total_pasals"]}')
    print(f'Ayats found: {result["stats"]["total_ayats"]}')
    print(f'Issues: {result["stats"]["total_issues"]}')
    if result['issues']:
        print(f'\nIssues:')
        for iss in result['issues'][:20]:
            print(f'  [{iss["issue_type"]}] p{iss["page"]}: {iss["line_text"][:60]}')
            print(f'    -> {iss["detail"]}')
    print(f'\nFirst 5 Pasals:')
    for p in result['pasals'][:5]:
        print(f'  Pasal {p["pasal_num"]} (BAB {p["bab"]}, {len(p["ayats"])} ayats)')
        for a in p['ayats'][:3]:
            print(f'    ({a["ayat_num"]}) {a["content"][:80]}...')
