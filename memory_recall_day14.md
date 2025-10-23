# Day 14 - Notes Table Setup

- A migration defines table structure using SQL.
- SERIAL = auto-incrementing integer.
- TEXT = variable-length string.
- TIMESTAMPTZ = timestamp with timezone (safe for UTC).
- DEFAULT NOW() auto-fills creation/update times.
- /d  notes shows structure; SELECT * FROM notes shows data in table.
- migrations/ folder stores SQL files for version control. 