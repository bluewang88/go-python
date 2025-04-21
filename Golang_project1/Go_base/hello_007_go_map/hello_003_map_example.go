package Go_base_map

import (
	"fmt"
	"strings"
)

// 写一个程序，统计一个字符串中每个单词出现的次数。比如："how do you do"中how=1 do=2 you=1。
func HelloMapExample() {
	str := "how do you do"

	// 将字符串按空格分割成单词切片
	wordSlice := strings.Fields(str)

	// 创建map用于统计单词出现次数
	wordCount := make(map[string]int)

	// 遍历单词切片，统计每个单词出现的次数
	for _, word := range wordSlice {
		// 已存在的单词次数加1，不存在的初始化为1
		wordCount[word]++
	}

	// 打印结果
	fmt.Println("单词出现次数统计:")
	for word, count := range wordCount {
		fmt.Printf("%s=%d ", word, count)
	}
	fmt.Println()
}
