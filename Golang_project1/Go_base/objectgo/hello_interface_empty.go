package objectgotest

// go语言的空接口

type EmptyInterface interface{}

// 空接口是没有任何方法的接口
// 空接口可以被任何类型实现
// 空接口可以用来表示任何类型的值
// 空接口可以用来实现多态，即同一个接口可以被不同的类型实现
// 空接口可以用来实现函数参数的多态
// 空接口可以用来实现函数返回值的多态
// 空接口可以用来实现函数的参数和返回值的多态

func HelloInterfaceEmpty(any ...EmptyInterface) EmptyInterface {
	// 空接口可以用来表示任何类型的值
	var i EmptyInterface
	for _, value := range any {
		i = value
		println(i)
	}
	return i
}

//空接口可以为作为map的值

func HelloInterfaceEmptyMap() {
	// 空接口可以用来表示任何类型的值
	var i EmptyInterface
	var m = make(map[string]EmptyInterface)
	x := make(map[string]interface{})
	y := make(map[string]EmptyInterface)
	// interface{} 是 Go 语言的内置类型名称，它是空接口的完整表示方式。在 Go 语言中，空接口的正确写法是包含花括号的 interface{}。
	// EmptyInterface 是你在代码中自定义的类型名称（通过 type EmptyInterface interface{}），它是一个命名类型，使用时不需要花括号。

	m["a"] = 1
	m["b"] = "hello"
	m["c"] = 3.14
	for k, v := range m {
		println(k, v)
	}
	i = m
	println(i)

	x["d"] = true
	x["e"] = []int{1, 2, 3}
	for k, v := range x {
		println(k, v)
	}
	i = x
	println(i)

	y["f"] = 4.5
	y["g"] = map[string]int{"a": 1, "b": 2}
}

// 空接口可以用来表示任何类型的值}
