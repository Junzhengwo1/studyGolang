package main

import "fmt"

func main() {

	var m = make(map[string]int)
	m["king"] = 23
	m["que"] = 24
	m["aa"] = 50
	fmt.Println(m["king"])
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
	// 删除map中的元素
	delete(m, "aa")
	fmt.Println(m)
}
