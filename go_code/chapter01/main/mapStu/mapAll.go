package main

import "fmt"

func main() {

	var m = make(map[string]int)
	m["king"] = 23
	m["que"] = 24
	fmt.Println(m["king"])
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

}
