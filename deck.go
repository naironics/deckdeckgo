package main

import (
	"fmt"
	"strings"
)

// Creating a type deck
type deck []string

// Method to return a new deck of 52 cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Method to deal cards, returns cards dealt of handSize and remainingCards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Method to print cards in the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Method to convert a deck to a String
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
