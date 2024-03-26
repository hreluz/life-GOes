package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age)

	myAge := &age

	fmt.Println(myAge)
	fmt.Println(*myAge)

	*myAge = 33
	fmt.Println(age)
	fmt.Println(*myAge)

	doubledAge := double(myAge)
	fmt.Println(doubledAge)
	fmt.Println(age)
	fmt.Println(myAge)
}

// receiving a pointer
func double(number *int) int {
	//accessing the value
	result := *number * 2
	return result
}
