package gorouting_test

// 互斥锁是对共享资源的互斥访问和操作，全称是Mutual Exclusion。
// 使用互斥锁的步骤：
// 1. 创建一个sync.Mutex实例
// 2. 在访问共享资源之前调用Lock()方法获取锁
// 3. 在访问完成后调用Unlock()方法释放锁
// 4. 使用defer语句确保在函数退出时释放锁
// 5. 注意：在同一个goroutine中，不能重复调用Lock()，否则会导致死锁
// 6. 在不同的goroutine中，可以使用同一个锁来保护共享资源

import (
	"sync"
)

func add(a int) int {
	return a + 1
}
func sub(a int) int {
	return a - 1
}

func HelloMutex1() {
	// 创建一个互斥锁
	var mu sync.Mutex

	// 创建一个WaitGroup实例
	var wg sync.WaitGroup

	// 定义共享变量
	counter := 0

	// 启动多个goroutine进行并发操作
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()   // 在每个goroutine完成时调用Done()
			defer mu.Unlock() // 使用defer确保解锁
			mu.Lock()         // 获取锁
			counter = add(1)  // 访问共享资源
			// mu.Unlock()      // 释放锁
		}()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		// 启动多个goroutine进行并发操作
		// 这里使用了defer来确保在函数退出时释放锁
		// 这里使用了匿名函数来传递参数
		go func() {
			defer wg.Done()   // 在每个goroutine完成时调用Done()
			defer mu.Unlock() // 使用defer确保解锁
			mu.Lock()         // 获取锁
			counter = sub(1)  // 访问共享资源
			// mu.Unlock()      // 释放锁

		}()
	}
	// 等待所有goroutine完成
	wg.Wait() // 等待所有的goroutine完成
	// 这里可以使用sync.WaitGroup来等待所有goroutine完成
	// 也可以使用time.Sleep来等待一段时间
	// 这里使用time.Sleep来等待一段时间
	// time.Sleep(2 * time.Second)
	// 打印最终计数
	println("最终计数:", counter)
	// 注意事项:
	// 1. 总是成对使用Lock和Unlock
	// 2. 确保在任何返回路径上都会调用Unlock，通常使用defer来确保
	// 3. 锁定后尽快解锁，减少持有锁的时间
	// 4. 避免在持有锁的情况下进行长时间的阻塞操作
	// 5. 避免在锁定的情况下调用其他可能会获取同一锁的函数
	// 6. 避免死锁：确保锁的获取顺序一致，避免循环依赖
	// 7. 使用读写锁（sync.RWMutex）来优化读操作频繁的场景
	// 8. 使用sync.Once来确保某个操作只执行一次
	// 9. 使用sync.Cond来实现条件变量，用于在某些条件下通知等待的goroutine
	// 10. 使用sync.Pool来实现对象池，减少内存分配和垃圾回收的开销
	// 11. 使用context.Context来实现超时和取消操作
	// 12. 使用channel来实现goroutine之间的通信和同步
	// 13. 使用atomic包来实现原子操作，避免使用锁
}
