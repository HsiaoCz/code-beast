package main

import "fmt"

// 开闭原则
// 对扩展开放
// 对修改关闭

// 平铺式的代码
// type Banker struct{}

// func (b *Banker) Save() { fmt.Println("银行职员进行了存款的业务...") }

// func (b *Banker) Trans() { fmt.Println("银行职员进行了转账的业务....") }

// func (b *Banker) Stack() { fmt.Println("银行职员进行了股票的业务....") }

// 这种平铺的设计  在增加职员的作用的时候 需要修改原有的类
// 这中修改是危险的

// 基于开闭原则的代码设计

type Banker interface {
	DoBuz()
}

type SaveBanker struct{}

func (s *SaveBanker) DoBuz() { fmt.Println("银行职员进行了存款的业务....") }

type TransBanker struct{}

func (t *TransBanker) DoBuz() { fmt.Println("银行职员进行了转账的业务....") }

type StackBanker struct{}

func (s *StackBanker) DoBuz() { fmt.Println("银行职员进行了股票的业务....") }
