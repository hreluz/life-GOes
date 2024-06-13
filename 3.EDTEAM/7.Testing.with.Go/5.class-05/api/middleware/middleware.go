package middleware

import (
	"net/http"

	"github.com/hreluz/go-testing-class-5/authorization"
	"github.com/labstack/echo"
)

// Authentication .
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "no permitido"})
		}

		return f(c)
	}
}
