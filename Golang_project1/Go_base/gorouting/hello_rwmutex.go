package gorouting_test

import (
	"fmt"
	"math/rand" // 使用math/rand而不是crypto/rand
	"sync"
	"time"
)

//读写锁是对共享资源的读写访问和操作，全称是Read Write Lock。
//读写锁允许多个读操作并发进行，但写操作是互斥的。
//读写锁的使用步骤：
// 1. 创建一个sync.RWMutex实例
// 2. 在读操作之前调用RLock()方法获取读锁
// 3. 在读操作完成后调用RUnlock()方法释放读锁
// 4. 在写操作之前调用Lock()方法获取写锁
// 5. 在写操作完成后调用Unlock()方法释放写锁
// 6. 使用defer语句确保在函数退出时释放锁
// 7. 注意：在同一个goroutine中，不能重复调用RLock()或Lock()，否则会导致死锁

func ReadMutex(mu *sync.RWMutex, id int) {
	mu.RLock() // 获取读锁
	defer mu.RUnlock()

	// 生成随机等待时间（1-3秒）
	waitTime := time.Duration(rand.Intn(3)+1) * time.Second

	// 记录开始时间
	start := time.Now()

	// 等待随机时间，模拟读操作
	time.Sleep(waitTime)

	// 计算实际等待时间
	elapsed := time.Since(start)

	fmt.Printf("Reader %d: 读取完成，花费时间: %v\n", id, elapsed)
}

func WriteMutex(mu *sync.RWMutex, id int) {
	mu.Lock() // 获取写锁
	defer mu.Unlock()

	// 记录开始时间
	start := time.Now()

	fmt.Printf("Writer %d: 开始写入\n", id)
	// 模拟写操作（固定为10秒）
	time.Sleep(10 * time.Second)

	// 计算实际等待时间
	elapsed := time.Since(start)

	fmt.Printf("Writer %d: 写入完成，花费时间: %v\n", id, elapsed)
}

func HelloRWMutex() {
	// 初始化随机数生成器
	// rand.Seed(time.Now().UnixNano())

	// 创建一个读写锁
	var mu sync.RWMutex

	// 创建WaitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	fmt.Println("=== 开始读写锁示例 ===")

	// 启动多个goroutine进行并发操作
	// 使用较少的goroutine使输出更清晰
	readersCount := 10
	writersCount := 2

	// 启动读取goroutine
	for i := 0; i < readersCount; i++ {
		wg.Add(1) // 添加到WaitGroup
		// 启动读取goroutine
		go func(id int) {
			defer wg.Done() // 在每个goroutine完成时调用Done()
			ReadMutex(&mu, id)
		}(i)
	}

	// 启动写入goroutine
	for i := 0; i < writersCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			WriteMutex(&mu, id)
		}(i)
	}

	// 主线程等待所有goroutine完成
	fmt.Println("等待所有读写操作完成...")
	wg.Wait()

	fmt.Println("=== 读写锁示例完成 ===")
}
