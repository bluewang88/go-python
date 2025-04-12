package gorpc_server

import (
	gorpcserverhelloworld "Golang_project1/Go_RPC/server/helloWorld"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func HelloHttpRpcServer() {
	_ = rpc.RegisterName("HelloService", &gorpcserverhelloworld.HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct { //结构体嵌入interface，需要结构体实现所有的接口方法才能表示结构体表示了有效的io.ReadWriteCloser类型
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		} //结构体初始化
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	println("Starting HTTP RPC server on port 1234...")
	http.ListenAndServe(":1234", nil)
}

//
