# python的鸭子类型

# 鸭子类型是指一种编程风格，强调对象的行为而不是对象的类型。
# 在Python中，如果一个对象具有某种方法或属性，那么我们就可以认为这个对象是某种类型的实例。
# 这种方式使得Python的代码更加灵活和动态。
# 例如，我们可以定义一个函数，它接受一个对象作为参数，然后调用这个对象的方法，而不关心这个对象的具体类型。
# 只要这个对象具有我们需要的方法或属性，我们就可以使用它。
# 下面是一个简单的示例，演示了鸭子类型的概念：
class Duck:
    def quack(self):
        print("Quack!")
    def swim(self):
        print("Duck is swimming!")
class Dog:
    def bark(self):
        print("Woof!")
    def swim(self):
        print("Dog is swimming!")
def make_it_quack(duck):
    duck.quack()
    duck.swim()
# 创建一个鸭子对象
duck = Duck()
# 创建一个狗对象
dog = Dog()
# 调用make_it_quack函数，传入鸭子对象
make_it_quack(duck)
# 调用make_it_quack函数，传入狗对象
make_it_quack(dog)
# 运行结果：
# Quack!
# Duck is swimming!
# Quack!
# Dog is swimming!
# 这个例子中，我们定义了一个make_it_quack函数，它接受一个对象作为参数。