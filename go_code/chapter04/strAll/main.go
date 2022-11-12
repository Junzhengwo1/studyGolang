package main

import (
	"fmt"
	"strconv"
)

/**
字符串的常用函数
*/
func main() {

	var str = "abc美"
	fmt.Println(len(str))
	a := []rune(str)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c \n", a[i])
	}
	b, _ := strconv.Atoi("12")
	fmt.Println(b)

}
