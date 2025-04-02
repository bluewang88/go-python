package Go_base

import (
	"fmt"
)

// HelloSliceCapGrow 演示切片的容量和增长
func HelloSliceCapGrow() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			//%-4d - 左对齐，宽度为4的十进制整数（用户特别关注的部分）
			oldCap = newCap
		}
	}
}
