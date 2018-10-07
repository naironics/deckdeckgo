package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Checking if deck has 52 cards
	if len(d) != 52 {
		t.Errorf("Expected 52 cards, but found %v", len(d))
	}

	// Checking if the first card in the deck is "Ace of Spades"
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as first card in the deck, but found %v", d[0])
	}

	// Checking if the last card in the deck is "King of Diamonds"
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected King of Diamonds as the last card in the deck, but found %v", d[len(d)-1])
	}

}
