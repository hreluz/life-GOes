package main

import "fmt"

//Array is a fixed number
//Slice is an array without length

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
