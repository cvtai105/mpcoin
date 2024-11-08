package domain

import (
	"time"

	"github.com/google/uuid"
)

type Chain struct {
	ID             uuid.UUID
	Name           string
	ChainID        string
	RPCURL         string
	WSURL          string
	NativeCurrency string
	NativeTokenID  uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ExplorerURL    string
	LastScanBlock  int64
}
