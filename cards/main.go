package main

var deckSize int

func main() {
	cards := newDeck()
	cards.saveToFile("my_deck")
	//fmt.Println(cards.toString())
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()
}
