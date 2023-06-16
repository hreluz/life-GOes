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

	var defaultFloat float64 = 1.23456789123456789
	var smallFloat float32 = 1.23456789123456789

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = 'a' //it is an integer, it converts as ASCII
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a' // same as rune, but only for smaller
	fmt.Println(firstByte)

	firstName := "Hector"
	lastName := "Lavoe"
	fullName := fmt.Sprintln(firstName + " " + lastName)
	newFullName := fmt.Sprintf("%v %v", firstName, lastName)

	fmt.Println("9" + "1")

	multipleLines := `This is 
	a text with
	many lines
	
	and more lines.`

	fmt.Println(multipleLines)

	age := 30

	fmt.Printf("Hi, my full name is %v and I am %v (Type: %T) years old.\n", fullName, age, age)

	fmt.Printf("%v\n", newFullName)

	fmt.Println("end")
}
