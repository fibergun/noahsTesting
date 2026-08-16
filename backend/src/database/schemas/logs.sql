CREATE TABLE IF NOT EXISTS logs (
                                    log_id INTEGER PRIMARY KEY AUTOINCREMENT,
                                    task_id INTEGER NOT NULL,
                                    user_id INTEGER NOT NULL,
                                    completed BOOLEAN NOT NULL DEFAULT FALSE,
                                    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                    UNIQUE(task_id, user_id)
);