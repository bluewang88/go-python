package Go_base

import (
	"fmt"
	"os"
)

// HelloDefer 展示defer的各种用法和注意事项
func HelloDefer() {
	fmt.Println("===== 1. defer基本用法 =====")
	deferBasic()

	fmt.Println("\n===== 2. defer执行顺序 =====")
	deferOrder()

	fmt.Println("\n===== 3. defer与返回值 =====")
	result := deferReturnValue()
	fmt.Printf("函数最终返回值: %d\n", result)

	fmt.Println("\n===== 4. defer与闭包 =====")
	deferClosure()

	fmt.Println("\n===== 5. defer的实际应用 =====")
	deferPracticalUse()

	fmt.Println("\n===== 6. defer与panic/recover =====")
	deferPanicRecover()
}

// deferBasic 展示defer的基本用法
func deferBasic() {
	// defer 语句会将函数推迟到外层函数返回之前执行
	// 常用于资源清理、解锁等操作
	fmt.Println("函数开始执行")

	// defer语句注册的函数会在周围函数返回前调用
	defer fmt.Println("defer: 这会在函数结束时执行")

	fmt.Println("函数继续执行")
	// 当函数执行完这一行，之前注册的defer函数会执行
}

// deferOrder 展示defer的执行顺序
func deferOrder() {
	// 重要特性：defer语句按照LIFO(后进先出)的顺序执行
	// 即：最后一个defer语句最先被执行

	fmt.Println("函数开始")

	// 注册多个defer语句
	defer fmt.Println("defer 1: 这是第一个defer，但会最后执行")
	defer fmt.Println("defer 2: 这是第二个defer，会在中间执行")
	defer fmt.Println("defer 3: 这是第三个defer，会最先执行")

	fmt.Println("函数结束")
	// 执行顺序将是: defer 3, defer 2, defer 1
}

// deferReturnValue 展示defer与返回值的关系
func deferReturnValue() int {
	// 重要注意点：defer语句可以读取并修改命名返回值，但不能修改匿名返回值

	// 使用命名返回值
	var result int = 0

	defer func() {
		// 这里可以修改命名返回值
		result++
		fmt.Printf("defer: 修改返回值，现在是 %d\n", result)
	}()

	// 设置返回值
	result = 100
	fmt.Printf("return前: 返回值设置为 %d\n", result)

	// 返回值会被defer修改
	return result
	// 实际返回值将是101，而不是100
}

// deferClosure 展示defer与闭包的交互
func deferClosure() {
	// 重要注意点：defer语句的参数在defer语句出现时就已经计算好了
	// 但是如果defer后面跟的是闭包，闭包中的变量会在实际执行时才计算

	i := 1

	// 方式1：defer直接调用函数并传参，参数值在此时固定
	defer fmt.Printf("defer直接传参: i = %d\n", i)

	// 方式2：defer一个闭包，闭包中的变量在实际执行时才求值
	defer func() {
		fmt.Printf("defer闭包: i = %d\n", i)
	}()

	// 修改变量i的值
	i = 99
	fmt.Printf("函数中: i变成了 %d\n", i)

	// 输出将是:
	// 函数中: i变成了 99
	// defer闭包: i = 99
	// defer直接传参: i = 1
}

// deferPracticalUse 展示defer的实际应用
func deferPracticalUse() {
	// 最常见的用法1：确保资源释放
	file, err := os.Open("hello_defer.go")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	// 无论后续代码是否出错，都能确保文件被关闭
	defer file.Close()

	// 模拟文件操作
	fmt.Println("文件已打开，并将在函数结束时自动关闭")

	// 用法2：跟踪函数的进入和退出
	defer fmt.Println("函数执行结束")

	// 用法3：延迟解锁
	// 在实际代码中可能是mutex.Unlock()
	defer fmt.Println("资源已解锁")

	fmt.Println("资源已锁定，进行操作...")
}

// deferPanicRecover 展示defer与panic/recover的结合使用
func deferPanicRecover() {
	// defer和recover结合是Go语言处理异常的方式

	// 设置recover，捕获可能的panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到panic: %v\n", r)
			fmt.Println("程序已恢复运行")
		}
	}()

	fmt.Println("函数正常执行")

	// 制造一个panic
	panic("发生了一个严重错误")

	// 下面的代码不会执行
	fmt.Println("这行代码不会执行")
}

// 其他defer注意事项:
// 1. defer的性能开销：虽然很小，但在性能关键的热点代码中需要注意
// 2. defer在goroutine中：每个goroutine的defer独立作用
// 3. 不要在循环中使用defer：可能导致资源延迟释放
// 4. defer和goto：goto跳过defer声明可能导致资源泄漏
// 5. 在大型项目中统一defer的使用风格，增强代码可读性
