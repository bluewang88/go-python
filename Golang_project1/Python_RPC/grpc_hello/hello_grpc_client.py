import sys
import os

# 直接添加proto目录到Python路径
proto_dir = os.path.join(os.path.dirname(os.path.abspath(__file__)), "proto")
sys.path.append(proto_dir)
print(f"添加proto目录到Python路径: {proto_dir}")

# --- 修改这里的导入语句 ---
# 直接导入模块，因为proto_dir已经在sys.path中
import helloworld_pb2_grpc
import helloworld_pb2
# --- 修改结束 ---

import grpc

if __name__ == '__main__':
    # --- 修改这里的调用语句 ---
    # 创建grpc通道
    with grpc.insecure_channel('localhost:50051') as channel:
        # 创建grpc客户端
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        # 调用grpc服务端
        response: helloworld_pb2.HelloReply = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))
        # --- 修改结束 ---
        print("Greeter client received: " + response.message)
    # --- 修改结束 ---

