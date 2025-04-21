package Go_Base_func

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// x 被初始化为 5。
// defer 注册了一个匿名函数，该函数会在 return 语句执行后运行。
// return x 将 x 的值（5）保存到返回值中。
// defer 的匿名函数执行，将 x 的值加 1，但此时 x 是局部变量，修改不会影响已经保存的返回值。
// 返回值是 5，因为 defer 修改的是局部变量 x，而不是返回值。

func f2() (x int) {
	defer func() {
		x++
	}()
	// 重要注意点：defer语句的参数在defer语句出现时就已经计算好了
	// 但是如果defer后面跟的是闭包，闭包中的变量会在实际执行时才计算
	return 5 // 将 5 赋值给命名返回值 x
}

// x 是命名返回值，初始值为 0。
// return 5 将 5 赋值给命名返回值 x。
// defer 的匿名函数执行，将 x 的值加 1。
// 最终返回 x 的值。
// 返回值是 6，因为 defer 修改了命名返回值 x。

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// x 被初始化为 5。
// return x 将 x 的值（5）赋值给命名返回值 y。
// defer 的匿名函数执行，将局部变量 x 的值加 1，但这不会影响返回值 y。
// 返回值：
// 返回值是 5，因为 defer 修改的是局部变量 x，而不是返回值 y。

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 命名返回值初始化

// hello_004_go_defer_1.go ) 是命名返回值，被自动初始化为 int 的零值 0
// defer 注册阶段

// 遇到 defer 语句，Go 会立即计算并保存延迟函数的参数值
// 此时 hello_004_go_defer_1.go ) 的值是 0，所以 0 被作为参数传递给延迟函数
// 这个参数值会被复制一份，在之后实际调用延迟函数时使用
// 返回值赋值阶段

// 执行到 return 5，将 5 赋值给命名返回值 hello_004_go_defer_1.go )
// 此时命名返回值 hello_004_go_defer_1.go ) 从 0 变成了 5
// defer 函数执行阶段

// 在函数即将返回前，执行之前注册的延迟函数
// 延迟函数使用之前保存的参数值 0，将其作为局部变量 hello_004_go_defer_1.go )
// 执行 hello_004_go_defer_1.go )，使得延迟函数内的参数 hello_004_go_defer_1.go ) 的值变为 1
// 关键点：这个 hello_004_go_defer_1.go ) 是参数副本，与外部的命名返回值 hello_004_go_defer_1.go ) 是不同的变量，尽管它们同名
// 函数返回阶段

// 延迟函数执行完毕后，函数 hello_004_go_defer_1.go ) 返回命名返回值 hello_004_go_defer_1.go ) 的最终值
// 由于延迟函数修改的是参数副本，不是命名返回值，所以返回值仍然是 5

func f5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// x 是命名返回值，初始值为 0。
// return 5 将 5 赋值给命名返回值 x。
// defer 注册的匿名函数接收了 x 的值（5）作为参数，但这是值传递，匿名函数中的 x 是一个副本。
// 匿名函数中的 x++ 只修改了副本，不影响外部的 x。
// 返回值：
// 返回值是 5，因为 defer 修改的是参数副本，而不是返回值 x。

func HelloDefer() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
}

// Output:
// 5
// 6
// 5
// 5
// 说明：
// 1. 在 f1() 中，defer 语句在函数返回之前执行，但由于 x 是局部变量，所以它的值在 defer 执行时仍然是 5。
// 2. 在 f2() 中，defer 语句在函数返回之前执行，并且 x 是命名返回值，所以它的值在 defer 执行时已经是 6。
// 3. 在 f3() 中，defer 语句在函数返回之前执行，但由于 x 是局部变量，所以它的值在 defer 执行时仍然是 5。
// 4. 在 f4() 中，defer 语句在函数返回之前执行，并且 x 是命名返回值，所以它的值在 defer 执行时仍然是 5。
