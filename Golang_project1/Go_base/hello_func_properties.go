package Go_base

import "fmt"

func Hello_func_properties_args(a, b string, floatargs float64) (sum int, count int, count2 float64) {
	sum = 1
	count = 2
	count2 = 3.0
	return
}

func Hello_func_properties() {
	MyFunc := Hello_func_properties_args

	// 函数作为参数传递
	_, _, _ = MyFunc("a", "b", 3.0)

	sum, count, count2 := MyFunc("a", "b", 3.0)
	fmt.Println(sum, count, count2)
	fmt.Printf("MyFunc is %T\n", MyFunc)

	//匿名函数
	MyFunc2 := func(a, b string, floatargs float64) int {
		sum := 1
		count := 2
		count2 := 3.0
		return sum + count + int(count2)
	}("1", "2", 3.0) //匿名函数调用

	fmt.Println(MyFunc2)
	fmt.Printf("MyFunc2 is %T\n", MyFunc2)
	fmt.Println("Hello from Hello_func_properties!")
	fmt.Println("End of function.")

}

// 在 Go 语言中，函数被称为**一等公民（First-Class Citizens）**，这意味着函数可以像其他变量一样被操作和使用。具体来说，函数作为一等公民具有以下特性：

// ---

// ### 1. **可以赋值给变量**
// 函数可以被赋值给变量，并通过该变量调用。例如：

// ```go
// package main

// import "fmt"

// func add(a, b int) int {
//     return a + b
// }

// func main() {
//     // 将函数赋值给变量
//     sum := add
//     fmt.Println(sum(3, 4)) // 输出：7
// }
// ```

// ---

// ### 2. **可以作为参数传递**
// 函数可以作为参数传递给另一个函数。例如：

// ```go
// package main

// import "fmt"

// // 定义一个函数类型
// type operation func(int, int) int

// func calculate(a, b int, op operation) int {
//     return op(a, b)
// }

// func add(a, b int) int {
//     return a + b
// }

// func main() {
//     result := calculate(5, 3, add) // 将函数作为参数传递
//     fmt.Println(result)           // 输出：8
// }
// ```

// ---

// ### 3. **可以作为返回值**
// 函数可以作为另一个函数的返回值。例如：

// ```go
// package main

// import "fmt"

// func multiplier(factor int) func(int) int {
//     return func(value int) int {
//         return value * factor
//     }
// }

// func main() {
//     double := multiplier(2) // 返回一个函数
//     fmt.Println(double(5))  // 输出：10
// }
// ```

// ---

// ### 4. **可以嵌套定义（匿名函数）**
// 函数可以在另一个函数内部定义，并作为匿名函数使用。例如：

// ```go
// package main

// import "fmt"

// func main() {
//     // 定义一个匿名函数并立即调用
//     result := func(a, b int) int {
//         return a + b
//     }(3, 4)
//     fmt.Println(result) // 输出：7
// }
// ```

// ---

// ### 5. **可以存储在数据结构中**
// 函数可以存储在数组、切片或映射等数据结构中。例如：

// ```go
// package main

// import "fmt"

// func add(a, b int) int {
//     return a + b
// }

// func subtract(a, b int) int {
//     return a - b
// }

// func main() {
//     operations := map[string]func(int, int) int{
//         "add":      add,
//         "subtract": subtract,
//     }

//     fmt.Println(operations["add"](10, 5))      // 输出：15
//     fmt.Println(operations["subtract"](10, 5)) // 输出：5
// }
// ```

// ---

// ### 总结
// 函数作为一等公民，意味着它们可以像普通变量一样被赋值、传递、返回和存储。这种特性使得 Go 语言在处理高阶函数、回调函数和函数式编程模式时非常灵活。
