package repository

import (
	"context"
	"errors"
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/repository"
	"sync"
	"time"
)

type CartRepo struct {
	data map[string]*domain.Cart
	mu   sync.Mutex
}

func (c *CartRepo) AddToCart(ctx context.Context, userID string, item domain.CartItem) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	cart, exists := c.data[userID]
	if !exists {
		cart = &domain.Cart{ID: "121312", UserID: userID, Items: []domain.CartItem{}}
		c.data[userID] = cart
	}
	cart.Items = append(cart.Items, item)
	cart.Total += item.Price * float64(item.Quantity)
	cart.CreatedAt = time.Now()
	return nil
}

func (c *CartRepo) GetCartByUserID(ctx context.Context, userID string) (*domain.Cart, error) {
	if cart, exists := c.data[userID]; exists {
		return cart, nil
	}

	return nil, errors.New("cart not found")
}

func (c *CartRepo) ClearCart(ctx context.Context, userID string) error {
	delete(c.data, userID)
	return nil
}

func NewCartRepository() repository.CartRepository {
	return &CartRepo{data: make(map[string]*domain.Cart)}
}
