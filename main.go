package main

func main() {
	ourDeck := newDeck(true)
	ourDeck.print()
	//ourDeck.shuffle()
	shuffle(ourDeck)
	ourDeck.print()

}
