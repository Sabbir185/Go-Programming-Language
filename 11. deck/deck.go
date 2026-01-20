package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	deck := deck{}
	deckEles := []string{"Ace", "King", "Salve"}
	deckValues := []string{"One", "Two", "Three"}
	for _, ele := range deckEles {
		for _, value := range deckValues {
			deck = append(deck, value+" of "+ele)
		}
	}
	return deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",\n")
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",\n")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randomIndex := r.Intn(len(d) - 1)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}
