package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 5
	b := 10

	sum := add(a, b)
	printNumber(sum)

	n1, n2 := generateRandomNumbers()

	fmt.Printf("\nThe random numbers are %v and %v.\n", n1, n2)

	fmt.Println("I'm executing inside a function")
}

func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func generateRandomNumbers() (r1 int, r2 int) {
	r1 = rand.Intn(100)
	r2 = rand.Intn(40)
	return
}

func printNumber(number int) {
	fmt.Printf("The number is %v ", number)
}
