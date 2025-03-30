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
 * 在函数中修改切片的长度和容量，不会影响到原切片】
 * 在 Go 语言中，切片是对底层数组的引用，但切片本身并不直接存储底层数组的地址。切片是一个结构体，包含以下三个字段：
 * 指针（Pointer）：指向底层数组的起始位置。
 * 长度（Length）：切片的长度。
 * 容量（Capacity）：切片的容量。

 * type slice struct {
 *   ptr *T    // 指向底层数组的指针
 *   len int   // 切片的长度
 *   cap int   // 切片的容量
 * }

切片与底层数组的关系
切片本身并不直接存储底层数组的地址，而是通过指针字段（ptr）指向底层数组的起始位置。
切片的长度和容量字段（len 和 cap）用于描述切片的使用范围和底层数组的可用范围。
*/

import (
	"fmt"
	"unsafe"
)

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
	subarray5 := array5[:3]         // 子切片
	// 通过数组创建切片
	// 这里的切片是一个半闭半开区间，包括第一个元素，但排除最后一个元素
	// 也就是slice5 = [2,3,4]
	fmt.Println("数组array5:", array5)
	fmt.Println("第五个切片slice5:", slice5)
	fmt.Println("子切片subarray5:", subarray5)
	fmt.Println("修改前的array5数组地址:", &array5)
	fmt.Printf("修改前的array5数组地址:%p\n", &array5) // 打印数组array5的地址
	fmt.Println("修改前的array5数组长度:", len(array5))
	fmt.Println("修改前的array5数组容量:", cap(array5))
	fmt.Println("修改前的slice5切片地址:", &slice5)
	fmt.Printf("修改前的slice5切片地址:%p\n", &slice5) // 打印slice的地址
	fmt.Println("修改前的slice5切片长度:", len(slice5))
	fmt.Println("修改前的slice5切片容量:", cap(slice5))
	fmt.Println("修改前的subarray5切片地址:", &subarray5)
	fmt.Printf("修改前的subarray5切片地址:%p\n", &subarray5)                   // 打印slice的地址
	fmt.Printf("修改前的subarray5切片第一个元素的地址就是底层数组的地址:%p\n", &subarray5[0]) // 打印slice的地址

	// Get the address of the first element of the slice
	if len(subarray5) > 0 {
		// ptr := unsafe.Pointer(&subarray5) //代码获取的是切片结构体 subarray5 本身的地址
		// ptr := unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&subarray5))) //这种直接将切片头地址转换为 uintptr 指针并解引用的方法是不安全的
		// header := (*reflect.SliceHeader)(unsafe.Pointer(&subarray5)) //
		arrayPtr := unsafe.SliceData(subarray5)
		// arrayPtr := unsafe.Pointer(header.Data)

		fmt.Printf("subarray5底层数组的地址: %p\n", arrayPtr)
	} else {
		fmt.Println("subarray5是空切片，无法获取底层数组地址")
	}

	fmt.Println("修改前的subarray5切片长度:", len(subarray5))
	fmt.Println("修改前的subarray5切片容量:", cap(subarray5))
	//修改数组创建的切片的元素会影响到原数组
	slice5[0] = 10
	fmt.Println("修改后的切片slice5:", slice5)
	fmt.Println("修改后的数组array5:", array5)
	fmt.Println("slice5切片长度:", len(slice5))
	fmt.Println("slice5切片容量:", cap(slice5))
	fmt.Println("slice5的数据类型:", fmt.Sprintf("%T", slice5)) // 打印slice的数据类型
	fmt.Println("array5的数据类型:", fmt.Sprintf("%T", array5)) // 打印slice的数据类型
	fmt.Println("slice5的地址:", &slice5)
	fmt.Printf("slice5的地址:%p\n", &slice5) // 打印slice的地址
	fmt.Println("array5的地址:", &array5)
	fmt.Printf("array5的地址:%p\n", &array5) // 打印slice的地址

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

	// 在切片中添加元素
	*slice6 = append(*slice6, 1)
	*slice6 = append(*slice6, 2)
	*slice6 = append(*slice6, 3)
	*slice6 = append(*slice6, 4)

	fmt.Println("通过slice6访问第六个切片slice6:", slice6)
	fmt.Println("通过*slice访问第六个切片slice6的值:", *slice6) // 通过指针访问切片的值
	fmt.Println("通过*slice6访问第六个切片slice6的长度:", len(*slice6))
	fmt.Println("通过*slice6访问第六个切片slice6的容量:", cap(*slice6))
	fmt.Println("slice6的数据类型:", fmt.Sprintf("%T", slice6))   // 打印slice的数据类型
	fmt.Println("*slice6的数据类型:", fmt.Sprintf("%T", *slice6)) // 打印slice的数据类型
	fmt.Println("slice6的地址:", &slice6)
	fmt.Printf("slice6的地址:%p\n", &slice6) // 打印slice的地址
	fmt.Println("*slice6的地址:", *slice6)
	fmt.Printf("*slice6的地址:%p\n", *slice6) // 打印slice的地址

	fmt.Println("----------------------")

	// 切片的第七种初始化方式：通过append函数创建
	fmt.Println("-----------第七种初始化方式：通过append函数创建--------------")
	slice7 := []int{} // 创建一个空切片
	// 通过append函数创建切片
	/*
			* Go中的切片扩容策略
			* 当切片的容量不足时，Go会自动扩容切片，扩容策略如下：
			* 如果新申请的容量大于2倍的原切片的容量，则直接使用新申请的容量。
			* 如果新申请的容量小于2倍的原切片的容量，则扩容为原来的2倍。
			* 如果新申请的容量小于原切片的容量，则扩容为原来的1.5倍。
			* 1. 如果切片的长度小于1024，则扩容为原来的2倍。
			* 2. 如果切片的长度大于等于1024，则扩容为原来的1.25倍。
			* 3. 如果切片的长度大于等于1024，则扩容为原来的1.5倍。
			* 4. 如果切片的长度大于等于1024，则扩容为原来的2倍。

			* Go 1.17切片扩容时会进行内存对齐，这个和内存分配策略相关。进行内存对齐之后，新 slice 的容量是要 大于等于老 slice 容量的 2倍或者1.25倍。
			* 当新切片需要的容量cap大于两倍扩容的容量，则直接按照新切片需要的容量扩容；
			* 当原 slice 容量 < 1024 的时候，新 slice 容量变成原来的 2 倍；
			* 当原 slice 容量 > 1024，进入一个循环，每次容量变成原来的1.25倍,直到大于期望容量。

			一、go1.18 之前：

		如果期望容量大于当前容量的两倍就会使用期望容量；
		如果当前切片的长度小于 1024 就会将容量翻倍；
		如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；
			二、go1.18 之后：

		如果期望容量大于当前容量的两倍就会使用期望容量；
		如果当前切片的长度小于阈值（默认 256）就会将容量翻倍；
		如果当前切片的长度大于等于阈值（默认 256），就会每次增加 25% 的容量，基准是 newcap + 3*threshold，直到新容量大于期望容量；

	*/
	fmt.Printf("slice7没增加元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 1)
	fmt.Printf("slice7第一次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 2)
	fmt.Printf("slice7第二次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 3)
	fmt.Printf("slice7第三次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 4)
	fmt.Printf("slice7第四次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 5)
	fmt.Printf("slice7第五次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 6)
	fmt.Printf("slice7第六次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 7)
	fmt.Printf("slice7第七次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 8)
	fmt.Printf("slice7第八次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 9)
	fmt.Printf("slice7第九次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
	slice7 = append(slice7, 10)
	fmt.Printf("slice7第十次增加一个元素，其长度:%d,容量是:%d\n", len(slice7), cap(slice7))
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
