package main

// import "fmt"

//Array is a fixed number
//Slice is an array without length

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	cards.saveToFile("my_cards")
}
