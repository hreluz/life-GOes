package main

import (
	"fmt"
)

func printNumbers(n int, done chan bool) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
	done <- true
}

func main() {

	done1 := make(chan bool)
	done2 := make(chan bool)
	done3 := make(chan bool)

	go printNumbers(5, done1)
	go printNumbers(10, done2)
	go printNumbers(15, done3)

	<-done1
	<-done2
	<-done3
}
