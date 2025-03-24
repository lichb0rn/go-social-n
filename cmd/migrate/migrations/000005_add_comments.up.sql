CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL PRIMARY KEY,
    post_id BIGSERIAL NOT NULL REFERENCES posts (id),
    user_id BIGSERIAL NOT NULL REFERENCES users (id),
    content TEXT NOT NULL,
    created_at TIMESTAMP(0)
    with
        time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
