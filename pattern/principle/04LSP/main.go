package main

import "fmt"

// 里氏代换原则
// 基类适用的，子类一定适用（子类可以扩展父类的功能，但不能改变父类原有的功能）
// 这个在go里面似乎不太需要这样考虑

type ClassA struct{}

func (ca *ClassA) Add(x int, y int) int {
	return x + y
}

type ClassB struct {
	ClassA
}

func (cb *ClassB) Add(a string, b string) string {
	return a + b
}

func main() {
	a := 10
	b := 20
	cb := ClassB{}
	fmt.Println(cb.ClassA.Add(a, b))
	x := "hello,"
	y := "hi"
	fmt.Println(cb.Add(x, y))
}
