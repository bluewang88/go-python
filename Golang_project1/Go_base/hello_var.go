package Go_base

import "fmt"

func HelloVar() {

	//批量声明变量
	var (
		a int
		b string
		c bool
	)

	fmt.Println("批量声明变量:", "a=", a, "b=", b, "c=", c)

	var foo1 int // 只声明，不做初始化
	fmt.Println("只声明，不做初始化:", "foo1=", foo1)
	fmt.Println("--------------------------------")
	var foo2 int = 42 // 声明的同时做初始化
	fmt.Println("声明的同时做初始化:", "foo2=", foo2)
	fmt.Println("--------------------------------")
	var foo3, bar3 int = 42, 1302 // 一次声明和初始化多个变量
	fmt.Println("一次声明和初始化多个变量:", "foo3=", foo3, "bar3=", bar3)
	fmt.Println("--------------------------------")
	var foo4 = 42 // 忽略类型，编译器自行推导
	fmt.Println("忽略类型，编译器自行推导:", "foo4=", foo4)
	fmt.Println("--------------------------------")
	foo5 := 42 // 简写，只能在函数或者方法体内使用，没有var关键字，变量类型也是隐式推导而来
	fmt.Println("简写，只能在函数或者方法体内使用:", "foo5=", foo5)
	fmt.Println("--------------------------------")
	const constant = "This is a constant" // 常量
	fmt.Println("常量:", "constant=", constant)
	fmt.Println("--------------------------------")

	// iota是常量计数器
	// iota的值从0开始，用于常量的数值递增
	// 常量组中，如果第一个常量没有赋值，则默认会赋值为0
	// 常量组中，则默认当常量没有明确赋值时，它会继承前一个常量的表达式
	// _ 表示跳过, 相当于被忽略
	// << 位运算符号, 左移运算符
	const (
		_      = iota      // 第一行：iota = 0，但使用_忽略这个值
		consta             // 第二行：iota = 1，因为没有明确赋值，所以采用iota的值，所以consta = 1
		constb             // 第三行：iota = 2，同上，所以constb = 2
		constc = 2 << iota // 第四行：iota = 3，所以这里计算 2 << 3 = 16
		constd             // 第五行：iota = 4，沿用上一行的表达式，所以是 2 << 4 = 32
		conste = 9         // 第六行：iota = 5，conste = 9
		constf             // 第七行：iota = 6，沿用上一行的表达式，所以是constf = 9
	)
	fmt.Println(consta, constb) // 1 2 (0被赋值给了_，相当于被跳过了)
	fmt.Println(constc, constd) // 16 32
	fmt.Println(conste, constf) // 9 9
	fmt.Println("--------------------------------")
}
