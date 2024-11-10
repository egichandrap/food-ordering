package handler

import (
	"food-ordering/internal/application/usecase/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MenuHandler struct {
	uc user.UC
}

func NewHandler(uc user.UC) *MenuHandler {
	return &MenuHandler{uc: uc}
}

func (h *MenuHandler) HandleMenu(c echo.Context) error {
	ctx := c.Request().Context()
	menus, err := h.uc.GetListMenu(ctx)
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
