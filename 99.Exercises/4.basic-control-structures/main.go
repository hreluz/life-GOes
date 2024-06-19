package main

import (
	"fmt"
	"math/rand"
)

func isOdd(n int) {
	if n%2 == 0 {
		fmt.Printf("%d is odd", n)
	} else {
		fmt.Printf("%d is not odd", n)
	}
}

func main() {
	isOdd(rand.Intn(100))
}
