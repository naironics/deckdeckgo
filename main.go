package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 12)

	//Printing hand and remainingCards
	hand.print()
	remainingCards.print()

	// Printing all the 52 cards in deck
	cards.print()
}
