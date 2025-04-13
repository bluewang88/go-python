package grpc_server_hello

import (
	proto "Golang_project1/Go_RPC/grpc_hello/proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloGrpcServer struct {
	proto.UnimplementedGreeterServer
}

// 实现SayHello方法
// 参数：ctx context.Context, req *proto.HelloRequest
// 返回值：*proto.HelloReply, error
// 方法体：
// 1. 打印接收到的请求信息
// 2. 返回响应信息
// 3. 返回错误信息
// 4. 返回nil
// ctx: 上下文，用于传递请求和响应, 可以理解为请求上下文

func (s *HelloGrpcServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println("GRPC Server SayHello Function Received: " + req.Name)
	return &proto.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

func HelloGrpcNewServer() {
	grpcServer := grpc.NewServer()                              // 创建grpc服务器
	proto.RegisterGreeterServer(grpcServer, &HelloGrpcServer{}) // 注册Greeter服务
	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("成功监听50051端口")
	if err := grpcServer.Serve(listen); err != nil { // 启动grpc服务器
		log.Fatalf("failed to serve: %v", err)
		fmt.Println("GRPC Server Start Failed")
		// log.Fatal(err)
		// return nil
	}
	log.Println("GRPC Server Start Success")

}
