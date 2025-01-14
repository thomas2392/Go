package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck lenght to be 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as first card, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Club" {
		t.Errorf("Expected King of Club as last card, but got %v", d[len(d)-1])
	}
}
