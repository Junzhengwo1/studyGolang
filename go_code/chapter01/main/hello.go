package main

import "fmt"

//go语言中引入的包合变量没会编译不通过
func main() {
	var num = 10
	fmt.Println("king")
	fmt.Println(num)

}
