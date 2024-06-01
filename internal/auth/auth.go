package auth

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/CreedsCode/ipfs-go-manager/internal/web3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidAPIKey  = errors.New("invalid API key")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrDeductingQuota = errors.New("could not deduct quota")
)

type AuthMiddleware struct {
	storage    Storage
	Web3Client *web3.Web3Client
}

func NewAuthMiddleware(storage Storage, web3Client *web3.Web3Client) *AuthMiddleware {
	return &AuthMiddleware{storage: storage, Web3Client: web3Client}
}

func (m *AuthMiddleware) Middleware(c *fiber.Ctx) error {
	apiKey := getAPIKey(c)
	if apiKey == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": ErrInvalidAPIKey.Error()})
	}

	user, err := m.storage.GetUser(apiKey)
	if err != nil {
		m.storage.CreateUser(apiKey, 100, false)
		user, err = m.storage.GetUser(apiKey)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
	}

	if user.Quota <= 0 {
		return c.Status(http.StatusPaymentRequired).JSON(fiber.Map{"status": "error", "message": ErrQuotaExceeded.Error()})
	}

	user.Quota--
	m.storage.UpdateQuota(apiKey, user.Quota)

	tx, err := m.Web3Client.DeductQuota("0x538522b81a333340a5E1605b7298E7d765781412", 1)
	if err != nil {
		return c.Status(http.StatusPaymentRequired).JSON(fiber.Map{"status": "error", "message": ErrDeductingQuota.Error()})
	}
	c.Locals("user", user)
	c.Locals("tx", tx)
	return c.Next()

}

func (am *AuthMiddleware) verifySignature(message, signature string) (common.Address, error) {
	if strings.HasPrefix(signature, "0x") {
		signature = signature[2:]
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to decode signature: %v", err)
	}

	msgHash := crypto.Keccak256Hash([]byte(message))
	pubKey, err := crypto.SigToPub(msgHash.Bytes(), sig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to recover public key from signature: %v", err)
	}

	addr := crypto.PubkeyToAddress(*pubKey)

	return addr, nil
}

// func (m *AuthMiddleware) deducQuota(userAddr common.Address) (string, error) {
// 	tx, err := m.Web3Client.Contract.UseQuota(m.Web3Client.Auth, userAddr, big.NewInt(1))
// 	if err != nil {
// 		fmt.Println("failed to deduct quota:", userAddr.Hex(), err.Error())
// 		return "", err
// 	}

// 	x := tx.Hash().Hex()
// 	fmt.Printf("Quota deducted for user %s. Transaction hash: %s\n", userAddr.Hex(), x)
// 	return x, nil
// }

// func (m *AuthMiddleware) AdminMiddleware(c *fiber.Ctx) error {
// 	apiKey := getAPIKey(c)
// 	if apiKey == "" {
// 		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": ErrInvalidAPIKey.Error()})
// 	}

// 	user, err := m.storage.GetUser(apiKey)
// 	if err != nil {
// 		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
// 	}

// 	if !user.Admin {
// 		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "admin privileges required"})
// 	}

// 	c.Locals("user", user)
// 	return c.Next()
// }

// func (m *AuthMiddleware) Storage() Storage {
// 	return m.storage
// }

func getAPIKey(c *fiber.Ctx) string {
	apiKey := c.Get("Authorization")
	fmt.Printf("api:", apiKey)
	if len(apiKey) > 7 && apiKey[:7] == "Bearer " {
		return apiKey[7:]
	}
	return ""
}
