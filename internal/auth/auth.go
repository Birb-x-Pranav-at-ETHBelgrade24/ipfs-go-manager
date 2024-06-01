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
	ErrNoSignature    = errors.New("missing signature")
	ErrSignature      = errors.New("invalid signature")
	ErrDeductingQuota = errors.New("issues deducting quota")
)

type AuthMiddleware struct {
	Web3Client *web3.Web3Client
	NonceStore *NonceStore
}

func NewAuthMiddleware(web3client *web3.Web3Client) *AuthMiddleware {
	return &AuthMiddleware{Web3Client: web3client, NonceStore: NewNonceStore()}
}

func (am *AuthMiddleware) Middleware(c *fiber.Ctx) error {
	var request struct {
		Message   string `json:"message"`
		Signature string `json:"signature"`
		Nonce     string `json:"nonce"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	userAddr, err := am.verifySignature(request.Message+request.Nonce, request.Signature)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	storedNounce := am.NonceStore.GetNonce(userAddr.Hex())
	if storedNounce == "" || storedNounce != request.Nonce {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid nonce"})
	}

	am.NonceStore.GenerateNonce(userAddr.Hex())

	c.Locals("user", userAddr.Hex())
	// tx, err := m.deducQuota(userAddr)
	// if err != nil {
	// 	return c.Status(http.StatusPaymentRequired).JSON(fiber.Map{"status": "error", "message": ErrDeductingQuota.Error()})
	// }

	// c.Locals("quota_tx", tx)
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
	if len(apiKey) > 7 && apiKey[:7] == "Bearer " {
		return apiKey[7:]
	}
	return ""
}
