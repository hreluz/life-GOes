//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello", name)
}

func printer(message string) {
	fmt.Println(message)
}

func add(n1, n2, n3 int) int {
	return n1 + n2 + n3
}

func r1() int {
	return 10
}

func r2() (int, int) {
	return 20, 30
}

func main() {
	greeting("Batman")

	printer("I am the night")

	n2, n3 := r2()

	number := add(r1(), n2, n3)

	fmt.Println(number)
}
