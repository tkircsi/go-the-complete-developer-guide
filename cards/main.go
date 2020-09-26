package main

func main() {
	cards := newDeck()
	cards.shuffle()

	hand, _ := deal(cards, 5)
	hand.printf()
	// hand.saveToFile("hand")
	// remainingDeck.saveToFile("remdeck")

}
