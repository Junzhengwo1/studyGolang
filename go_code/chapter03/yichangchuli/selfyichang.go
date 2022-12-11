package main

import (
	"errors"
	"fmt"
)

func main() {
	err := testErr(10, 0)
	if err != nil {
		fmt.Println("self Error Is:", err)
	}
	fmt.Println("正常执行下面的逻辑。。。")
}

func testErr(a int, b int) (err error) {
	if b == 0 {
		return errors.New("除数不能为零……")
	} else {
		result := a / b
		fmt.Println(result)
		return nil
	}
}
