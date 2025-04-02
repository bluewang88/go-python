package Go_base

import (
	"fmt"
)

func HelloSliceFuncParam() {
	// 函数参数传递的是值
	// 函数参数传递的是值，所以函数内部对slice的修改不会影响调用者
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)

	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}

// 当切片作为函数参数传递时：
// 是值传递：复制了一份切片结构体（包含指针、长度和容量）
// 但不是深拷贝：没有复制底层数组的数据

// 值传递与引用传递的定义
// 值传递(Pass by Value):

// 函数接收的是参数的一个副本
// 函数内部对参数的修改不会影响原始变量
// 整个数据结构被复制到函数中
// 引用传递(Pass by Reference):

// 传递的是变量的引用（地址）
// 函数内部对参数的修改直接影响原始变量
// 不复制数据，仅传递访问原始数据的方式
func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	for _, i := range s {
		fmt.Println(i)

	}

	for i := range s {
		s[i] += 1 // 这里的s[i]是对原始切片的引用，所以会改变原始切片的值
	}
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}
