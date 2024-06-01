package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NodeUrl                string
	StorageContractAddress string
	HostAddress            string
	IPFSHostAddress        string
	PrivateKey             string
	ChainId                string
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	private_key := getEnv("ADMIN_API_KEY", "")
	if private_key == "" {
		fmt.Println("NO ADMIN_API_KEY SET")
		os.Exit(1)
	}
	return Config{
		HostAddress:            fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST", "127.0.0.1"), getEnv("PORT", "3000")),
		IPFSHostAddress:        getEnv("IPFS_API_ADDRESS", "/ip4/127.0.0.1/tcp"),
		PrivateKey:             private_key,
		NodeUrl:                getEnv("NODE_URL", "http://127.0.0.1:8545"),
		StorageContractAddress: getEnv("STORAGE_CONTRACT_ADDRESS", ""),
		ChainId:                getEnv("CHAIN_ID", "31337"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
