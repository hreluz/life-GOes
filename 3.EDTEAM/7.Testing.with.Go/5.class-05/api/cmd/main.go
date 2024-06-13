package main

import (
	"log"

	"github.com/hreluz/go-testing-class-5/authorization"
	"github.com/hreluz/go-testing-class-5/handler"
	"github.com/hreluz/go-testing-class-5/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)

	log.Println("Servidor iniciado en el puerto 8080")

	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
