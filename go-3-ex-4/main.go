package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	for _, r := range ranks {
		for _, s := range suits {
			fmt.Printf("%c%c\t", s, r)
		}
		fmt.Println()
	}

	cards := make(map[rune][]string, len(suits))
	for _, s := range suits {
		list := make([]string, 0, len(ranks))
		for _, r := range ranks {
			list = append(list, fmt.Sprintf("%c%c", s, r))
		}
		cards[s] = list
	}

	_ = cards
}
