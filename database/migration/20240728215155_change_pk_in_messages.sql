-- +goose Up
ALTER TABLE messages DROP CONSTRAINT messages_pkey;
ALTER TABLE messages ADD PRIMARY KEY (chat_id, id);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE messages DROP CONSTRAINT messages_pkey;
ALTER TABLE messages ADD PRIMARY KEY (id);
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
