package main

import (
	"Golang_project1/Go_base"
)

func main() {
	// fmt.Println("通过fmt.Println函数打印Hello, World!")
	// Go_base.HelloVar()
	// Go_base.HelloDefaultVar()
	// Go_base.HelloDatatypeConvert()
	// Go_base.HelloStrconv()
	// Go_base.ArgsDemo()
	// Go_base.FlagDemo()

	// Go_base.PointerReceiverDemo()
	// Go_base.ValueReceiverDemo()
	Go_base.Hello_gorouties()
}

// Run the code
// go run hello.go
// Output: Hello, World!
// The output is the string "Hello, World!".

//先编译再运行
// go build hello.go
//会生成一个hello的可执行文件
// 运行 ./hello
// Output: Hello, World!
