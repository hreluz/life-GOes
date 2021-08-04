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
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Harpert",
		contactInfo: contactInfo{
			email:   "jim@harpert.com",
			zipCode: 833353,
		},
	}

	//&variable = give me the memory address of the value this variable is pointing at
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

//receiver
//*pointer = give me the value this memory address is pointing at
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
