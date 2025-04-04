package objectgo

import "fmt"

//type的使用方式
//go不支持面向对象编程，面向对象编程特征：1.封装 2.继承 3.多态
// 但是可以使用type来定义新的类型，
// 1. 给一个类型起一个别名 type <newType> = <oldType>
// 2. 用一个已有的类型作为基础，创建一个新的类型 type <newType> <oldType>
// 3. 定义一个新的结构体类型 type <newType> struct { ... }
// 4. 定义一个新的接口类型 type <newType> interface { ... }
// 5. 定义一个新的函数类型 type <newType> func(<params>) <returnType>
// 6. 定义一个新的数组类型 type <newType> [<size>]<elementType>
// 7. 定义一个新的切片类型 type <newType> []<elementType>
// 8. 定义一个新的映射类型 type <newType> map[<keyType>]<valueType>
// 9. 定义一个新的通道类型 type <newType> chan <elementType>
// 10. 定义一个新的指针类型 type <newType> *<elementType>

func HelloType() {
	// 1. 给一个类型起一个别名 type <newType> = <oldType>
	type myInt int
	var a myInt = 10
	// println(a)
	// 这里的a是myInt类型，底层是int类型
	fmt.Println("a的值为", a)
	fmt.Printf("%T\n", a) // 输出类型

	// 2. 用一个已有的类型作为基础，创建一个新的类型 type <newType> <oldType>
	type myInt2 = int
	var b myInt2 = 20
	// println(b)
	fmt.Printf("%T\n", b) // 输出类型

	var c myStruct
	c.name = "张三"
	c.age = 18
	fmt.Printf("%T\n", c)
	// println(c.name, c.age)

	c.SayHello() // 调用实现了myInterface接口的结构体myStruct的方法

	e := myInterface(c)   // 将myStruct类型的变量c转换为myInterface类型
	fmt.Printf("%T\n", e) // 输出类型

	// 5. 定义一个新的函数类型 type <newType> func(<params>) <returnType>
	type myFunc func(int, int) int
	var d myFunc = func(a, b int) int {
		return a + b
	}
	result := d(1, 2)
	fmt.Println("函数类型的结果为", result)

}

// 3. 定义一个新的结构体类型 type <newType> struct { ... }
type myStruct struct {
	name string
	age  int
}

// 4. 定义一个新的接口类型 type <newType> interface { ... }
type myInterface interface {
	SayHello() // 定义一个方法
}

//定义一个实现了myInterface接口的结构体myStruct

func (m myStruct) SayHello() {
	fmt.Println("Hello, my name is", m.name)
}
