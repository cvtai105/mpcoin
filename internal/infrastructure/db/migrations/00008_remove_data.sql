-- +goose Up
-- +goose StatementBegin
-- Delete tokens related to the specified chains
DELETE FROM tokens 
WHERE chain_id IN (
    SELECT id 
    FROM chains 
    WHERE chain_id IN ('11511', '84532', '421613')
);

-- Delete the chains
DELETE FROM chains 
WHERE chain_id IN ('11511', '84532', '421613');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Re-insert chains
INSERT INTO chains (id, name, chain_id, rpc_url, native_currency) VALUES
(uuid_generate_v4(), 'Sepolia', '11511', 'https://eth-sepolia.g.alchemy.com/v2/demo', 'ETH'),
(uuid_generate_v4(), 'Base Sepolia', '84532', 'https://sepolia.base.org', 'ETH'),
(uuid_generate_v4(), 'Arbitrum Sepolia', '421613', 'https://sepolia.arbitrum.io/rpc', 'ETH');

-- Re-insert tokens
INSERT INTO tokens (id, chain_id, contract_address, name, symbol, decimals)
SELECT 
    uuid_generate_v4(),
    chains.id,
    '0x1',
    'Ethereum',
    'ETH',
    18
FROM chains
WHERE chains.chain_id IN ('11511', '84532', '421613');
-- +goose StatementEnd
