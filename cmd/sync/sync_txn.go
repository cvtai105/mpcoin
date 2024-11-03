package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chain struct {
	ID             int64
	Name           string
	ChainID        string
	RPCURL         string
	NativeCurrency string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ExplorerURL    string
	BlockIndex     int64
}

// StartSyncJob initializes sync jobs for each chain
func StartSyncJob(chains []Chain) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, chain := range chains {
		wg.Add(1)
		go func(chain Chain) {
			defer wg.Done()
			fmt.Printf("Starting sync job for chain: %s\n", chain.Name)
			// Call your setup and sync functions here with ctx for cancellation support
			err := setupAndSync(ctx, chain)
			if err != nil {
				fmt.Printf("Error syncing chain %s: %v\n", chain.Name, err)
			}
		}(chain)
	}
	wg.Wait()
}
// FetchBlockData simulates fetching and processing block data
func FetchBlockData(chain Chain, blockIndex int64) {
	fmt.Printf("Fetching block data for chain %s from block %d\n", chain.Name, blockIndex)
	// Implement block and transaction scanning here
}

// setupAndSync simulates scanning from a start block to end block once, then listening for new blocks
func setupAndSync(ctx context.Context, chain Chain) error {

	// Initial scan from start block to the latest block
	fmt.Printf("Scanning blocks from block %d to end on chain: %s\n", 1000, chain.Name)
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond) // Simulate initial block scan

	
	fmt.Printf("Setting up WebSocket connection for chain: %s\n", chain.Name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate connection setup delay

	// Channel to signal end of listening
	done := make(chan struct{})

	// Goroutine for listening to new blockcreated events
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Block listener for chain %s was cancelled\n", chain.Name)
				close(done)
				return
			case <-done:
				fmt.Printf("Stopping block listener for chain %s\n", chain.Name)
				return
			default:
				// Simulate receiving a blockcreated event
				fmt.Printf("Listen on chain %s\n", chain.Name)
				time.Sleep(time.Duration(5000) * time.Millisecond) // Simulate time until next block is created
				fmt.Printf("New block created on chain %s\n", chain.Name)

				// Trigger async transaction scan for the new block
				go asyncScanNewBlock(ctx, chain)
			}
		}
	}()

	// Wait until we are done listening or context is cancelled
	<-done
	fmt.Printf("Completed sync for chain: %s\n", chain.Name)
	return nil
}

// asyncScanNewBlock simulates scanning transactions for a new block asynchronously
func asyncScanNewBlock(ctx context.Context, chain Chain) {
	select {
	case <-ctx.Done():
		fmt.Printf("Async scan for new block on chain %s was cancelled\n", chain.Name)
		return
	default:
		fmt.Printf("Asynchronously scanning transactions for new block on chain: %s\n", chain.Name)
		time.Sleep(time.Duration(2000) * time.Millisecond) // Simulate async transaction scanning
		fmt.Printf("Completed async transaction scan for new block on chain: %s\n", chain.Name)
	}
}

func main() {
	chains := loadChainsFromDatabase()
	StartSyncJob(chains)
}

func loadChainsFromDatabase() []Chain {
	return []Chain{
		{
			ID:             1,
			Name:           "Ethereum",
			ChainID:        "1",
			RPCURL:         "https://mainnet.infura.io/v3/your_project_id",
			NativeCurrency: "ETH",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			ExplorerURL:    "https://etherscan.io",
			BlockIndex:     0,
		},
		{
			ID:             2,
			Name:           "Sepolia Testnet",
			ChainID:        "11155111",
			RPCURL:         "https://sepolia-testnet.infura.io/v3/your_project_id",
			NativeCurrency: "SEP",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			ExplorerURL:    "https://sepolia-explorer.io",
			BlockIndex:     0,
		},
	}

}
