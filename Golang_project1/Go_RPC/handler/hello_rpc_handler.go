package gorpc_handler

// 解决名称冲突的问题加上handler/
const HelloRpcServiceName = "handler/HelloService"

type HelloRpcHandler struct {
}

func (h *HelloRpcHandler) Hello(request string, reply *string) error {
	*reply = "Hello, " + request + " from Go RPC/handler"
	return nil
}
