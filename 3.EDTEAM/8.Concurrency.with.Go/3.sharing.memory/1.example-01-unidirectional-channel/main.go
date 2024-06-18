package main

import (
	"fmt"
	"time"
)

func main() {
	// number := make(chan int, 2)
	// signal := make(chan struct{})
	// go receive(signal, number)
	// send(number)

	// <-signal

	number3 := make(chan int)
	signal3 := make(chan struct{})
	go receive3(signal3, number3)
	send3(number3)
	signal3 <- struct{}{}
}

func send(number chan<- int) {
	number <- 1
	number <- 2
	number <- 3
	number <- 4
	number <- 5
	number <- 6
	close(number)
}

func send3(number3 chan<- int) {
	number3 <- 1
	number3 <- 2
	number3 <- 3
	time.Sleep(1)
	number3 <- 4
}

func receive3(signal3 <-chan struct{}, number3 <-chan int) {
	for {
		select {
		case v := <-number3:
			fmt.Println(v)
		case <-signal3:
			return
		default:
			fmt.Println("... waiting")
		}
	}
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

	signal <- struct{}{}
}
