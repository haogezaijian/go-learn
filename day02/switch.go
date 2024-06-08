package main

import (
	"fmt"
	"math/rand"
)

/**
switch 结构
*/

func main() {
	num := 60

	/**
	每一个case分支都是唯一的，从上之下主逐一执行，直到匹配或进入default条件为止。
	一旦成功的匹配到某个分支，在执行完相应的代码后就会退出整个switch代码块，不需要特别使用break
	因此，程序不会自定去执行下一个分支的代码。如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的
	*/
	switch num { //num可以是任何类型
	case 59:
		fmt.Println("不及格") //无需使用break退出
		fallthrough        //继续执行后续的分支
	case 60:
		fmt.Println("及格")
	default:
		fmt.Println("....")
	}

	/**
	switch 语句的第二种形式是不提供任何被判断的值（实际上默认为判断是否为 true），
	然后在每个 case 分支中进行测试不同的条件。当任一分支的测试结果为 true 时，该分支的代码会被执行。
	*/
	score := rand.Intn(100)
	switch {
	case score > 60:
		fmt.Println("及格")
	case score < 60:
		fmt.Println("不及格")
	default:
		fmt.Println("....")
	}

	/**
	switch 语句的第三种形式是包含一个初始化语句
	*/
	switch res := len("abcdefg"); {
	case res > 10:
		fmt.Println(10)
	case res < 10:
		fmt.Println(0)
	default:
		fmt.Println(-1)

	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		break
		fallthrough //加了break，fallthrough会失效
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
