package main

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of spades" // type string is inferred by assignment
	// card := newCard()
	// card = "Five of diamonds"

	// fmt.Println(card)

	// slice
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
