package main

import "fmt"

// 结构体 定义在方法外面可以全局使用

type Student struct {
	id   int
	name string
	age  int
}

func main() {

	// 结构体名 定义结构体变量
	var s = Student{23, "kou", 22}
	fmt.Println(s)
}
