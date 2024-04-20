package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("\nFunction with params: ")
	hello("Bruce", "Wayne")
	emoji := "🐶"
	change(&emoji)
	fmt.Println(emoji)

	fmt.Println("\nFunction with return param: ")
	total := sum(10, 20)
	fmt.Println(total)

	fmt.Println("\nFunction with many return params: ")
	text := "Batman"
	min, may := convert(text)
	fmt.Println(min, may)

	fmt.Println("\nFunction with errors: ")
	readFile()
	res, _ := division(10, 2)
	fmt.Println("\n", res)

	_, err := division(10, 0)
	fmt.Println("error: ", err)

	fmt.Println("\nFunction that accept a function as param: ")
	nums := []int{1, 4, 70, 5, 67, 90, 2}
	filterNums := filter(nums, func(num int) bool {
		return num <= 10
	})
	fmt.Println(filterNums)

	fmt.Println("\nFunction that accept a function as param and return a function: ")
	x := helloThere("Obi Wan Kenobi")
	fmt.Println(x(", said Greivous"))

	fmt.Println("\nVariadic function: ")
	fmt.Println(sumVariadic(1, 2, 3, 4, 5))

	fmt.Println("\nAnonymous function: ")
	anon := func(text string) {
		fmt.Println("Hello there ", text)
	}
	anon("Grievious")
}

func change(value *string) {
	*value = "😆"
}

func hello(firstName string, lastName string) {
	fmt.Printf("Hello %s %s", firstName, lastName)
	fmt.Println("\n")
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}

func readFile() {
	content, err := ioutil.ReadFile("./goo.mod")

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(string(content))
}

func division(dividend, divisor int) (result int, err error) {
	if divisor == 0 {
		result = 0
		err = errors.New("You cannot divide by zero")
		return
	}

	result = dividend / divisor
	return
}

func filter(nums []int, callback func(int) bool) []int {

	result := []int{}

	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	return result
}

func helloThere(name string) func(string) string {
	return func(text string) string {
		return "Hello " + name + text
	}
}

func sumVariadic(nums ...int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	return totalSum
}
