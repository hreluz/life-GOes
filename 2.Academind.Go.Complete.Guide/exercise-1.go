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

// 5) Calculate the difference ("age") between the two years and
//   store it in a new variable. Output that result in the command line.

// 6) Overwrite the value stored in the "current year" variable with
//   the previous value + 1 (i.e. next year). Calculate the next year,
//   don't just change the number.
//   Repeat step 5) with that new value (without changing any of the previous code).

// Try all those steps on your own first, then have a look at my solution
// lecture to compare your solution to mine!
