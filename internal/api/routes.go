package api

import (
	"fmt"

	"github.com/CreedsCode/ipfs-go-manager/internal/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/ipfs/kubo/client/rpc"
)

func SetupRoutes(app *fiber.App, node *rpc.HttpApi, authMiddleWare *auth.AuthMiddleware) {
	app.Get("/pincid/:cid", authMiddleWare.Middleware, func(c *fiber.Ctx) error {
		user := c.Locals("user").(string)
		fmt.Println("AAAAAAAAA===========", user)
		return pinCIDHandler(c, node)

	})

	app.Post("/upload", authMiddleWare.Middleware, func(c *fiber.Ctx) error {
		return uploadFileHandler(c, node)
	})
}
