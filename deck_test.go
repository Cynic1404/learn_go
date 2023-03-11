package main

import "testing"

func TestNewDeck54(t *testing.T) {
	d := newDeck(true)
	if len(d) != 54 {
		t.Error("wrong amount of cards")
	}

}

func TestNewDeck52(t *testing.T) {
	d := newDeck(false)
	if len(d) != 52 {
		t.Error("wrong amount of cards")
	}
}
