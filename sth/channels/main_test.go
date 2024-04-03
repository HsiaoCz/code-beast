package main

import (
	"fmt"
	"testing"
)

func TestChannels(t *testing.T) {
	msgch := make(chan int, 128)

	msgch <- 1
	msgch <- 2
	msgch <- 3

	msgch <- 4

	close(msgch)

	for {
		msg := <-msgch
		fmt.Println("the message ---->", msg)
	}
}
