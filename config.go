package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HostAddress     string
	IPFSHOSTAddress string
	DBUser          string
	DBPassword      string
	DBAddress       string
	DBName          string
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		HostAddress:     fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST", "http://localhost"), getEnv("PORT", "8081")),
		IPFSHOSTAddress: fmt.Sprintf("%s:%s", getEnv("IPFS_HOST", "http://localhost"), getEnv("IPFS_HOST_PORT", "8081")),
		DBUser:          getEnv("DB_USER", "root"),
		DBPassword:      getEnv("DB_PASSWORD", "mypassword"),
		DBName:          getEnv("DB_NAME", "backend"),
		DBAddress:       fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
