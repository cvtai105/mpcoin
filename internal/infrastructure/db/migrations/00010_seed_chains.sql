-- +goose Up
-- +goose StatementBegin
-- Insert into chains table
INSERT INTO chains (id, name, chain_id, rpc_url, ws_url, last_scan_block_number, is_active) 
VALUES (
    '2773fa12-645a-45d0-80a2-79cf5a2ecf96', 
    'Sepolia', 
    '11155111', 
    'https://sepolia.infura.io/v3/6c89fb7fa351451f939eea9da6bee755', 
    'wss://sepolia.infura.io/ws/v3/6d3cfcab0e3a442eb3c890ae4422f76d', 
    -1,
    TRUE
);
INSERT INTO chains (id, name, chain_id, rpc_url, ws_url, last_scan_block_number, is_active) 
VALUES (
    '2773fa12-645a-45d0-80a2-79cf5a2ecf97', 
    'Linea Sepolia', 
    '59141', 
    'https://linea-sepolia.infura.io/v3/6c89fb7fa351451f939eea9da6bee755', 
    'wss://linea-sepolia.infura.io/ws/v3/6d3cfcab0e3a442eb3c890ae4422f76d', 
    -1,
    TRUE
);

-- Insert into tokens table
INSERT INTO tokens (id, chain_id, name, symbol, decimals, contract_address) 
VALUES (
    '2773fa12-645a-45d0-80a2-79cf5a2ecf98', 
    '2773fa12-645a-45d0-80a2-79cf5a2ecf96', 
    'SepoliaETH', 
    'ETH', 
    18, 
    '0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43'
);

INSERT INTO tokens (id, chain_id, name, symbol, decimals, contract_address) 
VALUES (
    '2773fa12-645a-45d0-80a2-79cf5a2ecf99', 
    '2773fa12-645a-45d0-80a2-79cf5a2ecf97', 
    'LineaETH', 
    'ETH', 
    18, 
    '0xe1a12515F9AB2764b887bF60B923Ca494EBbB2d6'
);

-- Update native_token_id in chains table
UPDATE chains 
SET native_token_id = '2773fa12-645a-45d0-80a2-79cf5a2ecf98' 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf96';

UPDATE chains 
SET native_token_id = '2773fa12-645a-45d0-80a2-79cf5a2ecf99' 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf97';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Revert the update of native_token_id in chains table
UPDATE chains 
SET native_token_id = NULL 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf96';
UPDATE chains 
SET native_token_id = NULL 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf97';

-- Delete from tokens table
DELETE FROM tokens 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf98';
DELETE FROM tokens 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf99';

-- Delete from chains table
DELETE FROM chains 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf96';
DELETE FROM chains 
WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf97';
-- +goose StatementEnd
