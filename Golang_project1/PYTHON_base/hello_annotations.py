import typing

#函数的__annotations__和typing.get_type_hints()方法都可以获取函数的参数和返回值类型
#函数的__annotations__属性是一个字典，键是参数名，值是参数类型
#typing.get_type_hints()方法是一个函数，接收一个函数作为参数，返回一个字典，键是参数名，值是参数类型
# #__annotations__叫做注解，typing.get_type_hints()叫做类型提示
def validate_input(obj, **kwargs):
    hints = typing.get_type_hints(obj)
    for para_name, para_type in hints.items():
        if para_name not in kwargs:
            raise TypeError(f"{para_name} is missing.")
            # raise的意思你可以理解为抛出异常，这里抛出的异常是TypeError
        if not isinstance(kwargs[para_name], para_type):
            raise TypeError(f"{para_name} must be {para_type}.")
    print(kwargs)


#python装饰器
#装饰器的作用是在不改变原函数的情况下，对函数的功能进行扩展
#装饰器的语法糖@，在函数定义的上一行加上@装饰器的名字
#装饰器的定义：装饰器是一个函数，接收一个函数作为参数，返回一个函数

def type_check(func):
    #wrapper函数的作用是接收原函数的参数，对参数进行校验，然后调用原函数
    @wraps(func)

    def wrapped_decorator(*args, **kwargs):
        func_args = getfullargspec(func)[0]
        kwargs.update(dict(zip(func_args, args))) # 将args转换为字典，然后更新到kwargs中
        validate_input(func, **kwargs) # 校验参数类型
        return func(**kwargs) # 调用原函数
    
    return wrapped_decorator # 返回wrapper函数
    # def wrapper(**kwargs):
    #     validate_input(func, **kwargs)
    #     return func(**kwargs)
    # return wrapper




@type_check
def hello(name: str, age: int) -> str:
    return f"Hello, {name}, you are {age} years old."


print(hello("Alice", 30))
print(hello.__annotations__)
print(typing.get_type_hints(hello))


    # Removed the second validation loop as it was unnecessary
validate_input(hello, name="Alice", age=30)  # Changed age back to int
validate_input(hello, name="Alice", age="30")  # Changed age back to str
# Output:  
# Hello, Alice, you are 30 years old.
# {'name': <class 'str'>, 'age': <class 'int'>}
# {'name': <class 'str'>, 'age': <class 'int'>}
# Traceback (most recent call last):
#   File "hello_annotations.py", line 17, in <module>
#     vaildate_input(hello, name="Alice", age="30")  # Changed age back to str
#   File "hello_annotations.py", line 13, in vaildate_input
#     raise TypeError(f"{key} must be {hints[key]}.")
# TypeError: age must be <class 'int'>.