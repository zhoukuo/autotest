package main

import (
	"fmt"
)

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {
	return a / b
}

type Interger int

func (a Interger) Add(b Interger) Interger {
	return a + b
}

func (a Interger) Sub(b Interger) Interger {
	return a - b
}

func (a Interger) Mul(b Interger) Interger {
	return a * b
}

func (a Interger) Div(b Interger) Interger {
	return a / b
}

func main() {
	// 1. 有这样一个数学表达式：(1 + 2) * 3 - 4

	// 2. 传统的过程式编程，可能这样写：
	var a1 = 1 + 2
	var b1 = a1 * 3
	var c1 = b1 - 4
	fmt.Println(c1)

	// 函数式编程要求使用函数，我们可以把运算过程定义为不同的函数，然后写成下面这样
	var a2, b2, c2, d2 = 1, 2, 3, 4
	fmt.Println(Sub(Mul(Add(a2, b2), c2), d2))

	// 对它进行变形，不难得到另一种写法，这基本就是自然语言的表达
	var a3, b3, c3, d3 Interger = 1, 2, 3, 4
	fmt.Println(a3.Add(b3).Mul(c3).Sub(d3))
}
