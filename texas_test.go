package main

import (
	"testing"
	. "texasHoldEmGO/texas"
)

func TestHandComparisonUsingSortInterface(t *testing.T) {
	hand1s := []Hand{
		// High Card cases - correct
		{Cards: [5]Card{{Diamonds, Queen}, {Spades, Nine}, {Clubs, Seven}, {Diamonds, Six}, {Hearts, Four}}},
		{Cards: [5]Card{{Hearts, Four}, {Spades, Nine}, {Clubs, Seven}, {Diamonds, Six}, {Diamonds, Queen}}},
		{Cards: [5]Card{{Clubs, Ace}, {Spades, King}, {Spades, Nine}, {Diamonds, Six}, {Hearts, Four}}},
		{Cards: [5]Card{{Diamonds, Six}, {Clubs, Ace}, {Hearts, Four}, {Spades, King}, {Spades, Nine}}},
		{Cards: [5]Card{{Clubs, Ace}, {Spades, King}, {Spades, Nine}, {Diamonds, Six}, {Hearts, Four}}},
		{Cards: [5]Card{{Spades, Nine}, {Spades, King}, {Clubs, Ace}, {Diamonds, Six}, {Hearts, Four}}},

		// One Pair cases - correct
		{Cards: [5]Card{{Diamonds, King}, {Spades, King}, {Hearts, Ten}, {Clubs, Eight}, {Clubs, Seven}}},
		{Cards: [5]Card{{Diamonds, King}, {Hearts, Ten}, {Clubs, Eight}, {Clubs, Seven}, {Spades, King}}},
		{Cards: [5]Card{{Diamonds, King}, {Spades, King}, {Hearts, Ten}, {Clubs, Eight}, {Clubs, Seven}}},
		{Cards: [5]Card{{Clubs, Eight}, {Diamonds, King}, {Spades, King}, {Hearts, Ten}, {Clubs, Seven}}},
		{Cards: [5]Card{{Hearts, Ace}, {Diamonds, Ace}, {Spades, Ten}, {Clubs, Nine}, {Clubs, Six}}},
		{Cards: [5]Card{{Clubs, Six}, {Clubs, Nine}, {Spades, Ten}, {Diamonds, Ace}, {Hearts, Ace}}},

		// Two Pair cases - correct
		{Cards: [5]Card{{Hearts, Queen}, {Diamonds, Queen}, {Diamonds, Six}, {Hearts, Six}, {Spades, Ace}}},
		{Cards: [5]Card{{Spades, Ace}, {Hearts, Queen}, {Diamonds, Queen}, {Diamonds, Six}, {Hearts, Six}}},
		{Cards: [5]Card{{Hearts, Ace}, {Spades, Ace}, {Diamonds, Six}, {Hearts, Six}, {Clubs, King}}},
		{Cards: [5]Card{{Clubs, King}, {Diamonds, Six}, {Hearts, Six}, {Hearts, Ace}, {Spades, Ace}}},
		{Cards: [5]Card{{Hearts, Queen}, {Diamonds, Queen}, {Clubs, Six}, {Diamonds, Six}, {Spades, Ace}}},
		{Cards: [5]Card{{Clubs, Six}, {Diamonds, Six}, {Hearts, Queen}, {Diamonds, Queen}, {Spades, Ace}}},

		// Three of a Kind cases - correct
		{Cards: [5]Card{{Hearts, Jack}, {Spades, Jack}, {Spades, Jack}, {Spades, Ace}, {Clubs, Eight}}},
		{Cards: [5]Card{{Spades, Ace}, {Clubs, Eight}, {Hearts, Jack}, {Spades, Jack}, {Spades, Jack}}},
		{Cards: [5]Card{{Diamonds, Three}, {Hearts, Three}, {Clubs, Three}, {Spades, Ace}, {Spades, Jack}}},
		{Cards: [5]Card{{Diamonds, Three}, {Spades, Ace}, {Hearts, Three}, {Spades, Jack}, {Clubs, Three}}},
		{Cards: [5]Card{{Hearts, Ace}, {Spades, Ace}, {Diamonds, Ace}, {Hearts, Ten}, {Spades, Five}}},
		{Cards: [5]Card{{Hearts, Ace}, {Spades, Ace}, {Hearts, Ten}, {Spades, Five}, {Diamonds, Five}}},

		// Straight cases - correct
		{Cards: [5]Card{{Hearts, Three}, {Spades, Four}, {Clubs, Five}, {Spades, Six}, {Diamonds, Seven}}},
		{Cards: [5]Card{{Spades, Six}, {Diamonds, Seven}, {Hearts, Three}, {Spades, Four}, {Clubs, Five}}},
		{Cards: [5]Card{{Hearts, Three}, {Spades, Four}, {Clubs, Five}, {Spades, Six}, {Diamonds, Seven}}},
		{Cards: [5]Card{{Clubs, Five}, {Spades, Six}, {Diamonds, Seven}, {Hearts, Three}, {Spades, Four}}},
		{Cards: [5]Card{{Hearts, Ace}, {Hearts, Two}, {Hearts, Three}, {Spades, Four}, {Clubs, Five}}},
		{Cards: [5]Card{{Hearts, Two}, {Hearts, Three}, {Spades, Four}, {Clubs, Five}, {Hearts, Ace}}},

		// Flush cases - correct
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, King}, {Diamonds, Ace}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ace}, {Diamonds, Ten}, {Diamonds, King}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Jack}, {Diamonds, King}}},
		{Cards: [5]Card{{Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Jack}, {Diamonds, King}, {Diamonds, Three}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Two}, {Diamonds, Five}}},
		{Cards: [5]Card{{Diamonds, Two}, {Diamonds, Five}, {Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ten}}},

		// Full House cases - correct
		{Cards: [5]Card{{Hearts, Queen}, {Spades, Queen}, {Hearts, Ten}, {Diamonds, Ten}, {Spades, Ten}}},
		{Cards: [5]Card{{Hearts, Ten}, {Diamonds, Ten}, {Spades, Ten}, {Hearts, Queen}, {Spades, Queen}}},
		{Cards: [5]Card{{Hearts, Ace}, {Spades, Ace}, {Diamonds, Queen}, {Hearts, Queen}, {Spades, Queen}}},
		{Cards: [5]Card{{Diamonds, Queen}, {Hearts, Queen}, {Spades, Queen}, {Hearts, Ace}, {Spades, Ace}}},
		{Cards: [5]Card{{Hearts, Queen}, {Spades, Queen}, {Diamonds, Queen}, {Hearts, Ten}, {Diamonds, Ten}}},
		{Cards: [5]Card{{Hearts, Queen}, {Hearts, Ten}, {Diamonds, Ten}, {Spades, Queen}, {Diamonds, Queen}}},

		// Four of a Kind cases - correct
		{Cards: [5]Card{{Hearts, Ten}, {Spades, Ten}, {Clubs, Ten}, {Diamonds, Ten}, {Spades, Eight}}},
		{Cards: [5]Card{{Clubs, Ten}, {Diamonds, Ten}, {Spades, Eight}, {Hearts, Ten}, {Spades, Ten}}},
		{Cards: [5]Card{{Spades, Five}, {Diamonds, Five}, {Clubs, Five}, {Hearts, Five}, {Hearts, Ace}}},
		{Cards: [5]Card{{Hearts, Ace}, {Clubs, Five}, {Diamonds, Five}, {Spades, Five}, {Hearts, Five}}},
		{Cards: [5]Card{{Hearts, Ten}, {Spades, Ten}, {Clubs, Ten}, {Diamonds, Ten}, {Hearts, Ace}}},
		{Cards: [5]Card{{Hearts, Ten}, {Hearts, Ace}, {Clubs, Ten}, {Spades, Ten}, {Diamonds, Ten}}},

		// Straight Flush cases - correct
		{Cards: [5]Card{{Hearts, Three}, {Hearts, Four}, {Hearts, Five}, {Hearts, Six}, {Hearts, Seven}}},
		{Cards: [5]Card{{Hearts, Three}, {Hearts, Four}, {Hearts, Five}, {Hearts, Seven}, {Hearts, Six}}},
		{Cards: [5]Card{{Hearts, Three}, {Hearts, Four}, {Hearts, Five}, {Hearts, Six}, {Hearts, Seven}}},
		{Cards: [5]Card{{Hearts, Seven}, {Hearts, Six}, {Hearts, Four}, {Hearts, Three}, {Hearts, Five}}},
		{Cards: [5]Card{{Spades, Six}, {Spades, Seven}, {Spades, Eight}, {Spades, Nine}, {Spades, Ten}}},
		{Cards: [5]Card{{Spades, Six}, {Spades, Ten}, {Spades, Seven}, {Spades, Eight}, {Spades, Nine}}},

		// Royal Flush cases
		{Cards: [5]Card{{Diamonds, Ten}, {Diamonds, Jack}, {Diamonds, Queen}, {Diamonds, King}, {Diamonds, Ace}}},
	}

	hand2s := []Hand{
		// High Card cases - prob. correct
		{Cards: [5]Card{{Hearts, Queen}, {Spades, Nine}, {Clubs, Eight}, {Diamonds, Six}, {Hearts, Four}}},
		{Cards: [5]Card{{Clubs, Eight}, {Diamonds, Six}, {Hearts, Queen}, {Spades, Nine}, {Hearts, Four}}},
		{Cards: [5]Card{{Hearts, Ace}, {Spades, Queen}, {Spades, Nine}, {Diamonds, Six}, {Hearts, Four}}},
		{Cards: [5]Card{{Hearts, Ace}, {Diamonds, Six}, {Spades, Queen}, {Hearts, Four}, {Spades, Nine}}},
		{Cards: [5]Card{{Hearts, Ace}, {Spades, King}, {Spades, Nine}, {Diamonds, Six}, {Hearts, Four}}},
		{Cards: [5]Card{{Hearts, Four}, {Hearts, Ace}, {Spades, Nine}, {Diamonds, Six}, {Spades, King}}},

		// One Pair cases - prob. correct
		{Cards: [5]Card{{Hearts, Eight}, {Clubs, Eight}, {Spades, King}, {Hearts, Ten}, {Clubs, Seven}}},
		{Cards: [5]Card{{Hearts, Ten}, {Hearts, Eight}, {Spades, King}, {Clubs, Seven}, {Clubs, Eight}}},
		{Cards: [5]Card{{Hearts, King}, {Spades, King}, {Hearts, Ten}, {Clubs, Eight}, {Clubs, Seven}}},
		{Cards: [5]Card{{Hearts, King}, {Clubs, Eight}, {Clubs, Seven}, {Spades, King}, {Hearts, Ten}}},
		{Cards: [5]Card{{Hearts, Ace}, {Diamonds, Ace}, {Spades, Ten}, {Clubs, Nine}, {Hearts, Seven}}},
		{Cards: [5]Card{{Hearts, Ace}, {Diamonds, Ace}, {Clubs, Nine}, {Spades, Ten}, {Hearts, Seven}}},

		// Two Pair cases - prob. correct
		{Cards: [5]Card{{Spades, Queen}, {Diamonds, Queen}, {Diamonds, Six}, {Hearts, Six}, {Spades, Ace}}},
		{Cards: [5]Card{{Spades, Queen}, {Diamonds, Queen}, {Spades, Ace}, {Diamonds, Six}, {Hearts, Six}}},
		{Cards: [5]Card{{Clubs, Queen}, {Diamonds, Queen}, {Diamonds, Six}, {Hearts, Six}, {Spades, Ace}}},
		{Cards: [5]Card{{Clubs, Queen}, {Diamonds, Queen}, {Spades, Ace}, {Diamonds, Six}, {Hearts, Six}}},
		{Cards: [5]Card{{Clubs, Ace}, {Spades, Ace}, {Hearts, King}, {Clubs, King}, {Diamonds, Queen}}},
		{Cards: [5]Card{{Diamonds, Queen}, {Hearts, King}, {Clubs, King}, {Clubs, Ace}, {Spades, Ace}}},

		// Three of a Kind cases - prob. correct
		{Cards: [5]Card{{Diamonds, Three}, {Hearts, Three}, {Clubs, Three}, {Spades, Ace}, {Spades, Jack}}},
		{Cards: [5]Card{{Diamonds, Three}, {Spades, Ace}, {Spades, Jack}, {Hearts, Three}, {Clubs, Three}}},
		{Cards: [5]Card{{Diamonds, Three}, {Hearts, Three}, {Spades, Three}, {Spades, Ace}, {Spades, Jack}}},
		{Cards: [5]Card{{Spades, Ace}, {Diamonds, Three}, {Hearts, Three}, {Spades, Three}, {Spades, Jack}}},
		{Cards: [5]Card{{Hearts, Ace}, {Spades, Ace}, {Diamonds, Ace}, {Spades, King}, {Hearts, Ten}}},
		{Cards: [5]Card{{Spades, King}, {Hearts, Ace}, {Spades, Ace}, {Diamonds, Ace}, {Hearts, Ten}}},

		// Straight cases
		{Cards: [5]Card{{Hearts, Two}, {Hearts, Three}, {Spades, Four}, {Clubs, Five}, {Spades, Six}}},
		{Cards: [5]Card{{Hearts, Three}, {Hearts, Two}, {Clubs, Five}, {Spades, Four}, {Spades, Six}}},
		{Cards: [5]Card{{Hearts, Three}, {Spades, Four}, {Clubs, Five}, {Spades, Six}, {Hearts, Seven}}},
		{Cards: [5]Card{{Hearts, Three}, {Spades, Six}, {Hearts, Seven}, {Spades, Four}, {Clubs, Five}}},
		{Cards: [5]Card{{Hearts, Two}, {Hearts, Three}, {Spades, Four}, {Clubs, Five}, {Hearts, Six}}},
		{Cards: [5]Card{{Hearts, Three}, {Spades, Four}, {Clubs, Five}, {Hearts, Six}, {Hearts, Two}}},

		// Flush cases - correct
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Two}, {Diamonds, Queen}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Queen}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Two}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Jack}, {Diamonds, King}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, King}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Jack}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Six}, {Diamonds, Ten}, {Diamonds, Jack}, {Diamonds, Ace}}},
		{Cards: [5]Card{{Diamonds, Three}, {Diamonds, Jack}, {Diamonds, Ace}, {Diamonds, Six}, {Diamonds, Ten}}},

		// Full House cases - prob. correct
		{Cards: [5]Card{{Hearts, Queen}, {Spades, Queen}, {Clubs, Queen}, {Hearts, Ten}, {Diamonds, Ten}}},
		{Cards: [5]Card{{Hearts, Queen}, {Diamonds, Ten}, {Spades, Queen}, {Diamonds, Ten}, {Clubs, Queen}}},
		{Cards: [5]Card{{Diamonds, Ace}, {Spades, Ace}, {Clubs, Queen}, {Hearts, Queen}, {Spades, Queen}}},
		{Cards: [5]Card{{Diamonds, Ace}, {Hearts, Queen}, {Spades, Queen}, {Spades, Ace}, {Clubs, Queen}}},
		{Cards: [5]Card{{Hearts, Queen}, {Spades, Queen}, {Clubs, Ten}, {Diamonds, Ten}, {Clubs, Ten}}},
		{Cards: [5]Card{{Spades, Queen}, {Hearts, Ten}, {Spades, Queen}, {Diamonds, Ten}, {Clubs, Ten}}},

		// Four of a Kind cases - correct
		{Cards: [5]Card{{Hearts, Ten}, {Spades, Ten}, {Clubs, Ten}, {Diamonds, Ten}, {Hearts, King}}},
		{Cards: [5]Card{{Clubs, Ten}, {Diamonds, Ten}, {Hearts, King}, {Hearts, Ten}, {Spades, Ten}}},
		{Cards: [5]Card{{Spades, Five}, {Diamonds, Five}, {Clubs, Five}, {Hearts, Five}, {Hearts, Ace}}},
		{Cards: [5]Card{{Spades, Five}, {Diamonds, Five}, {Clubs, Five}, {Hearts, Five}, {Hearts, Ace}}},
		{Cards: [5]Card{{Hearts, Ten}, {Spades, Ten}, {Clubs, Ten}, {Diamonds, Ten}, {Diamonds, Jack}}},
		{Cards: [5]Card{{Spades, Ten}, {Clubs, Ten}, {Diamonds, Ten}, {Diamonds, Jack}, {Hearts, Ten}}},

		// Straight Flush cases - correct
		{Cards: [5]Card{{Hearts, Two}, {Hearts, Three}, {Hearts, Four}, {Hearts, Five}, {Hearts, Six}}},
		{Cards: [5]Card{{Hearts, Four}, {Hearts, Five}, {Hearts, Two}, {Hearts, Three}, {Hearts, Six}}},
		{Cards: [5]Card{{Hearts, Three}, {Hearts, Four}, {Hearts, Five}, {Hearts, Six}, {Hearts, Seven}}},
		{Cards: [5]Card{{Hearts, Three}, {Hearts, Seven}, {Hearts, Four}, {Hearts, Five}, {Hearts, Six}}},
		{Cards: [5]Card{{Spades, Seven}, {Spades, Eight}, {Spades, Nine}, {Spades, Ten}, {Spades, Jack}}},
		{Cards: [5]Card{{Spades, Seven}, {Spades, Eight}, {Spades, Jack}, {Spades, Nine}, {Spades, Ten}}},

		// Royal Flush cases - correct
		{Cards: [5]Card{{Diamonds, Ten}, {Diamonds, Jack}, {Diamonds, Queen}, {Diamonds, King}, {Diamonds, Ace}}},
	}

	expected := []string{
		"hand 2 > hand 1", // High Card
		"hand 2 > hand 1",
		"hand 1 > hand 2",
		"hand 1 > hand 2",
		"hand 1 = hand 2",
		"hand 1 = hand 2",

		"hand 1 > hand 2", // One Pair
		"hand 1 > hand 2",
		"hand 1 = hand 2",
		"hand 1 = hand 2",
		"hand 2 > hand 1",
		"hand 2 > hand 1",

		"hand 1 = hand 2", // Two Pair
		"hand 1 = hand 2",
		"hand 1 > hand 2",
		"hand 1 > hand 2",
		"hand 2 > hand 1",
		"hand 2 > hand 1",

		"hand 1 > hand 2", // Three of a Kind
		"hand 1 > hand 2",
		"hand 1 = hand 2",
		"hand 1 = hand 2",
		"hand 2 > hand 1",
		"hand 2 > hand 1",

		"hand 1 > hand 2", // Straight
		"hand 1 > hand 2",
		"hand 1 = hand 2",
		"hand 1 = hand 2",
		"hand 2 > hand 1",
		"hand 2 > hand 1",

		"hand 1 > hand 2", // Flush
		"hand 1 > hand 2",
		"hand 1 = hand 2",
		"hand 1 = hand 2",
		"hand 2 > hand 1",
		"hand 2 > hand 1",

		"hand 2 > hand 1", // Full House
		"hand 2 > hand 1",
		"hand 1 = hand 2",
		"hand 1 = hand 2",
		"hand 1 > hand 2",
		"hand 1 > hand 2",

		"hand 2 > hand 1",
		"hand 2 > hand 1",
		"hand 1 = hand 2",
		"hand 1 = hand 2",
		"hand 1 > hand 2", // Four of A Kind
		"hand 1 > hand 2",

		"hand 1 > hand 2", // Straight Flush
		"hand 1 > hand 2",
		"hand 1 = hand 2",
		"hand 1 = hand 2",
		"hand 2 > hand 1",
		"hand 2 > hand 1",

		"hand 1 = hand 2", // Royal Flush
	}

	tests := [55]struct {
		hand1    Hand
		hand2    Hand
		expected string // "hand1 < hand2", "hand1 > hand2", or "hand1 == hand2"
	}{}

	print(len(tests))
	for i := range tests {
		tests[i].hand1 = hand1s[i]
		tests[i].hand2 = hand2s[i]
		tests[i].expected = expected[i]
	}

	for _, test := range tests {
		// Evaluate the hands
		EvaluateFullHand(&test.hand1)
		EvaluateFullHand(&test.hand2)

		// Compare using Less
		handList := HandList{test.hand1, test.hand2}
		hand1LessHand2 := handList.Less(0, 1)
		hand2LessHand1 := handList.Less(1, 0)

		// Determine the result
		var result string
		if hand1LessHand2 {
			result = "hand 2 > hand 1"
		} else if hand2LessHand1 {
			result = "hand 1 > hand 2"
		} else {
			result = "hand 1 = hand 2"
		}

		// Check the result
		if result != test.expected {
			t.Errorf("For hands %v and %v, expected %s, but got %s", test.hand1, test.hand2, test.expected, result)
		}
	}
}
