'''
这个Python文件主要有以下作用：

1. 作为一个Python模块示例
   - 演示了模块导入机制
   - 展示了全局代码与函数定义的区别

2. 定义了一个可复用的hello()函数
   - 该函数在被调用时会打印"hello world from python function"
   
3. 演示Python模块的双重用途
   - 可以作为独立脚本运行
   - 也可以被其他Python脚本导入使用
   
4. 演示了Python的__name__机制
   - 通过判断__name__是否为"__main__"来区分直接运行和被导入的情况

5. 展示了Python注释的使用方式
   - 单行注释(#)
   - 文档字符串(docstring)
'''

# 这是一个块注释的例子
# 使用多个以#开头的行
# 每行前面都需要加上#符号
print("hello", end="")  # hello
print(" world form hello.py global")  # world


def hello():
    """
    这是hello函数的文档字符串
    可以通过__doc__属性访问
    """
    print("hello", end="")
    print("hello world from python function")


# 通过__name__属性判断是否是主程序
# 如果是主程序则执行hello函数
# 如果是模块则不执行hello函数
if __name__ == "__main__":
    hello()

