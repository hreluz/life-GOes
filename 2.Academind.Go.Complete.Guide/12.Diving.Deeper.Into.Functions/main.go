package main

import (
	"fmt"
)

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	moreNumbers := []int{5, 1, 2}

	dNumbers := transformNumbers(&numbers, double)
	fmt.Println(dNumbers)

	tNumbers := transformNumbers(&numbers, triple)
	fmt.Println(tNumbers)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

	// anonymous function
	transformedAnonymousNumbers := tranformWithAnonymousValues(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformedAnonymousNumbers)
}

func tranformWithAnonymousValues(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// func(int) int
func getTransformerFunction(numbers *[]int) transformFunc {

	if (*numbers)[0] == 1 {
		return double
	}

	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
