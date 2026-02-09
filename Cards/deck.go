package main

import "fmt"

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "K"}
	cards := deck{}
	for _, s := range cardSuits {
		for _, value := range cardValues {
			card := s + " of " + value
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i+1, c)
	}
}
