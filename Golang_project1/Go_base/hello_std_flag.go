package Go_base

//学习go标准库flag的用法

import (
	"fmt"
	"os"
)

func ArgsDemo() {
	//os.Args变量获取命令行参数,是一个[]string类型的切片Slice

	if len(os.Args) > 0 {
		//打印os.Args变量
		fmt.Println("-----------打印os.Args变量---开始------")
		fmt.Println("os.Args内容是:", os.Args)
		//遍历os.Args
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%s\n", index, arg)
		}
		fmt.Println("-----------打印os.Args变量---结束------")
	}
}

// ➜  Golang_project1 git:(master) ✗ go run hello_main.go hwob ne 12 78 0 %
// -----------打印os.Args变量---开始------
// os.Args内容是: [/tmp/go-build369014129/b001/exe/hello_main hwob ne 12 78 0 %]
// args[0]=/tmp/go-build369014129/b001/exe/hello_main
// args[1]=hwob
// args[2]=ne
// args[3]=12
// args[4]=78
// args[5]=0
// args[6]=%
// -----------打印os.Args变量---结束------
