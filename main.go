package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card :=  newCard()
	cards := newDeck()

	//Deal
	hand, remaingingCard := deal(cards, 4)

	cards.print()
	fmt.Println("========================")
	hand.print()
	fmt.Println("========================")
	remaingingCard.print()
}
