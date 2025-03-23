package Go_base

import "fmt"

// fibonacci_select 函数用于生成斐波那契数列，并通过通道进行通信。
// 该函数使用 select 语句来监听两个通道：c 和 quit。
// 当 c 通道可写时，将当前的斐波那契数列值发送到 c 通道，并更新数列的下一个值。
// 当 quit 通道接收到信号时，函数打印 "quit" 并退出。
// 参数:
//
//	c: 用于发送斐波那契数列值的通道。
//	quit: 用于接收退出信号的通道。
func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	// 无限循环，用于生成斐波那契数列。
	for {
		//select 语句用于监听多个通道的操作，并在有操作时执行相应的分支。
		//select 语句会一直等待，直到某个分支的通道操作完成，然后执行该分支的代码块。
		//如果有多个分支的通道操作同时完成，select 会随机选择一个执行。
		select {
		/// 如果 c 通道可写，则发送当前的斐波那契数列值到 c 通道。
		case c <- x:
			// 发送当前的斐波那契数列值到 c 通道，并更新数列的下一个值。
			x, y = y, x+y
		// 如果 quit 通道接收到信号，则打印 "quit" 并退出函数。
		case <-quit:
			// 接收到退出信号，打印 "quit" 并退出函数。
			fmt.Println("quit")
			return
		}
	}
}

// Hello_select_goroute 函数演示了如何使用 goroutine 和 select 语句来生成斐波那契数列。
// 该函数创建了两个通道 c 和 quit，并启动一个 goroutine 来接收斐波那契数列的值。
// goroutine 会接收 10 个斐波那契数列的值并打印它们，然后发送退出信号到 quit 通道。
// 最后，调用 fibonacci_select 函数来生成斐波那契数列并通过通道进行通信。
// 该函数实现了经典的"生产者-消费者"模式
// 无限循环中使用select阻塞等待事件发生
// 当c通道可写入时，发送当前斐波那契数并计算下一个
// 当quit通道收到消息时，终止函数执行
func Hello_select_goroute() {
	c := make(chan int)
	quit := make(chan int)
	// 启动一个 goroutine 来接收斐波那契数列的值。
	// 启动匿名函数作为消费者协程
	// 消费者打印10个斐波那契数后发送退出信号
	go func() {
		// 接收并打印 10 个斐波那契数列的值，然后发送退出信号到 quit 通道。
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 从通道 c 接收斐波那契数列的值并打印。
		}
		quit <- 0 // 发送退出信号到 quit 通道。
	}()

	//主协程调用fibonacci_select作为生产者
	fibonacci_select(c, quit)
}
