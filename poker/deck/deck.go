package deck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Suit int

const (
	Spades Suit = iota
	Harts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Harts:
		return "HARTS"
	case Diamonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	default:
		panic("invalid card suit")
	}
}

type Card struct {
	Suit  Suit
	Value int
}

func (c Card) String() string {
	value := strconv.Itoa(c.Value)
	if value == "1" {
		value = "ACE"
	}
	return fmt.Sprintf("%s of %s %s", value, c.Suit, suitToUnicode(c.Suit))
}

type Deck [52]Card

func New() Deck {
	var (
		nSuits = 4
		nCards = 13
		d      = [52]Card{}
	)

	x := 0
	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			d[x] = NewCard(Suit(i), j+1)
			x++
		}
	}
	return Shuffle(d, rand.New(rand.NewSource(time.Now().UnixNano())))
}

func Shuffle(d Deck, randn *rand.Rand) Deck {
	for i := 0; i < len(d); i++ {
		r := randn.Intn(i + 1)

		if r != i {
			d[i], d[r] = d[r], d[i]
		}
	}
	return d
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("the value of the card cannot be higher than 13")
	}
	return Card{
		Suit:  s,
		Value: v,
	}
}

func suitToUnicode(s Suit) string {
	switch s {
	case Spades:
		return "♠"
	case Harts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("invalid card suit")
	}
}
