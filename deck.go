package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var cardSuites = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
var cardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func newDeck(AddJokers bool) deck {
	res := deck{}
	for _, s := range cardSuites {
		for _, v := range cardValues {
			res = append(res, v+" of "+s)
		}
	}
	if AddJokers {
		res = append(res, "Joker", "Joker")
	}
	return res
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, "\n")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) []string {
	read, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(read), "\n")
}

func (d deck) shuffle() deck {
	for i := 0; i < len(d); i++ {
		randomIndex := rand.Intn(len(d))
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
	return d
}

func shuffle(d deck) deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < len(d); i++ {
		randomIndex := r.Intn(len(d) - 1)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
	return d
}
