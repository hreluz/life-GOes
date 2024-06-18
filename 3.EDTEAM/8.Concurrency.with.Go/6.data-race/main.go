package main

import (
	"fmt"
	"sync"
)

func main() {
	courses := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)

	go func() {
		mu.Lock()
		courses["go from zero"] = "Intermediate"
		mu.Unlock()
		wg.Done()
	}()

	go func() {
		mu.Lock()
		courses["go concurrency"] = "Advanced"
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(courses)
}
