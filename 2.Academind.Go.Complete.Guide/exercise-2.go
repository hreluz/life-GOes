package main

import "fmt"

func main() {
	var pi float64 = 3.14
	radius := 5

	circumference := 2 * pi * float64(radius)
	fmt.Println(circumference)

	message := fmt.Sprintf("For a radius of %v, the circle circumference is %2f", radius, circumference)
	fmt.Println(message)

}
