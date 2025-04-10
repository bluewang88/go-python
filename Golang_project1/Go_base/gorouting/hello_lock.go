package gorouting_test

import (
	"fmt"
	"sync"
	"time"
)

/*
锁 - 并发编程下的资源竞争
互斥锁是用于保护共享资源的基本工具
在多线程环境中，多个goroutine可能会同时访问共享资源，导致数据不一致或竞争条件。
互斥锁（sync.Mutex）用于确保同一时间只有一个goroutine可以访问共享资源。
使用互斥锁的基本步骤：
1. 创建一个sync.Mutex实例
2. 在访问共享资源之前调用Lock()方法获取锁
3. 在访问完成后调用Unlock()方法释放锁
4. 使用defer语句确保在函数返回时释放锁
注意事项：
1. 确保Lock和Unlock成对使用
2. 避免在持有锁的情况下进行长时间的操作
3. 避免死锁：确保所有goroutine以相同的顺序获取锁
4. 使用sync.RWMutex实现读写锁，允许多个读操作并发进行，但写操作是互斥的
5. 使用sync.Once确保某个操作只执行一次，常用于单例模式或一次性初始化
6. 使用sync.WaitGroup等待一组goroutine完成
7. 使用sync.Cond实现条件变量，用于在某些条件下通知等待的goroutine
8. 使用sync.Pool实现对象池，减少内存分配和垃圾回收的开销
9. 使用context.Context实现超时和取消操作
10. 使用channel实现goroutine之间的通信和同步
11. 使用atomic包实现原子操作，避免使用锁
*/
// HelloMutex 展示互斥锁同步goroutine的各种用法和注意事项
func HelloMutex() {
	fmt.Println("===== 1. 互斥锁基本用法 =====")
	basicMutexUsage()

	fmt.Println("\n===== 2. 保护共享数据结构 =====")
	protectSharedData()

	fmt.Println("\n===== 3. 读写锁RWMutex =====")
	rwMutexDemo()

	fmt.Println("\n===== 4. 锁的粒度选择 =====")
	lockGranularity()

	fmt.Println("\n===== 5. 避免死锁 =====")
	avoidDeadlock()

	fmt.Println("\n===== 6. sync.Once示例 =====")
	onceDemo()
}

// basicMutexUsage 展示互斥锁的基本用法
func basicMutexUsage() {
	// 定义一个互斥锁
	var mu sync.Mutex

	// 定义共享变量
	counter := 0

	// 创建WaitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 模拟一些工作
			time.Sleep(time.Millisecond * 10)

			// 互斥锁保护临界区
			mu.Lock()
			// 以下是临界区(critical section)，同一时间只有一个goroutine可以执行
			counter++
			fmt.Printf("Goroutine %d: counter = %d\n", id, counter)
			// 解锁
			mu.Unlock()
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Printf("最终计数: %d\n", counter)

	// 注意事项:
	// 1. 总是成对使用Lock和Unlock
	// 2. 确保在任何返回路径上都会调用Unlock，通常使用defer来确保
	// 3. 锁定后尽快解锁，减少持有锁的时间
}

// protectSharedData 展示如何使用互斥锁保护复杂的共享数据结构
func protectSharedData() {
	// 定义一个线程安全的计数器类型
	type SafeCounter struct {
		mu     sync.Mutex
		counts map[string]int
	}

	// 创建计数器实例
	counter := SafeCounter{
		counts: make(map[string]int),
	}

	// 定义增加计数的方法
	increment := func(key string) {
		counter.mu.Lock()
		defer counter.mu.Unlock() // 使用defer确保解锁

		counter.counts[key]++
	}

	// 定义获取计数的方法
	getValue := func(key string) int {
		counter.mu.Lock()
		defer counter.mu.Unlock()

		return counter.counts[key]
	}

	// 启动多个goroutine并发增加计数
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment("someKey")
		}()
	}

	wg.Wait()
	fmt.Printf("someKey的计数: %d\n", getValue("someKey"))

	// 注意事项:
	// 1. 将互斥锁与其保护的数据封装在同一个结构中是好的实践
	// 2. 对外提供方法而不是直接暴露数据和锁
	// 3. 避免锁的传递和暴露，这会导致很难追踪锁的状态
	// 4. 如果结构体包含多个独立的数据域，考虑使用多个锁以增加并发性
}

// rwMutexDemo 展示读写锁(RWMutex)的用法
func rwMutexDemo() {
	// 读写锁允许多个读操作并发进行，但写操作是互斥的
	var rwMu sync.RWMutex

	// 共享数据
	data := make(map[string]string)

	// 写入函数
	write := func(key, value string) {
		rwMu.Lock() // 写锁定
		defer rwMu.Unlock()

		// 模拟耗时写操作
		time.Sleep(time.Millisecond * 10)
		data[key] = value
		fmt.Printf("写入: %s = %s\n", key, value)
	}

	// 读取函数
	read := func(key string) string {
		rwMu.RLock() // 读锁定
		defer rwMu.RUnlock()

		// 模拟耗时读操作
		time.Sleep(time.Millisecond * 5)
		value := data[key]
		fmt.Printf("读取: %s = %s\n", key, value)
		return value
	}

	// 先执行一些写入
	var wg sync.WaitGroup

	// 启动3个写入goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			write(fmt.Sprintf("key%d", id), fmt.Sprintf("value%d", id))
		}(i)
	}

	// 确保写入完成
	wg.Wait()

	// 启动10个并发读取goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			keyID := id % 3 // 确保读取已写入的key
			read(fmt.Sprintf("key%d", keyID))
		}(i)
	}

	wg.Wait()

	// 注意事项:
	// 1. 读写锁适用于读多写少的场景
	// 2. 当有活跃的读锁时，写锁会被阻塞
	// 3. 当有活跃的写锁时，所有读锁和写锁都会被阻塞
	// 4. Go 1.9引入的sync.Map可能是更好的选择，适用于特定场景
}

// lockGranularity 展示锁的粒度选择问题
func lockGranularity() {
	// 粗粒度锁示例 - 整个数据结构使用一个锁
	type CoarseGrainedCache struct {
		mu    sync.Mutex
		cache map[string]string
	}

	type CacheSegment struct {
		mu    sync.Mutex
		value string
	}

	// 细粒度锁示例 - 每个键值对使用单独的锁
	type FineGrainedCache struct {
		segments map[string]*CacheSegment
		mu       sync.Mutex // 仅用于segments map的修改
	}

	// 创建粗粒度缓存
	coarse := &CoarseGrainedCache{
		cache: make(map[string]string),
	}

	// 创建细粒度缓存
	fine := &FineGrainedCache{
		segments: make(map[string]*CacheSegment),
	}

	// 粗粒度缓存的操作
	coarseSet := func(key, value string) {
		coarse.mu.Lock()
		defer coarse.mu.Unlock()

		// 模拟耗时操作
		time.Sleep(time.Millisecond * 10)
		coarse.cache[key] = value
	}

	// 细粒度缓存的操作
	fineSet := func(key, value string) {
		fine.mu.Lock()
		segment, exists := fine.segments[key]
		if !exists {
			segment = &CacheSegment{}
			fine.segments[key] = segment
		}
		fine.mu.Unlock()

		// 仅锁定特定的段
		segment.mu.Lock()
		defer segment.mu.Unlock()

		// 模拟耗时操作
		time.Sleep(time.Millisecond * 10)
		segment.value = value
	}

	fmt.Println("粗粒度锁和细粒度锁比较:")
	fmt.Println("- 粗粒度锁: 实现简单，但并发度低")
	fmt.Println("- 细粒度锁: 实现复杂，但并发度高")

	// 演示两种缓存的使用
	var wg sync.WaitGroup

	// 使用粗粒度锁
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			coarseSet(key, value)
			fmt.Printf("粗粒度缓存设置: %s=%s\n", key, value)
		}(i)
	}

	// 使用细粒度锁
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			fineSet(key, value)
			fmt.Printf("细粒度缓存设置: %s=%s\n", key, value)
		}(i)
	}

	wg.Wait()

	// 注意事项:
	// 1. 粗粒度锁容易导致锁竞争，降低并发性
	// 2. 细粒度锁提高并发性，但增加复杂性和死锁风险
	// 3. 从简单开始，根据性能测试结果再决定是否需要更细的锁粒度
	// 4. 考虑无锁数据结构或其他并发控制方法(如通道)
}

// avoidDeadlock 展示如何避免死锁
func avoidDeadlock() {
	// 死锁产生的条件:
	// 1. 互斥: 资源不能被共享，一次只有一个进程可以使用
	// 2. 请求与保持: 进程已经保持了至少一个资源，又提出了新的资源请求
	// 3. 不可剥夺: 进程获得的资源，在未使用完之前，不能被剥夺
	// 4. 循环等待: 若干进程之间形成头尾相接的循环等待资源关系

	// 避免死锁的方法:
	fmt.Println("避免死锁的方法:")
	fmt.Println("1. 总是按照相同的顺序获取锁")
	fmt.Println("2. 使用超时机制")
	fmt.Println("3. 死锁检测工具")
	fmt.Println("4. 避免嵌套锁")

	// 示例: 按固定顺序获取锁
	var lockA, lockB sync.Mutex

	// 正确方式: 总是先获取lockA再获取lockB
	safeFunction := func() {
		lockA.Lock()
		defer lockA.Unlock()

		// 执行一些操作
		time.Sleep(time.Millisecond)

		lockB.Lock()
		defer lockB.Unlock()

		// 使用两把锁保护的资源
		fmt.Println("安全地使用两个锁")
	}

	// 错误示例(可能导致死锁) - 不同goroutine获取锁的顺序不同
	// 在实际代码中应避免这种模式
	// deadlockRisk := func() {
	// 	// goroutine1获取锁的顺序: A -> B
	// 	go func() {
	// 		lockA.Lock()
	// 		defer lockA.Unlock()

	// 		time.Sleep(time.Millisecond) // 增加死锁可能性

	// 		lockB.Lock()
	// 		defer lockB.Unlock()

	// 		fmt.Println("使用锁A和B")
	// 	}()

	// 	// goroutine2获取锁的顺序: B -> A
	// 	go func() {
	// 		lockB.Lock()
	// 		defer lockB.Unlock()

	// 		time.Sleep(time.Millisecond) // 增加死锁可能性

	// 		lockA.Lock()
	// 		defer lockA.Unlock()

	// 		fmt.Println("使用锁B和A")
	// 	}()
	// }

	// 执行安全函数
	safeFunction()

	// 注释掉有死锁风险的函数调用
	// deadlockRisk() // 可能导致死锁

	// 其他避免死锁的技巧:
	// 1. 使用带超时的锁获取操作，如使用context或select+time.After
	// 2. 保持锁的使用简单，遵循最小权限原则
	// 3. 考虑使用通道(channel)代替显式锁
	// 4. 使用Go的race detector和死锁检测工具
}

// onceDemo 展示sync.Once的用法
func onceDemo() {
	// sync.Once确保某个函数只执行一次，常用于单例模式或一次性初始化
	var once sync.Once

	// 初始化函数
	initFunc := func() {
		fmt.Println("初始化操作 - 这只会执行一次")
		// 模拟耗时初始化
		time.Sleep(time.Millisecond * 100)
	}

	// 启动多个goroutine尝试执行初始化
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Goroutine %d 尝试执行初始化\n", id)
			once.Do(initFunc) // 只有第一个调用会执行initFunc
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}

	wg.Wait()

	// 注意事项:
	// 1. sync.Once比互斥锁+标志位更高效，也更安全
	// 2. sync.Once对每个函数是独立的，如果需要多个一次性初始化，需要多个Once实例
	// 3. 一旦Do方法返回，即使initFunc发生panic，该Once实例也被视为已使用
	// 4. 常见用途包括懒加载、单例模式、一次性配置加载等
}
