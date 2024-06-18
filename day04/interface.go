package main

import "fmt"

/**
接口
*/

/*
*
Go 语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念
但是 Go 语言里有非常灵活的 接口 概念，通过它可以实现很多面向对象的特性。
接口提供了一种方式来 说明 对象的行为：如果谁能搞定这件事，它就可以用在这儿。
接口定义了一组方法（方法集），但是这些方法不包含（实现）代码
它们没有被实现（它们是抽象的）。接口里也不能包含变量。

通过如下格式定义接口：

	type Namer interface {
	    Method1(param_list) return_type
	    Method2(param_list) return_type
	    ...
	}
*/
type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rect struct {
	length, width float32
}

func (r Rect) Area() float32 {
	return r.length * r.width
}

//接口嵌套接口
/**
一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样
*/

/**
type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}
*/

//类型断言：如何检测和转换接口变量的类型
/**
一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型
即运行时在变量中存储的值的实际类型
在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型
常我们可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值：
v := varI.(T)       // unchecked type assertion

接口变量的类型也可以使用一种特殊形式的 switch 来检测：type-switch

switch t := areaIntf.(type) {
case *Square:
    fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
    fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
    fmt.Printf("nil value: nothing to check?\n")
default:
    fmt.Printf("Unexpected type %T\n", t)
}
*/

func main() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Println(areaIntf.Area())

	//多态
	r := Rect{5, 3}
	q := &Square{5}

	shapes := []Shaper{r, q}
	for n, _ := range shapes {
		fmt.Println(shapes[n])
		fmt.Println(shapes[n].Area())
	}
}
