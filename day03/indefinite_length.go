package main

import "fmt"

/**
不定长参数
*/

/**
如果函数的最后一个参数是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0
*/

func myFunc(a int, b string, arg ...int) {
	if len(arg) == 0 {
		return
	}
	min := arg[0]

	for _, v := range arg {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
	fmt.Println(arg)
}

func main() {
	myFunc(10, "abc", 10, 20, 30)
}
