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

	jim.updateName("Jimmy")
	jim.print()
}

//receiver
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

//receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
