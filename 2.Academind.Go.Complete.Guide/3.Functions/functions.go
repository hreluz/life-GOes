package main

import (
	"fmt"
	"math/rand"
)

func main() {
	printNumber(add(generateRandomNumbers()))
}

func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}

func generateRandomNumbers() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)
	return
}
