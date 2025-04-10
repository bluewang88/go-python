package gorouting_test

import (
	"fmt"
	"sync"
	"time"
)

// HelloChannel 展示Go语言中通道(channel)的各种用法和注意事项
func HelloChannel() {
	fmt.Println("===== 1. 基本通道用法 =====")
	basicChannelUsage()

	fmt.Println("\n===== 2. 无缓冲通道 vs 有缓冲通道 =====")
	bufferedVsUnbuffered()

	fmt.Println("\n===== 3. 通道的方向 =====")
	channelDirection()

	fmt.Println("\n===== 4. select语句与多通道操作 =====")
	selectStatement()

	fmt.Println("\n===== 5. 通道的关闭与遍历 =====")
	closeAndRange()

	fmt.Println("\n===== 6. 超时处理 =====")
	timeoutPattern()

	fmt.Println("\n===== 7. 工作池模式 =====")
	workerPool()

	fmt.Println("\n===== 8. 通道的常见错误和陷阱 =====")
	commonPitfalls()
}

// basicChannelUsage 展示通道的基本用法
func basicChannelUsage() {
	// 创建一个整型通道
	ch := make(chan int)

	// 启动一个goroutine从通道接收数据
	go func() {
		// 从通道接收值
		value := <-ch
		fmt.Println("接收到的值:", value)
	}()

	// 向通道发送数据
	fmt.Println("发送值: 42")
	ch <- 42

	// 给接收goroutine一点时间处理
	time.Sleep(100 * time.Millisecond)

	// 注意事项:
	// 1. 通道操作是同步的：发送操作会阻塞直到有人接收
	// 2. 默认通道是无缓冲的，发送和接收必须同时准备好
	// 3. 向nil通道发送或从nil通道接收将永远阻塞
	// 4. 从已关闭的通道接收会立即返回零值
	// 5. 向已关闭的通道发送会导致panic
}

// bufferedVsUnbuffered 展示有缓冲通道和无缓冲通道的区别
func bufferedVsUnbuffered() {
	// 无缓冲通道：发送方必须等待接收方准备好
	unbuffered := make(chan int)

	// 有缓冲通道：可以发送指定数量的值而不需要接收方准备好
	buffered := make(chan int, 3) // 缓冲大小为3

	// 无缓冲通道演示
	fmt.Println("无缓冲通道演示:")
	go func() {
		fmt.Println("接收者准备从无缓冲通道接收")
		value := <-unbuffered
		fmt.Println("从无缓冲通道接收到:", value)
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("发送者发送到无缓冲通道")
	unbuffered <- 10
	time.Sleep(100 * time.Millisecond)

	// 有缓冲通道演示
	fmt.Println("\n有缓冲通道演示:")
	// 发送3个值，不会阻塞因为有足够的缓冲空间
	fmt.Println("发送3个值到有缓冲通道")
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("发送完成")

	// 尝试发送第4个值会阻塞，直到有空间
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("开始从有缓冲通道接收")
		fmt.Println("接收:", <-buffered)
		fmt.Println("接收:", <-buffered)
		fmt.Println("接收:", <-buffered)
		fmt.Println("接收:", <-buffered) // 接收第4个值
	}()

	fmt.Println("尝试发送第4个值(会阻塞直到有空间)")
	buffered <- 4
	fmt.Println("第4个值发送成功")

	// 注意事项:
	// 1. 无缓冲通道提供强同步保证：发送和接收是同时发生的
	// 2. 有缓冲通道允许发送一定数量的值而不阻塞
	// 3. 当缓冲区已满时，发送会阻塞；当缓冲区为空时，接收会阻塞
	// 4. 缓冲大小应基于预期负载和性能要求来选择
	// 5. 使用过大的缓冲区可能掩盖并发问题
}

// channelDirection 展示通道的方向（单向通道）
func channelDirection() {
	// 创建一个双向通道
	ch := make(chan int)

	// 启动发送者goroutine
	go sender(ch)

	// 启动接收者goroutine
	go receiver(ch)

	// 等待完成
	time.Sleep(100 * time.Millisecond)

	// 注意事项:
	// 1. 单向通道增加了类型安全性
	// 2. 可以将双向通道转换为单向通道，但反之不行
	// 3. 单向通道在函数签名中特别有用，明确表达意图
	// 4. 即使通道是单向的，也可以关闭它(但只能由发送者关闭)
}

// sender 只能向通道发送数据
func sender(ch chan<- int) {
	// chan<- 表示只能发送的通道
	ch <- 42
	// <-ch // 编译错误：不能从只发送通道接收
}

// receiver 只能从通道接收数据
func receiver(ch <-chan int) {
	// <-chan 表示只能接收的通道
	value := <-ch
	fmt.Println("接收者收到:", value)
	// ch <- 100 // 编译错误：不能向只接收通道发送
}

// selectStatement 展示select语句的用法
func selectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个发送者
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "来自通道1"
	}()

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch2 <- "来自通道2"
	}()

	// 使用select同时等待多个通道
	fmt.Println("等待来自任一通道的消息...")

	// 第一次select
	select {
	case msg1 := <-ch1:
		fmt.Println("收到:", msg1)
	case msg2 := <-ch2:
		fmt.Println("收到:", msg2)
	}

	// 再等待一会儿，让另一个通道也能发送
	time.Sleep(50 * time.Millisecond)

	// 第二次select，带默认分支
	select {
	case msg1 := <-ch1:
		fmt.Println("收到:", msg1)
	case msg2 := <-ch2:
		fmt.Println("收到:", msg2)
	default:
		fmt.Println("没有通道准备好")
	}

	// 注意事项:
	// 1. select语句在多个通道操作中非常有用
	// 2. 如果多个case同时就绪，select会随机选择一个
	// 3. default分支在没有通道就绪时执行，使select变为非阻塞
	// 4. 空select{}会永远阻塞
	// 5. select可以用于实现超时、取消和优先级等模式
}

// closeAndRange 展示通道的关闭和遍历
func closeAndRange() {
	ch := make(chan int, 5)

	// 向通道发送一系列值
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// 关闭通道表示不会再有值发送
	close(ch)

	// 尝试发送到已关闭的通道会导致panic
	// ch <- 100 // 会panic

	// 从已关闭的通道接收:
	// 1. 首先获取所有已发送的值
	// 2. 然后返回通道类型的零值
	// 3. 第二个返回值表示通道是否已关闭

	fmt.Println("从关闭的通道接收值:")

	// 方法1: 显式检查通道是否关闭
	for {
		value, ok := <-ch
		if !ok {
			// 通道已关闭且为空
			fmt.Println("通道已关闭")
			break
		}
		fmt.Println("接收:", value)
	}

	// 重新创建并填充通道
	ch = make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i * 10
	}
	close(ch)

	// 方法2: 使用range遍历通道直到关闭
	fmt.Println("\n使用range遍历通道:")
	for value := range ch {
		fmt.Println("接收:", value)
	}

	// 注意事项:
	// 1. 只有发送者应该关闭通道，接收者不应关闭
	// 2. 关闭已关闭的通道会导致panic
	// 3. 发送到已关闭的通道会导致panic
	// 4. 接收者可以使用第二个返回值检查通道是否关闭
	// 5. range会自动处理通道关闭的情况
	// 6. 通道关闭是一种广播机制：所有等待接收的goroutine都会收到通知
}

// timeoutPattern 展示如何使用select和time.After实现超时
func timeoutPattern() {
	ch := make(chan string)

	// 模拟可能需要长时间运行的操作
	go func() {
		// 延迟5秒，模拟长时间操作
		time.Sleep(2 * time.Second)
		ch <- "操作完成"
	}()

	// 等待结果，但最多等待1秒
	fmt.Println("等待结果，超时时间1秒...")
	select {
	case result := <-ch:
		fmt.Println("收到结果:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("操作超时")
	}

	// 再等待足够长的时间，让首次模拟操作完成
	time.Sleep(2 * time.Second)

	// 再次尝试，使用更长的超时时间
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "第二次操作完成"
	}()

	fmt.Println("\n等待结果，超时时间3秒...")
	select {
	case result := <-ch:
		fmt.Println("收到结果:", result)
	case <-time.After(3 * time.Second):
		fmt.Println("操作超时")
	}

	// 注意事项:
	// 1. time.After返回一个在指定时间后发送当前时间的通道
	// 2. 可以与select结合使用，优雅地处理超时
	// 3. 超时模式对于避免程序无限等待非常有用
	// 4. 在生产环境中，应考虑使用context包进行超时和取消控制
	// 5. 对于多次使用的情况，应该使用time.Timer，它更高效
}

// workerPool 展示工作池模式的实现
func workerPool() {
	// 任务通道
	jobs := make(chan int, 100)
	// 结果通道
	results := make(chan int, 100)

	// 启动3个工作goroutine
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 发送9个任务
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) // 关闭任务通道，表示没有更多任务

	// 启动一个goroutine等待所有工作者完成并关闭结果通道
	go func() {
		wg.Wait()
		close(results)
	}()

	// 从结果通道收集所有结果
	fmt.Println("收集结果:")
	for result := range results {
		fmt.Println("结果:", result)
	}

	// 注意事项:
	// 1. 工作池模式适用于并行处理独立任务
	// 2. 任务通道用于分发工作，结果通道用于收集结果
	// 3. 通过关闭任务通道告知工作者没有更多工作
	// 4. 使用WaitGroup等待所有工作者完成
	// 5. 必须在所有发送者完成后才能关闭通道
	// 6. 通道的关闭是广播机制，所有工作者都会收到通知
}

// worker 是工作池中的工作函数
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Printf("工作者 %d 开始处理任务 %d\n", id, j)
		// 模拟处理任务
		time.Sleep(100 * time.Millisecond)
		// 发送结果
		results <- j * 2
		fmt.Printf("工作者 %d 完成任务 %d\n", id, j)
	}
}

// commonPitfalls 展示通道使用中的常见错误和陷阱
func commonPitfalls() {
	fmt.Println("通道使用中的常见错误和陷阱:")

	// 1. 死锁示例(注释掉以避免程序停止)
	fmt.Println("1. 死锁 - 向无缓冲通道发送而没有接收者")
	// deadlockExample()

	// 2. 向已关闭的通道发送
	fmt.Println("\n2. 向已关闭的通道发送(会导致panic)")
	// sendToClosedChannel()

	// 3. 关闭已关闭的通道
	fmt.Println("\n3. 关闭已关闭的通道(会导致panic)")
	// closeClosedChannel()

	// 4. nil通道
	fmt.Println("\n4. nil通道操作会永久阻塞")
	// nilChannelExample()

	// 5. 泄漏的goroutine
	fmt.Println("\n5. goroutine泄漏")
	// 在实际代码中goroutine泄漏会导致内存泄漏

	fmt.Println("\n避免这些问题的建议:")
	fmt.Println("- 确保每个发送操作都有匹配的接收操作")
	fmt.Println("- 使用defer确保通道在适当的时候关闭")
	fmt.Println("- 明确通道的所有权：通常只有发送者应该关闭通道")
	fmt.Println("- 使用带超时的select模式避免永久阻塞")
	fmt.Println("- 考虑使用有缓冲通道减少阻塞风险")
	fmt.Println("- 使用context包进行超时和取消控制")
}

// 以下是展示常见错误的函数(注释掉实际调用以避免程序崩溃)

// deadlockExample 展示死锁情况
func deadlockExample() {
	ch := make(chan int) // 无缓冲通道

	// 发送操作会阻塞，因为没有接收者
	// 这会导致死锁
	// ch <- 1

	// 正确的方式是确保有接收者
	go func() {
		value := <-ch
		fmt.Println("接收到:", value)
	}()

	ch <- 1 // 现在安全了，因为有goroutine在接收
}

// sendToClosedChannel 展示向已关闭的通道发送的错误
func sendToClosedChannel() {
	ch := make(chan int)
	close(ch)

	// 这会导致panic
	// ch <- 1
}

// closeClosedChannel 展示关闭已关闭通道的错误
func closeClosedChannel() {
	ch := make(chan int)
	close(ch)

	// 这会导致panic
	// close(ch)
}

// nilChannelExample 展示nil通道的阻塞行为
func nilChannelExample() {
	var ch chan int // 默认为nil

	// 以下操作会永久阻塞
	// <-ch      // 从nil通道接收会阻塞
	// ch <- 1   // 向nil通道发送会阻塞

	// close(ch) // 关闭nil通道会导致panic
	// 注意：nil通道在select语句中不会阻塞
	// select {
	// case <-ch: // 永久阻塞
	// 	fmt.Println("从nil通道接收")
	// }
	// case ch <- 1: // 永久阻塞
	// 	fmt.Println("向nil通道发送")
	// }

	fmt.Println(ch)
}
