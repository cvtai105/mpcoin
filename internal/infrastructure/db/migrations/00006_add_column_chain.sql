-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE chains 
ADD COLUMN native_token_id UUID NULL;
ALTER TABLE chains 
ADD CONSTRAINT fk_chains_native_token_id FOREIGN KEY (native_token_id) REFERENCES tokens(id);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
ALTER TABLE chains 
DROP CONSTRAINT fk_chains_native_token_id;
ALTER TABLE chains 
DROP COLUMN native_token_id;
