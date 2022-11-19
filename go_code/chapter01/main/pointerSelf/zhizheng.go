package main

import "fmt"

/*
*
指针 理解 在 go 语言中 只限制 取地址 和 取值
*/
func main() {

	//指针
	var inr int = 520
	var ptr *int = &inr
	fmt.Println("int的地址：", &inr)
	fmt.Println("int的地址：", ptr)
	//取出指针的值
	fmt.Printf("ptr指向的值=%v", *ptr)
	inr = 100
	fmt.Printf("ptr指向的值=%v", *ptr)

}
