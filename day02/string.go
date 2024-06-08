package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings" //内置多种处理字符串的函数
)

/**
字符串
Go 中的字符串根据需要占用 1 至 4 个字节
*/

/*
*
Go 这样做的好处是不仅减少了内存和硬盘空间占用，
同时也不用像其它语言那样需要对使用 UTF-8 字符集的文本进行编码和解码。
*/
func main() {
	str := "Hello,World"
	n := len(str)
	fmt.Println(n)
	fmt.Println(str)

	f := strings.HasPrefix(str, "hello")
	fmt.Println(f)

	f = strings.HasSuffix(str, "world")
	fmt.Println(f)

	f = strings.Contains(str, "hel")
	fmt.Println(f)

	f = strings.Contains(str, "hei")
	fmt.Println(f)

	//获取子字符串在父字符串中出现的,返回第一次出现的索引
	index := strings.Index(str, "l")
	fmt.Println(index)

	//返回最后一次出现的索引
	index = strings.LastIndex(str, "l")
	fmt.Println(index)

	/**
	Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，
	并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new
	*/
	replace := strings.Replace(str, "l", "i", 1)
	fmt.Println(replace)
	//字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容
	fmt.Println(str)

	replace = strings.ReplaceAll(str, "l", "L")
	fmt.Println(replace)

	//统计字符串出现的次数
	count := strings.Count(str, "l")
	fmt.Println(count)

	//重复字符串,Repeat 用于重复 count 次字符串 s 并返回一个新的字符串
	repeat := strings.Repeat(str, count)
	fmt.Println(repeat)

	//修改字符串大小写
	//小写
	lower := strings.ToLower(str)
	//大写
	upper := strings.ToUpper(str)

	fmt.Println(lower)
	fmt.Println(upper)

	//修剪字符串 使用 strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号；如果你想要剔除指定字符，则可以使用 strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉
	str2 := "   !!! Hello,World   !!!   "
	space := strings.TrimSpace(str2)
	fmt.Println(space)

	trim := strings.Trim(space, "!")
	fmt.Println(trim)

	//分割字符串
	split := strings.Fields(str2) //使用空白作为分隔符
	fmt.Println(split)
	fmt.Println(len(split))

	for _, v := range split {
		fmt.Println(v + ",")
	}

	//strings.Split(s, sep) 自定义分割符号对字符串分割，返回 slice
	split = strings.Split(str2, "l")
	fmt.Println(split)
	fmt.Println(len(split))

	for _, v := range split {
		fmt.Println(v + ",")
	}
	//拼接slice，效率优于+
	join := strings.Join(split, "")
	fmt.Println(join)

	join = strings.Join(split, ";")
	fmt.Println(join)

	/**
	字符串与其他类型转换
	与字符串相关的类型转换都是通过 strconv 包实现的。
	任何类型 T 转换为字符串总是成功的
	*/
	i := 10
	itoa := strconv.Itoa(i)
	fmt.Println(reflect.TypeOf(itoa))
	fmt.Println(itoa)

	//将字符串转换为其它类型 tp 并不总是可能的，可能会在运行时抛出错误
	atoi, _ := strconv.Atoi(itoa)
	fmt.Println(atoi)
	fmt.Println(reflect.TypeOf(atoi))

	float, _ := strconv.ParseFloat(itoa, 64)
	fmt.Println(reflect.TypeOf(float))
	fmt.Println(float)

}
