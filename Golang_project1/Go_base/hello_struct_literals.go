package Go_base

// 结构体字面量
// 使用 Name: 语法可以仅列出部分字段（字段名的顺序无关）。

// 特殊的前缀 & 返回一个指向结构体的指针。

// 结构体字面量可以在函数调用时使用，
// 也可以在函数返回时使用。

type VertexLiteral struct {
	X int
	Y int
}

var (
	v1 = VertexLiteral{1, 2}       // X: 1, Y: 2
	v2 = VertexLiteral{X: 1}       // Y:0 被隐式地赋予零值
	v3 = VertexLiteral{}           // X: 0, Y: 0
	v4 = &VertexLiteral{1, 2}      // X: 1, Y: 2 创建一个指向结构体的指针
	v5 = &VertexLiteral{X: 1}      // Y: 0
	v6 = &VertexLiteral{}          // X: 0, Y: 0
	v7 = VertexLiteral{X: 1, Y: 2} // X: 1, Y: 2
	v8 = VertexLiteral{X: 1}       // Y: 0
	v9 = VertexLiteral{}           // X: 0, Y: 0
)

func HelloStructLiterals() {
	// 结构体字面量
	// 使用 Name: 语法可以仅列出部分字段（字段名的顺序无关）。
	// 特殊的前缀 & 返回一个指向结构体的指针。
	// 结构体字面量可以在函数调用时使用，
	// 也可以在函数返回时使用。
	v1 := VertexLiteral{1, 2}       // X: 1, Y: 2
	v2 := VertexLiteral{X: 1}       // Y:0 被隐式地赋予零值
	v3 := VertexLiteral{}           // X: 0, Y: 0
	v4 := &VertexLiteral{1, 2}      // X: 1, Y: 2 创建	一个指向结构体的指针
	v5 := &VertexLiteral{X: 1}      // Y: 0
	v6 := &VertexLiteral{}          // X: 0, Y: 0
	v7 := VertexLiteral{X: 1, Y: 2} // X: 1, Y: 2
	v8 := VertexLiteral{X: 1}       // Y: 0
	v9 := VertexLiteral{}           // X: 0, Y: 0
	println(v1.X, v1.Y)             // 1 2
	println(v2.X, v2.Y)             // 1 0
	println(v3.X, v3.Y)             // 0 0
	println(v4.X, v4.Y)             // 1 2
	println(v5.X, v5.Y)             // 1 0
	println(v6.X, v6.Y)             // 0 0
	println(v7.X, v7.Y)             // 1 2
	println(v8.X, v8.Y)             // 1 0
	println(v9.X, v9.Y)             // 0 0
	println(v1)                     // {1 2}
	println(v2)                     // {1 0}
	println(v3)                     // {0 0}
	println(v4)                     // &{1 2}
	println(v5)                     // &{1 0}
	println(v6)                     // &{0 0}
	println(v7)                     // {1 2}
	println(v8)                     // {1 0}
	println(v9)                     // {0 0}
}
