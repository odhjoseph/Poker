package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const Hearts = 1
const Diamonds = 2
const Clubs = 3
const Spades = 4

//const Joker = 11, Queen = 12, King = 13
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

func dealCard(deck Deck) Card {
	deckHand += 1
	return deck[deckHand-1]
}

//rewrite in bytes
func showCard(suit int, cardNumber int) string {
	var cardFileName string

	switch cardNumber {

	case 1:
		cardFileName += "A"
	case 11:
		cardFileName += "J"
	case 12:
		cardFileName += "Q"
	case 13:
		cardFileName += "K"
	}

	//Switch 2-10???
	if cardNumber <= 10 && cardNumber > 1 {
		cardFileName += strconv.Itoa(cardNumber)
	}

	switch suit {

	case 1:
		cardFileName += "C"
	case 2:
		cardFileName += "H"
	case 3:
		cardFileName += "D"
	case 4:
		cardFileName += "S"
	}
	cardFileName += ".svg"

	return cardFileName

}

func main() {
	//deck creation
	var deck Deck
	//var cardImage *canvas.Image
	cardImage := canvas.NewImageFromFile("faces/" + "space.svg")
	//cardImage.SetMinSize(fyne.Size())

	deck = generatePlayingDeck(deck)
	Shuffle(deck)

	myApp := app.New()
	w := myApp.NewWindow("Poker")

	dealCard := widget.NewButton("Deal Card", func() {
		newCard := dealCard(deck)
		cardName := showCard(newCard.Suit, newCard.Number)
		cardImage = canvas.NewImageFromFile("faces/" + cardName)
		cardImage.FillMode = canvas.ImageFillOriginal
	})

	fmt.Println(dealCard)

	w.SetContent(
		cardImage,
	)

	w.Resize(200, 300)

	w.ShowAndRun()

}
