package user

import (
	"context"
	"food-ordering/internal/domain"
)

type User interface {
	FetchAllMenu(ctx context.Context) ([]domain.Menu, error)
	OrderMenu(ctx context.Context, order domain.Order) error
}
