package main

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, my beloved ones!")
	})

	// "Connect" to local node
	node, shouldReturn := connectLocalNode()
	if shouldReturn {
		return
	}

	app.Get("/pincid/:cid", func(c *fiber.Ctx) error {
		cid := c.Params("cid")
		err := pinCID(cid, node)
		if err != nil {
			fmt.Printf("error pinning cid: %s", cid)
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		return c.Status(200).JSON(fiber.Map{"status": "success"})
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		// Step 2: Open the file
		src, err := file.Open()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
		defer src.Close()

		// Step 3: Convert to a format suitable for files.NewReaderFile
		fileReader := multipartFileToFileReader(src, file.Size)

		// Step 4: Create a context
		ctx := context.Background()

		// Step 5: Add the file to the node
		cidFile, err := node.Unixfs().Add(ctx, fileReader)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		fmt.Printf("Added file to peer with CID %s\n", cidFile.String())

		return c.Status(200).JSON(fiber.Map{"status": "success", "cid": cidFile.String()})
	})

	log.Fatal(app.Listen(":3000"))
}

func pinCID(toPinCid string, node *rpc.HttpApi) error {
	// Pin a given file by its CID handle long running maybe like this?
	// ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	// defer cancel()

	ctx := context.Background()

	path, err := cidStringToPath(toPinCid)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = node.Pin().Add(ctx, path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("CID pinned successfully")
	return nil
}

func cidStringToPath(toPinCid string) (path.ImmutablePath, error) {
	cid, err := cid.Parse(toPinCid)
	if err != nil {
		fmt.Printf(err.Error())
		return path.ImmutablePath{}, err
	}

	path := path.FromCid(cid)

	return path, nil
}

func connectLocalNode() (*rpc.HttpApi, bool) {
	addr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/5001")

	if err != nil {
		fmt.Println(err.Error())
		return nil, true
	}

	node, err := rpc.NewApi(addr)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, true
	}
	return node, false
}

// multipartFileToFileReader converts a multipart.File to a files.File for use with the IPFS node
func multipartFileToFileReader(file multipart.File, size int64) files.File {
	return files.NewReaderFile(&fileReader{file, size})
}

// fileReader is a wrapper to implement the files.File interface
type fileReader struct {
	file multipart.File
	size int64
}

func (fr *fileReader) Read(p []byte) (n int, err error) {
	return fr.file.Read(p)
}

func (fr *fileReader) Close() error {
	return fr.file.Close()
}

func (fr *fileReader) Size() (int64, error) {
	return fr.size, nil
}
