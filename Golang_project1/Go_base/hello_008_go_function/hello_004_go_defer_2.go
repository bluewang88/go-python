package Go_Base_func

import (
	"fmt"
)

func calc3(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func HelloDefer2() {
	x := 1
	y := 2
	defer calc3("AA", x, calc3("A", x, y))
	x = 10
	defer calc3("BB", x, calc3("B", x, y))
	y = 20
}

// 输出结果:
// A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 1 3 4
