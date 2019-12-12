package main

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("deck_sands")

	// //DEAL
	// hand, remaingingCard := deal(cards, 4)

	//SAVE DECK TO FILE
	// cards.saveToFile("deck_sands")

	cards.print()

}
