package resto

import (
	"context"
	"errors"
	"fmt"
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/repository"
	"food-ordering/internal/domain/usecase"
	"food-ordering/internal/infrastructure/helper"
	"time"
)

type cartUseCaseImpl struct {
	cartRepo  repository.CartRepository
	orderRepo repository.OrderRepository
}

func (c *cartUseCaseImpl) AddToCart(ctx context.Context, userID string, item domain.CartItem) error {
	return c.cartRepo.AddToCart(ctx, userID, item)
}

func (c *cartUseCaseImpl) CheckoutCart(ctx context.Context, userID string) (*domain.Order, error) {
	cart, err := c.cartRepo.GetCartByUserID(ctx, userID)
	if err != nil || len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	fmt.Println("cart", cart)

	order := domain.Order{
		ID:         helper.GenerateOrderID(),
		CertID:     cart.ID,
		TotalPrice: cart.Total,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	orderID, err := c.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// Clear the cart after checkout
	_ = c.cartRepo.ClearCart(ctx, userID)

	order.ID = orderID
	return &order, nil
}

func NewCartUseCase(cartRepository repository.CartRepository, orderRepository repository.OrderRepository) usecase.CartUseCase {
	return &cartUseCaseImpl{
		cartRepo:  cartRepository,
		orderRepo: orderRepository,
	}
}
