// this is a package, the main package
package main

import "fmt"

// comment one line

/*
	multiline
	comment
*/

func main() {
	// Declaring some variables
	var dog, cat, number = "🐶", "🐱", 1
	rabbit := "🐰"
	cat, face := "cat", ":)"

	fmt.Println(dog, cat, number, rabbit, face)

	// Declaring  a constanst, PI
	const pi = 3.14
	fmt.Printf("Using constant %.2f", pi)

	// basic data types: bool, string, numeric
	var a bool = true
	fmt.Printf("\nTipo: %T Valor: %v ", a, a)

	var b string = "batman"
	fmt.Printf("\nTipo: %T Valor: %v ", b, b)

	// uint8 will have a max value of 255, uint onlt allows positive numbers
	var c uint8 = 100
	fmt.Printf("\nTipo: %T Valor: %v ", c, c)

	// int8 will have a  min value of -125 and a max value of 127
	var d int8 = 125
	fmt.Printf("\nTipo: %T Valor: %v ", d, d)

	// byte is an alias for uint8 (0 to 255)
	var e byte = 254
	fmt.Printf("\nTipo: %T Valor: %v ", e, e)

	// rune is an alias for int32 (-2147483648 to 2147483647), represents a unicode string
	var f rune = 'a'
	fmt.Printf("\nTipo: %T Valor: %v ", f, f)

	var g float64 = 23.56
	fmt.Printf("\nTipo: %T Valor: %v ", g, g)

	var n1 uint8 = 100
	var n2 uint16 = 3000
	s := uint16(n1) + n2
	fmt.Printf("\nTipo: %T Valor: %v ", s, s)
	fmt.Printf("\n%T", n1)

	_ = 234
	var _ string = "test"

	var m string
	fmt.Printf("\nTipo: %T, %q ", m, m)

	var n int
	fmt.Printf("\nTipo: %T, %v ", n, n)

}
