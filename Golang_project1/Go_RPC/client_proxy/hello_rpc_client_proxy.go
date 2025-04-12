package gorpc_client_proxy

import (
	gorpc_handler "Golang_project1/Go_RPC/handler"
	"fmt"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// 在go中没有对象、类，就以为没有初始化方法
// 通过构造函数接收服务器地址(addr)，实现了依赖注入，增强了代码灵活性和可测试性
func NewHelloServiceStub(protocol, addr string) *HelloServiceStub {
	fmt.Println("------------rpc client proxy connect to server --------------")
	client, err := rpc.Dial(protocol, addr)
	if err != nil {
		fmt.Println("Error connecting to server")
		panic(err)
	}
	return &HelloServiceStub{
		Client: client,
	}
}

/*
嵌入类型的初始化规则
1. 结构体中嵌入了一个类型，称为匿名字段
2. 使用类型名作为字段名（本例使用的方法）
3. 通过嵌入类型的字段名来访问嵌入类型的方法
4. 不指定字段名（直接传值）：
*/

// 创建RPC客户端代理
func (s *HelloServiceStub) Hello(request string, reply *string) error {
	fmt.Println("------------rpc client proxy call Hello method --------------")
	return s.Call(gorpc_handler.HelloRpcServiceName+".Hello", request, reply)
}
