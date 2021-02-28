package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of "deck", which is a slice of strings
type deck []string

func newDeck() deck {
	deck := deck{}
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, value+" of "+suit)
		}
	}
	return deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	if len(d) < 1 {
		d = newDeck()
	}
	hand := d[:handSize]
	remainingDeck := d[handSize:]
	return hand, remainingDeck
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func readFromFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func loadDeck(filename string) deck {
	savedDeck := strings.Split(readFromFile((filename)), ",")
	return deck(savedDeck)
}
