package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of spades" // type string is inferred by assignment
	card := newCard()
	// card = "Five of diamonds"

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
