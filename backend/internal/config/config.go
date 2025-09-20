package config

import (
	"os"
	"strconv"
)

type Config struct {
	Environment string         `json:"environment"`
	Server      ServerConfig   `json:"server"`
	Database    DatabaseConfig `json:"database"`
	Redis       RedisConfig    `json:"redis"`
	Blockchain  BlockchainConfig `json:"blockchain"`
}

type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type BlockchainConfig struct {
	RPCUrl          string `json:"rpc_url"`
	PrivateKey      string `json:"private_key"`
	ContractAddress string `json:"contract_address"`
	ChainID         int64  `json:"chain_id"`
}

func Load() (*Config, error) {
	config := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "nft_capture_game"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Blockchain: BlockchainConfig{
			RPCUrl:          getEnv("BLOCKCHAIN_RPC_URL", "http://localhost:8545"),
			PrivateKey:      getEnv("BLOCKCHAIN_PRIVATE_KEY", ""),
			ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
			ChainID:         getEnvAsInt64("CHAIN_ID", 1337),
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}