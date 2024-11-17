package utils

import (
	"context"
	"log"
	"mpc/internal/infrastructure/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UseCvTaiSampleData(dbPool *pgxpool.Pool, cfg *config.Config) {
	if cfg.DB.UseCvTaiSample {
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
