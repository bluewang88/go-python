package Go_base

import (
	"fmt"
)

// Hello_range_and_close 函数演示了如何使用带缓冲区的信道，并通过 range 遍历信道中的值。
// 该函数首先创建一个大小为 2 的带缓冲区信道，并向其中发送两个整数值。
// 然后关闭信道，并使用 range 遍历信道中的值，最后打印出“信道已关闭”的提示信息。
func Hello_range_and_close() {
	ch := make(chan int, 2) // 创建一个带缓冲区的信道，大小为 2。
	ch <- 1
	ch <- 2
	close(ch)

	// 使用 range 遍历信道中的值，直到信道关闭
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("信道已关闭")
}

// fibonacci 函数生成斐波那契数列，并将结果发送到信道中。
// 参数 n 表示生成的斐波那契数列的长度，c 是用于接收斐波那契数列的信道。
// 该函数在生成完所有数列后关闭信道。
// 斐波那契数列（Fibonacci Sequence）是一个经典的数学序列，定义如下：

// 定义：

// 数列的前两项为 0 和 1。
// 从第三项开始，每一项都等于前两项之和。
// 数学表示：

// ( F(0) = 0 )
// ( F(1) = 1 )
// ( F(n) = F(n-1) + F(n-2) ) （当 ( n \geq 2 ) 时）
// 数列示例：

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Hello_range_and_close_demo 函数演示了如何使用 goroutine 和信道生成斐波那契数列。
// 该函数创建一个大小为 10 的带缓冲区信道，并启动一个 goroutine 来生成斐波那契数列。
// 然后使用 range 遍历信道中的值，并打印出每个斐波那契数。
func Hello_range_and_close_demo() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
