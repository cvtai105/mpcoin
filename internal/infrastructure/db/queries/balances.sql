-- name: GetBalancesByWalletId :many
SELECT 
    b.balance,
    t.id AS token_id,
    t.name AS token_name,
    t.symbol AS token_symbol,
    t.decimals,
    b.updated_at
FROM 
    balances b
JOIN 
    tokens t ON b.token_id = t.id
WHERE 
    b.wallet_id = $1;
