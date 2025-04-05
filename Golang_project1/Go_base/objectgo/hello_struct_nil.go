package objectgotest

// 结构体初始化方式
// 结构体初始化的零值

import "fmt"

// 1. 使用new关键字创建结构体指针
// 2. 使用&结构体变量创建结构体指针
// 3. 使用结构体字面量创建结构体指针

type Person struct {
	X int
	Y int
}

func HelloStructNil() {
	// 1. 使用new关键字创建结构体指针
	p1 := new(Person)
	fmt.Println(p1)

	// 2. 使用&结构体变量创建结构体指针
	p2 := Person{}
	fmt.Println(p2)

	// 3. 使用结构体字面量创建结构体指针
	p3 := Person{X: 1, Y: 2}
	fmt.Println(p3)
}
