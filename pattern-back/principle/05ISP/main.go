package main

// 接口隔离 强调 不要把所有的方法封装到一个接口上
// 将大的接口设计成小的接口
// 不把所有的操作封装到一个接口上

// 假设A类 需要func1 func2 func5
// 假设B类 需要func1 func3 func4

type SomeISP interface {
	Func1()
	Func2()
	Func3()
	Func4()
	Func5()
}

// 类A 需要的是1 2 5 但是它还需要实现3 4
// 类B 需要的是1 3 4 但是它还需要实现2 5

func main() {

}
