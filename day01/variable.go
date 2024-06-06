package main

import "fmt"

//声明变量的一般形式是使用 var 关键字

var a int = 10 //声明一个int类型的变量

// 这种写法一般用于声明全局变量
var (
	b   int    = 20
	c   bool   = false
	str string = "str"
)

// 声明局部变量,函数体内
func getNum() int {
	a := 10 //优先使用函数体内的局部变量
	return a
}

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(str)

	fmt.Println(getNum())
}
