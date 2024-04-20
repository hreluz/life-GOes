package main

import "fmt"

func main() {
	//(size cannot be changed)
	fmt.Println("------------- ARRAYS -------------")
	var food [3]string
	food[0] = "🍕"
	food[1] = "🍔"
	food[2] = "🌭"

	fmt.Println(food)

	food2 := [3]string{"🍕", "🍔", "🌭"}
	fmt.Println(food2)
}
