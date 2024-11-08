-- name: GetBalancesByUserId :many
SELECT 
    b.balance,
    t.*
FROM 
    balances b
JOIN 
    wallets w ON b.wallet_id = w.id
JOIN 
    tokens t ON b.token_id = t.id
JOIN 
    users u ON w.user_id = u.id
WHERE 
    u.id = $1;

-- name: UpdateBalance :one
UPDATE balances
SET balance = $1
FROM wallets
WHERE balances.wallet_id = wallets.id
  AND wallets.address = $2 
  AND balances.token_id = $3
RETURNING *;

