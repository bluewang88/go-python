"""
ZeroRPC 示例代码

ZeroRPC 是一个轻量级的、高效的远程过程调用(RPC)框架，基于 ZeroMQ 和 MessagePack。
它允许你像调用本地函数一样调用远程函数，并提供了强大的网络通信能力。

主要特性:
1. 基于 ZeroMQ 的高性能网络通信
2. 使用 MessagePack 进行高效序列化
3. 支持多种通信模式：请求-响应、发布-订阅等
4. 跨语言支持：Python、Node.js、Ruby 等
5. 支持异步调用和流式响应
6. 内置心跳和超时机制
7. 易于扩展和集成
"""

# 安装所需的包（如果尚未安装）:
# pip install zerorpc pyzmq msgpack

import zerorpc
import threading
import time
import random
import sys
from datetime import datetime

# 服务器地址配置
SERVER_ADDR = "tcp://127.0.0.1:4242"

class AdvancedMathService(object):
    """
    高级数学服务类
    
    这个类提供了各种数学运算，展示 ZeroRPC 如何暴露 Python 类的功能。
    ZeroRPC 会自动将这个类的所有公共方法（非下划线开头）作为 RPC 方法暴露。
    """
    
    def __init__(self):
        self.calculation_count = 0
        self.last_accessed = datetime.now()
        print("数学服务初始化完成")
    
    def add(self, x, y):
        """
        加法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 和 y 的和
        """
        self._update_stats()
        print(f"服务器执行: {x} + {y}")
        return x + y
    
    def subtract(self, x, y):
        """
        减法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 减 y 的差
        """
        self._update_stats()
        print(f"服务器执行: {x} - {y}")
        return x - y
    
    def multiply(self, x, y):
        """
        乘法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 和 y 的乘积
        """
        self._update_stats()
        print(f"服务器执行: {x} * {y}")
        return x * y
    
    def divide(self, x, y):
        """
        除法运算
        
        参数:
            x: 第一个操作数
            y: 第二个操作数
            
        返回:
            x 除以 y 的商
            
        异常:
            如果 y 为 0，将抛出 ValueError 异常
        """
        self._update_stats()
        if y == 0:
            print(f"服务器执行: {x} / {y} - 错误: 除数不能为零")
            raise ValueError("除数不能为零")
        
        print(f"服务器执行: {x} / {y}")
        return x / y
    
    def power(self, base, exponent):
        """
        计算幂
        
        参数:
            base: 底数
            exponent: 指数
            
        返回:
            base 的 exponent 次方
        """
        self._update_stats()
        print(f"服务器执行: {base} ^ {exponent}")
        return base ** exponent
    
    def fibonacci(self, n):
        """
        计算斐波那契数列第 n 项
        
        参数:
            n: 位置 (0-indexed)
            
        返回:
            第 n 项斐波那契数
        """
        self._update_stats()
        print(f"服务器执行: fibonacci({n})")
        
        if n < 0:
            raise ValueError("输入必须是非负整数")
        
        if n <= 1:
            return n
        
        a, b = 0, 1
        for _ in range(2, n + 1):
            a, b = b, a + b
        return b
    
    def generate_random_numbers(self, n, min_val=0, max_val=100):
        """
        生成指定范围内的随机数列表
        
        参数:
            n: 要生成的随机数数量
            min_val: 最小值（默认为0）
            max_val: 最大值（默认为100）
            
        返回:
            随机数列表
            
        这个方法演示了 ZeroRPC 在传递复杂数据结构时的能力。
        """
        self._update_stats()
        print(f"服务器执行: 生成 {n} 个随机数，范围 [{min_val}, {max_val}]")
        
        return [random.randint(min_val, max_val) for _ in range(n)]
    
    def stream_counter(self, start=0, end=10, delay=1.0):
        """
        流式返回从 start 到 end 的计数
        
        参数:
            start: 起始值
            end: 结束值
            delay: 每个数之间的延迟（秒）
            
        返回:
            生成器，产生从 start 到 end 的数字
            
        这个方法演示了 ZeroRPC 的流式响应功能。
        客户端可以在服务器生成完整响应前开始接收和处理数据。
        """
        self._update_stats()
        print(f"服务器开始流式传输，从 {start} 到 {end}，延迟 {delay}秒")
        
        for i in range(start, end + 1):
            print(f"服务器流式传输: {i}")
            yield i
            time.sleep(delay)  # 模拟耗时操作或实时数据产生
    
    def get_stats(self):
        """
        获取服务统计信息
        
        返回:
            包含服务统计信息的字典
        """
        stats = {
            "calculation_count": self.calculation_count,
            "last_accessed": self.last_accessed.isoformat(),
            "uptime_seconds": (datetime.now() - self.last_accessed).total_seconds(),
            "server_time": datetime.now().isoformat()
        }
        print(f"服务器统计信息: {stats}")
        return stats
    
    def _update_stats(self):
        """
        更新服务统计信息
        
        注意: 以下划线开头的方法不会作为 RPC 方法暴露
        """
        self.calculation_count += 1
        self.last_accessed = datetime.now()

def run_server():
    """
    启动 ZeroRPC 服务器
    
    创建 ZeroRPC 服务器，绑定到指定地址，并开始接受连接。
    """
    server = zerorpc.Server(AdvancedMathService())
    
    # 配置心跳选项（可选）
    # - heartbeat: 客户端和服务器之间的心跳间隔（默认5秒）
    # - passive_heartbeat: 是否使用被动心跳模式
    server.heartbeat = 10  # 10秒心跳
    
    # 绑定服务器到指定地址
    # ZeroMQ 支持多种传输协议:
    # - tcp://host:port - TCP 连接
    # - ipc://path - 进程间通信
    # - inproc://name - 进程内通信
    server.bind(SERVER_ADDR)
    
    print(f"ZeroRPC 服务器已启动，地址: {SERVER_ADDR}")
    print("使用 Ctrl+C 结束服务")
    
    try:
        # 阻塞运行服务器
        server.run()
    except KeyboardInterrupt:
        pass
    finally:
        server.close()
        print("服务器已关闭")

def run_client():
    """
    运行 ZeroRPC 客户端
    
    连接到服务器并调用远程方法，展示不同的调用模式和功能。
    """
    # 等待一段时间，让服务器有机会启动
    time.sleep(1)
    
    try:
        # 创建 ZeroRPC 客户端
        # ZeroRPC 客户端自动处理连接、消息序列化、心跳等
        client = zerorpc.Client()
        
        # 设置超时（秒）
        client.timeout = 10
        
        # 连接到服务器
        client.connect(SERVER_ADDR)
        
        # 使用分隔线使输出更清晰
        def section(title):
            print(f"\n{'=' * 5} {title} {'=' * 5}")
        
        section("基础算术运算")
        # 调用远程方法
        # ZeroRPC 客户端将自动序列化调用，发送到服务器
        # 然后等待响应，并返回反序列化后的结果
        print(f"10 + 20 = {client.add(10, 20)}")
        print(f"50 - 30 = {client.subtract(50, 30)}")
        print(f"7 * 8 = {client.multiply(7, 8)}")
        print(f"100 / 4 = {client.divide(100, 4)}")
        
        section("高级数学运算")
        print(f"2^10 = {client.power(2, 10)}")
        print(f"斐波那契数列第 10 项 = {client.fibonacci(10)}")
        
        section("异常处理")
        # ZeroRPC 将服务器异常传播到客户端
        try:
            result = client.divide(1, 0)
        except zerorpc.exceptions.RemoteError as e:
            # RemoteError 包含原始异常的类型和消息
            print(f"捕获远程异常: {e.__class__.__name__}: {e}")
        
        section("复杂数据结构")
        # ZeroRPC 可以传输复杂的数据结构
        # MessagePack 支持: 原始值、字符串、数组和映射
        random_numbers = client.generate_random_numbers(5, 1, 100)
        print(f"5个随机数: {random_numbers}")
        
        # 可以传递和接收更复杂的嵌套结构
        numbers = client.generate_random_numbers(3, 1, 10)
        squares = [client.power(n, 2) for n in numbers]
        print(f"数字: {numbers}")
        print(f"平方: {squares}")
        
        section("流式响应")
        # ZeroRPC 支持流式响应（生成器）
        # 服务器可以逐步产生结果，客户端可以逐步处理
        print("开始接收流式数据:")
        for i in client.stream_counter(1, 5, 0.5):
            print(f"收到流数据: {i}")
        
        section("服务统计")
        # 获取服务器统计信息
        stats = client.get_stats()
        print("服务器统计信息:")
        for key, value in stats.items():
            print(f"  {key}: {value}")
        
    except zerorpc.exceptions.TimeoutExpired:
        print("连接超时: 无法连接到服务器或操作超时")
    except zerorpc.exceptions.LostRemote:
        print("连接丢失: 服务器可能已关闭或网络中断")
    except zerorpc.exceptions.RemoteError as e:
        print(f"远程错误: {e}")
    except Exception as e:
        print(f"错误: {e}")
    finally:
        # 关闭客户端连接
        if 'client' in locals():
            client.close()

def main():
    """
    主函数，协调服务器和客户端
    """
    # 在单独的线程中启动服务器
    server_thread = threading.Thread(target=run_server)
    server_thread.daemon = True  # 将线程设为守护线程，这样主线程结束时它也会结束
    server_thread.start()
    
    # 等待服务器启动
    time.sleep(1)
    
    # 在主线程中运行客户端
    print("运行 ZeroRPC 客户端...")
    run_client()
    
    # 等待用户输入，让服务器继续运行
    print("\n按 Enter 键退出...")
    input()

if __name__ == "__main__":
    main()