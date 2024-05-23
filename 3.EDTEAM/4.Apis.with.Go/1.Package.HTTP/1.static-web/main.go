package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Server started on: http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
