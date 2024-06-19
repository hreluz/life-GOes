package main

import (
	"fmt"
)

func sendString(s string, ch chan string) {
	ch <- s
	// close(ch)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendString("hello from channel 1", ch1)
	go sendString("hello from channel 2", ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}
