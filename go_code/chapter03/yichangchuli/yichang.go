package main

import (
	"fmt"
)

/*
*
defer 加 recover 机制处理异常
*/
func main() {
	test()
	fmt.Println("上面的除法操作执行成功。。。")
	fmt.Println("正常执行下面的逻辑。。。")

}

func test() {
	defer func() {
		// 调用recover内置函数
		err := recover()
		if err != nil {
			fmt.Println("错误已经捕获……")
			fmt.Println("errIs……", err)
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
