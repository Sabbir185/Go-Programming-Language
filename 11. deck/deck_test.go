package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 9 {
		t.Errorf("Expected deck length of 9, but got %d", len(deck))
	}
	if deck[0] != "One of Ace" {
		t.Errorf("Expected first card to be 'One of Ace', but got '%s'", deck[0])
	}
	if deck[len(deck)-1] != "Three of Salve" {
		t.Errorf("Expected last card to be 'Three of Salve', but got '%s'", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Errorf("Expected deck length is 9, but get %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
