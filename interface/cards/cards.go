package main

import (
	"fmt"
	"math/rand"
	"time"
)

// const Hearts = 1
// const Diamonds = 2
// const Clubs = 3
// const Spades = 4
//Joker = 11, Queen = 12, King = 13
//Ace = 1

//Keep track of deck without having to alter array
var deckHand = 0

type Card struct {
	Suit   int
	Number int
}

type Deck []Card

func generatePlayingDeck(playingDeck Deck) Deck {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 13; j++ {
			card := Card{
				Suit:   i,
				Number: j,
			}
			playingDeck = append(playingDeck, card)
		}
	}
	return playingDeck
}

func Shuffle(slice Deck) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(slice); i > 0; i-- {
		randIndex := r.Intn(i)
		slice[i-1], slice[randIndex] = slice[randIndex], slice[i-1]
	}
}

func createTable(tabledeck Deck, deck Deck) Deck {
	return tabledeck
}

func main() {
	var deck Deck
	deck = generatePlayingDeck(deck)
	Shuffle(deck)
	fmt.Println(len(deck))
}
