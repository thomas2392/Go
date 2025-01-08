package main

var deckSize int

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
