package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	cards := newDeck()
	// hand, remainCard := deal(cards, 5)
	// hand.print()
	// fmt.Println(hand.toString())
	// remainCard.print()
	// cards.saveToFile("my_cards")
	// newDeckfromfile("mycards").print()
	// cards.print()
	cards.shuffle()
	cards.print()
}
func caard() string {
	return "Ace of Spades"
}
