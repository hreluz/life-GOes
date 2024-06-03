package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/division", division)
	e.Start(":8080")
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"greeting": "Hello world"})
}

func division(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)

	if f == 0 {
		return c.String(http.StatusBadRequest, "Value cannot be zero")
	}

	r := 3000 / f
	return c.String(http.StatusOK, strconv.Itoa(r))
}
