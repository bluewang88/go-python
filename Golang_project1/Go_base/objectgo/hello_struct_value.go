package objectgotest

import "fmt"

//struct的值传递
// 结构体的值传递和指针传递
// 1. 结构体的值传递：将结构体的值传递给函数，函数内部对结构体的修改不会影响到原结构体
// 2. 结构体的指针传递：将结构体的指针传递给函数，函数内部对结构体的修改会影响到原结构体
// 3. 结构体的值传递和指针传递的区别在于：值传递会复制一份结构体的值，指针传递只会复制结构体的地址

type Student struct {
	Name  string
	Age   int
	Sex   string
	Score float64
}

func (student *Student) SetName(name string) {
	student.Name = name
}

func (student Student) SetAge(age int) {
	student.Age = age
}

func HelloStructValue() {
	student := Student{
		Name:  "张三",
		Age:   18,
		Sex:   "男",
		Score: 90.0,
	}

	student.SetName("李四") // 修改结构体
	student.SetAge(20)    // 修改结构体

	fmt.Println(student)
}
