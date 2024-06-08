package main

import "fmt"

/**
指针
*/

/*
*
一个指针变量可以指向任何一个值的内存地址
它指向那个值的内存地址，在 32 位机器上占用 4 个字节，
在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关
*/
func main() {
	num := 10
	ptr := &num

	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 20
	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	//无法获取一个文字或常量的地址
	const i = 5
	//ptr1 := &i   //error:cannot take the address of i
	//ptr2 := &10  //error:cannot take the address of 10
	//ptr3 := &"hello"

}
