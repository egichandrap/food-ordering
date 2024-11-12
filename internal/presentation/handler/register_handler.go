package handler

import (
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	uc usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) HandleRegister(c echo.Context) error {
	ctx := c.Request().Context()
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	token, err := h.uc.Register(ctx, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
