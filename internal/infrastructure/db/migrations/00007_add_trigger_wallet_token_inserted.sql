-- +goose Up
-- Up migration
-- +goose StatementBegin
-- Function to add a balance entry for all wallets when a token is inserted
CREATE OR REPLACE FUNCTION create_balance_for_all_wallets()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO balances (id, wallet_id, chain_id, token_id, balance, updated_at)
    SELECT gen_random_uuid(), w.id, NEW.chain_id, NEW.id, 0, NOW()
    FROM wallets w;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call the function after a token is inserted
CREATE TRIGGER after_token_insert
AFTER INSERT ON tokens
FOR EACH ROW
EXECUTE FUNCTION create_balance_for_all_wallets();


-- Function to add a balance entry for all tokens when a wallet is inserted
CREATE OR REPLACE FUNCTION create_balance_for_all_tokens()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO balances (id, wallet_id, chain_id, token_id, balance, updated_at)
    SELECT gen_random_uuid(), NEW.id, t.chain_id, t.id, 0, NOW()
    FROM tokens t;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call the function after a wallet is inserted
CREATE TRIGGER after_wallet_insert
AFTER INSERT ON wallets
FOR EACH ROW
EXECUTE FUNCTION create_balance_for_all_tokens();
-- +goose StatementEnd
-- +goose Down
-- SQL script to drop the triggers
DROP TRIGGER IF EXISTS after_token_insert ON tokens;
DROP FUNCTION IF EXISTS create_balance_for_all_wallets();
DROP TRIGGER IF EXISTS after_wallet_insert ON wallets;
DROP FUNCTION IF EXISTS create_balance_for_all_tokens();


