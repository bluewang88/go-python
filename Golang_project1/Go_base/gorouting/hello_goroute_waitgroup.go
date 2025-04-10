package gorouting_test

//如何解决主协程在子协程执行完之前退出的问题？
// 使用sync.WaitGroup来等待所有的goroutine完成
// WaitGroup是一个计数器，用于等待一组goroutine完成
import (
	"fmt"
	"sync"
)

// WaitGroup的使用步骤：
// 1. 创建一个WaitGroup实例
// 2. 在每个goroutine开始之前调用Add(1)
// 3. 在每个goroutine完成时调用Done()
// 4. 在主线程中调用Wait()，等待所有的goroutine完成
func HelloGoroutineWaitGroup() {
	var wg sync.WaitGroup // 创建一个WaitGroup实例

	for i := 0; i < 100; i++ {
		wg.Add(1) // 在每个goroutine开始之前调用Add(1)
		go func(n int) {
			defer wg.Done() // 在每个goroutine完成时调用Done()
			fmt.Println("Hello, goroutine!", n)
		}(i)
	}

	// 主线程继续执行
	fmt.Println("Hello, main!")

	wg.Wait() // 等待所有的goroutine完成
}
