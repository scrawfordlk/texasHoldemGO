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
		hand.HandVal = EvaluateFullHand(hand)
	}

	// Compare hand values first
	if h[i].HandVal != h[j].HandVal {
		return h[i].HandVal < h[j].HandVal
	}

	// If hand values are the same, compare tie-breakers
	for idx := 0; idx < len(h[i].TieBreakers) && idx < len(h[j].TieBreakers); idx++ {
		if h[i].TieBreakers[idx] > h[j].TieBreakers[idx] {
			return true
		} else if h[i].TieBreakers[idx] < h[j].TieBreakers[idx] {
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

// Example usage of the above code
func main() {
	// Sample hands (for demonstration purposes)
	hand1 := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Ten},
			{Suit: Hearts, Rank: Jack},
			{Suit: Hearts, Rank: Queen},
			{Suit: Hearts, Rank: King},
			{Suit: Hearts, Rank: Ace},
		},
		HandVal:     RoyalFlush,
		TieBreakers: []int{Ace, King, Queen, Jack, Ten},
	}

	hand2 := Hand{
		Cards: [5]Card{
			{Suit: Spades, Rank: Ten},
			{Suit: Spades, Rank: Jack},
			{Suit: Spades, Rank: Queen},
			{Suit: Spades, Rank: King},
			{Suit: Spades, Rank: Ace},
		},
		HandVal:     RoyalFlush,
		TieBreakers: []int{Ace, King, Queen, Jack, Ten},
	}

	hand3 := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Ace},
			{Suit: Hearts, Rank: King},
			{Suit: Hearts, Rank: Queen},
			{Suit: Hearts, Rank: Jack},
			{Suit: Hearts, Rank: Ten},
		},
		HandVal:     RoyalFlush,
		TieBreakers: []int{Ace, King, Queen, Jack, Ten},
	}

	hands := []Hand{hand1, hand2, hand3}

	// Sort hands
	SortHands(hands)

	// Print sorted hands
	for _, hand := range hands {
		fmt.Println(hand)
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
