package main

import "fmt"

/**
值类型
*/

func main() {

	//所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：
	a := 10
	b := 10.0
	c := "str"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//数组和结构体类型等复合类型也是值类型
	arr := [3]int{10, 20, 30}
	fmt.Println(arr)

	type str struct {
		age  int
		name string
	}

	s := str{10, "haoge"}
	fmt.Println(s)

	//当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
	j := 10
	i := j
	fmt.Println(i)
	fmt.Println(j)

	i = 20
	fmt.Println(i)
	fmt.Println(j)

	arr1 := arr
	arr1[1] = 100
	fmt.Println(arr)
	fmt.Println(arr1)
}
