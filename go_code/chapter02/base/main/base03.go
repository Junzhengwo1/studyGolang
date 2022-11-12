package main

import "fmt"

func main() {
	var name string

	var age int
	var sal float64

	fmt.Println("请输入姓名：")
	s, _ := fmt.Scanln(&name)
	fmt.Printf("%T %v", s, name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
}
