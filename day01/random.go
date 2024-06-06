package main

import (
	"fmt"
	"math/rand"
)

/**
随机数
*/

func main() {
	a := rand.Int()
	fmt.Println(a)

	f := rand.Float32()
	fmt.Println(f)

	b := rand.Intn(5) //生产0到4的随机数
	fmt.Println(b)
}
