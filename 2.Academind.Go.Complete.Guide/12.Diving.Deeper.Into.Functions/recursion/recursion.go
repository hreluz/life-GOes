package main

import (
	"fmt"
)

func main() {

	fact := factorial(5)
	fmt.Println(fact)

	factWithRecursion := factorialWithRecursion(10)
	fmt.Println(factWithRecursion)

}

func factorialWithRecursion(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}
