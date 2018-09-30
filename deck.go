package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

// Method to write / save deck contents to a file [0666 = All Permissions]
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Method to read deck from a file
func readDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// if for some reason, the file reading fails, we handle it here
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
