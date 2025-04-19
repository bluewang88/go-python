package Go_base_map

import (
	"fmt"
)

func HelloMapSlice() {
	var mapSlice = make([]map[string]string, 3) //创建切片，其中元素为map
	for index, value := range mapSlice {
		fmt.Println("index:", index, "value:", value)
	}
	//对切片中的元素进行初始化

	mapSlice[0] = make(map[string]string, 10) //创建map，初始容量是 10
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["age"] = "18"
	mapSlice[0]["address"] = "沙河"

	for index, value := range mapSlice {
		fmt.Println("index:", index, "value:", value)
		fmt.Printf("index:%d value:%v\n", index, value) //输出为：index:0 value:map[address:沙河 age:18 name:小王子]
		//%d表示十进制整数，%v表示值的默认格式
		for k, v := range value {
			fmt.Printf("key:%s value:%s\n", k, v)
		}
	}

	var mapSlice2 = make(map[string][]string, 3)
	fmt.Println(mapSlice2)
	mapSlice2["name"] = []string{"小王子", "小红"}
	mapSlice2["age"] = []string{"18", "19"}
	mapSlice2["address"] = []string{"沙河", "北京"}
	fmt.Println(mapSlice2)
	for k, v := range mapSlice2 {
		fmt.Printf("key:%s value:%v\n", k, v)
		for _, v1 := range v {
			fmt.Printf("value:%s\n", v1)
		}
	}
	fmt.Println("mapSlice2的地址是:", &mapSlice2)
	fmt.Printf("mapSlice2的地址是:%p\n", &mapSlice2)
}
