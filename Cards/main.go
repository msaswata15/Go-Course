package main

import "fmt"

func main() {
	cards := deck{"Five of Diamonds"}
	cards = append(cards, "Six of spades")

	for i, c := range deal(newDeck(), 10) {
		fmt.Println()

		fmt.Println("Printing ", i+1, "Half")
		fmt.Println()

		for _, v := range c {
			fmt.Println(v)
		}
	}
	fmt.Println(newDeck().toString())
	newDeck().saveToFile("My_cards")
	newDeckFromFile("My_cards").print()
}
func newCard() {
	fmt.Println("Five of Diamonds")
}
