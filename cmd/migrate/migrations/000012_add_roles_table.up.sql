CREATE TABLE IF NOT EXISTS roles (
    ID BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    level int NOT NULL DEFAULT 0,
    description TEXT
);

INSERT INTO roles (name, description, level) VALUES
('admin', 'Administrator role with all permissions', 3),
('user', 'A user can create posts and comments', 1),
('moderator', 'A moderator can update other user posts', 2);