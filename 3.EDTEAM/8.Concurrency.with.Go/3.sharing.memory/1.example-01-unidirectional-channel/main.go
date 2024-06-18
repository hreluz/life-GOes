package main

import "fmt"

func main() {
	number := make(chan int, 2)
	signal := make(chan struct{})
	go receive(signal, number)
	send(number)

	<-signal
}

func send(number chan<- int) {
	number <- 1
	number <- 2
	number <- 3
}

func receive(signal chan<- struct{}, number <-chan int) {
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)

	signal <- struct{}{}
}
