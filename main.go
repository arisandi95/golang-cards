package main

func main() {
	// NEW DECK 
	cards := newDeck()
	
	// NEW DECK FROM FILE
	// cards := newDeckFromFile("deck_sands")

	//// DEAL
	// hand, remaingingCard := deal(cards, 4)

	//SAVE DECK TO FILE
	// cards.saveToFile("deck_sands")

	// KOCOK KARTU / SWIPE
	cards.swipe()

	//CETAK
	// cards.print()

}
