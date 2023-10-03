package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Heart", "Spade", "Diamond", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), " ,")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckfromfile(filename string) deck {
	bslice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bslice), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newpos := r.Intn(len(d) - 1)
		d[i], d[newpos] = d[newpos], d[i]
	}
}
