package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

//this is a receiver
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//this is a receiver
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//a function
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

//this is a receiver
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New((source))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
