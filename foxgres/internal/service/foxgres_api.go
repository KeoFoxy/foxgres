package service

import (
	"context"

	"foxgres/internal/storage"
)

type Service interface {
	FindUser(ctx context.Context, login int, pass string) (int, error)
}

type service struct {
	userStorage storage.UserStorage
}

func NewService(us storage.UserStorage) Service {
	return &service{
		userStorage: us,
	}
}

func (s *service) FindUser(ctx context.Context, login int, pass string) (int, error) {
	return s.userStorage.FindUser(ctx, login, pass)
}
