package Go_base

import "fmt"

type VertexTest struct {
	X int
	Y int
}

func HelloStructPointer() {
	// 结构体指针
	v := VertexTest{1, 2}
	p := &v
	p.X = 1e9        // 通过指针修改结构体的值
	fmt.Println(v.X) // 直接访问结构体的值
	fmt.Println(p.X) // 通过指针访问结构体的值
	// 结构体指针可以直接使用点号访问字段
	fmt.Println((*p).X) // 通过指针访问结构体的值
	(*p).X = 1e9
	fmt.Println(v.X)
	fmt.Println(p.X)
	fmt.Println(v.Y)
	fmt.Println((*p).X)
	// 结构体指针的使用
	// 结构体指针可以直接使用点号访问字段

}
