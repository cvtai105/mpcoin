package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App        AppConfig
	DB         DBConfig
	JWT        JWTConfig
	Redis      RedisConfig
	Ethereum   EthereumConfig
	Kafka      KafkaConfig
	MailConfig MailConfig
}

type DBConfig struct {
	ConnStr string `mapstructure:"CONN_STR"`
	UseCvTaiSample bool `mapstructure:"DB_CVTAI_SAMPLE"`
}

type AppConfig struct {
	Port int `mapstructure:"PORT"`
}

type JWTConfig struct {
	SecretKey     string        `mapstructure:"SECRET_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
}

type EthereumConfig struct {
	URL       string `mapstructure:"ETHEREUM_URL"`
	SecretKey string `mapstructure:"ETHEREUM_SECRET_KEY"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"REDIS_HOST"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type KafkaConfig struct {
	Brokers 			[]string 	`mapstructure:"BROKERS"`
	Topic   			string   	`mapstructure:"TOPIC"`
	SyncGroupId 		string 		`mapstructure:"KAFKA_SYNCHRONIZE_GROUP_ID"`
	WalletCreatedTopic 	string 		`mapstructure:"KAFKA_WALLET_CREATED_TOPIC"`
	TransactionFoundTopic string 		`mapstructure:"KAFKA_TRANSACTION_FOUND_TOPIC"`
}

type MailConfig struct {
	SMTPHost      string `mapstructure:"SMTP_HOST"`
	SMTPPort      int    `mapstructure:"SMTP_PORT"`
	SMTPUsername  string `mapstructure:"SMTP_USERNAME"`
	SMTPPassword  string `mapstructure:"SMTP_PASSWORD"`
	FromEmail     string `mapstructure:"FROM_EMAIL"`
	OTPExpiration int    `mapstructure:"OTP_EXPIRATION"`
}

// Define default values
// Define default values
var defaults = map[string]string{
	"DB.CONN_STR":        "postgres://viet:123@localhost:5432/mpcoin?sslmode=disable",
	"DB.MAX_CONNECTIONS": "10",
	"APP.PORT":           "8080",
	"APP.ENV":            "development",
	"JWT.SECRET_KEY":     "chirp-chirp",
	"JWT.TOKEN_DURATION": "1h",
	"REDIS.ADDR":         "localhost:6379",
	"REDIS.PASSWORD":     "",
	"REDIS.DB":           "0",
	"KAFKA.BROKERS":      "localhost:29092",
	"KAFKA.TOPIC":        "mpc",
	// "ETHEREUM.URL":        "https://sepolia.infura.io/v3/<INFURA_PROJECT_ID>",
	// "ETHEREUM.SECRET_KEY": "<INFURA_SECRET_KEY>",
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	
	// Set environment variable names to match .env file
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	
	if err := viper.ReadInConfig(); err != nil {
		// app could not run without a .env file, so I commented this block
		// if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// 	return nil, fmt.Errorf("error reading config file: %w", err)
		// }
		fmt.Println("No .env file found. Using environment variables.")
		viper.AutomaticEnv()
	}

	// Set default values if not provided in .env or environment
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	// Manually set values for fields that don't match the default structure
	config.DB.ConnStr = viper.GetString("CONN_STR")
	config.DB.UseCvTaiSample = viper.GetBool("DB_USE_CVTAI_SAMPLE")
	
	config.Ethereum.URL = viper.GetString("ETHEREUM_URL")
	config.Ethereum.SecretKey = viper.GetString("ETHEREUM_SECRET_KEY")

	config.Kafka.Brokers = viper.GetStringSlice("BROKERS")
	config.Kafka.WalletCreatedTopic = viper.GetString("KAFKA_WALLET_CREATED_TOPIC")
	config.Kafka.TransactionFoundTopic = viper.GetString("KAFKA_TRANSACTION_FOUND_TOPIC")
	config.Kafka.SyncGroupId = viper.GetString("KAFKA_SYNCHRONIZE_GROUP_ID")

	config.Redis.Addr = viper.GetString("REDIS_HOST")
	config.Redis.Password = viper.GetString("REDIS_PASSWORD")
	config.Redis.DB = viper.GetInt("REDIS_DB")

	// Set default values if not provided
	if config.JWT.TokenDuration == 0 {
		config.JWT.TokenDuration = 24 * time.Hour // Default to 24 hours
	}

	log.Printf("Config loaded")
	return &config, nil
}
