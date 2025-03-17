package Go_base

import (
	"fmt"
	"reflect"
)

// HelloDatatypeConvert 函数用于演示Go语言中的数据类型转换。
// 该函数展示了Go语言中基本数据类型的显式转换、常量到变量的隐式转换，
// 以及底层结构相同的类型之间的转换。
// 函数中还展示了浮点数到整数的转换，以及不同整数类型之间的转换。
// 最后，函数总结了数据类型转换的一些限制和注意事项。
func HelloDatatypeConvert() {
	fmt.Println("---------------------------------打印基本数据类型转换---------------------------------")

	// Go语言不支持变量之间的隐式转换，必须进行显式转换。
	// 以下代码会报错，因为不能直接将int类型的变量赋值给int8类型的变量。
	// var a int = 10
	// var b int8 = a
	// fmt.Println(b)

	// 常量到变量之间支持隐式转换，但前提是不能有数据丢失。
	// 10.0是常量，也是float64类型，不能直接赋值给int类型，但可以强制转换。
	var a int = 10.0
	fmt.Println("类型转换var a int = 10.0输出a的值为", a, "a的类型为", reflect.TypeOf(a))

	// 显式类型转换示例：将float64类型的变量转换为int类型。
	var floatVal float64 = 8.1
	var b int = int(floatVal)
	fmt.Println("类型转换var b int = int(floatVal)输出b的值为", b, "b的类型为", reflect.TypeOf(b))

	// 浮点数转换为整数类型示例。
	c := 5.0
	d := int(c)
	fmt.Println("类型转换var d int = int(c)输出d的值为", d, "d的类型为", reflect.TypeOf(d))

	// Go允许在底层结构相同的两个类型之间进行转换。
	// IT类型的底层是int类型。
	type IT int

	// e的类型为IT，底层是int。
	var e IT = 5

	// 将e(IT)转换为int，f现在是int类型。
	f := int(e)

	// 将f(int)转换为IT，g现在是IT类型。
	g := IT(f)
	fmt.Println("类型转换var g IT = IT(f)输出g的值为", g, "g的类型为", reflect.TypeOf(g))

	// 不同整数类型之间的转换示例。
	var h int32 = 1
	var i int64 = 3
	i = int64(h) + i
	fmt.Println("类型转换var i int64 = int64(h) + i输出i的值为", i, "i的类型为", reflect.TypeOf(i))

	/*
	   数据类型转换的限制和注意事项：
	   - 不是所有数据类型都能转换，例如字母格式的string类型"abcd"转换为int会失败。
	   - 低精度转换为高精度时是安全的，高精度的值转换为低精度时会丢失精度。
	   - 简单的转换方式不能对int(float)和string进行互转，跨大类型转换可以使用strconv包提供的函数。
	*/

	fmt.Println("---------------------------------类型转换结束---------------------------------")
}

// 浮点数类型转换
// 浮点数类型转换需要使用float64()函数
// 以下代码会报错
// var a float32 = 10.1
// var b float64 = a
