package main

import (
	"context"
	"log"
	"os"
	"time"

	proto "Golang_project1/Go_RPC/grpc_token/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	ADDRESS = "localhost:50051" // 服务器地址
)

// tokenAuth 实现自定义凭证
type tokenAuth struct {
	token string
}

// GetRequestMetadata 实现凭证接口，返回metadata
func (t *tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.token,
	}, nil
}

// RequireTransportSecurity 指定是否需要传输安全性
func (t *tokenAuth) RequireTransportSecurity() bool {
	return false // 本示例中不需要传输安全性
}

// AuthClientInterceptor 客户端拦截器，用于添加认证metadata
func AuthClientInterceptor(token string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Login方法不需要认证
		if method == "/AuthService/Login" {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		// 为其他方法添加token到context
		log.Println("客户端拦截器: 添加认证token到metadata")
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func main() {
	// 创建不安全的连接（在生产环境应使用TLS）
	conn, err := grpc.Dial(
		ADDRESS,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(AuthClientInterceptor("")), // 初始无token
	)
	if err != nil {
		log.Fatalf("无法连接到服务器: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewAuthServiceClient(conn)

	// 步骤1: 登录获取token
	token, err := login(client)
	if err != nil {
		log.Fatalf("登录失败: %v", err)
		os.Exit(1)
	}

	// 更新连接使用新的token
	conn.Close()
	conn, err = grpc.Dial(
		ADDRESS,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(AuthClientInterceptor(token)),
	)
	if err != nil {
		log.Fatalf("无法重新连接到服务器: %v", err)
	}
	defer conn.Close()
	client = proto.NewAuthServiceClient(conn)

	// 步骤2: 使用token调用需要认证的方法
	callAuthenticatedMethod(client, token)
}

// login 处理登录并返回token
func login(client proto.AuthServiceClient) (string, error) {
	log.Println("尝试登录...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 创建登录请求
	loginReq := &proto.LoginRequest{
		Username: "admin",
		Password: "password",
	}

	// 发送登录请求
	resp, err := client.Login(ctx, loginReq)
	if err != nil {
		return "", err
	}

	log.Printf("登录成功: %s", resp.Message)
	log.Printf("获取到token: %s", resp.Token)
	return resp.Token, nil
}

// callAuthenticatedMethod 调用需要认证的方法
func callAuthenticatedMethod(client proto.AuthServiceClient, token string) {
	log.Println("调用需要认证的方法...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 方法1: 使用拦截器自动添加token (在创建连接时已配置)
	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "gRPC用户"})
	if err != nil {
		log.Fatalf("调用SayHello失败: %v", err)
		return
	}
	log.Printf("SayHello响应(使用拦截器): %s", resp.Message)

	// 方法2: 手动添加token到metadata
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", token))
	resp, err = client.SayHello(ctx, &proto.HelloRequest{Name: "gRPC用户(手动认证)"})
	if err != nil {
		log.Fatalf("调用SayHello失败(手动认证): %v", err)
		return
	}
	log.Printf("SayHello响应(手动添加token): %s", resp.Message)
}
