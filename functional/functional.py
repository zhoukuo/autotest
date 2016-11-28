# coding=utf-8


def add(a, b):
    return a+b


def sub(a, b):
    return a-b


def mul(a, b):
    return a*b


def div(a, b):
    return a/b


class Interger(object):

    def __init__(self, i):
        super(Interger, self).__init__()
        self.value = i

    def add(self, a):
        self.value = self.value + a.value
        return self

    def sub(self, a):
        self.value = self.value - a.value
        return self

    def mul(self, a):
        self.value = self.value * a.value
        return self

    def div(self, a):
        self.value = self.value / a.value
        return self


def main():

    # 1. 有这样一个数学表达式：(1 + 2) * 3 - 4

    # 2. 传统的过程式编程，可能这样写：
    a = 1 + 2
    b = a * 3
    c = b - 4
    result = c
    print result

    # 函数式编程要求使用函数，我们可以把运算过程定义为不同的函数，然后写成下面这样
    a = 1
    b = 2
    c = 3
    d = 4
    result = sub(mul(add(a, b), c), d)
    print result

    # 对它进行变形，不难得到另一种写法，这基本就是自然语言的表达
    a = Interger(1)
    b = Interger(2)
    c = Interger(3)
    d = Interger(4)
    a.add(b).mul(c).sub(d)
    print a.value


if __name__ == '__main__':
    main()
