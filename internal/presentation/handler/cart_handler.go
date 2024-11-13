package handler

import (
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CartHandler struct {
	uc usecase.CartUseCase
}

func NewCartHandler(uc usecase.CartUseCase) *CartHandler {
	return &CartHandler{uc: uc}
}

func (h *CartHandler) HandleCrtItem(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("user_id").(string)

	var cartItem domain.CartItem
	if err := c.Bind(&cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	c.Logger().Infof("Bound cart item: %+v", cartItem)

	if err := h.uc.AddToCart(ctx, userID, cartItem); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "item added to cart"})

}

func (h *CartHandler) HandleCheckout(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("user_id").(string)

	var cartItem domain.Order
	if err := c.Bind(&cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	c.Logger().Infof("Bound cart item: %+v", cartItem)

	order, err := h.uc.CheckoutCart(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := map[string]interface{}{
		"status_code": "200:OK",
		"message":     "success",
		"order": map[string]interface{}{
			"data": order,
		},
	}

	return c.JSON(http.StatusOK, response)

}
