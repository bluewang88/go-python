package Go_base

import (
	"fmt"
	"time"
)

// 定义一个函数，用于在 goroutine 中打印字符串
func Hello_gorouties_say(s string) {
	// 循环 5 次
	for i := 0; i < 5; i++ {
		// 暂停 100 毫秒
		time.Sleep(100 * time.Millisecond)
		// 打印字符串
		fmt.Println(s)
	}
}

func Hello_gorouties() {
	// 启动一个 goroutine 运行 Hello_gorouties_say 函数
	go Hello_gorouties_say(" run in gorouties")
	// 运行 Hello_gorouties_say 函数
	Hello_gorouties_say(" run in main")
}
