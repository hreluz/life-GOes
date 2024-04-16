package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func generateValue(c chan int) {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c <- randN.Intn(10)
}

func main() {

	c := make(chan int)

	// running both generated numbers at the same time, rather than wait until the next one
	go generateValue(c)
	go generateValue(c)
	go generateValue(c)
	go generateValue(c)
	go generateValue(c)

	sum := 0
	i := 0

	for num := range c {
		sum += num
		i++

		if i == 4 {
			close(c)
		}
	}

	fmt.Println(sum)
}
