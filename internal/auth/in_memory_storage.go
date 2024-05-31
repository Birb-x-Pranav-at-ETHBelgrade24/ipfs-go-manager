package auth

import "sync"

type InMemoryStorage struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		users: make(map[string]*User),
	}
}

func (s *InMemoryStorage) GetUser(apiKey string) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[apiKey]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *InMemoryStorage) CreateUser(apiKey string, quota int, admin bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[apiKey] = &User{
		APIKey: apiKey,
		Quota:  quota,
		Admin:  admin,
	}

	return nil
}

func (s *InMemoryStorage) UpdateQuota(apiKey string, quota int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[apiKey]
	if !exists {
		return ErrUserNotFound
	}

	user.Quota = quota
	return nil
}
