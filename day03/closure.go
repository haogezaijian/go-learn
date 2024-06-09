package main

import "fmt"

/**
闭包
*/

func main() {
	func() {
		sum := 0
		for i := 1; i < 100; i++ {
			sum += i
		}
	}()
	//表示参数列表的第一对括号必须紧挨着关键字 func，因为匿名函数没有名称。
	//花括号 {} 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。

	f()

	fv := func() { fmt.Println("hello,world") }
	fv()

}

// 将匿名函数复制给变量并进行调用
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Println(i) }
		g(i)
	}
}

/**
defer 语句和匿名函数

关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值。

匿名函数还可以配合 go 关键字来作为 goroutine 使用
*/
