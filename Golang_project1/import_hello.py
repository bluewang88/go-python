#导入hello.py模块中的hello函数
#导入hello.py其中
# print("hello", end="")  # hello
# print(" world form hello.py global")  # world
#是在全局作用于中的，所以在导入hello.py模块时会直接执行这两行代码，不用等到下面的hello()函数调用时才执行
import hello
# # #输出分割线
print("------------------import hello 成功----------------------")
# # #调用hello.py模块中的hello函数
# print("调用hello.py模块中的hello函数,使用hello.hello()，输出：")
hello.hello()

# # #输出分割线
print("--------------------------------------------------")


# #导入hello.py模块中的hello函数并重命名为h
from hello import hello as h

#输出分割线
print("--------------------from hello import hello as h 成功------------------------")
# # #调用hello.py模块中的hello函数
print("导入hello.py模块中的hello函数并重命名为h,使用h()，输出：")
h()