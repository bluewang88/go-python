#！/usr/bin/env python3
# -*- coding: utf-8 -*-


def hello_basic_datatype():
    '''
    打印python的基本数据类型
    '''
    # type() 查看变量数据类型   
    # isinstance() 判断变量数据类型
    # 基本数据类型
    #int
    # 整数类型
    a = 1
    print("变量a=1,a的数据类型：", type(a), "a的值：", a, "占用字节数：", a.__sizeof__())
    # 整数类型可以表示为二进制、八进制、十进制和十六进制
    # # 二进制表示为0b开头
    # b = 0b1010
    # print("变量b=0b1010,b的数据类型：", type(b), "b的值：", b)
    # # 八进制表示为0o开头
    # c = 0o123
    # print("变量c=0o123,c的数据类型：", type(c), "c的值：", c)   

    #float
    # 浮点数类型
    b = 1.1
    print("变量b=1.1,b的数据类型：", type(b), "b的值：", b, "占用字节数：", b.__sizeof__())

    #bool
    # 布尔类型
    c = True
    print("变量c=True,c的数据类型：", type(c), "c的值：", c, "占用字节数：", c.__sizeof__())

    # complex 复数
    d = 1+2j
    print("变量d=1+2j,d的数据类型：", type(d), "d的值：", d, "占用字节数：", d.__sizeof__())

    #string 字符串
    d = "Hello, World!"
    print("变量d=Hello, World!,d的数据类型：", type(d), "d的值：", d, "占用字节数：", d.__sizeof__())

    #list 列表
    # 列表是可变的序列类型，可以添加、删除、修改元素
    e = [1, 2, 3, 4, 5]
    print("变量e=[1, 2, 3, 4, 5],e的数据类型：", type(e), "e的值：", e, "占用字节数：", e.__sizeof__())

    #tuple 元组
    # 元组是不可变的序列类型，一旦创建，元素不能被修改  
    f = (1, 2, 3, 4, 5)
    print("变量f=(1, 2, 3, 4, 5),f的数据类型：", type(f), "f的值：", f, "占用字节数：", f.__sizeof__())

    #dict 字典
    # 字典是键值对的集合，键值对之间用逗号分隔，键和值之间用冒号分隔
    g = {"name": "张三", "age": 18}
    print("变量g={'name': '张三', 'age': 18},g的数据类型：", type(g), "g的值：", g, "占用字节数：", g.__sizeof__())

    #set 集合
    # 集合是不可变的无序的唯一元素的集合
    h = {1, 2, 3, 4, 5}
    print("变量h={1, 2, 3, 4, 5},h的数据类型：", type(h), "h的值：", h, "占用字节数：", h.__sizeof__())

    #None 空类型
    i = None
    print("变量i=None,i的数据类型：", type(i), "i的值：", i, "占用字节数：", i.__sizeof__())        


# 整数类型
# 浮点数类型
# 布尔类型
# 字符串类型
# 列表类型
# 元组类型
# 字典类型
# 集合类型
# 空类型
