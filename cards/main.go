package main

func main() {
	cards := newDeckFromFile("my_cards") //a slice of type strings
	cards.print()
}
