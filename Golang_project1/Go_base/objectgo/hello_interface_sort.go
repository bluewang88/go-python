package objectgotest

import (
	"fmt"
	"sort"
)

// 1. 使用内置sort包和sort.Interface接口实现排序
// sort.Interface接口要求实现三个方法：
// - Len() int：返回集合长度
// - Less(i, j int) bool：比较第i和第j个元素
// - Swap(i, j int)：交换第i和第j个元素

// 示例1：自定义整数切片排序
type IntSlice []int

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j] // 升序排序
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// 示例2：对结构体切片按照不同字段排序
type SortableStudent struct {
	Name  string
	Age   int
	Score float64
}

// 按年龄排序
type ByAge []SortableStudent

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// 按分数排序（降序）
type ByScore []SortableStudent

func (s ByScore) Len() int           { return len(s) }
func (s ByScore) Less(i, j int) bool { return s[i].Score > s[j].Score } // 注意这里是降序
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// 示例3：实现多字段排序 - 先按名字，再按年龄
type ByNameThenAge []SortableStudent

func (a ByNameThenAge) Len() int      { return len(a) }
func (a ByNameThenAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByNameThenAge) Less(i, j int) bool {
	if a[i].Name != a[j].Name {
		return a[i].Name < a[j].Name
	}
	return a[i].Age < a[j].Age
}

// 对外暴露的函数，展示接口排序的用法
func HelloInterfaceSort() {
	fmt.Println("===== 1. 整数切片排序 =====")
	nums := IntSlice{5, 2, 6, 3, 1, 4}
	fmt.Println("排序前:", nums)
	sort.Sort(nums)
	fmt.Println("排序后:", nums)

	// 逆序排序
	sort.Sort(sort.Reverse(nums))
	fmt.Println("逆序排序:", nums)

	fmt.Println("\n===== 2. 结构体切片按不同字段排序 =====")
	students := []SortableStudent{
		{"张三", 20, 85.5},
		{"李四", 18, 92.0},
		{"王五", 22, 78.5},
		{"赵六", 19, 88.0},
		{"钱七", 21, 90.5},
	}

	// 按年龄排序
	fmt.Println("原始学生列表:")
	printStudents(students)

	sort.Sort(ByAge(students))
	fmt.Println("\n按年龄排序（升序）:")
	printStudents(students)

	// 按分数排序（降序）
	sort.Sort(ByScore(students))
	fmt.Println("\n按分数排序（降序）:")
	printStudents(students)

	// 多字段排序
	// 添加一些同名学生进行测试
	students = append(students, SortableStudent{"张三", 19, 76.5})
	students = append(students, SortableStudent{"李四", 22, 81.0})

	sort.Sort(ByNameThenAge(students))
	fmt.Println("\n按姓名+年龄复合排序:")
	printStudents(students)

	fmt.Println("\n===== 3. 使用Go 1.8+的便捷排序函数 =====")
	// 使用sort.Slice函数，可以无需定义类型
	people := []struct {
		Name string
		Age  int
	}{
		{"张三", 35},
		{"李四", 20},
		{"王五", 45},
		{"赵六", 25},
	}

	// 直接使用匿名函数定义排序条件
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("按年龄排序的人员列表:")
	for _, p := range people {
		fmt.Printf("  %s: %d岁\n", p.Name, p.Age)
	}
}

// 辅助函数，打印学生列表
func printStudents(students []SortableStudent) {
	for _, s := range students {
		fmt.Printf("  姓名: %-4s 年龄: %2d 分数: %.1f\n", s.Name, s.Age, s.Score)
	}
}
