package main

import (
	"fmt"
	"testing"
)

func TestOrderbook(t *testing.T) {
	ob := NewOrderbook()

	buyOrder := NewOrder(true, 10)
	ob.PlaceOrder(18_000, buyOrder)

	fmt.Printf("%+v", ob.Bids[0])
}

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 10)
	buyOrderC := NewOrder(true, 9)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)
	fmt.Println(l)
}
