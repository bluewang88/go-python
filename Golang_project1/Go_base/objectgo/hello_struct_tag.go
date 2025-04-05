package objectgotest

// / go语言的结构体标签
// / 结构体标签是一个字符串，用来描述结构体字段的元数据
// / 结构体标签的语法是`json:"name"`，其中json是标签的名称，name是标签的值
// / 结构体标签可以用于json序列化和反序列化、数据库操作、表单验证等场景
// / 结构体标签可以通过反射获取
// / 结构体标签的值可以是任意字符串，可以包含空格、逗号等特殊字符
// / 结构体标签的值可以是多个标签，用空格分隔
// / 结构体标签的值可以是键值对，用冒号分隔
// / 结构体标签的值可以是多个键值对，用空格分隔

type Persontag struct {
	Name string `json:"name" xml:"name" yaml:"name"`
	Age  int    `json:"age" xml:"age" yaml:"age"`
}

func HelloStructTag() {
	p := Persontag{
		Name: "John",
		Age:  30,
	}
	// 获取结构体标签的值
	name := p.Name
	age := p.Age
	println(name, age)
}
