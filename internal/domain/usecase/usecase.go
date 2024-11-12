package usecase

import (
	"context"
	"food-ordering/internal/domain"
)

type CatalogUseCase interface {
	GetCatalogMenu(ctx context.Context) ([]domain.Menu, error)
}

type CartUseCase interface {
	AddToCart(ctx context.Context, userID string, item domain.CartItem) error
	CheckoutCart(ctx context.Context, userID string) (*domain.Order, error)
}

type OrderUseCase interface {
	GetOrder(ctx context.Context, orderID string) (*domain.Order, error)
}
