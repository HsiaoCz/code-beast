package main

import (
	"fmt"
	"time"
)

func main() {
	msgch := make(chan string, 1)
	msgch <- fetchResource(10)
	result := <-msgch
	fmt.Println(result)

	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3

	for num := range ch {
		fmt.Println("the number ---->", num)
	}
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
