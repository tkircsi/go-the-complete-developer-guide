package main

import "fmt"

func main() {
	remainingDeck := newDeck()
	remainingDeck.shuffle()
	var hand deck

	for i := 0; i < 2; i++ {
		hand, remainingDeck = deal(remainingDeck, 5)
		fmt.Print("Hand ", i, ": ")
		hand.printf()
	}
	fmt.Println(remainingDeck)

	v := Contains(remainingDeck, "6♠︎")
	fmt.Println(v)
}
