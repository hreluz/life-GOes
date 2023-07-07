package main

import (
	"fmt"
	"math/rand"
)

func main() {
	printNumber(add(generateRandomNumbers()))
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}

func generateRandomNumbers() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}
