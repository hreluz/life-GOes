package main

import "fmt"

func main() {
	printNumber(add(5, 6))
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}
