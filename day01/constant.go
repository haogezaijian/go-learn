package main

import "fmt"

/**
常量使用关键字 const 定义，用于存储不会改变的数据。

存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
*/

const A = 10
const B = "Hello"
const FLAG = true

//常量的值必须是能够在编译时就能够确定的

const c1 = 2 + 3

// const c2 = getNumber() //错误 getNumber()在编译期间属于未知
const c3 = len("abc") //正确 len()是内置方法

// 常量用作枚举
const (
	UNKNOWN = 0
	FEMALE  = 1
	MALE    = 2
)

func getNumber() int {
	return 10
}
func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(FLAG)

	//A = 20  //导致报错 cannot assign to A
	fmt.Println(A)
}
