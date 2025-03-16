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

	// 常量到变量之间支持隐式转换，但是不可以有数据丢失
	var a int = 10.0 //10.0是常量，也是float64类型，不能直接赋值给int类型,但是可以强制转换
	fmt.Println("类型转换var a int = 10.0输出a的值为", a, "a的类型为", reflect.TypeOf(a))

	//显式类型转换
	var floatVal float64 = 8.1
	var b int = int(floatVal) //8.1是float64类型，需要强制转换为int类型
	fmt.Println("类型转换var b int = int(floatVal)输出b的值为", b, "b的类型为", reflect.TypeOf(b))

	// 浮点数
	c := 5.0
	// 转换为int类型
	d := int(c)
	fmt.Println("类型转换var d int = int(c)输出d的值为", d, "d的类型为", reflect.TypeOf(d))

	// Go允许在底层结构相同的两个类型之间互转。例如：
	// IT类型的底层是int类型
	type IT int

	// e的类型为IT，底层是int
	var e IT = 5

	// 将e(IT)转换为int，f现在是int类型
	f := int(e)

	// 将f(int)转换为IT，g现在是IT类型
	g := IT(f)
	fmt.Println("类型转换var g IT = IT(f)输出g的值为", g, "g的类型为", reflect.TypeOf(g))

	var h int32 = 1
	var i int64 = 3
	i = int64(h) + i
	fmt.Println("类型转换var i int64 = int64(h) + i输出i的值为", i, "i的类型为", reflect.TypeOf(i))

	/*
	   不是所有数据类型都能转换的，例如字母格式的string类型"abcd"转换为int肯定会失败
	   低精度转换为高精度时是安全的，高精度的值转换为低精度时会丢失精度。例如int32转换为int16，float32转换为int
	   这种简单的转换方式不能对int(float)和string进行互转，要跨大类型转换，可以使用strconv包提供的函数
	*/

	fmt.Println("---------------------------------类型转换结束---------------------------------")
}

// 浮点数类型转换
// 浮点数类型转换需要使用float64()函数
// 以下代码会报错
// var a float32 = 10.1
// var b float64 = a
