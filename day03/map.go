package main

import (
	"fmt"
)

/**
map
*/

/**
map 是引用类型，可以使用如下声明：
var map1 map[keytype]valuetype
var map1 map[string]int
*/

/**
key 可以是任意可以用 == 或者！= 操作符比较的类型，
比如 string、int、float。所以切片和结构体不能作为 key
但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，
这样可以通过结构体的域计算出唯一的数字或者字符串的 key
*/

func main() {
	mapList := map[string]int{"one": 1, "tow": 2}
	m := make(map[string]int)
	fmt.Println(mapList)
	fmt.Println(mapList["one"])
	fmt.Println(m)
	/**
	不要使用 new，永远用 make 来构造 map
	 如果你错误的使用 new () 分配了一个引用对象，
	你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址
	*/

	//value可以是任意类型
	mf := map[int]func() int{
		1: func() int {
			return 1
		},
		2: func() int {
			return 2
		},
		3: func() int {
			return 3
		},
	}

	fmt.Println(mf[1]())

	//判断key是否存在
	v, has := mapList["tow"]
	if has {
		fmt.Println(v)
	}

	//简化写法
	if _, ok := mapList["tow"]; ok {
		fmt.Println(mapList["tow"])
	}

	//删除map中的元素
	delete(mapList, "now")

	for k, v := range mapList {
		fmt.Println(k)
		fmt.Println(v)
	}

}
