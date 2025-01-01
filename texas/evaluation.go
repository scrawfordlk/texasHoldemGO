package texas

// EvaluateFullHand evaluates the hand and returns the hand type as HandVal
func EvaluateFullHand(hand Hand) HandVal {
	evaluationMatrix := EvaluateHand(hand)

	if IsFlush(evaluationMatrix) && IsStraight(evaluationMatrix) {
		// If the hand is a straight flush, check for royal flush (Ace-high straight flush)
		if evaluationMatrix[4][12] > 0 { // Ace present in the hand
			return RoyalFlush
		}
		return StraightFlush
	} else if IsFourOfAKind(evaluationMatrix) {
		return FourOfAKind
	} else if IsFullHouse(evaluationMatrix) {
		return FullHouse
	} else if IsFlush(evaluationMatrix) {
		return Flush
	} else if IsStraight(evaluationMatrix) {
		return Straight
	} else if IsThreeOfAKind(evaluationMatrix) {
		return ThreeOfAKind
	} else if IsTwoPair(evaluationMatrix) {
		return TwoPair
	} else if IsOnePair(evaluationMatrix) {
		return OnePair
	} else {
		return HighCard
	}
}

func EvaluateHand(hand Hand) [5][14]int {
	var EvaluationMatrix [5][14]int // Initialize matrix with extra row and column for sums

	for _, card := range hand.Cards {
		EvaluationMatrix[card.Suit][card.Rank]++
		EvaluationMatrix[card.Suit][13]++ // Increment sum for the suit
		EvaluationMatrix[4][card.Rank]++  // Increment sum for the rank
		EvaluationMatrix[4][13]++         // Increment total sum
	}

	return EvaluationMatrix
}

func IsTwoPair(evaluationMatrix [5][14]int) bool {
	// Check sums of ranks in the extra row
	pairCount := 0
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] >= 2 {
			pairCount++
		}
	}

	// If there are at least two pairs, return true
	return pairCount >= 2
}

func IsFlush(evaluationMatrix [5][14]int) bool {
	for suit := 0; suit < 4; suit++ {
		if evaluationMatrix[suit][13] == 5 {
			return true
		}
	}

	return false
}

func IsStraight(evaluationMatrix [5][14]int) bool {
	s := ""

	// Build a binary representation of ranks
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] > 0 {
			s += "1"
		} else {
			s += "0"
		}
	}

	// Append the ace (as rank 0) to the end for wrap-around straight
	s = string(s[12]) + s

	// Check for "11111" in string
	for i := 0; i <= len(s)-5; i++ {
		if s[i:i+5] == "11111" {
			return true
		}
	}

	return false
}

func IsFourOfAKind(evaluationMatrix [5][14]int) bool {
	// Check the rank row (4th row) for a count of 4
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 4 {
			return true
		}
	}
	return false
}

func IsFullHouse(evaluationMatrix [5][14]int) bool {
	hasThree := false
	hasPair := false

	// Check the rank row (4th row) for a count of 3 or 2
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 3 {
			hasThree = true
		} else if evaluationMatrix[4][rank] == 2 {
			hasPair = true
		}
	}

	// A full house needs both a three of a kind and a pair
	return hasThree && hasPair
}

func IsThreeOfAKind(evaluationMatrix [5][14]int) bool {
	// Check the rank row (4th row) for a count of 3
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 3 {
			return true
		}
	}
	return false
}

func IsOnePair(evaluationMatrix [5][14]int) bool {
	// Check the rank row (4th row) for a count of 2
	pairCount := 0
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 2 {
			pairCount++
		}
	}

	// A hand with exactly one pair
	return pairCount == 1
}
