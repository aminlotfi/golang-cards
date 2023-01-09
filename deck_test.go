package main

import (
    "testing"
    "os"
)

func TestNewDeck(t *testing.T) {
    d := newDeck()

    if len(d) != 52 {
        t.Errorf("Expected deck length of 52, but got %v", len(d))
    }

    if d[0] != "Ace of Spades" {
        t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
    }

    if d[len(d) - 1] != "King of Clubs" {
        t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d) - 1])
    }
}
// Testing the read and write to file
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
    os.Remove("_deckTesting")

    deck := newDeck()

    deck.saveToFile("_deckTesting")

    loadedDeck := newDeckFromFile("_deckTesting")

    if len(loadedDeck) != 52 {
        t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
    }

    os.Remove("_deckTesting")
}