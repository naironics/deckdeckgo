package main

import "fmt"

// Creating a type deck
type deck []string

// Method to print cards in the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
