package main

type Suit int
type Rank int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Suit Suit
	Rank Rank
}
