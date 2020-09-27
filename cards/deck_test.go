package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Checking the length of the deck must be 52
	if len(d) != 52 {
		t.Errorf("Expected length of deck is 52, but got %d", len(d))
	}

	// Checking the existance of some random cards
	cards := []string{"Q♦︎", "8♠︎", "8♣︎", "5♥︎", "K♣︎"}
	for _, c := range cards {
		if !Contains(d, c) {
			t.Errorf("Missing card %v", c)
		}
	}

	// Checking the first and last card of the deck.
	if !(d[0] == "A♦︎" && d[len(d)-1] == "K♠︎") {
		t.Errorf("Expected first card of the deck is %v and last card is %v", "A♦︎", "K♠︎")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	f := "_decktesting.deck"
	os.Remove(f)

	d := newDeck()
	d.saveToFile(f)

	loadedDeck := newDeckFromFile(f)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of deck is 52, but got %d", len(loadedDeck))
	}

	os.Remove(f)
}
