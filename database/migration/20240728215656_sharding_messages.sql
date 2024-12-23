-- +goose Up
SELECT create_distributed_table('messages', 'chat_id');
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
SELECT drop_distributed_table('messages');
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
