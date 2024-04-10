package main

import "fmt"

type transfromFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	dNumbers := transformNumbers(&numbers, double)
	fmt.Println(dNumbers)

	tNumbers := transformNumbers(&numbers, triple)
	fmt.Println(tNumbers)
}

func transformNumbers(numbers *[]int, transform transfromFunc) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
