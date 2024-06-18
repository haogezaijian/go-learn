package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	pracise01()
	pracise02()
	pracise03()
	pracise04()
}

/*
*
使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来
*/
func pracise01() {
	//创建一个双向链表
	l := list.New()

	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	for elem := l.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value)
	}

}

/*
*
通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
*/
func pracise02() {
	sizeof := unsafe.Sizeof(float64(10))
	fmt.Println(sizeof)
}

/*
*
定义结构体 Address 和 VCard，后者包含一个人的名字、
地址编号、出生日期和图像，试着选择正确的数据类型。
构建一个自己的 vcard 并打印它的内容。
*/
type Address struct {
	addr string
}

type Vcard struct {
	name     string
	birthDay string
	img      string
	addr     *Address
}

func pracise03() {
	vcard := new(Vcard)
	address := new(Address)
	address.addr = "杭州"
	vcard.addr = address
	vcard.name = "Haoge"
	vcard.birthDay = "2020-01-01"
	vcard.img = "1.jpg"
	fmt.Println(vcard.name)
	fmt.Println(vcard.birthDay)
	fmt.Println(vcard.img)
	fmt.Println(vcard.addr.addr)
}

/*
*
定义一个 Rectangle 结构体，它的长和宽是 int 类型，并定义方法 Area() 和 Primeter()，然后进行测试。
*/
type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Width)
}
func pracise04() {
	rect := &Rectangle{10, 5}
	fmt.Println(rect.Area())
	fmt.Println(rect.Perimeter())
}
