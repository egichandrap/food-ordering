package repository

import (
	"context"
	"food-ordering/internal/domain"
)

type CatalogRepository interface {
	FetchAllMenu(ctx context.Context) ([]domain.Menu, error)
}

type CartRepository interface {
	AddToCart(ctx context.Context, userID string, item domain.CartItem) error
	GetCartByUserID(ctx context.Context, userID string) (*domain.Cart, error)
	ClearCart(ctx context.Context, userID string) error
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (string, error)
	GetOrderById(ctx context.Context, orderID string) (*domain.Order, error)
}
