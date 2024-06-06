package main

import "fmt"

/**
引用类型
*/

/**
更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。

一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。
*/

/*
*
当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。

如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响
*/
func main() {
	//在 Go 语言中，指针属于引用类型，其它的引用类型还包括 slices，maps和 channel。

	a := 10
	ptr := &a

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 20
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	m := map[string]int{"a": 10, "b": 20}
	fmt.Println(m)
	fmt.Println(m["a"])

	m1 := m
	m1["b"] = 100

	fmt.Println(m)
	fmt.Println(m1)
}
