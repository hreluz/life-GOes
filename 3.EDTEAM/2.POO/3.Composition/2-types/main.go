package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// declare alias
type myAlias = course

// type definition, a type definition is always empty
type newCourse course

type newBool bool

func (b newBool) String() string {

	if b {
		return "It is TRUE"
	}

	return "It is False"
}

func main() {
	c := myAlias{name: "Go"}
	// c.Print()
	d := newCourse{name: "Go"}

	fmt.Printf("Type is: %T\n", c)
	fmt.Printf("Type is: %T\n", d)

	var b newBool = true
	fmt.Println(b.String())
}
