package handler

import (
	// "net/http"

	// "github.com/hreluz/echo-framework/middleware"
	"github.com/hreluz/echo-framework/middleware"
	"github.com/labstack/echo/v4"
)

func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("/v1/person")
	person.Use(middleware.Authentication)

	person.POST("", h.create)

	// mux.HandleFunc("/v1/persons/get-all", middleware.Log(h.getAll))
	// mux.HandleFunc("/v1/persons/update", h.update)
	// mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	// mux.HandleFunc("/v1/persons/get-by-id", h.getById)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
