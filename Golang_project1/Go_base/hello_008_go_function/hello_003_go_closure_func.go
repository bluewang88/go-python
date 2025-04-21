package Go_Base_func

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func HelloClosureFunc() {
	f1, f2 := calc(10) //调用 calc(10) 时，base 被初始化为 10
	//f1 和 f2 分别是 add 和 sub 的闭包函数，它们共享同一个 base 变量。
	fmt.Println(f1(1), f2(2)) //11 9
	//f1(1)：调用 add，base += 1，base 变为 11，返回 11。
	// f2(2)：调用 sub，base -= 2，base 变为 9，返回 9。
	fmt.Println(f1(3), f2(4)) //12 8
	// 	f1(3)：调用 add，base += 3，base 从 9 变为 12，返回 12。
	// f2(4)：调用 sub，base -= 4，base 从 12 变为 8，返回 8。
	fmt.Println(f1(5), f2(6)) //13 7
	//	f1(5)：调用 add，base += 5，base 从 8 变为 13，返回 13。
	// f2(6)：调用 sub，base -= 6，base 从 13 变为 7，返回 7。
}
