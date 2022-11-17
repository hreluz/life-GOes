package main

// import "fmt"

//Array is a fixed number
//Slice is an array without length

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}
