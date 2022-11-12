package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/**
go语言中数据类型
todo go语言中 首字母的大小写来区别是否是私有的
1、基本数据类型 四个：
	1)数值  默认值  0 todo 保小保大原则
	2)字符 byte  默认值
	3)bool  默认值 false
	4)string 默认值 ""
2、复杂数据类型：
	1)pointer
	2)数组
	3)结构体
	4)管道
	5)函数
	6)切片
	7)接口
	8)map

*/
func main() {

	var i = 1
	fmt.Println(i)

	var a int8 = -128
	fmt.Println(a)
	fmt.Printf("%T %d \n", a, unsafe.Sizeof(a))

	//todo 浮点型
	var price float32 = 12.34
	fmt.Println(price)

	//todo 字符类型 一般字符就用byte来存；特殊的就用int
	var c byte = 'k'
	var c2 int = '北'
	fmt.Printf("c=%c %c \n", c, c2)

	//todo 字符串使用
	var str = "king"
	fmt.Println(len(str))

	var k = `var a int8 = -128
	fmt.Println(a)
	fmt.Printf("%T %d \n" ,a,unsafe.Sizeof(a))

	//浮点型
	var price float32 = 12.34
	fmt.Println(price)

	//字符类型 一般字符就用byte来存；特殊的就用int
	var c byte = 'k'
	var c2 int = '北'
	fmt.Printf("c=%c %c \n",c,c2)

	//字符串使用
	var str string = "king"
	fmt.Println(len(str))`
	fmt.Println(k)

	//todo 字符串拼接
	var w = "hello"
	var e = "world"

	var t = "king" + "prince"
	fmt.Println(w + e)
	fmt.Println(t)

	//todo 基本数据类型的相互转换
	var g = 100.00
	//把g转成float 类型
	f := int8(g)
	fmt.Println(f)

	//将int转成string
	var v = 100
	re := fmt.Sprintf("%d", v)
	fmt.Printf("%T %q", re, re)

	var v2 = 222
	str = strconv.FormatInt(int64(v2), 10) //十进制
	fmt.Printf("type %T str %q  \n", str, str)
	str = strconv.Itoa(v2)
	fmt.Println(str)

	//将string转成其他基本数据类型
	var s = "100"
	fmt.Printf("v = %v  %T \n", s, s)
	parseInt, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("type is %T \n", parseInt)

	var s2 = "100.2134"
	fmt.Printf("v = %v  %T \n", s2, s2)
	float, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("type is %T \n", float)

	//指针
	var inr = 520
	var ptr *int = &inr
	fmt.Println("int的地址：", &inr)
	fmt.Println("int的地址：", ptr)
	//取出指针的值
	fmt.Printf("ptr指向的值=%v", *ptr)

	var userName string = "king"
	fmt.Println(userName)

	//键盘输入语句
	fmt.Scanln()

}
