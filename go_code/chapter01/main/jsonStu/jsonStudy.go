package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Json
func main() {

	var s = []int{1, 2, 3}
	fmt.Println(s)
	// JSON 序列化
	marshal, _ := json.Marshal(s)

	fmt.Println(string(marshal))
	fmt.Println(marshal)
	err := os.WriteFile("json2", marshal, 0666)
	if err != nil {
		return
	}

	var nameMap = map[string]string{"king": "kou"}
	bytes, _ := json.Marshal(nameMap)
	fmt.Println(nameMap)
	fmt.Println(string(bytes))
	_ = os.WriteFile("map", bytes, 0666)

	file, err := os.ReadFile("map")

	// 反序列化
	fmt.Println(string(file))
	var info = make(map[string]string)
	_ = json.Unmarshal(file, &info)
	fmt.Println(info)
}
