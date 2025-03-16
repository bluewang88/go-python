package Go_base

import "fmt"

func HelloDefaultVar() {
	//字符串类型
	var testString string
	//整型
	var testint int
	//float
	var testfloat float32
	//切片类型
	var tesslice []int
	//函数类型
	var testfunc func(int, int) int
	//map类型
	var testmap map[string]int
	//接口类型
	var testinterface interface{}
	//指针类型
	var testpoint *int
	//通道类型
	var testchan chan int
	//bool
	var testbool bool
	//打印变量的默认值
	fmt.Println("-----------打印变量的默认值---开始------")
	fmt.Println("testvar的Strng值：", testString)
	fmt.Println("testint的值：", testint)
	fmt.Println("testfloat的值：", testfloat)
	fmt.Println("tesslice的值：", tesslice)
	//在Go语言中，func类型的变量如果不加括号直接传递给fmt.Println，会被视为函数值而非调用。
	//因此，编译器会报错，提示testfunc是一个函数值，未被调用。
	// fmt.Println("testfunc的值：", testfunc)
	//打印函数值的类型信息
	fmt.Printf("testfunc的值：%T\n", testfunc)
	fmt.Println("testmap的值：", testmap)
	fmt.Println("testinterface的值：", testinterface)
	fmt.Println("testpoint的值：", testpoint)
	fmt.Println("testchan的值：", testchan)
	fmt.Println("testbool的值：", testbool)
	fmt.Println("-----------打印变量的默认值--结束-------")
}
