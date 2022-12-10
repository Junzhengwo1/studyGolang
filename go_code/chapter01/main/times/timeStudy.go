package main

import (
	"fmt"
	"time"
)

/*
*
日期相关的
*/
func main() {

	now := time.Now()
	fmt.Printf("%v ~~~~  %T \n", now, now)

	println(now.Year())

}
