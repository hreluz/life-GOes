package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)

	// myAge := &age

	var myAge *int
	myAge = &age

	fmt.Println(myAge)
	fmt.Println(*myAge)

	// Save value and get memory address
	*myAge = 33

	// Print memory address
	fmt.Println(myAge)

	// Print value
	fmt.Println(*myAge)

	// Since address was copied to myAge and then myAge assigned 33, age is now 33
	fmt.Println(age)

	// sending a pointer value
	doubledAge := double(myAge)
	fmt.Println(doubledAge)
	fmt.Println(age)
}

// it expects a pointer
func double(number *int) int {
	// accessing to the int value
	result := *number * 2
	return result
}
