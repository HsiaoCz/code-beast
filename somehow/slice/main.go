package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	chen(s)
	fmt.Println(s)
}

func chen(s []int) {
	s[2] = 4
}
