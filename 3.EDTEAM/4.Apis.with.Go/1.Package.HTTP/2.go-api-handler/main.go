package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register the route and the handler
	http.HandleFunc("/hello", hello)
	// Uploads a server in the port 8080
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world")
}
