package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := &Person{"Batman", 35}
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d", p.Age)
}
