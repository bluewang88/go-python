package gorpc_server_proxy

import (
	gorpc_handler "Golang_project1/Go_RPC/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	// 注册处理逻辑handler
	println("--------Registering HelloRpcService in rpc server proxy...-------")
	err := rpc.RegisterName(gorpc_handler.HelloRpcServiceName, srv)
	if err != nil {
		println("Error registering HelloRpcService:", err)
		panic(err)
	}
	println("HelloRpcService registered successfully.")
	return err
}
