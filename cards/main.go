package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of spades" // type string is inferred by assignment
	// card := newCard()
	// card = "Five of diamonds"

	// fmt.Println(card)

	// slice
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
