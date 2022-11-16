package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice os strings
type deck []string

// Create and return a list of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d refers to this or self
//this is a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//a function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//this is areceiver
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
