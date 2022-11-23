package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName)
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)

}

// When no assigning values there are zero values
// TYPE 	ZERO VALUE
// STRING 	""
// INT 	0
// FLOAT	0
// BOOL 	FALSE
