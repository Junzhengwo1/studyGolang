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
	// 先转成切片
	a := []rune(str)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c \n", a[i])
	}
	// 将string 转成数字
	b, _ := strconv.Atoi("12")
	i := 1
	b += i
	fmt.Println(b)

	sum := strconv.Itoa(i)
	fmt.Println("-----" + sum)
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

	//  ====================strings包 字符串长度
	println(len(k))

	//====================strings包 字符串遍历
	for i2 := range k {
		println(i2)
	}

	for _, v := range k {
		//fmt.Println(index)
		fmt.Printf("%c \n", v)
	}

	//====================strings包 统计字符串中 这个字符的个数

	kins := "chenkingabcdking"
	println(strings.Count(kins, "king"))

	//====================strings包 不区分大小比较
	cbc := "nihao"

	bcd := "NIHao1"

	println(strings.EqualFold(cbc, bcd))

}
