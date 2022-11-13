package fiebo

func Fbo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fbo(n-1) + Fbo(n-2)
	}
}

/**
猴子吃桃子问题
*/

func EatPeach(n int) int {
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (EatPeach(n+1) + 1) * 2
	}
}
