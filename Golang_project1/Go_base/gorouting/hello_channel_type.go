package gorouting_test

import "fmt"

/*
channel提供了一种协程之间的通信机制，（java、python鸥鸟线程之间的通信机制是消息队列）
channel 可以存放几乎任何数据类型，包括基本类型、复合类型、自定义类型、接口类型、函数类型，甚至是其他 channel 类型
channel 是一种引用类型，类似于指针，使用时需要使用 make 函数来创建
Go语言中的引用类型包括切片、映射、通道、函数、接口、指针和结构体指针。这些类型在内存中存储的是对数据的引用，而不是数据本身，因此具有共享内存和高效传递的特点。
channel的默认值是 nil，表示没有初始化的通道。nil 通道不能用于发送或接收数据，尝试这样做会导致运行时错误。
*/

func HelloChannelType() {

	//消息队列或者说是channel在声明时需要指明通道可以存放的数据类型

	//定义channel
	// var msg chan int

	// //初始化channel，无缓冲
	// msg = make(chan int)

	//初始化channel，有缓冲

	msg_cache := make(chan int, 1) // 缓冲区大小为10

	// go func() {
	// 	msg <- 42 // 向通道msg发送数据42
	// }()
	// msg <- 42      // 向通道msg发送数据42,会报错fatal error: all goroutines are asleep - deadlock!
	// value := <-msg // 从通道msg接收数据
	// fmt.Println(value)

	msg_cache <- 42      // 向通道msg_cache发送数据42
	value := <-msg_cache // 从通道msg_cache接收数据
	fmt.Println(value)   // 打印接收到的数据

}
