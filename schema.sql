CREATE TABLE emoji_counter(
  emoji TEXT PRIMARY KEY,
  click_count INTEGER NOT NULL DEFAULT 0
);
CREATE TABLE slide_state(
  id INTEGER PRIMARY KEY CHECK(id = 1),
  current_index INTEGER NOT NULL DEFAULT 0,
  updated_at TEXT NOT NULL DEFAULT(strftime('%Y-%m-%dT%H:%M:%fZ','now'))
);
