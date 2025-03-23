#python 的函数参数和返回值类型声明

import typing
import hello_annotations
def add(a: int, b: int) -> int:
    hello_annotations.validate_input(add, a=a, b=b)
    return a + b


def add2(a: int, b: float=3.5) -> int:
    return a + b


if __name__ == '__main__':
    print(add(1, 2)) #程序直接运行，不导入到其他文件中运行