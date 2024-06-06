package main

import fm "fmt"

/**
go 函数
*/
//定义一个无参数，无返回值的函数
func fun1() {
	fm.Println("无参数，无返回值的函数")
}

// 定义一个有参数的函数
func fun2(x, y int) {
	fm.Println(x + y)
}

// 定义一个有参数有返回值的函数
func fun3(x, y int) int {
	return x + y
}

func main() {
	fun1()
	fun2(10, 20)
	fm.Println(fun3(10, 20))
}
