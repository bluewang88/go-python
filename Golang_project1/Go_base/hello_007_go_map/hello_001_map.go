package Go_base_map

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
	fmt.Println("-----------1.直接初始化,通过字面值--------------")
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
	fmt.Println("-----------2.通过make函数创建map--------------")
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
	fmt.Println("-----------3.定义一个空的map--------------")
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
	fmt.Println("------------------------------")

	// 	map 容量与长度的区别
	// 容量而非长度：10 这个参数指定的是 map 的初始容量，而不是长度。
	// 不限制大小：与切片不同，map 指定初始容量后，并不会限制 map 可以存储的元素数量。当元素数量超过初始容量时，map 会自动扩容。
	// 性能优化：指定合理的初始容量是一种性能优化手段，可以减少 map 在增长过程中的重新哈希操作。

	// 4.通过 make 创建map，并同时指定长度
	fmt.Println("-----------4.通过 make 创建map，并同时指定长度--------------")
	map4 := make(map[string]int, 10) // 创建一个长度为10的map
	fmt.Println("map4:", map4)
	fmt.Println("map4的地址是:", &map4)
	fmt.Printf("map4的地址是:%p\n", &map4)
	fmt.Println("map4的长度是:", len(map4))
	fmt.Println("map4的值是:", map4["x"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map4的值是:", map4["y"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map4的值是:", map4["z"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map4的值是:", map4["a"]) // 如果key不存在，返回值类型的零值
	fmt.Println("map4的类型是:", fmt.Sprintf("%T", map4))
	fmt.Println("------------------------------")

	// 5.通过数组创建map
	fmt.Println("-----------5.通过数组创建map--------------")
	var arr = [3]string{"a", "b", "c"}
	map5 := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		map5[arr[i]] = i + 1
	}
	fmt.Println("map5:", map5)
	fmt.Println("map5的地址是:", &map5)
	fmt.Printf("map5的地址是:%p\n", &map5)
	fmt.Println("map5的长度是:", len(map5)) // map没有容量的概念
	fmt.Println("map5的值是:", map5["a"])  // 如果key不存在，返回值类型的零值
	fmt.Println("map5的值是:", map5["b"])  // 如果key不存在，返回值类型的零值
	fmt.Println("map5的值是:", map5["c"])  // 如果key不存在，返回值类型的零值
	fmt.Println("map5的值是:", map5["d"])  // 如果key不存在，返回值类型的零值
	fmt.Println("map5的类型是:", fmt.Sprintf("%T", map5))
	fmt.Println("------------------------------")

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
