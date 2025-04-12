package gorpc_server

import (
	gorpcserverhelloworld "Golang_project1/Go_RPC/server/helloWorld"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func HelloJsonRpcServer() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	_ = rpc.RegisterName("HelloService", &gorpcserverhelloworld.HelloService{})

	println("Starting JSON RPC server on port 1234...")

	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 处理连接
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
