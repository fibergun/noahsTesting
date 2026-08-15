CREATE TABLE IF NOT EXISTS tasks (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     task TEXT NOT NULL,
                                     group_id INTEGER NOT NULL,
                                     user_id INTEGER NOT NULL,
                                     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                    UNIQUE(group_id, task)
);