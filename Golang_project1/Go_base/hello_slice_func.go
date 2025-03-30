package Go_base

import "fmt"

// 定义了一个函数类型
type MathFunc func(float64, float64) float64

func HelloSliceFunc() {
	// 定义一个存储函数的slice
	funcs := []MathFunc{
		func(a, b float64) float64 {
			fmt.Printf("%f add %f = ", a, b)
			return a + b
		},
		func(a, b float64) float64 {
			fmt.Printf("%f subtract %f = ", a, b)
			return a - b
		},
		func(a, b float64) float64 {
			fmt.Printf("%f multiply %f = ", a, b)
			return a * b
		},
	}

	// 遍历并调用slice中的函数
	for _, f := range funcs {
		result := f(2, 3)
		fmt.Println(result)
	}
}
