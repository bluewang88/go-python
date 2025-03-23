#编写一个计算器




def add(x, y):
    return x + y

def sub(x, y):
    return x - y

def mul(x, y):
    return x * y

def div(x, y):
    if y == 0:
        raise ValueError("Cannot divide by zero")
    return x / y

op_dict = {
    "+": add,
    "-": sub,
    "*": mul,
    "/": div
}

a = int(input("Enter a number: "))
b = int(input("Enter another number: "))
op = input("Enter an operator: ")


#函数可以当普通变量使用、还可以当返回值，这个特性就是一等公民特性
func = op_dict[op]
result = func(a, b)
print("The result is:", result)
    