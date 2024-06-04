package main

import (
	"log"
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hreluz/echo-framework/authorization"
	"github.com/hreluz/echo-framework/handler"
	"github.com/hreluz/echo-framework/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")

	if err != nil {
		log.Fatalf("certificates could not be loaded: %v", err)
	}

	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	log.Println("Server started in port 8080")

	err = e.Start(":8080")

	if err != nil {
		log.Printf("Error in the server: %v\n", err)
	}
}
