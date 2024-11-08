package main

import (
	"context"
	"log"
	_ "mpc/docs"
	"mpc/internal/delivery/http"
	_ "mpc/internal/domain"
	"mpc/internal/infrastructure/auth"
	"mpc/internal/infrastructure/config"
	"mpc/internal/infrastructure/db"
	"mpc/internal/infrastructure/ethereum"
	"mpc/internal/infrastructure/kafka"
	"mpc/internal/infrastructure/logger"
	"mpc/internal/infrastructure/redis"
	"mpc/internal/repository/postgres"
	"mpc/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
)

// @title MPC API
// @version 1.0
// @description This is the API documentation for the MPC project.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// db
	dbPool, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()
	useCvTaiSampleData(dbPool, cfg)

	// redis
	redisClient, err := redis.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to initialize Redis client: %v", err)
	}
	defer redisClient.Close()

	// kafka
	kafkaProducer, err := kafka.NewKafkaProducer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()
	walletCreatedPublisher, err := kafka.NewKafkaProducer(cfg, kafka.WithTopic(cfg.Kafka.WalletCreatedTopic))
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}
	defer walletCreatedPublisher.Close()

	// logger
	log := logger.NewLogger()

	// jwt
	jwtConfig := auth.NewJWTConfig(cfg.JWT.SecretKey, cfg.JWT.TokenDuration, cfg.JWT.TokenDuration*30)
	jwtService := auth.NewJWTService(jwtConfig, *redisClient)

	// ethereum
	ethClient, err := ethereum.NewEthereumClient(cfg.Ethereum.URL, cfg.Ethereum.SecretKey)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum client: %v", err)
	}

	// repository
	userRepo := postgres.NewUserRepo(dbPool)
	walletRepo := postgres.NewWalletRepo(dbPool)
	transactionRepo := postgres.NewTransactionRepo(dbPool)
	balanceRepo := postgres.NewBalanceRepo(dbPool)
	chainRepo := postgres.NewChainRepo(dbPool)

	// usecase
	walletUC := usecase.NewWalletUC(walletRepo, ethClient, walletCreatedPublisher)
	authUC := usecase.NewAuthUC(userRepo, walletUC, *jwtService)
	userUC := usecase.NewUserUC(userRepo)
	txnUC := usecase.NewTxnUC(transactionRepo, ethClient, walletUC, *redisClient, kafkaProducer)
	balanceUC := usecase.NewBalanceUC(balanceRepo, walletRepo, chainRepo, ethClient)

	// router
	router := http.NewRouter(&userUC, &walletUC, &txnUC, &authUC, &balanceUC, jwtService, log)

	log.Fatal(router.Run(":8080"))
}


func useCvTaiSampleData (dbPool *pgxpool.Pool, cfg *config.Config) {
	if(cfg.DB.UseCvTaiSample){
		log.Printf("Use CV Tai sample data")
		_, err := dbPool.Exec(context.Background(), "DELETE FROM chains")
		if err != nil {
			log.Printf("Unable to delete chains table: %v\n", err)
		}
		_, err = dbPool.Exec(context.Background(), "DELETE FROM tokens")
		if err != nil {
			log.Printf("Unable to delete tokens table: %v\n", err)
		}
		// the previous code will also delete transactoins data :))

	
		// Insert sample data
		//chain 1
		_, err = dbPool.Exec(context.Background(), "INSERT INTO chains (id, name, chain_id, rpc_url, ws_url, last_scan_block_number) VALUES ('2773fa12-645a-45d0-80a2-79cf5a2ecf96', 'Sepolia', '11155111', 'https://sepolia.infura.io/v3/6c89fb7fa351451f939eea9da6bee755', 'wss://sepolia.infura.io/ws/v3/6d3cfcab0e3a442eb3c890ae4422f76d', -1)")
		if err != nil {
			log.Printf("Unable to insert sample data into chains table: %v\n", err)
		}
		_, err = dbPool.Exec(context.Background(), "INSERT INTO tokens (id, chain_id, name, symbol, decimals, contract_address) VALUES ('2773fa12-645a-45d0-80a2-79cf5a2ecf98', '2773fa12-645a-45d0-80a2-79cf5a2ecf96', 'SepoliaETH', 'ETH', 18, '0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43' )")
		if err != nil {
			log.Printf("Unable to insert sample data into tokens table: %v\n", err)
		}
		_, err = dbPool.Exec(context.Background(), "UPDATE chains SET native_token_id = '2773fa12-645a-45d0-80a2-79cf5a2ecf98' WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf96'")
		if err != nil {
			log.Printf("Unable to update native_token_id in chains table: %v\n", err)
		}


		//chain 2
		// _, err = dbPool.Exec(context.Background(), "INSERT INTO chains (id, name, chain_id, rpc_url, ws_url, last_scan_block_number) VALUES ('2773fa12-645a-45d0-80a2-79cf5a2ecf97', 'Linea Sepolia', '59141', 'https://linea-sepolia.infura.io/v3/6c89fb7fa351451f939eea9da6bee755', 'wss://linea-sepolia.infura.io/ws/v3/6d3cfcab0e3a442eb3c890ae4422f76d', -1)")
		// if err != nil {
		// 	log.Printf("Unable to insert sample data into chains table: %v\n", err)
		// }
		// _, err = dbPool.Exec(context.Background(), "INSERT INTO tokens (id, chain_id, name, symbol, decimals,contract_address) VALUES ('2773fa12-645a-45d0-80a2-79cf5a2ecf99', '2773fa12-645a-45d0-80a2-79cf5a2ecf97', 'LineaETH', 'ETH', 18, '0xe1a12515F9AB2764b887bF60B923Ca494EBbB2d6')")
		// if err != nil {
		// 	log.Printf("Unable to insert sample data into tokens table: %v\n", err)
		// }
		// _, err = dbPool.Exec(context.Background(), "UPDATE chains SET native_token_id = '2773fa12-645a-45d0-80a2-79cf5a2ecf99' WHERE id = '2773fa12-645a-45d0-80a2-79cf5a2ecf97'")
		// if err != nil {
		// 	log.Printf("Unable to update native_token_id in chains table: %v\n", err)
		// }
	}
}