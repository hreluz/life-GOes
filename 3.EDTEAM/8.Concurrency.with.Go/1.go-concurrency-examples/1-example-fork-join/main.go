package main

import (
	"fmt"
	"time"
)

func hello() int {
	fmt.Println("This is a hello world")
	return 1
}

func main() {
	go hello()

	go func() {
		fmt.Println("An anonymous function")
	}()

	time.Sleep(time.Millisecond)

	fmt.Println("This is the end")
}
