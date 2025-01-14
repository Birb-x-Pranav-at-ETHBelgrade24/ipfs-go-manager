package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/CreedsCode/ipfs-go-manager/internal/api"
	"github.com/CreedsCode/ipfs-go-manager/internal/auth"
	"github.com/CreedsCode/ipfs-go-manager/internal/config"
	"github.com/CreedsCode/ipfs-go-manager/internal/ipfs"
	"github.com/CreedsCode/ipfs-go-manager/internal/web3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			c.Set("Access-Control-Allow-Headers", "*")
			c.Set("Access-Control-Allow-Origin", "*")
			c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	storage := auth.NewInMemoryStorage()
	storage.CreateUser(config.Env.PrivateKey, 1000, true)

	chainId := new(big.Int)
	chainId.SetString(config.Env.ChainId, 10)

	web3Client, err := web3.NewWeb3Client(config.Env.NodeUrl, config.Env.PrivateKey, config.Env.StorageContractAddress, chainId)
	if err != nil {
		log.Fatalf("failed to create web3 client: %v", err)
	}

	authMiddleWare := auth.NewAuthMiddleware(storage, web3Client)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, my beloved ones!")
	})

	node, err := ipfs.ConnectLocalNode(config.Env.IPFSHostAddress)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	api.SetupRoutes(app, node, authMiddleWare)
	api.SetupAdminRoutes(app, authMiddleWare)

	log.Fatal(app.Listen(config.Env.HostAddress))
}
