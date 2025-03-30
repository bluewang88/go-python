package Go_base

import (
	"fmt"
)

// go中的map对应python的dict
// map 是一种内置的数据结构，用于存储键值对的集合。
// go语言中的map的key和value在声明的时候就要指定类型

// map 映射将键映射到值。
// 映射的零值为 nil 。nil 映射既没有键，也不能添加键。
// make 函数会返回给定类型的映射，并将其初始化备用。

// map是无序的key-value对的集合
// map是引用类型，传递的是map的地址
// map的key是不可变类型，值是可变类型
// map的key可以是string、int、float、bool、指针、结构体等不可变类型
// map的key类型需要支持==和!=操作符
// map的值可以是任意类型，包括切片、结构体、函数等
// map的key是唯一的，不能重复
// map的值可以重复

// map的初始化方式
// 1. 直接初始化
// 2. 通过make函数创建，指定长度
// 3. 通过make函数创建，指定长度和容量
// 4. 通过make函数创建，指定长度为0和容量不为0
// 5. 通过数组创建map

func HelloMap() {
	// 1. 直接初始化,通过字面值
	fmt.Println("-----------直接初始化,通过字面值--------------")
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("map1:", map1)
	fmt.Println("map1的地址是:", &map1)
	fmt.Printf("map1的地址是:%p\n", &map1)
	fmt.Println("map1的长度是:", len(map1))
	// fmt.Println("map1的容量是:", cap(map1)) // map没有容量的概念
	fmt.Println("map1的值是:", map1["a"])
	fmt.Println("map1的值是:", map1["b"])
	fmt.Println("map1的值是:", map1["c"])
	fmt.Println("map1的值是:", map1["d"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map1的类型是:", fmt.Sprintf("%T", map1))
	fmt.Println("------------------------------")

	// 2. 通过make函数创建，makeh函数可以创建slice、map和channel
	fmt.Println("-----------通过make函数创建map--------------")
	map2 := make(map[string]int)
	map2["x"] = 10
	map2["y"] = 20
	map2["z"] = 30
	fmt.Println("map2:", map2)
	fmt.Println("map2的地址是:", &map2)
	fmt.Printf("map2的地址是:%p\n", &map2)
	fmt.Println("map2的长度是:", len(map2))
	fmt.Println("map2的值是:", map2["x"])
	fmt.Println("map2的值是:", map2["y"])
	fmt.Println("map2的值是:", map2["z"])
	fmt.Println("map2的值是:", map2["a"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map2的类型是:", fmt.Sprintf("%T", map2))
	fmt.Println("------------------------------")

	// 3.定义一个空的map
	fmt.Println("-----------定义一个空的map--------------")
	map3 := make(map[string]int)
	fmt.Println("map3:", map3)
	fmt.Println("map3的地址是:", &map3)
	fmt.Printf("map3的地址是:%p\n", &map3)
	fmt.Println("map3的长度是:", len(map3))
	fmt.Println("map3的值是:", map3["x"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map3的值是:", map3["y"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map3的值是:", map3["z"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map3的值是:", map3["a"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map3的类型是:", fmt.Sprintf("%T", map3))

	// 删除map中的某个元素
	fmt.Println("-----------删除map中的某个元素--------------")
	delete(map3, "x")
	fmt.Println("删除map3中的x元素后map3:", map3)
	fmt.Println("删除map3中的x元素后map3的长度是:", len(map3))
	fmt.Println("删除map3中的x元素后map3的值是:", map3["x"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的x元素后map3的值是:", map3["y"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的x元素后map3的值是:", map3["z"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的x元素后map3的值是:", map3["a"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的x元素后map3的类型是:", fmt.Sprintf("%T", map3))
	// 删除map中的所有元素
	fmt.Println("-----------删除map中的所有元素--------------")
	for k := range map3 {
		delete(map3, k)
	}
	fmt.Println("删除map3中的所有元素后map3:", map3)
	fmt.Println("删除map3中的所有元素后map3的长度是:", len(map3))
	fmt.Println("删除map3中的所有元素后map3的值是:", map3["x"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的所有元素后map3的值是:", map3["y"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的所有元素后map3的值是:", map3["z"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的所有元素后map3的值是:", map3["a"]) // 如果key不存在，返回值类型的零值
	fmt.Println("删除map3中的所有元素后map3的类型是:", fmt.Sprintf("%T", map3))
	// 在map中添加元素
	fmt.Println("-----------在map中添加元素--------------")
	map3["x"] = 10
	map3["y"] = 20
	map3["z"] = 30
	fmt.Println("添加元素后map3:", map3)
	fmt.Println("添加元素后map3的长度是:", len(map3))
	fmt.Println("添加元素后map3的值是:", map3["x"]) // 如果key不存在，返回值类型的零值
	fmt.Println("添加元素后map3的值是:", map3["y"]) // 如果key不存在，返回值类型的零值
	fmt.Println("添加元素后map3的值是:", map3["z"]) // 如果key不存在，返回值类型的零值

}
