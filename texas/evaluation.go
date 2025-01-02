package texas

// EvaluateFullHand evaluates the hand and returns the hand type as HandVal
func EvaluateFullHand(hand *Hand) HandVal {
	evaluationMatrix := EvaluateHand(hand)
	hand.SortCards() // needed for Tiebreaker calc

	if IsFlush(hand, evaluationMatrix) && IsStraight(hand, evaluationMatrix) {
		// If the hand is a straight flush, check for royal flush (Ace-high straight flush)
		if evaluationMatrix[4][12] > 0 && evaluationMatrix[4][11] > 0 { // Ace and King present in the hand
			hand.HandVal = RoyalFlush
			return RoyalFlush
		}

		hand.HandVal = StraightFlush
		hand.TieBreakers[0] = hand.Cards[4].Rank
		return StraightFlush
	} else if IsFourOfAKind(hand, evaluationMatrix) {
		return FourOfAKind
	} else if IsFullHouse(hand, evaluationMatrix) {
		return FullHouse
	} else if IsFlush(hand, evaluationMatrix) {
		return Flush
	} else if IsStraight(hand, evaluationMatrix) {
		return Straight
	} else if IsThreeOfAKind(hand, evaluationMatrix) {
		return ThreeOfAKind
	} else if IsTwoPair(hand, evaluationMatrix) {
		return TwoPair
	} else if IsOnePair(hand, evaluationMatrix) {
		return OnePair
	} else {
		hand.TieBreakers[0] = hand.Cards[4].Rank
		hand.TieBreakers[1] = hand.Cards[3].Rank
		hand.TieBreakers[2] = hand.Cards[2].Rank
		hand.TieBreakers[3] = hand.Cards[1].Rank
		hand.TieBreakers[4] = hand.Cards[0].Rank
		return HighCard
	}
}

func EvaluateHand(hand *Hand) *[5][14]int {
	var EvaluationMatrix [5][14]int // Initialize matrix with extra row and column for sums

	for _, card := range hand.Cards {
		EvaluationMatrix[card.Suit][card.Rank]++
		EvaluationMatrix[card.Suit][13]++ // Increment sum for the suit
		EvaluationMatrix[4][card.Rank]++  // Increment sum for the rank
		EvaluationMatrix[4][13]++         // Increment total sum
	}

	return &EvaluationMatrix
}

func IsTwoPair(hand *Hand, evaluationMatrix *[5][14]int) bool {
	twoPairs := [2]Rank{Two}
	// Check sums of ranks in the extra row
	pairCount := 0
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] >= 2 {
			twoPairs[pairCount] = Rank(rank)
			pairCount++
		}
	}

	// If there are at least two pairs, return true
	if pairCount >= 2 {
		hand.HandVal = TwoPair
		if twoPairs[0] > twoPairs[1] {
			hand.TieBreakers[0] = twoPairs[0]
			hand.TieBreakers[1] = twoPairs[1]
		} else {
			hand.TieBreakers[0] = twoPairs[1]
			hand.TieBreakers[1] = twoPairs[0]
		}

		hand.TieBreakers[2] = getLoneRank(hand)
		return true
	} else {
		return false
	}
}

func IsFlush(hand *Hand, evaluationMatrix *[5][14]int) bool {
	for suit := 0; suit < 4; suit++ {
		if evaluationMatrix[suit][13] == 5 {
			hand.HandVal = Flush
			hand.TieBreakers[0] = hand.Cards[4].Rank
			hand.TieBreakers[1] = hand.Cards[3].Rank
			hand.TieBreakers[2] = hand.Cards[2].Rank
			hand.TieBreakers[3] = hand.Cards[1].Rank
			hand.TieBreakers[4] = hand.Cards[0].Rank
			return true
		}
	}

	return false
}

func IsStraight(hand *Hand, evaluationMatrix *[5][14]int) bool {
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
			hand.HandVal = Straight
			if hand.Cards[0].Rank == Two && hand.Cards[4].Rank == Ace {
				hand.TieBreakers[0] = hand.Cards[3].Rank
			} else {
				hand.TieBreakers[0] = hand.Cards[4].Rank
			}
			return true
		}
	}

	return false
}

func IsFourOfAKind(hand *Hand, evaluationMatrix *[5][14]int) bool {
	// Check the rank row (4th row) for a count of 4
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 4 {
			hand.HandVal = FourOfAKind
			hand.TieBreakers[0] = Rank(rank)
			for _, card := range hand.Cards {
				if card.Rank != Rank(rank) {
					hand.TieBreakers[1] = card.Rank
				}
			}
			return true
		}
	}
	return false
}

func getLoneRank(hand *Hand) Rank {
	arr := [13]int{0}
	for _, card := range hand.Cards {
		arr[card.Rank]++
	}

	for i := range arr {
		if arr[i] == 1 {
			return Rank(i)
		}
	}

	return Ace
}

func IsFullHouse(hand *Hand, evaluationMatrix *[5][14]int) bool {
	hasThree := false
	hasPair := false
	tripleRank := Ace
	pairRank := Ace

	// Check the rank row (4th row) for a count of 3 or 2
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 3 {
			hasThree = true
			tripleRank = Rank(rank)
		} else if evaluationMatrix[4][rank] == 2 {
			hasPair = true
			pairRank = Rank(pairRank)
		}
	}

	// A full house needs both a three of a kind and a pair
	if hasThree && hasPair {
		hand.HandVal = FullHouse
		hand.TieBreakers[0] = tripleRank
		hand.TieBreakers[1] = pairRank
		return true
	} else {
		return false
	}
}

func IsThreeOfAKind(hand *Hand, evaluationMatrix *[5][14]int) bool {
	// Check the rank row (4th row) for a count of 3
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 3 {
			hand.HandVal = ThreeOfAKind
			hand.TieBreakers[0] = Rank(rank)
			assignLowestTwoLoneCardsInDescOrder(hand)
			return true
		}
	}
	return false
}

func assignLowestTwoLoneCardsInDescOrder(hand *Hand) {
	arr := [13]int{0}
	for _, card := range hand.Cards {
		arr[card.Rank]++
	}

	count := 0
	for i := range arr {
		if arr[i] == 1 {
			hand.TieBreakers[2-count] = Rank(i)
		}
	}
}

func IsOnePair(hand *Hand, evaluationMatrix *[5][14]int) bool {
	// Check the rank row (4th row) for a count of 2
	pairRank := Two
	pairCount := 0
	for rank := 0; rank < 13; rank++ {
		if evaluationMatrix[4][rank] == 2 {
			pairRank = Rank(rank)
			pairCount++
		}
	}

	if pairCount == 1 {
		hand.HandVal = OnePair
		hand.TieBreakers[0] = pairRank
		assignAllButPairDesc(hand)
		return true
	} else {
		return false
	}
}

func assignAllButPairDesc(hand *Hand) {
	count := 4
	for _, card := range hand.Cards {
		if card.Rank != hand.TieBreakers[0] { // if not pair
			hand.TieBreakers[count] = card.Rank
			count = count - 1
		}
	}
}

func GetHighestRank(hand *Hand) Rank {
	rank := Two
	for _, card := range hand.Cards {
		if card.Rank > rank {
			rank = card.Rank
		}
	}

	hand.TieBreakers[0] = rank
	assignRest(hand)

	return rank
}

func assignRest(hand *Hand) {
	arr := [13]int{0}

	count := 0
	for i := range arr {
		if arr[12-i] == 1 {
			hand.TieBreakers[1+count] = Rank(i)
			count++
		}
	}
}
