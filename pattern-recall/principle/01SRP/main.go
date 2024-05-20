package main

// 不符合单一职责原则的接口
type PhoneCall interface {
	Dial(phoneNumber int64)
	Talk(somebody struct{})
	Hangup()
}

// 可以拆分成两个接口
// 一个接口负责协议管理
type IConnectionMgr interface {
	Dial(phoneNumber int64)
	Hangup()
}

// 一个接口负责数据传输
type DataTransferMgr interface {
	Talk(somebody struct{})
}

func main() {
}
