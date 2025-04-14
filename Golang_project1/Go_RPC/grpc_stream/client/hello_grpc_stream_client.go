package main

import (
	protostream "Golang_project1/Go_RPC/grpc_stream/proto"
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		panic(err)
	}
	defer conn.Close()
	client := protostream.NewGreeterClient(conn) // 创建客户端
	replay, err := client.GetStream(context.Background(), &protostream.StreamReqData{Data: " from go client world"})
	// replay, err := client.SayHello(nil, &proto.HelloRequest{Name: "world"})
	if err != nil {
		fmt.Println("Error calling GetStream:", err)
		return
	}
	// 服务端流模式
	fmt.Println("------------服务度端流模式GetStream---------")
	for {
		res, err := replay.Recv() // 接收服务器响应
		if err != nil {
			fmt.Println("Error receiving response:", err)
			break
		}
		fmt.Println("Response from server:", res.Data)
	}

	// 创建客户端流
	fmt.Println("------------客户端流模式PutStream---------")
	client_stream, err := client.PutStream(context.Background())
	if err != nil {
		fmt.Println("Error creating client stream:", err)
		return
	}

	for i := 0; i < 10; i++ {
		err := client_stream.Send(&protostream.StreamReqData{Data: fmt.Sprintf("Message %d from stream client", i)})
		if err != nil {
			fmt.Println("Error sending data:", err)
			break
		}
	}

	res, err := client_stream.CloseAndRecv()
	if err != nil {
		fmt.Println("Error closing client stream:", err)
		return
	}
	fmt.Println("Response from server:", res.Data)

	// 双向流模式
	fmt.Println("------------双向流模式AllStream---------")
	all_stream, err := client.AllStream(context.Background())
	if err != nil {
		fmt.Println("Error creating all stream:", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	// 启动一个 goroutine 来处理发送数据
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			err := all_stream.Send(&protostream.StreamReqData{
				Data: "Hello from go rpc stream client allstream func" + fmt.Sprintf(" %v", i),
			})
			if err != nil {
				fmt.Println("Error sending data:", err)
				return
			}
		}
		all_stream.CloseSend()
	}()
	// 启动一个 goroutine 来处理接收数据
	go func() {
		defer wg.Done()
		for {
			// 接收服务器响应
			res, err := all_stream.Recv()
			if err != nil {
				fmt.Println("Error receiving response:", err)
				break
			}
			fmt.Println("Response from server:", res.Data)
		}
	}()
	// 阻塞主 goroutine，等待所有 goroutine 完成
	wg.Wait()

}
