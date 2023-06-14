package main

import "fmt"

func main() {
	// var greetingText string
	// greetingText = "Hello World!"

	// var greetingText string = "Hello World"

	greetingText := "Hello World"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greetingText)
	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3
	fmt.Println(newNumber)
}
