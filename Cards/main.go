package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of diamonds"
	// card := newCard()
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "siz")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
