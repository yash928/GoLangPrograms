package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}
	if d[0] != "Ace of Heart" {
		t.Errorf("Expected first card of Ace of Heart but got %v", d[0])
	}
	if d[len(d)-1] != "Five of Club" {
		t.Errorf("Expected last card of Five of Club but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndSaveFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckfromfile("_decktesting")
	if len(loadedDeck) != 20 {
		t.Errorf("Expected loaded deck length of 20 but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
