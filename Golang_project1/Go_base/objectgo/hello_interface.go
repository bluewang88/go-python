package objectgotest

import "fmt"

// / go语言的接口,是实现OOP特性多态的基础
// / 接口是一组方法的集合，接口定义了一个类型应该具有的方法，但不实现这些方法
// / 接口可以被任何实现了这些方法的类型所实现
// / 接口可以被用来实现多态，即同一个接口可以被不同的类型实现
// / 多态的含义就是不同对象对同一消息使用相同的方法名但做出不同的响应

type Speaker interface {
	Speak() string //方法只是声明
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// 多态函数
func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

// 使用

func HelloInterface() {
	dog := Dog{"Rover"}
	cat := Cat{"Whiskers"}
	MakeSpeak(dog) // 输出: Rover says Woof!
	MakeSpeak(cat) // 输出: Whiskers says Meow!

	animal := []Speaker{dog, cat} // 创建一个包含不同类型的切片 类型为Speaker
	for _, a := range animal {
		fmt.Println(a.Speak())
	}

	var s Speaker = dog    // 接口类型的变量可以存储任何实现了该接口的类型的值
	fmt.Println(s.Speak()) // 输出: Rover says Woof!
	// 下面的代码会报错，因为s是一个接口类型，不能直接赋值给一个具体类型
	// dog = s // 错误: cannot use s (type Speaker) as type Dog  in assignment: Speaker does not implement Dog (missing Speak method)

}
