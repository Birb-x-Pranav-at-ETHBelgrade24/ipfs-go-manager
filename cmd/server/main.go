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
)

func main() {
	app := fiber.New()

	chainId := new(big.Int)
	_, success := chainId.SetString(config.Env.ChainId, 10)
	if !success {
		fmt.Println("Error converting chainId")
		return
	}

	web3Client, err := web3.NewWeb3Client(config.Env.NodeUrl, config.Env.PrivateKey, config.Env.StorageContractAddress, chainId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	authMiddleWare := auth.NewAuthMiddleware(web3Client)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, my beloved ones!")
	})

	node, err := ipfs.ConnectLocalNode(config.Env.IPFSHostAddress)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	api.SetupRoutes(app, node, authMiddleWare)
	api.NonceRoutes(app, node, authMiddleWare)
	// api.SetupAdminRoutes(app, authMiddleWare)

	log.Fatal(app.Listen(config.Env.HostAddress))
}
