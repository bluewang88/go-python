package Go_base

import "fmt"

/*
 * Go中的数组是值类型，而不是引用类型。
 * 这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。
 * 如果对新变量进行了更改，则不会在原始数组中反映。
 * 这与切片不同，切片是引用类型，当它们被分配给一个新变量时，将把原始切片的引用分配给新变量。
 * 这意味着对新变量的更改将反映在原始切片中。
 * 这使得数组在需要保护数据不被意外修改时非常有用。
 */

func HelloArrayValueSemantics() {
	// 定义一个数组
	var arr [5]int

	// 赋值
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// 输出数组
	fmt.Println("原始数组:", arr)

	// 创建一个新变量，并赋值给新变量
	newArr := arr

	// 修改新变量的值
	newArr[0] = 10

	// 输出原始数组和修改后的数组
	fmt.Println("原始数组:", arr)
	fmt.Println("修改后的数组:", newArr)

	// 调用函数，传递数组
	fmt.Println("----------------------")
	fmt.Println("调用函数ChangeArrayFunc前的原始数组:", arr)
	ChangeArrayFunc(arr)
	// 输出原始数组
	fmt.Println("函数调用后的原始数组:", arr)

	fmt.Println("----------------------")
}

//数组作为函数参数的传递方式，其也是值传递

func ChangeArrayFunc(arr [5]int) {
	arr[0] = 10
	fmt.Println("函数内修改后的数组:", arr)
}
