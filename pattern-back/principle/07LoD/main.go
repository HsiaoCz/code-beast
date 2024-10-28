package main

import "fmt"

// 迪米特法则
// 一个对象应该对其他对象尽可能少的了解
// 从而降低各个系统之间的耦合 提高系统的可维护性
// 迪米特法则 又叫最少知道原则
// 如果两个软件实体无须直接通信，那么就不应当发生直接的相互调用，可以通过第三方转发该调用

type Student struct {
	Id   string
	Name string
}

type Class struct {
	Id       string
	Name     string
	Students []Student
}

func (c Class) PrintStudents() {
	for _, student := range c.Students {
		fmt.Println(student)
	}
}

type School struct {
	Id      string
	Name    string
	Classes []Class
}

// 这里 school 与 student 没有直接的关系
func (s School) PrintAllStudents() {

	// 高耦合度的写法
	// for _, class := range s.Classes {
	//  for _, student := range class.Students {
	//   fmt.Println(student)
	//  }
	// }

	// 降低了依赖关系的写法
	for _, class := range s.Classes {
		class.PrintStudents()
	}
}
