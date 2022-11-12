package main

import (
	"fmt"
	"studyGolang/go_code/chapter03/fiebo"
)

func main() {

	//add:= pack.Add(1, 3)
	//fmt.Println(add)

	//test(4)

	//t(4)

	/**
	递归的案例
	*/
	fbo := fiebo.Fbo(5)
	fmt.Println(fbo)

	peach := fiebo.EatPeach(1)
	fmt.Println(peach)
}

/*
*
递归调用
*/
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

/*
*
递归调用
*/
func t(n int) {
	if n > 2 {
		n--
		t(n)
	} else {
		fmt.Println("n=", n)
	}

	/**
	  匿名函数
	*/

}
