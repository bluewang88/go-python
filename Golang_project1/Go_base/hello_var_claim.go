package Go_base

import "fmt"

func HelloVarClaim() {
	fmt.Println("-----------变量声明---开始------")
	//变量声明
	var name string = "John"
	var age int = 25
	var isStudent bool = true
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("isStudent:", isStudent)
	fmt.Println("-----------变量声明---结束------")

	//短变量声明
	//短变量声明使用 := 操作符，可以同时声明和初始化变量。
	//短变量声明的变量类型由编译器根据初始值推断。
	//短变量声明只能在函数内部使用。
	//短变量声明的变量名必须是一个新的变量名，不能与已存在的变量名相同。
	//短变量声明的变量名必须是一个有效的标识符。
	//短变量声明的变量名必须是一个有效的标识符。
	//短变量声明的变量名必须是一个有效的标识符。
	fmt.Println("-----------短变量声明---开始------")
	shortName := "John"
	shortAge := 25
	shortIsStudent := true
	fmt.Println("shortName:", shortName)
	fmt.Println("shortAge:", shortAge)
	fmt.Println("shortIsStudent:", shortIsStudent)
	fmt.Println("-----------短变量声明---结束------")
}
