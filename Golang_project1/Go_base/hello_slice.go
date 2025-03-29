package Go_base

// 每个数组的大小都是固定的。而切片则为数组元素提供了动态大小的、灵活的视角。 在实践中，切片比数组更常用。

// 类型 []T 表示一个元素类型为 T 的切片。.

// 切片通过两个下标来界定，一个下界和一个上界，二者以冒号分隔：

// a[low : high]
// 它会选出一个半闭半开区间，包括第一个元素，但排除最后一个元素。

// 以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：

// a[1:4]

/*
 * 切片的初始化方式
 * 1. 直接初始化
 * 2. 通过make函数创建，指定长度
 * 3. 通过make函数创建，指定长度和容量
 * 4. 通过make函数创建，指定长度为0和容量不为0
 * 5. 通过数组创建切片
 * 6. 通过new函数创建
 * 7. 通过append函数创建
 * 8. 通过copy函数创建
 * 切片是引用类型，传递的是切片的地址
 * 通过切片传递参数，传递的是切片的引用
 * 也就是slice的地址
 * 在函数中修改切片的元素，会影响到原切片
 * 在函数中修改切片的长度和容量，不会影响到原切片
 */

import "fmt"

func replaceSlice(slice []int) {
	slice[0] = 100
	fmt.Println("replaceSlice函数中参数的切片地址是:", &slice)
	fmt.Printf("replaceSlice函数中参数的切片地址是:%p\n", &slice)
	fmt.Println("replaceSlice函数中修改slice后的切片值:", slice)
	fmt.Println("replaceSlice函数中修改slice后的切片长度:", len(slice))
	fmt.Println("replaceSlice函数中修改slice后的切片容量:", cap(slice))
}

func HelloSlice() {
	// 切片的第一种初始化方式： 直接初始化

	fmt.Println("-----------第一种初始化方式：直接初始化--------------")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片slice:", slice)
	fmt.Println("原始切片slice长度:", len(slice))
	fmt.Println("原始切片slice容量:", cap(slice))

	fmt.Println("----------------------")

	//切片的第二种初始化方式：通过make函数创建，指定长度
	fmt.Println("-----------第二种初始化方式：通过make函数创建--------------")
	slice2 := make([]int, 5)
	fmt.Println("第二个切片slice2:", slice2)
	fmt.Println("slice2切片长度:", len(slice2))
	fmt.Println("slice2切片容量:", cap(slice2))
	fmt.Println("----------------------")

	// 切片的第三种初始化方式：通过make函数创建，指定长度和容量
	fmt.Println("-----------第三种初始化方式：通过make函数创建，指定长度和容量--------------")
	slice3 := make([]int, 5, 10) // 第一个参数是长度，第二个参数是容量
	fmt.Println("第三个切片slice3:", slice3)
	fmt.Println("slice3切片长度:", len(slice3))
	fmt.Println("slice3切片容量:", cap(slice3))
	fmt.Println("----------------------")

	// 切片的第四种初始化方式：通过make函数创建，指定长度为0和容量不为0
	fmt.Println("-----------第四种初始化方式：通过make函数创建，指定长度为0和容量不为0--------------")
	slice4 := make([]int, 0, 10) // 第一个参数是长度，第二个参数是容量
	fmt.Println("第四个切片slice4:", slice4)
	fmt.Println("slice4切片长度:", len(slice4))
	fmt.Println("slice4切片容量:", cap(slice4))
	fmt.Println("----------------------")

	// 切片的第五种初始化方式：通过数组创建切片
	fmt.Println("-----------第五种初始化方式：通过数组创建切片--------------")
	array5 := [5]int{1, 2, 3, 4, 5} // 数组
	slice5 := array5[1:4]           // 切片
	// 通过数组创建切片
	// 这里的切片是一个半闭半开区间，包括第一个元素，但排除最后一个元素
	// 也就是slice5 = [2,3,4]
	fmt.Println("数组array5:", array5)
	fmt.Println("第五个切片slice5:", slice5)
	//修改数组创建的切片的元素会影响到原数组
	slice5[0] = 10
	fmt.Println("修改后的切片slice5:", slice5)
	fmt.Println("修改后的数组array5:", array5)

	fmt.Println("----------------------")

	// 切片的第六种初始化方式：通过new函数创建
	fmt.Println("-----------第六种初始化方式：通过new函数创建--------------")
	slice6 := new([]int) // 创建一个切片指针，new函数返回的是一个指向切片的指针
	// 这里的slice6是一个切片指针
	// 通过new函数创建的切片指针，默认值为nil
	fmt.Println("通过slice6访问第六个切片slice6:", slice6)
	fmt.Println("通过*slice访问第六个切片slice6的值:", *slice6) // 通过指针访问切片的值
	fmt.Println("silce6切片长度:", len(*slice6))
	fmt.Println("silce6切片容量:", cap(*slice6))
	fmt.Println("slice的数据类型:", fmt.Sprintf("%T", slice6))   // 打印slice的数据类型
	fmt.Println("*slice的数据类型:", fmt.Sprintf("%T", *slice6)) // 打印slice的数据类型
	fmt.Println("----------------------")

	// 切片的第七种初始化方式：通过append函数创建
	fmt.Println("-----------第七种初始化方式：通过append函数创建--------------")
	slice7 := []int{} // 创建一个空切片
	// 通过append函数创建切片
	slice7 = append(slice7, 1)
	slice7 = append(slice7, 2)
	slice7 = append(slice7, 3)
	slice7 = append(slice7, 4)
	slice7 = append(slice7, 5)
	slice7 = append(slice7, 6)
	slice7 = append(slice7, 7)
	slice7 = append(slice7, 8)
	slice7 = append(slice7, 9)
	slice7 = append(slice7, 10)
	fmt.Println("第七个切片slice7:", slice7)
	fmt.Println("第七个切片slice7长度:", len(slice7))
	fmt.Println("第七个切片slice7容量:", cap(slice7))
	fmt.Println("----------------------")
	// 切片的第八种初始化方式：通过copy函数创建
	fmt.Println("-----------第八种初始化方式：通过copy函数创建--------------")
	slice8 := make([]int, 5) // 创建一个切片
	// 通过copy函数创建切片
	slice8 = append(slice8, 1)
	slice8 = append(slice8, 2)
	slice8 = append(slice8, 3)
	slice8 = append(slice8, 4)
	slice8 = append(slice8, 5)
	slice8 = append(slice8, 6)
	slice8 = append(slice8, 7)
	slice8 = append(slice8, 8)
	slice8 = append(slice8, 9)
	slice8 = append(slice8, 10)
	slice9 := make([]int, 5) // 创建一个切片
	// 通过copy函数创建切片
	copy(slice9, slice8) // 将slice8的值复制到slice9
	fmt.Println("第八个切片slice8:", slice8)
	fmt.Println("第八个切片slice9:", slice9)
	fmt.Println("第八个切片slice8长度:", len(slice8))
	fmt.Println("第八个切片slice8容量:", cap(slice8))
	fmt.Println("第八个切片slice9长度:", len(slice9))
	fmt.Println("第八个切片slice9容量:", cap(slice9))
	fmt.Println("----------------------")

	// 修改切片的元素
	slice[0] = 10
	fmt.Println("修改后的切片:", slice)

	// 获取切片的长度和容量
	fmt.Println("切片长度:", len(slice))
	fmt.Println("切片容量:", cap(slice))

	// 添加元素到切片
	slice = append(slice, 6)
	fmt.Println("添加元素后的切片:", slice)

	// 删除切片中的元素
	slice = append(slice[:1], slice[2:]...) // 删除下标为1的元素
	fmt.Println("删除元素后的切片:", slice)

	// 通过切片传递参数
	fmt.Println("-----------通过切片传递参数--------------")
	// 通过切片传递参数
	// 这里的slice是一个切片
	// 通过切片传递参数，传递的是切片的引用
	// 也就是slice的地址
	// 在函数中修改切片的元素，会影响到原切片
	fmt.Println("在传递参数前原切片变量slice5的值为:", slice5)
	fmt.Println("在传递参数前切片变量slice5的地址为:", &slice5)
	fmt.Printf("在传递参数前切片变量slice5的地址是:%p\n", &slice5)
	fmt.Println("在传递参数前切片变量slice5的长度为:", len(slice5))
	fmt.Println("在传递参数前切片变量slice5的容量为:", cap(slice5))
	replaceSlice(slice5)
	fmt.Println("在传递参数后原切片变量slice5的值为:", slice5)
	fmt.Println("在传递参数后切片变量slice5的地址为:", &slice5)
	fmt.Printf("在传递参数后切片变量slice5的地址是:%p\n", &slice5)
	fmt.Println("在传递参数后切片变量slice5的长度为:", len(slice5))
	fmt.Println("在传递参数后切片变量slice5的容量为:", cap(slice5))

}
