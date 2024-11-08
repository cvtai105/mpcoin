package db

import (
	"mpc/cmd/sync/models"
	"time"

	"github.com/google/uuid"
)

func LoadWalletsFromDatabase() []models.Wallet {
	return []models.Wallet{
		{
			ID:      uuid.New(),
			Address: "0x2F2c845811F14d0F70FaB20F30031816D20081C6",
		},
	}
}

func LoadChainsFromDatabase() []models.Chain {
	return []models.Chain{
		{
			ID:             uuid.MustParse("2773fa12-645a-45d0-80a2-79cf5a2ecf96"),
			Name:           "Ethereum",
			ChainID:        "1",
			RPCURL:         "https://mainnet.infura.io/v3/6c89fb7fa351451f939eea9da6bee755",
			WSURL:          "wss://mainnet.infura.io/ws/v3/6c89fb7fa351451f939eea9da6bee755",
			NativeCurrency: "ETH",
			NativeTokenID: 	uuid.MustParse("2773fa12-645a-45d0-80a2-79cf5a2ecf96"),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			ExplorerURL:    "https://etherscan.io",
			LastScanBlock:  -1,
		},
		{
			ID:             uuid.MustParse("2773fa12-645a-45d0-80a2-79cf5a2ecf96"),
			Name:           "Sepolia",
			ChainID:        "11155111",
			RPCURL:         "https://sepolia.infura.io/v3/6c89fb7fa351451f939eea9da6bee755",
			WSURL:          "wss://sepolia.infura.io/ws/v3/6c89fb7fa351451f939eea9da6bee755",
			NativeCurrency: "SEP",
			NativeTokenID: 	uuid.MustParse("2773fa12-645a-45d0-80a2-79cf5a2ecf96"),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			ExplorerURL:    "https://sepolia.etherscan.io",
			LastScanBlock:  -1,
		},
	}
}
