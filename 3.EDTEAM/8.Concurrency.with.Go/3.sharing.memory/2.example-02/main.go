package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// run server in this directory
// go test -bench=.

var urls = []string{
	"http://localhost:1234?duration=3",
	"http://localhost:1234?duration=1",
	"http://localhost:1234?duration=10",
	"http://localhost:1234?duration=5",
	"http://localhost:1234?duration=2",
	"http://localhost:1234?duration=2",
}

func main() {
	// fetchConcurrent(urls)
	// fetchConcurrentSCP(urls)
	fetchConcurrentCancelation(urls)
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

func fetchConcurrentSCP(urls []string) {
	signal := make(chan struct{})

	for _, url := range urls {
		go func(u string) {
			fetch(u, "csp")
			signal <- struct{}{}
		}(url)
	}

	<-signal
	<-signal
	<-signal
}

func fetchConcurrentCancelation(urls []string) {
	done := make(chan struct{})

	for _, url := range urls {
		go func(u string) {
			fetch(u, "cancellation")

			select {
			case <-done:
				return
			}
		}(url)
	}

	select {
	case <-time.After(time.Second * 4):
		close(done)
	}
}

func fetch(url, typeFetch string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("failed url of type %s: %s, err: %v", typeFetch, url, err)
	}
	log.Printf("%s of type %s: %d ", url, typeFetch, resp.StatusCode)
}
