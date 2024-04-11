package closures

import (
	"fmt"
)

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
