package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of diamonds"
	// card := newCard()
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "siz")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)

	// cards := newDeck()
	// // cards.print()

	// // Multiple Return
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// Type Conversion from string to by byte array only  (not to int it will not work in this case)
	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))
	// cards := newDeck()
	// fmt.Println(cards.toString())

	//  save file
	// fmt.Println(cards.saveToFile("sanket.txt"))

	// Read File
	// fmt.Println(newDeckFromFile("sanket.txt"))
	// fmt.Println(newDeckFromFile("sankt.txt"))

	// shuffle card

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
