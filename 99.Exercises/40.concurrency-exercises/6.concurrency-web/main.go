package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	url        string
	bodyLength int
}

func fetch(url string, results chan Result, done chan bool) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	results <- Result{url: url, bodyLength: len(body)}
	done <- true
}

func main() {

	urls := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.x.com/",
	}

	results := make(chan Result)
	done := make(chan bool)

	for _, url := range urls {
		go fetch(url, results, done)
	}

	go func() {
		for i := 0; i < len(urls); i++ {
			<-done
		}

		close(results)
	}()

	for r := range results {
		fmt.Printf("Web: %s, Content length: %d\n\n", r.url, r.bodyLength)
	}
}
