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

type Personconf struct {
	Name    string
	Age     int
	Hobbies []string       // 切片 - 引用类型
	Scores  map[string]int // 映射 - 引用类型
	Address *string        // 指针 - 引用类型
	Friends []*Personconf  // 指针切片 - 引用类型
}

// 初始化时的关键点
// 切片类型初始化：

// 使用字面量 []Type{}
// 使用 make([]Type, length, capacity)
// Map 类型初始化：

// 使用字面量 map[KeyType]ValueType{}
// 使用 make(map[KeyType]ValueType)
// 指针类型初始化：

// 取某个变量的地址 &variable
// 使用 new(Type)
// 嵌套引用类型：

// 根据需要逐层初始化

func HelloStructWithReferenceTypes() {
	// 方法1：声明时直接初始化所有字段
	address := "北京市海淀区"
	p1 := Personconf{
		Name:    "张三",
		Age:     25,
		Hobbies: []string{"读书", "游泳"},
		Scores:  map[string]int{"数学": 90, "语文": 85},
		Address: &address,
		Friends: []*Personconf{}, // 初始化为空切片
	}
	fmt.Println(p1)

	// 方法2：先声明，再分别初始化各引用类型字段
	var p2 Personconf
	p2.Name = "李四"
	p2.Age = 30
	p2.Hobbies = make([]string, 0) // 或 p2.Hobbies = []string{}
	p2.Scores = make(map[string]int)

	addr := "上海市浦东新区"
	p2.Address = &addr
	p2.Friends = make([]*Personconf, 0)

	fmt.Println(p2)

	// 方法3：使用 new() 函数创建指针，再初始化
	p3 := new(Personconf)
	p3.Name = "王五"
	p3.Hobbies = []string{"篮球"}
	p3.Scores = make(map[string]int)

	fmt.Println(p3)

	// 示例使用
	p2.Hobbies = append(p2.Hobbies, "旅游")
	p2.Scores["英语"] = 92
}
