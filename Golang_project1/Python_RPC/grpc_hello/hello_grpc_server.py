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
from concurrent import futures

# 定义Greeter服务
class Greeter(helloworld_pb2_grpc.GreeterServicer):
    # 实现SayHello方法
    def SayHello(self, request, context):
        # 返回HelloReply消息
        return helloworld_pb2.HelloReply(message=f"Hello, {request.name}!")

# 启动grpc服务
def serve():
    # 创建grpc服务器
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # 注册Greeter服务
    helloworld_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    # 设置服务器地址和端口
    server.add_insecure_port('[::]:50051') # 添加监听地址和端口
    print("gRPC server listening on port 50051")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve() # 确保调用serve函数