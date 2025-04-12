import sys
import os

# 清晰地定义路径
current_dir = os.path.dirname(os.path.abspath(__file__))
parent_dir = os.path.dirname(current_dir)
grandparent_dir = os.path.dirname(parent_dir)

# 添加相关路径到Python搜索路径
sys.path.append(parent_dir)  # 添加protobuf_test目录
print(f"添加的路径: {parent_dir}")
print(f"Python搜索路径: {sys.path}")  # 打印出搜索路径便于调试

try:
    from protobuf_test import hello_pb2
    print("导入成功!")
except ImportError as e:
    print(f"导入错误: {e}")
    
    # 尝试直接导入
    try:
        sys.path.append(os.path.join(parent_dir, "protobuf_test"))
        import hello_pb2
        print("直接导入成功!")
    except ImportError as e2:
        print(f"直接导入也失败: {e2}")

request = hello_pb2.HelloRequest()

request.name = "zhangsan"

serialized_request = request.SerializeToString() # 序列化
print(serialized_request)
print(len(serialized_request)) # 打印序列化后的字节长度

res_json = {    
    "name": "zhangsan"
} # 模拟一个字典


print("res_json: ", res_json)
import json

print("json.dumps(res_json): ", json.dumps(res_json)) # 将字典转换为JSON字符串
print(len(json.dumps(res_json))) # 打印JSON字符串的长度

# 反序列化,通过字符串反向生成对象

request2 = hello_pb2.HelloRequest()
request2.ParseFromString(serialized_request) # 反序列化
print("request2: ", request2)
print("request2.name: ",request2.name)
print("Serialized request2:", request2.SerializeToString())