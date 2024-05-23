package main

import (
	"log"
	"net/http"

	"github.com/hreluz/go-api-crud/handler"
	"github.com/hreluz/go-api-crud/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	log.Println("Server started in port 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Printf("Error in the server: %v\n", err)
	}
}
