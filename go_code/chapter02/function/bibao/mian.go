package main

/*
*
闭包：返回的匿名函数+匿名函数以外的变量
*/
func main() {
	f := getSum()
	println(f(1))
	println(f(2))
}

func getSum() func(int) int {
	var sum int = 10
	return func(i int) int {
		sum = sum + i
		return sum
	}
}
