package main

import "fmt"

/**
编写一个函数，接受两个整数，然后返回他们的和，积与差。
编写两个版本，一个是非命名返回值，一个是命名返回值
*/

func sum1(x, y int) int {
	sum := x + y
	return sum
}
func sum2(x, y int) (sum int) {
	sum = x + y
	return
}

func product1(x, y int) int {
	product := x * y
	return product
}

func product2(x, y int) (product int) {
	product = x * y
	return
}

func sub1(x, y int) int {
	sub := x - y
	return sub
}

func sub2(x, y int) (sub int) {
	sub = x - y
	return
}

func main() {
	fmt.Println(sub1(10, 20))
	fmt.Println(sub2(10, 20))

	fmt.Println(sum1(10, 20))
	fmt.Println(sum2(10, 20))

	fmt.Println(product1(10, 20))
	fmt.Println(product2(10, 20))

}
