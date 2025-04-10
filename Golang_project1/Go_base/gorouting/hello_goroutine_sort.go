package gorouting_test

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 从命令行获取一组数字，对其排序
func HelloGoroutineSort() {
	fmt.Println("Hello, Goroutine Sort!Please input a group of numbers to sort:")
	//获取用户输入，并存入切片s []int
	var input string
	fmt.Println("Please input a group of numbers to sort (e.g., 3,1,4,1,5):")
	fmt.Scanln(&input)
	s := strings.Split(input, ",")

	var numbers []int
	for _, str := range s {
		n, err := strconv.Atoi(str)
		if err == nil {
			numbers = append(numbers, n)
		}
	}

	for _, v := range numbers {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			println("Sorted value:", n)
		}(v)
	}
	// 主线程继续执行
	time.Sleep(time.Second * 10)
}
