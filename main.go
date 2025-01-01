package main

import (
	"fmt"
	. "texasHoldEmGO/texas"
)

func main() {
	// Define various test hands to evaluate

	// 1. Royal Flush (Ace-high straight flush)
	royalFlush := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Ten},
			{Suit: Hearts, Rank: Jack},
			{Suit: Hearts, Rank: Queen},
			{Suit: Hearts, Rank: King},
			{Suit: Hearts, Rank: Ace},
		},
	}
	fmt.Println("Royal Flush:", EvaluateFullHand(royalFlush)) // Expect: RoyalFlush

	// 2. Straight Flush (Five consecutive cards of the same suit)
	straightFlush := Hand{
		Cards: [5]Card{
			{Suit: Spades, Rank: Seven},
			{Suit: Spades, Rank: Eight},
			{Suit: Spades, Rank: Nine},
			{Suit: Spades, Rank: Ten},
			{Suit: Spades, Rank: Jack},
		},
	}
	fmt.Println("Straight Flush:", EvaluateFullHand(straightFlush)) // Expect: StraightFlush

	// 3. Four of a Kind
	fourOfAKind := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Four},
			{Suit: Diamonds, Rank: Four},
			{Suit: Spades, Rank: Four},
			{Suit: Clubs, Rank: Four},
			{Suit: Hearts, Rank: Ace},
		},
	}
	fmt.Println("Four of a Kind:", EvaluateFullHand(fourOfAKind)) // Expect: FourOfAKind

	// 4. Full House (Three of a kind and a pair)
	fullHouse := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: King},
			{Suit: Diamonds, Rank: King},
			{Suit: Spades, Rank: King},
			{Suit: Clubs, Rank: Nine},
			{Suit: Hearts, Rank: Nine},
		},
	}
	fmt.Println("Full House:", EvaluateFullHand(fullHouse)) // Expect: FullHouse

	// 5. Flush (Five cards of the same suit)
	flush := Hand{
		Cards: [5]Card{
			{Suit: Diamonds, Rank: Two},
			{Suit: Diamonds, Rank: Five},
			{Suit: Diamonds, Rank: Eight},
			{Suit: Diamonds, Rank: Jack},
			{Suit: Diamonds, Rank: Queen},
		},
	}
	fmt.Println("Flush:", EvaluateFullHand(flush)) // Expect: Flush

	// 6. Straight (Five consecutive cards, not of the same suit)
	straight := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Seven},
			{Suit: Diamonds, Rank: Eight},
			{Suit: Spades, Rank: Nine},
			{Suit: Clubs, Rank: Ten},
			{Suit: Hearts, Rank: Jack},
		},
	}
	fmt.Println("Straight:", EvaluateFullHand(straight)) // Expect: Straight

	// 7. Three of a Kind
	threeOfAKind := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Ten},
			{Suit: Diamonds, Rank: Ten},
			{Suit: Spades, Rank: Ten},
			{Suit: Clubs, Rank: Three},
			{Suit: Hearts, Rank: Seven},
		},
	}
	fmt.Println("Three of a Kind:", EvaluateFullHand(threeOfAKind)) // Expect: ThreeOfAKind

	// 8. Two Pair
	twoPair := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Queen},
			{Suit: Diamonds, Rank: Queen},
			{Suit: Clubs, Rank: Seven},
			{Suit: Spades, Rank: Seven},
			{Suit: Hearts, Rank: Ace},
		},
	}
	fmt.Println("Two Pair:", EvaluateFullHand(twoPair)) // Expect: TwoPair

	// 9. One Pair
	onePair := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Eight},
			{Suit: Diamonds, Rank: Eight},
			{Suit: Spades, Rank: King},
			{Suit: Clubs, Rank: Queen},
			{Suit: Hearts, Rank: Four},
		},
	}
	fmt.Println("One Pair:", EvaluateFullHand(onePair)) // Expect: OnePair

	// 10. High Card (No pairs or other hands)
	highCard := Hand{
		Cards: [5]Card{
			{Suit: Hearts, Rank: Two},
			{Suit: Diamonds, Rank: Five},
			{Suit: Spades, Rank: Nine},
			{Suit: Clubs, Rank: Jack},
			{Suit: Hearts, Rank: King},
		},
	}
	fmt.Println("High Card:", EvaluateFullHand(highCard)) // Expect: HighCard
}
