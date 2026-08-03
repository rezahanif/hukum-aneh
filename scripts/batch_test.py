"""Batch test: run extractor + parser on 1 sample from each folder, collect flaws."""
import os
import sys
import json
from datetime import datetime

sys.path.insert(0, "/home/z/my-project/scripts")
from extractor import extract_lines_for_parsing
from parser_statute import StatuteParser

SAMPLES_DIR = "/home/z/my-project/download/samples"
OUTPUT_PATH = "/home/z/my-project/download/batch_test_results.json"

# Map folder names to doc types for context
def get_doc_type(folder_name):
    if folder_name == 'uud-1945':
        return 'constitution'
    elif folder_name == 'Putusan-MK':
        return 'court_ruling'
    elif folder_name.startswith('JDIH_'):
        return 'ministerial_regulation'
    elif folder_name in ('uu', 'pp', 'perppu', 'perpres', 'perda', 'inpres', 'keppres', 'tap_mpr', 'peraturan'):
        return 'statute'
    return 'unknown'


def find_sample_pdfs():
    """Find one PDF per subfolder in samples dir."""
    samples = {}
    for folder in sorted(os.listdir(SAMPLES_DIR)):
        folder_path = os.path.join(SAMPLES_DIR, folder)
        if not os.path.isdir(folder_path):
            continue
        pdfs = [f for f in os.listdir(folder_path) if f.lower().endswith('.pdf')]
        if pdfs:
            # pick first one
            samples[folder] = os.path.join(folder_path, pdfs[0])
    return samples


def test_one(folder, pdf_path):
    """Test parser on one PDF, return results dict."""
    result = {
        'folder': folder,
        'doc_type': get_doc_type(folder),
        'filename': os.path.basename(pdf_path),
        'errors': [],
        'flaws': [],
        'parser_output': None,
    }
    try:
        lines, stats = extract_lines_for_parsing(pdf_path)
        result['extraction_stats'] = {
            'total_lines': stats['total_lines'],
            'unique_sizes': stats['unique_sizes'],
            'most_common_size': stats['most_common_size'],
            'body_sizes': stats['body_sizes'],
            'bold_sizes': stats['bold_sizes'],
            'size_distribution': stats['size_distribution'],
        }

        parser = StatuteParser(lines, stats)
        parsed = parser.parse()
        result['parser_output'] = parsed['stats']

        # ---- FLAW DETECTION ----

        # FLAW 1: Duplicate pasal numbers
        pasal_nums = [p['pasal_num'] for p in parsed['pasals']]
        seen = set()
        dupes = []
        for pn in pasal_nums:
            if pn in seen:
                dupes.append(pn)
            seen.add(pn)
        if dupes:
            result['flaws'].append({
                'flaw': 'DUPLICATE_PASAL',
                'detail': f'Duplicate pasal numbers: {list(set(dupes))}',
                'severity': 'high',
            })

        # FLAW 2: Ayat content starts with the ayat marker itself
        for p in parsed['pasals']:
            for a in p['ayats']:
                content = a['content']
                # Check if content starts with the ayat number repeated
                if content.startswith(f"({a['ayat_num']})"):
                    result['flaws'].append({
                        'flaw': 'AYAT_MARKER_IN_CONTENT',
                        'detail': f'Pasal {p["pasal_num"]} ayat ({a["ayat_num"]}) content starts with ayat marker',
                        'severity': 'medium',
                    })
                    break  # one example per pasal is enough

        # FLAW 3: Pasal with 0 ayats (might be a false positive or content not captured)
        for p in parsed['pasals']:
            if len(p['ayats']) == 0:
                result['flaws'].append({
                    'flaw': 'EMPTY_PASAL',
                    'detail': f'Pasal {p["pasal_num"]} has 0 ayats',
                    'severity': 'medium',
                })

        # FLAW 4: Font size gap analysis - if only 1-2 sizes, font signal is weak
        if len(stats['unique_sizes']) <= 2:
            result['flaws'].append({
                'flaw': 'WEAK_FONT_SIGNAL',
                'detail': f'Only {len(stats["unique_sizes"])} font sizes detected ({stats["unique_sizes"]}), heading detection relies almost entirely on regex',
                'severity': 'low',
            })

        # FLAW 5: Too many loose matches (parsing instability)
        loose_matches = [i for i in parsed['issues'] if 'LOOSE_MATCH' in i['issue_type']]
        if len(loose_matches) > len(parsed['pasals']) * 0.3:
            result['flaws'].append({
                'flaw': 'HIGH_LOOSE_MATCH_RATIO',
                'detail': f'{len(loose_matches)} loose matches out of {len(parsed["pasals"])} pasals ({len(loose_matches)/max(len(parsed["pasals"]),1)*100:.0f}%)',
                'severity': 'medium',
            })

        # FLAW 6: Very long lines in ayat content (possible paragraph bleed)
        for p in parsed['pasals'][:10]:  # check first 10
            for a in p['ayats']:
                if len(a['content']) > 500:
                    result['flaws'].append({
                        'flaw': 'LONG_AYAT_CONTENT',
                        'detail': f'Pasal {p["pasal_num"]} ayat ({a["ayat_num"]}) has {len(a["content"])} chars (possible paragraph bleed)',
                        'severity': 'low',
                    })
                    break

        # FLAW 7: Missing BAB assignment
        pasals_no_bab = [p for p in parsed['pasals'] if p['bab'] is None]
        if pasals_no_bab and parsed['pasals']:
            result['flaws'].append({
                'flaw': 'PASALS_WITHOUT_BAB',
                'detail': f'{len(pasals_no_bab)}/{len(parsed["pasals"])} pasals have no BAB assigned',
                'severity': 'low' if len(pasals_no_bab) < 3 else 'medium',
            })

        # Collect all parser issues as info
        if parsed['issues']:
            result['parser_issues'] = [
                {k: v for k, v in iss.items()}
                for iss in parsed['issues'][:10]
            ]

    except Exception as e:
        result['errors'].append(str(e))
        import traceback
        result['traceback'] = traceback.format_exc()

    return result


def main():
    samples = find_sample_pdfs()
    print(f'Found {len(samples)} sample folders to test')
    print('=' * 80)

    all_results = []
    for folder, path in samples.items():
        print(f'\n>>> Testing: {folder} ({os.path.basename(path)})')
        result = test_one(folder, path)
        all_results.append(result)

        if result['errors']:
            print(f'  ERROR: {result["errors"][0][:100]}')
        else:
            stats = result['parser_output']
            print(f'  Lines: {result["extraction_stats"]["total_lines"]} | '
                  f'Font sizes: {result["extraction_stats"]["unique_sizes"]} | '
                  f'Body: {result["extraction_stats"]["body_sizes"]}')
            print(f'  Pasals: {stats["total_pasals"]} | Ayats: {stats["total_ayats"]} | Issues: {stats["total_issues"]}')
            if result['flaws']:
                for f in result['flaws']:
                    print(f'  FLAW [{f["severity"]}] {f["flaw"]}: {f["detail"][:100]}')
            else:
                print(f'  No flaws detected')

    # Save full results
    with open(OUTPUT_PATH, 'w') as f:
        json.dump(all_results, f, indent=2, ensure_ascii=False, default=str)
    print(f'\nFull results saved to {OUTPUT_PATH}')

    # Summary
    print(f'\n{"=" * 80}')
    print('SUMMARY')
    print(f'{"=" * 80}')
    total_flaws = 0
    total_errors = 0
    for r in all_results:
        n_flaws = len(r['flaws'])
        n_errors = len(r['errors'])
        total_flaws += n_flaws
        total_errors += n_errors
        status = 'OK' if n_flaws == 0 and n_errors == 0 else f'{n_flaws} flaws' + (f' + ERROR' if n_errors else '')
        print(f'  {r["folder"]:20s} | {status}')
    print(f'\n  Total: {len(all_results)} folders, {total_flaws} flaws, {total_errors} errors')


if __name__ == '__main__':
    main()
