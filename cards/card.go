package cards

type Cards struct {
	Number   int
	CardSuit Suit
}

type Suit struct {
	Hearts   string
	Spades   string
	Clubs    string
	Diamonds string
}
