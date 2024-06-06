package main

import "fmt"

/*
*
go 数据类型
*/
func main() {
	//基本数据类型
	var a int //使用var声明的变量会自动初始化为该类型的零值
	fmt.Println(a)
	var b float64 //变量声明后没有使用也会报错
	var c float32
	fmt.Println(b)
	fmt.Println(c)
	var d bool
	fmt.Println(d)
	var e string
	fmt.Println(e)
	//结构化类型
	var f struct { //结构体
		a int
		b string
	}
	fmt.Println(f)
	var g [5]int //数组
	fmt.Println(g)

	var h []int //切片
	fmt.Println(h)

	var i map[string]int //map
	fmt.Println(i)
}
