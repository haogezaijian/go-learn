package main

/**
Go 语言不存在隐式类型转换，因此所有的转换都必须显式说明
*/

import "fmt"

func main() {
	var a = 10.0
	var b int

	b = int(a)
	fmt.Println(b)

	//当从一个取值范围较大的转换到取值范围较小的类型时（例如将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失（截断）的情况。
	var c = 10.02
	var d int
	d = int(c)
	fmt.Println(d) //精度丢失
}
