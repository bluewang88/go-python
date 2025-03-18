#海象运算符
#在python3.10中新增的海象运算符，用于简化if语句
#语法：:=
#作用：在判断条件时，直接在条件表达式中进行变量赋值
#示例：
#if (n := len(a)) > 10:
#    print(f"List is too long ({n} elements, expected <= 10)")
#else:
#    print(f"List is ok")

my_list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]

# 不使用海象运算符
n = len(my_list)
if n > 10:
    print(f"List is too long ({n} elements)")

# 使用海象运算符
if (n := len(my_list)) > 10:
    print(f"List is too long ({n} elements)")


