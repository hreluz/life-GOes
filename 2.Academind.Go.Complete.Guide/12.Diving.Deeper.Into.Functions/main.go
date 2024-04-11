package main

import (
	"fmt"
)

func main() {

	// numbers := []int{1, 10, 25}

	sum := sumup(1, 10, 25)

	fmt.Println(sum)

}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return startingValue + sum
}
