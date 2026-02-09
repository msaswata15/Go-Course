package main

import (
	"fmt"
	"os"

	"strings"
)

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
func deal(d deck, hsize int) []deck {
	hand := d[:hsize]
	remaining := d[hsize:]

	return []deck{hand, remaining}
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) print() {
	for i, c := range d {
		fmt.Println(i+1, c)
	}
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return deck{}
	}

	s := string(bs)
	cards := strings.Split(s, ",")
	return deck(cards)
}
