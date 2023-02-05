package main

import "fmt"

func main() {
	cards := newDeck()
	// cards := newDeckFromFile(("my_cards"))
	fmt.Println(deal(cards, 8))
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")

}
