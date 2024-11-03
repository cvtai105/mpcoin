-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE transactions 
ALTER COLUMN wallet_id SET DATA TYPE UUID,
ALTER COLUMN wallet_id DROP NOT NULL;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
ALTER TABLE transactions 
ALTER COLUMN wallet_id SET DATA TYPE UUID,
ALTER COLUMN wallet_id SET NOT NULL;
