# 导入 XML-RPC 服务器和客户端模块
from xmlrpc.server import SimpleXMLRPCServer
from xmlrpc.client import ServerProxy
import threading
import time

# 定义服务器地址和端口
HOST = 'localhost'
PORT = 8080

"""
RPC (Remote Procedure Call) 简介:
- RPC 允许一个程序调用另一个地址空间(通常是网络上的另一台计算机)的子程序，就像调用本地程序一样
- XML-RPC 是通过 HTTP 协议传输 XML 格式数据来实现 RPC 的一种方式
- 主要组件:
  1. 服务端: 提供可被远程调用的函数或方法
  2. 客户端: 发起远程调用，并获取结果
  3. 服务端存根: 处理客户端请求，调用本地方法并返回结果
  4. 客户端存根: 客户端存根是客户端程序与远程服务器通信的代理，它封装了远程方法调用，并提供了与远程服务器通信所需的所有信息
"""

class CalculatorService:
    """计算服务类，包含可以被远程调用的方法"""
    
    def add(self, x, y):
        """加法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 和 y 的和
        """
        print(f"服务器执行: {x} + {y}")
        return x + y
    
    def subtract(self, x, y):
        """减法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 减去 y 的差
        """
        print(f"服务器执行: {x} - {y}")
        return x - y
    
    def multiply(self, x, y):
        """乘法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 和 y 的乘积
        """
        print(f"服务器执行: {x} * {y}")
        return x * y
    
    def divide(self, x, y):
        """除法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 除以 y 的商
            
        异常:
            如果 y 为 0，将抛出 ZeroDivisionError 异常
        """
        if y == 0:
            raise ValueError("除数不能为零")
        print(f"服务器执行: {x} / {y}")
        return x / y

def start_server():
    """启动 XML-RPC 服务器"""
    # 创建服务器实例
    server = SimpleXMLRPCServer((HOST, PORT), allow_none=True, logRequests=True)
    
    # 注册内省函数，允许客户端查询服务器提供的方法
    server.register_introspection_functions()
    
    # 创建计算器服务实例
    calculator = CalculatorService()
    
    # 将整个实例注册到服务器，所有公共方法都将可用
    server.register_instance(calculator)
    
    print(f"XML-RPC 服务器已启动，监听地址: {HOST}:{PORT}")
    
    # 开始接收和处理请求
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        print("服务器已关闭")

def run_client():
    """运行 XML-RPC 客户端"""
    # 等待服务器启动
    time.sleep(1)
    
    # 创建客户端代理
    # ServerProxy 对象创建了一个连接到远程服务器的代理对象
    # 所有对这个对象的方法调用都会被转发到远程服务器
    client = ServerProxy(f"http://{HOST}:{PORT}")
    
    try:
        # 调用远程方法
        # 这些调用看起来像本地函数调用，但实际上:
        # 1. 客户端将调用信息编码为 XML
        # 2. 通过 HTTP 发送到服务器
        # 3. 服务器解码请求，执行对应方法
        # 4. 将结果编码后返回给客户端
        # 5. 客户端解码结果并返回给调用者
        
        # 调用加法
        result = client.add(10, 20)
        print(f"10 + 20 = {result}")
        
        # 调用减法
        result = client.subtract(30, 10)
        print(f"30 - 10 = {result}")
        
        # 调用乘法
        result = client.multiply(5, 6)
        print(f"5 * 6 = {result}")
        
        # 调用除法
        result = client.divide(100, 5)
        print(f"100 / 5 = {result}")
        
        # 处理异常情况
        try:
            result = client.divide(10, 0)
        except Exception as e:
            print(f"异常处理: {e}")
        
        # 获取服务器支持的方法列表
        methods = client.system.listMethods()
        print("服务器支持的方法:")
        for method in methods:
            print(f"  - {method}")
            # 获取方法的帮助信息
            if method.startswith('system.'):
                continue
            help_text = client.system.methodHelp(method)
            if help_text:
                print(f"    帮助: {help_text}")
            
    except ConnectionRefusedError:
        print("连接服务器失败，确保服务器已启动")
    except Exception as e:
        print(f"出现错误: {e}")

def main():
    """主函数，启动服务器和客户端"""
    # 在单独的线程中启动服务器
    server_thread = threading.Thread(target=start_server, daemon=True)
    server_thread.start()
    
    # 在主线程中运行客户端
    run_client()
    
    # 等待用户输入以保持程序运行
    input("按 Enter 键退出...")

if __name__ == "__main__":
    main()



