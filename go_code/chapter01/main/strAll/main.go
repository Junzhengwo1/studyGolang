package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
字符串的常用函数
*/
func main() {

	var str string = "abc美king"
	fmt.Println(len(str))
	a := []rune(str)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c \n", a[i])
	}
	// 将string 转成数字
	b, _ := strconv.Atoi("12")
	i := 1
	b += i
	fmt.Println(b)

	// 切片 左闭右开
	s := str[0:8]
	println(s) //abc美ki
	// 字符串拼接 +
	king := "kou"

	println(str + king)

	//  ====================strings包 一系列的字符串转换过程
	k := "chen"
	// s := "yu"
	replace := strings.Replace(k, str, king, 2)
	println(replace)
}
