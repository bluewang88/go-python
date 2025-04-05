package objectgotest

//接口的类型断言

// 类型断言是Go语言中的一种机制，用于检查接口类型的值是否实现了某个具体类型的方法
// 类型断言的语法是：value, ok := interfaceValue.(ConcreteType)
// 其中，interfaceValue是一个接口类型的值，ConcreteType是一个具体类型
// value是interfaceValue的具体类型的值，ok是一个布尔值，表示类型断言是否成功
// 如果类型断言成功，ok为true，value为interfaceValue的具体类型的值
// 如果类型断言失败，ok为false，value为ConcreteType的零值
// 类型断言可以用于接口类型的值的转换
// 类型断言可以用于接口类型的值的检查

// 两种断言语法:

// 不安全断言: value := interface.(Type) - 如果失败会引发panic
// 安全断言: value, ok := interface.(Type) - 返回布尔值表示是否成功
// Type Switch:

// 使用 switch value := x.(type) 进行多类型判断
// 比连续的if-else断言更简洁高效
// 接口转换:

// 可以将一个接口类型转换为另一个接口类型（如果实现了相应方法）
// 例如: var obj Object = shape.(Object)
// 空接口断言:

// interface{} 可以存储任何类型的值
// 使用类型断言提取具体类型
// 注意事项:

// 只能对接口类型进行断言
// 断言为不可能的类型会导致panic（如果不使用ok模式）
// 应该优先使用带ok的安全断言形式

import (
	"fmt"
	"math"
)

// 定义几个接口
type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

// 定义具体类型
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Cube struct {
	Side float64
}

func (c Cube) Area() float64 {
	return 6 * c.Side * c.Side
}

func (c Cube) Volume() float64 {
	return c.Side * c.Side * c.Side
}

func HelloInterfaceTypeAssertion() {
	// 创建一个包含不同Shape实现的切片
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Cube{Side: 2},
	}

	fmt.Println("===== 1. 基本类型断言 =====")
	// 遍历并使用类型断言
	for i, shape := range shapes {
		fmt.Printf("形状 #%d: 面积 = %.2f\n", i, shape.Area())

		// 类型断言方式1：无检查（不安全）
		// circle := shape.(Circle) // 如果断言失败会引发panic
		// fmt.Printf("  作为圆形: 半径 = %.2f\n", circle.Radius)

		// 类型断言方式2：带检查（安全）
		if circle, ok := shape.(Circle); ok {
			fmt.Printf("  是圆形: 半径 = %.2f\n", circle.Radius)
		} else if rectangle, ok := shape.(Rectangle); ok {
			fmt.Printf("  是矩形: 宽 = %.2f, 高 = %.2f\n", rectangle.Width, rectangle.Height)
		} else if cube, ok := shape.(Cube); ok {
			fmt.Printf("  是立方体: 边长 = %.2f, 体积 = %.2f\n", cube.Side, cube.Volume())
		}
	}

	fmt.Println("\n===== 2. 类型断言用于接口转换 =====")
	// 检查Shape是否也实现了Object接口
	for _, shape := range shapes {
		// 尝试将Shape转换为Object接口
		if obj, ok := shape.(Object); ok {
			fmt.Printf("形状也是一个物体，体积 = %.2f\n", obj.Volume())
		} else {
			fmt.Printf("形状不是一个物体（未实现Volume方法）\n")
		}
	}

	fmt.Println("\n===== 3. Type Switch类型选择 =====")
	// 使用type switch进行类型判断
	for i, shape := range shapes {
		fmt.Printf("形状 #%d: ", i)

		// 使用type switch替代多个if语句
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("是圆形，半径 = %.2f\n", s.Radius)
		case Rectangle:
			fmt.Printf("是矩形，面积 = %.2f\n", s.Area())
		case Cube:
			fmt.Printf("是立方体，表面积 = %.2f, 体积 = %.2f\n", s.Area(), s.Volume())
		default:
			fmt.Println("未知类型")
		}
	}

	fmt.Println("\n===== 4. 空接口和任意类型断言 =====")
	// 使用空接口存储任意类型
	var anyValues []interface{} = []interface{}{
		42,
		"Hello, World",
		true,
		Circle{Radius: 3},
		[]int{1, 2, 3},
		map[string]int{"one": 1, "two": 2},
	}

	for i, value := range anyValues {
		fmt.Printf("值 #%d: ", i)

		// 对空接口进行类型断言
		switch v := value.(type) {
		case int:
			fmt.Printf("整数 %d\n", v)
		case string:
			fmt.Printf("字符串 \"%s\"\n", v)
		case bool:
			fmt.Printf("布尔值 %t\n", v)
		case Circle:
			fmt.Printf("圆形，面积 = %.2f\n", v.Area())
		case []int:
			fmt.Printf("整数切片，长度 = %d\n", len(v))
		case map[string]int:
			fmt.Printf("字符串到整数的映射，大小 = %d\n", len(v))
		default:
			fmt.Printf("未知类型 %T\n", v)
		}
	}

	fmt.Println("\n===== 5. 类型断言的常见错误 =====")
	// 错误示例1：对非接口类型使用类型断言
	// var num int = 10
	// _ = num.(string) // 编译错误：不能对非接口类型使用类型断言

	// 错误示例2：断言为不可能的类型（会导致panic）
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("  捕获到panic: 类型断言失败")
			}
		}()

		var s Shape = Circle{Radius: 1}
		// 尝试断言为一个不可能的类型（没有任何实现关系）
		// _ = s.(string) // 这会导致panic
		fmt.Println("  这一行不会执行", s)
	}()
}
