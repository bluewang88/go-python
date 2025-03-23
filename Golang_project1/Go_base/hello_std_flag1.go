package Go_base

import (
	"flag"
	"fmt"
)

// flag包支持的参数数据类型有bool、int、int64、uint、uint64、float、float64、string、duration
// 定义命令行参数
//格式：
// flag.Type(flag名, 默认值, 帮助信息)
//举例：
// var ip = flag.Int("flagname", 1234, "help message for flagname")
// flag.TypeVar(&变量, flag名, 默认值, 帮助信息)
// flag.TypeVarP(&变量, flag名, flag别名, 默认值, 帮助信息)

func FlagDemo() {
	var ip = flag.Int("ip", 1234, "help message for ip")
	flag.Parse()
	fmt.Println("IP:", *ip)
}
