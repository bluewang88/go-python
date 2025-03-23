package Go_base

import "fmt"

// 信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
// 信道是 Go 语言中用于 Goroutine 之间通信的核心机制，类似于其他语言中的消息队列或管道。
// 信道的类型决定了可以发送和接收的数据类型，例如 `chan int` 表示只能传递 `int` 类型的数据。

func Hello_buffered_channels() {
	// 创建一个带缓冲区的信道，大小为 2。
	messages := make(chan string, 2)

	// 发送两个消息到信道中。
	messages <- "buffered"
	messages <- "channel"

	// 从信道中接收两个消息。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
