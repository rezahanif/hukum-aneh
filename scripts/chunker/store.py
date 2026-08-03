import sqlite3, os, json, sys

DB_PATH = "/home/z/my-project/download/chunker_test.db"


def init_db(path=DB_PATH):
    """Create SQLite schema (mirrors planned PostgreSQL schema)."""
    conn = sqlite3.connect(path)
    c = conn.cursor()
    
    c.executescript("""
    CREATE TABLE IF NOT EXISTS documents (
        id            INTEGER PRIMARY KEY AUTOINCREMENT,
        folder        TEXT NOT NULL,
        filename      TEXT NOT NULL,
        doc_type      TEXT NOT NULL,
        title         TEXT,
        reg_number    TEXT,
        reg_year      INT,
        raw_text      TEXT,
        file_hash     TEXT,
        pasal_count   INT,
        ayat_count    INT,
        issue_count   INT,
        created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS chunks (
        id              INTEGER PRIMARY KEY AUTOINCREMENT,
        document_id     INTEGER NOT NULL REFERENCES documents(id),
        parent_chunk_id INTEGER REFERENCES chunks(id),
        chunk_level     TEXT NOT NULL,  -- 'preamble', 'bab', 'pasal', 'ayat'
        hierarchy_path  TEXT,         -- 'BAB II > Pasal 5'
        bab_num         TEXT,
        bab_title       TEXT,
        pasal_num       TEXT,
        ayat_num        INT,
        content         TEXT NOT NULL,
        token_count     INT,
        metadata        TEXT,         -- JSON
        created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS parse_issues (
        id            INTEGER PRIMARY KEY AUTOINCREMENT,
        document_id   INTEGER REFERENCES documents(id),
        issue_type    TEXT,
        pasal         TEXT,
        detail        TEXT,
        line_index    INT,
        raw_text      TEXT,
        resolved      BOOLEAN DEFAULT 0
    );

    CREATE INDEX IF NOT EXISTS idx_chunks_doc ON chunks(document_id);
    CREATE INDEX IF NOT EXISTS idx_chunks_level ON chunks(chunk_level);
    CREATE INDEX IF NOT EXISTS idx_chunks_pasal ON chunks(pasal_num);
    """)
    conn.commit()
    return conn


def rough_token_count(text):
    """Rough token estimate: ~4 chars per token for Indonesian text."""
    return len(text) // 4


def store_parsed_document(conn, parsed_doc, folder, raw_text, file_hash, reg_number="", reg_year=None):
    """Store a parsed document and all its chunks into SQLite."""
    c = conn.cursor()
    
    # Count totals
    total_pasals = len(parsed_doc.all_pasals)
    total_ayat = sum(len(p.ayat_list) for p in parsed_doc.all_pasals)
    total_issues = len(parsed_doc.issues)
    
    # Insert document
    c.execute("""
        INSERT INTO documents (folder, filename, doc_type, title, reg_number, reg_year, 
                           raw_text, file_hash, pasal_count, ayat_count, issue_count)
        VALUES (?, ?, 'statute', ?, ?, ?, ?, ?, ?, ?, ?)
    """, (folder, parsed_doc.filename, parsed_doc.title, reg_number, reg_year,
          raw_text, file_hash, total_pasals, total_ayat, total_issues))
    
    doc_id = c.lastrowid
    
    # Insert preamble chunk
    if parsed_doc.preamble_text:
        c.execute("""
            INSERT INTO chunks (document_id, chunk_level, hierarchy_path, content, token_count)
            VALUES (?, ?, ?, ?)
        """, (doc_id, 'preamble', 'Preamble', parsed_doc.preamble_text, rough_token_count(parsed_doc.preamble_text)))
    
    # Insert BAB chunks (for hierarchy context)
    bab_for_pasal = {}  # map pasal_number -> bab info
    for bab in parsed_doc.babs:
        bab_for_pasal.update({p.number: (bab.number, bab.title) for p in bab.pasal_list})
    
    # Insert Pasal (parent) and Ayat (child) chunks
    for pasal in parsed_doc.all_pasals:
        bab_info = bab_for_pasal.get(pasal.number, (None, None))
        bab_num, bab_title = bab_info
        
        hier_path = f"BAB {bab_num}" if bab_num else ""
        if bab_title:
            hier_path += f" > {bab_title}"
        hier_path += f" > Pasal {pasal.number}" if hier_path else f"Pasal {pasal.number}"
        
        # Parent chunk: full Pasal
        c.execute("""
            INSERT INTO chunks (document_id, chunk_level, hierarchy_path, 
                               bab_num, bab_title, pasal_num, content, token_count)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)
        """, (doc_id, 'pasal', hier_path, bab_num, bab_title, pasal.number,
              pasal.raw_text, rough_token_count(pasal.raw_text)))
        pasal_chunk_id = c.lastrowid
        
        # Child chunks: per Ayat
        for ayat in pasal.ayat_list:
            ayat_hier = f"{hier_path} > Ayat ({ayat.number})"
            c.execute("""
                INSERT INTO chunks (document_id, parent_chunk_id, chunk_level, hierarchy_path,
                                   bab_num, bab_title, pasal_num, ayat_num, content, token_count)
                VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            """, (doc_id, pasal_chunk_id, 'ayat', ayat_hier, bab_num, bab_title,
                  pasal.number, ayat.number, ayat.text,
                  rough_token_count(ayat.text)))
    
    # Insert parse issues
    for issue in parsed_doc.issues:
        c.execute("""
            INSERT INTO parse_issues (document_id, issue_type, pasal, detail, line_index, raw_text)
            VALUES (?, ?, ?, ?, ?, ?)
        """, (doc_id, issue["type"], issue.get("pasal", ""), 
              json.dumps(issue, ensure_ascii=False), issue.get("line"), ""))
    
    conn.commit()
    return doc_id


def print_summary(conn):
    """Print a summary of what's in the database."""
    c = conn.cursor()
    
    docs = c.execute("SELECT id, filename, pasal_count, ayat_count, issue_count FROM documents").fetchall()
    print(f"\n{'='*60}")
    print(f"Documents: {len(docs)}")
    for d in docs:
        print(f"  [{d[0]}] {d[1]}: {d[2]} pasal, {d[3]} ayat, {d[4]} issues")
    
    chunks = c.execute("SELECT chunk_level, COUNT(*) FROM chunks GROUP BY chunk_level").fetchall()
    print(f"\nChunks by level:")
    for level, count in chunks:
        print(f"  {level:12s}: {count}")
    
    issues = c.execute("SELECT COUNT(*) FROM parse_issues").fetchone()[0]
    print(f"\nParse issues: {issues}")
    
    print(f"{'='*60}")
