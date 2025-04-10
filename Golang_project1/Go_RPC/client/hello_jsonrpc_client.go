package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string

	err = client.Call("HelloService.Hello", "World", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
