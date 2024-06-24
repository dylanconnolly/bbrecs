CREATE TABLE IF NOT EXISTS group_users(
    id serial PRIMARY KEY,
    group_id uuid REFERENCES groups(id) ON DELETE CASCADE,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    created_at timestamp NOT NULL DEFAULT current_timestamp
);