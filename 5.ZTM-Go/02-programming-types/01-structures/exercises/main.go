//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func (r Rectangle) width() int {
	return r.b.x - r.a.x
}

func (r Rectangle) length() int {
	return r.a.y - r.b.y
}

func (r Rectangle) area() int {
	return r.width() * r.length()
}

func (r Rectangle) printArea() {
	fmt.Printf("The area is %d\n", r.area())
}

func (r Rectangle) perimeter() int {
	return r.width()*2 + r.length()*2
}

func (r Rectangle) printPerimeter() {
	fmt.Printf("The perimeter is %d\n", r.perimeter())
}

func main() {
	r := Rectangle{Coordinate{0, 7}, Coordinate{10, 0}}
	r.printArea()
	r.printPerimeter()
}
