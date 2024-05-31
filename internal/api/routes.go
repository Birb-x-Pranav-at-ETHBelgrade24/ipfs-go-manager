package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ipfs/kubo/client/rpc"
)

func SetupRoutes(app *fiber.App, node *rpc.HttpApi) {
	app.Get("/pincid/:cid", func(c *fiber.Ctx) error {
		return pinCIDHandler(c, node)
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		return uploadFileHandler(c, node)
	})
}
