package Go_base

//slice删除元素
import "fmt"

// 删除切片中指定索引的元素
func deleteSliceElement(slice []int, index int) []int {
	// 检查索引是否在范围内
	if index < 0 || index >= len(slice) {
		fmt.Println("索引超出范围")
		return slice
	}

	// 删除指定索引的元素
	slice = append(slice[:index], slice[index+1:]...)

	return slice
}

// 删除切片中指定值的元素
func deleteSliceValue(slice []int, value int) []int {
	// 遍历切片，找到指定值的索引
	for i, v := range slice {
		if v == value {
			// 删除指定值的元素
			slice = append(slice[:i], slice[i+1:]...)
			break // 只删除第一个匹配的元素
		}
	}

	return slice
}

// 删除切片中所有指定值的元素
func deleteAllSliceValue(slice []int, value int) []int {
	// 遍历切片，找到指定值的索引
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			// 删除指定值的元素
			slice = append(slice[:i], slice[i+1:]...)
			i-- // 调整索引以处理删除后的元素
		}
	}

	return slice
}

// 删除切片中指定范围的元素
func deleteSliceRange(slice []int, start, end int) []int {
	// 检查范围是否在切片长度内
	if start < 0 || end >= len(slice) || start > end {
		fmt.Println("范围超出切片长度")
		return slice
	}

	// 删除指定范围的元素
	slice = append(slice[:start], slice[end+1:]...)

	return slice
}

// 删除切片中所有指定范围的元素
func deleteAllSliceRange(slice []int, start, end int) []int {
	// 检查范围是否在切片长度内
	if start < 0 || end >= len(slice) || start > end {
		fmt.Println("范围超出切片长度")
		return slice
	}

	// 删除指定范围的元素
	slice = append(slice[:start], slice[end+1:]...)

	return slice
}

// 删除切片中所有指定范围的元素

func HelloSliceDeleteElem() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 11, 12, 13, 14, 15}
	fmt.Println("原始切片:", slice)
	fmt.Println("原始切片长度:", len(slice))
	fmt.Println("原始切片容量:", cap(slice))
	// 删除索引为2的元素
	slice = deleteSliceElement(slice, 2)
	fmt.Println("删除索引为2的元素:", slice)
	fmt.Println("删除索引为2的元素后切片长度:", len(slice))
	fmt.Println("删除索引为2的元素后切片容量:", cap(slice))

	// 删除值为5的元素
	slice = deleteSliceValue(slice, 5)
	fmt.Println("删除值为5的元素:", slice)
	fmt.Println("删除值为5的元素后切片长度:", len(slice))
	fmt.Println("删除值为5的元素后切片容量:", cap(slice))

	// 删除所有值为5的元素
	slice = deleteAllSliceValue(slice, 5)
	fmt.Println("删除所有值为5的元素:", slice)
	fmt.Println("删除所有值为5的元素后切片长度:", len(slice))
	fmt.Println("删除所有值为5的元素后切片容量:", cap(slice))

	// 删除索引为3到6的元素
	slice = deleteSliceRange(slice, 3, 6)
	fmt.Println("删除索引为3到6的元素:", slice)
	fmt.Println("删除索引为3到6的元素后切片长度:", len(slice))
	fmt.Println("删除索引为3到6的元素后切片容量:", cap(slice))

	// 删除所有索引为3到6的元素
	slice = deleteAllSliceRange(slice, 3, 6)
	fmt.Println("删除所有索引为3到6的元素:", slice)
	fmt.Println("删除所有索引为3到6的元素后切片长度:", len(slice))
	fmt.Println("删除所有索引为3到6的元素后切片容量:", cap(slice))

	// 删除索引为0到2的元素
	slice = deleteAllSliceRange(slice, 0, 2)
	fmt.Println("删除索引为0到2的元素:", slice)
	fmt.Println("删除索引为0到2的元素后切片长度:", len(slice))
	fmt.Println("删除索引为0到2的元素后切片容量:", cap(slice))

	// 删除索引为0到1的元素
	slice = deleteAllSliceRange(slice, 0, 1)

	fmt.Println("删除索引为0到1的元素:", slice)
	fmt.Println("删除索引为0到1的元素后切片长度:", len(slice))
	fmt.Println("删除索引为0到1的元素后切片容量:", cap(slice))

}
