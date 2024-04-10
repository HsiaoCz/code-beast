package main

import "fmt"

type User struct {
	Email string
	Age   int
}

func main() {
	var user *User
	fmt.Println(user)
}
