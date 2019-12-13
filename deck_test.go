package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	firstCard := "Clubs Of Ace"
	lastCard := "Spades Of King"

	if len(deck) != 52 {
		t.Errorf("Gagal, seharusnya output 52 , output saat ini %v", len(deck))
	}

	if deck[0] != firstCard {
		t.Errorf("kartu lu bukan "+firstCard+", tapi ini %v", deck[0])
	}

	if deck[len(deck)-1] != lastCard {
		t.Errorf("kartu lu bukan "+lastCard+", tapi ini %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	openDeck := newDeckFromFile("_decktesting")
	if len(openDeck) != 52 {
		t.Errorf("Gagal, seharusnya output 52 , output saat ini %v", len(deck))
	}

	os.Remove("_decktesting")
}

//  LESSON 33