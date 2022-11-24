package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName)
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@party.com",
			zipCode: 9400,
		},
	}

	// jim.print()
	// jimPointer := &jim
	// jimPointer.updateName("Hector")
	// jim.print()

	// Same as:

	jim.updateName("Hector")
	jim.print()
}

// * person  = this is a type desc . It means we are working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstNaame string) {
	(*pointerToPerson).firstName = newFirstNaame
	// *pointerToPerson - This is an operator - it means we want to manipulate the value the pointer is referencing
}

//receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
