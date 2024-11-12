package repository

import (
	"context"
	"food-ordering/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (string, error)
	GetUserByUsername(username string) (*domain.User, error)
}
