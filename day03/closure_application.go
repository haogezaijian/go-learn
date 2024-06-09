package main

import "fmt"

/**
闭包应用
*/

func main() {
	p2 := Add2()
	fmt.Println(p2(3))

	p3 := Adder(2)
	fmt.Println(p3(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

//可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点
