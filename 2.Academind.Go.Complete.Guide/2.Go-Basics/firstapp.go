package main

import "fmt"

func main() {

	// var greetingText string
	// greetingText = "Hello World"

	// var greetingText string = "Heello world !"

	greetingText := "Hello there"
	luckyNUmber := 17
	evenMoreLuckyNumber := luckyNUmber + 5

	fmt.Println(greetingText)
	fmt.Println(luckyNUmber)

	evenMoreLuckyNumber = 99
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNUmber) / 3
	fmt.Println(newNumber)

	// it has to be single quotes
	var firstRune rune = '€'
	fmt.Println(firstRune)         //8364
	fmt.Println(string(firstRune)) // EURO SIGN

	var firstByte byte = 'a' //only numbers and letters
	fmt.Println(firstByte)
}
