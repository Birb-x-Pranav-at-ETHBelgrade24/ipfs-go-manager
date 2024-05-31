package auth

import "errors"

type User struct {
	APIKey string
	Quota  int
	Admin  bool
}

type Storage interface {
	GetUser(apiKey string) (*User, error)
	CreateUser(apiKey string, quota int, admin bool) error
	UpdateQuota(apiKey string, quota int) error
}

var ErrUserNotFound = errors.New("user not found")
var ErrQuotaExceeded = errors.New("quota exceeded")
