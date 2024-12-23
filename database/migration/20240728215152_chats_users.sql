-- +goose Up
CREATE TABLE chats_users
(
    id            SERIAL PRIMARY KEY,
    chat_id       INTEGER REFERENCES chats (id) ON DELETE CASCADE,
    user_id       INTEGER NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE chats_users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
