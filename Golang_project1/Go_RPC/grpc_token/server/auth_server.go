package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "Golang_project1/Go_RPC/grpc_token/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 常量定义
const (
	PORT        = ":50051"            // 服务监听端口
	VALID_TOKEN = "auth-token-123456" // 有效的令牌示例
)

// AuthServer 实现 AuthService 接口
type AuthServer struct {
	proto.UnimplementedAuthServiceServer
}

// Login 处理登录请求，验证用户名密码并返回令牌
func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	log.Printf("收到登录请求: 用户名=%s, 密码=%s", req.Username, req.Password)

	// 模拟用户验证逻辑 (在实际应用中应使用安全的验证机制)
	if req.Username == "admin" && req.Password == "password" {
		return &proto.LoginResponse{
			Token:   VALID_TOKEN,
			Message: "登录成功",
		}, nil
	}

	return nil, status.Error(codes.Unauthenticated, "用户名或密码错误")
}

// SayHello 是一个需要认证的方法示例
func (s *AuthServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("收到SayHello请求: name=%s", req.Name)
	return &proto.HelloResponse{
		Message: fmt.Sprintf("你好, %s! 这是一个经过身份验证的响应。", req.Name),
	}, nil
}

// AuthInterceptor 是服务端的一元拦截器，用于处理认证
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Login方法不需要认证
	if info.FullMethod == "/AuthService/Login" {
		log.Println("跳过Login方法的认证")
		return handler(ctx, req)
	}

	// 从metadata中提取token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "缺少metadata")
	}

	// 获取authorization值
	values := md.Get("authorization")
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "缺少authorization token")
	}

	// 验证token
	token := values[0]
	if token != VALID_TOKEN {
		return nil, status.Error(codes.Unauthenticated, "无效的token")
	}

	log.Println("认证成功，继续处理请求")
	// 调用实际的处理程序
	return handler(ctx, req)
}

// StreamAuthInterceptor 是服务端的流式拦截器，用于处理流式API的认证
func StreamAuthInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// 从上下文中获取metadata
	ctx := ss.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "缺少metadata")
	}

	// 获取authorization值
	values := md.Get("authorization")
	if len(values) == 0 {
		return status.Error(codes.Unauthenticated, "缺少authorization token")
	}

	// 验证token
	token := values[0]
	if token != VALID_TOKEN {
		return status.Error(codes.Unauthenticated, "无效的token")
	}

	log.Println("流式认证成功，继续处理请求")
	// 调用实际的处理程序
	return handler(srv, ss)
}

func main() {
	// 创建监听器
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}

	// 创建gRPC服务器，注册拦截器
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(AuthInterceptor),
		grpc.StreamInterceptor(StreamAuthInterceptor),
	)

	// 注册服务
	proto.RegisterAuthServiceServer(grpcServer, &AuthServer{})

	log.Printf("服务器启动，监听端口%s", PORT)
	// 启动服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
