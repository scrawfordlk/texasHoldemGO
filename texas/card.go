package texas

// Define Suit and Rank as integers
type Suit int
type Rank int

// Define the suits
const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

// Define the ranks
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

// Define the Card struct
type Card struct {
	Suit Suit
	Rank Rank
}

// Define the Hand struct which contains 5 cards
type Hand struct {
	Cards       [5]Card
	HandVal     HandVal
	TieBreakers []int
}
