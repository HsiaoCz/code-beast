package main

import (
	"fmt"
	"time"
)

func main() {
	SetTime()
}

// 定时器 指定的时间触发一个事件
func SetTime() {
	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	fmt.Println("hello my man")
}

// 定时触发的计时器
func SetTicker() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			fmt.Println("hello")
		default:
			fmt.Println("hi")
		}

	}
}
