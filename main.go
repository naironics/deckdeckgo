package main

func main() {
	// reading deck of cards from file
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
