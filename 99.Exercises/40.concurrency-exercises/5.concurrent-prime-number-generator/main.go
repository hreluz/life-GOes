package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func generatePrimes(primes chan<- int) {
	for i := 2; ; i++ {
		if isPrime(i) {
			primes <- i
		}
	}
}

func main() {

	primes := make(chan int)

	go generatePrimes(primes)

	for i := 0; i < 10; i++ {
		fmt.Println("Prime: ", <-primes)
	}
}
