import json
d = json.load(open('/home/z/my-project/download/batch_test_results.json'))
for r in d:
    stats = r.get('extraction_stats') or {}
    pout = r.get('parser_output') or {}
    errs = r.get('errors', [])
    err_str = errs[0][:80] if errs else 'none'
    folder = r['folder']
    sizes = str(stats.get('unique_sizes', []))[:50]
    pasals = str(pout.get('total_pasals', 'ERR'))
    ayats = str(pout.get('total_ayats', 'ERR'))
    flaws = len(r['flaws'])
    print(f'{folder:20s} | sizes={sizes:50s} | P={pasals:>4s} | A={ayats:>4s} | flaws={flaws:3d} | err={err_str}')