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
	person := e.Group("/person")
	person.POST("", create)
	person.GET("/:id", get)
	person.PUT("/:id", update)
	person.DELETE("/:id", delete)

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

func create(c echo.Context) error {
	return c.String(http.StatusOK, "created")
}

func update(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "updated "+id)
}

func delete(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "deleted "+id)
}

func get(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "query "+id)
}
