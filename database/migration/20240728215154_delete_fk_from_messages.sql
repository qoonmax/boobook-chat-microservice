-- +goose Up
ALTER TABLE messages DROP CONSTRAINT messages_chat_id_fkey;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE messages ADD CONSTRAINT messages_chat_id_fkey FOREIGN KEY (chat_id) REFERENCES chats(id);
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
