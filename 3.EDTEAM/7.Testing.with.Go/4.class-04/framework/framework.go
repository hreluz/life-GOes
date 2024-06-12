package framework

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func Get(c echo.Context) error {
	p := Person{
		"Batman",
		35,
	}

	return c.JSON(http.StatusOK, p)
}
