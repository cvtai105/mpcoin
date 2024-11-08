// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transactions.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (id, wallet_id , chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, wallet_id, chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status, tx_hash, created_at, updated_at, from_address
`

type CreateTransactionParams struct {
	ID        pgtype.UUID
	WalletID  pgtype.UUID
	ChainID   pgtype.UUID
	ToAddress string
	Amount    string
	TokenID   pgtype.UUID
	GasPrice  pgtype.Text
	GasLimit  pgtype.Text
	Nonce     pgtype.Int8
	Status    string
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.ID,
		arg.WalletID,
		arg.ChainID,
		arg.ToAddress,
		arg.Amount,
		arg.TokenID,
		arg.GasPrice,
		arg.GasLimit,
		arg.Nonce,
		arg.Status,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.ChainID,
		&i.ToAddress,
		&i.Amount,
		&i.TokenID,
		&i.GasPrice,
		&i.GasLimit,
		&i.Nonce,
		&i.Status,
		&i.TxHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FromAddress,
	)
	return i, err
}

const getPaginatedTransactions = `-- name: GetPaginatedTransactions :many
SELECT id, wallet_id, chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status, tx_hash, created_at, updated_at, from_address
FROM transactions
WHERE from_address = $1
   OR to_address = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type GetPaginatedTransactionsParams struct {
	FromAddress pgtype.Text
	Limit       int32
	Offset      int32
}

func (q *Queries) GetPaginatedTransactions(ctx context.Context, arg GetPaginatedTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getPaginatedTransactions, arg.FromAddress, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.WalletID,
			&i.ChainID,
			&i.ToAddress,
			&i.Amount,
			&i.TokenID,
			&i.GasPrice,
			&i.GasLimit,
			&i.Nonce,
			&i.Status,
			&i.TxHash,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FromAddress,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, wallet_id, chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status, tx_hash, created_at, updated_at, from_address FROM transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id pgtype.UUID) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.ChainID,
		&i.ToAddress,
		&i.Amount,
		&i.TokenID,
		&i.GasPrice,
		&i.GasLimit,
		&i.Nonce,
		&i.Status,
		&i.TxHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FromAddress,
	)
	return i, err
}

const getTransactionsByWalletID = `-- name: GetTransactionsByWalletID :many
SELECT id, wallet_id, chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status, tx_hash, created_at, updated_at, from_address FROM transactions
WHERE wallet_id = $1
`

func (q *Queries) GetTransactionsByWalletID(ctx context.Context, walletID pgtype.UUID) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsByWalletID, walletID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.WalletID,
			&i.ChainID,
			&i.ToAddress,
			&i.Amount,
			&i.TokenID,
			&i.GasPrice,
			&i.GasLimit,
			&i.Nonce,
			&i.Status,
			&i.TxHash,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FromAddress,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertSetteledTransaction = `-- name: InsertSetteledTransaction :one
INSERT INTO transactions (id, wallet_id , chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status, tx_hash, from_address)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, wallet_id, chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status, tx_hash, created_at, updated_at, from_address
`

type InsertSetteledTransactionParams struct {
	ID          pgtype.UUID
	WalletID    pgtype.UUID
	ChainID     pgtype.UUID
	ToAddress   string
	Amount      string
	TokenID     pgtype.UUID
	GasPrice    pgtype.Text
	GasLimit    pgtype.Text
	Nonce       pgtype.Int8
	Status      string
	TxHash      pgtype.Text
	FromAddress pgtype.Text
}

func (q *Queries) InsertSetteledTransaction(ctx context.Context, arg InsertSetteledTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, insertSetteledTransaction,
		arg.ID,
		arg.WalletID,
		arg.ChainID,
		arg.ToAddress,
		arg.Amount,
		arg.TokenID,
		arg.GasPrice,
		arg.GasLimit,
		arg.Nonce,
		arg.Status,
		arg.TxHash,
		arg.FromAddress,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.ChainID,
		&i.ToAddress,
		&i.Amount,
		&i.TokenID,
		&i.GasPrice,
		&i.GasLimit,
		&i.Nonce,
		&i.Status,
		&i.TxHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FromAddress,
	)
	return i, err
}

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transactions 
SET (status, tx_hash, gas_price, gas_limit, nonce) = ($2, $3, $4, $5, $6)
WHERE id = $1
RETURNING id, wallet_id, chain_id, to_address, amount, token_id, gas_price, gas_limit, nonce, status, tx_hash, created_at, updated_at, from_address
`

type UpdateTransactionParams struct {
	ID       pgtype.UUID
	Status   string
	TxHash   pgtype.Text
	GasPrice pgtype.Text
	GasLimit pgtype.Text
	Nonce    pgtype.Int8
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransaction,
		arg.ID,
		arg.Status,
		arg.TxHash,
		arg.GasPrice,
		arg.GasLimit,
		arg.Nonce,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.ChainID,
		&i.ToAddress,
		&i.Amount,
		&i.TokenID,
		&i.GasPrice,
		&i.GasLimit,
		&i.Nonce,
		&i.Status,
		&i.TxHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FromAddress,
	)
	return i, err
}
