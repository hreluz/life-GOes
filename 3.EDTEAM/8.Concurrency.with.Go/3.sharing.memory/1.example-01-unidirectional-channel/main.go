package main

import (
	"fmt"
)

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
	number <- 4

	close(number)
}

func receive(signal chan<- struct{}, number <-chan int) {
	// for i := 0; i < 10; i++ {
	// 	v, ok := <-number
	// 	if ok {
	// 		fmt.Printf("%d %t\n", v, ok)
	// 	}
	// }

	for v := range number {
		fmt.Println(v)
	}
	// fmt.Println(<-number)
	// fmt.Println(<-number)
	// fmt.Println(<-number)

	signal <- struct{}{}
}
