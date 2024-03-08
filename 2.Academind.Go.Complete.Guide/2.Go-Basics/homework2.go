package main

import "fmt"

func main() {
	pi := 3.1415
	var radius float64 = 5

	circumference := 2 * pi * radius
	fmt.Println(circumference)
	fmt.Printf("For a radius of %v, the circle circumference is %.2f ", radius, circumference)
}
