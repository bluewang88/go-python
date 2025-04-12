package main

import (
	gorpc_client_proxy "Golang_project1/Go_RPC/client_proxy"
	"fmt"
)

func main() {
	fmt.Println("Json Optimized Client Connecting to server on 1234...")
	// conn, err := net.Dial("tcp", ":1234")
	// if err != nil {
	// 	fmt.Println("Error connecting to server")
	// 	panic(err)
	// }
	// defer conn.Close()
	// // Create a new JSON-RPC client
	// fmt.Println("Creating JSON-RPC client...")
	// client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	// var reply string
	fmt.Println("RPC client ready to init a proxy")
	client := gorpc_client_proxy.NewHelloServiceStub("tcp", "localhost:1234")
	var reply string
	defer client.Close()
	// Call the Hello method on the HelloService
	fmt.Println("RPC Client Calling Hello method...")
	err := client.Hello("World", &reply)
	if err != nil {
		fmt.Println("Error calling Hello method:", err)
		panic(err)
	}
	fmt.Println(reply)
}
