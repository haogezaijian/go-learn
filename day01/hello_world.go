package main

import fm "fmt" //给引入的包起别名

//导入多个包
/*
	type1
	import "fmt"
	import "os"
*/
/*
	type2
	import "fmt";import "os"
*/
/**
	type3
	import(
	"fmt"
	"os"
)
*/
/*
	type4
	import("fmt";"os")
*/
func main() {
	//导入包如果未使用会报错
	fm.Println("Hello World")
}
