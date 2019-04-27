package main

import "fmt"

func main() {
	// card := newCard()
	// fmt.Println(card)
	cards := []string{"satu", "dua", "tiga"}
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

// create function newCard() with return of string
func newCard() string {
	return "Five of Diamonds"
}
