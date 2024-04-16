package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func generateValue(c chan int, limit chan int) {
	limit <- 1
	fmt.Println("Generating value...")
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c <- randN.Intn(10)
	<-limit
}

func main() {

	c := make(chan int)
	limiter := make(chan int, 3)

	// running both generated numbers at the same time, rather than wait until the next one
	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)

	sum := 0
	i := 0

	for num := range c {
		sum += num
		i++

		if i == 6 {
			close(c)
		}
	}

	fmt.Println(sum)

	x := make(chan int)
	y := make(chan int)
	limiter2 := make(chan int, 3)

	go generateValue(x, limiter2)
	go generateValue(y, limiter2)

	var a int
	var b int

	select {
	case a = <-x:
		fmt.Printf("X finished first, value is %v ", a)
	case b = <-y:
		fmt.Printf("Y finished first, value is %v ", b)
	}
}
