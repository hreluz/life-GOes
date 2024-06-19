package main

import "fmt"

func main() {

	var age int = 25
	var number64 float64 = 24.3
	var message string = "This is a message"
	var isOpen bool = false

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("A float 64 number: %.2f\n", number64)
	fmt.Printf("This is the message: '%s'\n", message)
	fmt.Printf("Boolean value is %v\n", isOpen)
}
