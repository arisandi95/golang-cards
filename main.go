package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card :=  newCard()
	card := []string{newCard(), "Five of hearts"}

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}