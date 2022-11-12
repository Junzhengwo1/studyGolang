package main

import "fmt"

/**
 go语言的转义字符：
	1、\t:表示制表符：对其操作
	2、\n:换行
	3、\\:一个\
	4、\":一个"
	5、\r:一个回车
*/

/*
	定义全局变量

*/
//var userName  = "king"
//var age  = 20
//var height = 174.5

var (
	userName = "king"
	age2     = 20
	height   = 174.5
)

func main() {
	/*定义变量三种方式
	1、var a int 默认值0
	2、var num = 10  编译器 默认会认为是 int 类型
	3、省略 var 直接 name := "tom" ->(注意冒号不能省略)
	*/

	var a int
	a = 10
	name := "tom"
	/*
		一次性声明多个变量
		var n1,n2,n3 int  这三个都是int
		var n,age,b = 12.3,100,"wan"
		n,age,b := 12.3,100,"wan"
	*/
	var n1, n2, n3 int
	var n, age, b = 12.3, 100, "wan"
	fmt.Println(n, age, b)
	fmt.Println(n1, n2, n3)
	fmt.Println("a=", a)
	fmt.Println("name=", name)

	fmt.Println(userName, age2, height)
}
