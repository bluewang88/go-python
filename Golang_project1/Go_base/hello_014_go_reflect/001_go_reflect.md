# go语言的反射

## 变量的内在机制
Go语言中的变量是分为两部分的:

类型信息：预先定义好的元信息。
值信息：程序运行过程中可动态变化的。

## 反射的概念
- 反射是指在程序运行期间对程序本身进行访问和修改的能力。
- 程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
- 反射就是程序在运行时，可以获取变量的类型、值等信息。
- 反射是Go语言的一个强大特性，它允许程序在运行时检查类型和变量的值。通过反射，我们可以动态地获取变量的类型、值以及其他信息。
- 支持反射的语言可以在程序编译期间将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。
- Go程序在运行期间使用reflect包访问程序的反射信息。
- 反射的实现是通过类型的接口实现的，Go语言中的所有类型都实现了`reflect.Type`和`reflect.Value`接口。
- 反射的实现是通过`reflect`包实现的，`reflect`包提供了对Go语言类型系统的访问。

## reflect包
在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的(我们在上一篇接口的博客中有介绍相关概念)。 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

### reflect.TypeOf
- 在Go语言中，使用`reflect.TypeOf()`函数可以获得任意值的`类型对象`（reflect.Type），程序通过类型对象可以访问任意值的类型信息。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}
```
- `reflect.TypeOf(x)`：返回x的类型对象。
- `v`：表示类型对象。
- `fmt.Printf("type:%v\n", v)`：打印类型对象的值。

### type name 和 type kind
在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）
在Go语言中，类型对象提供了`Name()`和`Kind()`两个方法，分别用于获取类型对象的名称和种类。
- `reflect.TypeOf(x)`：返回x的类型对象。
- `v.Name()`：返回类型对象的名称。
- `v.Kind()`：返回类型对象的种类。
  
```go
package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct
}
```
- 在 Go 反射中，只有命名类型才会有名称，而指针、切片等复合类型没有独立的名称。
- Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
- 指针类型是复合类型，它们没有直接的类型名称，这就是为什么 t.Name() 返回空字符串。

Kind 类型如下：
```go
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
```

### reflect.ValueOf
- 在Go语言中，使用`reflect.ValueOf()`函数可以获得任意值的值对象（reflect.Value），程序通过值对象可以访问任意值的值信息。

|方法|说明|
| --- | --- |
|Interface() interface {} | 	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型 |
|Int() int64 | 将值以 int 类型返回，所有有符号整型均可以此方式返回 |
|Uint() uint64 | 将值以 uint 类型返回，所有无符号整型均可以此方式返回 |
|Float() float64 | 将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回 |
|Bool() bool | 将值以 bool 类型返回 |
|Bytes() []bytes | 将值以字节数组 []bytes 类型返回 |
|String() string | 将值以字符串类型返回 |

### 通过反射获取值
```go
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func main() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}
```

### Elem（）
- `Elem()`方法用于获取指针指向的值对象。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func main() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
```

- `v.Elem()`：获取指针指向的值对象。
- `v.Elem().SetInt(200)`：修改指针指向的值对象的值。
- `v.SetInt(200)`：修改的是副本，reflect包会引发panic。
- `v.Elem().SetInt(200)`：修改指针指向的值对象的值。
- `fmt.Println(a)`：打印修改后的值。

### isNil 和 isValid

- `IsNil()`方法用于判断值对象是否为nil。
- `IsValid()`方法用于判断值对象是否有效。

IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。


## 结构体反射

任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。

reflect.Type中与获取结构体成员相关的的方法如下表所示。
|方法|	说明|
| --- | --- |
|Field(i int) StructField	|根据索引，返回索引对应的结构体字段的信息。|
|NumField() int	|返回结构体成员字段数量。|
|FieldByName(name string) (StructField, bool)	|根据给定字符串返回字符串对应的结构体字段的信息。|
|FieldByIndex(index []int) StructField	|多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。|
|FieldByNameFunc(match func(string) bool) (StructField,bool)	|根据传入的匹配函数匹配需要的字段。|
|NumMethod() int	|返回该类型的方法集中方法的数目|
|Method(int) Method	|返回该类型方法集中的第i个方法|
|MethodByName(string)(Method, bool)	|根据方法名返回该类型方法集中的方法|


### StructField类型
- `StructField`是一个结构体类型，表示结构体的字段信息。
- `StructField`结构体的定义如下：
```go
type StructField struct {
	Name      string      // 字段名称
	PkgPath   string      // 包路径
	Type      Type        // 字段类型
	Tag       StructTag   // 字段标签
	Offset    uintptr     // 字段在结构体中的偏移量
	Index     []int       // 字段在结构体中的索引
	Anonymous bool        // 是否匿名字段
}
```

结构体反射示例
```go
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}



// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

```


## 反射的应用场景

- 反射可以用于实现通用的函数库，比如序列化和反序列化、ORM框架等。
- 反射可以用于实现动态类型检查，比如在运行时检查变量的类型。
- 反射可以用于实现动态调用，比如在运行时调用函数。
- 反射可以用于实现动态创建对象，比如在运行时创建对象。
- 反射可以用于实现动态注册，比如在运行时注册函数。
- 反射可以用于实现动态配置，比如在运行时读取配置文件。
- 反射可以用于实现动态代理，比如在运行时代理函数。
- 反射可以用于实现动态注入，比如在运行时注入依赖。
- 反射可以用于实现动态路由，比如在运行时路由请求。
- 反射可以用于实现动态编译，比如在运行时编译代码。
- 反射可以用于实现动态加载，比如在运行时加载模块。
- 反射可以用于实现动态链接，比如在运行时链接库。
- 反射可以用于实现动态卸载，比如在运行时卸载模块。
- 反射可以用于实现动态更新，比如在运行时更新代码。
- 反射可以用于实现动态热更新，比如在运行时热更新代码。
- 反射可以用于实现动态调试，比如在运行时调试代码。