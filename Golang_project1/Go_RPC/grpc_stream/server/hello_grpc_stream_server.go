package grpc_stream_server

import (
	protostream "Golang_project1/Go_RPC/grpc_stream/proto"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":50051" // 端口号

type StreamServer struct {
	protostream.UnimplementedGreeterServer
}

// func (s *StreamServer) mustEmbedUnimplementedGreeterServer() {}

func (s *StreamServer) GetStream(req *protostream.StreamReqData, res_stream grpc.ServerStreamingServer[protostream.StreamResData]) error {
	// fmt.Println("GRPC Server GetStream Function Received context: " + ctx.Value("key").(string))
	fmt.Println("GO GRPC Server GetStream function received: " + req.Data)
	// 处理单向流请求
	// 返回响应数据
	for i := 0; i < 10; i++ {
		err := res_stream.Send(&protostream.StreamResData{
			Data: "Hello from server" + fmt.Sprintf(" %v", time.Now().Unix()),
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1) // 每秒发送一次数据
	}
	fmt.Println("GRPC Server GetStream Function Send Finished")
	return nil
}
func (s *StreamServer) PutStream(client_stream grpc.ClientStreamingServer[protostream.StreamReqData, protostream.StreamResData]) error {
	// fmt.Println("GRPC Server PutStream Function Received context: " + client_stream.Context().Value("key").(string))
	// 处理客户端流请求
	for {
		// 接收客户端请求
		req, err := client_stream.Recv()
		if err != nil {
			fmt.Println("GRPC Server PutStream Function Received error: ", err)
			// When we reach the end of the stream or encounter an error,
			// send a single response and close the stream
			return client_stream.SendAndClose(&protostream.StreamResData{Data: "Hello from RPC stream server"})

		}
		fmt.Println("GO GRPC Server PutStream function received: " + req.Data)
		// 处理请求
		// Don't send responses within the loop for client streaming
	}
}

func (s *StreamServer) AllStream(allstream grpc.BidiStreamingServer[protostream.StreamReqData, protostream.StreamResData]) error {
	// fmt.Println("GRPC Server AllStream Function Received context: " + allstream.Context().Err().Error().Context().Value("key").(string))
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 启动一个 goroutine 来处理发送数据
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			err := allstream.Send(&protostream.StreamResData{
				Data: "Hello from go rpc stream server allstream func" + fmt.Sprintf(" %v", time.Now().Unix()),
			})
			if err != nil {
				fmt.Println("Error sending data:", err)
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()

	// 启动一个 goroutine 来处理接收数据
	go func() {
		defer wg.Done()
		for {
			// 接收客户端请求
			req, err := allstream.Recv()
			if err != nil {
				fmt.Println("GRPC Server AllStream Function Received error: ", err)
				return
			}
			fmt.Println("GO GRPC Server AllStream function received: " + req.Data)
			// 处理请求
			// 返回响应数据
			if err := allstream.Send(&protostream.StreamResData{Data: "Hello from server"}); err != nil {
				return
			}
		}
	}()

	wg.Wait()
	return nil
	// // 处理双向流请求
	// for {
	// 	// 接收客户端请求
	// 	req, err := stream.Recv()
	// 	if err != nil {
	// 		fmt.Println("GRPC Server AllStream Function Received error: ", err)
	// 		return err
	// 	}
	// 	fmt.Println("GO GRPC Server AllStream function received: " + req.Data)
	// 	// 处理请求
	// 	// 返回响应数据
	// 	if err := stream.Send(&protostream.StreamResData{Data: "Hello from server"}); err != nil {
	// 		return err
	// 	}
	// }
}

func HelloRpcStreamServer() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		panic(err)
	}
	defer listen.Close()
	fmt.Println("GO GRPC Stream Server成功监听50051端口")
	// 创建grpc服务器
	grpcServer := grpc.NewServer()
	// 注册Greeter服务
	protostream.RegisterGreeterServer(grpcServer, &StreamServer{})
	fmt.Println("GO GRPC Stream Server注册Greeter服务成功")
	// 启动grpc服务器
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println("GRPC Server Start Failed")
		return
	}
	fmt.Println("GRPC Server Start Success")
}
