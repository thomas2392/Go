package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamond", "Hearts", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four",
		"Five", "Six", "Seve", "Eigth", "Nine", "Jack", "Queen",
		"King"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			card := value + " of " + suite
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	data := []byte(d.toString())
	return os.WriteFile(fileName, data, 0666)
}
