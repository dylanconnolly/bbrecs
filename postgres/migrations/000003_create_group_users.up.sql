CREATE TABLE IF NOT EXISTS group_users(
    group_id uuid,
    user_id uuid,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);