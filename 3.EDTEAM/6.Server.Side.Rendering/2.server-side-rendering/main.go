package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	myFunctions := template.FuncMap{
		"usd": func(a float64) string {
			return fmt.Sprintf("$%.fUSD", a)
		},
	}
	_ = myFunctions

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("public/imgs"))))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error when uploading to server: %v", err)
	}
}
