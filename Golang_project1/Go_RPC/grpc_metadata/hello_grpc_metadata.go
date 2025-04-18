package grpcmetadata_test

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 定义一个简单的 gRPC 服务接口
type HelloService interface {
	SayHello(ctx context.Context, request *HelloRequest) (*HelloResponse, error)
}

// 请求消息结构
type HelloRequest struct {
	Name string
}

// 响应消息结构
type HelloResponse struct {
	Message string
}

// 服务实现
type helloServer struct {
}

// SayHello 实现了 HelloService 接口
func (s *helloServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	// 从上下文中获取 metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("无法获取 metadata")
	}

	// 打印 metadata
	fmt.Println("收到的 metadata:", md)

	// 向客户端返回响应
	return &HelloResponse{
		Message: fmt.Sprintf("你好, %s!", req.Name),
	}, nil
}

// 客户端 metadata 使用示例
func HelloGrpcMetadata() {
	// 创建 metadata 的第一种方式
	md := metadata.New(map[string]string{
		"key1": "value1",
		"key2": "value2",
	})

	// 创建 metadata 的第二种方式, 使用 Pairs 函数
	// Pairs 函数的参数是可变参数, 需要成对出现
	// 例如: "key1", "value1", "key2", "value2"
	// 也可以使用 Pairs 函数创建一个空的 metadata
	// 例如: metadata.Pairs()
	// key 不区分大小写，会统一转换为小写
	md2 := metadata.Pairs(
		"key1", "value1", // 即使是 key 和 value 之间也使用逗号分隔
		// 也可以使用空格分隔
		// "key1 value1",
		// "key2 value2",
		"key2", "value2",
	)

	fmt.Println("第一种方式创建的 metadata:", md)
	fmt.Println("第二种方式创建的 metadata:", md2)

	// 创建到服务端的连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("无法连接到服务器: %v", err)
	}
	defer conn.Close()

	// 创建 client stub
	client := NewHelloServiceClient(conn)

	// 创建请求
	request := &HelloRequest{Name: "gRPC用户"}

	// 新建一个有 metadata 的 context
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 发送 RPC 请求
	response, err := client.SayHello(ctx, request)
	if err != nil {
		fmt.Println("调用 SayHello 出错:", err)
		return
	}

	fmt.Println("收到响应:", response.Message)
}

// HelloServiceClient 接口
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

// 客户端结构体
type helloServiceClient struct {
	cc *grpc.ClientConn
}

// 创建新的客户端
func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc: cc}
}

// 客户端方法实现
func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// 启动服务器
func StartServer() {
	// 监听指定端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 注册服务实现
	RegisterHelloServiceServer(s, &helloServer{})

	log.Println("服务器启动在 :50051")

	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

// RegisterHelloServiceServer 注册服务
func RegisterHelloServiceServer(s *grpc.Server, srv interface{}) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

// 服务描述
var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloService",
	HandlerType: (*HelloService)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

// 请求处理函数
func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloService).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloService).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// 主函数示例
func ExampleMetadata() {
	// 在实际应用中，服务器应该在单独的进程中运行
	// 这里仅作为示例，在同一个函数中演示
	go StartServer()

	// 等待服务器启动
	fmt.Println("等待服务器启动...")

	// 调用客户端函数
	HelloGrpcMetadata()
}
