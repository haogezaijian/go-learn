package main

import (
	"fmt"
	"io"
	"log"
)

/**
defer  关键字
*/

func main() {
	function1()
	a()
	func1("go")

}

func function1() {
	fmt.Println("start")
	defer function2()
	fmt.Println("end")

}

func function2() {
	fmt.Println("defer")
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Print("func1(%q) = %d,%v", s, n, err)
	}()
	return 7, io.EOF
}
