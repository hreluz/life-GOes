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

func add(num1 int, num2 int) int {
	return num1 + num2
}

func generateRandomNumbers() (int, int) {
	return rand.Intn(100), rand.Intn(40)
}

func printNumber(number int) {
	fmt.Printf("The number is %v ", number)
}
