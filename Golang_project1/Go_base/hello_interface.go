package Go_base

import (
	"fmt"
	"math"
)

// Abser 是一个接口，定义了 Abs 方法，返回一个 float64 类型的值。
type Abser interface {
	Abs() float64
}

// MyFloat 是一个自定义的 float64 类型，实现了 Abser 接口。
type MyFloat float64

// Abs 方法返回 MyFloat 类型的绝对值。
// 如果 MyFloat 为负数，则返回其相反数，否则返回其本身。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// VertexInterface 是一个结构体，包含两个 float64 类型的字段 X 和 Y。
type VertexInterface struct {
	X, Y float64
}

// Abs 方法返回 VertexInterface 类型的欧几里得范数（即向量的长度）。
// 该方法通过计算 X 和 Y 的平方和的平方根来实现。
func (v *VertexInterface) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// InterfaceDemo 演示了接口的使用方式，展示了如何通过接口调用不同类型的 Abs 方法。
// 该函数首先定义了一个 Abser 接口类型的变量 a，然后分别将 MyFloat 类型和 *Vertex 类型的值赋给 a，
// 最后尝试将 Vertex 类型的值赋给 a，但由于 Vertex 类型没有实现 Abser 接口，因此会报错。
func InterfaceDemo() {
	var a Abser
	f := MyFloat(-math.Sqrt2)  // -1.4142135623730951
	v := VertexInterface{3, 4} // {3, 4}

	// 将 MyFloat 类型的值赋给 a，因为 MyFloat 实现了 Abser 接口
	a = f

	// 将 *VertexInterface 类型的值赋给 a，因为 *VertexInterface 实现了 Abser 接口
	a = &v

	// 尝试将 VertexInterface 类型的值赋给 a，但由于 VertexInterface 类型没有实现 Abser 接口，因此会报错
	// a = v // This line will cause a compile error

	// 调用 Abs 方法，因为 a 是一个实现了 Abser 接口的变量
	fmt.Println(a.Abs())
}
