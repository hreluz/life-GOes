package main

import (
	"log"
	"net/http"
	"sync"
)

// run server in this directory

var urls = []string{
	"http://localhost:1234?duration=3",
	"http://localhost:1234?duration=1",
	"http://localhost:1234?duration=5",
}

func main() {
	fetchConcurrent(urls)
	fetchSequential(urls)
}

func fetchSequential(urls []string) {
	for _, url := range urls {
		fetch(url, "sequential")
	}
}

func fetchConcurrent(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(u string) {
			fetch(u, "concurrent")
			wg.Done()
		}(url)
	}

	wg.Wait()
}

func fetch(url, typeFetch string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("failed url of type %s: %s, err: %v", typeFetch, url, err)
	}
	log.Printf("%s of type %s: %d ", url, typeFetch, resp.StatusCode)
}
