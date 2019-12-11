package main

import "fmt"

type deck []string

func newDeck() deck{
	cards := deck{}

	var cardSuits = []string {"Clubs", "Diamonds", "Hearts", "Spades"}
	var cardValues = []string {"Ace", "2", "3","4", "5", "6", "7", "8", 
								"9", "10", "Jack", "Queen", "King"}
	
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuit+" Of "+cardValue)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	} 
}

func deal(cards deck, handOfSize int) (deck, deck){
	return cards[:handOfSize], cards[handOfSize:]
}