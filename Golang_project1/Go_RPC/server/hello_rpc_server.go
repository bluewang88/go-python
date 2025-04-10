package gorpc_server

import (
	"net"
	"net/rpc"
)

//注册接口

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改指针来实现的
	//request 是传入的参数
	*reply = "Hello, " + request
	return nil
}

//rpc使用服务
// 1.实例化一个server
// 2.注册服务，注册处理逻辑handler
// 3.监听端口，启动服务

func HelloRpcServer() {
	//实例化一个server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// 注册服务,注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})

	defer listener.Close()

	//监听端口，启动服务
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	// 处理连接
	//callid
	//序列化和反序列化
	rpc.ServeConn(conn)
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		continue
	// 	}
	// 	go rpc.ServeConn(conn)
	// }
}
