package gorpc_server

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func HelloJsonRpcServer() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	_ = rpc.RegisterName("HelloService", &HelloService{})

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
