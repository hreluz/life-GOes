package main

import "fmt"

func main() {
	number := make(chan int)
	go receive(number)
	send(number)
}

func send(number chan<- int) {
	number <- 10
}

func receive(number <-chan int) {
	fmt.Println(<-number)
}
