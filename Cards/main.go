package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of diamonds"
	// card := newCard()
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "siz")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
