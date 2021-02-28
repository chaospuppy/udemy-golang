package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("newDeck should return deck of length 52; got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in deck to be Ace of Spades, got %s", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card in deck to be King of Diamonds, got %s", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	testFile := "_decktest"
	os.Remove(testFile)

	d := newDeck()
	d.saveToFile(testFile)
	ld := loadDeck(testFile)

	if len(ld) != 52 {
		t.Errorf("newDeck should return deck of length 52; got %d", len(d))
	}

	os.Remove(testFile)
}
