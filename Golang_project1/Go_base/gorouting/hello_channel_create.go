package gorouting_test

import (
	"fmt"
	"time"
)

// HelloChannelCreate 展示Go语言中通道(channel)创建的各种详细用法和注意点
func HelloChannelCreate() {
	fmt.Println("===== 1. 通道的基本创建 =====")
	basicChannelCreation()

	fmt.Println("\n===== 2. 有缓冲通道的创建 =====")
	bufferedChannelCreation()

	fmt.Println("\n===== 3. 通道的零值(nil通道) =====")
	nilChannelBehavior()

	fmt.Println("\n===== 4. 单向通道的创建 =====")
	directionalChannelCreation()

	fmt.Println("\n===== 5. 通道的类型转换 =====")
	channelTypeConversion()

	fmt.Println("\n===== 6. 通道的大小和容量 =====")
	channelSizeAndCapacity()

	fmt.Println("\n===== 7. 通道作为函数参数 =====")
	channelAsParameter()

	fmt.Println("\n===== 8. 通道和select的使用 =====")
	channelWithSelect()

	fmt.Println("\n===== 9. 通道创建的最佳实践 =====")
	channelBestPractices()
}

// basicChannelCreation 展示通道的基本创建方法
func basicChannelCreation() {
	// 1. 使用make函数创建通道
	// 语法: make(chan 元素类型)
	// 这会创建一个无缓冲的通道

	// 创建一个传递整数的通道
	intChan := make(chan int)

	// 创建一个传递字符串的通道
	stringChan := make(chan string)

	// 创建一个传递结构体的通道
	type Person struct {
		Name string
		Age  int
	}
	personChan := make(chan Person)

	// 创建一个传递接口的通道
	interfaceChan := make(chan interface{})

	// 启动goroutine使用这些通道
	go func() {
		// 发送数据到各种通道
		intChan <- 42
		stringChan <- "Hello"
		personChan <- Person{"Alice", 30}
		interfaceChan <- "任何类型都可以"
	}()

	// 从通道接收数据
	fmt.Println("整数通道:", <-intChan)
	fmt.Println("字符串通道:", <-stringChan)
	fmt.Println("结构体通道:", <-personChan)
	fmt.Println("接口通道:", <-interfaceChan)

	// 注意事项:
	// 1. 默认创建的是无缓冲通道，发送操作会阻塞直到有接收方准备好
	// 2. 通道的元素类型可以是任何类型，包括内置类型、结构体、接口等
	// 3. 通道本身也是引用类型，类似于切片和映射
	// 4. 无缓冲通道提供了强同步保证，发送和接收是同时发生的
}

// bufferedChannelCreation 展示有缓冲通道的创建
func bufferedChannelCreation() {
	// 有缓冲通道的创建
	// 语法: make(chan 元素类型, 缓冲大小)

	// 创建一个缓冲大小为3的整数通道
	bufferedChan := make(chan int, 3)

	// 发送数据到有缓冲通道
	// 可以连续发送3个值而不阻塞
	fmt.Println("向有缓冲通道发送数据:")
	for i := 1; i <= 3; i++ {
		bufferedChan <- i
		fmt.Printf("发送: %d\n", i)
	}

	// 尝试发送第4个值会阻塞，因为缓冲区已满
	// 启动goroutine来接收，让发送可以继续
	go func() {
		time.Sleep(time.Second) // 等待一秒，让阻塞情况更明显
		for i := 1; i <= 4; i++ {
			value := <-bufferedChan
			fmt.Printf("接收: %d\n", value)
		}
	}()

	fmt.Println("尝试发送第4个值(会阻塞直到有接收)")
	bufferedChan <- 4 // 会阻塞直到上面的goroutine开始接收
	fmt.Println("第4个值已发送")

	// 等待接收goroutine完成
	time.Sleep(2 * time.Second)

	// 注意事项:
	// 1. 缓冲大小必须是非负整数
	// 2. 缓冲大小为0的通道等同于无缓冲通道
	// 3. 有缓冲通道在缓冲区未满时，发送操作不会阻塞
	// 4. 有缓冲通道在缓冲区非空时，接收操作不会阻塞
	// 5. 适当选择缓冲大小很重要:
	//    - 太小: 可能导致不必要的阻塞和性能问题
	//    - 太大: 可能掩盖程序中的并发问题，并消耗更多内存
	// 6. 有缓冲通道可以用作信号量或限制并发数量
}

// nilChannelBehavior 展示nil通道的行为
func nilChannelBehavior() {
	// nil通道: 未初始化的通道变量的零值是nil
	var nilChan chan int // 零值为nil

	fmt.Println("nil通道的行为:")
	fmt.Printf("nilChan == nil: %v\n", nilChan == nil)

	// 以下操作会导致永久阻塞，实际代码中已被注释
	// 1. 从nil通道接收数据会永久阻塞
	// fmt.Println(<-nilChan) // 永久阻塞

	// 2. 向nil通道发送数据会永久阻塞
	// nilChan <- 1 // 永久阻塞

	// 3. 关闭nil通道会导致panic
	// close(nilChan) // panic: close of nil channel

	// nil通道在select语句中的用法
	fmt.Println("nil通道在select中的行为:")
	selectWithNilChannel()

	// 注意事项:
	// 1. 永远不要使用未初始化的通道(除非你有特殊目的)
	// 2. nil通道在select语句中特别有用，相当于禁用该case
	// 3. 确保在使用前正确初始化通道
	// 4. 可以通过将通道设置为nil来禁用特定的通信路径
}

// selectWithNilChannel 展示nil通道在select中的行为
func selectWithNilChannel() {
	var nilChan chan int // nil通道
	activeChan := make(chan int)

	// 在另一个goroutine中发送数据
	go func() {
		time.Sleep(100 * time.Millisecond)
		activeChan <- 42
	}()

	// 在select中，nil通道的case永远不会被选中
	select {
	case v := <-nilChan:
		// 这个case永远不会被选中
		fmt.Println("从nil通道接收:", v)
	case v := <-activeChan:
		fmt.Println("从活动通道接收:", v)
	}
}

// directionalChannelCreation 展示单向通道的创建
func directionalChannelCreation() {
	// 创建双向通道
	bidirectional := make(chan int)

	// 单向发送通道
	var sendOnly chan<- int = bidirectional

	// 单向接收通道
	var recvOnly <-chan int = bidirectional

	// 使用这些通道
	go func() {
		// 向发送通道发送数据
		sendOnly <- 42

		// 尝试从发送通道接收数据会导致编译错误
		// value := <-sendOnly // 编译错误
	}()

	// 从接收通道接收数据
	value := <-recvOnly
	fmt.Println("从接收通道接收:", value)

	// 尝试向接收通道发送数据会导致编译错误
	// recvOnly <- 100 // 编译错误

	// 注意事项:
	// 1. 单向通道主要用于提高类型安全性
	// 2. 可以将双向通道转换为单向通道，但反过来不行
	// 3. 单向通道在函数参数和返回值中特别有用
	// 4. 单向发送通道: chan<- T
	// 5. 单向接收通道: <-chan T
	// 6. 通道转换是单向的: 双向 -> 单向，而非 单向 -> 双向
}

// channelTypeConversion 展示通道的类型转换
func channelTypeConversion() {
	// 创建各种类型的通道
	intChan := make(chan int)

	// 通道类型转换示例
	// 1. 双向 -> 单向发送
	sendChan := chan<- int(intChan)

	// 2. 双向 -> 单向接收
	recvChan := <-chan int(intChan)

	// 使用转换后的通道
	go func() {
		sendChan <- 100
		// <-sendChan // 编译错误: 不能从单向发送通道接收
	}()

	// fmt.Println("从接收通道接收:", <-recvChan) // recvChan <- 200 // 编译错误: 不能向单向接收通道发送

	fmt.Println("转换后的通道:", sendChan, recvChan)

	// 不支持的转换示例（编译错误）
	// 不能将单向通道转换回双向通道
	// biChan := chan int(sendChan) // 编译错误
	// biChan := chan int(recvChan) // 编译错误

	// 不能将通道转换为不同元素类型的通道
	// stringChan := chan string(intChan) // 编译错误

	// 注意事项:
	// 1. 通道转换只允许从双向通道到单向通道
	// 2. 不允许从单向通道到双向通道的转换
	// 3. 不允许改变通道的元素类型
	// 4. 类型转换是静态的，发生在编译时
	// 5. 转换后的通道引用相同的底层通道
}

// channelSizeAndCapacity 展示通道的大小和容量
func channelSizeAndCapacity() {
	// 创建一个缓冲大小为5的通道
	ch := make(chan int, 5)

	// 检查初始状态
	fmt.Printf("初始: 长度=%d, 容量=%d\n", len(ch), cap(ch))

	// 发送一些值
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("发送后: 长度=%d, 容量=%d\n", len(ch), cap(ch))
	}

	// 接收一些值
	<-ch
	<-ch
	fmt.Printf("接收后: 长度=%d, 容量=%d\n", len(ch), cap(ch))

	// 注意事项:
	// 1. len(channel): 返回通道缓冲区中当前的元素数量
	// 2. cap(channel): 返回通道缓冲区的容量
	// 3. 对无缓冲通道，cap()返回0
	// 4. len()总是小于或等于cap()
	// 5. 可以在不阻塞的情况下发送len() < cap()个元素
	// 6. 编译时无法知道通道的长度，这是运行时属性
}

// channelAsParameter 展示通道作为函数参数的用法
func channelAsParameter() {
	// 创建双向通道
	ch := make(chan string)

	// 启动生产者和消费者
	go producer(ch)
	go consumer(ch)

	// 等待通信完成
	time.Sleep(500 * time.Millisecond)

	// 注意事项:
	// 1. 通道作为函数参数时通常使用单向类型，增加类型安全性
	// 2. 通道是引用类型，函数接收的是原始通道的引用
	// 3. 可以通过参数指定通道的方向，限制函数的操作权限
	// 4. 通常，生产者接收发送通道，消费者接收接收通道
	// 5. 关闭通道的责任应该在生产者一方
}

// producer 函数只需要向通道发送数据
func producer(ch chan<- string) {
	// chan<- 表示只能发送的通道
	ch <- "Message 1"
	ch <- "Message 2"
	close(ch) // 生产者负责关闭通道
}

// consumer 函数只需要从通道接收数据
func consumer(ch <-chan string) {
	// <-chan 表示只能接收的通道
	for msg := range ch {
		fmt.Println("消费:", msg)
	}
}

// channelWithSelect 展示通道创建与select语句的结合使用
func channelWithSelect() {
	// 创建多个通道
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	// 启动多个发送者
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自通道1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "来自通道2"
	}()

	// 启动结束信号
	go func() {
		time.Sleep(200 * time.Millisecond)
		done <- true
	}()

	// 使用select处理多个通道
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("收到ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("收到ch2:", msg2)
		case <-done:
			fmt.Println("收到结束信号")
			return
		default:
			// 非阻塞操作
			fmt.Println("没有通道就绪，继续...")
			time.Sleep(10 * time.Millisecond)
		}
	}

	// 注意事项:
	// 1. select使得处理多个通道变得容易
	// 2. 如果多个case就绪，select会随机选择一个
	// 3. default使select变为非阻塞操作
	// 4. select{}（没有case的select）会永远阻塞
	// 5. 通常使用done通道来发送终止信号
}

// channelBestPractices 展示通道创建的最佳实践
func channelBestPractices() {
	fmt.Println("通道创建的最佳实践:")

	// 1. 适当的缓冲区大小
	fmt.Println("1. 适当的缓冲区大小:")
	fmt.Println("   - 无缓冲通道: 用于需要同步的场景")
	fmt.Println("   - 有缓冲通道: 根据预期的生产消费速率设置缓冲区大小")

	// 2. 通道所有权
	fmt.Println("2. 通道所有权原则:")
	fmt.Println("   - 创建者通常拥有通道")
	fmt.Println("   - 拥有者负责关闭通道")
	fmt.Println("   - 尽量不要从接收方关闭通道")

	// 3. 单向通道的使用
	fmt.Println("3. 使用单向通道增加类型安全性:")
	fmt.Println("   - 函数参数尽量使用单向通道类型")
	fmt.Println("   - 限制功能只执行必要的操作")

	// 4. 通道的关闭和检查
	fmt.Println("4. 通道的关闭和检查:")
	fmt.Println("   - 不要关闭已关闭的通道")
	fmt.Println("   - 使用comma-ok语法检查通道是否已关闭")
	fmt.Println("   - 使用defer确保通道在适当时机关闭")

	// 5. 避免通道泄漏
	fmt.Println("5. 避免通道泄漏:")
	fmt.Println("   - 确保每个生产者和消费者goroutine都能正确退出")
	fmt.Println("   - 使用done通道发送终止信号")
	fmt.Println("   - 考虑使用context包管理goroutine生命周期")

	// 6. 性能考虑
	fmt.Println("6. 性能考虑:")
	fmt.Println("   - 通道操作有开销，不要用于细粒度的并发控制")
	fmt.Println("   - 避免频繁创建和销毁通道")
	fmt.Println("   - 对于高性能场景，考虑预先分配足够的通道")

	// 7. 错误处理
	fmt.Println("7. 错误处理:")
	fmt.Println("   - 可以使用带错误信息的结构体通过通道传递")
	fmt.Println("   - 考虑使用专门的错误通道")
}
