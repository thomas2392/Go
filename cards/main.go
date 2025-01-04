package main

import "fmt"

var deckSize int

func main() {
	cards := []string{"Ace of Diamond", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
