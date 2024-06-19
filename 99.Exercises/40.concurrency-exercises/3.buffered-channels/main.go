package main

import "fmt"

func printNumbers(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch)
}

func main() {

	ch := make(chan int, 3)
	go printNumbers(ch)

	for n := range ch {
		fmt.Printf("Received: %d\n", n)
	}
}
