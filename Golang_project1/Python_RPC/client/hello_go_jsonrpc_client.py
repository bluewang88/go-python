import json
import socket


request = {
    "method": "HelloService.Hello",
    "params": ["python client rpc"],
    "id": 0
}

client = socket.create_connection(("127.0.0.1", 1234))
client.sendall(json.dumps(request).encode())


# 获取服务器返回的数据
rsp = client.recv(1024)
rsp = json.loads(rsp.decode())
print(rsp)