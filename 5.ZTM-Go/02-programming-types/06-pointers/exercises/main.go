//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Status bool

const (
	StatusActive   Status = true
	StatusInactive Status = false
)

type SecurityTag struct {
	status Status
}

func switchStatus(st *SecurityTag, s Status) {
	st.status = s
}

func checkout(sAll *[]SecurityTag) {
	for i := range *sAll {
		(*sAll)[i].status = StatusInactive
	}
}

// - Create at least 4 items, all with active security tags
// - Store them in a slice or array
// - Deactivate any one security tag in the array/slice
// - Call the checkout() function to deactivate all tags
// - Print out the array/slice after each change
func main() {

	sAll := make([]SecurityTag, 4)

	fmt.Println(sAll)

	switchStatus(&sAll[0], StatusActive)

	fmt.Println(sAll)

	checkout(&sAll)

	fmt.Println(sAll)
}
