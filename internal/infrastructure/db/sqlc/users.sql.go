// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING id, email, password_hash, created_at, updated_at, name, avatar_url
`

type CreateUserParams struct {
	Email        string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AvatarUrl,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, password_hash, created_at, updated_at, name, avatar_url FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AvatarUrl,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password_hash, created_at, updated_at, name, avatar_url FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AvatarUrl,
	)
	return i, err
}

const getUserWithWallet = `-- name: GetUserWithWallet :one
SELECT u.id, u.email, w.id, w.address
FROM users u
LEFT JOIN wallets w ON u.id = w.user_id
WHERE u.id = $1 LIMIT 1
`

type GetUserWithWalletRow struct {
	ID      pgtype.UUID
	Email   string
	ID_2    pgtype.UUID
	Address pgtype.Text
}

func (q *Queries) GetUserWithWallet(ctx context.Context, id pgtype.UUID) (GetUserWithWalletRow, error) {
	row := q.db.QueryRow(ctx, getUserWithWallet, id)
	var i GetUserWithWalletRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.ID_2,
		&i.Address,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET email = $2, password_hash = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, email, password_hash, created_at, updated_at, name, avatar_url
`

type UpdateUserParams struct {
	ID           pgtype.UUID
	Email        string
	PasswordHash string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.ID, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.AvatarUrl,
	)
	return i, err
}
