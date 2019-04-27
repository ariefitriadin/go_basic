package main

import "fmt"

// create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck {
	// initial var cards with empty type of deck
	// it can be done like this as well :   var cards deck
	// var cards deck
	cards := deck{}
	// initialize var cardSuits with slice of type string
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// the sign "_" in looping, is as an exchanges with explicit variable name like i or j or etc
	// which is we must use that explicit var name within loop
	// with sign "_" we don't have to use that variable, it is used just for index process internally
	// in a loop
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// create receiver for new function print() in type of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
