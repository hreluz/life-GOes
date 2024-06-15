package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "Hello !"
		// fmt.Println(<-message)
	}()

	fmt.Println(<-message)
}
