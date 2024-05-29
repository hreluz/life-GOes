package main

import (
	"log"
	"net/http"

	"github.com/hreluz/go-api-crud/authorization"
	"github.com/hreluz/go-api-crud/handler"
	"github.com/hreluz/go-api-crud/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub ")

	if err != nil {
		log.Fatalf("certificates could not be loaded: %v", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Server started in port 8080")

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Printf("Error in the server: %v\n", err)
	}
}
