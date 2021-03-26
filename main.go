//Aces not included currently
//Unique Hands

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Hearts = 1
const Spades = 2
const Clubs = 3
const Diamonds = 4

type Card struct {
	CardSuit int
	Number   int
}

func randomCard() Card {
	rand.Seed(time.Now().UnixNano())

	return Card{
		CardSuit: rand.Intn(4) + 1,
		Number:   rand.Intn(10) + 1,
	}
}

func generateFromDeck(amountofCards int) Card {

	return Card{}
}

func main() {
	Deck := []Card{}
}
