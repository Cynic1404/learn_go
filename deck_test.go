package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeckWithJokers(t *testing.T) {
	d := newDeck(true)
	if len(d) != 54 {
		t.Error("wrong amount of cards")
	}
	jokerCounter := 0
	for i := 0; i < len(d); i++ {
		if d[i] == "Joker" {
			jokerCounter++
		}
	}
	if jokerCounter < 2 {
		fmt.Println("Less than 2 Jokers")
	}
}

func TestNewDeck52(t *testing.T) {
	d := newDeck(false)
	if len(d) != 52 {
		t.Error("wrong amount of cards")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck(true)
	d.saveToFile("_decktesting")
	deckFromFile := newDeckFromFile("_decktesting")
	fmt.Println(len(deckFromFile) == 54)
}
