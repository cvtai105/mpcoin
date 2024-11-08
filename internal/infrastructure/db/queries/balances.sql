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
