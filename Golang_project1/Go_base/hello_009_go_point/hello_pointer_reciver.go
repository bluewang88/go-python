package Go_base_point

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X, Y float64
}

// typeof returns the type of v as a string
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
func (v Vertex) Abs() float64 {
	//使用数学公式 √(x² + y²) 计算模
	// return math.Sqrt(v.X*v.X + v.Y*v.Y) //计算v的模

	return v.X + v.Y

}

func (v *Vertex) ScalePoint(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleValue(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerReceiverDemo() {
	fmt.Println("----------------值变量调用指针Receiver方法----------------")
	v := Vertex{3, 4}                                        //创建一个Vertex类型的变量v,结构体实例化
	fmt.Println("结构体Vertex的变量v的值是：", v, "他的类型是：", typeof(v)) // Print the value and type of the Vertex structure variable v
	v.ScalePoint(10)
	fmt.Println("调用func (v *Vertex) ScalePoint(f float64) 方法后，结构体Vertex的变量v中x的值是：", v.X, "y的值是：", v.Y) // Print the value and type of the Vertex structure variable v
	fmt.Println("--------------------------------")

	fmt.Println("----------------指针变量调用指针Receiver方法----------------")
	vpointer := &Vertex{3, 4}                                                     //创建一个Vertex类型的指针vpointer,结构体实例化
	fmt.Println("结构体Vertex的指针vpointer的值是：", vpointer, "他的类型是：", typeof(vpointer)) // Print the value and type of the Vertex structure pointer vpointer
	vpointer.ScalePoint(10)
	fmt.Println("调用func (v *Vertex) ScalePoint(f float64) 方法后，结构体Vertex的指针vpointer中x的值是：", vpointer.X, "y的值是：", vpointer.Y) // Print the value and type of the Vertex structure pointer vpointer
	fmt.Println("--------------------------------")
}

// ➜  Golang_project1 git:(master) ✗ go run hello_main.go
// 结构体Vertex的变量v的值是： {3 4} 他的类型是： Go_base.Vertex
// 调用func (v *Vertex) ScalePoint(f float64) 方法后，结构体Vertex的变量v中x的值是： 30 y的值是： 40
// ➜  Golang_project1 git:(master) ✗ go run hello_main.go
// 结构体Vertex的变量v的值是： {3 4} 他的类型是： Go_base.Vertex
// 调用func (v *Vertex) ScalePoint(f float64) 方法后，结构体Vertex的变量v中x的值是： 30 y的值是： 40
// 结构体Vertex的指针vpointer的值是： &{3 4} 他的类型是： *Go_base.Vertex
// 调用func (v *Vertex) ScalePoint(f float64) 方法后，结构体Vertex的指针vpointer中x的值是： 30 y的值是： 40

func ValueReceiverDemo() {
	fmt.Println("----------------值变量调用值Receiver方法----------------")
	v := Vertex{3, 4}                                        //创建一个Vertex类型的变量v,结构体实例化
	fmt.Println("结构体Vertex的变量v的值是：", v, "他的类型是：", typeof(v)) // Print the value and type of the Vertex structure variable v
	v.ScaleValue(10)
	fmt.Println("调用func (v Vertex) ScaleValue(f float64) 方法后，结构体Vertex的变量v中x的值是：", v.X, "y的值是：", v.Y) // Print the value and type of the Vertex structure variable v
	fmt.Println("--------------------------------")

	fmt.Println("----------------指针变量调用值Receiver方法----------------")
	vpointer := &v
	fmt.Println("结构体Vertex的指针vpointer的值是：", vpointer, "他的类型是：", typeof(vpointer)) // Print the value and type of the Vertex structure pointer vpointer
	vpointer.ScaleValue(10)
	fmt.Println("调用func (v Vertex) ScaleValue(f float64) 方法后，结构体Vertex的指针vpointer中x的值是：", vpointer.X, "y的值是：", vpointer.Y) // Print the value and type of the Vertex structure pointer vpointer
	fmt.Println("原始的结构体Vertex的变量v的值是：", v, "他的类型是：", typeof(v))                                                            // Print the value and type of the Vertex structure variable v
	fmt.Println("结构体Vertex的指针vpointer所指向变量v中x的值是：", v.X, "y的值是：", v.Y)
	fmt.Println("--------------------------------")
}

//----------------------------------------
// 值接收者方法（如 func (v Vertex) ScaleValue(f float64)）会接收一个结构体的副本，而不是原始结构体。
// 当通过指针变量调用值接收者方法时，Go语言会自动解引用指针，并将指针所指向的变量的副本传递给方法。
//----------------------------------------

// ➜  Golang_project1 git:(master) ✗ go run hello_main.go
// ----------------值变量调用指针Receiver方法----------------
// 结构体Vertex的变量v的值是： {3 4} 他的类型是： Go_base.Vertex
// 调用func (v *Vertex) ScalePoint(f float64) 方法后，结构体Vertex的变量v中x的值是： 30 y的值是： 40
// --------------------------------
// ----------------指针变量调用指针Receiver方法----------------
// 结构体Vertex的指针vpointer的值是： &{3 4} 他的类型是： *Go_base.Vertex
// 调用func (v *Vertex) ScalePoint(f float64) 方法后，结构体Vertex的指针vpointer中x的值是： 30 y的值是： 40
// --------------------------------
// ----------------值变量调用值Receiver方法----------------
// 结构体Vertex的变量v的值是： {3 4} 他的类型是： Go_base.Vertex
// 调用func (v Vertex) ScaleValue(f float64) 方法后，结构体Vertex的变量v中x的值是： 3 y的值是： 4
// --------------------------------
// ----------------指针变量调用值Receiver方法----------------
// 结构体Vertex的指针vpointer的值是： &{3 4} 他的类型是： *Go_base.Vertex
// 调用func (v Vertex) ScaleValue(f float64) 方法后，结构体Vertex的指针vpointer中x的值是： 3 y的值是： 4
// 原始的结构体Vertex的变量v的值是： {3 4} 他的类型是： Go_base.Vertex
// 结构体Vertex的指针vpointer所指向变量v中x的值是： 3 y的值是： 4
// --------------------------------
