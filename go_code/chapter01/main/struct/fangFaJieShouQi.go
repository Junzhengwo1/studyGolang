package main

import "fmt"

type Pes struct {
	id   int
	name string
	age  int
}

// 方法接收器
// 表达了该方法唯一属于 Pes
func (s Pes) read() {
	fmt.Println(s.name + "学生正在读书……")

}

// 继承 通过匿名字段实现继承

type p struct {
	Pes
	addr string
}

func main() {

	p := Pes{22, "que", 23}
	fmt.Println(p.name)
	p.read()

}
