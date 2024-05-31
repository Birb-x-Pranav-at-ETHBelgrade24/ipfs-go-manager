package auth

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidAPIKey = errors.New("invalid API key")
	ErrUnauthorized  = errors.New("unauthorized")
)

type AuthMiddleware struct {
	storage Storage
}

func NewAuthMiddleware(storage Storage) *AuthMiddleware {
	return &AuthMiddleware{storage: storage}
}

func (m *AuthMiddleware) Middleware(c *fiber.Ctx) error {
	apiKey := getAPIKey(c)
	if apiKey == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": ErrInvalidAPIKey.Error()})
	}

	user, err := m.storage.GetUser(apiKey)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if user.Quota <= 0 {
		return c.Status(http.StatusPaymentRequired).JSON(fiber.Map{"status": "error", "message": ErrQuotaExceeded.Error()})
	}

	user.Quota--
	m.storage.UpdateQuota(apiKey, user.Quota)

	c.Locals("user", user)
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
	if len(apiKey) > 7 && apiKey[:7] == "Bearer " {
		return apiKey[7:]
	}
	return ""
}
