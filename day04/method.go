package main

import (
	"fmt"
	"math"
)

/**
方法
*/

/**
Go 方法是作用在接收者（receiver）上的一个函数，
接收者是某种类型的变量。因此方法是一种特殊类型的函数。
接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，
可以是 int、bool、string 或数组的别名类型。
但是接收者不能是一个接口类型，因为接口是一个抽象定义，
但是方法却是具体实现；如果这样做会引发一个编译错误：invalid receiver type…。
最后接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。

*/

type TwoInts struct {
	a int
	b int
}

// 非结构体类型上方法例子
type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

//指针作为接受者
/**
鉴于性能的原因，recv 最常见的是一个指向 receiver_type 的指针
如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法
*/
type D struct {
	thing int
}

func (d *D) change() {
	d.thing = 1
}
func (d D) write() string {
	return fmt.Sprint(d)
}

//内嵌类型的方法和继承
/**
当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌
这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型
这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果
*/

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

/*
*
内嵌将一个已存在类型的字段和方法注入到了另一个类型里：匿名字段上的方法 “晋升” 成为了外层类型的方法。
当然类型可以有只作用于本身实例而不作用于内嵌 “父” 类型上的方法，可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法
*/
func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100.
}

//如何在类型中嵌入功能
/**
主要有两种方法来实现在类型中嵌入功能：
A：聚合（或组合）：包含一个所需功能类型的具名字段。
B：内嵌：内嵌（匿名地）所需功能类型
*/

type Log struct {
	msg string
}

type Customer1 struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}
func (l *Log) String() string {
	return l.msg
}
func (c *Customer1) Log() *Log {
	return c.log
}

// 方法A
func A() {
	c := new(Customer1)
	c.Name = "Haoge name"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	c = &Customer1{"Haoge name", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}

type Customer2 struct {
	Name string
	Log
}

func (c *Customer2) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}

// 方法B
func B() {
	c := &Customer2{"Haoge name", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}

//如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：
//runtime.SetFinalizer(obj, func(obj *typeObj))

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	fmt.Println(IntVector{1, 2, 3}.Sum())

	//指针方法和值方法都可以在指针或非指针上被调用
	var d1 D
	d1.change()
	fmt.Println(d1.write())

	d2 := new(D)
	d2.change()
	fmt.Println(d2.write())

	n := &NamedPoint{Point{3, 4}, "haoge"}
	fmt.Println(n.Abs())

	A()
	B()
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
