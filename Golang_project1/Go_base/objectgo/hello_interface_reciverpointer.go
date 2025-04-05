package objectgotest

import "fmt"

//当结构体方法使用指针接收者时，只有该结构体的指针类型实现了接口
// 结构体的值类型没有实现该接口

// 定义一个接口
type Updater interface {
	Update(newValue string)
	GetValue() string
}

// 定义一个结构体
type Document struct {
	content string
}

// 使用指针接收者实现Update方法
// 这里必须使用指针接收者，因为我们需要修改结构体的状态
func (d *Document) Update(newValue string) {
	d.content = newValue // 修改结构体的字段
}

// 获取值的方法也使用指针接收者保持一致性
func (d *Document) GetValue() string {
	return d.content
}

// 演示函数
func HelloInterfaceReceiverPointer() {
	// 创建Document实例
	doc := &Document{content: "初始内容"}

	// 将指针赋值给接口变量
	var updater Updater = doc

	// 初始状态
	fmt.Printf("初始内容: %s\n", updater.GetValue())

	// 通过接口调用Update方法
	updater.Update("已更新的内容")

	// 验证内容已被修改
	fmt.Printf("更新后内容: %s\n", updater.GetValue())

	// 注意：以下代码是错误的，会导致编译错误
	// var updater2 Updater = Document{content: "不能工作"}
	// 错误原因：Document类型的值不能赋值给Updater接口
	// 只有*Document类型实现了Updater接口

	// 正确的做法
	doc2 := Document{content: "另一个文档"}
	var updater2 Updater = &doc2 // 必须取地址
	fmt.Printf("另一个文档: %s\n", updater2.GetValue())
}
