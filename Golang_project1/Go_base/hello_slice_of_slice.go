package Go_base

//切片可以包含任何类型，当然也包括其他切片。

//切片的切片是一个二维切片，类似于数组的数组。
//切片的切片可以用来表示矩阵、表格等数据结构。

import "fmt"

func HelloSliceOfSlice() {
	//创建一个二维切片
	//已经在外层指定了切片的类型为 [][]int 时，内部元素的类型 []int 就不需要重复指定了，
	//因为 Go 编译器可以从外层声明推断出内层元素类型。
	// sliceOfSlice := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }

	//创建一个二维切片
	sliceOfSlice2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	//创建一个三维切片
	sliceOfSlice3 := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{10, 11, 12},
			{13, 14, 15},
			{16, 17, 18},
		},
		{
			{19, 20, 21},
			{22, 23, 24},
			{25, 26, 27},
		},
	}

	//打印二维切片
	fmt.Println("---------------打印二维切片-------------------")
	// fmt.Println("sliceOfSlice:", sliceOfSlice)
	fmt.Println("sliceOfSlice2:", sliceOfSlice2)

	//打印二维切片的元素
	fmt.Println("sliceOfSlice2[0]:", sliceOfSlice2[0])
	fmt.Println("sliceOfSlice2[1]:", sliceOfSlice2[1])
	fmt.Println("sliceOfSlice2[2]:", sliceOfSlice2[2])
	fmt.Println("sliceOfSlice2[0][0]:", sliceOfSlice2[0][0])
	fmt.Println("sliceOfSlice2[0][1]:", sliceOfSlice2[0][1])
	fmt.Println("sliceOfSlice2[0][2]:", sliceOfSlice2[0][2])
	fmt.Println("sliceOfSlice2[1][0]:", sliceOfSlice2[1][0])
	fmt.Println("sliceOfSlice2[1][1]:", sliceOfSlice2[1][1])
	fmt.Println("sliceOfSlice2[1][2]:", sliceOfSlice2[1][2])
	fmt.Printf("siliceOfSlice2的类型是：%T\n", sliceOfSlice2)
	fmt.Printf("siliceOfSlice2[0]的类型是：%T\n", sliceOfSlice2[0])
	fmt.Printf("siliceOfSlice2[0][0]的类型是：%T\n", sliceOfSlice2[0][0])
	fmt.Println("----------------------")

	fmt.Println("---------------打印三维切片-------------------")
	fmt.Println("sliceOfSlice3:", sliceOfSlice3)
	fmt.Println("sliceOfSlice3[0]:", sliceOfSlice3[0])
	fmt.Println("sliceOfSlice3[1]:", sliceOfSlice3[1])
	fmt.Println("sliceOfSlice3[2]:", sliceOfSlice3[2])
	fmt.Println("sliceOfSlice3[0][0]:", sliceOfSlice3[0][0])
	fmt.Println("sliceOfSlice3[0][1]:", sliceOfSlice3[0][1])
	fmt.Println("sliceOfSlice3[0][2]:", sliceOfSlice3[0][2])
	fmt.Println("sliceOfSlice3[1][0]:", sliceOfSlice3[1][0])
	fmt.Println("sliceOfSlice3[1][1]:", sliceOfSlice3[1][1])
	fmt.Println("sliceOfSlice3[1][2]:", sliceOfSlice3[1][2])
	fmt.Println("sliceOfSlice3[2][0]:", sliceOfSlice3[2][0])
	fmt.Println("sliceOfSlice3[2][1]:", sliceOfSlice3[2][1])
	fmt.Println("sliceOfSlice3[2][2]:", sliceOfSlice3[2][2])
	fmt.Printf("siliceOfSlice3的类型是：%T\n", sliceOfSlice3)
	fmt.Printf("siliceOfSlice3[0]的类型是：%T\n", sliceOfSlice3[0])
	fmt.Printf("siliceOfSlice3[0][0]的类型是：%T\n", sliceOfSlice3[0][0])
	fmt.Printf("siliceOfSlice3[0][0][0]的类型是：%T\n", sliceOfSlice3[0][0][0])
	fmt.Println("----------------------")
}
