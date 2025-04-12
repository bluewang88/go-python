package rpcoptimizedhello

import (
	gorpc_handler "Golang_project1/Go_RPC/handler"
	gorpc_server_proxy "Golang_project1/Go_RPC/server_proxy"
	"fmt"
	"net"
	"net/rpc"
)

func HelloRpcServerOptimized() {
	// 创建一个TCP监听器，实例化server端
	println("-----Starting JSON RPC Optimized server on port 1234...------")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	// 注册处理逻辑handler
	fmt.Println("--------Registering HelloRpcService...-------")
	// _ = rpc.RegisterName(gorpc_handler.HelloRpcServiceName, &gorpc_handler.HelloRpcHandler{})
	if err := gorpc_server_proxy.RegisterHelloService(&gorpc_handler.HelloRpcHandler{}); err != nil {
		println("Error registering HelloRpcService:", err)
		panic(err)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept() // 监听连接

		if err != nil {
			println("Error accepting connection:", err)
			panic(err)
		}
		println("New connection established")
		// go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //使用json编码器
		go rpc.ServeConn(conn) //使用gob编码器
		// gorpcserverhelloworld.ServeCodec(conn)

	}

	// println("Shutting down the server...")
}
