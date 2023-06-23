package main

import "fmt"

func main() {

	firstName := "Hector"

	var lastName string
	lastName = "Lavoe"

	fmt.Println(firstName)
	fmt.Println(lastName)

	var currentYear int = 2023
	birthYear := 1946

	age := currentYear - birthYear
	fmt.Println(age)

	currentYear += 1
	age = currentYear - birthYear
	fmt.Println(age)

}

// := -> declare and define the variable in one step
// only var -> declares the variable

// For int values, the default value is 0, for float it's 0.0 and for string values (text), it's an empty string("")
