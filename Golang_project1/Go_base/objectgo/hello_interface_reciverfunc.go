package objectgotest

import "fmt"

// 如果结构体实现接口的方法中有任何一个使用了指针接收者，
// 那么只能将该结构体的指针赋值给接口变量，不能将结构体的值直接赋值给接口变量。这是因为接口需要实现所有方法，而值类型无法调用指针接收者的方法。
// 指针接收者的优先级：

// 如果任何一个接口方法使用了指针接收者，那么只有结构体的指针类型才实现了该接口
// 即使其他方法使用值接收者，你也必须使用结构体指针赋值给接口变量

// 定义一个接口
type Processor interface {
	Process()
	Display()
}

// 定义一个结构体
type Data struct {
	Value string
}

// 使用指针接收者实现Process方法
func (d *Data) Process() {
	d.Value = "已处理: " + d.Value
	fmt.Println("使用指针接收者处理数据")
}

// 使用值接收者实现Display方法
func (d Data) Display() {
	fmt.Println("显示数据:", d.Value)
}

func HelloInterfaceReceiverfunc() {
	// 创建结构体实例
	data := Data{Value: "原始数据"}

	// 尝试将值赋给接口
	// var p1 Processor = data // 这行会导致编译错误!

	// 正确方式：使用指针
	var p2 Processor = &data

	p2.Process()
	p2.Display()
}
