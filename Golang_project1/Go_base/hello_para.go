package Go_base

//通过省略号动态设置多个参数值

import (
	"fmt"
)

func hello_sum(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func Hello_para() {
	fmt.Println(hello_sum(1, 2, 3, 4, 5))
	fmt.Println(hello_sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	hello_slice := []int{1, 2, 3, 4, 5}
	fmt.Println(hello_sum(hello_slice...))

	hello_slice2 := []int{6, 7, 8, 9, 10}
	fmt.Println(hello_sum(hello_slice2...))

	hello_arry := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", hello_arry)
	fmt.Println(hello_sum(hello_arry[:]...))

}
