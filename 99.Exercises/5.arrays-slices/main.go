package main

import "fmt"

func main() {
	arrIntegers := [5]int{2, 3, 55, 71, 2}
	slice := arrIntegers[1:4]
	fmt.Println("Array: ", arrIntegers)
	fmt.Println("Slice: ", slice)
}
