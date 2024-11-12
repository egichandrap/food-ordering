package usecase

import (
	"context"
	"food-ordering/internal/domain"
)

type UserUseCase interface {
	Register(ctx context.Context, user domain.User) (string, error)
}
