package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:8080"

func main() {
	lc := loginClient(url+"/v1/login", "batman@gmail.com", "123456")
	// fmt.Println(lc)

	person := Person{
		Name:        "Bruce Wayne",
		Age:         35,
		Communities: []Community{Community{Name: "Bats"}, Community{Name: "More bats"}},
	}

	gr := createPerson(url+"/v1/person", lc.Data.Token, &person)
	fmt.Println(gr)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Fatalf("Request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("Request: %v", err)
	}

	return response
}
