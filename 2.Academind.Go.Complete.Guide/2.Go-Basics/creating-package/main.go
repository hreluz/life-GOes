package main

import (
	"fmt"

	"hlavoe.com/myorganization/firstapp/greeting"
)

func main() {
	luckyNUmber := 17
	evenMoreLuckyNumber := luckyNUmber + 5
	evenMoreLuckyNumber +=1

	fmt.Println(greeting.GreetingText)
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

	firstName := "James"
	lastName := "Bond"
	longText := `
		this 
		is 
		a long
		text
	`
	fmt.Println(firstName + " " + lastName)
	fmt.Println("9" + "1")
	fmt.Println(longText)

	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 32
	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old ", fullName, age, age)
}
