package main

import (
	"fmt"
	"time"
)

/**
日期和时间
*/

/**
time 包为我们提供了一个数据类型 time.Time（作为值使用）以及显示和测量时间和日期的功能函数。
*/

func main() {
	//获取当前时间
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
}
