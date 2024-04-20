package main

import (
	"fmt"
)

func main() {
	// pointers
	fmt.Println("------------- POINTERS -------------")
	fruit := "🍎"
	var p *string
	p = &fruit
	fmt.Printf("Type %T, Value: %s, Address: %v", fruit, fruit, &fruit)
	fmt.Printf("\nType %T, Value: %s, Derefence: %v", p, *p, p)
	fmt.Print("\n", &fruit == p)

	*p = "🍏"
	fmt.Printf("Type %T, Value: %s, Address: %v", fruit, fruit, &fruit)
	fmt.Printf("\nType %T, Value: %s, Derefence: %v", p, *p, p)
}
