package objectgotest

// 设置类的主要作用是将方法绑定到类型。class主要是封装数据和方法
// 1. struct类型可以存储一组数据
// 2. struct类型可以定义方法

type Course struct {
	Name  string
	Price int
	Url   string
}

// 结构体的方法
func (c Course) CourseResistor() {
	println("Hello, I am a CourseResistor")
	println("Course Name:", c.Name)
	println("Course Price:", c.Price)
	println("Course URL:", c.Url)
}

func HelloStructFunc() {
	c := Course{
		Name:  "Go",
		Price: 100,
		Url:   "https://www.imooc.com"}
	c.CourseResistor()
}
