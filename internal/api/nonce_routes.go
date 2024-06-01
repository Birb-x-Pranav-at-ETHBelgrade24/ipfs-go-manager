package api

import (
	"github.com/CreedsCode/ipfs-go-manager/internal/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/ipfs/kubo/client/rpc"
)

func NonceRoutes(app *fiber.App, node *rpc.HttpApi, authMiddleWare *auth.AuthMiddleware) {
	app.Get("/nonce/:address", func(c *fiber.Ctx) error {
		address := c.Params("address")
		nonce := authMiddleWare.NonceStore.GenerateNonce(address)
		return c.JSON(fiber.Map{"nonce": nonce})
	})
}
