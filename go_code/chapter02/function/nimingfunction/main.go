package main

import "fmt"

/*
*
全局匿名函数的使用
*/
var (
	f = func(n1 int, n2 int) int { return n1 + n2 }
)

/*
*
匿名函数使用
*/
func main() {

	// 定义的同时调用
	res := func(n1 int, n2 int) int { return n1 + n2 }(1, 2)

	fmt.Println(res)

	fmt.Println(f(1, 3))
}
