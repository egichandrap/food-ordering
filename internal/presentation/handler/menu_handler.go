package handler

import (
	"food-ordering/internal/domain/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CatalogHandler struct {
	uc usecase.CatalogUseCase
}

func NewCatalogHandler(uc usecase.CatalogUseCase) *CatalogHandler {
	return &CatalogHandler{uc: uc}
}

func (h *CatalogHandler) HandleCatalogMenu(c echo.Context) error {
	ctx := c.Request().Context()
	menus, err := h.uc.GetCatalogMenu(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status_code": "500:Internal Server Error",
			"message":     "Failed to fetch menus",
		})
	}

	response := map[string]interface{}{
		"status_code": "200:OK",
		"message":     "success",
		"menu": map[string]interface{}{
			"data": menus,
		},
	}

	return c.JSON(http.StatusOK, response)

}
