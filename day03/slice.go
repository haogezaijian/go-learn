package main

import (
	"fmt"
	"sort"
)

/**
切片
*/

//切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量

func main() {
	s := []int{3, 2, 1}
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s1 := s //切片是引用类型

	s1[0] = 10
	fmt.Println(s)
	fmt.Println(s1)

	//使用make创建切片
	s2 := make([]int, 10) //len =10  cap = 10
	fmt.Println(s2)

	s3 := make([]int, 10, 20) //len = 10  cap = 20
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	s4 := s3
	s4[0] = 100
	fmt.Println(s3)
	fmt.Println(s4)

	//切片的复制和追加
	sliceCopy()
	sliceAppend()
	//搜索切片中的元素，以及排序切片

	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	//想要在数组或切片中搜索一个元素，该数组或切片必须先被排序
	index := sort.SearchInts(s, 10)
	fmt.Println(index) //返回索引
}

func sliceCopy() {
	source := []int{1, 2, 3, 4, 5}
	target := make([]int, 10)

	n := copy(target, source) //返回值为复制元素个数
	fmt.Println(target)
	fmt.Println(n)
}

func sliceAppend() {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1, 6, 7, 8, 9)
	fmt.Println(s1)
}
