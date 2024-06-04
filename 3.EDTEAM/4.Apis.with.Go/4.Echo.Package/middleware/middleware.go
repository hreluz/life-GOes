package middleware

import (
	"net/http"

	"github.com/hreluz/echo-framework/authorization"
	"github.com/labstack/echo/v4"
)

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)

		if err != nil {
			return forbidden(c)
		}

		return f(c)
	}
}

func forbidden(c echo.Context) error {
	return c.JSON(http.StatusForbidden, map[string]string{"error": "You are not allowed"})
}
