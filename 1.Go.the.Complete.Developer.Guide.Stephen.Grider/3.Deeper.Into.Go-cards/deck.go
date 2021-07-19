package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//receiver of type deck
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

//function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//receiver of type deck
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//receiver of type deck
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

//receiver
func (d deck) shuffle() {
	//when we used before rand.Intn(len(d) - 1)
	// we always had the same order because we use the same number = len(d) -1

	//seed = source of randomness
	//time.Now().UnixNano() = Generates a unique 64 number = the seed
	source := rand.NewSource(time.Now().UnixNano())
	//we use the source as the basis to our random
	r := rand.New((source))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
