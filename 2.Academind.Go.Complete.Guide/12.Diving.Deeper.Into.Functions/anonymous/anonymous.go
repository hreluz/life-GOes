package anonymous

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	// anonymous function
	transformedAnonymousNumbers := tranformWithAnonymousValues(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformedAnonymousNumbers)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func tranformWithAnonymousValues(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
