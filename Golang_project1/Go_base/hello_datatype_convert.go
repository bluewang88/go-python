package Go_base

import (
	"fmt"
	"reflect"
)

func HelloDatatypeConvert() {
	fmt.Println("---------------------------------打印基本数据类型转换---------------------------------")

	//go语言不支持变量之间的隐式转换，需要显示转换
	//以下代码会报错
	// var a int = 10
	// var b int8 = a
	// fmt.Println(b)

	var a int = 10.0 //10.0是float64类型，不能直接赋值给int类型,但是可以强制转换
	fmt.Println("类型转换var a int = 10.0输出a的值为", a, "a的类型为", reflect.TypeOf(a))

	// 整数类型转换
	// var a int = 10
	// var b int8 = 20
	// var c int16 = 30
	// var d int32 = 40
	// var e int64 = 50
	fmt.Println("---------------------------------类型转换结束---------------------------------")
}

// 浮点数类型转换
// 浮点数类型转换需要使用float64()函数
// 以下代码会报错
// var a float32 = 10.1
// var b float64 = a
