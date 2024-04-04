package main

import "fmt"

func main() {

	var i *int

	var p interface{}

	p = i

	//i, ok := p.(*int)
	//if ok {
	//	fmt.Println(i)
	//}

	//if i == nil {
	//	fmt.Println("i is nill")
	//}

	if p != nil {
		fmt.Println("p is not nill")
		fmt.Println(p)
	}
}
