package main

import (
	"fmt"
	"go-learn-hao/structPack"
	"reflect"
	"strings"
)

/**
结构和方法
*/

/*
*
结构体定义

	type identifier struct {
	    field1 type1
	    field2 type2
	    ...
	}
*/
type struct1 struct {
	i1 int
	f1 float32
	s1 string
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
}

// 结构体工厂
type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

// 带标签的结构体
type TagType struct {
	filed1 bool   "An important answer"
	filed2 string "The name of the thing"
	filed3 int    "How nuch there are"
}

// 匿名字段和内嵌结构体
type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.s1 = "haoge"

	fmt.Println(ms.s1)
	fmt.Println(ms.f1)
	fmt.Println(ms.i1)
	fmt.Println(ms)

	//方式一
	var pers1 Person
	pers1.firstName = "Haoge"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Println(pers1.firstName)
	fmt.Println(pers1.lastName)
	//方式二
	pers2 := new(Person)
	pers2.firstName = "Haoge"
	pers2.lastName = "Woodward"
	upPerson(pers2)
	fmt.Println(pers2.firstName)
	fmt.Println(pers2.lastName)
	//方式三
	pers3 := &Person{"Haoge", "Woodward"}
	upPerson(pers3)
	fmt.Println(pers3.firstName)
	fmt.Println(pers3.lastName)

	f := NewFile(10, "hao")
	fmt.Println(f)
	//使用自定义包中的结构体
	sp := new(structPack.ExpStruct)
	sp.Mf1 = 10.0
	sp.Mi1 = 10
	fmt.Println(sp)

	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {

		refTag(tt, i)
	}

	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println(outer2)
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixFiled := ttType.Field(ix)
	fmt.Println(ixFiled.Tag)
}
