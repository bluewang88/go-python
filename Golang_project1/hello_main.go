package main

// "Golang_project1/Go_base"
import (
	// objectgotest "Golang_project1/Go_base/objectgo"
	gorouting_test "Golang_project1/Go_base/gorouting"
)

func main() {
	// fmt.Println("通过fmt.Println函数打印Hello, World!")
	// Go_base.HelloVar()
	// Go_base.HelloDefaultVar()
	// Go_base.HelloArray() // 数组
	// Go_base.HelloSlice() // 切片
	// Go_base.HelloSlice2()
	// Go_base.HelloSliceCapGrow() // 切片的容量和增长
	// Go_base.HelloSlicePointers() // 切片指针
	// Go_base.HelloSliceLenCap() // 切片长度和容量
	// Go_base.HelloSliceOfSlice() // 切片的切片
	// Go_base.HelloSliceFunc() // 切片存储函数
	// Go_base.HelloSliceFuncParam() // 切片作为函数参数
	// Go_base.HelloSliceDeleteElem() // 切片删除元素
	// Go_base.HelloArrayValueSemantics() //数组赋值是值传递
	// Go_base.HelloMap() // map
	// Go_base.HelloDatatypeConvert()
	// Go_base.HelloStrconv()
	// Go_base.ArgsDemo()
	// Go_base.FlagDemo()

	// objectgotest.HelloType() // type的使用方式

	// objectgotest.HelloInterface()

	// gorouting_test.HelloMutex1() // 互斥锁
	// gorouting_test.HelloRWMutex() // 读写锁

	// gorouting_test.HelloGoRoutineWorld()\
	gorouting_test.HelloGoroutineSort() // goroutine的使用

	// gorouting_test.HelloChannelType()

	// gorouting_test.HelloChannel() // 通道的使用

	// Go_base.PointerReceiverDemo()
	// Go_base.ValueReceiverDemo()
	// Go_base.Hello_gorouties()
	// Go_base.Hello_func_properties()
	// Go_gin.HelloGin()
	// Go_gin.HelloGinDefaultRouter()
	// Go_gin.HelloGinHttpFunc()
	// Go_gin.HelloRouteGroup()
	// Go_gin.HelloUrlVar()
	// Go_gin.HelloUrlStrict() // 通过uri约束条件获取url参数
	// Go_gin.HelloGetParam() // 通过get和post获取参数
	// Go_gin.HelloProtoBuf() // 返回protobuf格式的数据
	// Go_gin.HelloJson() // 返回json格式的数据
	// Go_gin.HelloProtoBuf() // 返回protobuf格式的数据
	// Go_gin.HelloFormValidation() // 表单验证
	// Go_gin.HelloFormValidationLogin() // 表单验证,login
	// Go_gin.HelloFormValidationRegistion() // 表单验证,registion

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
