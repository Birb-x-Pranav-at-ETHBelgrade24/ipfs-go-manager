package api

import (
	"context"
	"fmt"

	"github.com/CreedsCode/ipfs-go-manager/internal/ipfs"
	"github.com/CreedsCode/ipfs-go-manager/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/ipfs/kubo/client/rpc"
)

func pinCIDHandler(c *fiber.Ctx, node *rpc.HttpApi) error {
	cid := c.Params("cid")
	err := ipfs.PinCID(cid, node)
	if err != nil {
		fmt.Printf("error pinning cid: %s", cid)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "tx": c.Locals("tx").(string)})
}

func uploadFileHandler(c *fiber.Ctx, node *rpc.HttpApi) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	defer src.Close()

	fileReader := models.MultipartFileToFileReader(src, file.Size)

	ctx := context.Background()
	cidFile, err := node.Unixfs().Add(ctx, fileReader)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	fmt.Printf("Added file to peer with CID %s\n", cidFile.String())

	return c.Status(200).JSON(fiber.Map{"status": "success", "cid": cidFile.String(), "tx": c.Locals("tx").(string)})
}
