package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"

	mockdb "mpc/cmd/sync/db"
	"mpc/cmd/sync/models"
	"mpc/internal/domain"
	"mpc/internal/infrastructure/config"
	"mpc/internal/infrastructure/ethereum"
	"mpc/internal/repository"
	"mpc/internal/repository/postgres"

	db "mpc/internal/infrastructure/db"
)

var (
	// wsApiKey = os.Getenv("WEBSOCKET_API_KEY")
	walletAddressMapId = make(map[string]uuid.UUID)
    // mu      sync.Mutex
)

func StartKafka(cfg *config.Config) {
	kafConf := kafka.ReaderConfig{
		Brokers:   cfg.Kafka.Brokers,
		GroupID:   cfg.Kafka.SyncGroupId,
		Topic:     cfg.Kafka.WalletCreatedTopic,
		// MaxBytes:  10, 
	}

	reader := kafka.NewReader(kafConf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
			fmt.Printf("Error reading kafka message: %v\n", err)
			continue
		}

		//Handle message
		walletCreatedEventHandle(string(m.Value))
	}
}

func walletCreatedEventHandle(message string){
	var data domain.Wallet
    err := json.Unmarshal([]byte(message), &data)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }
	//update wallet address hash table
	// mu.Lock()
	walletAddressMapId[data.Address] = data.ID
	// mu.Unlock()
}

func persistUsersTransactions(ctx context.Context, transactions []domain.Transaction, ethRepo repository.EthereumRepository, tnxRepo repository.TransactionRepository, chain models.Chain) error {
	usersTransactions := []domain.Transaction{}
	for _, tnx := range transactions {

		// chỉ lấy các tnx từ address lạ chuyển tới address của user
		if(walletAddressMapId[tnx.ToAddress] == uuid.Nil){	
			continue
		}
		if(walletAddressMapId[tnx.FromAddress] != uuid.Nil){
			continue
		}


		fmt.Printf("Transaction: %s\n", tnx.TxHash)
		fmt.Printf("From: %s\n", tnx.FromAddress)
		fmt.Printf("To: %s\n", tnx.ToAddress)

		tnxReceipt, err := ethRepo.GetTransactionReceipt(ctx, common.HexToHash(tnx.TxHash))
		if err != nil {
			fmt.Printf("Error getting transaction receipt: %v\n", err)
			return err
		}
		if(tnxReceipt.Status == 0){
			tnx.Status = domain.StatusFailed
		}else{
			tnx.Status = domain.StatusSuccess
		}
		tnx.ChainID = chain.ID
		tnx.TokenID = chain.NativeTokenID
		usersTransactions = append(usersTransactions, tnx)
	}

	err := tnxRepo.InsertSettledTransactions(ctx, usersTransactions)
	if err != nil {
		fmt.Printf("Error inserting transactions: %v\n", err)
		return err
	}

	return nil
}

func SyncChainData(
	ctx context.Context, 
	chain models.Chain, 
	ethRepo repository.EthereumRepository,
	wsEthRepo repository.EthereumRepository,
	tnxRepo repository.TransactionRepository) error {
	// Initial scan from start block to the latest block
	if chain.LastScanBlock == -1 {
		fmt.Printf("Chain %s added.\n", chain.Name)
	} else {
		go func(blockNum uint64, ethRepo repository.EthereumRepository, tnxRepo repository.TransactionRepository, chain models.Chain) {
			fmt.Printf("Chain %s is scanning from block %d to end\n", chain.Name, chain.LastScanBlock)
			transactions, err := ethRepo.GetTransactionsStartFrom(blockNum)
			if err != nil {
				fmt.Printf("Error getting transactions: %v\n", err)
				return
			}

			err = persistUsersTransactions(ctx, transactions, ethRepo, tnxRepo, chain)
			if err != nil {
				fmt.Printf("Error persisting transactions: %v\n", err)
				return
			}
		}(uint64(chain.LastScanBlock), ethRepo, tnxRepo, chain)
	}

	fmt.Printf("Setting up WebSocket connection for chain: %s\n", chain.Name)
	headers := make(chan *types.Header)
	sub, err := wsEthRepo.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Failed to subscribe to new head: %v", err)
	}

	// Lắng nghe sự kiện khi có block mới
	fmt.Println("Listening for new blocks...")
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Context canceled, stopping block listener for chain %s\n", chain.Name)
			return nil
		case err := <-sub.Err():
			fmt.Printf("Subscription error on chain %s: %v. Retrying...\n", chain.Name, err)
			time.Sleep(5 * time.Second) // Retry delay before reconnecting
			sub, err = wsEthRepo.SubscribeNewHead(ctx, headers)
			if err != nil {
				log.Printf("Failed to resubscribe to new head for chain %s: %v", chain.Name, err)
				continue
			}
		case header := <-headers:
			go func(header *types.Header, ethRepo repository.EthereumRepository, tnxRepo repository.TransactionRepository, chain models.Chain) {
				fmt.Printf("Chain %s: New block %d\n", chain.Name, header.Number.Uint64())
				transactions, err := ethRepo.GetTransactionsInBlock(header.Number.Uint64())
				if err != nil {
					fmt.Printf("Error getting transactions for block %d on chain %s: %v\n", header.Number.Uint64(), chain.Name, err)
					return
				}

				if len(transactions) > 0 {
					fmt.Printf("Chain %s: Found %d navtive transactions in block %d\n", chain.Name, len(transactions), header.Number.Uint64())
				}

				err = persistUsersTransactions(ctx, transactions,ethRepo, tnxRepo, chain)
				if err != nil {
					fmt.Printf("Error persisting transactions: %v\n", err)
					return
				}
			}(header, wsEthRepo, tnxRepo, chain)
		}
	}

}

//to-do: add arguments hanlder to run specific chain sync job
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return
	}

	dbPool, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	transactionRepo := postgres.NewTransactionRepo(dbPool)
	go StartKafka(cfg)

	chains := mockdb.LoadChainsFromDatabase()
	wallets := mockdb.LoadWalletsFromDatabase()

	for _, wallet := range wallets {
		walletAddressMapId[wallet.Address] = wallet.ID
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, chain := range chains {
		wg.Add(1)
		ethRepo, err := ethereum.NewEthereumClient(chain.RPCURL, cfg.Ethereum.SecretKey)
		if err != nil {
			log.Fatalf("Failed to initialize Ethereum client: %v", err)
			wg.Done()
			continue
		}
		wsEthRepo, err := ethereum.NewEthereumClient(chain.WSURL, cfg.Ethereum.SecretKey)
		if err != nil {
			log.Fatalf("Failed to initialize Ethereum client: %v", err)
			wg.Done()
			continue
		}
		
		go func(chain models.Chain, ethRepo repository.EthereumRepository, tnxRepo repository.TransactionRepository) {
			defer wg.Done()
			fmt.Printf("Sync: Starting sync job for chain: %s\n", chain.Name)
			err := SyncChainData(ctx, chain, ethRepo,wsEthRepo, tnxRepo)
			if err != nil {
				fmt.Printf("Sync: Error syncing chain %s: %v\n", chain.Name, err)
			}
		}(chain, ethRepo, transactionRepo)
	}
	wg.Wait()
}

