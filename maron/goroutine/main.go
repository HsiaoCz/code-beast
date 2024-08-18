package main

import (
	"fmt"
	"time"
)

func main() {
	// result :=fetchResource()
	// fmt.Println("the result:", result)
	// channel always blocking when is full
	resultch := make(chan string)

	go fetchResource(1, resultch)

	result := <-resultch
	fmt.Println(result)
	time.Sleep(time.Second * 3)
}

func fetchResource(n int, resultch chan string) {
	time.Sleep(time.Second * 2)
	resultch <- fmt.Sprintf("result %d", n)
}
