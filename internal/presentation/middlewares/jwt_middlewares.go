package middlewares

import (
	"food-ordering/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		userID, err := utils.ValidateJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		c.Set("user_id", userID)
		return next(c)
	}
}
