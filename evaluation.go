package main

func EvaluateHand(hand []Card) [5][14]int {
	var EvaluationMatrix [5][14]int // Initialize matrix with extra row and column for sums

	for _, card := range hand {
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
