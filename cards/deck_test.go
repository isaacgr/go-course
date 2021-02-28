package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
}

func TestFirstCard(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := newDeck()
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be 'Four of Clubs', but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
