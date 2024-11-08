package domain

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        uuid.UUID
	WalletID  uuid.UUID
	ChainID   uuid.UUID
	TokenID   uuid.UUID
	Balance   float64
	UpdatedAt time.Time
}

type GetBalanceResponse struct {
	TokenID uuid.UUID   `json:"token_id"`
	TokenName string    `json:"token_name"`
	ContractAddress string `json:"contract_address"`
	TokenSymbol string  `json:"token_symbol"`
	Decimals int64      `json:"decimals"`
	Balance float64     `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateBalanceParams struct {
	Address string
	TokenID uuid.UUID
	Balance float64
}
