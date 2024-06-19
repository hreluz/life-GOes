package main

import (
	"fmt"
	"time"
)

func printNumber(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumber(10)
	go printNumber(5)
	go printNumber(15)

	time.Sleep(2 * time.Second)
}
