package texas

import (
	"fmt"
	"sort"
)

// Implement Len(), Less(i, j int), Swap(i, j int) for sort.Interface
type HandList []Hand

func (h HandList) Len() int {
	return len(h)
}

func (h HandList) Less(i, j int) bool {
	for _, hand := range h {
		hand.HandVal = EvaluateFullHand(&hand)
	}

	// Compare hand values first
	if h[i].HandVal != h[j].HandVal {
		return h[i].HandVal < h[j].HandVal
	}

	// If hand values are the same, compare tie-breakers
	for idx := 0; idx < len(h[i].TieBreakers) && idx < len(h[j].TieBreakers); idx++ {
		if h[i].TieBreakers[idx] < h[j].TieBreakers[idx] {
			return true
		} else if h[i].TieBreakers[idx] > h[j].TieBreakers[idx] {
			return false
		}
	}

	// If all tie-breakers are the same, hands are equal
	return false
}

func (h HandList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Function to sort the hands
func SortHands(hands []Hand) {
	sort.Sort(HandList(hands))
}

func (s Suit) String() string {
	switch s {
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Spades:
		return "Spades"
	default:
		return "Unknown"
	}
}

func (r Rank) String() string {
	switch r {
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		return "Unknown"
	}
}

// String method for printing hands
func (h Hand) String() string {
	result := ""
	for _, card := range h.Cards {
		result += fmt.Sprintf("%v-%v ", card.Suit, card.Rank)
	}
	return result
}

// Sorting function for cards (sort by Rank, then Suit)
func (h *Hand) SortCards() {
	sort.Slice(h.Cards[:], func(i, j int) bool {
		// First sort by Rank
		if h.Cards[i].Rank != h.Cards[j].Rank {
			return h.Cards[i].Rank < h.Cards[j].Rank
		}
		// If ranks are equal, sort by Suit
		return h.Cards[i].Suit < h.Cards[j].Suit
	})
}
