package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/CreedsCode/ipfs-go-manager/internal/web3"
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

func (m *AuthMiddleware) AdminMiddleware(c *fiber.Ctx) error {
	apiKey := getAPIKey(c)
	if apiKey == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": ErrInvalidAPIKey.Error()})
	}

	user, err := m.storage.GetUser(apiKey)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if !user.Admin {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "admin privileges required"})
	}

	c.Locals("user", user)
	return c.Next()
}

func (m *AuthMiddleware) Storage() Storage {
	return m.storage
}

func getAPIKey(c *fiber.Ctx) string {
	apiKey := c.Get("Authorization")
	fmt.Printf("api:", apiKey)
	if len(apiKey) > 7 && apiKey[:7] == "Bearer " {
		return apiKey[7:]
	}
	return ""
}
