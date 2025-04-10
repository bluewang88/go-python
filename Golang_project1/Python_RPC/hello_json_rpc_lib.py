"""
JSON-RPC 使用 jsonrpclib 库的示例

JSON-RPC 是一种基于 JSON 的轻量级远程过程调用(RPC)协议。
相比 XML-RPC，它更加轻量、高效，并且使用 JSON 作为数据格式。

主要特点:
1. 数据格式为 JSON，更小更快，解析更高效
2. 支持批量请求和通知（不需要响应的请求）
3. 支持同步和异步调用模式
4. 内容类型为 application/json
5. 支持 HTTP、TCP、Unix Socket 等多种传输方式
"""

# 安装 jsonrpclib-pelix（如果尚未安装）:
# pip install jsonrpclib-pelix

import threading
import time
from jsonrpclib.SimpleJSONRPCServer import SimpleJSONRPCServer
from jsonrpclib import Server

# 设置服务器地址和端口
HOST = 'localhost'
PORT = 8080

class AdvancedCalculatorService:
    """
    高级计算器服务
    
    提供多种计算功能的服务类，所有方法都可以被远程调用。
    这个类演示了 JSON-RPC 在面向对象编程中的应用。
    """
    
    def __init__(self):
        # 计算历史记录
        self.history = []
    
    def add(self, x, y):
        """
        执行加法运算
        
        参数:
            x (number): 第一个操作数
            y (number): 第二个操作数
            
        返回:
            number: x 和 y 的和
        """
        result = x + y
        self.history.append(f"add: {x} + {y} = {result}")
        print(f"服务器执行加法: {x} + {y} = {result}")
        return result
    
    def subtract(self, x, y):
        """
        执行减法运算
        
        参数:
            x (number): 第一个操作数
            y (number): 第二个操作数
            
        返回:
            number: x 减去 y 的差
        """
        result = x - y
        self.history.append(f"subtract: {x} - {y} = {result}")
        print(f"服务器执行减法: {x} - {y} = {result}")
        return result
    
    def multiply(self, x, y):
        """
        执行乘法运算
        
        参数:
            x (number): 第一个操作数
            y (number): 第二个操作数
            
        返回:
            number: x 和 y 的乘积
        """
        result = x * y
        self.history.append(f"multiply: {x} * {y} = {result}")
        print(f"服务器执行乘法: {x} * {y} = {result}")
        return result
    
    def divide(self, x, y):
        """
        执行除法运算
        
        参数:
            x (number): 第一个操作数
            y (number): 第二个操作数
            
        返回:
            number: x 除以 y 的商
            
        异常:
            ValueError: 如果 y 为 0，将抛出异常
        """
        if y == 0:
            error_message = "除数不能为零"
            self.history.append(f"divide: 错误 - {error_message}")
            print(f"服务器执行除法: 错误 - {error_message}")
            raise ValueError(error_message)
        
        result = x / y
        self.history.append(f"divide: {x} / {y} = {result}")
        print(f"服务器执行除法: {x} / {y} = {result}")
        return result
    
    def power(self, base, exponent):
        """
        计算幂
        
        参数:
            base (number): 底数
            exponent (number): 指数
            
        返回:
            number: base 的 exponent 次方
        """
        result = base ** exponent
        self.history.append(f"power: {base} ^ {exponent} = {result}")
        print(f"服务器执行幂运算: {base} ^ {exponent} = {result}")
        return result
    
    def get_history(self):
        """
        获取计算历史记录
        
        返回:
            list: 计算历史记录列表
        """
        print("服务器返回历史记录")
        return self.history
    
    def clear_history(self):
        """
        清除计算历史记录
        
        返回:
            bool: 操作是否成功
        """
        self.history = []
        print("服务器清除历史记录")
        return True
    
    def batch_calculate(self, operations):
        """
        批量执行计算操作
        
        参数:
            operations (list): 操作列表，每个操作是一个字典，包含 'op', 'x', 'y' 键
                'op': 操作类型，可以是 'add', 'subtract', 'multiply', 'divide', 'power'
                'x': 第一个操作数
                'y': 第二个操作数
                
        返回:
            list: 每个操作的结果列表
            
        这个方法展示了 JSON-RPC 处理复杂数据结构的能力。
        """
        results = []
        for op in operations:
            try:
                if op['op'] == 'add':
                    results.append(self.add(op['x'], op['y']))
                elif op['op'] == 'subtract':
                    results.append(self.subtract(op['x'], op['y']))
                elif op['op'] == 'multiply':
                    results.append(self.multiply(op['x'], op['y']))
                elif op['op'] == 'divide':
                    results.append(self.divide(op['x'], op['y']))
                elif op['op'] == 'power':
                    results.append(self.power(op['x'], op['y']))
                else:
                    results.append({"error": f"未知操作: {op['op']}"})
            except Exception as e:
                results.append({"error": str(e)})
        
        return results

def start_server():
    """
    启动 JSON-RPC 服务器
    
    创建并配置 JSON-RPC 服务器，注册服务实例，启动服务循环。
    """
    # 创建 JSON-RPC 服务器
    # SimpleJSONRPCServer 类似于 XML-RPC 的 SimpleXMLRPCServer，但使用 JSON 作为数据格式
    server = SimpleJSONRPCServer((HOST, PORT))
    
    # 创建计算器服务实例
    calculator = AdvancedCalculatorService()
    
    # 注册实例 - 所有实例的公共方法都将被暴露为 RPC 方法
    # 客户端可以通过方法名直接调用这些方法
    server.register_instance(calculator)
    
    print(f"JSON-RPC 服务器已启动，监听地址: {HOST}:{PORT}")
    print("使用 Ctrl+C 关闭服务器")
    
    # 启动服务器循环
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        print("服务器已关闭")

def run_client():
    """
    运行 JSON-RPC 客户端
    
    连接到服务器并调用远程方法，展示不同的调用方式和处理响应的方法。
    """
    # 等待服务器启动
    time.sleep(1)
    
    # 创建 JSON-RPC 客户端
    # Server 函数创建一个代理对象，所有对它的调用都将被转发到远程服务器
    client = Server(f"http://{HOST}:{PORT}")
    
    try:
        print("\n=== 基本算术运算 ===")
        # 调用远程方法 - 这些看起来像本地函数调用，但实际上是:
        # 1. 客户端构造 JSON-RPC 请求
        # 2. 通过 HTTP 发送请求
        # 3. 服务器解析请求，执行相应方法
        # 4. 结果编码为 JSON 并返回
        # 5. 客户端解析 JSON 响应
        
        # 加法
        result = client.add(10, 20)
        print(f"10 + 20 = {result}")
        
        # 减法
        result = client.subtract(30, 15)
        print(f"30 - 15 = {result}")
        
        # 乘法
        result = client.multiply(5, 7)
        print(f"5 * 7 = {result}")
        
        # 除法
        result = client.divide(100, 4)
        print(f"100 / 4 = {result}")
        
        # 幂运算
        result = client.power(2, 10)
        print(f"2 ^ 10 = {result}")
        
        print("\n=== 异常处理 ===")
        # 处理服务器端异常
        # JSON-RPC 将服务器异常转换为客户端异常
        try:
            result = client.divide(10, 0)
        except Exception as e:
            print(f"除以零的异常处理: {e}")
        
        print("\n=== 复杂数据结构 ===")
        # 传递和接收复杂数据结构
        # JSON-RPC 可以传输各种 JSON 兼容的数据类型: 
        # 数字、字符串、布尔值、数组、对象、null
        operations = [
            {'op': 'add', 'x': 10, 'y': 5},
            {'op': 'subtract', 'x': 20, 'y': 8},
            {'op': 'multiply', 'x': 3, 'y': 4},
            {'op': 'divide', 'x': 15, 'y': 3},
            {'op': 'power', 'x': 2, 'y': 3},
            {'op': 'divide', 'x': 1, 'y': 0}  # 故意产生错误
        ]
        
        results = client.batch_calculate(operations)
        print("批量计算结果:")
        for i, result in enumerate(results):
            op = operations[i]
            print(f"  操作 {i+1}: {op} → {result}")
        
        print("\n=== 历史记录 ===")
        # 获取历史记录
        history = client.get_history()
        print("计算历史记录:")
        for entry in history:
            print(f"  {entry}")
        
        # 清除历史记录
        success = client.clear_history()
        print(f"清除历史记录: {'成功' if success else '失败'}")
        
        # 确认历史记录已清除
        history = client.get_history()
        print(f"清除后的历史记录数量: {len(history)}")
        
    except ConnectionRefusedError:
        print("连接服务器失败，确保服务器已启动")
    except Exception as e:
        print(f"出现错误: {e}")

def main():
    """
    主函数，协调服务器和客户端的启动和执行
    """
    # 在后台线程中启动服务器
    server_thread = threading.Thread(target=start_server, daemon=True)
    server_thread.start()
    
    # 在主线程中运行客户端
    run_client()
    
    # 保持程序运行，直到用户手动终止
    print("\n服务器仍在后台运行。按 Enter 键退出...")
    input()

if __name__ == "__main__":
    main()