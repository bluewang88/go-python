package Go_base

/*
* 切片就像数组的引用 切片并不存储任何数据，它只是描述了底层数组中的一段。
* 切片的底层数组是动态分配的，并且可以动态扩展。更改切片的元素会修改其底层数组中对应的元素。
* 和它共享底层数组的切片都会观测到这些修改。
 */

import "fmt"

func HelloSlicePointers() {
	names := [4]string{ // 定义一个数组
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // 切片，包含数组的前两个元素,a 的值是对 names 数组的引用
	fmt.Println("Slice a:", a)
	b := names[1:3] // 切片，包含数组的第二和第三个元素
	fmt.Println("Slice b:", b)
	fmt.Println(a, b)

	b[0] = "XXX" // 修改 b 的第一个元素
	fmt.Println("After b[0] = \"XXX\"")
	fmt.Println(a, b)
	fmt.Println(names)

	//给切片a添加元素
	fmt.Println("-----------给切片a添加元素--------------")
	a = append(a, "sjwl")
	fmt.Println("After appending to slice a:", a)
	fmt.Println("After appending to slice b:", b)
	fmt.Println("names:", names)

	fmt.Println("-----------给切片b添加元素--------------")

	b = append(b, "YYY")
	fmt.Println("After appending to slice a:", a)
	fmt.Println("After appending to slice b:", b)
	fmt.Println("names:", names)
	// 通过切片a和b的地址来查看切片的地址

	fmt.Println("-----------通过切片a和b的地址来查看切片的地址--------------")
	fmt.Println("Slice a address:", &a)
	fmt.Printf("Slice a address:%p\n", &a)
	fmt.Println("Slice b address:", &b)
	fmt.Printf("Slice b address:%p\n", &b)

	fmt.Println("-----------通过切片a和b的地址来查看切片的值--------------")
	// 通过切片a和b的地址来查看切片的值
	fmt.Println("Slice a value:", a)
	fmt.Println("Slice b value:", b)
}
