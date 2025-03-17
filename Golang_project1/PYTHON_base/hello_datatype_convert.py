# python的数据类型转换
#     # int() 转换为整数
#     # float() 转换为浮点数
#     # bool() 转换为布尔类型
#     # str() 转换为字符串
#     # list() 转换为列表
#     # tuple() 转换为元组
#     # dict() 转换为字典
#     # set() 转换为集合
#     # NoneType() 转换为空类型
def hello_datatype_convert():
    print("---------------------------------打印数据类型转换--------开始-------------------------")
    # int() 转换为整数
    a = 1.1
    print("变量a=1.1,a的数据类型：", type(a), "a的值：", a, "占用字节数：", a.__sizeof__())
    a = int(a)
    print("变量a=1.1,a的数据类型：", type(a), "a的值：", a, "占用字节数：", a.__sizeof__())
    # float() 转换为浮点数
    b = 1
    print("变量b=1,b的数据类型：", type(b), "b的值：", b, "占用字节数：", b.__sizeof__())
    b = float(b)
    print("变量b=1,b的数据类型：", type(b), "b的值：", b, "占用字节数：", b.__sizeof__())
    # bool() 转换为布尔类型 0为False,非0为True
    c = 0
    print("变量c=0,c的数据类型：", type(c), "c的值：", c, "占用字节数：", c.__sizeof__())
    c = bool(c)
    print("变量c=0,c的数据类型：", type(c), "c的值：", c, "占用字节数：", c.__sizeof__())
    # str() 转换为字符串
    d = 1
    print("变量d=1,d的数据类型：", type(d), "d的值：", d, "占用字节数：", d.__sizeof__())
    d = str(d)
    print("变量d=1,d的数据类型：", type(d), "d的值：", d, "占用字节数：", d.__sizeof__())
    # list() 转换为列表
    e = (1, 2, 3, 4, 5)
    print("变量e=(1, 2, 3, 4, 5),e的数据类型：", type(e), "e的值：", e, "占用字节数：", e.__sizeof__())
    e = list(e)
    print("变量e=(1, 2, 3, 4, 5),e的数据类型：", type(e), "e的值：", e, "占用字节数：", e.__sizeof__())
    # tuple() 转换为元组
    f = [1, 2, 3, 4, 5]
    print("变量f=[1, 2, 3, 4, 5],f的数据类型：", type(f), "f的值：", f, "占用字节数：", f.__sizeof__())
    f = tuple(f)
    print("变量f=[1, 2, 3, 4, 5],f的数据类型：", type(f), "f的值：", f, "占用字节数：", f.__sizeof__())
    # dict() 转换为字典
    g = [("name", "张三"), ("age", 18)]
    print("变量g=[('name', '张三'), ('age', 18)],g的数据类型：", type(g), "g的值：", g, "占用字节数：", g.__sizeof__())
    g = dict(g)
    print("变量g=[('name', '张三'), ('age', 18)],g的数据类型：", type(g), "g的值：", g, "占用字节数：", g.__sizeof__())
    # set() 转换为集合
    h = [1, 2, 3, 4, 5]
    print("变量h=[1, 2, 3, 4, 5],h的数据类型：", type(h), "h的值：", h, "占用字节数：", h.__sizeof__())
    h = set(h)
    print("变量h=[1, 2, 3, 4, 5],h的数据类型：", type(h), "h的值：", h, "占用字节数：", h.__sizeof__())
    # NoneType() 转换为空类型
    i = 1
    print("变量i=1,i的数据类型：", type(i), "i的值：", i, "占用字节数：", i.__sizeof__())
    i = None
    print("变量i=1,i的数据类型：", type(i), "i的值：", i, "占用字节数：", i.__sizeof__())
    print("---------------------------------打印数据类型转换--------结束-------------------------")