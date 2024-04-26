package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hello I am %s\n", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Bye I am %s\n", p)
}

func (p Person) String() string {
	return "Hello I am a person and my name is " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hello I am %s\n", t)
}

func (t Text) Bye() {
	fmt.Printf("Bye I am %s\n", t)
}

// set clear that the gs must need to implete the methods on interface Greeter
func greetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\tValue: %v, Type : %T\n", g, g)
	}
}

func byeAll(bs ...Byer) {
	for _, b := range bs {
		b.Bye()
		fmt.Printf("\tValue: %v, Type : %T\n", b, b)
	}
}

type GreeterByer interface {
	Greeter
	Byer
}

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}

}

func main() {
	// var g Greeter = Person{Name: "Batman"}
	// var g Greeter = Text("Bruce Wayne")
	// g.Greet()

	p := Person{Name: "Superman"}
	t := Text("Wonder Woman")

	// greetAll(p, t)
	// byeAll(p, t)
	All(p, t)

	fmt.Println(p)
}
