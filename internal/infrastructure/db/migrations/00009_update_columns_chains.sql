-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE chains ADD COLUMN is_active BOOLEAN NULL;
ALTER TABLE users ADD COLUMN name VARCHAR(255) NULL;
ALTER TABLE users ADD COLUMN avatar_url VARCHAR(255) NULL;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
ALTER TABLE chains DROP COLUMN is_active;
ALTER TABLE users DROP COLUMN name;
ALTER TABLE users DROP COLUMN avatar_url;
