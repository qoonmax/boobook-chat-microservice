-- +goose Up
CREATE TABLE messages
(
    id            SERIAL PRIMARY KEY,
    chat_id       INTEGER REFERENCES chats (id) ON DELETE CASCADE,
    user_id       INTEGER NOT NULL,
    body          TEXT NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE messages;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
