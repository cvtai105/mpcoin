-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE chains ADD COLUMN native_token_id UUID NULL;
ALTER TABLE chains ADD CONSTRAINT fk_chains_native_token_id FOREIGN KEY (native_token_id) REFERENCES tokens(id);
ALTER TABLE chains ADD COLUMN ws_url VARCHAR(255) NULL;
ALTER TABLE chains ADD COLUMN last_scan_block_number BIGINT NULL;
ALTER TABLE chains ALTER COLUMN native_currency DROP NOT NULL;
ALTER TABLE transactions ALTER COLUMN from_address TYPE VARCHAR(42);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
ALTER TABLE chains DROP CONSTRAINT fk_chains_native_token_id;
ALTER TABLE chains DROP COLUMN native_token_id;
ALTER TABLE chains DROP COLUMN ws_url;
ALTER TABLE chains DROP COLUMN last_scan_block_number;
ALTER TABLE chains ALTER COLUMN native_currency SET NOT NULL;
ALTER TABLE transactions ALTER COLUMN from_address TYPE VARCHAR(255);
