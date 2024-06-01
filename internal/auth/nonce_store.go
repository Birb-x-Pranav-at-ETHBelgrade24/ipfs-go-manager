package auth

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

type NonceStore struct {
	nonces map[string]string
	mu     sync.Mutex
}

func NewNonceStore() *NonceStore {
	return &NonceStore{nonces: make(map[string]string)}
}

func (ns *NonceStore) GenerateNonce(address string) string {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	nonceBytes := make([]byte, 16)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		panic(err)
	}
	nonce := hex.EncodeToString(nonceBytes)

	ns.nonces[address] = nonce
	go ns.expireNonce(address, nonce, 5*time.Minute)

	return nonce
}

func (ns *NonceStore) GetNonce(address string) string {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	return ns.nonces[address]
}

func (ns *NonceStore) expireNonce(address, nonce string, duration time.Duration) {
	time.Sleep(duration)
	ns.mu.Lock()
	defer ns.mu.Unlock()
	if ns.nonces[address] == nonce {
		delete(ns.nonces, address)
	}
}
