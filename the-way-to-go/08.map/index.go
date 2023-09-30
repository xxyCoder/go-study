package main

import "fmt"

func main() {
	var mapLit map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	// 引用类型，内存用make分配
	mapCreated := make(map[string]float32, 5)

	fmt.Println(mapLit, mapCreated)

	mapCreated["k1"] = 4.5
	// 测试键值对是否存在
	if v, ok := mapCreated["k1"]; ok {
		fmt.Println("存在", v)
	}
	// 删除键值对，如果键不存在，也不会触发错误
	delete(mapLit, "one")
}
