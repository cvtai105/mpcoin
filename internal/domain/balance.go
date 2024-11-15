package domain

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Balance struct {
	ID        uuid.UUID
	WalletID  uuid.UUID
	ChainID   uuid.UUID
	TokenID   uuid.UUID
	Balance   big.Int
	UpdatedAt time.Time
}

type GetBalanceResponse struct {
	TokenID uuid.UUID   `json:"token_id"`
	TokenName string    `json:"token_name"`
	ContractAddress string `json:"contract_address"`
	TokenSymbol string  `json:"token_symbol"`
	ChainID uuid.UUID   `json:"chain_id"`
	Decimals int64      `json:"decimals"`
	Balance big.Int     `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateBalanceParams struct {
	Balance *big.Int
	Address common.Address
	TokenID uuid.UUID
}
