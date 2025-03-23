package Go_base

//通过省略号动态设置多个参数值

import (
	"fmt"
)
funcc sum(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func Hello_para() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}