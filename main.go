package main

func main() {
	// reading deck of cards from file
	cards := readDeckFromFile("mycards.txt")
	cards.print()
}
