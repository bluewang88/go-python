package Go_base

type VertexTest struct {
	X int
	Y int
}

func HelloStructPointer() {
	// 结构体指针
	v := VertexTest{1, 2}
	p := &v
	p.X = 1e9    // 通过指针修改结构体的值
	println(v.X) // 直接访问结构体的值
	println(p.X) // 通过指针访问结构体的值
	// 结构体指针可以直接使用点号访问字段
	println((*p).X) // 通过指针访问结构体的值
	(*p).X = 1e9
	println(v.X)
	println(p.X)
	println((*p).X)
	// 结构体指针的使用
	// 结构体指针可以直接使用点号访问字段

}
