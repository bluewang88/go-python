package main

import (
	"Golang_project1/Go_RPC/grpc_hello/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		panic(err)
	}
	defer conn.Close()
	client := proto.NewGreeterClient(conn) // 创建客户端
	replay, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "world"})
	// replay, err := client.SayHello(nil, &proto.HelloRequest{Name: "world"})
	if err != nil {
		fmt.Println("Error calling SayHello:", err)
		return
	}
	fmt.Println("Response from server:", replay.Message)

}
