package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "os"
    "time"
    "math/rand"
)

type deck []string

func newDeck() deck {
    cards := deck{}

    cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

    for _, suite := range cardSuites {
        for _, value := range cardValues {
            cards = append(cards, value + " of " + suite)
        }
    }

    return cards
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i + 1, card)
    }
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    return deck(strings.Split(string(bs), ","))
}