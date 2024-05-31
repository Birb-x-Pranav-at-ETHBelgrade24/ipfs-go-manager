package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HostAddress     string
	IPFSHostAddress string
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		HostAddress:     fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST", "127.0.0.1"), getEnv("PORT", "3000")),
		IPFSHostAddress: getEnv("IPFS_API_ADDRESS", "/ip4/127.0.0.1/tcp"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
