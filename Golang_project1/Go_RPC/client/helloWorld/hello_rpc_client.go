package main

import "net/rpc"

func main() {
	//建立连接
	client, err := rpc.Dial("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "World", &reply)
	if err != nil {
		panic(err)
	}

	println(reply)
}
