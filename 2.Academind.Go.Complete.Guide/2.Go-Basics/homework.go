package main

import "fmt"

func main() {
	var name string
	lastName := "Bond"

	name = "James"

	fmt.Println(name, lastName)

	var currentYear int
	birthYear := 2000

	currentYear = 2024

	age := currentYear - birthYear
	fmt.Println(age)
	
	currentYear = currentYear + 1

	age = currentYear - birthYear

	fmt.Println(age)

}
