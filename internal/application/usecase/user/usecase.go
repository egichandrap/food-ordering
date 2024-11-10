package user

import (
	"context"
	"food-ordering/internal/domain"
)

type UC interface {
	GetListMenu(ctx context.Context) ([]domain.Menu, error)
}
