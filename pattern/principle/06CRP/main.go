package main

import "fmt"

// 合成复用原则

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 给小猫添加睡觉的方法
// 通过继承的方式

type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 给小猫添加睡觉的方法
// 通过组合的方式

type CatC struct {
	Ca Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c := Cat{}
	c.Eat()

	cb := CatB{c}
	cb.Sleep()
	cb.Eat() //继承

	// 组合
	cc := CatC{Ca: c}
	cc.Sleep()
	cc.Ca.Eat() // 组合
}
