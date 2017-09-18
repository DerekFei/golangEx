package main

func main() {
	cards := newDeck() //a slice of type strings
	cards.shuffle()
	cards.print()
}
