package Go_base

import (
	"fmt"
)

func Hello_function() {
	fmt.Println("Hello, World!")
}

func Hello_function_args(a, b string, floatargs float64) (int, int, float64) {
	fmt.Println(a, b, floatargs)
	return 0, 1, 2.0
}

func Hello_function_args_return(a, b string, floatargs float64) (sum int, count int, count2 float64) {
	fmt.Println(a, b, floatargs)
	sum = 0
	count = 1
	count2 = floatargs
	return sum, count, count2
}

func Hello_function_args_return2(a, b string, floatargs float64) (sum int, count int, count2 float64) {
	fmt.Println(a, b, floatargs)
	sum = 0
	count = 1
	count2 = floatargs
	return

}
