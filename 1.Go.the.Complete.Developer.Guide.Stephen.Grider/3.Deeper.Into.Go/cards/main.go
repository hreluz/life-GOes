package main

//Array is a fixed number
//Slice is an array without length

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
