-- +goose Up
CREATE TABLE chats
(
    id            SERIAL PRIMARY KEY,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE chats;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
