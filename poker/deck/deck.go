package deck

type Suit int

const (
	Spades Suit = iota
	Harts
	Diamonds
	Clibs
)

type Card struct {
	Suit  Suit
	Value int
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
