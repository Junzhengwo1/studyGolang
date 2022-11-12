package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
循环操作
*/

/**
while
和
do
while
go语言中没有while do while
*/
func main() {

	//for i := 0; i < 10; i++ {
	//	fmt.Println("hello world")
	//}

	//统计三个班的成绩情况，每个班的有五个同学，去求每个班的平均成绩
	//var total =0.0
	//for j := 1; j <=3 ; j++ {
	//	var sum  = 0.0
	//	for i := 1; i <= 5; i++ {
	//		var score float64
	//		fmt.Printf("请输入第%d班 第%d个学生的出成绩 \n",j,i)
	//		_, _ = fmt.Scanln(&score)
	//		//成绩求和
	//		sum += score
	//	}
	//	fmt.Printf("第%d班的平均分是 %v \n",j,sum/5)
	//	total+=sum
	//}
	//fmt.Printf("总成绩是%v,各个班的平均分是%v \n",total,total/3)

	/**
	打印金字塔案例
	*/
	//var total int = 9
	//for i := 1; i <= total; i++ {
	//	//在打印星号前先打印空格
	//	for k := 1; k <= total-i; k++ {
	//		fmt.Print(" ")
	//	}
	//	for j := 1; j <= 2*i-1; j++ {
	//		if j==1 || j==2*i-1 || i==total{
	//			fmt.Print("*")
	//		}else {
	//			fmt.Print(" ")
	//		}
	//	}
	//	fmt.Println(" ")
	//}

	/**
	打印九九乘法表
	*/
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%v * %v = %v \t", j, i, j*i)
	//	}
	//	fmt.Println(" ")
	//}

	/**
	  break 终止循环
	  continue 结束本次循环
	*/
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println(count)
}
