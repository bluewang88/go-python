#python的类型声明
def hello_datatype_claim():
    '''
    python的类型声明
    '''
    a: int = 1
    b: str = 'hello'
    c: bool = True

    print(a, b, c)
    print(type(a), type(b), type(c))
    print(a.__sizeof__(), b.__sizeof__(), c.__sizeof__())

    # python3.10中新增的类型声明
    d: int
    d = 10
    print(d)
    print(type(d))
    print(d.__sizeof__()) 

    