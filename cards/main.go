package main

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
